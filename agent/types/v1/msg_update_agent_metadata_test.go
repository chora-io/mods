package v1

import (
	"strconv"
	"strings"
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdateAgentMetadata struct {
	t   gocuke.TestingT
	msg *MsgUpdateAgentMetadata
	err error
}

func TestMsgUpdateAgentMetadata(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateAgentMetadata{}).
		Path("./msg_update_agent_metadata.feature").
		Run()
}

func (s *msgUpdateAgentMetadata) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdateAgentMetadata) Message(a gocuke.DocString) {
	s.msg = &MsgUpdateAgentMetadata{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdateAgentMetadata) NewMetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.NewMetadata = strings.Repeat("x", int(length))
}

func (s *msgUpdateAgentMetadata) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateAgentMetadata) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateAgentMetadata) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
