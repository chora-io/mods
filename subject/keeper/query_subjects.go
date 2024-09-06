package keeper

import (
	"context"

	"cosmossdk.io/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

	subjectv1 "github.com/chora-io/mods/subject/api/v1"
	v1 "github.com/chora-io/mods/subject/types/v1"
	"github.com/chora-io/mods/subject/utils"
)

// Subjects implements the Query/Subjects method.
func (k Keeper) Subjects(ctx context.Context, req *v1.QuerySubjectsRequest) (*v1.QuerySubjectsResponse, error) {

	// set index for table lookup
	index := subjectv1.SubjectAddressIndexKey{}

	// set pagination for table lookup
	pgnReq, err := utils.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err // internal error
	}

	// get subjects from subject table
	it, err := k.ss.SubjectTable().List(ctx, index, ormlist.Paginate(pgnReq))
	if err != nil {
		return nil, err // internal error
	}

	// set subjects for query response
	subjects := make([]*v1.QuerySubjectsResponse_Subject, 0, 10)
	for it.Next() {
		v, err := it.Value()
		if err != nil {
			return nil, err // internal error
		}

		address := sdk.AccAddress(v.Address).String()
		steward := sdk.AccAddress(v.Steward).String()

		subject := v1.QuerySubjectsResponse_Subject{
			Address:  address,
			Steward:  steward,
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
	return &v1.QuerySubjectsResponse{
		Subjects:   subjects,
		Pagination: pgnRes,
	}, nil
}
