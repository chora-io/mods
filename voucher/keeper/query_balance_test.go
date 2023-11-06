package keeper

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	voucherv1 "github.com/chora-io/mods/voucher/api/v1"
	v1 "github.com/chora-io/mods/voucher/types/v1"
)

type queryBalance struct {
	*baseSuite
	res *v1.QueryBalanceResponse
	err error
}

func TestQueryBalance(t *testing.T) {
	gocuke.NewRunner(t, &queryBalance{}).
		Path("./features/query_balance.feature").
		Run()
}

func (s *queryBalance) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *queryBalance) Balance(a gocuke.DocString) {
	var balance voucherv1.Balance
	err := jsonpb.UnmarshalString(a.Content, &balance)
	require.NoError(s.t, err)

	err = s.k.ss.BalanceTable().Insert(s.sdkCtx, &voucherv1.Balance{
		Id:         balance.Id,
		Address:    balance.Address,
		Amount:     balance.Amount,
		Expiration: balance.Expiration,
	})
	require.NoError(s.t, err)
}

func (s *queryBalance) QueryBalance(a gocuke.DocString) {
	var req v1.QueryBalanceRequest
	err := jsonpb.UnmarshalString(a.Content, &req)
	require.NoError(s.t, err)

	s.res, s.err = s.k.Balance(s.sdkCtx, &req)
}

func (s *queryBalance) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *queryBalance) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *queryBalance) ExpectResponse(a gocuke.DocString) {
	var expected v1.QueryBalanceResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
