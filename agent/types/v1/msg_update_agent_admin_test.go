package v1

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdateAgentAdmin struct {
	t   gocuke.TestingT
	msg *MsgUpdateAgentAdmin
	err error
}

func TestMsgUpdateAgentAdmin(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateAgentAdmin{}).
		Path("./msg_update_agent_admin.feature").
		Run()
}

func (s *msgUpdateAgentAdmin) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdateAgentAdmin) Message(a gocuke.DocString) {
	s.msg = &MsgUpdateAgentAdmin{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdateAgentAdmin) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateAgentAdmin) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateAgentAdmin) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}