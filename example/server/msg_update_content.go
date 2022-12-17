package server

import (
	"context"

	types "github.com/choraio/mods/example/types/v1"
)

// UpdateContent implements the Msg/UpdateContent method.
func (s Server) UpdateContent(ctx context.Context, msg *types.MsgUpdateContent) (*types.MsgUpdateContentResponse, error) {
	panic("not implemented")
}
