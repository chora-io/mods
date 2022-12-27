package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/choraio/mods/content/types/v1"
)

// Delete implements Msg/Delete.
func (s Server) Delete(ctx context.Context, req *v1.MsgDelete) (*v1.MsgDeleteResponse, error) {
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

	// delete content from content table
	err = s.ss.ContentTable().Delete(ctx, content)
	if err != nil {
		return nil, err // internal error
	}

	// emit delete event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventDelete{
		Id: content.Id,
	}); err != nil {
		return nil, err // internal error
	}

	// return delete response
	return &v1.MsgDeleteResponse{
		Id: content.Id,
	}, nil
}
