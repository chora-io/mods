package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/choraio/mods/content/types/v1"
)

// Update implements the Msg/Update method.
func (s Server) Update(ctx context.Context, req *v1.MsgUpdate) (*v1.MsgUpdateResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get curator account from bech32
	curator, err := sdk.AccAddressFromBech32(req.Curator)
	if err != nil {
		return nil, err // internal error
	}

	// get content from content table
	content, err := s.ss.ContentTable().Get(ctx, req.Id)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf(
				"content with id %d: %s", req.Id, err,
			)
		}
		return nil, err // internal error
	}

	// convert content curator to account
	contentCurator := sdk.AccAddress(content.Curator)

	// verify curator is content curator
	if !contentCurator.Equals(curator) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"curator %s: content curator %s", curator, contentCurator.String(),
		)
	}

	// set new content metadata
	content.Metadata = req.NewMetadata

	// update content in content table
	err = s.ss.ContentTable().Update(ctx, content)
	if err != nil {
		return nil, err // internal error
	}

	// emit update event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventUpdate{
		Id: content.Id,
	}); err != nil {
		return nil, err // internal error
	}

	// return update response
	return &v1.MsgUpdateResponse{
		Id: content.Id,
	}, nil
}
