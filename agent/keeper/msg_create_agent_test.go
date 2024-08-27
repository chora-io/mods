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

type msgCreateAgent struct {
	*baseSuite
	res *v1.MsgCreateAgentResponse
	err error
}

func TestMsgCreateAgent(t *testing.T) {
	gocuke.NewRunner(t, &msgCreateAgent{}).
		Path("./msg_create_agent.feature").
		Run()
}

func (s *msgCreateAgent) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgCreateAgent) AgentSequence(a gocuke.DocString) {
	var as agentv1.AgentSequence
	err := jsonpb.UnmarshalString(a.Content, &as)
	require.NoError(s.t, err)

	err = s.k.ss.AgentSequenceTable().Save(s.sdkCtx, &agentv1.AgentSequence{
		Sequence: as.Sequence,
	})
	require.NoError(s.t, err)
}

func (s *msgCreateAgent) MsgCreateAgent(a gocuke.DocString) {
	var msg v1.MsgCreateAgent
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.CreateAgent(s.sdkCtx, &msg)
}

func (s *msgCreateAgent) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgCreateAgent) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgCreateAgentResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgCreateAgent) ExpectStateAgent(a gocuke.DocString) {
	var expected agentv1.Agent
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.k.ss.AgentTable().Get(s.sdkCtx, expected.Address)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Address, actual.Address)
	require.Equal(s.t, expected.Admin, actual.Admin)
	require.Equal(s.t, expected.Metadata, actual.Metadata)
}

func (s *msgCreateAgent) ExpectEventCreateAgent(a gocuke.DocString) {
	var expected v1.EventCreateAgent
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
