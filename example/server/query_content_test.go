package server

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	examplev1 "github.com/choraio/mods/example/api/v1"
	v1 "github.com/choraio/mods/example/types/v1"
)

type content struct {
	*baseSuite
	res *v1.QueryContentResponse
	err error
}

func TestContent(t *testing.T) {
	gocuke.NewRunner(t, &content{}).
		Path("./features/query_content.feature").
		Run()
}

func (s *content) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *content) Content(a gocuke.DocString) {
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

func (s *content) QueryContent(a gocuke.DocString) {
	var req v1.QueryContentRequest
	err := jsonpb.UnmarshalString(a.Content, &req)
	require.NoError(s.t, err)

	s.res, s.err = s.srv.Content(s.ctx, &req)
}

func (s *content) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *content) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *content) ExpectResponse(a gocuke.DocString) {
	var expected v1.QueryContentResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
