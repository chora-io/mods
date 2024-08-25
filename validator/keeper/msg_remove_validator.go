package keeper

import (
	"context"

	"cosmossdk.io/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/chora-io/mods/validator/types/v1"
)

// RemoveValidator implements Msg/RemoveValidator.
func (k Keeper) RemoveValidator(ctx context.Context, req *v1.MsgRemoveValidator) (*v1.MsgRemoveValidatorResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	adminAddress := k.admin.String()
	if adminAddress != req.Admin {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"admin: expected %s: received %s", adminAddress, req.Admin,
		)
	}

	// get validator from validator table
	validator, err := k.ss.ValidatorTable().Get(ctx, req.Address)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf(
				"validator with address %s: %s", req.Address, err,
			)
		}
		return nil, err // internal error
	}

	// get signing info from validator signing info table
	signingInfo, err := k.ss.ValidatorSigningInfoTable().Get(ctx, req.Address)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf(
				"validator signing info with address %s", req.Address,
			)
		}
		return nil, err // internal error
	}

	// delete validator from validator table
	err = k.ss.ValidatorTable().Delete(ctx, validator)
	if err != nil {
		return nil, err // internal error
	}

	// delete validator from validator signing info table
	err = k.ss.ValidatorSigningInfoTable().Delete(ctx, signingInfo)
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventRemoveValidator{
		Address: req.Address,
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgRemoveValidatorResponse{
		Address: req.Address,
	}, nil
}
