import glob
from pathlib import Path
from string import Template
import json
from dysvm_server import get_module_dict, dyslang, build_sandbox


msg_template = Template(
    """
// Keeper: $mod_keeper
// Types: $mod_types
// $keeper_import
func (rpcservice *RpcService) $function_name(_ *http.Request, msg *$mod_types.$req_type, response *$mod_types.$resp_type) (err error) {
	//handler := $mod_keeper.NewMsgServerImpl(rpcservice.k.$mod_keeper).$keeper_function
    $handler
	//
	defer func() {
		if r := recover(); r != nil {
			err = sdkerrors.Wrapf(types.RpcError, "RPC ERROR: %+v", msg)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}
	if !msg.GetSigners()[0].Equals(rpcservice.ScriptAddress) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", rpcservice.ScriptAddress)
	}

	sdkCtx := sdk.UnwrapSDKContext(rpcservice.ctx)

	cachedCtx, Write := sdkCtx.CacheContext()

	r, err := handler(sdk.WrapSDKContext(cachedCtx), msg)
	if err != nil {
		return err
	}
	Write()
	*response = *r
	return nil
}
"""
)

query_template = Template(
    """
// Keeper: $mod_keeper
// Types: $mod_types
// $keeper_import
func (rpcservice *RpcService) $function_name(_ *http.Request, msg *$mod_types.$req_type, response *$mod_types.$resp_type) (err error) {
    $handler
	if err != nil {
		return err
	}
	*response = *r
	return nil
}
"""
)


def to_camel_case(text):
    s = text.replace("-", " ").replace("_", " ")
    s = s.split()
    if len(text) == 0:
        return text
    ret = s[0] + "".join(i.capitalize() for i in s[1:])
    return ret


types = set()
keepers = set()
code = []
schemas = {}
for file_path in sorted(glob.glob("vue/src/**/protomodule.json", recursive=True)):
    if "crisis" in file_path:
        continue
    data = json.load(open(file_path))
    for services in data["Pkg"]["Services"]:
        for d in services["RPCFuncs"]:
            if services["Name"] == "Msg":
                command = (
                    f'{ data["Pkg"]["Name"] }/send{ services["Name"] }{ d["Name"] }'
                )

            if services["Name"] == "Query":
                command = f'{ data["Pkg"]["Name"] }/{ services["Name"] }{ d["Name"] }'

            req_schema = json.load(
                open(
                    str(
                        Path(file_path).parent
                        / Path("module/types")
                        / Path(f"{d['RequestType']}.json")
                    )
                )
            )

            for k, v in req_schema["definitions"].items():
                if "properties" in v:
                    for p in list(v["properties"]):
                        cc = to_camel_case(p)
                        if p != cc:
                            v["properties"][cc] = v["properties"].pop(p)

            res_schema = json.load(
                open(
                    str(
                        Path(file_path).parent
                        / Path("module/types")
                        / Path(f"{d['ReturnsType']}.json")
                    )
                )
            )
            schemas[command] = {
                "module_name": data["Pkg"]["Name"],
                "name": d["Name"],
                "service_name": services["Name"],
                "request_schema": req_schema,
                "resp_schema": res_schema,
                "http_rules": d["HTTPRules"],
            }
            # print(d)
            service_name = services["Name"]
            if service_name == "Msg":
                template = msg_template
                function_name = f"{data['Pkg']['Name']}sendMsg{d['Name']}".replace(
                    ".", ""
                ).capitalize()
                handler_template = Template(
                    "handler := $mod_keeper.NewMsgServerImpl(rpcservice.k.$mod_keeper).$keeper_function"
                )
            elif service_name == "Query":
                template = query_template
                handler_template = Template(
                    """r, err := rpcservice.k.$mod_keeper.$keeper_function(rpcservice.ctx, msg)"""
                )
                function_name = f"{data['Pkg']['Name']}Query{d['Name']}".replace(
                    ".", ""
                ).capitalize()
            else:
                raise Exception(f"oh no: {service_name}")
            mod_types = f"{data['Pkg']['Name']}types".lower().replace(".", "")
            mod_keeper = f"{data['Pkg']['Name']}keeper".lower().replace(".", "")

            if "dyson" in mod_keeper:
                keeper_import = data["Pkg"]["GoImportName"].replace("/types", "")
                if service_name == "Msg":
                    handler_template = Template(
                        "handler := rpcservice.m.$keeper_function"
                    )
                elif service_name == "Query":
                    handler_template = Template(
                        """r, err := rpcservice.k.$keeper_function(rpcservice.ctx, msg)"""
                    )
                types.add(f'{mod_types} "{data["Pkg"]["GoImportName"]}"')
            else:
                if "vesting" in mod_keeper:
                    keeper_import = data["Pkg"]["GoImportName"].replace("/types", "")
                    if service_name == "Msg":
                        handler_template = Template(
                            "handler := $mod_keeper.NewMsgServerImpl(rpcservice.k.accountKeeper, rpcservice.k.bankkeeper).$keeper_function"
                        )
                elif "ibcapplication" in mod_keeper:
                    # none because it uses the one on the keeper
                    keeper_import = None
                    if service_name == "Msg":
                        handler_template = Template(
                            "handler := rpcservice.k.$mod_keeper.$keeper_function"
                        )
                elif ("staking" in mod_keeper) and (service_name == "Query"):
                    mod_keeper = f"{data['Pkg']['Name']}querier".lower().replace(
                        ".", ""
                    )
                    keeper_import = None
                elif "authz" in mod_keeper:
                    # none because it uses the one on the keeper
                    keeper_import = None
                    if service_name == "Msg":
                        handler_template = Template(
                            "handler := rpcservice.k.$mod_keeper.$keeper_function"
                        )
                elif "feegrant" in mod_keeper:
                    keeper_import = data["Pkg"]["GoImportName"] + "/keeper"

                else:
                    keeper_import = data["Pkg"]["GoImportName"].replace(
                        "/types", "/keeper"
                    )
                if keeper_import:
                    keepers.add(f'{mod_keeper} "{keeper_import}"')
                types.add(f'{mod_types} "{data["Pkg"]["GoImportName"]}"')

            template_data = dict(
                function_name=function_name,
                mod_types=mod_types,
                req_type=d["RequestType"],
                resp_type=d["ReturnsType"],
                mod_keeper=mod_keeper,
                keeper_function=d["Name"],
                keeper_import=keeper_import,
            )
            handler = handler_template.safe_substitute(**template_data)
            template_data["handler"] = handler

            code.append(template.safe_substitute(**template_data))


imports = "\n".join(types | keepers)

funcs = "\n".join(code)


out = f"""package keeper

import (
    "net/http"
    sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/org/dyson/x/dyson/types"
    { imports }
)

{ funcs }
"""


with open("x/dyson/keeper/rpcserver.go", "w") as f:
    f.write(out)

# override schema to use textareas
schemas["dyson/sendMsgCreateScript"]["request_schema"]["definitions"][
    "MsgCreateScript"
]["properties"]["code"]["format"] = "textarea"

schemas["dyson/sendMsgUpdateScript"]["request_schema"]["definitions"][
    "MsgUpdateScript"
]["properties"]["code"]["format"] = "textarea"

schemas["dyson/QueryQueryScript"]["request_schema"]["definitions"]["MsgRun"][
    "properties"
]["extraLines"]["format"] = "textarea"

schemas["dyson/QueryQueryScript"]["resp_schema"]["definitions"]["MsgRunResponse"][
    "properties"
]["response"]["format"] = "textarea"

schemas["dyson/sendMsgRun"]["request_schema"]["definitions"]["MsgRun"]["properties"][
    "extraLines"
]["format"] = "textarea"

schemas["dyson/sendMsgRun"]["resp_schema"]["definitions"]["MsgRunResponse"][
    "properties"
]["response"]["format"] = "textarea"

schemas["dyson/sendMsgUpdateStorage"]["request_schema"]["definitions"][
    "MsgUpdateStorage"
]["properties"]["data"]["format"] = "textarea"

schemas["dyson/sendMsgCreateStorage"]["request_schema"]["definitions"][
    "MsgCreateStorage"
]["properties"]["data"]["format"] = "textarea"


with open("vue/src/views/command_schema.json", "w") as f:
    json.dump(schemas, f, indent=2)

# Generate docs
dyslang_modules = get_module_dict()
for mod_key, mod in dyslang_modules.items():
    for imp_key, imp in mod.items():
        mod[imp_key] = imp.__doc__

with open("vue/src/views/dyslang_modules.json", "w") as f:
    json.dump(dyslang_modules, f, indent=2)


sandbox = build_sandbox("port", "creator", "address", "amount", "block_info")
print(sandbox.scope)
print(sandbox.modules)
