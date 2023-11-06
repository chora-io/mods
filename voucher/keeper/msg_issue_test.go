package keeper

import (
	"testing"
	"time"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	voucherv1 "github.com/chora-io/mods/voucher/api/v1"
	v1 "github.com/chora-io/mods/voucher/types/v1"
	"github.com/chora-io/mods/voucher/utils"
)

type msgIssue struct {
	*baseSuite
	res *v1.MsgIssueResponse
	err error
}

func TestMsgIssue(t *testing.T) {
	gocuke.NewRunner(t, &msgIssue{}).
		Path("./features/msg_issue.feature").
		Run()
}

func (s *msgIssue) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgIssue) BlockTime(a string) {
	blockTime, err := time.Parse(time.RFC3339, a)
	require.NoError(s.t, err)

	s.sdkCtx = s.sdkCtx.WithBlockTime(blockTime)
}

func (s *msgIssue) Voucher(a gocuke.DocString) {
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

func (s *msgIssue) Balance(a gocuke.DocString) {
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

func (s *msgIssue) MsgIssue(a gocuke.DocString) {
	var msg v1.MsgIssue
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.Issue(s.sdkCtx, &msg)
}

func (s *msgIssue) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgIssue) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgIssue) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgIssueResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgIssue) ExpectStateBalance(a gocuke.DocString) {
	var expected voucherv1.Balance
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.k.ss.BalanceTable().Get(s.sdkCtx, expected.Id, expected.Address, expected.Expiration)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Id, actual.Id)
	require.Equal(s.t, expected.Address, actual.Address)
	require.Equal(s.t, expected.Amount, actual.Amount)
	require.Equal(s.t, expected.Expiration, actual.Expiration)
}

func (s *msgIssue) ExpectEventIssue(a gocuke.DocString) {
	var expected v1.EventIssue
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
