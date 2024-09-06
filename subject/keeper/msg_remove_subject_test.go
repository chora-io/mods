package keeper

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	subjectv1 "github.com/chora-io/mods/subject/api/v1"
	v1 "github.com/chora-io/mods/subject/types/v1"
	"github.com/chora-io/mods/subject/utils"
)

type msgRemoveSubject struct {
	*baseSuite
	res *v1.MsgRemoveSubjectResponse
	err error
}

func TestMsgRemoveSubject(t *testing.T) {
	gocuke.NewRunner(t, &msgRemoveSubject{}).
		Path("./msg_remove_subject.feature").
		Run()
}

func (s *msgRemoveSubject) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgRemoveSubject) Subject(a gocuke.DocString) {
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

func (s *msgRemoveSubject) MsgRemoveSubject(a gocuke.DocString) {
	var msg v1.MsgRemoveSubject
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.RemoveSubject(s.sdkCtx, &msg)
}

func (s *msgRemoveSubject) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgRemoveSubject) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgRemoveSubject) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgRemoveSubjectResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgRemoveSubject) ExpectNoStateSubject(a gocuke.DocString) {
	var expected subjectv1.Subject
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.k.ss.SubjectTable().Get(s.sdkCtx, expected.Address)
	require.Nil(s.t, actual)
	require.EqualError(s.t, err, "not found")
}

func (s *msgRemoveSubject) ExpectEventRemoveSubject(a gocuke.DocString) {
	var expected v1.EventRemoveSubject
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
