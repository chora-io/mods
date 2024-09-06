package keeper

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	validatorv1 "github.com/chora-io/mods/validator/api/v1"
	v1 "github.com/chora-io/mods/validator/types/v1"
	"github.com/chora-io/mods/validator/utils"
)

type msgCreateValidator struct {
	*baseSuite
	res *v1.MsgCreateValidatorResponse
	err error
}

func TestMsgCreateValidator(t *testing.T) {
	gocuke.NewRunner(t, &msgCreateValidator{}).
		Path("./msg_create_validator.feature").
		Run()
}

func (s *msgCreateValidator) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgCreateValidator) Admin(a string) {
	require.Equal(s.t, s.admin.String(), a)
}

func (s *msgCreateValidator) MsgCreateValidator(a gocuke.DocString) {
	var msg v1.MsgCreateValidator
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.CreateValidator(s.sdkCtx, &msg)
}

func (s *msgCreateValidator) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgCreateValidator) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgCreateValidator) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgCreateValidatorResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgCreateValidator) ExpectStateValidator(a gocuke.DocString) {
	var expected validatorv1.Validator
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.k.ss.ValidatorTable().Get(s.sdkCtx, expected.Address)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Address, actual.Address)
	require.Equal(s.t, expected.Metadata, actual.Metadata)
}

func (s *msgCreateValidator) ExpectEventAdd(a gocuke.DocString) {
	var expected v1.EventCreateValidator
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
