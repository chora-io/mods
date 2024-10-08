package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateValidator{}

// ValidateBasic performs stateless validation on MsgCreateValidator.
func (m MsgCreateValidator) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Operator); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("operator: %s", err)
	}

	if m.Metadata == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("metadata: empty string is not allowed")
	}

	if len(m.Metadata) > MetadataMaxLength {
		return sdkerrors.ErrInvalidRequest.Wrapf("metadata: exceeds max length %d", MetadataMaxLength)
	}

	return nil
}

// GetSigners returns the expected signers for MsgCreateValidator.
func (m MsgCreateValidator) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgCreateValidator) Route() string {
	return sdk.MsgTypeURL(&m)
}

// Type implements the LegacyMsg interface.
func (m MsgCreateValidator) Type() string {
	return sdk.MsgTypeURL(&m)
}
