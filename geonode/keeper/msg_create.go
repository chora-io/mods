package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	geonodev1 "github.com/chora-io/mods/geonode/api/v1"
	v1 "github.com/chora-io/mods/geonode/types/v1"
)

// Create implements Msg/Create.
func (k Keeper) Create(ctx context.Context, req *v1.MsgCreate) (*v1.MsgCreateResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get account from curator address
	curator, err := sdk.AccAddressFromBech32(req.Curator)
	if err != nil {
		return nil, err // internal error
	}

	// insert node into node table
	id, err := k.ss.NodeTable().InsertReturningId(ctx, &geonodev1.Node{
		Curator:  curator,
		Metadata: req.Metadata,
	})
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventCreate{
		Id: id,
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgCreateResponse{
		Id: id,
	}, nil
}
