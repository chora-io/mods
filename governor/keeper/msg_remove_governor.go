package keeper

import (
	"context"

	"cosmossdk.io/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/chora-io/mods/governor/types/v1"
)

// RemoveGovernor implements Msg/RemoveGovernor.
func (k Keeper) RemoveGovernor(ctx context.Context, req *v1.MsgRemoveGovernor) (*v1.MsgRemoveGovernorResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get account address from address
	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err // internal error
	}

	// get governor from governor table
	governor, err := k.ss.GovernorTable().Get(ctx, address)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf(
				"governor with address %s: %s", req.Address, err,
			)
		}
		return nil, err // internal error
	}

	// delete governor from governor table
	err = k.ss.GovernorTable().Delete(ctx, governor)
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventRemoveGovernor{
		Address: req.Address,
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgRemoveGovernorResponse{
		Address: req.Address,
	}, nil
}
