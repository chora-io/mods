package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateSubjectSteward{}

// ValidateBasic performs stateless validation on MsgUpdateSubjectSteward.
func (m MsgUpdateSubjectSteward) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("address: %s", err)
	}

	if _, err := sdk.AccAddressFromBech32(m.Steward); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("steward: %s", err)
	}

	if _, err := sdk.AccAddressFromBech32(m.NewSteward); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("new steward: %s", err)
	}

	return nil
}

// GetSigners returns the expected signers for MsgUpdateSubjectSteward.
func (m MsgUpdateSubjectSteward) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Steward)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgUpdateSubjectSteward) Route() string {
	return sdk.MsgTypeURL(&m)
}

// Type implements the LegacyMsg interface.
func (m MsgUpdateSubjectSteward) Type() string {
	return sdk.MsgTypeURL(&m)
}
