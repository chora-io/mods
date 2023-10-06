package keeper

import (
	"context"

	"cosmossdk.io/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

	contentv1 "github.com/choraio/mods/content/api/v1"
	v1 "github.com/choraio/mods/content/types/v1"
	"github.com/choraio/mods/content/utils"
)

// Contents implements the Query/Contents method.
func (k Keeper) Contents(ctx context.Context, req *v1.QueryContentsRequest) (*v1.QueryContentsResponse, error) {

	// set the index for table lookup
	index := contentv1.ContentIdIndexKey{}

	// set the pagination for table lookup
	pgnReq, err := utils.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err // internal error
	}

	// get contents from content table
	it, err := k.ss.ContentTable().List(ctx, index, ormlist.Paginate(pgnReq))
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

		content := v1.QueryContentsResponse_Content{
			Id:       v.Id,
			Curator:  curator,
			Metadata: v.Metadata,
		}

		contents = append(contents, &content)
	}

	// set the pagination for query response
	pgnRes, err := utils.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err // internal error
	}

	// return query response
	return &v1.QueryContentsResponse{
		Contents:   contents,
		Pagination: pgnRes,
	}, nil
}
