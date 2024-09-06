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

type msgUpdateContentCurator struct {
	*baseSuite
	res *v1.MsgUpdateContentCuratorResponse
	err error
}

func TestMsgUpdateContentCurator(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateContentCurator{}).
		Path("./msg_update_content_curator.feature").
		Run()
}

func (s *msgUpdateContentCurator) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgUpdateContentCurator) Content(a gocuke.DocString) {
	var content contentv1.Content
	err := jsonpb.UnmarshalString(a.Content, &content)
	require.NoError(s.t, err)

	err = s.k.ss.ContentTable().Insert(s.sdkCtx, &contentv1.Content{
		Curator: content.Curator,
		Hash:    content.Hash,
	})
	require.NoError(s.t, err)
}

func (s *msgUpdateContentCurator) MsgUpdateContentCurator(a gocuke.DocString) {
	var msg v1.MsgUpdateContentCurator
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.UpdateContentCurator(s.sdkCtx, &msg)
}

func (s *msgUpdateContentCurator) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateContentCurator) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgUpdateContentCurator) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgUpdateContentCuratorResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgUpdateContentCurator) ExpectStateContent(a gocuke.DocString) {
	var expected contentv1.Content
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.k.ss.ContentTable().Get(s.sdkCtx, expected.Hash)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Curator, actual.Curator)
	require.Equal(s.t, expected.Hash, actual.Hash)
}

func (s *msgUpdateContentCurator) ExpectEventUpdateContentCurator(a gocuke.DocString) {
	var expected v1.EventUpdateContentCurator
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
