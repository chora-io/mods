package server

import (
	"context"

	"cosmossdk.io/math"
	voucherv1 "github.com/choraio/mods/voucher/api/v1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/choraio/mods/voucher/types/v1"
)

// Balance implements the Query/Balance method.
func (s Server) Balance(ctx context.Context, req *v1.QueryBalanceRequest) (*v1.QueryBalanceResponse, error) {

	// get account from address
	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("address: %s", err)
	}

	// set index for table lookup
	index := voucherv1.BalanceIdAddressExpirationIndexKey{}.WithIdAddress(req.Id, address)

	// get balance from balance table
	it, err := s.ss.BalanceTable().List(ctx, index)
	if err != nil {
		return nil, err // internal error
	}

	// set initial amounts for query response
	amounts := make([]*v1.QueryBalanceResponse_Amount, 0, 10)

	// set initial total amount for query response
	totalAmount, err := math.LegacyNewDecFromStr("0")
	if err != nil {
		return nil, err // internal error
	}

	// set amounts and total amount for query response
	for it.Next() {
		v, err := it.Value()
		if err != nil {
			return nil, err // internal error
		}

		expiration := v.Expiration.AsTime()

		amount := v1.QueryBalanceResponse_Amount{
			Amount:     v.Amount,
			Expiration: &expiration,
		}

		amounts = append(amounts, &amount)

		dec, err := math.LegacyNewDecFromStr(v.Amount)
		if err != nil {
			return nil, err // internal error
		}

		totalAmount = totalAmount.Add(dec)
	}

	// return query response
	return &v1.QueryBalanceResponse{
		Id:          req.Id,
		Address:     req.Address,
		TotalAmount: totalAmount.String(),
		Amounts:     amounts,
	}, nil
}
