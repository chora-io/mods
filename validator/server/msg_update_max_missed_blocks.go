package server

import (
	"context"

	v1 "github.com/choraio/mods/validator/types/v1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// UpdateMaxMissedBlocks implements Msg/UpdateMaxMissedBlocks.
func (s Server) UpdateMaxMissedBlocks(ctx context.Context, req *v1.MsgUpdateMaxMissedBlocks) (*v1.MsgUpdateMaxMissedBlocksResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	authorityAddress := s.authority.String()
	if authorityAddress != req.Authority {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"authority: expected %s: received %s", authorityAddress, req.Authority,
		)
	}

	// get max missed blocks from max missed blocks table
	maxMissedBlocks, err := s.ss.MaxMissedBlocksTable().Get(ctx)
	if err != nil {
		return nil, err // internal error
	}

	// set max missed blocks to requested max missed blocks
	maxMissedBlocks.MaxMissedBlocks = req.MaxMissedBlocks

	// update max missed blocks in max missed blocks table
	err = s.ss.MaxMissedBlocksTable().Save(ctx, maxMissedBlocks)
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventUpdateMaxMissedBlocks{
		MaxMissedBlocks: req.MaxMissedBlocks,
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgUpdateMaxMissedBlocksResponse{
		MaxMissedBlocks: req.MaxMissedBlocks,
	}, nil
}
