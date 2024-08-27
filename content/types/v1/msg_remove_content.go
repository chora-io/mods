package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRemoveContent{}

// ValidateBasic performs stateless validation on MsgRemoveContent.
func (m MsgRemoveContent) ValidateBasic() error {
	if m.Id == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("id: empty or zero is not allowed")
	}

	if _, err := sdk.AccAddressFromBech32(m.Curator); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("curator: %s", err)
	}

	return nil
}

// GetSigners returns the expected signers for MsgRemoveContent.
func (m MsgRemoveContent) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Curator)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgRemoveContent) Route() string {
	return sdk.MsgTypeURL(&m)
}

// Type implements the LegacyMsg interface.
func (m MsgRemoveContent) Type() string {
	return sdk.MsgTypeURL(&m)
}
