package server

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	contentv1 "github.com/choraio/mods/content/api/v1"
	v1 "github.com/choraio/mods/content/types/v1"
	"github.com/choraio/mods/content/utils"
)

type msgCreate struct {
	*baseSuite
	res *v1.MsgCreateResponse
	err error
}

func TestMsgCreate(t *testing.T) {
	gocuke.NewRunner(t, &msgCreate{}).
		Path("./features/msg_create.feature").
		Run()
}

func (s *msgCreate) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgCreate) MsgCreate(a gocuke.DocString) {
	var msg v1.MsgCreate
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.srv.Create(s.sdkCtx, &msg)
}

func (s *msgCreate) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgCreate) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgCreateResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgCreate) ExpectStateContent(a gocuke.DocString) {
	var expected contentv1.Content
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.srv.ss.ContentTable().Get(s.sdkCtx, expected.Id)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Id, actual.Id)
	require.Equal(s.t, expected.Curator, actual.Curator)
	require.Equal(s.t, expected.Metadata, actual.Metadata)
}

func (s *msgCreate) ExpectEventCreate(a gocuke.DocString) {
	var expected v1.EventCreate
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
