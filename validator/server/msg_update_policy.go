package server

import (
	"context"

	v1 "github.com/choraio/mods/validator/types/v1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// UpdatePolicy implements Msg/UpdatePolicy.
func (s Server) UpdatePolicy(ctx context.Context, req *v1.MsgUpdatePolicy) (*v1.MsgUpdatePolicyResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	authorityAddress := s.authority.String()
	if authorityAddress != req.Authority {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"authority: expected %s: received %s", authorityAddress, req.Authority,
		)
	}

	// get max missed blocks from max missed blocks table
	policy, err := s.ss.PolicyTable().Get(ctx)
	if err != nil {
		return nil, err // internal error
	}

	// set signed blocks window
	policy.SignedBlocksWindow = req.SignedBlocksWindow

	// set min signed per window
	policy.MinSignedPerWindow = req.MinSignedPerWindow

	// update max missed blocks in max missed blocks table
	err = s.ss.PolicyTable().Save(ctx, policy)
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventUpdatePolicy{
		SignedBlocksWindow: req.SignedBlocksWindow,
		MinSignedPerWindow: req.MinSignedPerWindow,
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgUpdatePolicyResponse{
		SignedBlocksWindow: req.SignedBlocksWindow,
		MinSignedPerWindow: req.MinSignedPerWindow,
	}, nil
}
