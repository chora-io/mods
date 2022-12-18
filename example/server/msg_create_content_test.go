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

type createContent struct {
	*baseSuite
	res *v1.MsgCreateContentResponse
	err error
}

func TestCreateContent(t *testing.T) {
	gocuke.NewRunner(t, &createContent{}).
		Path("./features/msg_create_content.feature").
		Run()
}

func (s *createContent) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *createContent) MsgCreateContent(a gocuke.DocString) {
	var msg v1.MsgCreateContent
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.srv.CreateContent(s.ctx, &msg)
}

func (s *createContent) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *createContent) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgCreateContentResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *createContent) ExpectStateContent(a gocuke.DocString) {
	var expected examplev1.Content
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.srv.ss.ContentTable().Get(s.ctx, expected.Id)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Id, actual.Id)
	require.Equal(s.t, expected.Creator, actual.Creator)
	require.Equal(s.t, expected.Hash, actual.Hash)
}

func (s *createContent) ExpectEventCreateContent(a gocuke.DocString) {
	var expected v1.EventCreateContent
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := testutil.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
