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

type msgUpdateCurator struct {
	*baseSuite
	res *v1.MsgUpdateCuratorResponse
	err error
}

func TestMsgUpdateCurator(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateCurator{}).
		Path("./features/msg_update_curator.feature").
		Run()
}

func (s *msgUpdateCurator) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgUpdateCurator) Node(a gocuke.DocString) {
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

func (s *msgUpdateCurator) MsgUpdateCurator(a gocuke.DocString) {
	var msg v1.MsgUpdateCurator
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.srv.UpdateCurator(s.ctx, &msg)
}

func (s *msgUpdateCurator) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateCurator) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgUpdateCurator) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgUpdateCuratorResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgUpdateCurator) ExpectStateNode(a gocuke.DocString) {
	var expected contentv1.Node
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.srv.ss.NodeTable().Get(s.ctx, expected.Id)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Id, actual.Id)
	require.Equal(s.t, expected.Curator, actual.Curator)
	require.Equal(s.t, expected.Metadata, actual.Metadata)
}

func (s *msgUpdateCurator) ExpectEventUpdateCurator(a gocuke.DocString) {
	var expected v1.EventUpdateCurator
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := testutil.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
