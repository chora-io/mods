package server

import (
	"context"

	types "github.com/choraio/mods/example/types/v1"
)

// DeleteContent implements the Msg/DeleteContent method.
func (s Server) DeleteContent(ctx context.Context, msg *types.MsgDeleteContent) (*types.MsgDeleteContentResponse, error) {
	panic("not implemented")
}
