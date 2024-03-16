package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/chora-io/mods/validator/types/v1"
)

// UpdatePolicy implements Msg/UpdatePolicy.
func (k Keeper) UpdatePolicy(ctx context.Context, req *v1.MsgUpdatePolicy) (*v1.MsgUpdatePolicyResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	authorityAddress := k.authority.String()
	if authorityAddress != req.Authority {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"authority: expected %s: received %s", authorityAddress, req.Authority,
		)
	}

	// get validator policy from policy table
	policy, err := k.ss.PolicyTable().Get(ctx)
	if err != nil {
		return nil, err // internal error
	}

	// set signed blocks window
	policy.SignedBlocksWindow = req.SignedBlocksWindow

	// set min signed per window
	policy.MinSignedPerWindow = req.MinSignedPerWindow

	// update validator policy in policy table
	err = k.ss.PolicyTable().Save(ctx, policy)
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
