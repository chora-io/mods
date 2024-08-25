package keeper

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	adminv1 "github.com/chora-io/mods/admin/api/v1"
	v1 "github.com/chora-io/mods/admin/types/v1"
)

type queryAdmin struct {
	*baseSuite
	res *v1.QueryAdminResponse
	err error
}

func TestQueryAdmin(t *testing.T) {
	gocuke.NewRunner(t, &queryAdmin{}).
		Path("./features/query_admin.feature").
		Run()
}

func (s *queryAdmin) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *queryAdmin) Admin(a gocuke.DocString) {
	var admin adminv1.Admin
	err := jsonpb.UnmarshalString(a.Content, &admin)
	require.NoError(s.t, err)

	err = s.k.ss.AdminTable().Save(s.sdkCtx, &adminv1.Admin{
		Address: admin.Address,
	})
	require.NoError(s.t, err)
}

func (s *queryAdmin) QueryAdmin(a gocuke.DocString) {
	var req v1.QueryAdminRequest
	err := jsonpb.UnmarshalString(a.Content, &req)
	require.NoError(s.t, err)

	s.res, s.err = s.k.Admin(s.sdkCtx, &req)
}

func (s *queryAdmin) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *queryAdmin) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *queryAdmin) ExpectResponse(a gocuke.DocString) {
	var expected v1.QueryAdminResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
