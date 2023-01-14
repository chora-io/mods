package server

import (
	"context"

	voucherv1 "github.com/choraio/mods/voucher/api/v1"
	v1 "github.com/choraio/mods/voucher/types/v1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types/v2/math"
)

// BalancesByVoucher implements the Query/BalancesByVoucher method.
func (s Server) BalancesByVoucher(ctx context.Context, req *v1.QueryBalancesByVoucherRequest) (*v1.QueryBalancesByVoucherResponse, error) {

	// set index for table lookup
	index := voucherv1.BalanceIdAddressExpirationIndexKey{}.WithId(req.Id)

	// get balance from balance table
	it, err := s.ss.BalanceTable().List(ctx, index)
	if err != nil {
		return nil, err // internal error
	}

	// set initial addresses
	addrs := make([]string, 0, 10)

	// set initial address to decimal map
	addrToDec := make(map[string]math.Dec)

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

		// get account from address
		addr := sdk.AccAddress(v.Address).String()

		if addrToDec[addr].Equal(math.Dec{}) {
			addrs = append(addrs, addr)
			addrToDec[addr] = dec
		} else {
			addrToDec[addr], err = addrToDec[addr].Add(dec)
			if err != nil {
				return nil, err // internal error
			}
		}
	}

	// declare total amounts for query response
	totalAmounts := make([]*v1.QueryBalancesByVoucherResponse_TotalAmount, 0, 10)

	// set total amounts for query response
	for _, addr := range addrs {
		ta := &v1.QueryBalancesByVoucherResponse_TotalAmount{
			Address:     addr,
			TotalAmount: addrToDec[addr].String(),
		}
		totalAmounts = append(totalAmounts, ta)
	}

	// return query response
	return &v1.QueryBalancesByVoucherResponse{
		Id:           req.Id,
		TotalAmounts: totalAmounts,
	}, nil
}
