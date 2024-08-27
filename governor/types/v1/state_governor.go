package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/chora-io/mods/governor/errors"
)

// Validate validates Governor.
func (m *Governor) Validate() error {
	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(m.Address).String()); err != nil {
		return errors.ErrParse.Wrapf("address: %s", err)
	}

	if m.Metadata == "" {
		return errors.ErrParse.Wrap("metadata: empty string is not allowed")
	}

	if len(m.Metadata) > MetadataMaxLength {
		return errors.ErrParse.Wrapf("metadata: exceeds max length %d", MetadataMaxLength)
	}

	return nil
}
