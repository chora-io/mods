package keeper

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	authorityv1 "github.com/chora-io/mods/authority/api/v1"
	v1 "github.com/chora-io/mods/authority/types/v1"
	"github.com/chora-io/mods/authority/utils"
)

type msgUpdate struct {
	*baseSuite
	res *v1.MsgUpdateResponse
	err error
}

func TestMsgUpdate(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdate{}).
		Path("./features/msg_update.feature").
		Run()
}

func (s *msgUpdate) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgUpdate) Authority(a gocuke.DocString) {
	var authority authorityv1.Authority
	err := jsonpb.UnmarshalString(a.Content, &authority)
	require.NoError(s.t, err)

	err = s.k.ss.AuthorityTable().Save(s.sdkCtx, &authorityv1.Authority{
		Address: authority.Address,
	})
	require.NoError(s.t, err)
}

func (s *msgUpdate) MsgUpdate(a gocuke.DocString) {
	var msg v1.MsgUpdate
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.Update(s.sdkCtx, &msg)
}

func (s *msgUpdate) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdate) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgUpdate) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgUpdateResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgUpdate) ExpectStateAuthority(a gocuke.DocString) {
	var expected authorityv1.Authority
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.k.ss.AuthorityTable().Get(s.sdkCtx)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Address, actual.Address)
}

func (s *msgUpdate) ExpectEventUpdate(a gocuke.DocString) {
	var expected v1.EventUpdate
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
