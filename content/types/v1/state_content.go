package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/chora-io/mods/content/errors"
)

// Validate validates Content.
func (m *Content) Validate() error {

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(m.Curator).String()); err != nil {
		return errors.ErrParse.Wrapf("curator: %s", err)
	}

	if m.Hash == "" {
		return errors.ErrParse.Wrap("hash: empty string is not allowed")
	}

	if len(m.Hash) > HashMaxLength {
		return errors.ErrParse.Wrapf("hash: exceeds max length %d", HashMaxLength)
	}

	return nil
}
