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
		&MsgCreateGovernor{},
		&MsgRemoveGovernor{},
		&MsgUpdateGovernorMetadata{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// RegisterLegacyAminoCodec registers legacy amino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateGovernor{}, "governor/MsgCreateGovernor", nil)
	cdc.RegisterConcrete(&MsgRemoveGovernor{}, "governor/MsgRemoveGovernor", nil)
	cdc.RegisterConcrete(&MsgUpdateGovernorMetadata{}, "governor/MsgUpdateGovernorMetadata", nil)
}
