package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgRemoveValidator struct {
	t   gocuke.TestingT
	msg *MsgRemoveValidator
	err error
}

func TestMsgRemoveValidator(t *testing.T) {
	gocuke.NewRunner(t, &msgRemoveValidator{}).
		Path("./features/msg_remove_validator.feature").
		Run()
}

func (s *msgRemoveValidator) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgRemoveValidator) Message(a gocuke.DocString) {
	s.msg = &MsgRemoveValidator{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgRemoveValidator) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgRemoveValidator) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgRemoveValidator) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
