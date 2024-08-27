package keeper

import (
	"context"

	governorv1 "github.com/chora-io/mods/governor/api/v1"
	v1 "github.com/chora-io/mods/governor/types/v1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CreateGovernor implements Msg/CreateGovernor.
func (k Keeper) CreateGovernor(ctx context.Context, req *v1.MsgCreateGovernor) (*v1.MsgCreateGovernorResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get account address from address
	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err // internal error
	}

	// insert governor into governor table
	err = k.ss.GovernorTable().Insert(ctx, &governorv1.Governor{
		Address:  address,
		Metadata: req.Metadata,
	})
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventCreateGovernor{
		Address: req.Address,
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgCreateGovernorResponse{
		Address: req.Address,
	}, nil
}
