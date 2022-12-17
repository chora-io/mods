package server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/choraio/mods/example/api/v1"
	"github.com/choraio/mods/example/types/v1"
)

// CreateContent implements Msg/CreateContent.
func (s Server) CreateContent(ctx context.Context, req *v1.MsgCreateContent) (*v1.MsgCreateContentResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get creator account from bech32
	creator, err := sdk.AccAddressFromBech32(req.Creator)
	if err != nil {
		return nil, err // internal error
	}

	// insert content into content table
	id, err := s.ss.ContentTable().InsertReturningID(ctx, &examplev1.Content{
		Creator: creator,
		Hash:    req.Hash,
	})
	if err != nil {
		return nil, err // internal error
	}

	// emit create content event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventCreateContent{
		Id: id,
	}); err != nil {
		return nil, err // internal error
	}

	// return create content response
	return &v1.MsgCreateContentResponse{
		Id: id,
	}, nil
}
