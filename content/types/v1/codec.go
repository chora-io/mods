package v1

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterInterfaces registers interfaces with the interface registry.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreate{},
		&MsgDelete{},
		&MsgUpdateCurator{},
		&MsgUpdateMetadata{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// RegisterLegacyAminoCodec registers legacy amino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgCreate{}, "content/v1/MsgCreate")
	legacy.RegisterAminoMsg(cdc, &MsgDelete{}, "content/v1/MsgDelete")
	legacy.RegisterAminoMsg(cdc, &MsgUpdateCurator{}, "content/v1/MsgUpdateCurator")
	legacy.RegisterAminoMsg(cdc, &MsgUpdateMetadata{}, "content/v1/MsgUpdateMetadata")
}
