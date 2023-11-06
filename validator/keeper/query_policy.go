package keeper

import (
	"context"

	v1 "github.com/chora-io/mods/validator/types/v1"
)

// Policy implements the Query/Policy method.
func (k Keeper) Policy(ctx context.Context, _ *v1.QueryPolicyRequest) (*v1.QueryPolicyResponse, error) {

	// get policy from singleton table
	policy, err := k.ss.PolicyTable().Get(ctx)
	if err != nil {
		return nil, err // internal error
	}

	// return query response
	return &v1.QueryPolicyResponse{
		SignedBlocksWindow: policy.SignedBlocksWindow,
		MinSignedPerWindow: policy.MinSignedPerWindow,
	}, nil
}
