package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	cosmosauthv1beta1types "github.com/cosmos/cosmos-sdk/x/auth/types"
	cosmosauthzv1beta1types "github.com/cosmos/cosmos-sdk/x/authz"
	cosmosbankv1beta1keeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	cosmosbankv1beta1types "github.com/cosmos/cosmos-sdk/x/bank/types"
	cosmosdistributionv1beta1keeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	cosmosdistributionv1beta1types "github.com/cosmos/cosmos-sdk/x/distribution/types"
	cosmosfeegrantv1beta1types "github.com/cosmos/cosmos-sdk/x/feegrant"
	cosmosfeegrantv1beta1keeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	cosmosgovv1keeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	cosmosgovv1types "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	cosmosgroupv1types "github.com/cosmos/cosmos-sdk/x/group"
	cosmosmintv1beta1types "github.com/cosmos/cosmos-sdk/x/mint/types"
	cosmosslashingv1beta1keeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	cosmosslashingv1beta1types "github.com/cosmos/cosmos-sdk/x/slashing/types"
	cosmosstakingv1beta1keeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	cosmosstakingv1beta1types "github.com/cosmos/cosmos-sdk/x/staking/types"
	cosmosupgradev1beta1keeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	cosmosupgradev1beta1types "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibcapplicationstransferv1types "github.com/cosmos/ibc-go/v5/modules/apps/transfer/types"
	"github.com/org/dyson/x/dyson/types"
	dysontypes "github.com/org/dyson/x/dyson/types"
	nameskeeper "github.com/org/dyson/x/names/keeper"
	namestypes "github.com/org/dyson/x/names/types"
	"net/http"
)

// Keeper: cosmosauthv1beta1keeper
// Types: cosmosauthv1beta1types
// github.com/cosmos/cosmos-sdk/x/auth/keeper
func (rpcservice *RpcService) Cosmosauthv1beta1queryaccounts(_ *http.Request, msg *cosmosauthv1beta1types.QueryAccountsRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosauthv1beta1keeper.Accounts(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosauthv1beta1keeper
// Types: cosmosauthv1beta1types
// github.com/cosmos/cosmos-sdk/x/auth/keeper
func (rpcservice *RpcService) Cosmosauthv1beta1queryaccount(_ *http.Request, msg *cosmosauthv1beta1types.QueryAccountRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosauthv1beta1keeper.Account(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosauthv1beta1keeper
// Types: cosmosauthv1beta1types
// github.com/cosmos/cosmos-sdk/x/auth/keeper
func (rpcservice *RpcService) Cosmosauthv1beta1queryaccountaddressbyid(_ *http.Request, msg *cosmosauthv1beta1types.QueryAccountAddressByIDRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosauthv1beta1keeper.AccountAddressByID(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosauthv1beta1keeper
// Types: cosmosauthv1beta1types
// github.com/cosmos/cosmos-sdk/x/auth/keeper
func (rpcservice *RpcService) Cosmosauthv1beta1queryparams(_ *http.Request, msg *cosmosauthv1beta1types.QueryParamsRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosauthv1beta1keeper.Params(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosauthv1beta1keeper
// Types: cosmosauthv1beta1types
// github.com/cosmos/cosmos-sdk/x/auth/keeper
func (rpcservice *RpcService) Cosmosauthv1beta1querymoduleaccounts(_ *http.Request, msg *cosmosauthv1beta1types.QueryModuleAccountsRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosauthv1beta1keeper.ModuleAccounts(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosauthv1beta1keeper
// Types: cosmosauthv1beta1types
// github.com/cosmos/cosmos-sdk/x/auth/keeper
func (rpcservice *RpcService) Cosmosauthv1beta1querymoduleaccountbyname(_ *http.Request, msg *cosmosauthv1beta1types.QueryModuleAccountByNameRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosauthv1beta1keeper.ModuleAccountByName(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosauthv1beta1keeper
// Types: cosmosauthv1beta1types
// github.com/cosmos/cosmos-sdk/x/auth/keeper
func (rpcservice *RpcService) Cosmosauthv1beta1querybech32prefix(_ *http.Request, msg *cosmosauthv1beta1types.Bech32PrefixRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosauthv1beta1keeper.Bech32Prefix(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosauthv1beta1keeper
// Types: cosmosauthv1beta1types
// github.com/cosmos/cosmos-sdk/x/auth/keeper
func (rpcservice *RpcService) Cosmosauthv1beta1queryaddressbytestostring(_ *http.Request, msg *cosmosauthv1beta1types.AddressBytesToStringRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosauthv1beta1keeper.AddressBytesToString(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosauthv1beta1keeper
// Types: cosmosauthv1beta1types
// github.com/cosmos/cosmos-sdk/x/auth/keeper
func (rpcservice *RpcService) Cosmosauthv1beta1queryaddressstringtobytes(_ *http.Request, msg *cosmosauthv1beta1types.AddressStringToBytesRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosauthv1beta1keeper.AddressStringToBytes(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosauthzv1beta1keeper
// Types: cosmosauthzv1beta1types
// None
func (rpcservice *RpcService) Cosmosauthzv1beta1querygrants(_ *http.Request, msg *cosmosauthzv1beta1types.QueryGrantsRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosauthzv1beta1keeper.Grants(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosauthzv1beta1keeper
// Types: cosmosauthzv1beta1types
// None
func (rpcservice *RpcService) Cosmosauthzv1beta1querygrantergrants(_ *http.Request, msg *cosmosauthzv1beta1types.QueryGranterGrantsRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosauthzv1beta1keeper.GranterGrants(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosauthzv1beta1keeper
// Types: cosmosauthzv1beta1types
// None
func (rpcservice *RpcService) Cosmosauthzv1beta1querygranteegrants(_ *http.Request, msg *cosmosauthzv1beta1types.QueryGranteeGrantsRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosauthzv1beta1keeper.GranteeGrants(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosauthzv1beta1keeper
// Types: cosmosauthzv1beta1types
// None
func (rpcservice *RpcService) Cosmosauthzv1beta1sendmsggrant(_ *http.Request, msg *cosmosauthzv1beta1types.MsgGrant, response *string) (err error) {
	//handler := cosmosauthzv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosauthzv1beta1keeper).Grant
	handler := rpcservice.k.cosmosauthzv1beta1keeper.Grant
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosauthzv1beta1keeper
// Types: cosmosauthzv1beta1types
// None
func (rpcservice *RpcService) Cosmosauthzv1beta1sendmsgexec(_ *http.Request, msg *cosmosauthzv1beta1types.MsgExec, response *string) (err error) {
	//handler := cosmosauthzv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosauthzv1beta1keeper).Exec
	handler := rpcservice.k.cosmosauthzv1beta1keeper.Exec
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosauthzv1beta1keeper
// Types: cosmosauthzv1beta1types
// None
func (rpcservice *RpcService) Cosmosauthzv1beta1sendmsgrevoke(_ *http.Request, msg *cosmosauthzv1beta1types.MsgRevoke, response *string) (err error) {
	//handler := cosmosauthzv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosauthzv1beta1keeper).Revoke
	handler := rpcservice.k.cosmosauthzv1beta1keeper.Revoke
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosbankv1beta1keeper
// Types: cosmosbankv1beta1types
// github.com/cosmos/cosmos-sdk/x/bank/keeper
func (rpcservice *RpcService) Cosmosbankv1beta1querybalance(_ *http.Request, msg *cosmosbankv1beta1types.QueryBalanceRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosbankv1beta1keeper.Balance(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosbankv1beta1keeper
// Types: cosmosbankv1beta1types
// github.com/cosmos/cosmos-sdk/x/bank/keeper
func (rpcservice *RpcService) Cosmosbankv1beta1queryallbalances(_ *http.Request, msg *cosmosbankv1beta1types.QueryAllBalancesRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosbankv1beta1keeper.AllBalances(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosbankv1beta1keeper
// Types: cosmosbankv1beta1types
// github.com/cosmos/cosmos-sdk/x/bank/keeper
func (rpcservice *RpcService) Cosmosbankv1beta1queryspendablebalances(_ *http.Request, msg *cosmosbankv1beta1types.QuerySpendableBalancesRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosbankv1beta1keeper.SpendableBalances(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosbankv1beta1keeper
// Types: cosmosbankv1beta1types
// github.com/cosmos/cosmos-sdk/x/bank/keeper
func (rpcservice *RpcService) Cosmosbankv1beta1querytotalsupply(_ *http.Request, msg *cosmosbankv1beta1types.QueryTotalSupplyRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosbankv1beta1keeper.TotalSupply(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosbankv1beta1keeper
// Types: cosmosbankv1beta1types
// github.com/cosmos/cosmos-sdk/x/bank/keeper
func (rpcservice *RpcService) Cosmosbankv1beta1querysupplyof(_ *http.Request, msg *cosmosbankv1beta1types.QuerySupplyOfRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosbankv1beta1keeper.SupplyOf(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosbankv1beta1keeper
// Types: cosmosbankv1beta1types
// github.com/cosmos/cosmos-sdk/x/bank/keeper
func (rpcservice *RpcService) Cosmosbankv1beta1queryparams(_ *http.Request, msg *cosmosbankv1beta1types.QueryParamsRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosbankv1beta1keeper.Params(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosbankv1beta1keeper
// Types: cosmosbankv1beta1types
// github.com/cosmos/cosmos-sdk/x/bank/keeper
func (rpcservice *RpcService) Cosmosbankv1beta1querydenommetadata(_ *http.Request, msg *cosmosbankv1beta1types.QueryDenomMetadataRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosbankv1beta1keeper.DenomMetadata(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosbankv1beta1keeper
// Types: cosmosbankv1beta1types
// github.com/cosmos/cosmos-sdk/x/bank/keeper
func (rpcservice *RpcService) Cosmosbankv1beta1querydenomsmetadata(_ *http.Request, msg *cosmosbankv1beta1types.QueryDenomsMetadataRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosbankv1beta1keeper.DenomsMetadata(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosbankv1beta1keeper
// Types: cosmosbankv1beta1types
// github.com/cosmos/cosmos-sdk/x/bank/keeper
func (rpcservice *RpcService) Cosmosbankv1beta1querydenomowners(_ *http.Request, msg *cosmosbankv1beta1types.QueryDenomOwnersRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosbankv1beta1keeper.DenomOwners(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosbankv1beta1keeper
// Types: cosmosbankv1beta1types
// github.com/cosmos/cosmos-sdk/x/bank/keeper
func (rpcservice *RpcService) Cosmosbankv1beta1sendmsgsend(_ *http.Request, msg *cosmosbankv1beta1types.MsgSend, response *string) (err error) {
	//handler := cosmosbankv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosbankv1beta1keeper).Send
	handler := cosmosbankv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosbankv1beta1keeper).Send
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosbankv1beta1keeper
// Types: cosmosbankv1beta1types
// github.com/cosmos/cosmos-sdk/x/bank/keeper
func (rpcservice *RpcService) Cosmosbankv1beta1sendmsgmultisend(_ *http.Request, msg *cosmosbankv1beta1types.MsgMultiSend, response *string) (err error) {
	//handler := cosmosbankv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosbankv1beta1keeper).MultiSend
	handler := cosmosbankv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosbankv1beta1keeper).MultiSend
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosdistributionv1beta1keeper
// Types: cosmosdistributionv1beta1types
// github.com/cosmos/cosmos-sdk/x/distribution/keeper
func (rpcservice *RpcService) Cosmosdistributionv1beta1queryparams(_ *http.Request, msg *cosmosdistributionv1beta1types.QueryParamsRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosdistributionv1beta1keeper.Params(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosdistributionv1beta1keeper
// Types: cosmosdistributionv1beta1types
// github.com/cosmos/cosmos-sdk/x/distribution/keeper
func (rpcservice *RpcService) Cosmosdistributionv1beta1queryvalidatoroutstandingrewards(_ *http.Request, msg *cosmosdistributionv1beta1types.QueryValidatorOutstandingRewardsRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosdistributionv1beta1keeper.ValidatorOutstandingRewards(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosdistributionv1beta1keeper
// Types: cosmosdistributionv1beta1types
// github.com/cosmos/cosmos-sdk/x/distribution/keeper
func (rpcservice *RpcService) Cosmosdistributionv1beta1queryvalidatorcommission(_ *http.Request, msg *cosmosdistributionv1beta1types.QueryValidatorCommissionRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosdistributionv1beta1keeper.ValidatorCommission(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosdistributionv1beta1keeper
// Types: cosmosdistributionv1beta1types
// github.com/cosmos/cosmos-sdk/x/distribution/keeper
func (rpcservice *RpcService) Cosmosdistributionv1beta1queryvalidatorslashes(_ *http.Request, msg *cosmosdistributionv1beta1types.QueryValidatorSlashesRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosdistributionv1beta1keeper.ValidatorSlashes(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosdistributionv1beta1keeper
// Types: cosmosdistributionv1beta1types
// github.com/cosmos/cosmos-sdk/x/distribution/keeper
func (rpcservice *RpcService) Cosmosdistributionv1beta1querydelegationrewards(_ *http.Request, msg *cosmosdistributionv1beta1types.QueryDelegationRewardsRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosdistributionv1beta1keeper.DelegationRewards(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosdistributionv1beta1keeper
// Types: cosmosdistributionv1beta1types
// github.com/cosmos/cosmos-sdk/x/distribution/keeper
func (rpcservice *RpcService) Cosmosdistributionv1beta1querydelegationtotalrewards(_ *http.Request, msg *cosmosdistributionv1beta1types.QueryDelegationTotalRewardsRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosdistributionv1beta1keeper.DelegationTotalRewards(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosdistributionv1beta1keeper
// Types: cosmosdistributionv1beta1types
// github.com/cosmos/cosmos-sdk/x/distribution/keeper
func (rpcservice *RpcService) Cosmosdistributionv1beta1querydelegatorvalidators(_ *http.Request, msg *cosmosdistributionv1beta1types.QueryDelegatorValidatorsRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosdistributionv1beta1keeper.DelegatorValidators(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosdistributionv1beta1keeper
// Types: cosmosdistributionv1beta1types
// github.com/cosmos/cosmos-sdk/x/distribution/keeper
func (rpcservice *RpcService) Cosmosdistributionv1beta1querydelegatorwithdrawaddress(_ *http.Request, msg *cosmosdistributionv1beta1types.QueryDelegatorWithdrawAddressRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosdistributionv1beta1keeper.DelegatorWithdrawAddress(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosdistributionv1beta1keeper
// Types: cosmosdistributionv1beta1types
// github.com/cosmos/cosmos-sdk/x/distribution/keeper
func (rpcservice *RpcService) Cosmosdistributionv1beta1querycommunitypool(_ *http.Request, msg *cosmosdistributionv1beta1types.QueryCommunityPoolRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosdistributionv1beta1keeper.CommunityPool(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosdistributionv1beta1keeper
// Types: cosmosdistributionv1beta1types
// github.com/cosmos/cosmos-sdk/x/distribution/keeper
func (rpcservice *RpcService) Cosmosdistributionv1beta1sendmsgsetwithdrawaddress(_ *http.Request, msg *cosmosdistributionv1beta1types.MsgSetWithdrawAddress, response *string) (err error) {
	//handler := cosmosdistributionv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosdistributionv1beta1keeper).SetWithdrawAddress
	handler := cosmosdistributionv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosdistributionv1beta1keeper).SetWithdrawAddress
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosdistributionv1beta1keeper
// Types: cosmosdistributionv1beta1types
// github.com/cosmos/cosmos-sdk/x/distribution/keeper
func (rpcservice *RpcService) Cosmosdistributionv1beta1sendmsgwithdrawdelegatorreward(_ *http.Request, msg *cosmosdistributionv1beta1types.MsgWithdrawDelegatorReward, response *string) (err error) {
	//handler := cosmosdistributionv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosdistributionv1beta1keeper).WithdrawDelegatorReward
	handler := cosmosdistributionv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosdistributionv1beta1keeper).WithdrawDelegatorReward
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosdistributionv1beta1keeper
// Types: cosmosdistributionv1beta1types
// github.com/cosmos/cosmos-sdk/x/distribution/keeper
func (rpcservice *RpcService) Cosmosdistributionv1beta1sendmsgwithdrawvalidatorcommission(_ *http.Request, msg *cosmosdistributionv1beta1types.MsgWithdrawValidatorCommission, response *string) (err error) {
	//handler := cosmosdistributionv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosdistributionv1beta1keeper).WithdrawValidatorCommission
	handler := cosmosdistributionv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosdistributionv1beta1keeper).WithdrawValidatorCommission
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosdistributionv1beta1keeper
// Types: cosmosdistributionv1beta1types
// github.com/cosmos/cosmos-sdk/x/distribution/keeper
func (rpcservice *RpcService) Cosmosdistributionv1beta1sendmsgfundcommunitypool(_ *http.Request, msg *cosmosdistributionv1beta1types.MsgFundCommunityPool, response *string) (err error) {
	//handler := cosmosdistributionv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosdistributionv1beta1keeper).FundCommunityPool
	handler := cosmosdistributionv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosdistributionv1beta1keeper).FundCommunityPool
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosfeegrantv1beta1keeper
// Types: cosmosfeegrantv1beta1types
// github.com/cosmos/cosmos-sdk/x/feegrant/keeper
func (rpcservice *RpcService) Cosmosfeegrantv1beta1queryallowance(_ *http.Request, msg *cosmosfeegrantv1beta1types.QueryAllowanceRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosfeegrantv1beta1keeper.Allowance(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosfeegrantv1beta1keeper
// Types: cosmosfeegrantv1beta1types
// github.com/cosmos/cosmos-sdk/x/feegrant/keeper
func (rpcservice *RpcService) Cosmosfeegrantv1beta1queryallowances(_ *http.Request, msg *cosmosfeegrantv1beta1types.QueryAllowancesRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosfeegrantv1beta1keeper.Allowances(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosfeegrantv1beta1keeper
// Types: cosmosfeegrantv1beta1types
// github.com/cosmos/cosmos-sdk/x/feegrant/keeper
func (rpcservice *RpcService) Cosmosfeegrantv1beta1queryallowancesbygranter(_ *http.Request, msg *cosmosfeegrantv1beta1types.QueryAllowancesByGranterRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosfeegrantv1beta1keeper.AllowancesByGranter(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosfeegrantv1beta1keeper
// Types: cosmosfeegrantv1beta1types
// github.com/cosmos/cosmos-sdk/x/feegrant/keeper
func (rpcservice *RpcService) Cosmosfeegrantv1beta1sendmsggrantallowance(_ *http.Request, msg *cosmosfeegrantv1beta1types.MsgGrantAllowance, response *string) (err error) {
	//handler := cosmosfeegrantv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosfeegrantv1beta1keeper).GrantAllowance
	handler := cosmosfeegrantv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosfeegrantv1beta1keeper).GrantAllowance
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosfeegrantv1beta1keeper
// Types: cosmosfeegrantv1beta1types
// github.com/cosmos/cosmos-sdk/x/feegrant/keeper
func (rpcservice *RpcService) Cosmosfeegrantv1beta1sendmsgrevokeallowance(_ *http.Request, msg *cosmosfeegrantv1beta1types.MsgRevokeAllowance, response *string) (err error) {
	//handler := cosmosfeegrantv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosfeegrantv1beta1keeper).RevokeAllowance
	handler := cosmosfeegrantv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosfeegrantv1beta1keeper).RevokeAllowance
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgovv1keeper
// Types: cosmosgovv1types
// github.com/cosmos/cosmos-sdk/x/feegrant/keeper
func (rpcservice *RpcService) Cosmosgovv1queryproposal(_ *http.Request, msg *cosmosgovv1types.QueryProposalRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosgovv1keeper.Proposal(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgovv1keeper
// Types: cosmosgovv1types
// github.com/cosmos/cosmos-sdk/x/feegrant/keeper
func (rpcservice *RpcService) Cosmosgovv1queryproposals(_ *http.Request, msg *cosmosgovv1types.QueryProposalsRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosgovv1keeper.Proposals(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgovv1keeper
// Types: cosmosgovv1types
// github.com/cosmos/cosmos-sdk/x/feegrant/keeper
func (rpcservice *RpcService) Cosmosgovv1queryvote(_ *http.Request, msg *cosmosgovv1types.QueryVoteRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosgovv1keeper.Vote(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgovv1keeper
// Types: cosmosgovv1types
// github.com/cosmos/cosmos-sdk/x/feegrant/keeper
func (rpcservice *RpcService) Cosmosgovv1queryvotes(_ *http.Request, msg *cosmosgovv1types.QueryVotesRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosgovv1keeper.Votes(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgovv1keeper
// Types: cosmosgovv1types
// github.com/cosmos/cosmos-sdk/x/feegrant/keeper
func (rpcservice *RpcService) Cosmosgovv1queryparams(_ *http.Request, msg *cosmosgovv1types.QueryParamsRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosgovv1keeper.Params(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgovv1keeper
// Types: cosmosgovv1types
// github.com/cosmos/cosmos-sdk/x/feegrant/keeper
func (rpcservice *RpcService) Cosmosgovv1querydeposit(_ *http.Request, msg *cosmosgovv1types.QueryDepositRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosgovv1keeper.Deposit(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgovv1keeper
// Types: cosmosgovv1types
// github.com/cosmos/cosmos-sdk/x/feegrant/keeper
func (rpcservice *RpcService) Cosmosgovv1querydeposits(_ *http.Request, msg *cosmosgovv1types.QueryDepositsRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosgovv1keeper.Deposits(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgovv1keeper
// Types: cosmosgovv1types
// github.com/cosmos/cosmos-sdk/x/feegrant/keeper
func (rpcservice *RpcService) Cosmosgovv1querytallyresult(_ *http.Request, msg *cosmosgovv1types.QueryTallyResultRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosgovv1keeper.TallyResult(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgovv1keeper
// Types: cosmosgovv1types
// github.com/cosmos/cosmos-sdk/x/gov/keeper
func (rpcservice *RpcService) Cosmosgovv1sendmsgsubmitproposal(_ *http.Request, msg *cosmosgovv1types.MsgSubmitProposal, response *string) (err error) {
	//handler := cosmosgovv1keeper.NewMsgServerImpl(rpcservice.k.cosmosgovv1keeper).SubmitProposal
	handler := cosmosgovv1keeper.NewMsgServerImpl(rpcservice.k.cosmosgovv1keeper).SubmitProposal
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgovv1keeper
// Types: cosmosgovv1types
// github.com/cosmos/cosmos-sdk/x/gov/keeper
func (rpcservice *RpcService) Cosmosgovv1sendmsgexeclegacycontent(_ *http.Request, msg *cosmosgovv1types.MsgExecLegacyContent, response *string) (err error) {
	//handler := cosmosgovv1keeper.NewMsgServerImpl(rpcservice.k.cosmosgovv1keeper).ExecLegacyContent
	handler := cosmosgovv1keeper.NewMsgServerImpl(rpcservice.k.cosmosgovv1keeper).ExecLegacyContent
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgovv1keeper
// Types: cosmosgovv1types
// github.com/cosmos/cosmos-sdk/x/gov/keeper
func (rpcservice *RpcService) Cosmosgovv1sendmsgvote(_ *http.Request, msg *cosmosgovv1types.MsgVote, response *string) (err error) {
	//handler := cosmosgovv1keeper.NewMsgServerImpl(rpcservice.k.cosmosgovv1keeper).Vote
	handler := cosmosgovv1keeper.NewMsgServerImpl(rpcservice.k.cosmosgovv1keeper).Vote
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgovv1keeper
// Types: cosmosgovv1types
// github.com/cosmos/cosmos-sdk/x/gov/keeper
func (rpcservice *RpcService) Cosmosgovv1sendmsgvoteweighted(_ *http.Request, msg *cosmosgovv1types.MsgVoteWeighted, response *string) (err error) {
	//handler := cosmosgovv1keeper.NewMsgServerImpl(rpcservice.k.cosmosgovv1keeper).VoteWeighted
	handler := cosmosgovv1keeper.NewMsgServerImpl(rpcservice.k.cosmosgovv1keeper).VoteWeighted
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgovv1keeper
// Types: cosmosgovv1types
// github.com/cosmos/cosmos-sdk/x/gov/keeper
func (rpcservice *RpcService) Cosmosgovv1sendmsgdeposit(_ *http.Request, msg *cosmosgovv1types.MsgDeposit, response *string) (err error) {
	//handler := cosmosgovv1keeper.NewMsgServerImpl(rpcservice.k.cosmosgovv1keeper).Deposit
	handler := cosmosgovv1keeper.NewMsgServerImpl(rpcservice.k.cosmosgovv1keeper).Deposit
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgroupv1keeper
// Types: cosmosgroupv1types
// None
func (rpcservice *RpcService) Cosmosgroupv1querygroupinfo(_ *http.Request, msg *cosmosgroupv1types.QueryGroupInfoRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosgroupv1keeper.GroupInfo(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgroupv1keeper
// Types: cosmosgroupv1types
// None
func (rpcservice *RpcService) Cosmosgroupv1querygrouppolicyinfo(_ *http.Request, msg *cosmosgroupv1types.QueryGroupPolicyInfoRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosgroupv1keeper.GroupPolicyInfo(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgroupv1keeper
// Types: cosmosgroupv1types
// None
func (rpcservice *RpcService) Cosmosgroupv1querygroupmembers(_ *http.Request, msg *cosmosgroupv1types.QueryGroupMembersRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosgroupv1keeper.GroupMembers(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgroupv1keeper
// Types: cosmosgroupv1types
// None
func (rpcservice *RpcService) Cosmosgroupv1querygroupsbyadmin(_ *http.Request, msg *cosmosgroupv1types.QueryGroupsByAdminRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosgroupv1keeper.GroupsByAdmin(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgroupv1keeper
// Types: cosmosgroupv1types
// None
func (rpcservice *RpcService) Cosmosgroupv1querygrouppoliciesbygroup(_ *http.Request, msg *cosmosgroupv1types.QueryGroupPoliciesByGroupRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosgroupv1keeper.GroupPoliciesByGroup(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgroupv1keeper
// Types: cosmosgroupv1types
// None
func (rpcservice *RpcService) Cosmosgroupv1querygrouppoliciesbyadmin(_ *http.Request, msg *cosmosgroupv1types.QueryGroupPoliciesByAdminRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosgroupv1keeper.GroupPoliciesByAdmin(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgroupv1keeper
// Types: cosmosgroupv1types
// None
func (rpcservice *RpcService) Cosmosgroupv1queryproposal(_ *http.Request, msg *cosmosgroupv1types.QueryProposalRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosgroupv1keeper.Proposal(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgroupv1keeper
// Types: cosmosgroupv1types
// None
func (rpcservice *RpcService) Cosmosgroupv1queryproposalsbygrouppolicy(_ *http.Request, msg *cosmosgroupv1types.QueryProposalsByGroupPolicyRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosgroupv1keeper.ProposalsByGroupPolicy(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgroupv1keeper
// Types: cosmosgroupv1types
// None
func (rpcservice *RpcService) Cosmosgroupv1queryvotebyproposalvoter(_ *http.Request, msg *cosmosgroupv1types.QueryVoteByProposalVoterRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosgroupv1keeper.VoteByProposalVoter(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgroupv1keeper
// Types: cosmosgroupv1types
// None
func (rpcservice *RpcService) Cosmosgroupv1queryvotesbyproposal(_ *http.Request, msg *cosmosgroupv1types.QueryVotesByProposalRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosgroupv1keeper.VotesByProposal(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgroupv1keeper
// Types: cosmosgroupv1types
// None
func (rpcservice *RpcService) Cosmosgroupv1queryvotesbyvoter(_ *http.Request, msg *cosmosgroupv1types.QueryVotesByVoterRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosgroupv1keeper.VotesByVoter(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgroupv1keeper
// Types: cosmosgroupv1types
// None
func (rpcservice *RpcService) Cosmosgroupv1querygroupsbymember(_ *http.Request, msg *cosmosgroupv1types.QueryGroupsByMemberRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosgroupv1keeper.GroupsByMember(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgroupv1keeper
// Types: cosmosgroupv1types
// None
func (rpcservice *RpcService) Cosmosgroupv1querytallyresult(_ *http.Request, msg *cosmosgroupv1types.QueryTallyResultRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosgroupv1keeper.TallyResult(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgroupv1keeper
// Types: cosmosgroupv1types
// None
func (rpcservice *RpcService) Cosmosgroupv1sendmsgcreategroup(_ *http.Request, msg *cosmosgroupv1types.MsgCreateGroup, response *string) (err error) {
	//handler := cosmosgroupv1keeper.NewMsgServerImpl(rpcservice.k.cosmosgroupv1keeper).CreateGroup
	handler := rpcservice.k.cosmosgroupv1keeper.CreateGroup
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgroupv1keeper
// Types: cosmosgroupv1types
// None
func (rpcservice *RpcService) Cosmosgroupv1sendmsgupdategroupmembers(_ *http.Request, msg *cosmosgroupv1types.MsgUpdateGroupMembers, response *string) (err error) {
	//handler := cosmosgroupv1keeper.NewMsgServerImpl(rpcservice.k.cosmosgroupv1keeper).UpdateGroupMembers
	handler := rpcservice.k.cosmosgroupv1keeper.UpdateGroupMembers
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgroupv1keeper
// Types: cosmosgroupv1types
// None
func (rpcservice *RpcService) Cosmosgroupv1sendmsgupdategroupadmin(_ *http.Request, msg *cosmosgroupv1types.MsgUpdateGroupAdmin, response *string) (err error) {
	//handler := cosmosgroupv1keeper.NewMsgServerImpl(rpcservice.k.cosmosgroupv1keeper).UpdateGroupAdmin
	handler := rpcservice.k.cosmosgroupv1keeper.UpdateGroupAdmin
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgroupv1keeper
// Types: cosmosgroupv1types
// None
func (rpcservice *RpcService) Cosmosgroupv1sendmsgupdategroupmetadata(_ *http.Request, msg *cosmosgroupv1types.MsgUpdateGroupMetadata, response *string) (err error) {
	//handler := cosmosgroupv1keeper.NewMsgServerImpl(rpcservice.k.cosmosgroupv1keeper).UpdateGroupMetadata
	handler := rpcservice.k.cosmosgroupv1keeper.UpdateGroupMetadata
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgroupv1keeper
// Types: cosmosgroupv1types
// None
func (rpcservice *RpcService) Cosmosgroupv1sendmsgcreategrouppolicy(_ *http.Request, msg *cosmosgroupv1types.MsgCreateGroupPolicy, response *string) (err error) {
	//handler := cosmosgroupv1keeper.NewMsgServerImpl(rpcservice.k.cosmosgroupv1keeper).CreateGroupPolicy
	handler := rpcservice.k.cosmosgroupv1keeper.CreateGroupPolicy
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgroupv1keeper
// Types: cosmosgroupv1types
// None
func (rpcservice *RpcService) Cosmosgroupv1sendmsgcreategroupwithpolicy(_ *http.Request, msg *cosmosgroupv1types.MsgCreateGroupWithPolicy, response *string) (err error) {
	//handler := cosmosgroupv1keeper.NewMsgServerImpl(rpcservice.k.cosmosgroupv1keeper).CreateGroupWithPolicy
	handler := rpcservice.k.cosmosgroupv1keeper.CreateGroupWithPolicy
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgroupv1keeper
// Types: cosmosgroupv1types
// None
func (rpcservice *RpcService) Cosmosgroupv1sendmsgupdategrouppolicyadmin(_ *http.Request, msg *cosmosgroupv1types.MsgUpdateGroupPolicyAdmin, response *string) (err error) {
	//handler := cosmosgroupv1keeper.NewMsgServerImpl(rpcservice.k.cosmosgroupv1keeper).UpdateGroupPolicyAdmin
	handler := rpcservice.k.cosmosgroupv1keeper.UpdateGroupPolicyAdmin
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgroupv1keeper
// Types: cosmosgroupv1types
// None
func (rpcservice *RpcService) Cosmosgroupv1sendmsgupdategrouppolicydecisionpolicy(_ *http.Request, msg *cosmosgroupv1types.MsgUpdateGroupPolicyDecisionPolicy, response *string) (err error) {
	//handler := cosmosgroupv1keeper.NewMsgServerImpl(rpcservice.k.cosmosgroupv1keeper).UpdateGroupPolicyDecisionPolicy
	handler := rpcservice.k.cosmosgroupv1keeper.UpdateGroupPolicyDecisionPolicy
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgroupv1keeper
// Types: cosmosgroupv1types
// None
func (rpcservice *RpcService) Cosmosgroupv1sendmsgupdategrouppolicymetadata(_ *http.Request, msg *cosmosgroupv1types.MsgUpdateGroupPolicyMetadata, response *string) (err error) {
	//handler := cosmosgroupv1keeper.NewMsgServerImpl(rpcservice.k.cosmosgroupv1keeper).UpdateGroupPolicyMetadata
	handler := rpcservice.k.cosmosgroupv1keeper.UpdateGroupPolicyMetadata
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgroupv1keeper
// Types: cosmosgroupv1types
// None
func (rpcservice *RpcService) Cosmosgroupv1sendmsgsubmitproposal(_ *http.Request, msg *cosmosgroupv1types.MsgSubmitProposal, response *string) (err error) {
	//handler := cosmosgroupv1keeper.NewMsgServerImpl(rpcservice.k.cosmosgroupv1keeper).SubmitProposal
	handler := rpcservice.k.cosmosgroupv1keeper.SubmitProposal
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgroupv1keeper
// Types: cosmosgroupv1types
// None
func (rpcservice *RpcService) Cosmosgroupv1sendmsgwithdrawproposal(_ *http.Request, msg *cosmosgroupv1types.MsgWithdrawProposal, response *string) (err error) {
	//handler := cosmosgroupv1keeper.NewMsgServerImpl(rpcservice.k.cosmosgroupv1keeper).WithdrawProposal
	handler := rpcservice.k.cosmosgroupv1keeper.WithdrawProposal
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgroupv1keeper
// Types: cosmosgroupv1types
// None
func (rpcservice *RpcService) Cosmosgroupv1sendmsgvote(_ *http.Request, msg *cosmosgroupv1types.MsgVote, response *string) (err error) {
	//handler := cosmosgroupv1keeper.NewMsgServerImpl(rpcservice.k.cosmosgroupv1keeper).Vote
	handler := rpcservice.k.cosmosgroupv1keeper.Vote
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgroupv1keeper
// Types: cosmosgroupv1types
// None
func (rpcservice *RpcService) Cosmosgroupv1sendmsgexec(_ *http.Request, msg *cosmosgroupv1types.MsgExec, response *string) (err error) {
	//handler := cosmosgroupv1keeper.NewMsgServerImpl(rpcservice.k.cosmosgroupv1keeper).Exec
	handler := rpcservice.k.cosmosgroupv1keeper.Exec
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosgroupv1keeper
// Types: cosmosgroupv1types
// None
func (rpcservice *RpcService) Cosmosgroupv1sendmsgleavegroup(_ *http.Request, msg *cosmosgroupv1types.MsgLeaveGroup, response *string) (err error) {
	//handler := cosmosgroupv1keeper.NewMsgServerImpl(rpcservice.k.cosmosgroupv1keeper).LeaveGroup
	handler := rpcservice.k.cosmosgroupv1keeper.LeaveGroup
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosmintv1beta1keeper
// Types: cosmosmintv1beta1types
// github.com/cosmos/cosmos-sdk/x/mint/keeper
func (rpcservice *RpcService) Cosmosmintv1beta1queryparams(_ *http.Request, msg *cosmosmintv1beta1types.QueryParamsRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosmintv1beta1keeper.Params(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosmintv1beta1keeper
// Types: cosmosmintv1beta1types
// github.com/cosmos/cosmos-sdk/x/mint/keeper
func (rpcservice *RpcService) Cosmosmintv1beta1queryinflation(_ *http.Request, msg *cosmosmintv1beta1types.QueryInflationRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosmintv1beta1keeper.Inflation(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosmintv1beta1keeper
// Types: cosmosmintv1beta1types
// github.com/cosmos/cosmos-sdk/x/mint/keeper
func (rpcservice *RpcService) Cosmosmintv1beta1queryannualprovisions(_ *http.Request, msg *cosmosmintv1beta1types.QueryAnnualProvisionsRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosmintv1beta1keeper.AnnualProvisions(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosslashingv1beta1keeper
// Types: cosmosslashingv1beta1types
// github.com/cosmos/cosmos-sdk/x/slashing/keeper
func (rpcservice *RpcService) Cosmosslashingv1beta1queryparams(_ *http.Request, msg *cosmosslashingv1beta1types.QueryParamsRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosslashingv1beta1keeper.Params(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosslashingv1beta1keeper
// Types: cosmosslashingv1beta1types
// github.com/cosmos/cosmos-sdk/x/slashing/keeper
func (rpcservice *RpcService) Cosmosslashingv1beta1querysigninginfo(_ *http.Request, msg *cosmosslashingv1beta1types.QuerySigningInfoRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosslashingv1beta1keeper.SigningInfo(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosslashingv1beta1keeper
// Types: cosmosslashingv1beta1types
// github.com/cosmos/cosmos-sdk/x/slashing/keeper
func (rpcservice *RpcService) Cosmosslashingv1beta1querysigninginfos(_ *http.Request, msg *cosmosslashingv1beta1types.QuerySigningInfosRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosslashingv1beta1keeper.SigningInfos(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosslashingv1beta1keeper
// Types: cosmosslashingv1beta1types
// github.com/cosmos/cosmos-sdk/x/slashing/keeper
func (rpcservice *RpcService) Cosmosslashingv1beta1sendmsgunjail(_ *http.Request, msg *cosmosslashingv1beta1types.MsgUnjail, response *string) (err error) {
	//handler := cosmosslashingv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosslashingv1beta1keeper).Unjail
	handler := cosmosslashingv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosslashingv1beta1keeper).Unjail
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosstakingv1beta1querier
// Types: cosmosstakingv1beta1types
// None
func (rpcservice *RpcService) Cosmosstakingv1beta1queryvalidators(_ *http.Request, msg *cosmosstakingv1beta1types.QueryValidatorsRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosstakingv1beta1querier.Validators(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosstakingv1beta1querier
// Types: cosmosstakingv1beta1types
// None
func (rpcservice *RpcService) Cosmosstakingv1beta1queryvalidator(_ *http.Request, msg *cosmosstakingv1beta1types.QueryValidatorRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosstakingv1beta1querier.Validator(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosstakingv1beta1querier
// Types: cosmosstakingv1beta1types
// None
func (rpcservice *RpcService) Cosmosstakingv1beta1queryvalidatordelegations(_ *http.Request, msg *cosmosstakingv1beta1types.QueryValidatorDelegationsRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosstakingv1beta1querier.ValidatorDelegations(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosstakingv1beta1querier
// Types: cosmosstakingv1beta1types
// None
func (rpcservice *RpcService) Cosmosstakingv1beta1queryvalidatorunbondingdelegations(_ *http.Request, msg *cosmosstakingv1beta1types.QueryValidatorUnbondingDelegationsRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosstakingv1beta1querier.ValidatorUnbondingDelegations(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosstakingv1beta1querier
// Types: cosmosstakingv1beta1types
// None
func (rpcservice *RpcService) Cosmosstakingv1beta1querydelegation(_ *http.Request, msg *cosmosstakingv1beta1types.QueryDelegationRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosstakingv1beta1querier.Delegation(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosstakingv1beta1querier
// Types: cosmosstakingv1beta1types
// None
func (rpcservice *RpcService) Cosmosstakingv1beta1queryunbondingdelegation(_ *http.Request, msg *cosmosstakingv1beta1types.QueryUnbondingDelegationRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosstakingv1beta1querier.UnbondingDelegation(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosstakingv1beta1querier
// Types: cosmosstakingv1beta1types
// None
func (rpcservice *RpcService) Cosmosstakingv1beta1querydelegatordelegations(_ *http.Request, msg *cosmosstakingv1beta1types.QueryDelegatorDelegationsRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosstakingv1beta1querier.DelegatorDelegations(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosstakingv1beta1querier
// Types: cosmosstakingv1beta1types
// None
func (rpcservice *RpcService) Cosmosstakingv1beta1querydelegatorunbondingdelegations(_ *http.Request, msg *cosmosstakingv1beta1types.QueryDelegatorUnbondingDelegationsRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosstakingv1beta1querier.DelegatorUnbondingDelegations(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosstakingv1beta1querier
// Types: cosmosstakingv1beta1types
// None
func (rpcservice *RpcService) Cosmosstakingv1beta1queryredelegations(_ *http.Request, msg *cosmosstakingv1beta1types.QueryRedelegationsRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosstakingv1beta1querier.Redelegations(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosstakingv1beta1querier
// Types: cosmosstakingv1beta1types
// None
func (rpcservice *RpcService) Cosmosstakingv1beta1querydelegatorvalidators(_ *http.Request, msg *cosmosstakingv1beta1types.QueryDelegatorValidatorsRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosstakingv1beta1querier.DelegatorValidators(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosstakingv1beta1querier
// Types: cosmosstakingv1beta1types
// None
func (rpcservice *RpcService) Cosmosstakingv1beta1querydelegatorvalidator(_ *http.Request, msg *cosmosstakingv1beta1types.QueryDelegatorValidatorRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosstakingv1beta1querier.DelegatorValidator(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosstakingv1beta1querier
// Types: cosmosstakingv1beta1types
// None
func (rpcservice *RpcService) Cosmosstakingv1beta1queryhistoricalinfo(_ *http.Request, msg *cosmosstakingv1beta1types.QueryHistoricalInfoRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosstakingv1beta1querier.HistoricalInfo(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosstakingv1beta1querier
// Types: cosmosstakingv1beta1types
// None
func (rpcservice *RpcService) Cosmosstakingv1beta1querypool(_ *http.Request, msg *cosmosstakingv1beta1types.QueryPoolRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosstakingv1beta1querier.Pool(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosstakingv1beta1querier
// Types: cosmosstakingv1beta1types
// None
func (rpcservice *RpcService) Cosmosstakingv1beta1queryparams(_ *http.Request, msg *cosmosstakingv1beta1types.QueryParamsRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosstakingv1beta1querier.Params(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosstakingv1beta1keeper
// Types: cosmosstakingv1beta1types
// github.com/cosmos/cosmos-sdk/x/staking/keeper
func (rpcservice *RpcService) Cosmosstakingv1beta1sendmsgcreatevalidator(_ *http.Request, msg *cosmosstakingv1beta1types.MsgCreateValidator, response *string) (err error) {
	//handler := cosmosstakingv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosstakingv1beta1keeper).CreateValidator
	handler := cosmosstakingv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosstakingv1beta1keeper).CreateValidator
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosstakingv1beta1keeper
// Types: cosmosstakingv1beta1types
// github.com/cosmos/cosmos-sdk/x/staking/keeper
func (rpcservice *RpcService) Cosmosstakingv1beta1sendmsgeditvalidator(_ *http.Request, msg *cosmosstakingv1beta1types.MsgEditValidator, response *string) (err error) {
	//handler := cosmosstakingv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosstakingv1beta1keeper).EditValidator
	handler := cosmosstakingv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosstakingv1beta1keeper).EditValidator
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosstakingv1beta1keeper
// Types: cosmosstakingv1beta1types
// github.com/cosmos/cosmos-sdk/x/staking/keeper
func (rpcservice *RpcService) Cosmosstakingv1beta1sendmsgdelegate(_ *http.Request, msg *cosmosstakingv1beta1types.MsgDelegate, response *string) (err error) {
	//handler := cosmosstakingv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosstakingv1beta1keeper).Delegate
	handler := cosmosstakingv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosstakingv1beta1keeper).Delegate
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosstakingv1beta1keeper
// Types: cosmosstakingv1beta1types
// github.com/cosmos/cosmos-sdk/x/staking/keeper
func (rpcservice *RpcService) Cosmosstakingv1beta1sendmsgbeginredelegate(_ *http.Request, msg *cosmosstakingv1beta1types.MsgBeginRedelegate, response *string) (err error) {
	//handler := cosmosstakingv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosstakingv1beta1keeper).BeginRedelegate
	handler := cosmosstakingv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosstakingv1beta1keeper).BeginRedelegate
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosstakingv1beta1keeper
// Types: cosmosstakingv1beta1types
// github.com/cosmos/cosmos-sdk/x/staking/keeper
func (rpcservice *RpcService) Cosmosstakingv1beta1sendmsgundelegate(_ *http.Request, msg *cosmosstakingv1beta1types.MsgUndelegate, response *string) (err error) {
	//handler := cosmosstakingv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosstakingv1beta1keeper).Undelegate
	handler := cosmosstakingv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosstakingv1beta1keeper).Undelegate
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosstakingv1beta1keeper
// Types: cosmosstakingv1beta1types
// github.com/cosmos/cosmos-sdk/x/staking/keeper
func (rpcservice *RpcService) Cosmosstakingv1beta1sendmsgcancelunbondingdelegation(_ *http.Request, msg *cosmosstakingv1beta1types.MsgCancelUnbondingDelegation, response *string) (err error) {
	//handler := cosmosstakingv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosstakingv1beta1keeper).CancelUnbondingDelegation
	handler := cosmosstakingv1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosstakingv1beta1keeper).CancelUnbondingDelegation
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosupgradev1beta1keeper
// Types: cosmosupgradev1beta1types
// github.com/cosmos/cosmos-sdk/x/upgrade/keeper
func (rpcservice *RpcService) Cosmosupgradev1beta1querycurrentplan(_ *http.Request, msg *cosmosupgradev1beta1types.QueryCurrentPlanRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosupgradev1beta1keeper.CurrentPlan(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosupgradev1beta1keeper
// Types: cosmosupgradev1beta1types
// github.com/cosmos/cosmos-sdk/x/upgrade/keeper
func (rpcservice *RpcService) Cosmosupgradev1beta1queryappliedplan(_ *http.Request, msg *cosmosupgradev1beta1types.QueryAppliedPlanRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosupgradev1beta1keeper.AppliedPlan(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosupgradev1beta1keeper
// Types: cosmosupgradev1beta1types
// github.com/cosmos/cosmos-sdk/x/upgrade/keeper
func (rpcservice *RpcService) Cosmosupgradev1beta1queryupgradedconsensusstate(_ *http.Request, msg *cosmosupgradev1beta1types.QueryUpgradedConsensusStateRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosupgradev1beta1keeper.UpgradedConsensusState(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosupgradev1beta1keeper
// Types: cosmosupgradev1beta1types
// github.com/cosmos/cosmos-sdk/x/upgrade/keeper
func (rpcservice *RpcService) Cosmosupgradev1beta1querymoduleversions(_ *http.Request, msg *cosmosupgradev1beta1types.QueryModuleVersionsRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosupgradev1beta1keeper.ModuleVersions(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosupgradev1beta1keeper
// Types: cosmosupgradev1beta1types
// github.com/cosmos/cosmos-sdk/x/upgrade/keeper
func (rpcservice *RpcService) Cosmosupgradev1beta1queryauthority(_ *http.Request, msg *cosmosupgradev1beta1types.QueryAuthorityRequest, response *string) (err error) {
	r, err := rpcservice.k.cosmosupgradev1beta1keeper.Authority(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosupgradev1beta1keeper
// Types: cosmosupgradev1beta1types
// github.com/cosmos/cosmos-sdk/x/upgrade/keeper
func (rpcservice *RpcService) Cosmosupgradev1beta1sendmsgsoftwareupgrade(_ *http.Request, msg *cosmosupgradev1beta1types.MsgSoftwareUpgrade, response *string) (err error) {
	//handler := cosmosupgradev1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosupgradev1beta1keeper).SoftwareUpgrade
	handler := cosmosupgradev1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosupgradev1beta1keeper).SoftwareUpgrade
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: cosmosupgradev1beta1keeper
// Types: cosmosupgradev1beta1types
// github.com/cosmos/cosmos-sdk/x/upgrade/keeper
func (rpcservice *RpcService) Cosmosupgradev1beta1sendmsgcancelupgrade(_ *http.Request, msg *cosmosupgradev1beta1types.MsgCancelUpgrade, response *string) (err error) {
	//handler := cosmosupgradev1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosupgradev1beta1keeper).CancelUpgrade
	handler := cosmosupgradev1beta1keeper.NewMsgServerImpl(rpcservice.k.cosmosupgradev1beta1keeper).CancelUpgrade
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: ibcapplicationstransferv1keeper
// Types: ibcapplicationstransferv1types
// None
func (rpcservice *RpcService) Ibcapplicationstransferv1querydenomtrace(_ *http.Request, msg *ibcapplicationstransferv1types.QueryDenomTraceRequest, response *string) (err error) {
	r, err := rpcservice.k.ibcapplicationstransferv1keeper.DenomTrace(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: ibcapplicationstransferv1keeper
// Types: ibcapplicationstransferv1types
// None
func (rpcservice *RpcService) Ibcapplicationstransferv1querydenomtraces(_ *http.Request, msg *ibcapplicationstransferv1types.QueryDenomTracesRequest, response *string) (err error) {
	r, err := rpcservice.k.ibcapplicationstransferv1keeper.DenomTraces(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: ibcapplicationstransferv1keeper
// Types: ibcapplicationstransferv1types
// None
func (rpcservice *RpcService) Ibcapplicationstransferv1queryparams(_ *http.Request, msg *ibcapplicationstransferv1types.QueryParamsRequest, response *string) (err error) {
	r, err := rpcservice.k.ibcapplicationstransferv1keeper.Params(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: ibcapplicationstransferv1keeper
// Types: ibcapplicationstransferv1types
// None
func (rpcservice *RpcService) Ibcapplicationstransferv1querydenomhash(_ *http.Request, msg *ibcapplicationstransferv1types.QueryDenomHashRequest, response *string) (err error) {
	r, err := rpcservice.k.ibcapplicationstransferv1keeper.DenomHash(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: ibcapplicationstransferv1keeper
// Types: ibcapplicationstransferv1types
// None
func (rpcservice *RpcService) Ibcapplicationstransferv1queryescrowaddress(_ *http.Request, msg *ibcapplicationstransferv1types.QueryEscrowAddressRequest, response *string) (err error) {
	r, err := rpcservice.k.ibcapplicationstransferv1keeper.EscrowAddress(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: ibcapplicationstransferv1keeper
// Types: ibcapplicationstransferv1types
// None
func (rpcservice *RpcService) Ibcapplicationstransferv1sendmsgtransfer(_ *http.Request, msg *ibcapplicationstransferv1types.MsgTransfer, response *string) (err error) {
	//handler := ibcapplicationstransferv1keeper.NewMsgServerImpl(rpcservice.k.ibcapplicationstransferv1keeper).Transfer
	handler := rpcservice.k.ibcapplicationstransferv1keeper.Transfer
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: dysonkeeper
// Types: dysontypes
// github.com/org/dyson/x/dyson
func (rpcservice *RpcService) Dysonqueryscheduledrun(_ *http.Request, msg *dysontypes.QueryGetScheduledRunRequest, response *string) (err error) {
	r, err := rpcservice.k.ScheduledRun(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: dysonkeeper
// Types: dysontypes
// github.com/org/dyson/x/dyson
func (rpcservice *RpcService) Dysonqueryscheduledrunall(_ *http.Request, msg *dysontypes.QueryAllScheduledRunRequest, response *string) (err error) {
	r, err := rpcservice.k.ScheduledRunAll(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: dysonkeeper
// Types: dysontypes
// github.com/org/dyson/x/dyson
func (rpcservice *RpcService) Dysonquerystorage(_ *http.Request, msg *dysontypes.QueryGetStorageRequest, response *string) (err error) {
	r, err := rpcservice.k.Storage(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: dysonkeeper
// Types: dysontypes
// github.com/org/dyson/x/dyson
func (rpcservice *RpcService) Dysonquerystorageall(_ *http.Request, msg *dysontypes.QueryAllStorageRequest, response *string) (err error) {
	r, err := rpcservice.k.StorageAll(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: dysonkeeper
// Types: dysontypes
// github.com/org/dyson/x/dyson
func (rpcservice *RpcService) Dysonqueryscript(_ *http.Request, msg *dysontypes.QueryGetScriptRequest, response *string) (err error) {
	r, err := rpcservice.k.Script(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: dysonkeeper
// Types: dysontypes
// github.com/org/dyson/x/dyson
func (rpcservice *RpcService) Dysonqueryschema(_ *http.Request, msg *dysontypes.QueryGetSchemaRequest, response *string) (err error) {
	r, err := rpcservice.k.Schema(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: dysonkeeper
// Types: dysontypes
// github.com/org/dyson/x/dyson
func (rpcservice *RpcService) Dysonquerywsgi(_ *http.Request, msg *dysontypes.QueryWsgiRequest, response *string) (err error) {
	r, err := rpcservice.k.Wsgi(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: dysonkeeper
// Types: dysontypes
// github.com/org/dyson/x/dyson
func (rpcservice *RpcService) Dysonqueryqueryscript(_ *http.Request, msg *dysontypes.MsgRun, response *string) (err error) {
	r, err := rpcservice.k.QueryScript(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: dysonkeeper
// Types: dysontypes
// github.com/org/dyson/x/dyson
func (rpcservice *RpcService) Dysonqueryscriptall(_ *http.Request, msg *dysontypes.QueryAllScriptRequest, response *string) (err error) {
	r, err := rpcservice.k.ScriptAll(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: dysonkeeper
// Types: dysontypes
// github.com/org/dyson/x/dyson
func (rpcservice *RpcService) Dysonqueryprefixstorage(_ *http.Request, msg *dysontypes.QueryPrefixStorageRequest, response *string) (err error) {
	r, err := rpcservice.k.PrefixStorage(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: dysonkeeper
// Types: dysontypes
// github.com/org/dyson/x/dyson
func (rpcservice *RpcService) Dysonqueryscheduledgaspriceandfeeatblock(_ *http.Request, msg *dysontypes.QueryScheduledGasPriceAndFeeAtBlockRequest, response *string) (err error) {
	r, err := rpcservice.k.ScheduledGasPriceAndFeeAtBlock(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: dysonkeeper
// Types: dysontypes
// github.com/org/dyson/x/dyson
func (rpcservice *RpcService) Dysonquerycron(_ *http.Request, msg *dysontypes.QueryGetCronRequest, response *string) (err error) {
	r, err := rpcservice.k.Cron(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: dysonkeeper
// Types: dysontypes
// github.com/org/dyson/x/dyson
func (rpcservice *RpcService) Dysonquerycronall(_ *http.Request, msg *dysontypes.QueryAllCronRequest, response *string) (err error) {
	r, err := rpcservice.k.CronAll(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: dysonkeeper
// Types: dysontypes
// github.com/org/dyson/x/dyson
func (rpcservice *RpcService) Dysonsendmsgcreatescheduledrun(_ *http.Request, msg *dysontypes.MsgCreateScheduledRun, response *string) (err error) {
	//handler := dysonkeeper.NewMsgServerImpl(rpcservice.k.dysonkeeper).CreateScheduledRun
	handler := rpcservice.m.CreateScheduledRun
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: dysonkeeper
// Types: dysontypes
// github.com/org/dyson/x/dyson
func (rpcservice *RpcService) Dysonsendmsgbettersubmitproposal(_ *http.Request, msg *dysontypes.MsgBetterSubmitProposal, response *string) (err error) {
	//handler := dysonkeeper.NewMsgServerImpl(rpcservice.k.dysonkeeper).BetterSubmitProposal
	handler := rpcservice.m.BetterSubmitProposal
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: dysonkeeper
// Types: dysontypes
// github.com/org/dyson/x/dyson
func (rpcservice *RpcService) Dysonsendmsgcreatestorage(_ *http.Request, msg *dysontypes.MsgCreateStorage, response *string) (err error) {
	//handler := dysonkeeper.NewMsgServerImpl(rpcservice.k.dysonkeeper).CreateStorage
	handler := rpcservice.m.CreateStorage
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: dysonkeeper
// Types: dysontypes
// github.com/org/dyson/x/dyson
func (rpcservice *RpcService) Dysonsendmsgupdatestorage(_ *http.Request, msg *dysontypes.MsgUpdateStorage, response *string) (err error) {
	//handler := dysonkeeper.NewMsgServerImpl(rpcservice.k.dysonkeeper).UpdateStorage
	handler := rpcservice.m.UpdateStorage
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: dysonkeeper
// Types: dysontypes
// github.com/org/dyson/x/dyson
func (rpcservice *RpcService) Dysonsendmsgdeletestorage(_ *http.Request, msg *dysontypes.MsgDeleteStorage, response *string) (err error) {
	//handler := dysonkeeper.NewMsgServerImpl(rpcservice.k.dysonkeeper).DeleteStorage
	handler := rpcservice.m.DeleteStorage
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: dysonkeeper
// Types: dysontypes
// github.com/org/dyson/x/dyson
func (rpcservice *RpcService) Dysonsendmsgcreatescript(_ *http.Request, msg *dysontypes.MsgCreateScript, response *string) (err error) {
	//handler := dysonkeeper.NewMsgServerImpl(rpcservice.k.dysonkeeper).CreateScript
	handler := rpcservice.m.CreateScript
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: dysonkeeper
// Types: dysontypes
// github.com/org/dyson/x/dyson
func (rpcservice *RpcService) Dysonsendmsgupdatescript(_ *http.Request, msg *dysontypes.MsgUpdateScript, response *string) (err error) {
	//handler := dysonkeeper.NewMsgServerImpl(rpcservice.k.dysonkeeper).UpdateScript
	handler := rpcservice.m.UpdateScript
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: dysonkeeper
// Types: dysontypes
// github.com/org/dyson/x/dyson
func (rpcservice *RpcService) Dysonsendmsgdeletescript(_ *http.Request, msg *dysontypes.MsgDeleteScript, response *string) (err error) {
	//handler := dysonkeeper.NewMsgServerImpl(rpcservice.k.dysonkeeper).DeleteScript
	handler := rpcservice.m.DeleteScript
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: dysonkeeper
// Types: dysontypes
// github.com/org/dyson/x/dyson
func (rpcservice *RpcService) Dysonsendmsgrun(_ *http.Request, msg *dysontypes.MsgRun, response *string) (err error) {
	//handler := dysonkeeper.NewMsgServerImpl(rpcservice.k.dysonkeeper).Run
	handler := rpcservice.m.Run
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: nameskeeper
// Types: namestypes
// github.com/org/dyson/x/names/keeper
func (rpcservice *RpcService) Namesqueryparams(_ *http.Request, msg *namestypes.QueryParamsRequest, response *string) (err error) {
	r, err := rpcservice.k.nameskeeper.Params(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: nameskeeper
// Types: namestypes
// github.com/org/dyson/x/names/keeper
func (rpcservice *RpcService) Namesqueryname(_ *http.Request, msg *namestypes.QueryGetNameRequest, response *string) (err error) {
	r, err := rpcservice.k.nameskeeper.Name(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: nameskeeper
// Types: namestypes
// github.com/org/dyson/x/names/keeper
func (rpcservice *RpcService) Namesquerynameall(_ *http.Request, msg *namestypes.QueryAllNameRequest, response *string) (err error) {
	r, err := rpcservice.k.nameskeeper.NameAll(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: nameskeeper
// Types: namestypes
// github.com/org/dyson/x/names/keeper
func (rpcservice *RpcService) Namesqueryresolve(_ *http.Request, msg *namestypes.QueryResolveRequest, response *string) (err error) {
	r, err := rpcservice.k.nameskeeper.Resolve(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: nameskeeper
// Types: namestypes
// github.com/org/dyson/x/names/keeper
func (rpcservice *RpcService) Namesquerygeneratecommit(_ *http.Request, msg *namestypes.QueryGenerateCommitRequest, response *string) (err error) {
	r, err := rpcservice.k.nameskeeper.GenerateCommit(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: nameskeeper
// Types: namestypes
// github.com/org/dyson/x/names/keeper
func (rpcservice *RpcService) Namesqueryexpirations(_ *http.Request, msg *namestypes.QueryGetExpirationsRequest, response *string) (err error) {
	r, err := rpcservice.k.nameskeeper.Expirations(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: nameskeeper
// Types: namestypes
// github.com/org/dyson/x/names/keeper
func (rpcservice *RpcService) Namesqueryexpirationsall(_ *http.Request, msg *namestypes.QueryAllExpirationsRequest, response *string) (err error) {
	r, err := rpcservice.k.nameskeeper.ExpirationsAll(rpcservice.ctx, msg)
	if err != nil {
		return err
	}
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: nameskeeper
// Types: namestypes
// github.com/org/dyson/x/names/keeper
func (rpcservice *RpcService) Namessendmsgregister(_ *http.Request, msg *namestypes.MsgRegister, response *string) (err error) {
	//handler := nameskeeper.NewMsgServerImpl(rpcservice.k.nameskeeper).Register
	handler := nameskeeper.NewMsgServerImpl(rpcservice.k.nameskeeper).Register
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: nameskeeper
// Types: namestypes
// github.com/org/dyson/x/names/keeper
func (rpcservice *RpcService) Namessendmsgupdatename(_ *http.Request, msg *namestypes.MsgUpdateName, response *string) (err error) {
	//handler := nameskeeper.NewMsgServerImpl(rpcservice.k.nameskeeper).UpdateName
	handler := nameskeeper.NewMsgServerImpl(rpcservice.k.nameskeeper).UpdateName
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: nameskeeper
// Types: namestypes
// github.com/org/dyson/x/names/keeper
func (rpcservice *RpcService) Namessendmsgdeletename(_ *http.Request, msg *namestypes.MsgDeleteName, response *string) (err error) {
	//handler := nameskeeper.NewMsgServerImpl(rpcservice.k.nameskeeper).DeleteName
	handler := nameskeeper.NewMsgServerImpl(rpcservice.k.nameskeeper).DeleteName
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: nameskeeper
// Types: namestypes
// github.com/org/dyson/x/names/keeper
func (rpcservice *RpcService) Namessendmsgreveal(_ *http.Request, msg *namestypes.MsgReveal, response *string) (err error) {
	//handler := nameskeeper.NewMsgServerImpl(rpcservice.k.nameskeeper).Reveal
	handler := nameskeeper.NewMsgServerImpl(rpcservice.k.nameskeeper).Reveal
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: nameskeeper
// Types: namestypes
// github.com/org/dyson/x/names/keeper
func (rpcservice *RpcService) Namessendmsgsetpriceandextend(_ *http.Request, msg *namestypes.MsgSetPriceAndExtend, response *string) (err error) {
	//handler := nameskeeper.NewMsgServerImpl(rpcservice.k.nameskeeper).SetPriceAndExtend
	handler := nameskeeper.NewMsgServerImpl(rpcservice.k.nameskeeper).SetPriceAndExtend
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: nameskeeper
// Types: namestypes
// github.com/org/dyson/x/names/keeper
func (rpcservice *RpcService) Namessendmsgofferto(_ *http.Request, msg *namestypes.MsgOfferTo, response *string) (err error) {
	//handler := nameskeeper.NewMsgServerImpl(rpcservice.k.nameskeeper).OfferTo
	handler := nameskeeper.NewMsgServerImpl(rpcservice.k.nameskeeper).OfferTo
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: nameskeeper
// Types: namestypes
// github.com/org/dyson/x/names/keeper
func (rpcservice *RpcService) Namessendmsgaccept(_ *http.Request, msg *namestypes.MsgAccept, response *string) (err error) {
	//handler := nameskeeper.NewMsgServerImpl(rpcservice.k.nameskeeper).Accept
	handler := nameskeeper.NewMsgServerImpl(rpcservice.k.nameskeeper).Accept
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: nameskeeper
// Types: namestypes
// github.com/org/dyson/x/names/keeper
func (rpcservice *RpcService) Namessendmsgbuy(_ *http.Request, msg *namestypes.MsgBuy, response *string) (err error) {
	//handler := nameskeeper.NewMsgServerImpl(rpcservice.k.nameskeeper).Buy
	handler := nameskeeper.NewMsgServerImpl(rpcservice.k.nameskeeper).Buy
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: nameskeeper
// Types: namestypes
// github.com/org/dyson/x/names/keeper
func (rpcservice *RpcService) Namessendmsgmintcoins(_ *http.Request, msg *namestypes.MsgMintCoins, response *string) (err error) {
	//handler := nameskeeper.NewMsgServerImpl(rpcservice.k.nameskeeper).MintCoins
	handler := nameskeeper.NewMsgServerImpl(rpcservice.k.nameskeeper).MintCoins
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}

// Keeper: nameskeeper
// Types: namestypes
// github.com/org/dyson/x/names/keeper
func (rpcservice *RpcService) Namessendmsgburncoins(_ *http.Request, msg *namestypes.MsgBurnCoins, response *string) (err error) {
	//handler := nameskeeper.NewMsgServerImpl(rpcservice.k.nameskeeper).BurnCoins
	handler := nameskeeper.NewMsgServerImpl(rpcservice.k.nameskeeper).BurnCoins
	//
	defer func() {
		if r := recover(); r != nil {

			err = sdkerrors.Wrapf(types.RpcError, "CHAIN ERROR: %T %+v", r, r)
		}
	}()
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	if len(msg.GetSigners()) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "this requires more than one signer and cannot be run from a script")
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
	*response = string(rpcservice.k.cdc.MustMarshalJSON(r))
	return nil
}
