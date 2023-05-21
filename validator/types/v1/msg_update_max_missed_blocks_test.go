package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdateMaxMissedBlocks struct {
	t   gocuke.TestingT
	msg *MsgUpdateMaxMissedBlocks
	err error
}

func TestMsgUpdateMaxMissedBlocks(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateMaxMissedBlocks{}).
		Path("./features/msg_update_max_missed_blocks.feature").
		Run()
}

func (s *msgUpdateMaxMissedBlocks) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdateMaxMissedBlocks) Message(a gocuke.DocString) {
	s.msg = &MsgUpdateMaxMissedBlocks{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdateMaxMissedBlocks) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateMaxMissedBlocks) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateMaxMissedBlocks) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
