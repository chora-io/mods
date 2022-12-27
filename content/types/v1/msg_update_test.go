package v1

import (
	"strconv"
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdate struct {
	t   gocuke.TestingT
	msg *MsgUpdate
	err error
}

func TestMsgUpdate(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdate{}).
		Path("./features/msg_update.feature").
		Run()
}

func (s *msgUpdate) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdate) Message(a gocuke.DocString) {
	s.msg = &MsgUpdate{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdate) NewHashWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.NewHash = strings.Repeat("x", int(length))
}

func (s *msgUpdate) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdate) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdate) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
