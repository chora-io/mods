package v1

import (
	"strconv"
	"strings"
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgRemoveContent struct {
	t   gocuke.TestingT
	msg *MsgRemoveContent
	err error
}

func TestMsgRemoveContent(t *testing.T) {
	gocuke.NewRunner(t, &msgRemoveContent{}).
		Path("./msg_remove_content.feature").
		Run()
}

func (s *msgRemoveContent) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgRemoveContent) Message(a gocuke.DocString) {
	s.msg = &MsgRemoveContent{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgRemoveContent) HashWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.Hash = strings.Repeat("x", int(length))
}

func (s *msgRemoveContent) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgRemoveContent) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgRemoveContent) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
