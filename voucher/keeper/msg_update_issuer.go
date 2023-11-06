package keeper

import (
	"context"

	"cosmossdk.io/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/chora-io/mods/voucher/types/v1"
)

// UpdateIssuer implements the Msg/UpdateIssuer method.
func (k Keeper) UpdateIssuer(ctx context.Context, req *v1.MsgUpdateIssuer) (*v1.MsgUpdateIssuerResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get account from issuer address
	issuer, err := sdk.AccAddressFromBech32(req.Issuer)
	if err != nil {
		return nil, err // internal error
	}

	// get voucher from voucher table
	voucher, err := k.ss.VoucherTable().Get(ctx, req.Id)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("voucher with id %d: %s", req.Id, err)
		}
		return nil, err // internal error
	}

	// get account from account bytes
	voucherIssuer := sdk.AccAddress(voucher.Issuer)

	// verify issuer is voucher issuer
	if !voucherIssuer.Equals(issuer) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"issuer %s: voucher issuer %s", issuer, voucherIssuer.String(),
		)
	}

	// get account from new issuer address
	newIssuer, err := sdk.AccAddressFromBech32(req.NewIssuer)
	if err != nil {
		return nil, err // internal error
	}

	// set new issuer
	voucher.Issuer = newIssuer

	// update voucher in voucher table
	err = k.ss.VoucherTable().Update(ctx, voucher)
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventUpdateIssuer{
		Id: voucher.Id,
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgUpdateIssuerResponse{
		Id: voucher.Id,
	}, nil
}
