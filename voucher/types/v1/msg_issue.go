package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"

	"github.com/regen-network/regen-ledger/types/v2/math"
)

var _ legacytx.LegacyMsg = &MsgIssue{}

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

	if _, err := math.NewPositiveDecFromString(m.Amount); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("amount: %s", err)
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

// GetSignBytes implements the LegacyMsg interface.
func (m MsgIssue) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&m))
}

// Route implements the LegacyMsg interface.
func (m MsgIssue) Route() string {
	return sdk.MsgTypeURL(&m)
}

// Type implements the LegacyMsg interface.
func (m MsgIssue) Type() string {
	return sdk.MsgTypeURL(&m)
}
