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

type msgCreateGovernor struct {
	*baseSuite
	res *v1.MsgCreateGovernorResponse
	err error
}

func TestMsgCreateGovernor(t *testing.T) {
	gocuke.NewRunner(t, &msgCreateGovernor{}).
		Path("./msg_create_governor.feature").
		Run()
}

func (s *msgCreateGovernor) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgCreateGovernor) Admin(a string) {
	require.Equal(s.t, s.admin.String(), a)
}

func (s *msgCreateGovernor) MsgCreateGovernor(a gocuke.DocString) {
	var msg v1.MsgCreateGovernor
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.CreateGovernor(s.sdkCtx, &msg)
}

func (s *msgCreateGovernor) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgCreateGovernor) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgCreateGovernor) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgCreateGovernorResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgCreateGovernor) ExpectStateGovernor(a gocuke.DocString) {
	var expected governorv1.Governor
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.k.ss.GovernorTable().Get(s.sdkCtx, expected.Address)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Address, actual.Address)
	require.Equal(s.t, expected.Metadata, actual.Metadata)
}

func (s *msgCreateGovernor) ExpectEventCreateGovernor(a gocuke.DocString) {
	var expected v1.EventCreateGovernor
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
