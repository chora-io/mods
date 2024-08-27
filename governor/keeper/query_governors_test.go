package keeper

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	governorv1 "github.com/chora-io/mods/governor/api/v1"
	v1 "github.com/chora-io/mods/governor/types/v1"
)

type queryGovernors struct {
	*baseSuite
	res *v1.QueryGovernorsResponse
	err error
}

func TestQueryGovernors(t *testing.T) {
	gocuke.NewRunner(t, &queryGovernors{}).
		Path("./query_governors.feature").
		Run()
}

func (s *queryGovernors) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *queryGovernors) Governor(a gocuke.DocString) {
	var governor governorv1.Governor
	err := jsonpb.UnmarshalString(a.Content, &governor)
	require.NoError(s.t, err)

	err = s.k.ss.GovernorTable().Insert(s.sdkCtx, &governorv1.Governor{
		Address:  governor.Address,
		Metadata: governor.Metadata,
	})
	require.NoError(s.t, err)
}

func (s *queryGovernors) QueryGovernors(a gocuke.DocString) {
	var req v1.QueryGovernorsRequest
	err := jsonpb.UnmarshalString(a.Content, &req)
	require.NoError(s.t, err)

	s.res, s.err = s.k.Governors(s.sdkCtx, &req)
}

func (s *queryGovernors) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *queryGovernors) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *queryGovernors) ExpectResponse(a gocuke.DocString) {
	var expected v1.QueryGovernorsResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
