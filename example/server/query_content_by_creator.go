package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types/v2/ormutil"

	"github.com/choraio/mods/example/api/v1"
	"github.com/choraio/mods/example/types/v1"
)

// ContentByCreator implements the Query/ContentByCreator method.
func (s Server) ContentByCreator(ctx context.Context, req *v1.QueryContentByCreatorRequest) (*v1.QueryContentByCreatorResponse, error) {

	// convert creator to account
	creator := sdk.AccAddress(req.Creator)

	// set the index for table lookup
	index := examplev1.ContentCreatorIndexKey{}.WithCreator(creator)

	// set the pagination for table lookup
	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf(
			"invalid pagination: %s", err,
		)
	}

	// get iterator for content by creator from content table
	it, err := s.ss.ContentTable().List(ctx, index, ormlist.Paginate(pg))
	if err != nil {
		return nil, err // internal error
	}

	// set content for content by creator response
	content := make([]*v1.QueryContentByCreatorResponse_Content, 0, 10)
	for it.Next() {
		v, err := it.Value()
		if err != nil {
			return nil, err // internal error
		}
		c := v1.QueryContentByCreatorResponse_Content{
			Id:   v.Id,
			Hash: v.Hash,
		}
		content = append(content, &c)
	}

	// return query content by creator response
	return &v1.QueryContentByCreatorResponse{
		Creator: creator.String(),
		Content: content,
	}, nil
}
