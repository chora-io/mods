package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRemoveGovernor{}

// ValidateBasic performs stateless validation on MsgRemoveGovernor.
func (m MsgRemoveGovernor) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("address: %s", err)
	}

	return nil
}

// GetSigners returns the expected signers for MsgRemoveGovernor.
func (m MsgRemoveGovernor) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Address)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgRemoveGovernor) Route() string {
	return sdk.MsgTypeURL(&m)
}

// Type implements the LegacyMsg interface.
func (m MsgRemoveGovernor) Type() string {
	return sdk.MsgTypeURL(&m)
}
