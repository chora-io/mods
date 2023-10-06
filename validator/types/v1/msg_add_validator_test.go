package v1

import (
	"strconv"
	"strings"
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgAddValidator struct {
	t   gocuke.TestingT
	msg *MsgAddValidator
	err error
}

func TestMsgAddValidator(t *testing.T) {
	gocuke.NewRunner(t, &msgAddValidator{}).
		Path("./features/msg_add_validator.feature").
		Run()
}

func (s *msgAddValidator) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgAddValidator) Message(a gocuke.DocString) {
	s.msg = &MsgAddValidator{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgAddValidator) MetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.Metadata = strings.Repeat("x", int(length))
}

func (s *msgAddValidator) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgAddValidator) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgAddValidator) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
