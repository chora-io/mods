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

type msgRemove struct {
	*baseSuite
	res *v1.MsgRemoveResponse
	err error
}

func TestMsgRemove(t *testing.T) {
	gocuke.NewRunner(t, &msgRemove{}).
		Path("./features/msg_remove.feature").
		Run()
}

func (s *msgRemove) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgRemove) Authority(a string) {
	require.Equal(s.t, s.authority.String(), a)
}

func (s *msgRemove) Validator(a gocuke.DocString) {
	var validator validatorv1.Validator
	err := jsonpb.UnmarshalString(a.Content, &validator)
	require.NoError(s.t, err)

	err = s.srv.ss.ValidatorTable().Insert(s.ctx, &validatorv1.Validator{
		Address: validator.Address,
	})
	require.NoError(s.t, err)
}

func (s *msgRemove) MsgRemove(a gocuke.DocString) {
	var msg v1.MsgRemove
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.srv.Remove(s.ctx, &msg)
}

func (s *msgRemove) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgRemove) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgRemove) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgRemoveResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgRemove) ExpectNoValidatorWithAddress(a string) {
	found, err := s.srv.ss.ValidatorTable().Has(s.ctx, a)
	require.NoError(s.t, err)
	require.False(s.t, found)
}

func (s *msgRemove) ExpectEventRemove(a gocuke.DocString) {
	var expected v1.EventRemove
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
