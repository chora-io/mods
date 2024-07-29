package keeper

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	authorityv1 "github.com/chora-io/mods/authority/api/v1"
	v1 "github.com/chora-io/mods/authority/types/v1"
)

type queryAuthority struct {
	*baseSuite
	res *v1.QueryAuthorityResponse
	err error
}

func TestQueryAuthority(t *testing.T) {
	gocuke.NewRunner(t, &queryAuthority{}).
		Path("./features/query_authority.feature").
		Run()
}

func (s *queryAuthority) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *queryAuthority) Authority(a gocuke.DocString) {
	var authority authorityv1.Authority
	err := jsonpb.UnmarshalString(a.Content, &authority)
	require.NoError(s.t, err)

	err = s.k.ss.AuthorityTable().Save(s.sdkCtx, &authorityv1.Authority{
		Address: authority.Address,
	})
	require.NoError(s.t, err)
}

func (s *queryAuthority) QueryAuthority(a gocuke.DocString) {
	var req v1.QueryAuthorityRequest
	err := jsonpb.UnmarshalString(a.Content, &req)
	require.NoError(s.t, err)

	s.res, s.err = s.k.Authority(s.sdkCtx, &req)
}

func (s *queryAuthority) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *queryAuthority) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *queryAuthority) ExpectResponse(a gocuke.DocString) {
	var expected v1.QueryAuthorityResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
