package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/choraio/mods/voucher/errors"
)

// Validate validates Voucher.
func (m *Voucher) Validate() error {
	if m.Id == 0 {
		return errors.ErrParse.Wrap("id: empty or zero is not allowed")
	}

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(m.Issuer).String()); err != nil {
		return errors.ErrParse.Wrapf("issuer: %s", err)
	}

	if m.Metadata == "" {
		return errors.ErrParse.Wrap("metadata: empty string is not allowed")
	}

	if len(m.Metadata) > MetadataMaxLength {
		return errors.ErrParse.Wrapf("metadata: exceeds max length %d", MetadataMaxLength)
	}

	return nil
}
