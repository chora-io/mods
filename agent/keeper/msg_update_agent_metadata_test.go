package keeper

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	agentv1 "github.com/chora-io/mods/agent/api/v1"
	v1 "github.com/chora-io/mods/agent/types/v1"
	"github.com/chora-io/mods/agent/utils"
)

type msgUpdateAgentMetadata struct {
	*baseSuite
	res *v1.MsgUpdateAgentMetadataResponse
	err error
}

func TestMsgUpdateAgentMetadata(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateAgentMetadata{}).
		Path("./msg_update_agent_metadata.feature").
		Run()
}

func (s *msgUpdateAgentMetadata) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgUpdateAgentMetadata) Agent(a gocuke.DocString) {
	var agent agentv1.Agent
	err := jsonpb.UnmarshalString(a.Content, &agent)
	require.NoError(s.t, err)

	err = s.k.ss.AgentTable().Insert(s.sdkCtx, &agentv1.Agent{
		Address:  agent.Address,
		Admin:    agent.Admin,
		Metadata: agent.Metadata,
	})
	require.NoError(s.t, err)
}

func (s *msgUpdateAgentMetadata) MsgUpdateAgentMetadata(a gocuke.DocString) {
	var msg v1.MsgUpdateAgentMetadata
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.UpdateAgentMetadata(s.sdkCtx, &msg)
}

func (s *msgUpdateAgentMetadata) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateAgentMetadata) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgUpdateAgentMetadata) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgUpdateAgentMetadataResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgUpdateAgentMetadata) ExpectStateAgent(a gocuke.DocString) {
	var expected agentv1.Agent
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.k.ss.AgentTable().Get(s.sdkCtx, expected.Address)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Address, actual.Address)
	require.Equal(s.t, expected.Admin, actual.Admin)
	require.Equal(s.t, expected.Metadata, actual.Metadata)
}

func (s *msgUpdateAgentMetadata) ExpectEventUpdateAgentMetadata(a gocuke.DocString) {
	var expected v1.EventUpdateAgentMetadata
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
