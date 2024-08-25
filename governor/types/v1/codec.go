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
		&MsgAddValidator{},
		&MsgRemoveValidator{},
		&MsgUpdatePolicy{},
		&MsgUpdateValidator{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// RegisterLegacyAminoCodec registers legacy amino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgAddValidator{}, "validator/MsgAddValidator", nil)
	cdc.RegisterConcrete(&MsgRemoveValidator{}, "validator/MsgRemoveValidator", nil)
	cdc.RegisterConcrete(&MsgUpdatePolicy{}, "validator/MsgUpdatePolicy", nil)
	cdc.RegisterConcrete(&MsgUpdateValidator{}, "validator/MsgUpdateValidator", nil)
}
