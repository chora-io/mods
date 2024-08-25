package keeper

import (
	"context"

	"cosmossdk.io/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/chora-io/mods/governor/types/v1"
)

// UpdateGovernor implements the Msg/UpdateGovernor method.
func (k Keeper) UpdateGovernor(ctx context.Context, req *v1.MsgUpdateGovernor) (*v1.MsgUpdateGovernorResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get governor from governor table
	governor, err := k.ss.GovernorTable().Get(ctx, req.Address)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("governor with address %s: %s", req.Address, err)
		}
		return nil, err // internal error
	}

	// set new governor metadata
	governor.Metadata = req.NewMetadata

	// update governor in governor table
	err = k.ss.GovernorTable().Update(ctx, governor)
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventUpdateGovernor{
		Address: req.Address,
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgUpdateGovernorResponse{
		Address: req.Address,
	}, nil
}
