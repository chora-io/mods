package server

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/types/v2/testutil"

	contentv1 "github.com/choraio/mods/geonode/api/v1"
	v1 "github.com/choraio/mods/geonode/types/v1"
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

func (s *msgUpdateMetadata) Node(a gocuke.DocString) {
	var node contentv1.Node
	err := jsonpb.UnmarshalString(a.Content, &node)
	require.NoError(s.t, err)

	id, err := s.srv.ss.NodeTable().InsertReturningID(s.ctx, &contentv1.Node{
		Curator:  node.Curator,
		Metadata: node.Metadata,
	})
	require.NoError(s.t, err)
	require.Equal(s.t, node.Id, id)
}

func (s *msgUpdateMetadata) MsgUpdateMetadata(a gocuke.DocString) {
	var msg v1.MsgUpdateMetadata
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.srv.UpdateMetadata(s.ctx, &msg)
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

func (s *msgUpdateMetadata) ExpectStateNode(a gocuke.DocString) {
	var expected contentv1.Node
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.srv.ss.NodeTable().Get(s.ctx, expected.Id)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Id, actual.Id)
	require.Equal(s.t, expected.Curator, actual.Curator)
	require.Equal(s.t, expected.Metadata, actual.Metadata)
}

func (s *msgUpdateMetadata) ExpectEventUpdateMetadata(a gocuke.DocString) {
	var expected v1.EventUpdateMetadata
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := testutil.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
