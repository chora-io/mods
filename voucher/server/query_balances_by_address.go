package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types/v2/math"

	voucherv1 "github.com/choraio/mods/voucher/api/v1"
	v1 "github.com/choraio/mods/voucher/types/v1"
	"github.com/choraio/mods/voucher/utils"
)

// BalancesByAddress implements the Query/BalancesByAddress method.
func (s Server) BalancesByAddress(ctx context.Context, req *v1.QueryBalancesByAddressRequest) (*v1.QueryBalancesByAddressResponse, error) {

	// set pagination for table lookup
	pgnReq, err := utils.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err // internal error
	}

	// get account from address
	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("address: %s", err)
	}

	// set index for table lookup
	index := voucherv1.BalanceAddressIndexKey{}.WithAddress(address)

	// get balance from balance table
	it, err := s.ss.BalanceTable().List(ctx, index, ormlist.Paginate(pgnReq))
	if err != nil {
		return nil, err // internal error
	}

	// set initial ids
	ids := make([]uint64, 0, 10)

	// set initial id to decimal map
	idToDec := make(map[uint64]math.Dec)

	// set amounts and total amount for query response
	for it.Next() {
		v, err := it.Value()
		if err != nil {
			return nil, err // internal error
		}

		dec, err := math.NewDecFromString(v.Amount)
		if err != nil {
			return nil, err // internal error
		}

		if idToDec[v.Id].Equal(math.Dec{}) {
			ids = append(ids, v.Id)
			idToDec[v.Id] = dec
		} else {
			idToDec[v.Id], err = idToDec[v.Id].Add(dec)
			if err != nil {
				return nil, err // internal error
			}
		}
	}

	// declare total amounts for query response
	totalAmounts := make([]*v1.QueryBalancesByAddressResponse_TotalAmount, 0, 10)

	// set total amounts for query response
	for _, id := range ids {
		ta := &v1.QueryBalancesByAddressResponse_TotalAmount{
			Id:          id,
			TotalAmount: idToDec[id].String(),
		}
		totalAmounts = append(totalAmounts, ta)
	}

	// set pagination for query response
	pgnRes, err := utils.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err // internal error
	}

	// return query response
	return &v1.QueryBalancesByAddressResponse{
		Address:      req.Address,
		TotalAmounts: totalAmounts,
		Pagination:   pgnRes,
	}, nil
}
