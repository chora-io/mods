package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

var _ legacytx.LegacyMsg = &MsgUpdateCurator{}

// ValidateBasic performs stateless validation on MsgUpdateCurator.
func (m MsgUpdateCurator) ValidateBasic() error {
	if m.Id == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("id: empty or zero is not allowed")
	}

	if _, err := sdk.AccAddressFromBech32(m.Curator); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("curator: %s", err)
	}

	if _, err := sdk.AccAddressFromBech32(m.NewCurator); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("new curator: %s", err)
	}

	return nil
}

// GetSigners returns the expected signers for MsgUpdateCurator.
func (m MsgUpdateCurator) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Curator)
	return []sdk.AccAddress{addr}
}

// GetSignBytes implements the LegacyMsg interface.
func (m MsgUpdateCurator) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&m))
}

// Route implements the LegacyMsg interface.
func (m MsgUpdateCurator) Route() string {
	return sdk.MsgTypeURL(&m)
}

// Type implements the LegacyMsg interface.
func (m MsgUpdateCurator) Type() string {
	return sdk.MsgTypeURL(&m)
}
