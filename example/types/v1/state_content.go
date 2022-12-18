package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/choraio/mods/example/errors"
)

// Validate validates Content.
func (m *Content) Validate() error {
	if m.Id == 0 {
		return errors.ErrParse.Wrap("id: empty or zero is not allowed")
	}

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(m.Creator).String()); err != nil {
		return errors.ErrParse.Wrapf("creator: %s", err)
	}

	if m.Hash == "" {
		return errors.ErrParse.Wrap("hash: empty string is not allowed")
	}

	if len(m.Hash) > HashMaxLength {
		return errors.ErrParse.Wrapf("hash: exceeds max length %d", HashMaxLength)
	}

	return nil
}
