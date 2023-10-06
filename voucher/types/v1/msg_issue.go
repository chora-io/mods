package v1

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgIssue{}

// ValidateBasic performs stateless validation on MsgIssue.
func (m MsgIssue) ValidateBasic() error {
	if m.Id == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("id: empty or zero is not allowed")
	}

	if _, err := sdk.AccAddressFromBech32(m.Issuer); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("issuer: %s", err)
	}

	if _, err := sdk.AccAddressFromBech32(m.Recipient); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("recipient: %s", err)
	}

	if m.Amount == "" {
		return sdkerrors.ErrInvalidRequest.Wrapf("amount: empty string is not allowed")
	}

	dec, err := math.LegacyNewDecFromStr(m.Amount)
	if err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("amount: %s", err)
	}
	if !dec.IsPositive() {
		return sdkerrors.ErrInvalidRequest.Wrapf("amount: expected a positive decimal, got %s: invalid decimal string", dec.String())
	}

	if m.Expiration == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("expiration: empty timestamp is not allowed")
	}

	if m.Metadata == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("metadata: empty string is not allowed")
	}

	if len(m.Metadata) > MetadataMaxLength {
		return sdkerrors.ErrInvalidRequest.Wrapf("metadata: exceeds max length %d", MetadataMaxLength)
	}

	return nil
}

// GetSigners returns the expected signers for MsgIssue.
func (m MsgIssue) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Issuer)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgIssue) Route() string {
	return sdk.MsgTypeURL(&m)
}

// Type implements the LegacyMsg interface.
func (m MsgIssue) Type() string {
	return sdk.MsgTypeURL(&m)
}
