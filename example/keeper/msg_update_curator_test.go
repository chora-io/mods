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

func (s *msgUpdateCurator) Content(a gocuke.DocString) {
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

func (s *msgUpdateCurator) MsgUpdateCurator(a gocuke.DocString) {
	var msg v1.MsgUpdateCurator
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.UpdateCurator(s.sdkCtx, &msg)
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

func (s *msgUpdateCurator) ExpectStateContent(a gocuke.DocString) {
	var expected examplev1.Content
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.k.ss.ContentTable().Get(s.sdkCtx, expected.Id)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Id, actual.Id)
	require.Equal(s.t, expected.Curator, actual.Curator)
	require.Equal(s.t, expected.Metadata, actual.Metadata)
}

func (s *msgUpdateCurator) ExpectEventUpdateCurator(a gocuke.DocString) {
	var expected v1.EventUpdateCurator
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
