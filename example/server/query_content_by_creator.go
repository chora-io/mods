package server

import (
	"context"

	types "github.com/choraio/mods/example/types/v1"
)

// ContentByCreator implements the Query/ContentByCreator method.
func (s Server) ContentByCreator(ctx context.Context, req *types.QueryContentByCreatorRequest) (*types.QueryContentByCreatorResponse, error) {
	panic("not implemented")
}
