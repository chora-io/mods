package v1

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdateIssuer struct {
	t   gocuke.TestingT
	msg *MsgUpdateIssuer
	err error
}

func TestMsgUpdateIssuer(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateIssuer{}).
		Path("./msg_update_issuer.feature").
		Run()
}

func (s *msgUpdateIssuer) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdateIssuer) Message(a gocuke.DocString) {
	s.msg = &MsgUpdateIssuer{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdateIssuer) ValidateMessage() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateIssuer) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateIssuer) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
