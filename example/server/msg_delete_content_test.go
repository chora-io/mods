package server

import (
	"strconv"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/types/v2/testutil"

	examplev1 "github.com/choraio/mods/example/api/v1"
	v1 "github.com/choraio/mods/example/types/v1"
)

type deleteContent struct {
	*baseSuite
	res *v1.MsgDeleteContentResponse
	err error
}

func TestDeleteContent(t *testing.T) {
	gocuke.NewRunner(t, &deleteContent{}).
		Path("./features/msg_delete_content.feature").
		Run()
}

func (s *deleteContent) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *deleteContent) Content(a gocuke.DocString) {
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

func (s *deleteContent) MsgDeleteContent(a gocuke.DocString) {
	var msg v1.MsgDeleteContent
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.srv.DeleteContent(s.ctx, &msg)
}

func (s *deleteContent) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *deleteContent) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *deleteContent) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgDeleteContentResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *deleteContent) ExpectNoStateContentWithId(a string) {
	id, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	found, err := s.srv.ss.ContentTable().Has(s.ctx, id)
	require.NoError(s.t, err)
	require.False(s.t, found)
}

func (s *deleteContent) ExpectEventDeleteContent(a gocuke.DocString) {
	var expected v1.EventDeleteContent
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := testutil.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
