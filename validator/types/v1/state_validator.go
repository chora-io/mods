package v1

import (
	"github.com/choraio/mods/validator/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Validate validates Validator.
func (m *Validator) Validate() error {
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
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
