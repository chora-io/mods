package server

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	contentv1 "github.com/choraio/mods/content/api/v1"
	v1 "github.com/choraio/mods/content/types/v1"
)

type queryContent struct {
	*baseSuite
	res *v1.QueryContentResponse
	err error
}

func TestQueryContent(t *testing.T) {
	gocuke.NewRunner(t, &queryContent{}).
		Path("./features/query_content.feature").
		Run()
}

func (s *queryContent) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *queryContent) Content(a gocuke.DocString) {
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

func (s *queryContent) QueryContent(a gocuke.DocString) {
	var req v1.QueryContentRequest
	err := jsonpb.UnmarshalString(a.Content, &req)
	require.NoError(s.t, err)

	s.res, s.err = s.srv.Content(s.ctx, &req)
}

func (s *queryContent) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *queryContent) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *queryContent) ExpectResponse(a gocuke.DocString) {
	var expected v1.QueryContentResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
