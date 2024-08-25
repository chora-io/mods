package keeper

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	ecosystemv1 "github.com/chora-io/mods/ecosystem/api/v1"
	v1 "github.com/chora-io/mods/ecosystem/types/v1"
)

type queryNodes struct {
	*baseSuite
	res *v1.QueryNodesResponse
	err error
}

func TestQueryNodes(t *testing.T) {
	gocuke.NewRunner(t, &queryNodes{}).
		Path("./features/query_nodes.feature").
		Run()
}

func (s *queryNodes) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *queryNodes) Node(a gocuke.DocString) {
	var node ecosystemv1.Node
	err := jsonpb.UnmarshalString(a.Content, &node)
	require.NoError(s.t, err)

	id, err := s.k.ss.NodeTable().InsertReturningId(s.sdkCtx, &ecosystemv1.Node{
		Curator:  node.Curator,
		Metadata: node.Metadata,
	})
	require.NoError(s.t, err)
	require.Equal(s.t, node.Id, id)
}

func (s *queryNodes) QueryNodes(a gocuke.DocString) {
	var req v1.QueryNodesRequest
	err := jsonpb.UnmarshalString(a.Content, &req)
	require.NoError(s.t, err)

	s.res, s.err = s.k.Nodes(s.sdkCtx, &req)
}

func (s *queryNodes) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *queryNodes) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *queryNodes) ExpectResponse(a gocuke.DocString) {
	var expected v1.QueryNodesResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
