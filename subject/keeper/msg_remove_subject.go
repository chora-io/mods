package keeper

import (
	"context"

	"cosmossdk.io/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/chora-io/mods/subject/types/v1"
)

// RemoveSubject implements Msg/RemoveSubject.
func (k Keeper) RemoveSubject(ctx context.Context, req *v1.MsgRemoveSubject) (*v1.MsgRemoveSubjectResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get account from subject address
	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err // internal error
	}

	// get account from steward address
	steward, err := sdk.AccAddressFromBech32(req.Steward)
	if err != nil {
		return nil, err // internal error
	}

	// get subject from subject table
	subject, err := k.ss.SubjectTable().Get(ctx, address)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("subject with address %s: %s", req.Address, err)
		}
		return nil, err // internal error
	}

	// get steward account from subject steward
	subjectSteward := sdk.AccAddress(subject.Steward)

	// verify steward is subject steward
	if !subjectSteward.Equals(steward) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"steward %s: subject steward %s", steward, subjectSteward,
		)
	}

	// delete subject from subject table
	err = k.ss.SubjectTable().Delete(ctx, subject)
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventRemoveSubject{
		Address: req.Address,
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgRemoveSubjectResponse{
		Address: req.Address,
	}, nil
}
