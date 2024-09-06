package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRemoveSubject{}

// ValidateBasic performs stateless validation on MsgRemoveSubject.
func (m MsgRemoveSubject) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("address: %s", err)
	}

	if _, err := sdk.AccAddressFromBech32(m.Steward); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("steward: %s", err)
	}

	return nil
}

// GetSigners returns the expected signers for MsgRemoveSubject.
func (m MsgRemoveSubject) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Steward)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgRemoveSubject) Route() string {
	return sdk.MsgTypeURL(&m)
}

// Type implements the LegacyMsg interface.
func (m MsgRemoveSubject) Type() string {
	return sdk.MsgTypeURL(&m)
}
