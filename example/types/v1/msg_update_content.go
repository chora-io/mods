package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

var _ legacytx.LegacyMsg = &MsgUpdateContent{}

// ValidateBasic performs stateless validation on MsgUpdateContent.
func (m MsgUpdateContent) ValidateBasic() error {
	if m.Id == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("id: empty or zero is not allowed")
	}

	if _, err := sdk.AccAddressFromBech32(m.Creator); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("creator: %s", err)
	}

	if m.NewHash == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("new hash: empty string is not allowed")
	}

	if len(m.NewHash) > HashMaxLength {
		return sdkerrors.ErrInvalidRequest.Wrapf("new hash: exceeds max length %d", HashMaxLength)
	}

	return nil
}

// GetSigners returns the expected signers for MsgUpdateContent.
func (m MsgUpdateContent) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Creator)
	return []sdk.AccAddress{addr}
}

// GetSignBytes implements the LegacyMsg interface.
func (m MsgUpdateContent) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&m))
}

// Route implements the LegacyMsg interface.
func (m MsgUpdateContent) Route() string {
	return sdk.MsgTypeURL(&m)
}

// Type implements the LegacyMsg interface.
func (m MsgUpdateContent) Type() string {
	return sdk.MsgTypeURL(&m)
}
