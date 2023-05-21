package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type validatorMissedBlocks struct {
	t                     gocuke.TestingT
	validatorMissedBlocks *ValidatorMissedBlocks
	err                   error
}

func TestValidatorMissedBlocks(t *testing.T) {
	gocuke.NewRunner(t, &validatorMissedBlocks{}).
		Path("./features/state_validator_missed_blocks.feature").
		Run()
}

func (s *validatorMissedBlocks) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *validatorMissedBlocks) ValidatorMissedBlocks(a gocuke.DocString) {
	s.validatorMissedBlocks = &ValidatorMissedBlocks{}
	err := jsonpb.UnmarshalString(a.Content, s.validatorMissedBlocks)
	require.NoError(s.t, err)
}

func (s *validatorMissedBlocks) ValidateValidatorMissedBlocks() {
	s.err = s.validatorMissedBlocks.Validate()
}

func (s *validatorMissedBlocks) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *validatorMissedBlocks) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
