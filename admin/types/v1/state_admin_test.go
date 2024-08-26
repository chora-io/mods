package v1

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type admin struct {
	t     gocuke.TestingT
	admin *Admin
	err   error
}

func TestAdmin(t *testing.T) {
	gocuke.NewRunner(t, &admin{}).
		Path("./state_admin.feature").
		Run()
}

func (s *admin) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *admin) Admin(a gocuke.DocString) {
	s.admin = &Admin{}
	err := jsonpb.UnmarshalString(a.Content, s.admin)
	require.NoError(s.t, err)
}

func (s *admin) ValidateAdmin() {
	s.err = s.admin.Validate()
}

func (s *admin) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *admin) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
