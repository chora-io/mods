package keeper

import (
	"context"
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/group"

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

	// get sequence for key generation
	agentSequence, err := k.ss.AgentSequenceTable().Get(ctx)
	if err != nil {
		return nil, err // internal error
	}

	// generate account derivation key
	accountKey := make([]byte, 10)
	binary.LittleEndian.PutUint64(accountKey, agentSequence.Sequence)

	// create module account using account derivation key
	agentAccount, err := authtypes.NewModuleCredential(group.ModuleName, accountKey)
	if err != nil {
		return nil, err
	}

	// insert agent into agent table
	err = k.ss.AgentTable().Insert(ctx, &agentv1.Agent{
		Address:  agentAccount.Address(),
		Admin:    admin,
		Metadata: req.Metadata,
	})
	if err != nil {
		return nil, err // internal error
	}

	// update agent sequence table
	err = k.ss.AgentSequenceTable().Save(ctx, &agentv1.AgentSequence{
		Sequence: agentSequence.Sequence + 1,
	})
	if err != nil {
		return nil, err // internal error
	}

	// get string address from agent account
	address := sdk.AccAddress(agentAccount.Address()).String()

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
