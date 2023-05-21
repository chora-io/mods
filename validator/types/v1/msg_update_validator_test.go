package v1

import (
	"strconv"
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdateValidator struct {
	t   gocuke.TestingT
	msg *MsgUpdateValidator
	err error
}

func TestMsgUpdateValidator(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateValidator{}).
		Path("./features/msg_update_validator.feature").
		Run()
}

func (s *msgUpdateValidator) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdateValidator) Message(a gocuke.DocString) {
	s.msg = &MsgUpdateValidator{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdateValidator) NewMetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.NewMetadata = strings.Repeat("x", int(length))
}

func (s *msgUpdateValidator) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateValidator) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateValidator) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
