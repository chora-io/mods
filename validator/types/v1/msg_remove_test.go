package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgRemove struct {
	t   gocuke.TestingT
	msg *MsgRemove
	err error
}

func TestMsgRemove(t *testing.T) {
	gocuke.NewRunner(t, &msgRemove{}).
		Path("./features/msg_remove.feature").
		Run()
}

func (s *msgRemove) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgRemove) Message(a gocuke.DocString) {
	s.msg = &MsgRemove{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgRemove) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgRemove) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgRemove) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
