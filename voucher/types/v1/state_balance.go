package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types/v2/math"

	"github.com/choraio/mods/voucher/errors"
)

// Validate validates Balance.
func (m *Balance) Validate() error {
	if m.Id == 0 {
		return errors.ErrParse.Wrap("id: empty or zero is not allowed")
	}

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(m.Address).String()); err != nil {
		return errors.ErrParse.Wrapf("address: %s", err)
	}

	if m.Amount == "" {
		return errors.ErrParse.Wrapf("amount: empty string is not allowed")
	}

	if _, err := math.NewPositiveDecFromString(m.Amount); err != nil {
		return errors.ErrParse.Wrapf("amount: %s", err)
	}

	if m.Expiration == nil {
		return errors.ErrParse.Wrap("expiration: empty timestamp is not allowed")
	}

	return nil
}
