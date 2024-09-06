package keeper

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	subjectv1 "github.com/chora-io/mods/subject/api/v1"
	v1 "github.com/chora-io/mods/subject/types/v1"
)

type querySubjects struct {
	*baseSuite
	res *v1.QuerySubjectsResponse
	err error
}

func TestQuerySubjects(t *testing.T) {
	gocuke.NewRunner(t, &querySubjects{}).
		Path("./query_subjects.feature").
		Run()
}

func (s *querySubjects) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *querySubjects) Subject(a gocuke.DocString) {
	var subject subjectv1.Subject
	err := jsonpb.UnmarshalString(a.Content, &subject)
	require.NoError(s.t, err)

	err = s.k.ss.SubjectTable().Insert(s.sdkCtx, &subjectv1.Subject{
		Steward:  subject.Steward,
		Address:  subject.Address,
		Metadata: subject.Metadata,
	})
	require.NoError(s.t, err)
}

func (s *querySubjects) QuerySubjects(a gocuke.DocString) {
	var req v1.QuerySubjectsRequest
	err := jsonpb.UnmarshalString(a.Content, &req)
	require.NoError(s.t, err)

	s.res, s.err = s.k.Subjects(s.sdkCtx, &req)
}

func (s *querySubjects) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *querySubjects) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *querySubjects) ExpectResponse(a gocuke.DocString) {
	var expected v1.QuerySubjectsResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
