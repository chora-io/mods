package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	validatorv1 "github.com/choraio/mods/validator/api/v1"
	v1 "github.com/choraio/mods/validator/types/v1"
)

// AddValidator implements Msg/AddValidator.
func (k Keeper) AddValidator(ctx context.Context, req *v1.MsgAddValidator) (*v1.MsgAddValidatorResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	authorityAddress := k.authority.String()
	if authorityAddress != req.Authority {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"authority: expected %s: received %s", authorityAddress, req.Authority,
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
