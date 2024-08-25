package keeper

import (
	"strconv"
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	examplev1 "github.com/chora-io/mods/example/api/v1"
	v1 "github.com/chora-io/mods/example/types/v1"
	"github.com/chora-io/mods/example/utils"
)

type msgDelete struct {
	*baseSuite
	res *v1.MsgDeleteResponse
	err error
}

func TestMsgDelete(t *testing.T) {
	gocuke.NewRunner(t, &msgDelete{}).
		Path("./features/msg_delete.feature").
		Run()
}

func (s *msgDelete) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgDelete) Content(a gocuke.DocString) {
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

func (s *msgDelete) MsgDelete(a gocuke.DocString) {
	var msg v1.MsgDelete
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.Delete(s.sdkCtx, &msg)
}

func (s *msgDelete) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgDelete) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgDelete) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgDeleteResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgDelete) ExpectNoStateContentWithId(a string) {
	id, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	found, err := s.k.ss.ContentTable().Has(s.sdkCtx, id)
	require.NoError(s.t, err)
	require.False(s.t, found)
}

func (s *msgDelete) ExpectEventDelete(a gocuke.DocString) {
	var expected v1.EventDelete
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
