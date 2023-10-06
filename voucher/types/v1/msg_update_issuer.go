package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateIssuer{}

// ValidateBasic performs stateless validation on MsgUpdateIssuer.
func (m MsgUpdateIssuer) ValidateBasic() error {
	if m.Id == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("id: empty or zero is not allowed")
	}

	if _, err := sdk.AccAddressFromBech32(m.Issuer); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("issuer: %s", err)
	}

	if _, err := sdk.AccAddressFromBech32(m.NewIssuer); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("new issuer: %s", err)
	}

	return nil
}

// GetSigners returns the expected signers for MsgUpdateIssuer.
func (m MsgUpdateIssuer) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Issuer)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgUpdateIssuer) Route() string {
	return sdk.MsgTypeURL(&m)
}

// Type implements the LegacyMsg interface.
func (m MsgUpdateIssuer) Type() string {
	return sdk.MsgTypeURL(&m)
}
