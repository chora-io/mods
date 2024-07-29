package v1

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type authority struct {
	t         gocuke.TestingT
	authority *Authority
	err       error
}

func TestAuthority(t *testing.T) {
	gocuke.NewRunner(t, &authority{}).
		Path("./features/state_authority.feature").
		Run()
}

func (s *authority) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *authority) Authority(a gocuke.DocString) {
	s.authority = &Authority{}
	err := jsonpb.UnmarshalString(a.Content, s.authority)
	require.NoError(s.t, err)
}

func (s *authority) ValidateAuthority() {
	s.err = s.authority.Validate()
}

func (s *authority) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *authority) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
