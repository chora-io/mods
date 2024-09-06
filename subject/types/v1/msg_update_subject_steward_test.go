package v1

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdateSubjectSteward struct {
	t   gocuke.TestingT
	msg *MsgUpdateSubjectSteward
	err error
}

func TestMsgUpdateSubjectSteward(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateSubjectSteward{}).
		Path("./msg_update_subject_steward.feature").
		Run()
}

func (s *msgUpdateSubjectSteward) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdateSubjectSteward) Message(a gocuke.DocString) {
	s.msg = &MsgUpdateSubjectSteward{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdateSubjectSteward) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateSubjectSteward) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateSubjectSteward) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
