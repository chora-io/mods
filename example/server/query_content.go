package server

import (
	"context"

	types "github.com/choraio/mods/example/types/v1"
)

// Content implements the Query/Content method.
func (s Server) Content(ctx context.Context, req *types.QueryContentRequest) (*types.QueryContentResponse, error) {
	panic("not implemented")
}
