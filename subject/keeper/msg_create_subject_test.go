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

type msgCreateSubject struct {
	*baseSuite
	res *v1.MsgCreateSubjectResponse
	err error
}

func TestMsgCreateSubject(t *testing.T) {
	gocuke.NewRunner(t, &msgCreateSubject{}).
		Path("./msg_create_subject.feature").
		Run()
}

func (s *msgCreateSubject) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgCreateSubject) SubjectSequence(a gocuke.DocString) {
	var as subjectv1.SubjectSequence
	err := jsonpb.UnmarshalString(a.Content, &as)
	require.NoError(s.t, err)

	err = s.k.ss.SubjectSequenceTable().Save(s.sdkCtx, &subjectv1.SubjectSequence{
		Sequence: as.Sequence,
	})
	require.NoError(s.t, err)
}

func (s *msgCreateSubject) MsgCreateSubject(a gocuke.DocString) {
	var msg v1.MsgCreateSubject
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.CreateSubject(s.sdkCtx, &msg)
}

func (s *msgCreateSubject) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgCreateSubject) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgCreateSubjectResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgCreateSubject) ExpectStateSubject(a gocuke.DocString) {
	var expected subjectv1.Subject
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.k.ss.SubjectTable().Get(s.sdkCtx, expected.Address)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Address, actual.Address)
	require.Equal(s.t, expected.Steward, actual.Steward)
	require.Equal(s.t, expected.Metadata, actual.Metadata)
}

func (s *msgCreateSubject) ExpectEventCreateSubject(a gocuke.DocString) {
	var expected v1.EventCreateSubject
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
