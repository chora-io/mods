package keeper

import (
	"context"

	"cosmossdk.io/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/chora-io/mods/admin/types/v1"
)

// UpdateAdmin implements the Msg/UpdateAdmin method.
func (k Keeper) UpdateAdmin(ctx context.Context, req *v1.MsgUpdateAdmin) (*v1.MsgUpdateAdminResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get account from admin address
	reqAdmin, err := sdk.AccAddressFromBech32(req.Admin)
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
	adminAdmin := sdk.AccAddress(admin.Admin)

	// verify admin is admin account
	if !reqAdmin.Equals(adminAdmin) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"admin %s: admin account %s", reqAdmin, adminAdmin,
		)
	}

	// get account from new admin address
	newAdmin, err := sdk.AccAddressFromBech32(req.NewAdmin)
	if err != nil {
		return nil, err // internal error
	}

	// set new admin
	admin.Admin = newAdmin

	// update admin in admin table
	err = k.ss.AdminTable().Save(ctx, admin)
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventUpdateAdmin{
		Admin:    reqAdmin.String(),
		NewAdmin: newAdmin.String(),
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgUpdateAdminResponse{
		Admin:    reqAdmin.String(),
		NewAdmin: newAdmin.String(),
	}, nil
}
