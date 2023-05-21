package server

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	validatorv1 "github.com/choraio/mods/validator/api/v1"
	v1 "github.com/choraio/mods/validator/types/v1"
	"github.com/choraio/mods/validator/utils"
)

type msgAdd struct {
	*baseSuite
	res *v1.MsgAddResponse
	err error
}

func TestMsgAdd(t *testing.T) {
	gocuke.NewRunner(t, &msgAdd{}).
		Path("./features/msg_add.feature").
		Run()
}

func (s *msgAdd) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgAdd) Authority(a string) {
	require.Equal(s.t, s.authority.String(), a)
}

func (s *msgAdd) MsgAdd(a gocuke.DocString) {
	var msg v1.MsgAdd
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.srv.Add(s.ctx, &msg)
}

func (s *msgAdd) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgAdd) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgAdd) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgAddResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgAdd) ExpectStateValidator(a gocuke.DocString) {
	var expected validatorv1.Validator
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.srv.ss.ValidatorTable().Get(s.ctx, expected.Address)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Address, actual.Address)
	require.Equal(s.t, expected.Metadata, actual.Metadata)
}

func (s *msgAdd) ExpectEventAdd(a gocuke.DocString) {
	var expected v1.EventAdd
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
