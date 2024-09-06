package v1

import (
	"strconv"
	"strings"
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgCreateValidator struct {
	t   gocuke.TestingT
	msg *MsgCreateValidator
	err error
}

func TestMsgCreateValidator(t *testing.T) {
	gocuke.NewRunner(t, &msgCreateValidator{}).
		Path("./msg_create_validator.feature").
		Run()
}

func (s *msgCreateValidator) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgCreateValidator) Message(a gocuke.DocString) {
	s.msg = &MsgCreateValidator{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgCreateValidator) MetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.Metadata = strings.Repeat("x", int(length))
}

func (s *msgCreateValidator) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgCreateValidator) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgCreateValidator) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
