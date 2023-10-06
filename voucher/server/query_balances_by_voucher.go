package server

import (
	"context"

	"cosmossdk.io/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"cosmossdk.io/math"

	voucherv1 "github.com/choraio/mods/voucher/api/v1"
	v1 "github.com/choraio/mods/voucher/types/v1"
	"github.com/choraio/mods/voucher/utils"
)

// BalancesByVoucher implements the Query/BalancesByVoucher method.
func (s Server) BalancesByVoucher(ctx context.Context, req *v1.QueryBalancesByVoucherRequest) (*v1.QueryBalancesByVoucherResponse, error) {

	// set index for table lookup
	index := voucherv1.BalanceIdAddressExpirationIndexKey{}.WithId(req.Id)

	// set pagination for table lookup
	pgnReq, err := utils.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err // internal error
	}

	// get balance from balance table
	it, err := s.ss.BalanceTable().List(ctx, index, ormlist.Paginate(pgnReq))
	if err != nil {
		return nil, err // internal error
	}

	// set initial addresses
	addrs := make([]string, 0, 10)

	// set initial address to decimal map
	addrToDec := make(map[string]math.LegacyDec)

	// set amounts and total amount for query response
	for it.Next() {
		v, err := it.Value()
		if err != nil {
			return nil, err // internal error
		}

		dec, err := math.LegacyNewDecFromStr(v.Amount)
		if err != nil {
			return nil, err // internal error
		}

		// get account from address
		addr := sdk.AccAddress(v.Address).String()

		if addrToDec[addr].IsNil() {
			addrs = append(addrs, addr)
			addrToDec[addr] = dec
		} else {
			addrToDec[addr] = addrToDec[addr].Add(dec)
		}
	}

	// declare total amounts for query response
	totalAmounts := make([]*v1.QueryBalancesByVoucherResponse_TotalAmount, 0, 10)

	// set total amounts for query response
	for _, addr := range addrs {
		totalAmount := &v1.QueryBalancesByVoucherResponse_TotalAmount{
			Address:     addr,
			TotalAmount: addrToDec[addr].String(),
		}

		totalAmounts = append(totalAmounts, totalAmount)
	}

	// set pagination for query response
	pgnRes, err := utils.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err // internal error
	}

	// return query response
	return &v1.QueryBalancesByVoucherResponse{
		Id:           req.Id,
		TotalAmounts: totalAmounts,
		Pagination:   pgnRes,
	}, nil
}
