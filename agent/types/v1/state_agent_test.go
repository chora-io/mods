package v1

import (
	"strconv"
	"strings"
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type agent struct {
	t     gocuke.TestingT
	agent *Agent
	err   error
}

func TestAgent(t *testing.T) {
	gocuke.NewRunner(t, &agent{}).
		Path("./state_agent.feature").
		Run()
}

func (s *agent) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *agent) Agent(a gocuke.DocString) {
	s.agent = &Agent{}
	err := jsonpb.UnmarshalString(a.Content, s.agent)
	require.NoError(s.t, err)
}

func (s *agent) MetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.agent.Metadata = strings.Repeat("x", int(length))
}

func (s *agent) ValidateAgent() {
	s.err = s.agent.Validate()
}

func (s *agent) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *agent) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
