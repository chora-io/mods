package keeper

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	agentv1 "github.com/chora-io/mods/agent/api/v1"
	v1 "github.com/chora-io/mods/agent/types/v1"
)

type queryAgent struct {
	*baseSuite
	res *v1.QueryAgentResponse
	err error
}

func TestQueryAgent(t *testing.T) {
	gocuke.NewRunner(t, &queryAgent{}).
		Path("./query_agent.feature").
		Run()
}

func (s *queryAgent) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *queryAgent) Agent(a gocuke.DocString) {
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

func (s *queryAgent) QueryAgent(a gocuke.DocString) {
	var req v1.QueryAgentRequest
	err := jsonpb.UnmarshalString(a.Content, &req)
	require.NoError(s.t, err)

	s.res, s.err = s.k.Agent(s.sdkCtx, &req)
}

func (s *queryAgent) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *queryAgent) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *queryAgent) ExpectResponse(a gocuke.DocString) {
	var expected v1.QueryAgentResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
