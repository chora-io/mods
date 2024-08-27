package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateGovernor{}

// ValidateBasic performs stateless validation on MsgCreateGovernor.
func (m MsgCreateGovernor) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("address: %s", err)
	}

	if m.Metadata == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("metadata: empty string is not allowed")
	}

	if len(m.Metadata) > MetadataMaxLength {
		return sdkerrors.ErrInvalidRequest.Wrapf("metadata: exceeds max length %d", MetadataMaxLength)
	}

	return nil
}

// GetSigners returns the expected signers for MsgCreateGovernor.
func (m MsgCreateGovernor) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Address)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgCreateGovernor) Route() string {
	return sdk.MsgTypeURL(&m)
}

// Type implements the LegacyMsg interface.
func (m MsgCreateGovernor) Type() string {
	return sdk.MsgTypeURL(&m)
}
