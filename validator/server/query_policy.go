package server

import (
	"context"

	v1 "github.com/choraio/mods/validator/types/v1"
)

// Policy implements the Query/Policy method.
func (s Server) Policy(ctx context.Context, _ *v1.QueryPolicyRequest) (*v1.QueryPolicyResponse, error) {

	// get policy from singleton table
	policy, err := s.ss.PolicyTable().Get(ctx)
	if err != nil {
		return nil, err // internal error
	}

	// return query response
	return &v1.QueryPolicyResponse{
		SignedBlocksWindow: policy.SignedBlocksWindow,
		MinSignedPerWindow: policy.MinSignedPerWindow,
	}, nil
}
