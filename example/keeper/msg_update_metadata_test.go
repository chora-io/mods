package keeper

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	examplev1 "github.com/chora-io/mods/example/api/v1"
	v1 "github.com/chora-io/mods/example/types/v1"
	"github.com/chora-io/mods/example/utils"
)

type msgUpdateMetadata struct {
	*baseSuite
	res *v1.MsgUpdateMetadataResponse
	err error
}

func TestMsgUpdateMetadata(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateMetadata{}).
		Path("./features/msg_update_metadata.feature").
		Run()
}

func (s *msgUpdateMetadata) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgUpdateMetadata) Content(a gocuke.DocString) {
	var example examplev1.Content
	err := jsonpb.UnmarshalString(a.Content, &example)
	require.NoError(s.t, err)

	id, err := s.k.ss.ContentTable().InsertReturningId(s.sdkCtx, &examplev1.Content{
		Curator:  example.Curator,
		Metadata: example.Metadata,
	})
	require.NoError(s.t, err)
	require.Equal(s.t, example.Id, id)
}

func (s *msgUpdateMetadata) MsgUpdateMetadata(a gocuke.DocString) {
	var msg v1.MsgUpdateMetadata
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.UpdateMetadata(s.sdkCtx, &msg)
}

func (s *msgUpdateMetadata) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateMetadata) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgUpdateMetadata) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgUpdateMetadataResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgUpdateMetadata) ExpectStateContent(a gocuke.DocString) {
	var expected examplev1.Content
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.k.ss.ContentTable().Get(s.sdkCtx, expected.Id)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Id, actual.Id)
	require.Equal(s.t, expected.Curator, actual.Curator)
	require.Equal(s.t, expected.Metadata, actual.Metadata)
}

func (s *msgUpdateMetadata) ExpectEventUpdateMetadata(a gocuke.DocString) {
	var expected v1.EventUpdateMetadata
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
