package keeper

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	contentv1 "github.com/chora-io/mods/content/api/v1"
	v1 "github.com/chora-io/mods/content/types/v1"
)

type queryContents struct {
	*baseSuite
	res *v1.QueryContentsResponse
	err error
}

func TestQueryContents(t *testing.T) {
	gocuke.NewRunner(t, &queryContents{}).
		Path("./query_contents.feature").
		Run()
}

func (s *queryContents) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *queryContents) Content(a gocuke.DocString) {
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

func (s *queryContents) QueryContents(a gocuke.DocString) {
	var req v1.QueryContentsRequest
	err := jsonpb.UnmarshalString(a.Content, &req)
	require.NoError(s.t, err)

	s.res, s.err = s.k.Contents(s.sdkCtx, &req)
}

func (s *queryContents) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *queryContents) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *queryContents) ExpectResponse(a gocuke.DocString) {
	var expected v1.QueryContentsResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
