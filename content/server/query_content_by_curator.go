package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types/v2/ormutil"

	contentv1 "github.com/choraio/mods/content/api/v1"
	v1 "github.com/choraio/mods/content/types/v1"
)

// ContentByCurator implements the Query/ContentByCurator method.
func (s Server) ContentByCurator(ctx context.Context, req *v1.QueryContentByCuratorRequest) (*v1.QueryContentByCuratorResponse, error) {

	// convert curator to account
	curator, err := sdk.AccAddressFromBech32(req.Curator)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("curator: %s", err)
	}

	// set the index for table lookup
	index := contentv1.ContentCuratorIndexKey{}.WithCurator(curator)

	// set the pagination for table lookup
	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err // internal error
	}

	// get iterator for content by curator from content table
	it, err := s.ss.ContentTable().List(ctx, index, ormlist.Paginate(pg))
	if err != nil {
		return nil, err // internal error
	}

	// set content for content by curator response
	content := make([]*v1.QueryContentByCuratorResponse_Content, 0, 10)
	for it.Next() {
		v, err := it.Value()
		if err != nil {
			return nil, err // internal error
		}
		c := v1.QueryContentByCuratorResponse_Content{
			Id:       v.Id,
			Metadata: v.Metadata,
		}
		content = append(content, &c)
	}

	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err // internal error
	}

	// return query content by curator response
	return &v1.QueryContentByCuratorResponse{
		Curator:    curator.String(),
		Content:    content,
		Pagination: pr,
	}, nil
}
