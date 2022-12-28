package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types/v2/ormutil"

	geonodev1 "github.com/choraio/mods/geonode/api/v1"
	v1 "github.com/choraio/mods/geonode/types/v1"
)

// NodesByCurator implements the Query/NodesByCurator method.
func (s Server) NodesByCurator(ctx context.Context, req *v1.QueryNodesByCuratorRequest) (*v1.QueryNodesByCuratorResponse, error) {

	// convert curator to account
	curator, err := sdk.AccAddressFromBech32(req.Curator)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("curator: %s", err)
	}

	// set the index for table lookup
	index := geonodev1.NodeCuratorIndexKey{}.WithCurator(curator)

	// set the pagination for table lookup
	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err // internal error
	}

	// get iterator for node by curator from node table
	it, err := s.ss.NodeTable().List(ctx, index, ormlist.Paginate(pg))
	if err != nil {
		return nil, err // internal error
	}

	// set nodes for query nodes by curator response
	nodes := make([]*v1.QueryNodesByCuratorResponse_Node, 0, 10)
	for it.Next() {
		v, err := it.Value()
		if err != nil {
			return nil, err // internal error
		}
		n := v1.QueryNodesByCuratorResponse_Node{
			Id:       v.Id,
			Metadata: v.Metadata,
		}
		nodes = append(nodes, &n)
	}

	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err // internal error
	}

	// return query nodes by curator response
	return &v1.QueryNodesByCuratorResponse{
		Curator:    curator.String(),
		Nodes:      nodes,
		Pagination: pr,
	}, nil
}
