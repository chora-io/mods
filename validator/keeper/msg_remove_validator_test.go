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

type msgRemoveValidator struct {
	*baseSuite
	res *v1.MsgRemoveValidatorResponse
	err error
}

func TestMsgRemoveValidator(t *testing.T) {
	gocuke.NewRunner(t, &msgRemoveValidator{}).
		Path("./features/msg_remove_validator.feature").
		Run()
}

func (s *msgRemoveValidator) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgRemoveValidator) Authority(a string) {
	require.Equal(s.t, s.authority.String(), a)
}

func (s *msgRemoveValidator) Validator(a gocuke.DocString) {
	var validator validatorv1.Validator
	err := jsonpb.UnmarshalString(a.Content, &validator)
	require.NoError(s.t, err)

	err = s.k.ss.ValidatorTable().Insert(s.sdkCtx, &validatorv1.Validator{
		Address: validator.Address,
	})
	require.NoError(s.t, err)
}

func (s *msgRemoveValidator) ValidatorSigningInfo(a gocuke.DocString) {
	var missedBlocks validatorv1.ValidatorSigningInfo
	err := jsonpb.UnmarshalString(a.Content, &missedBlocks)
	require.NoError(s.t, err)

	err = s.k.ss.ValidatorSigningInfoTable().Insert(s.sdkCtx, &validatorv1.ValidatorSigningInfo{
		Address:      missedBlocks.Address,
		MissedBlocks: missedBlocks.MissedBlocks,
	})
	require.NoError(s.t, err)
}

func (s *msgRemoveValidator) MsgRemoveValidator(a gocuke.DocString) {
	var msg v1.MsgRemoveValidator
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.RemoveValidator(s.sdkCtx, &msg)
}

func (s *msgRemoveValidator) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgRemoveValidator) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgRemoveValidator) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgRemoveValidatorResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgRemoveValidator) ExpectNoValidatorWithAddress(a string) {
	found, err := s.k.ss.ValidatorTable().Has(s.sdkCtx, a)
	require.NoError(s.t, err)
	require.False(s.t, found)
}

func (s *msgRemoveValidator) ExpectNoValidatorSigningInfoWithAddress(a string) {
	found, err := s.k.ss.ValidatorSigningInfoTable().Has(s.sdkCtx, a)
	require.NoError(s.t, err)
	require.False(s.t, found)
}

func (s *msgRemoveValidator) ExpectEventRemove(a gocuke.DocString) {
	var expected v1.EventRemoveValidator
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
