package v1

import (
	"strconv"
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgCreateContent struct {
	t   gocuke.TestingT
	msg *MsgCreateContent
	err error
}

func TestMsgCreateContent(t *testing.T) {
	gocuke.NewRunner(t, &msgCreateContent{}).
		Path("./features/msg_create_content.feature").
		Run()
}

func (s *msgCreateContent) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgCreateContent) TheMessage(a gocuke.DocString) {
	s.msg = &MsgCreateContent{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgCreateContent) HashWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.Hash = strings.Repeat("x", int(length))
}

func (s *msgCreateContent) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgCreateContent) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgCreateContent) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
