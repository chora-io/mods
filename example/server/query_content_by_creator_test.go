package server

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	examplev1 "github.com/choraio/mods/example/api/v1"
	v1 "github.com/choraio/mods/example/types/v1"
)

type contentByCreator struct {
	*baseSuite
	res *v1.QueryContentByCreatorResponse
	err error
}

func TestContentByCreator(t *testing.T) {
	gocuke.NewRunner(t, &contentByCreator{}).
		Path("./features/query_content_by_creator.feature").
		Run()
}

func (s *contentByCreator) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *contentByCreator) Content(a gocuke.DocString) {
	var content examplev1.Content
	err := jsonpb.UnmarshalString(a.Content, &content)
	require.NoError(s.t, err)

	id, err := s.srv.ss.ContentTable().InsertReturningID(s.ctx, &examplev1.Content{
		Creator: content.Creator,
		Hash:    content.Hash,
	})
	require.NoError(s.t, err)
	require.Equal(s.t, content.Id, id)
}

func (s *contentByCreator) QueryContentByCreator(a gocuke.DocString) {
	var req v1.QueryContentByCreatorRequest
	err := jsonpb.UnmarshalString(a.Content, &req)
	require.NoError(s.t, err)

	s.res, s.err = s.srv.ContentByCreator(s.ctx, &req)
}

func (s *contentByCreator) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *contentByCreator) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *contentByCreator) ExpectResponse(a gocuke.DocString) {
	var expected v1.QueryContentByCreatorResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
