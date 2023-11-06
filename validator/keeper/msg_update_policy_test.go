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

type msgUpdatePolicy struct {
	*baseSuite
	res *v1.MsgUpdatePolicyResponse
	err error
}

func TestMsgUpdatePolicy(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdatePolicy{}).
		Path("./features/msg_update_policy.feature").
		Run()
}

func (s *msgUpdatePolicy) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *msgUpdatePolicy) Authority(a string) {
	require.Equal(s.t, s.authority.String(), a)
}

func (s *msgUpdatePolicy) Policy(a gocuke.DocString) {
	var policy validatorv1.Policy
	err := jsonpb.UnmarshalString(a.Content, &policy)
	require.NoError(s.t, err)

	err = s.k.ss.PolicyTable().Save(s.sdkCtx, &policy)
	require.NoError(s.t, err)
}

func (s *msgUpdatePolicy) MsgUpdatePolicy(a gocuke.DocString) {
	var msg v1.MsgUpdatePolicy
	err := jsonpb.UnmarshalString(a.Content, &msg)
	require.NoError(s.t, err)

	s.res, s.err = s.k.UpdatePolicy(s.sdkCtx, &msg)
}

func (s *msgUpdatePolicy) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdatePolicy) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *msgUpdatePolicy) ExpectPolicy(a gocuke.DocString) {
	var expected validatorv1.Policy
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	policy, err := s.k.ss.PolicyTable().Get(s.sdkCtx)
	require.NoError(s.t, err)
	require.Equal(s.t, expected.SignedBlocksWindow, policy.SignedBlocksWindow)
	require.Equal(s.t, expected.MinSignedPerWindow, policy.MinSignedPerWindow)
}

func (s *msgUpdatePolicy) ExpectEventUpdatePolicy(a gocuke.DocString) {
	var expected v1.EventUpdatePolicy
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	actual, found := utils.GetEvent(&expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = utils.MatchEvent(&expected, actual)
	require.NoError(s.t, err)
}

func (s *msgUpdatePolicy) ExpectResponse(a gocuke.DocString) {
	var expected v1.MsgUpdatePolicyResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
