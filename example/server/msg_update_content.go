package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/choraio/mods/example/types/v1"
)

// UpdateContent implements the Msg/UpdateContent method.
func (s Server) UpdateContent(ctx context.Context, req *v1.MsgUpdateContent) (*v1.MsgUpdateContentResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get creator account from bech32
	creator, err := sdk.AccAddressFromBech32(req.Creator)
	if err != nil {
		return nil, err // internal error
	}

	// get content from content table
	content, err := s.ss.ContentTable().Get(ctx, req.Id)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf(
				"content with id %s: %s", req.Id, err,
			)
		}
		return nil, err // internal error
	}

	// convert content creator to account
	contentCreator := sdk.AccAddress(content.Creator)

	// verify creator is content creator
	if !contentCreator.Equals(creator) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"content creator %s: %s", contentCreator.String(), err,
		)
	}

	// set new content hash
	content.Hash = req.NewHash

	// update content in content table
	err = s.ss.ContentTable().Update(ctx, content)
	if err != nil {
		return nil, err // internal error
	}

	// emit update content event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventUpdateContent{
		Id: content.Id,
	}); err != nil {
		return nil, err // internal error
	}

	// return update content response
	return &v1.MsgUpdateContentResponse{
		Id: content.Id,
	}, nil
}
