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

type msgRemoveAgent struct {
	*baseSuite
	res *v1.MsgRemoveAgentResponse
	err error
}

func TestMsgRemoveAgent(t *testing.T) {
	gocuke.NewRunner(t, &msgRemoveAgent{}).
		Path("./msg_remove_agent.feature").
		Run()
}

func (s *msgRemoveAgent) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgRemoveAgent) Agent(a gocuke.DocString) {
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

func (s *msgRemoveAgent) MsgRemoveAgent(a gocuke.DocString) {
	var msg v1.MsgRemoveAgent
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.RemoveAgent(s.sdkCtx, &msg)
}

func (s *msgRemoveAgent) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgRemoveAgent) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgRemoveAgent) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgRemoveAgentResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgRemoveAgent) ExpectNoStateAgent(a gocuke.DocString) {
	var expected agentv1.Agent
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.k.ss.AgentTable().Get(s.sdkCtx, expected.Address)
	require.Nil(s.t, actual)
	// require.EqualError(s.t, s.err, "not found")
}

func (s *msgRemoveAgent) ExpectEventRemoveAgent(a gocuke.DocString) {
	var expected v1.EventRemoveAgent
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
