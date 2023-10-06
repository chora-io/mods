package v1

import (
	"strconv"
	"strings"
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type voucher struct {
	t       gocuke.TestingT
	voucher *Voucher
	err     error
}

func TestVoucher(t *testing.T) {
	gocuke.NewRunner(t, &voucher{}).
		Path("./features/state_voucher.feature").
		Run()
}

func (s *voucher) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *voucher) Voucher(a gocuke.DocString) {
	s.voucher = &Voucher{}
	err := jsonpb.UnmarshalString(a.Content, s.voucher)
	require.NoError(s.t, err)
}

func (s *voucher) MetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.voucher.Metadata = strings.Repeat("x", int(length))
}

func (s *voucher) ValidateVoucher() {
	s.err = s.voucher.Validate()
}

func (s *voucher) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *voucher) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}
