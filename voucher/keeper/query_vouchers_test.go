package keeper

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	voucherv1 "github.com/chora-io/mods/voucher/api/v1"
	v1 "github.com/chora-io/mods/voucher/types/v1"
)

type queryVouchers struct {
	*baseSuite
	res *v1.QueryVouchersResponse
	err error
}

func TestQueryVouchers(t *testing.T) {
	gocuke.NewRunner(t, &queryVouchers{}).
		Path("./features/query_vouchers.feature").
		Run()
}

func (s *queryVouchers) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *queryVouchers) Voucher(a gocuke.DocString) {
	var voucher voucherv1.Voucher
	err := jsonpb.UnmarshalString(a.Content, &voucher)
	require.NoError(s.t, err)

	id, err := s.k.ss.VoucherTable().InsertReturningId(s.sdkCtx, &voucherv1.Voucher{
		Issuer:   voucher.Issuer,
		Metadata: voucher.Metadata,
	})
	require.NoError(s.t, err)
	require.Equal(s.t, voucher.Id, id)
}

func (s *queryVouchers) QueryVouchers(a gocuke.DocString) {
	var req v1.QueryVouchersRequest
	err := jsonpb.UnmarshalString(a.Content, &req)
	require.NoError(s.t, err)

	s.res, s.err = s.k.Vouchers(s.sdkCtx, &req)
}

func (s *queryVouchers) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *queryVouchers) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *queryVouchers) ExpectResponse(a gocuke.DocString) {
	var expected v1.QueryVouchersResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
