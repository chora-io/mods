package server

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	geonodev1 "github.com/choraio/mods/geonode/api/v1"
	v1 "github.com/choraio/mods/geonode/types/v1"
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
	var node geonodev1.Node
	err := jsonpb.UnmarshalString(a.Content, &node)
	require.NoError(s.t, err)

	id, err := s.srv.ss.NodeTable().InsertReturningID(s.ctx, &geonodev1.Node{
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

	s.res, s.err = s.srv.Nodes(s.ctx, &req)
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
