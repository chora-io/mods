package server

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	geonodev1 "github.com/choraio/mods/geonode/api/v1"
	v1 "github.com/choraio/mods/geonode/types/v1"
)

type queryNodeByCurator struct {
	*baseSuite
	res *v1.QueryNodeByCuratorResponse
	err error
}

func TestQueryNodeByCurator(t *testing.T) {
	gocuke.NewRunner(t, &queryNodeByCurator{}).
		Path("./features/query_node_by_curator.feature").
		Run()
}

func (s *queryNodeByCurator) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *queryNodeByCurator) Node(a gocuke.DocString) {
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

func (s *queryNodeByCurator) QueryNodeByCurator(a gocuke.DocString) {
	var req v1.QueryNodeByCuratorRequest
	err := jsonpb.UnmarshalString(a.Content, &req)
	require.NoError(s.t, err)

	s.res, s.err = s.srv.NodeByCurator(s.ctx, &req)
}

func (s *queryNodeByCurator) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *queryNodeByCurator) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *queryNodeByCurator) ExpectResponse(a gocuke.DocString) {
	var expected v1.QueryNodeByCuratorResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
