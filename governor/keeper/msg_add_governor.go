package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	governorv1 "github.com/chora-io/mods/governor/api/v1"
	v1 "github.com/chora-io/mods/governor/types/v1"
)

// AddGovernor implements Msg/AddGovernor.
func (k Keeper) AddGovernor(ctx context.Context, req *v1.MsgAddGovernor) (*v1.MsgAddGovernorResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	adminAddress := k.admin.String()
	if adminAddress != req.Admin {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"admin: expected %s: received %s", adminAddress, req.Admin,
		)
	}

	// insert governor into governor table
	err := k.ss.GovernorTable().Insert(ctx, &governorv1.Governor{
		Address:  req.Address,
		Metadata: req.Metadata,
	})
	if err != nil {
		return nil, err // internal error
	}

	// insert governor into governor signing info table
	err = k.ss.GovernorSigningInfoTable().Insert(ctx, &governorv1.GovernorSigningInfo{
		Address: req.Address,
	})
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventAddGovernor{
		Address: req.Address,
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgAddGovernorResponse{
		Address: req.Address,
	}, nil
}
