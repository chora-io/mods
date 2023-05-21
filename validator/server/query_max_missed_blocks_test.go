package server

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	validatorv1 "github.com/choraio/mods/validator/api/v1"
	v1 "github.com/choraio/mods/validator/types/v1"
)

type queryMaxMissedBlocks struct {
	*baseSuite
	res *v1.QueryMaxMissedBlocksResponse
	err error
}

func TestQueryMaxMissedBlocks(t *testing.T) {
	gocuke.NewRunner(t, &queryMaxMissedBlocks{}).
		Path("./features/query_max_missed_blocks.feature").
		Run()
}

func (s *queryMaxMissedBlocks) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *queryMaxMissedBlocks) MaxMissedBlocks(a gocuke.DocString) {
	var maxMissedBlocks validatorv1.MaxMissedBlocks
	err := jsonpb.UnmarshalString(a.Content, &maxMissedBlocks)
	require.NoError(s.t, err)

	err = s.srv.ss.MaxMissedBlocksTable().Save(s.ctx, &validatorv1.MaxMissedBlocks{
		MaxMissedBlocks: maxMissedBlocks.MaxMissedBlocks,
	})
	require.NoError(s.t, err)
}

func (s *queryMaxMissedBlocks) QueryMaxMissedBlocks() {
	var req v1.QueryMaxMissedBlocksRequest
	s.res, s.err = s.srv.MaxMissedBlocks(s.ctx, &req)
}

func (s *queryMaxMissedBlocks) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *queryMaxMissedBlocks) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *queryMaxMissedBlocks) ExpectResponse(a gocuke.DocString) {
	var expected v1.QueryMaxMissedBlocksResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
