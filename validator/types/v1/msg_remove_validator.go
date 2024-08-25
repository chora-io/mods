package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRemoveValidator{}

// ValidateBasic performs stateless validation on MsgRemoveValidator.
func (m MsgRemoveValidator) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Admin); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("admin: %s", err)
	}

	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("address: %s", err)
	}

	return nil
}

// GetSigners returns the expected signers for MsgRemoveValidator.
func (m MsgRemoveValidator) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Admin)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgRemoveValidator) Route() string {
	return sdk.MsgTypeURL(&m)
}

// Type implements the LegacyMsg interface.
func (m MsgRemoveValidator) Type() string {
	return sdk.MsgTypeURL(&m)
}
