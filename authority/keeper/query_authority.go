package keeper

import (
	"context"

	"cosmossdk.io/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/chora-io/mods/authority/types/v1"
)

// Authority implements the Query/Authority method.
func (k Keeper) Authority(ctx context.Context, req *v1.QueryAuthorityRequest) (*v1.QueryAuthorityResponse, error) {

	// get authority from authority table
	authority, err := k.ss.AuthorityTable().Get(ctx)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("authority")
		}
		return nil, err // internal error
	}

	// get account from account bytes
	curator := sdk.AccAddress(authority.Address)

	// return query response
	return &v1.QueryAuthorityResponse{
		Authority: curator.String(),
	}, nil
}
