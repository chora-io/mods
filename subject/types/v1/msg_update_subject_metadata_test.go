package v1

import (
	"strconv"
	"strings"
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdateSubjectMetadata struct {
	t   gocuke.TestingT
	msg *MsgUpdateSubjectMetadata
	err error
}

func TestMsgUpdateSubjectMetadata(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateSubjectMetadata{}).
		Path("./msg_update_subject_metadata.feature").
		Run()
}

func (s *msgUpdateSubjectMetadata) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdateSubjectMetadata) Message(a gocuke.DocString) {
	s.msg = &MsgUpdateSubjectMetadata{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdateSubjectMetadata) NewMetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.NewMetadata = strings.Repeat("x", int(length))
}

func (s *msgUpdateSubjectMetadata) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateSubjectMetadata) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateSubjectMetadata) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
