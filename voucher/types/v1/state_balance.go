package v1

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/chora-io/mods/voucher/errors"
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

	dec, err := math.LegacyNewDecFromStr(m.Amount)
	if err != nil {
		return errors.ErrParse.Wrapf("amount: %s", err)
	}
	if !dec.IsPositive() {
		return errors.ErrParse.Wrapf("amount: expected a positive decimal, got %s: invalid decimal string", dec.String())
	}

	if m.Expiration == nil {
		return errors.ErrParse.Wrap("expiration: empty timestamp is not allowed")
	}

	return nil
}
