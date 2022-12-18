package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgDeleteContent struct {
	t   gocuke.TestingT
	msg *MsgDeleteContent
	err error
}

func TestMsgDeleteContent(t *testing.T) {
	gocuke.NewRunner(t, &msgDeleteContent{}).
		Path("./features/msg_delete_content.feature").
		Run()
}

func (s *msgDeleteContent) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgDeleteContent) Message(a gocuke.DocString) {
	s.msg = &MsgDeleteContent{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgDeleteContent) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgDeleteContent) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgDeleteContent) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
