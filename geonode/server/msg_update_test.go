package server

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/types/v2/testutil"

	geonodev1 "github.com/choraio/mods/geonode/api/v1"
	v1 "github.com/choraio/mods/geonode/types/v1"
)

type msgUpdate struct {
	*baseSuite
	res *v1.MsgUpdateResponse
	err error
}

func TestMsgUpdate(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdate{}).
		Path("./features/msg_update.feature").
		Run()
}

func (s *msgUpdate) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgUpdate) Node(a gocuke.DocString) {
	var node geonodev1.Node
	err := jsonpb.UnmarshalString(a.Content, &node)
	require.NoError(s.t, err)

	id, err := s.srv.ss.NodeTable().InsertReturningID(s.ctx, &geonodev1.Node{
		Curator:  node.Curator,
		Metadata: node.Metadata,
	})
	require.NoError(s.t, err)
	require.Equal(s.t, node.Id, id)
}

func (s *msgUpdate) MsgUpdate(a gocuke.DocString) {
	var msg v1.MsgUpdate
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.srv.Update(s.ctx, &msg)
}

func (s *msgUpdate) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdate) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgUpdate) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgUpdateResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgUpdate) ExpectStateNode(a gocuke.DocString) {
	var expected geonodev1.Node
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.srv.ss.NodeTable().Get(s.ctx, expected.Id)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Id, actual.Id)
	require.Equal(s.t, expected.Curator, actual.Curator)
	require.Equal(s.t, expected.Metadata, actual.Metadata)
}

func (s *msgUpdate) ExpectEventUpdate(a gocuke.DocString) {
	var expected v1.EventUpdate
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := testutil.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
