package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	agentv1 "github.com/chora-io/mods/agent/api/v1"
	v1 "github.com/chora-io/mods/agent/types/v1"
)

// CreateAgent implements Msg/CreateAgent.
func (k Keeper) CreateAgent(ctx context.Context, req *v1.MsgCreateAgent) (*v1.MsgCreateAgentResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get account from admin address
	admin, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, err // internal error
	}

	// TODO: generate address
	address := "address"

	// insert agent into agent table
	err = k.ss.AgentTable().Insert(ctx, &agentv1.Agent{
		Address:  address,
		Admin:    admin,
		Metadata: req.Metadata,
	})
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventCreateAgent{
		Address: address,
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgCreateAgentResponse{
		Address: address,
	}, nil
}
