package v1

import (
	"strconv"
	"strings"
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdateContentMetadata struct {
	t   gocuke.TestingT
	msg *MsgUpdateContentMetadata
	err error
}

func TestMsgUpdateContentMetadata(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateContentMetadata{}).
		Path("./msg_update_content_metadata.feature").
		Run()
}

func (s *msgUpdateContentMetadata) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdateContentMetadata) Message(a gocuke.DocString) {
	s.msg = &MsgUpdateContentMetadata{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdateContentMetadata) NewMetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.NewMetadata = strings.Repeat("x", int(length))
}

func (s *msgUpdateContentMetadata) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateContentMetadata) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateContentMetadata) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
