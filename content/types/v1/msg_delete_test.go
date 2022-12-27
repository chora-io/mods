package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgDelete struct {
	t   gocuke.TestingT
	msg *MsgDelete
	err error
}

func TestMsgDelete(t *testing.T) {
	gocuke.NewRunner(t, &msgDelete{}).
		Path("./features/msg_delete.feature").
		Run()
}

func (s *msgDelete) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgDelete) Message(a gocuke.DocString) {
	s.msg = &MsgDelete{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgDelete) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgDelete) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgDelete) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
