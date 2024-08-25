package keeper

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	governorv1 "github.com/chora-io/mods/governor/api/v1"
	v1 "github.com/chora-io/mods/governor/types/v1"
)

type queryGovernor struct {
	*baseSuite
	res *v1.QueryGovernorResponse
	err error
}

func TestQueryGovernor(t *testing.T) {
	gocuke.NewRunner(t, &queryGovernor{}).
		Path("./features/query_governor.feature").
		Run()
}

func (s *queryGovernor) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *queryGovernor) Governor(a gocuke.DocString) {
	var governor governorv1.Governor
	err := jsonpb.UnmarshalString(a.Content, &governor)
	require.NoError(s.t, err)

	err = s.k.ss.GovernorTable().Insert(s.sdkCtx, &governorv1.Governor{
		Address:  governor.Address,
		Metadata: governor.Metadata,
	})
	require.NoError(s.t, err)
}

func (s *queryGovernor) QueryGovernor(a gocuke.DocString) {
	var req v1.QueryGovernorRequest
	err := jsonpb.UnmarshalString(a.Content, &req)
	require.NoError(s.t, err)

	s.res, s.err = s.k.Governor(s.sdkCtx, &req)
}

func (s *queryGovernor) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *queryGovernor) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *queryGovernor) ExpectResponse(a gocuke.DocString) {
	var expected v1.QueryGovernorResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
