package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/types/v2/ormutil"

	contentv1 "github.com/choraio/mods/content/api/v1"
	v1 "github.com/choraio/mods/content/types/v1"
)

// Contents implements the Query/Contents method.
func (s Server) Contents(ctx context.Context, req *v1.QueryContentsRequest) (*v1.QueryContentsResponse, error) {

	// set the index for table lookup
	index := contentv1.ContentIdIndexKey{}

	// set the pagination for table lookup
	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err // internal error
	}

	// get contents from content table
	it, err := s.ss.ContentTable().List(ctx, index, ormlist.Paginate(pg))
	if err != nil {
		return nil, err // internal error
	}

	// set contents for query response
	contents := make([]*v1.QueryContentsResponse_Content, 0, 10)
	for it.Next() {
		v, err := it.Value()
		if err != nil {
			return nil, err // internal error
		}

		curator := sdk.AccAddress(v.Curator).String()

		c := v1.QueryContentsResponse_Content{
			Id:       v.Id,
			Curator:  curator,
			Metadata: v.Metadata,
		}

		contents = append(contents, &c)
	}

	// set the pagination for query response
	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err // internal error
	}

	// return query response
	return &v1.QueryContentsResponse{
		Contents:   contents,
		Pagination: pr,
	}, nil
}
