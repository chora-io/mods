package server

import (
	"context"

	"cosmossdk.io/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	geonodev1 "github.com/choraio/mods/geonode/api/v1"
	v1 "github.com/choraio/mods/geonode/types/v1"
	"github.com/choraio/mods/geonode/utils"
)

// NodesByCurator implements the Query/NodesByCurator method.
func (s Server) NodesByCurator(ctx context.Context, req *v1.QueryNodesByCuratorRequest) (*v1.QueryNodesByCuratorResponse, error) {

	// get account from curator address
	curator, err := sdk.AccAddressFromBech32(req.Curator)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("curator: %s", err)
	}

	// set index for table lookup
	index := geonodev1.NodeCuratorIndexKey{}.WithCurator(curator)

	// set pagination for table lookup
	pgnReq, err := utils.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err // internal error
	}

	// get nodes by curator from node table
	it, err := s.ss.NodeTable().List(ctx, index, ormlist.Paginate(pgnReq))
	if err != nil {
		return nil, err // internal error
	}

	// set nodes for query response
	nodes := make([]*v1.QueryNodesByCuratorResponse_Node, 0, 10)
	for it.Next() {
		v, err := it.Value()
		if err != nil {
			return nil, err // internal error
		}

		node := v1.QueryNodesByCuratorResponse_Node{
			Id:       v.Id,
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
	return &v1.QueryNodesByCuratorResponse{
		Curator:    curator.String(),
		Nodes:      nodes,
		Pagination: pgnRes,
	}, nil
}
