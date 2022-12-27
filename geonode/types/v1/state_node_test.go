package v1

import (
	"strconv"
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type node struct {
	t    gocuke.TestingT
	node *Node
	err  error
}

func TestNode(t *testing.T) {
	gocuke.NewRunner(t, &node{}).
		Path("./features/state_node.feature").
		Run()
}

func (s *node) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *node) Node(a gocuke.DocString) {
	s.node = &Node{}
	err := jsonpb.UnmarshalString(a.Content, s.node)
	require.NoError(s.t, err)
}

func (s *node) MetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.node.Metadata = strings.Repeat("x", int(length))
}

func (s *node) ValidateNode() {
	s.err = s.node.Validate()
}

func (s *node) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *node) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
