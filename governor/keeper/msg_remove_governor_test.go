package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	governorv1 "github.com/chora-io/mods/governor/api/v1"
	v1 "github.com/chora-io/mods/governor/types/v1"
	"github.com/chora-io/mods/governor/utils"
)

type msgRemoveGovernor struct {
	*baseSuite
	res *v1.MsgRemoveGovernorResponse
	err error
}

func TestMsgRemoveGovernor(t *testing.T) {
	gocuke.NewRunner(t, &msgRemoveGovernor{}).
		Path("./msg_remove_governor.feature").
		Run()
}

func (s *msgRemoveGovernor) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgRemoveGovernor) Admin(a string) {
	require.Equal(s.t, s.admin.String(), a)
}

func (s *msgRemoveGovernor) Governor(a gocuke.DocString) {
	var governor governorv1.Governor
	err := jsonpb.UnmarshalString(a.Content, &governor)
	require.NoError(s.t, err)

	err = s.k.ss.GovernorTable().Insert(s.sdkCtx, &governorv1.Governor{
		Address: governor.Address,
	})
	require.NoError(s.t, err)
}

func (s *msgRemoveGovernor) MsgRemoveGovernor(a gocuke.DocString) {
	var msg v1.MsgRemoveGovernor
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.RemoveGovernor(s.sdkCtx, &msg)
}

func (s *msgRemoveGovernor) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgRemoveGovernor) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgRemoveGovernor) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgRemoveGovernorResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}

func (s *msgRemoveGovernor) ExpectNoGovernorWithAddress(a string) {
	found, err := s.k.ss.GovernorTable().Has(s.sdkCtx, sdk.AccAddress(a))
	require.NoError(s.t, err)
	require.False(s.t, found)
}

func (s *msgRemoveGovernor) ExpectEventRemove(a gocuke.DocString) {
	var expected v1.EventRemoveGovernor
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}
