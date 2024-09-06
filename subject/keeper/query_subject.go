package keeper

import (
	"context"

	"cosmossdk.io/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/chora-io/mods/subject/types/v1"
)

// Subject implements the Query/Subject method.
func (k Keeper) Subject(ctx context.Context, req *v1.QuerySubjectRequest) (*v1.QuerySubjectResponse, error) {

	// get subject account from address
	account, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err // internal error
	}

	// get subject from subject table
	subject, err := k.ss.SubjectTable().Get(ctx, account)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("subject with address %s", req.Address)
		}
		return nil, err // internal error
	}

	// get account from account bytes
	steward := sdk.AccAddress(subject.Steward)

	// return query response
	return &v1.QuerySubjectResponse{
		Steward:  steward.String(),
		Address:  account.String(),
		Metadata: subject.Metadata,
	}, nil
}
