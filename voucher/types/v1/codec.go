package v1

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterInterfaces registers interfaces with the interface registry.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreate{},
		&MsgIssue{},
		&MsgUpdateIssuer{},
		&MsgUpdateMetadata{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// RegisterLegacyAminoCodec registers legacy amino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreate{}, "voucher/MsgCreate", nil)
	cdc.RegisterConcrete(&MsgIssue{}, "voucher/MsgIssue", nil)
	cdc.RegisterConcrete(&MsgUpdateIssuer{}, "voucher/MsgUpdateIssuer", nil)
	cdc.RegisterConcrete(&MsgUpdateMetadata{}, "voucher/MsgUpdateMetadata", nil)
}
