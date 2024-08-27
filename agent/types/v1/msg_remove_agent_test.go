package v1

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgRemoveAgent struct {
	t   gocuke.TestingT
	msg *MsgRemoveAgent
	err error
}

func TestMsgRemoveAgent(t *testing.T) {
	gocuke.NewRunner(t, &msgRemoveAgent{}).
		Path("./msg_remove_agent.feature").
		Run()
}

func (s *msgRemoveAgent) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgRemoveAgent) Message(a gocuke.DocString) {
	s.msg = &MsgRemoveAgent{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgRemoveAgent) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgRemoveAgent) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgRemoveAgent) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
