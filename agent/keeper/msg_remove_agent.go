package keeper

import (
	"context"

	"cosmossdk.io/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/chora-io/mods/agent/types/v1"
)

// RemoveAgent implements Msg/RemoveAgent.
func (k Keeper) RemoveAgent(ctx context.Context, req *v1.MsgRemoveAgent) (*v1.MsgRemoveAgentResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get account from agent address
	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err // internal error
	}

	// get account from admin address
	admin, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, err // internal error
	}

	// get agent from agent table
	agent, err := k.ss.AgentTable().Get(ctx, address)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("agent with address %s: %s", req.Address, err)
		}
		return nil, err // internal error
	}

	// get admin account from agent admin
	agentAdmin := sdk.AccAddress(agent.Admin)

	// verify admin is agent admin
	if !agentAdmin.Equals(admin) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"admin %s: agent admin %s", admin, agentAdmin,
		)
	}

	// delete agent from agent table
	err = k.ss.AgentTable().Delete(ctx, agent)
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventRemoveAgent{
		Address: req.Address,
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgRemoveAgentResponse{
		Address: req.Address,
	}, nil
}
