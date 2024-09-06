package v1

import (
	"strconv"
	"strings"
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type subject struct {
	t       gocuke.TestingT
	subject *Subject
	err     error
}

func TestSubject(t *testing.T) {
	gocuke.NewRunner(t, &subject{}).
		Path("./state_subject.feature").
		Run()
}

func (s *subject) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *subject) Subject(a gocuke.DocString) {
	s.subject = &Subject{}
	err := jsonpb.UnmarshalString(a.Content, s.subject)
	require.NoError(s.t, err)
}

func (s *subject) MetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.subject.Metadata = strings.Repeat("x", int(length))
}

func (s *subject) ValidateSubject() {
	s.err = s.subject.Validate()
}

func (s *subject) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *subject) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
