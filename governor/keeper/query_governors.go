package keeper

import (
	"context"

	"cosmossdk.io/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

	governorv1 "github.com/chora-io/mods/governor/api/v1"
	v1 "github.com/chora-io/mods/governor/types/v1"
	"github.com/chora-io/mods/governor/utils"
)

// Governors implements the Query/Governors method.
func (k Keeper) Governors(ctx context.Context, req *v1.QueryGovernorsRequest) (*v1.QueryGovernorsResponse, error) {

	// set the index for table lookup
	index := governorv1.GovernorAddressIndexKey{}

	// set the pagination for table lookup
	pgnReq, err := utils.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err // internal error
	}

	// get governors from governor table
	it, err := k.ss.GovernorTable().List(ctx, index, ormlist.Paginate(pgnReq))
	if err != nil {
		return nil, err // internal error
	}

	// set governors for query response
	governors := make([]*v1.QueryGovernorsResponse_Governor, 0, 10)
	for it.Next() {
		v, err := it.Value()
		if err != nil {
			return nil, err // internal error
		}

		// get governor account from address
		address := sdk.AccAddress(v.Address)

		governor := v1.QueryGovernorsResponse_Governor{
			Address:  address.String(),
			Metadata: v.Metadata,
		}

		governors = append(governors, &governor)
	}

	// set the pagination for query response
	pgnRes, err := utils.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err // internal error
	}

	// return query response
	return &v1.QueryGovernorsResponse{
		Governors:  governors,
		Pagination: pgnRes,
	}, nil
}
