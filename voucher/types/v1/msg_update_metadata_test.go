package v1

import (
	"strconv"
	"strings"
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdateMetadata struct {
	t   gocuke.TestingT
	msg *MsgUpdateMetadata
	err error
}

func TestMsgUpdateMetadata(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateMetadata{}).
		Path("./features/msg_update_metadata.feature").
		Run()
}

func (s *msgUpdateMetadata) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdateMetadata) Message(a gocuke.DocString) {
	s.msg = &MsgUpdateMetadata{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdateMetadata) NewMetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.NewMetadata = strings.Repeat("x", int(length))
}

func (s *msgUpdateMetadata) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateMetadata) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateMetadata) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
