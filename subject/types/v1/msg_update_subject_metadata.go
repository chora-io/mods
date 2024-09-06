package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateSubjectMetadata{}

// ValidateBasic performs stateless validation on MsgUpdateSubjectMetadata.
func (m MsgUpdateSubjectMetadata) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("address: %s", err)
	}

	if _, err := sdk.AccAddressFromBech32(m.Steward); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("steward: %s", err)
	}

	if m.NewMetadata == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("new metadata: empty string is not allowed")
	}

	if len(m.NewMetadata) > MetadataMaxLength {
		return sdkerrors.ErrInvalidRequest.Wrapf("new metadata: exceeds max length %d", MetadataMaxLength)
	}

	return nil
}

// GetSigners returns the expected signers for MsgUpdateSubjectMetadata.
func (m MsgUpdateSubjectMetadata) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Steward)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgUpdateSubjectMetadata) Route() string {
	return sdk.MsgTypeURL(&m)
}

// Type implements the LegacyMsg interface.
func (m MsgUpdateSubjectMetadata) Type() string {
	return sdk.MsgTypeURL(&m)
}
