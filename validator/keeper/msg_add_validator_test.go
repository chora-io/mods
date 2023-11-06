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

type msgAddValidator struct {
	*baseSuite
	res *v1.MsgAddValidatorResponse
	err error
}

func TestMsgAddValidator(t *testing.T) {
	gocuke.NewRunner(t, &msgAddValidator{}).
		Path("./features/msg_add_validator.feature").
		Run()
}

func (s *msgAddValidator) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgAddValidator) Authority(a string) {
	require.Equal(s.t, s.authority.String(), a)
}

func (s *msgAddValidator) MsgAddValidator(a gocuke.DocString) {
	var msg v1.MsgAddValidator
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.AddValidator(s.sdkCtx, &msg)
}

func (s *msgAddValidator) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgAddValidator) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgAddValidator) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgAddValidatorResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgAddValidator) ExpectStateValidator(a gocuke.DocString) {
	var expected validatorv1.Validator
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, err := s.k.ss.ValidatorTable().Get(s.sdkCtx, expected.Address)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Address, actual.Address)
	require.Equal(s.t, expected.Metadata, actual.Metadata)
}

func (s *msgAddValidator) ExpectEventAdd(a gocuke.DocString) {
	var expected v1.EventAddValidator
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
