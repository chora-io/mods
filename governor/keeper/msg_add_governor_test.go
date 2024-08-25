package keeper

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	governorv1 "github.com/chora-io/mods/governor/api/v1"
	v1 "github.com/chora-io/mods/governor/types/v1"
	"github.com/chora-io/mods/governor/utils"
)

type msgAddGovernor struct {
	*baseSuite
	res *v1.MsgAddGovernorResponse
	err error
}

func TestMsgAddGovernor(t *testing.T) {
	gocuke.NewRunner(t, &msgAddGovernor{}).
		Path("./features/msg_add_governor.feature").
		Run()
}

func (s *msgAddGovernor) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgAddGovernor) Admin(a string) {
	require.Equal(s.t, s.admin.String(), a)
}

func (s *msgAddGovernor) MsgAddGovernor(a gocuke.DocString) {
	var msg v1.MsgAddGovernor
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.AddGovernor(s.sdkCtx, &msg)
}

func (s *msgAddGovernor) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgAddGovernor) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgAddGovernor) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgAddGovernorResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgAddGovernor) ExpectStateGovernor(a gocuke.DocString) {
	var expected governorv1.Governor
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.k.ss.GovernorTable().Get(s.sdkCtx, expected.Address)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Address, actual.Address)
	require.Equal(s.t, expected.Metadata, actual.Metadata)
}

func (s *msgAddGovernor) ExpectEventAdd(a gocuke.DocString) {
	var expected v1.EventAddGovernor
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
