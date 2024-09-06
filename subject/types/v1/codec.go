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
		&MsgCreateSubject{},
		&MsgRemoveSubject{},
		&MsgUpdateSubjectMetadata{},
		&MsgUpdateSubjectSteward{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// RegisterLegacyAminoCodec registers legacy amino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateSubject{}, "subject/MsgCreateSubject", nil)
	cdc.RegisterConcrete(&MsgRemoveSubject{}, "subject/MsgRemoveSubject", nil)
	cdc.RegisterConcrete(&MsgUpdateSubjectMetadata{}, "subject/MsgUpdateSubjectMetadata", nil)
	cdc.RegisterConcrete(&MsgUpdateSubjectSteward{}, "subject/MsgUpdateSubjectSteward", nil)
}
