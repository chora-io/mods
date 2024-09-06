package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	validatorv1 "github.com/chora-io/mods/validator/api/v1"
	v1 "github.com/chora-io/mods/validator/types/v1"
)

// CreateValidator implements Msg/CreateValidator.
func (k Keeper) CreateValidator(ctx context.Context, req *v1.MsgCreateValidator) (*v1.MsgCreateValidatorResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	operator := k.admin.String()
	if operator != req.Operator {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"operator: expected %s: received %s", operator, req.Operator,
		)
	}

	// TODO: generate validator address
	validator := req.Operator

	// insert validator into validator table
	err := k.ss.ValidatorTable().Insert(ctx, &validatorv1.Validator{
		Address:  validator,
		Metadata: req.Metadata,
	})
	if err != nil {
		return nil, err // internal error
	}

	// insert validator into validator signing info table
	err = k.ss.ValidatorSigningInfoTable().Insert(ctx, &validatorv1.ValidatorSigningInfo{
		Address: validator,
	})
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventCreateValidator{
		Address: validator,
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgCreateValidatorResponse{
		Address: validator,
	}, nil
}
