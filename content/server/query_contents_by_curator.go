package server

import (
	"context"

	"cosmossdk.io/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	contentv1 "github.com/choraio/mods/content/api/v1"
	v1 "github.com/choraio/mods/content/types/v1"
	"github.com/choraio/mods/content/utils"
)

// ContentsByCurator implements the Query/ContentsByCurator method.
func (s Server) ContentsByCurator(ctx context.Context, req *v1.QueryContentsByCuratorRequest) (*v1.QueryContentsByCuratorResponse, error) {

	// get account from curator address
	curator, err := sdk.AccAddressFromBech32(req.Curator)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("curator: %s", err)
	}

	// set index for table lookup
	index := contentv1.ContentCuratorIndexKey{}.WithCurator(curator)

	// set pagination for table lookup
	pgnReq, err := utils.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err // internal error
	}

	// get contents by curator from content table
	it, err := s.ss.ContentTable().List(ctx, index, ormlist.Paginate(pgnReq))
	if err != nil {
		return nil, err // internal error
	}

	// set contents for query response
	contents := make([]*v1.QueryContentsByCuratorResponse_Content, 0, 10)
	for it.Next() {
		v, err := it.Value()
		if err != nil {
			return nil, err // internal error
		}

		content := v1.QueryContentsByCuratorResponse_Content{
			Id:       v.Id,
			Metadata: v.Metadata,
		}

		contents = append(contents, &content)
	}

	// set pagination for query response
	pgnRes, err := utils.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err // internal error
	}

	// return query response
	return &v1.QueryContentsByCuratorResponse{
		Curator:    curator.String(),
		Contents:   contents,
		Pagination: pgnRes,
	}, nil
}
