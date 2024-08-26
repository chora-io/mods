package keeper

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	agentv1 "github.com/chora-io/mods/agent/api/v1"
	v1 "github.com/chora-io/mods/agent/types/v1"
)

type queryAgentsByAdmin struct {
	*baseSuite
	res *v1.QueryAgentsByAdminResponse
	err error
}

func TestQueryAgentsByAdmin(t *testing.T) {
	gocuke.NewRunner(t, &queryAgentsByAdmin{}).
		Path("./query_agents_by_admin.feature").
		Run()
}

func (s *queryAgentsByAdmin) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *queryAgentsByAdmin) Agent(a gocuke.DocString) {
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

func (s *queryAgentsByAdmin) QueryAgentsByAdmin(a gocuke.DocString) {
	var req v1.QueryAgentsByAdminRequest
	err := jsonpb.UnmarshalString(a.Content, &req)
	require.NoError(s.t, err)

	s.res, s.err = s.k.AgentsByAdmin(s.sdkCtx, &req)
}

func (s *queryAgentsByAdmin) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *queryAgentsByAdmin) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *queryAgentsByAdmin) ExpectResponse(a gocuke.DocString) {
	var expected v1.QueryAgentsByAdminResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
