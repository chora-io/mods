package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/choraio/mods/example/types/v1"
)

// DeleteContent implements Msg/DeleteContent.
func (s Server) DeleteContent(ctx context.Context, req *v1.MsgDeleteContent) (*v1.MsgDeleteContentResponse, error) {
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

	// delete content from content table
	err = s.ss.ContentTable().Delete(ctx, content)
	if err != nil {
		return nil, err // internal error
	}

	// emit delete content event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventDeleteContent{
		Id: content.Id,
	}); err != nil {
		return nil, err // internal error
	}

	// return delete content response
	return &v1.MsgDeleteContentResponse{
		Id: content.Id,
	}, nil
}
