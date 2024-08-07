package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdate{}

// ValidateBasic performs stateless validation on MsgUpdate.
func (m MsgUpdate) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("authority: %s", err)
	}

	if _, err := sdk.AccAddressFromBech32(m.NewAuthority); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("new authority: %s", err)
	}

	return nil
}

// GetSigners returns the expected signers for MsgUpdate.
func (m MsgUpdate) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgUpdate) Route() string {
	return sdk.MsgTypeURL(&m)
}

// Type implements the LegacyMsg interface.
func (m MsgUpdate) Type() string {
	return sdk.MsgTypeURL(&m)
}
