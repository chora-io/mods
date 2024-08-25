package module

import (
	"context"
	"encoding/json"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"cosmossdk.io/core/appmodule"
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/chora-io/mods/governor"
	"github.com/chora-io/mods/governor/cmd"
	"github.com/chora-io/mods/governor/genesis"
	"github.com/chora-io/mods/governor/keeper"
	v1 "github.com/chora-io/mods/governor/types/v1"
)

// ConsensusVersion is the module consensus version.
const ConsensusVersion = 1

var (
	_ appmodule.AppModule = AppModule{}

	_ module.AppModuleBasic      = AppModule{}
	_ module.HasConsensusVersion = AppModule{}
	_ module.HasGenesis          = AppModule{}
	_ module.HasInvariants       = AppModule{}
	_ module.HasServices         = AppModule{}
)

// AppModule implements the AppModule interface.
type AppModule struct {
	cdc codec.BinaryCodec
	k   keeper.Keeper
}

// NewAppModule returns a new AppModule instance.
func NewAppModule(cdc codec.BinaryCodec, k keeper.Keeper) AppModule {
	return AppModule{
		cdc: cdc,
		k:   k,
	}
}

// IsAppModule implements the appmodule.AppModule/IsAppModule.
func (am AppModule) IsAppModule() {}

// IsOnePerModuleType implements appmodule.AppModule/IsOnePerModuleType.
func (am AppModule) IsOnePerModuleType() {}

// Name implements module.AppModuleBasic/Name.
func (am AppModule) Name() string {
	return governor.ModuleName
}

// RegisterGRPCGatewayRoutes implements module.AppModuleBasic/RegisterGRPCGatewayRoutes.
func (am AppModule) RegisterGRPCGatewayRoutes(clientCtx sdkclient.Context, mux *runtime.ServeMux) {
	err := v1.RegisterQueryHandlerClient(context.Background(), mux, v1.NewQueryClient(clientCtx))
	if err != nil {
		panic(err)
	}
}

// RegisterInterfaces implements module.AppModuleBasic/RegisterInterfaces.
func (am AppModule) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	v1.RegisterInterfaces(registry)
}

// RegisterLegacyAminoCodec implements module.AppModuleBasic/RegisterLegacyAminoCodec.
func (am AppModule) RegisterLegacyAminoCodec(amino *codec.LegacyAmino) {
	v1.RegisterLegacyAminoCodec(amino)
}

// ConsensusVersion implements module.HasConsensusVersion/ConsensusVersion.
func (am AppModule) ConsensusVersion() uint64 {
	return ConsensusVersion
}

// InitGenesis implements module.HasGenesis/InitGenesis.
func (am AppModule) InitGenesis(ctx sdk.Context, jsonCodec codec.JSONCodec, message json.RawMessage) {
	err := am.k.InitGenesis(ctx, jsonCodec, message)
	if err != nil {
		panic(err)
	}
}

// ExportGenesis implements module.HasGenesis/ExportGenesis.
func (am AppModule) ExportGenesis(ctx sdk.Context, jsonCodec codec.JSONCodec) json.RawMessage {
	export, err := am.k.ExportGenesis(ctx, jsonCodec)
	if err != nil {
		panic(err)
	}
	return export
}

// DefaultGenesis implements module.HasGenesis/DefaultGenesis.
func (am AppModule) DefaultGenesis(_ codec.JSONCodec) json.RawMessage {
	return nil
}

// ValidateGenesis implements module.HasGenesis/ValidateGenesis.
func (am AppModule) ValidateGenesis(_ codec.JSONCodec, _ sdkclient.TxEncodingConfig, bz json.RawMessage) error {
	return genesis.ValidateGenesis(bz)
}

// RegisterInvariants implements module.HasInvariants/RegisterInvariants.
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// RegisterServices implements module.HasServices/RegisterServices.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	v1.RegisterMsgServer(cfg.MsgServer(), am.k)
	v1.RegisterQueryServer(cfg.QueryServer(), am.k)
}

// GetTxCmd returns the transaction commands for the module
func (am AppModule) GetTxCmd() *cobra.Command {
	return cmd.TxCmd()
}
