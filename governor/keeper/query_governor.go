package keeper

import (
	"context"

	"cosmossdk.io/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/chora-io/mods/governor/types/v1"
)

// Governor implements the Query/Governor method.
func (k Keeper) Governor(ctx context.Context, req *v1.QueryGovernorRequest) (*v1.QueryGovernorResponse, error) {

	// get account address from address
	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err // internal error
	}

	// get governor from governor table
	governor, err := k.ss.GovernorTable().Get(ctx, address)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("governor with address %s", req.Address)
		}
		return nil, err // internal error
	}

	// return query response
	return &v1.QueryGovernorResponse{
		Address:  req.Address,
		Metadata: governor.Metadata,
	}, nil
}
