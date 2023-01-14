package v1

import (
	"strconv"
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgIssue struct {
	t   gocuke.TestingT
	msg *MsgIssue
	err error
}

func TestMsgIssue(t *testing.T) {
	gocuke.NewRunner(t, &msgIssue{}).
		Path("./features/msg_issue.feature").
		Run()
}

func (s *msgIssue) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgIssue) Message(a gocuke.DocString) {
	s.msg = &MsgIssue{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgIssue) MetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.Metadata = strings.Repeat("x", int(length))
}

func (s *msgIssue) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgIssue) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgIssue) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
