package v1

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdateContentCurator struct {
	t   gocuke.TestingT
	msg *MsgUpdateContentCurator
	err error
}

func TestMsgUpdateContentCurator(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateContentCurator{}).
		Path("./msg_update_content_curator.feature").
		Run()
}

func (s *msgUpdateContentCurator) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdateContentCurator) Message(a gocuke.DocString) {
	s.msg = &MsgUpdateContentCurator{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdateContentCurator) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateContentCurator) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateContentCurator) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
