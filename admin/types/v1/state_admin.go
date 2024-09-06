package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/chora-io/mods/admin/errors"
)

// Validate validates Admin.
func (m *Admin) Validate() error {
	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(m.Admin).String()); err != nil {
		return errors.ErrParse.Wrapf("address: %s", err)
	}

	return nil
}
