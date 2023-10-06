package keeper

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	contentv1 "github.com/choraio/mods/content/api/v1"
	v1 "github.com/choraio/mods/content/types/v1"
)

type queryContentsByCurator struct {
	*baseSuite
	res *v1.QueryContentsByCuratorResponse
	err error
}

func TestQueryContentsByCurator(t *testing.T) {
	gocuke.NewRunner(t, &queryContentsByCurator{}).
		Path("./features/query_contents_by_curator.feature").
		Run()
}

func (s *queryContentsByCurator) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *queryContentsByCurator) Content(a gocuke.DocString) {
	var content contentv1.Content
	err := jsonpb.UnmarshalString(a.Content, &content)
	require.NoError(s.t, err)

	id, err := s.k.ss.ContentTable().InsertReturningId(s.sdkCtx, &contentv1.Content{
		Curator:  content.Curator,
		Metadata: content.Metadata,
	})
	require.NoError(s.t, err)
	require.Equal(s.t, content.Id, id)
}

func (s *queryContentsByCurator) QueryContentsByCurator(a gocuke.DocString) {
	var req v1.QueryContentsByCuratorRequest
	err := jsonpb.UnmarshalString(a.Content, &req)
	require.NoError(s.t, err)

	s.res, s.err = s.k.ContentsByCurator(s.sdkCtx, &req)
}

func (s *queryContentsByCurator) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *queryContentsByCurator) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *queryContentsByCurator) ExpectResponse(a gocuke.DocString) {
	var expected v1.QueryContentsByCuratorResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
