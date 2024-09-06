package keeper

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	subjectv1 "github.com/chora-io/mods/subject/api/v1"
	v1 "github.com/chora-io/mods/subject/types/v1"
)

type querySubjectsBySteward struct {
	*baseSuite
	res *v1.QuerySubjectsByStewardResponse
	err error
}

func TestQuerySubjectsBySteward(t *testing.T) {
	gocuke.NewRunner(t, &querySubjectsBySteward{}).
		Path("./query_subjects_by_steward.feature").
		Run()
}

func (s *querySubjectsBySteward) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *querySubjectsBySteward) Subject(a gocuke.DocString) {
	var subject subjectv1.Subject
	err := jsonpb.UnmarshalString(a.Content, &subject)
	require.NoError(s.t, err)

	err = s.k.ss.SubjectTable().Insert(s.sdkCtx, &subjectv1.Subject{
		Address:  subject.Address,
		Steward:  subject.Steward,
		Metadata: subject.Metadata,
	})
	require.NoError(s.t, err)
}

func (s *querySubjectsBySteward) QuerySubjectsBySteward(a gocuke.DocString) {
	var req v1.QuerySubjectsByStewardRequest
	err := jsonpb.UnmarshalString(a.Content, &req)
	require.NoError(s.t, err)

	s.res, s.err = s.k.SubjectsBySteward(s.sdkCtx, &req)
}

func (s *querySubjectsBySteward) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *querySubjectsBySteward) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *querySubjectsBySteward) ExpectResponse(a gocuke.DocString) {
	var expected v1.QuerySubjectsByStewardResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
