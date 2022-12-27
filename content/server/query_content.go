package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/choraio/mods/content/types/v1"
)

// Content implements the Query/Content method.
func (s Server) Content(ctx context.Context, req *v1.QueryContentRequest) (*v1.QueryContentResponse, error) {

	// get content from content table
	content, err := s.ss.ContentTable().Get(ctx, req.Id)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf(
				"content with id %d", req.Id,
			)
		}
		return nil, err // internal error
	}

	// convert curator bytes to account
	curator := sdk.AccAddress(content.Curator)

	// return query content response
	return &v1.QueryContentResponse{
		Id:      content.Id,
		Curator: curator.String(),
		Hash:    content.Hash,
	}, nil
}
