package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

var _ legacytx.LegacyMsg = &MsgRemove{}

// ValidateBasic performs stateless validation on MsgRemove.
func (m MsgRemove) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("authority: %s", err)
	}

	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("address: %s", err)
	}

	return nil
}

// GetSigners returns the expected signers for MsgRemove.
func (m MsgRemove) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

// GetSignBytes implements the LegacyMsg interface.
func (m MsgRemove) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&m))
}

// Route implements the LegacyMsg interface.
func (m MsgRemove) Route() string {
	return sdk.MsgTypeURL(&m)
}

// Type implements the LegacyMsg interface.
func (m MsgRemove) Type() string {
	return sdk.MsgTypeURL(&m)
}
