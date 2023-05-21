package server

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	validatorv1 "github.com/choraio/mods/validator/api/v1"
	v1 "github.com/choraio/mods/validator/types/v1"
	"github.com/choraio/mods/validator/utils"
)

type msgUpdateMaxMissedBlocks struct {
	*baseSuite
	res *v1.MsgUpdateMaxMissedBlocksResponse
	err error
}

func TestMsgUpdateMaxMissedBlocks(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateMaxMissedBlocks{}).
		Path("./features/msg_update_max_missed_blocks.feature").
		Run()
}

func (s *msgUpdateMaxMissedBlocks) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgUpdateMaxMissedBlocks) Authority(a string) {
	require.Equal(s.t, s.authority.String(), a)
}

func (s *msgUpdateMaxMissedBlocks) MaxMissedBlocks(a gocuke.DocString) {
	var maxMissedBlocks validatorv1.MaxMissedBlocks
	err := jsonpb.UnmarshalString(a.Content, &maxMissedBlocks)
	require.NoError(s.t, err)

	err = s.srv.ss.MaxMissedBlocksTable().Save(s.ctx, &maxMissedBlocks)
	require.NoError(s.t, err)
}

func (s *msgUpdateMaxMissedBlocks) MsgUpdateMaxMissedBlocks(a gocuke.DocString) {
	var msg v1.MsgUpdateMaxMissedBlocks
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.srv.UpdateMaxMissedBlocks(s.ctx, &msg)
}

func (s *msgUpdateMaxMissedBlocks) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateMaxMissedBlocks) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgUpdateMaxMissedBlocks) ExpectMaxMissedBlocks(a gocuke.DocString) {
	var expected validatorv1.MaxMissedBlocks
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	maxMissedBlocks, err := s.srv.ss.MaxMissedBlocksTable().Get(s.ctx)
	require.NoError(s.t, err)
	require.Equal(s.t, expected.MaxMissedBlocks, maxMissedBlocks.MaxMissedBlocks)
}

func (s *msgUpdateMaxMissedBlocks) ExpectEventUpdateMaxMissedBlocks(a gocuke.DocString) {
	var expected v1.EventUpdateMaxMissedBlocks
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}

func (s *msgUpdateMaxMissedBlocks) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgUpdateMaxMissedBlocksResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
