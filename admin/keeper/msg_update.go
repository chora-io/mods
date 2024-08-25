package keeper

import (
	"context"

	"cosmossdk.io/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/chora-io/mods/admin/types/v1"
)

// Update implements the Msg/Update method.
func (k Keeper) Update(ctx context.Context, req *v1.MsgUpdate) (*v1.MsgUpdateResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get account from admin address
	msgSigner, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, err // internal error
	}

	// get admin from admin table
	admin, err := k.ss.AdminTable().Get(ctx)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("admin: %s", err)
		}
		return nil, err // internal error
	}

	// get account from account bytes
	adminAdmin := sdk.AccAddress(admin.Address)

	// verify admin is admin account
	if !adminAdmin.Equals(msgSigner) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"admin %s: admin account %s", msgSigner, adminAdmin.String(),
		)
	}

	// get account from new admin address
	newAdmin, err := sdk.AccAddressFromBech32(req.NewAdmin)
	if err != nil {
		return nil, err // internal error
	}

	// set new admin
	admin.Address = newAdmin

	// update admin in admin table
	err = k.ss.AdminTable().Save(ctx, admin)
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventUpdate{
		Admin:    msgSigner.String(),
		NewAdmin: newAdmin.String(),
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgUpdateResponse{
		Admin:    msgSigner.String(),
		NewAdmin: newAdmin.String(),
	}, nil
}
