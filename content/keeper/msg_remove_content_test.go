package keeper

import (
	"strconv"
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	contentv1 "github.com/chora-io/mods/content/api/v1"
	v1 "github.com/chora-io/mods/content/types/v1"
	"github.com/chora-io/mods/content/utils"
)

type msgRemoveContent struct {
	*baseSuite
	res *v1.MsgRemoveContentResponse
	err error
}

func TestMsgRemoveContent(t *testing.T) {
	gocuke.NewRunner(t, &msgRemoveContent{}).
		Path("./msg_remove_content.feature").
		Run()
}

func (s *msgRemoveContent) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgRemoveContent) Content(a gocuke.DocString) {
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

func (s *msgRemoveContent) MsgRemoveContent(a gocuke.DocString) {
	var msg v1.MsgRemoveContent
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.RemoveContent(s.sdkCtx, &msg)
}

func (s *msgRemoveContent) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgRemoveContent) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgRemoveContent) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgRemoveContentResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgRemoveContent) ExpectNoStateContentWithId(a string) {
	id, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	found, err := s.k.ss.ContentTable().Has(s.sdkCtx, id)
	require.NoError(s.t, err)
	require.False(s.t, found)
}

func (s *msgRemoveContent) ExpectEventRemoveContent(a gocuke.DocString) {
	var expected v1.EventRemoveContent
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
