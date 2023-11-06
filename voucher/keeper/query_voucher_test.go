package keeper

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	voucherv1 "github.com/chora-io/mods/voucher/api/v1"
	v1 "github.com/chora-io/mods/voucher/types/v1"
)

type queryVoucher struct {
	*baseSuite
	res *v1.QueryVoucherResponse
	err error
}

func TestQueryVoucher(t *testing.T) {
	gocuke.NewRunner(t, &queryVoucher{}).
		Path("./features/query_voucher.feature").
		Run()
}

func (s *queryVoucher) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *queryVoucher) Voucher(a gocuke.DocString) {
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

func (s *queryVoucher) QueryVoucher(a gocuke.DocString) {
	var req v1.QueryVoucherRequest
	err := jsonpb.UnmarshalString(a.Content, &req)
	require.NoError(s.t, err)

	s.res, s.err = s.k.Voucher(s.sdkCtx, &req)
}

func (s *queryVoucher) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *queryVoucher) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *queryVoucher) ExpectResponse(a gocuke.DocString) {
	var expected v1.QueryVoucherResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
