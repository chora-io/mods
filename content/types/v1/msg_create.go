package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreate{}

// ValidateBasic performs stateless validation on MsgCreate.
func (m MsgCreate) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Curator); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("curator: %s", err)
	}

	if m.Metadata == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("metadata: empty string is not allowed")
	}

	if len(m.Metadata) > MetadataMaxLength {
		return sdkerrors.ErrInvalidRequest.Wrapf("metadata: exceeds max length %d", MetadataMaxLength)
	}

	return nil
}

// GetSigners returns the expected signers for MsgCreate.
func (m MsgCreate) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Curator)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgCreate) Route() string {
	return sdk.MsgTypeURL(&m)
}

// Type implements the LegacyMsg interface.
func (m MsgCreate) Type() string {
	return sdk.MsgTypeURL(&m)
}
