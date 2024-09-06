package keeper

import (
	"context"

	"cosmossdk.io/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/chora-io/mods/admin/types/v1"
)

// Admin implements the Query/Admin method.
func (k Keeper) Admin(ctx context.Context, req *v1.QueryAdminRequest) (*v1.QueryAdminResponse, error) {

	// get admin from admin table
	admin, err := k.ss.AdminTable().Get(ctx)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("admin")
		}
		return nil, err // internal error
	}

	// get account from account bytes
	curator := sdk.AccAddress(admin.Address)

	// return query response
	return &v1.QueryAdminResponse{
		Admin: curator.String(),
	}, nil
}
