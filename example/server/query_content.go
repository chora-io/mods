package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/choraio/mods/example/types/v1"
)

// Content implements the Query/Content method.
func (s Server) Content(ctx context.Context, req *v1.QueryContentRequest) (*v1.QueryContentResponse, error) {

	// get content from content table
	content, err := s.ss.ContentTable().Get(ctx, req.Id)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf(
				"content with id %s: %s", req.Id, err,
			)
		}
		return nil, err // internal error
	}

	// convert creator bytes to account
	creator := sdk.AccAddress(content.Creator)

	// return query content response
	return &v1.QueryContentResponse{
		Id:      content.Id,
		Creator: creator.String(),
		Hash:    content.Hash,
	}, nil
}
