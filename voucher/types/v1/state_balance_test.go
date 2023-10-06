package v1

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type balance struct {
	t       gocuke.TestingT
	balance *Balance
	err     error
}

func TestBalance(t *testing.T) {
	gocuke.NewRunner(t, &balance{}).
		Path("./features/state_balance.feature").
		Run()
}

func (s *balance) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *balance) Balance(a gocuke.DocString) {
	s.balance = &Balance{}
	err := jsonpb.UnmarshalString(a.Content, s.balance)
	require.NoError(s.t, err)
}

func (s *balance) ValidateBalance() {
	s.err = s.balance.Validate()
}

func (s *balance) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *balance) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
