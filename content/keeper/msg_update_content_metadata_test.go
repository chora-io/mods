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

type msgUpdateContentMetadata struct {
	*baseSuite
	res *v1.MsgUpdateContentMetadataResponse
	err error
}

func TestMsgUpdateContentMetadata(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateContentMetadata{}).
		Path("./msg_update_content_metadata.feature").
		Run()
}

func (s *msgUpdateContentMetadata) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgUpdateContentMetadata) Content(a gocuke.DocString) {
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

func (s *msgUpdateContentMetadata) MsgUpdateContentMetadata(a gocuke.DocString) {
	var msg v1.MsgUpdateContentMetadata
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.UpdateContentMetadata(s.sdkCtx, &msg)
}

func (s *msgUpdateContentMetadata) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateContentMetadata) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgUpdateContentMetadata) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgUpdateContentMetadataResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgUpdateContentMetadata) ExpectStateContent(a gocuke.DocString) {
	var expected contentv1.Content
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.k.ss.ContentTable().Get(s.sdkCtx, expected.Id)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Id, actual.Id)
	require.Equal(s.t, expected.Curator, actual.Curator)
	require.Equal(s.t, expected.Metadata, actual.Metadata)
}

func (s *msgUpdateContentMetadata) ExpectEventUpdateContentMetadata(a gocuke.DocString) {
	var expected v1.EventUpdateContentMetadata
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
