package server

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/types/v2/testutil"

	examplev1 "github.com/choraio/mods/example/api/v1"
	v1 "github.com/choraio/mods/example/types/v1"
)

type updateContent struct {
	*baseSuite
	res *v1.MsgUpdateContentResponse
	err error
}

func TestUpdateContent(t *testing.T) {
	gocuke.NewRunner(t, &updateContent{}).
		Path("./features/msg_update_content.feature").
		Run()
}

func (s *updateContent) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *updateContent) Content(a gocuke.DocString) {
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

func (s *updateContent) MsgUpdateContent(a gocuke.DocString) {
	var msg v1.MsgUpdateContent
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.srv.UpdateContent(s.ctx, &msg)
}

func (s *updateContent) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *updateContent) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *updateContent) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgUpdateContentResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *updateContent) ExpectStateContent(a gocuke.DocString) {
	var expected examplev1.Content
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.srv.ss.ContentTable().Get(s.ctx, expected.Id)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Id, actual.Id)
	require.Equal(s.t, expected.Creator, actual.Creator)
	require.Equal(s.t, expected.Hash, actual.Hash)
}

func (s *updateContent) ExpectEventUpdateContent(a gocuke.DocString) {
	var expected v1.EventUpdateContent
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := testutil.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
