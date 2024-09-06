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
		&MsgCreateContent{},
		&MsgRemoveContent{},
		&MsgUpdateContentCurator{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// RegisterLegacyAminoCodec registers legacy amino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgCreateContent{}, "content/v1/MsgCreateContent")
	legacy.RegisterAminoMsg(cdc, &MsgRemoveContent{}, "content/v1/MsgRemoveContent")
	legacy.RegisterAminoMsg(cdc, &MsgUpdateContentCurator{}, "content/v1/MsgUpdateContentCurator")
}
