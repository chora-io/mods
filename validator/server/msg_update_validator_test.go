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

type msgUpdateValidator struct {
	*baseSuite
	res *v1.MsgUpdateValidatorResponse
	err error
}

func TestMsgUpdateValidator(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateValidator{}).
		Path("./features/msg_update_validator.feature").
		Run()
}

func (s *msgUpdateValidator) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgUpdateValidator) Validator(a gocuke.DocString) {
	var validator validatorv1.Validator
	err := jsonpb.UnmarshalString(a.Content, &validator)
	require.NoError(s.t, err)

	err = s.srv.ss.ValidatorTable().Insert(s.ctx, &validatorv1.Validator{
		Address: validator.Address,
	})
	require.NoError(s.t, err)
}

func (s *msgUpdateValidator) MsgUpdateValidator(a gocuke.DocString) {
	var msg v1.MsgUpdateValidator
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.srv.UpdateValidator(s.ctx, &msg)
}

func (s *msgUpdateValidator) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateValidator) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgUpdateValidator) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgUpdateValidatorResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgUpdateValidator) ExpectStateValidator(a gocuke.DocString) {
	var expected validatorv1.Validator
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.srv.ss.ValidatorTable().Get(s.ctx, expected.Address)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Address, actual.Address)
	require.Equal(s.t, expected.Metadata, actual.Metadata)
}

func (s *msgUpdateValidator) ExpectEventUpdateValidator(a gocuke.DocString) {
	var expected v1.EventUpdateValidator
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}