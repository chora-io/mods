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

type msgUpdateSubjectSteward struct {
	*baseSuite
	res *v1.MsgUpdateSubjectStewardResponse
	err error
}

func TestMsgUpdateSubjectSteward(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateSubjectSteward{}).
		Path("./msg_update_subject_steward.feature").
		Run()
}

func (s *msgUpdateSubjectSteward) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgUpdateSubjectSteward) Subject(a gocuke.DocString) {
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

func (s *msgUpdateSubjectSteward) MsgUpdateSubjectSteward(a gocuke.DocString) {
	var msg v1.MsgUpdateSubjectSteward
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.UpdateSubjectSteward(s.sdkCtx, &msg)
}

func (s *msgUpdateSubjectSteward) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateSubjectSteward) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgUpdateSubjectSteward) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgUpdateSubjectStewardResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgUpdateSubjectSteward) ExpectStateSubject(a gocuke.DocString) {
	var expected subjectv1.Subject
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.k.ss.SubjectTable().Get(s.sdkCtx, expected.Address)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Address, actual.Address)
	require.Equal(s.t, expected.Steward, actual.Steward)
	require.Equal(s.t, expected.Metadata, actual.Metadata)
}

func (s *msgUpdateSubjectSteward) ExpectEventUpdateSubjectSteward(a gocuke.DocString) {
	var expected v1.EventUpdateSubjectSteward
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
