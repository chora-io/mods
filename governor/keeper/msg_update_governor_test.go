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

type msgUpdateGovernor struct {
	*baseSuite
	res *v1.MsgUpdateGovernorResponse
	err error
}

func TestMsgUpdateGovernor(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateGovernor{}).
		Path("./features/msg_update_governor.feature").
		Run()
}

func (s *msgUpdateGovernor) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgUpdateGovernor) Governor(a gocuke.DocString) {
	var governor governorv1.Governor
	err := jsonpb.UnmarshalString(a.Content, &governor)
	require.NoError(s.t, err)

	err = s.k.ss.GovernorTable().Insert(s.sdkCtx, &governorv1.Governor{
		Address: governor.Address,
	})
	require.NoError(s.t, err)
}

func (s *msgUpdateGovernor) MsgUpdateGovernor(a gocuke.DocString) {
	var msg v1.MsgUpdateGovernor
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.UpdateGovernor(s.sdkCtx, &msg)
}

func (s *msgUpdateGovernor) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateGovernor) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgUpdateGovernor) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgUpdateGovernorResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgUpdateGovernor) ExpectStateGovernor(a gocuke.DocString) {
	var expected governorv1.Governor
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.k.ss.GovernorTable().Get(s.sdkCtx, expected.Address)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Address, actual.Address)
	require.Equal(s.t, expected.Metadata, actual.Metadata)
}

func (s *msgUpdateGovernor) ExpectEventUpdateGovernor(a gocuke.DocString) {
	var expected v1.EventUpdateGovernor
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
