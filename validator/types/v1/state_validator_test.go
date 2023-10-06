package v1

import (
	"strconv"
	"strings"
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type validator struct {
	t         gocuke.TestingT
	validator *Validator
	err       error
}

func TestValidator(t *testing.T) {
	gocuke.NewRunner(t, &validator{}).
		Path("./features/state_validator.feature").
		Run()
}

func (s *validator) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *validator) Validator(a gocuke.DocString) {
	s.validator = &Validator{}
	err := jsonpb.UnmarshalString(a.Content, s.validator)
	require.NoError(s.t, err)
}

func (s *validator) MetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.validator.Metadata = strings.Repeat("x", int(length))
}

func (s *validator) ValidateValidator() {
	s.err = s.validator.Validate()
}

func (s *validator) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *validator) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
