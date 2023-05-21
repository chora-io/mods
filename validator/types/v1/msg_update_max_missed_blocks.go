package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

var _ legacytx.LegacyMsg = &MsgUpdateMaxMissedBlocks{}

// ValidateBasic performs stateless validation on MsgUpdateMaxMissedBlocks.
func (m MsgUpdateMaxMissedBlocks) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("authority: %s", err)
	}

	return nil
}

// GetSigners returns the expected signers for MsgUpdateMaxMissedBlocks.
func (m MsgUpdateMaxMissedBlocks) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

// GetSignBytes implements the LegacyMsg interface.
func (m MsgUpdateMaxMissedBlocks) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&m))
}

// Route implements the LegacyMsg interface.
func (m MsgUpdateMaxMissedBlocks) Route() string {
	return sdk.MsgTypeURL(&m)
}

// Type implements the LegacyMsg interface.
func (m MsgUpdateMaxMissedBlocks) Type() string {
	return sdk.MsgTypeURL(&m)
}
