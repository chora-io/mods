package server

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	contentv1 "github.com/choraio/mods/content/api/v1"
	v1 "github.com/choraio/mods/content/types/v1"
)

type queryContentByCurator struct {
	*baseSuite
	res *v1.QueryContentByCuratorResponse
	err error
}

func TestQueryContentByCurator(t *testing.T) {
	gocuke.NewRunner(t, &queryContentByCurator{}).
		Path("./features/query_content_by_curator.feature").
		Run()
}

func (s *queryContentByCurator) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *queryContentByCurator) Content(a gocuke.DocString) {
	var content contentv1.Content
	err := jsonpb.UnmarshalString(a.Content, &content)
	require.NoError(s.t, err)

	id, err := s.srv.ss.ContentTable().InsertReturningID(s.ctx, &contentv1.Content{
		Curator:  content.Curator,
		Metadata: content.Metadata,
	})
	require.NoError(s.t, err)
	require.Equal(s.t, content.Id, id)
}

func (s *queryContentByCurator) QueryContentByCurator(a gocuke.DocString) {
	var req v1.QueryContentByCuratorRequest
	err := jsonpb.UnmarshalString(a.Content, &req)
	require.NoError(s.t, err)

	s.res, s.err = s.srv.ContentByCurator(s.ctx, &req)
}

func (s *queryContentByCurator) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *queryContentByCurator) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *queryContentByCurator) ExpectResponse(a gocuke.DocString) {
	var expected v1.QueryContentByCuratorResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
