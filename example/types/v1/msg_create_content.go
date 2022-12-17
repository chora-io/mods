package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

var _ legacytx.LegacyMsg = &MsgCreateContent{}

// ValidateBasic performs stateless validation on MsgCreateContent.
func (m MsgCreateContent) ValidateBasic() error {
	if m.Creator == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("creator cannot be empty")
	}

	if _, err := sdk.AccAddressFromBech32(m.Creator); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("creator: %s", err)
	}

	if m.Hash == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("hash: cannot be empty")
	}

	if len(m.Hash) > HashMaxLength {
		return sdkerrors.ErrInvalidRequest.Wrapf("hash: exceeds max length %d", HashMaxLength)
	}

	return nil
}

// GetSigners returns the expected signers for MsgCreateContent.
func (m MsgCreateContent) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Creator)
	return []sdk.AccAddress{addr}
}

// GetSignBytes implements the LegacyMsg interface.
func (m MsgCreateContent) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&m))
}

// Route implements the LegacyMsg interface.
func (m MsgCreateContent) Route() string {
	return sdk.MsgTypeURL(&m)
}

// Type implements the LegacyMsg interface.
func (m MsgCreateContent) Type() string {
	return sdk.MsgTypeURL(&m)
}
