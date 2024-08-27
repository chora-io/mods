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
		&MsgCreateAgent{},
		&MsgRemoveAgent{},
		&MsgUpdateAgentAdmin{},
		&MsgUpdateAgentMetadata{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// RegisterLegacyAminoCodec registers legacy amino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateAgent{}, "agent/MsgCreateAgent", nil)
	cdc.RegisterConcrete(&MsgRemoveAgent{}, "agent/MsgRemoveAgent", nil)
	cdc.RegisterConcrete(&MsgUpdateAgentAdmin{}, "agent/MsgUpdateAgentAdmin", nil)
	cdc.RegisterConcrete(&MsgUpdateAgentMetadata{}, "agent/MsgUpdateAgentMetadata", nil)
}
