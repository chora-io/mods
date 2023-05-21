package v1

import (
	"strconv"
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type content struct {
	t       gocuke.TestingT
	content *Validator
	err     error
}

func TestValidator(t *testing.T) {
	gocuke.NewRunner(t, &content{}).
		Path("./features/state_validator.feature").
		Run()
}

func (s *content) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *content) Validator(a gocuke.DocString) {
	s.content = &Validator{}
	err := jsonpb.UnmarshalString(a.Content, s.content)
	require.NoError(s.t, err)
}

func (s *content) MetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.content.Metadata = strings.Repeat("x", int(length))
}

func (s *content) ValidateValidator() {
	s.err = s.content.Validate()
}

func (s *content) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *content) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
