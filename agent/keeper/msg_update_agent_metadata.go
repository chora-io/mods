package keeper

import (
	"context"

	"cosmossdk.io/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/chora-io/mods/agent/types/v1"
)

// UpdateAgentMetadata implements the Msg/UpdateAgentMetadata method.
func (k Keeper) UpdateAgentMetadata(ctx context.Context, req *v1.MsgUpdateAgentMetadata) (*v1.MsgUpdateAgentMetadataResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get account from admin address
	admin, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, err // internal error
	}

	// get agent from agent table
	agent, err := k.ss.AgentTable().Get(ctx, req.Address)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("agent with address %s: %s", req.Address, err)
		}
		return nil, err // internal error
	}

	// get account from account bytes
	agentAdmin := sdk.AccAddress(agent.Admin)

	// verify admin is agent admin
	if !agentAdmin.Equals(admin) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"admin %s: agent admin %s", admin, agentAdmin.String(),
		)
	}

	// set new agent metadata
	agent.Metadata = req.NewMetadata

	// update agent in agent table
	err = k.ss.AgentTable().Update(ctx, agent)
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventUpdateAgentMetadata{
		Address: agent.Address,
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgUpdateAgentMetadataResponse{
		Address: agent.Address,
	}, nil
}
