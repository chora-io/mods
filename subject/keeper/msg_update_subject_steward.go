package keeper

import (
	"context"

	"cosmossdk.io/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/chora-io/mods/subject/types/v1"
)

// UpdateSubjectSteward implements the Msg/UpdateSubjectSteward method.
func (k Keeper) UpdateSubjectSteward(ctx context.Context, req *v1.MsgUpdateSubjectSteward) (*v1.MsgUpdateSubjectStewardResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get subject account from address
	account, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err // internal error
	}

	// get account from steward address
	steward, err := sdk.AccAddressFromBech32(req.Steward)
	if err != nil {
		return nil, err // internal error
	}

	// get subject from subject table
	subject, err := k.ss.SubjectTable().Get(ctx, account)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("subject with address %s: %s", req.Address, err)
		}
		return nil, err // internal error
	}

	// get account from account bytes
	subjectSteward := sdk.AccAddress(subject.Steward)

	// verify steward is subject steward
	if !subjectSteward.Equals(steward) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"steward %s: subject steward %s", steward, subjectSteward.String(),
		)
	}

	// get account from new steward address
	newSteward, err := sdk.AccAddressFromBech32(req.NewSteward)
	if err != nil {
		return nil, err // internal error
	}

	// set new steward
	subject.Steward = newSteward

	// update subject in subject table
	err = k.ss.SubjectTable().Update(ctx, subject)
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventUpdateSubjectSteward{
		Address: req.Address,
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgUpdateSubjectStewardResponse{
		Address: req.Address,
	}, nil
}
