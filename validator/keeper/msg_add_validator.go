package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	validatorv1 "github.com/chora-io/mods/validator/api/v1"
	v1 "github.com/chora-io/mods/validator/types/v1"
)

// AddValidator implements Msg/AddValidator.
func (k Keeper) AddValidator(ctx context.Context, req *v1.MsgAddValidator) (*v1.MsgAddValidatorResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	adminAddress := k.admin.String()
	if adminAddress != req.Admin {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"admin: expected %s: received %s", adminAddress, req.Admin,
		)
	}

	// insert validator into validator table
	err := k.ss.ValidatorTable().Insert(ctx, &validatorv1.Validator{
		Address:  req.Address,
		Metadata: req.Metadata,
	})
	if err != nil {
		return nil, err // internal error
	}

	// insert validator into validator signing info table
	err = k.ss.ValidatorSigningInfoTable().Insert(ctx, &validatorv1.ValidatorSigningInfo{
		Address: req.Address,
	})
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventAddValidator{
		Address: req.Address,
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgAddValidatorResponse{
		Address: req.Address,
	}, nil
}
