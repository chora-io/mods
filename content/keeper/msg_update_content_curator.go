package keeper

import (
	"context"

	"cosmossdk.io/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/chora-io/mods/content/types/v1"
)

// UpdateContentCurator implements the Msg/UpdateContentCurator method.
func (k Keeper) UpdateContentCurator(ctx context.Context, req *v1.MsgUpdateContentCurator) (*v1.MsgUpdateContentCuratorResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get account from curator address
	curator, err := sdk.AccAddressFromBech32(req.Curator)
	if err != nil {
		return nil, err // internal error
	}

	// get content from content table
	content, err := k.ss.ContentTable().Get(ctx, req.Hash)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("content with hash %s: %s", req.Hash, err)
		}
		return nil, err // internal error
	}

	// get account from account bytes
	contentCurator := sdk.AccAddress(content.Curator)

	// verify curator is content curator
	if !contentCurator.Equals(curator) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"curator %s: content curator %s", curator, contentCurator.String(),
		)
	}

	// get account from new curator address
	newCurator, err := sdk.AccAddressFromBech32(req.NewCurator)
	if err != nil {
		return nil, err // internal error
	}

	// set new curator
	content.Curator = newCurator

	// update content in content table
	err = k.ss.ContentTable().Update(ctx, content)
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventUpdateContentCurator{
		Hash: content.Hash,
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgUpdateContentCuratorResponse{
		Hash: content.Hash,
	}, nil
}
