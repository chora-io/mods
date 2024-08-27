package v1

import (
	"strconv"
	"strings"
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdateGovernorMetadata struct {
	t   gocuke.TestingT
	msg *MsgUpdateGovernorMetadata
	err error
}

func TestMsgUpdateGovernorMetadata(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateGovernorMetadata{}).
		Path("./msg_update_governor_metadata.feature").
		Run()
}

func (s *msgUpdateGovernorMetadata) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdateGovernorMetadata) Message(a gocuke.DocString) {
	s.msg = &MsgUpdateGovernorMetadata{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdateGovernorMetadata) NewMetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.NewMetadata = strings.Repeat("x", int(length))
}

func (s *msgUpdateGovernorMetadata) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateGovernorMetadata) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateGovernorMetadata) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
