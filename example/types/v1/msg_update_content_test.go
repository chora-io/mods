package v1

import (
	"strconv"
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdateContent struct {
	t   gocuke.TestingT
	msg *MsgUpdateContent
	err error
}

func TestMsgUpdateContent(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateContent{}).
		Path("./features/msg_update_content.feature").
		Run()
}

func (s *msgUpdateContent) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdateContent) Message(a gocuke.DocString) {
	s.msg = &MsgUpdateContent{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdateContent) NewHashWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.NewHash = strings.Repeat("x", int(length))
}

func (s *msgUpdateContent) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateContent) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateContent) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
