package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdatePolicy struct {
	t   gocuke.TestingT
	msg *MsgUpdatePolicy
	err error
}

func TestMsgUpdatePolicy(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdatePolicy{}).
		Path("./features/msg_update_max_missed_blocks.feature").
		Run()
}

func (s *msgUpdatePolicy) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdatePolicy) Message(a gocuke.DocString) {
	s.msg = &MsgUpdatePolicy{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdatePolicy) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdatePolicy) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdatePolicy) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
