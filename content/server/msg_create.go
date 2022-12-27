package server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	contentv1 "github.com/choraio/mods/content/api/v1"
	v1 "github.com/choraio/mods/content/types/v1"
)

// Create implements Msg/Create.
func (s Server) Create(ctx context.Context, req *v1.MsgCreate) (*v1.MsgCreateResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get curator account from bech32
	curator, err := sdk.AccAddressFromBech32(req.Curator)
	if err != nil {
		return nil, err // internal error
	}

	// insert content into content table
	id, err := s.ss.ContentTable().InsertReturningID(ctx, &contentv1.Content{
		Curator: curator,
		Hash:    req.Hash,
	})
	if err != nil {
		return nil, err // internal error
	}

	// emit create event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventCreate{
		Id: id,
	}); err != nil {
		return nil, err // internal error
	}

	// return create response
	return &v1.MsgCreateResponse{
		Id: id,
	}, nil
}
