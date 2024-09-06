package keeper

import (
	"context"

	"cosmossdk.io/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/chora-io/mods/content/types/v1"
)

// Content implements the Query/Content method.
func (k Keeper) Content(ctx context.Context, req *v1.QueryContentRequest) (*v1.QueryContentResponse, error) {

	// get content from content table
	content, err := k.ss.ContentTable().Get(ctx, req.Hash)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("content with hash %s", req.Hash)
		}
		return nil, err // internal error
	}

	// get account from account bytes
	curator := sdk.AccAddress(content.Curator)

	// return query response
	return &v1.QueryContentResponse{
		Curator: curator.String(),
		Hash:    content.Hash,
	}, nil
}
