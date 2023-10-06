package server

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	voucherv1 "github.com/choraio/mods/voucher/api/v1"
	v1 "github.com/choraio/mods/voucher/types/v1"
)

type queryBalancesByAddress struct {
	*baseSuite
	res *v1.QueryBalancesByAddressResponse
	err error
}

func TestQueryBalancesByAddress(t *testing.T) {
	gocuke.NewRunner(t, &queryBalancesByAddress{}).
		Path("./features/query_balances_by_address.feature").
		Run()
}

func (s *queryBalancesByAddress) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *queryBalancesByAddress) Balance(a gocuke.DocString) {
	var balance voucherv1.Balance
	err := jsonpb.UnmarshalString(a.Content, &balance)
	require.NoError(s.t, err)

	err = s.srv.ss.BalanceTable().Insert(s.sdkCtx, &voucherv1.Balance{
		Id:         balance.Id,
		Address:    balance.Address,
		Amount:     balance.Amount,
		Expiration: balance.Expiration,
	})
	require.NoError(s.t, err)
}

func (s *queryBalancesByAddress) QueryBalancesByAddress(a gocuke.DocString) {
	var req v1.QueryBalancesByAddressRequest
	err := jsonpb.UnmarshalString(a.Content, &req)
	require.NoError(s.t, err)

	s.res, s.err = s.srv.BalancesByAddress(s.sdkCtx, &req)
}

func (s *queryBalancesByAddress) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *queryBalancesByAddress) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *queryBalancesByAddress) ExpectResponse(a gocuke.DocString) {
	var expected v1.QueryBalancesByAddressResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
