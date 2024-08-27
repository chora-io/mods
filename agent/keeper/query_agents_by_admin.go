package keeper

import (
	"context"

	"cosmossdk.io/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	agentv1 "github.com/chora-io/mods/agent/api/v1"
	v1 "github.com/chora-io/mods/agent/types/v1"
	"github.com/chora-io/mods/agent/utils"
)

// AgentsByAdmin implements the Query/AgentsByAdmin method.
func (k Keeper) AgentsByAdmin(ctx context.Context, req *v1.QueryAgentsByAdminRequest) (*v1.QueryAgentsByAdminResponse, error) {

	// get account from admin address
	admin, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("admin: %s", err)
	}

	// set index for table lookup
	index := agentv1.AgentAdminIndexKey{}.WithAdmin(admin)

	// set pagination for table lookup
	pgnReq, err := utils.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err // internal error
	}

	// get agents by admin from agent table
	it, err := k.ss.AgentTable().List(ctx, index, ormlist.Paginate(pgnReq))
	if err != nil {
		return nil, err // internal error
	}

	// set agents for query response
	agents := make([]*v1.QueryAgentsByAdminResponse_Agent, 0, 10)
	for it.Next() {
		v, err := it.Value()
		if err != nil {
			return nil, err // internal error
		}

		address := sdk.AccAddress(v.Address).String()

		agent := v1.QueryAgentsByAdminResponse_Agent{
			Address:  address,
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
	return &v1.QueryAgentsByAdminResponse{
		Admin:      admin.String(),
		Agents:     agents,
		Pagination: pgnRes,
	}, nil
}
