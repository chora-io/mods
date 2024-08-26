package v1

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdate struct {
	t   gocuke.TestingT
	msg *MsgUpdateAdmin
	err error
}

func TestMsgUpdateAdmin(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdate{}).
		Path("./msg_update_admin.feature").
		Run()
}

func (s *msgUpdate) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdate) Message(a gocuke.DocString) {
	s.msg = &MsgUpdateAdmin{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdate) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdate) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdate) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
