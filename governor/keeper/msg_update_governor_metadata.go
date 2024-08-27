package keeper

import (
	"context"

	"cosmossdk.io/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/chora-io/mods/governor/types/v1"
)

// UpdateGovernorMetadata implements the Msg/UpdateGovernorMetadata method.
func (k Keeper) UpdateGovernorMetadata(ctx context.Context, req *v1.MsgUpdateGovernorMetadata) (*v1.MsgUpdateGovernorMetadataResponse, error) {
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
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventUpdateGovernorMetadata{
		Address: req.Address,
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgUpdateGovernorMetadataResponse{
		Address: req.Address,
	}, nil
}
