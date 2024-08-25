package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateValidator{}

// ValidateBasic performs stateless validation on MsgUpdateValidator.
func (m MsgUpdateValidator) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("address: %s", err)
	}

	if m.NewMetadata == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("new metadata: empty string is not allowed")
	}

	if len(m.NewMetadata) > MetadataMaxLength {
		return sdkerrors.ErrInvalidRequest.Wrapf("new metadata: exceeds max length %d", MetadataMaxLength)
	}

	return nil
}

// GetSigners returns the expected signers for MsgUpdateValidator.
func (m MsgUpdateValidator) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Address)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgUpdateValidator) Route() string {
	return sdk.MsgTypeURL(&m)
}

// Type implements the LegacyMsg interface.
func (m MsgUpdateValidator) Type() string {
	return sdk.MsgTypeURL(&m)
}
