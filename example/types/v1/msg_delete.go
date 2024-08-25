package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDelete{}

// ValidateBasic performs stateless validation on MsgDelete.
func (m MsgDelete) ValidateBasic() error {
	if m.Id == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("id: empty or zero is not allowed")
	}

	if _, err := sdk.AccAddressFromBech32(m.Curator); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("curator: %s", err)
	}

	return nil
}

// GetSigners returns the expected signers for MsgDelete.
func (m MsgDelete) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Curator)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgDelete) Route() string {
	return sdk.MsgTypeURL(&m)
}

// Type implements the LegacyMsg interface.
func (m MsgDelete) Type() string {
	return sdk.MsgTypeURL(&m)
}
