package server

import (
	"context"

	types "github.com/choraio/mods/example/types/v1"
)

// CreateContent implements the Msg/CreateContent method.
func (s Server) CreateContent(ctx context.Context, msg *types.MsgCreateContent) (*types.MsgCreateContentResponse, error) {
	panic("not implemented")
}
