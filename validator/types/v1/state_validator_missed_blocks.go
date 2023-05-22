package v1

import (
	"github.com/choraio/mods/validator/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Validate validates ValidatorSigningInfo.
func (m *ValidatorSigningInfo) Validate() error {
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return errors.ErrParse.Wrapf("address: %s", err)
	}

	return nil
}
