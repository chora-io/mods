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

	adminAddress := k.admin.String()
	if adminAddress != req.Admin {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"admin: expected %s: received %s", adminAddress, req.Admin,
		)
	}

	// get governor from governor table
	governor, err := k.ss.GovernorTable().Get(ctx, req.Address)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf(
				"governor with address %s: %s", req.Address, err,
			)
		}
		return nil, err // internal error
	}

	// get signing info from governor signing info table
	signingInfo, err := k.ss.GovernorSigningInfoTable().Get(ctx, req.Address)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf(
				"governor signing info with address %s", req.Address,
			)
		}
		return nil, err // internal error
	}

	// delete governor from governor table
	err = k.ss.GovernorTable().Delete(ctx, governor)
	if err != nil {
		return nil, err // internal error
	}

	// delete governor from governor signing info table
	err = k.ss.GovernorSigningInfoTable().Delete(ctx, signingInfo)
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
