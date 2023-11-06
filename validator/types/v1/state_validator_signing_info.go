package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/chora-io/mods/validator/errors"
)

// Validate validates ValidatorSigningInfo.
func (m *ValidatorSigningInfo) Validate() error {
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return errors.ErrParse.Wrapf("address: %s", err)
	}

	return nil
}
