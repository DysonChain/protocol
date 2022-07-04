// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUpdateName } from "./types/names/tx";
import { MsgDeleteName } from "./types/names/tx";
import { MsgCreateName } from "./types/names/tx";
import { MsgRegister } from "./types/names/tx";
const types = [
    ["/names.MsgUpdateName", MsgUpdateName],
    ["/names.MsgDeleteName", MsgDeleteName],
    ["/names.MsgCreateName", MsgCreateName],
    ["/names.MsgRegister", MsgRegister],
];
export const MissingWalletError = new Error("wallet is required");
export const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    let client;
    if (addr) {
        client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    }
    else {
        client = await SigningStargateClient.offline(wallet, { registry });
    }
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgUpdateName: (data) => ({ typeUrl: "/names.MsgUpdateName", value: MsgUpdateName.fromPartial(data) }),
        msgDeleteName: (data) => ({ typeUrl: "/names.MsgDeleteName", value: MsgDeleteName.fromPartial(data) }),
        msgCreateName: (data) => ({ typeUrl: "/names.MsgCreateName", value: MsgCreateName.fromPartial(data) }),
        msgRegister: (data) => ({ typeUrl: "/names.MsgRegister", value: MsgRegister.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
