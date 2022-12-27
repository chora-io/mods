package v1

import (
	"strconv"
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgCreate struct {
	t   gocuke.TestingT
	msg *MsgCreate
	err error
}

func TestMsgCreate(t *testing.T) {
	gocuke.NewRunner(t, &msgCreate{}).
		Path("./features/msg_create.feature").
		Run()
}

func (s *msgCreate) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgCreate) Message(a gocuke.DocString) {
	s.msg = &MsgCreate{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgCreate) HashWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.Hash = strings.Repeat("x", int(length))
}

func (s *msgCreate) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgCreate) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgCreate) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
