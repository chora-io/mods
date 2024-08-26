package keeper

import (
	"context"

	"cosmossdk.io/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

	agentv1 "github.com/chora-io/mods/agent/api/v1"
	v1 "github.com/chora-io/mods/agent/types/v1"
	"github.com/chora-io/mods/agent/utils"
)

// Agents implements the Query/Agents method.
func (k Keeper) Agents(ctx context.Context, req *v1.QueryAgentsRequest) (*v1.QueryAgentsResponse, error) {

	// set index for table lookup
	index := agentv1.AgentAddressIndexKey{}

	// set pagination for table lookup
	pgnReq, err := utils.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err // internal error
	}

	// get agents from agent table
	it, err := k.ss.AgentTable().List(ctx, index, ormlist.Paginate(pgnReq))
	if err != nil {
		return nil, err // internal error
	}

	// set agents for query response
	agents := make([]*v1.QueryAgentsResponse_Agent, 0, 10)
	for it.Next() {
		v, err := it.Value()
		if err != nil {
			return nil, err // internal error
		}

		admin := sdk.AccAddress(v.Admin).String()

		agent := v1.QueryAgentsResponse_Agent{
			Address:  v.Address,
			Admin:    admin,
			Metadata: v.Metadata,
		}

		agents = append(agents, &agent)
	}

	// set pagination for query response
	pgnRes, err := utils.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err // internal error
	}

	// return query response
	return &v1.QueryAgentsResponse{
		Agents:     agents,
		Pagination: pgnRes,
	}, nil
}
