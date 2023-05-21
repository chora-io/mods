package server

import (
	"context"

	v1 "github.com/choraio/mods/validator/types/v1"
)

// MaxMissedBlocks implements the Query/MaxMissedBlocks method.
func (s Server) MaxMissedBlocks(ctx context.Context, _ *v1.QueryMaxMissedBlocksRequest) (*v1.QueryMaxMissedBlocksResponse, error) {

	// get max missed blocks from max missed blocks table
	maxMissedBlocks, err := s.ss.MaxMissedBlocksTable().Get(ctx)
	if err != nil {
		return nil, err // internal error
	}

	// return query response
	return &v1.QueryMaxMissedBlocksResponse{
		MaxMissedBlocks: maxMissedBlocks.MaxMissedBlocks,
	}, nil
}
