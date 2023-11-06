package keeper

import (
	"context"

	"cosmossdk.io/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

	geonodev1 "github.com/chora-io/mods/geonode/api/v1"
	v1 "github.com/chora-io/mods/geonode/types/v1"
	"github.com/chora-io/mods/geonode/utils"
)

// Nodes implements the Query/Nodes method.
func (k Keeper) Nodes(ctx context.Context, req *v1.QueryNodesRequest) (*v1.QueryNodesResponse, error) {

	// set index for table lookup
	index := geonodev1.NodeIdIndexKey{}

	// set pagination for table lookup
	pgnReq, err := utils.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err // internal error
	}

	// get nodes from node table
	it, err := k.ss.NodeTable().List(ctx, index, ormlist.Paginate(pgnReq))
	if err != nil {
		return nil, err // internal error
	}

	// set nodes for query response
	nodes := make([]*v1.QueryNodesResponse_Node, 0, 10)
	for it.Next() {
		v, err := it.Value()
		if err != nil {
			return nil, err // internal error
		}

		curator := sdk.AccAddress(v.Curator).String()

		node := v1.QueryNodesResponse_Node{
			Id:       v.Id,
			Curator:  curator,
			Metadata: v.Metadata,
		}

		nodes = append(nodes, &node)
	}

	// set pagination for query response
	pgnRes, err := utils.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err // internal error
	}

	// return query response
	return &v1.QueryNodesResponse{
		Nodes:      nodes,
		Pagination: pgnRes,
	}, nil
}
