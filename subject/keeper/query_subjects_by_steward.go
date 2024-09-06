package keeper

import (
	"context"

	"cosmossdk.io/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	subjectv1 "github.com/chora-io/mods/subject/api/v1"
	v1 "github.com/chora-io/mods/subject/types/v1"
	"github.com/chora-io/mods/subject/utils"
)

// SubjectsBySteward implements the Query/SubjectsBySteward method.
func (k Keeper) SubjectsBySteward(ctx context.Context, req *v1.QuerySubjectsByStewardRequest) (*v1.QuerySubjectsByStewardResponse, error) {

	// get account from steward address
	steward, err := sdk.AccAddressFromBech32(req.Steward)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("steward: %s", err)
	}

	// set index for table lookup
	index := subjectv1.SubjectStewardIndexKey{}.WithSteward(steward)

	// set pagination for table lookup
	pgnReq, err := utils.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err // internal error
	}

	// get subjects by steward from subject table
	it, err := k.ss.SubjectTable().List(ctx, index, ormlist.Paginate(pgnReq))
	if err != nil {
		return nil, err // internal error
	}

	// set subjects for query response
	subjects := make([]*v1.QuerySubjectsByStewardResponse_Subject, 0, 10)
	for it.Next() {
		v, err := it.Value()
		if err != nil {
			return nil, err // internal error
		}

		address := sdk.AccAddress(v.Address).String()

		subject := v1.QuerySubjectsByStewardResponse_Subject{
			Address:  address,
			Metadata: v.Metadata,
		}

		subjects = append(subjects, &subject)
	}

	// set pagination for query response
	pgnRes, err := utils.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err // internal error
	}

	// return query response
	return &v1.QuerySubjectsByStewardResponse{
		Steward:    steward.String(),
		Subjects:   subjects,
		Pagination: pgnRes,
	}, nil
}
