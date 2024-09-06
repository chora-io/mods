package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	contentv1 "github.com/chora-io/mods/content/api/v1"
	v1 "github.com/chora-io/mods/content/types/v1"
)

// CreateContent implements Msg/CreateContent.
func (k Keeper) CreateContent(ctx context.Context, req *v1.MsgCreateContent) (*v1.MsgCreateContentResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get account from curator address
	curator, err := sdk.AccAddressFromBech32(req.Curator)
	if err != nil {
		return nil, err // internal error
	}

	// insert content into content table
	err = k.ss.ContentTable().Insert(ctx, &contentv1.Content{
		Curator: curator,
		Hash:    req.Hash,
	})
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventCreateContent{
		Hash: req.Hash,
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgCreateContentResponse{
		Hash: req.Hash,
	}, nil
}
