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

type msgUpdateGovernorMetadata struct {
	*baseSuite
	res *v1.MsgUpdateGovernorMetadataResponse
	err error
}

func TestMsgUpdateGovernorMetadata(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateGovernorMetadata{}).
		Path("./msg_update_governor_metadata.feature").
		Run()
}

func (s *msgUpdateGovernorMetadata) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgUpdateGovernorMetadata) Governor(a gocuke.DocString) {
	var governor governorv1.Governor
	err := jsonpb.UnmarshalString(a.Content, &governor)
	require.NoError(s.t, err)

	err = s.k.ss.GovernorTable().Insert(s.sdkCtx, &governorv1.Governor{
		Address: governor.Address,
	})
	require.NoError(s.t, err)
}

func (s *msgUpdateGovernorMetadata) MsgUpdateGovernorMetadata(a gocuke.DocString) {
	var msg v1.MsgUpdateGovernorMetadata
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.UpdateGovernorMetadata(s.sdkCtx, &msg)
}

func (s *msgUpdateGovernorMetadata) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateGovernorMetadata) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgUpdateGovernorMetadata) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgUpdateGovernorMetadataResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgUpdateGovernorMetadata) ExpectStateGovernor(a gocuke.DocString) {
	var expected governorv1.Governor
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.k.ss.GovernorTable().Get(s.sdkCtx, expected.Address)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Address, actual.Address)
	require.Equal(s.t, expected.Metadata, actual.Metadata)
}

func (s *msgUpdateGovernorMetadata) ExpectEventUpdateGovernor(a gocuke.DocString) {
	var expected v1.EventUpdateGovernorMetadata
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
