package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdateCurator struct {
	t   gocuke.TestingT
	msg *MsgUpdateCurator
	err error
}

func TestMsgUpdateCurator(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateCurator{}).
		Path("./features/msg_update_curator.feature").
		Run()
}

func (s *msgUpdateCurator) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdateCurator) Message(a gocuke.DocString) {
	s.msg = &MsgUpdateCurator{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdateCurator) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateCurator) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateCurator) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
