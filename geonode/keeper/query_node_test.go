package keeper

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	geonodev1 "github.com/choraio/mods/geonode/api/v1"
	v1 "github.com/choraio/mods/geonode/types/v1"
)

type queryNode struct {
	*baseSuite
	res *v1.QueryNodeResponse
	err error
}

func TestQueryNode(t *testing.T) {
	gocuke.NewRunner(t, &queryNode{}).
		Path("./features/query_node.feature").
		Run()
}

func (s *queryNode) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *queryNode) Node(a gocuke.DocString) {
	var node geonodev1.Node
	err := jsonpb.UnmarshalString(a.Content, &node)
	require.NoError(s.t, err)

	id, err := s.k.ss.NodeTable().InsertReturningId(s.sdkCtx, &geonodev1.Node{
		Curator:  node.Curator,
		Metadata: node.Metadata,
	})
	require.NoError(s.t, err)
	require.Equal(s.t, node.Id, id)
}

func (s *queryNode) QueryNode(a gocuke.DocString) {
	var req v1.QueryNodeRequest
	err := jsonpb.UnmarshalString(a.Content, &req)
	require.NoError(s.t, err)

	s.res, s.err = s.k.Node(s.sdkCtx, &req)
}

func (s *queryNode) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *queryNode) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *queryNode) ExpectResponse(a gocuke.DocString) {
	var expected v1.QueryNodeResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
