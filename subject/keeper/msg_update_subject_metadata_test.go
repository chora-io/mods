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

type msgUpdateSubjectMetadata struct {
	*baseSuite
	res *v1.MsgUpdateSubjectMetadataResponse
	err error
}

func TestMsgUpdateSubjectMetadata(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateSubjectMetadata{}).
		Path("./msg_update_subject_metadata.feature").
		Run()
}

func (s *msgUpdateSubjectMetadata) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgUpdateSubjectMetadata) Subject(a gocuke.DocString) {
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

func (s *msgUpdateSubjectMetadata) MsgUpdateSubjectMetadata(a gocuke.DocString) {
	var msg v1.MsgUpdateSubjectMetadata
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.UpdateSubjectMetadata(s.sdkCtx, &msg)
}

func (s *msgUpdateSubjectMetadata) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateSubjectMetadata) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgUpdateSubjectMetadata) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgUpdateSubjectMetadataResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgUpdateSubjectMetadata) ExpectStateSubject(a gocuke.DocString) {
	var expected subjectv1.Subject
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.k.ss.SubjectTable().Get(s.sdkCtx, expected.Address)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Address, actual.Address)
	require.Equal(s.t, expected.Steward, actual.Steward)
	require.Equal(s.t, expected.Metadata, actual.Metadata)
}

func (s *msgUpdateSubjectMetadata) ExpectEventUpdateSubjectMetadata(a gocuke.DocString) {
	var expected v1.EventUpdateSubjectMetadata
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
