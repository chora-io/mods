package keeper

import (
	"context"

	"cosmossdk.io/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/chora-io/mods/agent/types/v1"
)

// Agent implements the Query/Agent method.
func (k Keeper) Agent(ctx context.Context, req *v1.QueryAgentRequest) (*v1.QueryAgentResponse, error) {

	// get agent from agent table
	agent, err := k.ss.AgentTable().Get(ctx, req.Address)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("agent with address %s", req.Address)
		}
		return nil, err // internal error
	}

	// get account from account bytes
	admin := sdk.AccAddress(agent.Admin)

	// return query response
	return &v1.QueryAgentResponse{
		Address:  agent.Address,
		Admin:    admin.String(),
		Metadata: agent.Metadata,
	}, nil
}
