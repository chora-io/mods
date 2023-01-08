package server

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	contentv1 "github.com/choraio/mods/voucher/api/v1"
	v1 "github.com/choraio/mods/voucher/types/v1"
	"github.com/choraio/mods/voucher/utils"
)

type msgUpdateIssuer struct {
	*baseSuite
	res *v1.MsgUpdateIssuerResponse
	err error
}

func TestMsgUpdateIssuer(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateIssuer{}).
		Path("./features/msg_update_issuer.feature").
		Run()
}

func (s *msgUpdateIssuer) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgUpdateIssuer) Voucher(a gocuke.DocString) {
	var voucher contentv1.Voucher
	err := jsonpb.UnmarshalString(a.Content, &voucher)
	require.NoError(s.t, err)

	id, err := s.srv.ss.VoucherTable().InsertReturningID(s.ctx, &contentv1.Voucher{
		Issuer:   voucher.Issuer,
		Metadata: voucher.Metadata,
	})
	require.NoError(s.t, err)
	require.Equal(s.t, voucher.Id, id)
}

func (s *msgUpdateIssuer) MsgUpdateIssuer(a gocuke.DocString) {
	var msg v1.MsgUpdateIssuer
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.srv.UpdateIssuer(s.ctx, &msg)
}

func (s *msgUpdateIssuer) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateIssuer) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgUpdateIssuer) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgUpdateIssuerResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgUpdateIssuer) ExpectStateVoucher(a gocuke.DocString) {
	var expected contentv1.Voucher
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.srv.ss.VoucherTable().Get(s.ctx, expected.Id)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Id, actual.Id)
	require.Equal(s.t, expected.Issuer, actual.Issuer)
	require.Equal(s.t, expected.Metadata, actual.Metadata)
}

func (s *msgUpdateIssuer) ExpectEventUpdateIssuer(a gocuke.DocString) {
	var expected v1.EventUpdateIssuer
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
