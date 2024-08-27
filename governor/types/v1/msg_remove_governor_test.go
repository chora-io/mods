package v1

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgRemoveGovernor struct {
	t   gocuke.TestingT
	msg *MsgRemoveGovernor
	err error
}

func TestMsgRemoveGovernor(t *testing.T) {
	gocuke.NewRunner(t, &msgRemoveGovernor{}).
		Path("./msg_remove_governor.feature").
		Run()
}

func (s *msgRemoveGovernor) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgRemoveGovernor) Message(a gocuke.DocString) {
	s.msg = &MsgRemoveGovernor{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgRemoveGovernor) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgRemoveGovernor) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgRemoveGovernor) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
