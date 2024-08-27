package v1

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type validatorSigningInfo struct {
	t                    gocuke.TestingT
	validatorSigningInfo *ValidatorSigningInfo
	err                  error
}

func TestValidatorSigningInfo(t *testing.T) {
	gocuke.NewRunner(t, &validatorSigningInfo{}).
		Path("./state_validator_signing_info.feature").
		Run()
}

func (s *validatorSigningInfo) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *validatorSigningInfo) ValidatorSigningInfo(a gocuke.DocString) {
	s.validatorSigningInfo = &ValidatorSigningInfo{}
	err := jsonpb.UnmarshalString(a.Content, s.validatorSigningInfo)
	require.NoError(s.t, err)
}

func (s *validatorSigningInfo) ValidateValidatorSigningInfo() {
	s.err = s.validatorSigningInfo.Validate()
}

func (s *validatorSigningInfo) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *validatorSigningInfo) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
