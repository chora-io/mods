package server

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	voucherv1 "github.com/choraio/mods/voucher/api/v1"
	v1 "github.com/choraio/mods/voucher/types/v1"
)

type queryBalancesByVoucher struct {
	*baseSuite
	res *v1.QueryBalancesByVoucherResponse
	err error
}

func TestQueryBalancesByVoucher(t *testing.T) {
	gocuke.NewRunner(t, &queryBalancesByVoucher{}).
		Path("./features/query_balances_by_voucher.feature").
		Run()
}

func (s *queryBalancesByVoucher) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *queryBalancesByVoucher) Balance(a gocuke.DocString) {
	var balance voucherv1.Balance
	err := jsonpb.UnmarshalString(a.Content, &balance)
	require.NoError(s.t, err)

	err = s.srv.ss.BalanceTable().Insert(s.ctx, &voucherv1.Balance{
		Id:         balance.Id,
		Address:    balance.Address,
		Amount:     balance.Amount,
		Expiration: balance.Expiration,
	})
	require.NoError(s.t, err)
}

func (s *queryBalancesByVoucher) QueryBalancesByVoucher(a gocuke.DocString) {
	var req v1.QueryBalancesByVoucherRequest
	err := jsonpb.UnmarshalString(a.Content, &req)
	require.NoError(s.t, err)

	s.res, s.err = s.srv.BalancesByVoucher(s.ctx, &req)
}

func (s *queryBalancesByVoucher) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *queryBalancesByVoucher) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *queryBalancesByVoucher) ExpectResponse(a gocuke.DocString) {
	var expected v1.QueryBalancesByVoucherResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
