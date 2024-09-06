package keeper

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	contentv1 "github.com/chora-io/mods/content/api/v1"
	v1 "github.com/chora-io/mods/content/types/v1"
	"github.com/chora-io/mods/content/utils"
)

type msgCreateContent struct {
	*baseSuite
	res *v1.MsgCreateContentResponse
	err error
}

func TestMsgCreateContent(t *testing.T) {
	gocuke.NewRunner(t, &msgCreateContent{}).
		Path("./msg_create_content.feature").
		Run()
}

func (s *msgCreateContent) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgCreateContent) MsgCreateContent(a gocuke.DocString) {
	var msg v1.MsgCreateContent
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.CreateContent(s.sdkCtx, &msg)
}

func (s *msgCreateContent) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgCreateContent) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgCreateContentResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgCreateContent) ExpectStateContent(a gocuke.DocString) {
	var expected contentv1.Content
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.k.ss.ContentTable().Get(s.sdkCtx, expected.Hash)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Curator, actual.Curator)
	require.Equal(s.t, expected.Hash, actual.Hash)
}

func (s *msgCreateContent) ExpectEventCreateContent(a gocuke.DocString) {
	var expected v1.EventCreateContent
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
