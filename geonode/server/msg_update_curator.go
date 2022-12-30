package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/choraio/mods/geonode/types/v1"
)

// UpdateCurator implements the Msg/UpdateCurator method.
func (s Server) UpdateCurator(ctx context.Context, req *v1.MsgUpdateCurator) (*v1.MsgUpdateCuratorResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get account from curator address
	curator, err := sdk.AccAddressFromBech32(req.Curator)
	if err != nil {
		return nil, err // internal error
	}

	// get node from node table
	node, err := s.ss.NodeTable().Get(ctx, req.Id)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("node with id %d: %s", req.Id, err)
		}
		return nil, err // internal error
	}

	// get account from account bytes
	nodeCurator := sdk.AccAddress(node.Curator)

	// verify curator is node curator
	if !nodeCurator.Equals(curator) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"curator %s: node curator %s", curator, nodeCurator.String(),
		)
	}

	// get account from new curator address
	newCurator, err := sdk.AccAddressFromBech32(req.NewCurator)
	if err != nil {
		return nil, err // internal error
	}

	// set new curator
	node.Curator = newCurator

	// update node in node table
	err = s.ss.NodeTable().Update(ctx, node)
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventUpdateCurator{
		Id: node.Id,
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgUpdateCuratorResponse{
		Id: node.Id,
	}, nil
}
