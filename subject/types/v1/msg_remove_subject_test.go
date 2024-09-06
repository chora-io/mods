package v1

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgRemoveSubject struct {
	t   gocuke.TestingT
	msg *MsgRemoveSubject
	err error
}

func TestMsgRemoveSubject(t *testing.T) {
	gocuke.NewRunner(t, &msgRemoveSubject{}).
		Path("./msg_remove_subject.feature").
		Run()
}

func (s *msgRemoveSubject) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgRemoveSubject) Message(a gocuke.DocString) {
	s.msg = &MsgRemoveSubject{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgRemoveSubject) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgRemoveSubject) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgRemoveSubject) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
