package keeper

import (
	"context"

	"cosmossdk.io/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/chora-io/mods/validator/types/v1"
)

// UpdateValidator implements the Msg/UpdateValidator method.
func (k Keeper) UpdateValidator(ctx context.Context, req *v1.MsgUpdateValidator) (*v1.MsgUpdateValidatorResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get validator from validator table
	validator, err := k.ss.ValidatorTable().Get(ctx, req.Address)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("validator with address %s: %s", req.Address, err)
		}
		return nil, err // internal error
	}

	// set new validator metadata
	validator.Metadata = req.NewMetadata

	// update validator in validator table
	err = k.ss.ValidatorTable().Update(ctx, validator)
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventUpdateValidator{
		Address: req.Address,
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgUpdateValidatorResponse{
		Address: req.Address,
	}, nil
}
