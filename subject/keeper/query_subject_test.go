package keeper

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	subjectv1 "github.com/chora-io/mods/subject/api/v1"
	v1 "github.com/chora-io/mods/subject/types/v1"
)

type querySubject struct {
	*baseSuite
	res *v1.QuerySubjectResponse
	err error
}

func TestQuerySubject(t *testing.T) {
	gocuke.NewRunner(t, &querySubject{}).
		Path("./query_subject.feature").
		Run()
}

func (s *querySubject) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *querySubject) Subject(a gocuke.DocString) {
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

func (s *querySubject) QuerySubject(a gocuke.DocString) {
	var req v1.QuerySubjectRequest
	err := jsonpb.UnmarshalString(a.Content, &req)
	require.NoError(s.t, err)

	s.res, s.err = s.k.Subject(s.sdkCtx, &req)
}

func (s *querySubject) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *querySubject) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *querySubject) ExpectResponse(a gocuke.DocString) {
	var expected v1.QuerySubjectResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
