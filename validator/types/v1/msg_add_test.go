package v1

import (
	"strconv"
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgAdd struct {
	t   gocuke.TestingT
	msg *MsgAdd
	err error
}

func TestMsgAdd(t *testing.T) {
	gocuke.NewRunner(t, &msgAdd{}).
		Path("./features/msg_add.feature").
		Run()
}

func (s *msgAdd) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgAdd) Message(a gocuke.DocString) {
	s.msg = &MsgAdd{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgAdd) MetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.Metadata = strings.Repeat("x", int(length))
}

func (s *msgAdd) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgAdd) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgAdd) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
