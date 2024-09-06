package keeper

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	adminv1 "github.com/chora-io/mods/admin/api/v1"
	v1 "github.com/chora-io/mods/admin/types/v1"
	"github.com/chora-io/mods/admin/utils"
)

type msgUpdateAdmin struct {
	*baseSuite
	res *v1.MsgUpdateAdminResponse
	err error
}

func TestMsgUpdateAdmin(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateAdmin{}).
		Path("./msg_update_admin.feature").
		Run()
}

func (s *msgUpdateAdmin) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgUpdateAdmin) Admin(a gocuke.DocString) {
	var admin adminv1.Admin
	err := jsonpb.UnmarshalString(a.Content, &admin)
	require.NoError(s.t, err)

	err = s.k.ss.AdminTable().Save(s.sdkCtx, &adminv1.Admin{
		Admin: admin.Admin,
	})
	require.NoError(s.t, err)
}

func (s *msgUpdateAdmin) MsgUpdateAdmin(a gocuke.DocString) {
	var msg v1.MsgUpdateAdmin
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.UpdateAdmin(s.sdkCtx, &msg)
}

func (s *msgUpdateAdmin) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateAdmin) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgUpdateAdmin) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgUpdateAdminResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgUpdateAdmin) ExpectStateAdmin(a gocuke.DocString) {
	var expected adminv1.Admin
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.k.ss.AdminTable().Get(s.sdkCtx)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Admin, actual.Admin)
}

func (s *msgUpdateAdmin) ExpectEventUpdate(a gocuke.DocString) {
	var expected v1.EventUpdateAdmin
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
