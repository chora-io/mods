package module

import (
	"context"
	"encoding/json"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"cosmossdk.io/core/appmodule"
	storetypes "cosmossdk.io/store/types"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/choraio/mods/validator"
	"github.com/choraio/mods/validator/genesis"
	"github.com/choraio/mods/validator/server"
	v1 "github.com/choraio/mods/validator/types/v1"
)

// ConsensusVersion is the module consensus version.
const ConsensusVersion = 1

var (
	_ appmodule.AppModule   = AppModule{}
	_ module.AppModuleBasic = AppModule{}
	_ module.HasGenesis     = AppModule{}
	_ module.HasServices    = AppModule{}
)

// AppModule implements the AppModule interface.
type AppModule struct {
	key storetypes.StoreKey
	srv server.Server
}

// NewAppModule returns a new module.
func NewAppModule(key storetypes.StoreKey, authority sdk.AccAddress) AppModule {
	srv := server.NewServer(key, authority)
	return AppModule{
		key: key,
		srv: srv,
	}
}

// ConsensusVersion returns ConsensusVersion.
func (am AppModule) ConsensusVersion() uint64 {
	return ConsensusVersion
}

// IsAppModule implements the appmodule.AppModule/IsAppModule.
func (am AppModule) IsAppModule() {}

// IsOnePerModuleType implements appmodule.AppModule/IsOnePerModuleType.
func (am AppModule) IsOnePerModuleType() {}

// Name implements module.AppModuleBasic/Name.
func (am AppModule) Name() string {
	return validator.ModuleName
}

// RegisterInterfaces implements module.AppModuleBasic/RegisterInterfaces.
func (am AppModule) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	v1.RegisterInterfaces(registry)
}

// RegisterGRPCGatewayRoutes implements module.AppModuleBasic/RegisterGRPCGatewayRoutes.
func (am AppModule) RegisterGRPCGatewayRoutes(clientCtx sdkclient.Context, mux *runtime.ServeMux) {
	err := v1.RegisterQueryHandlerClient(context.Background(), mux, v1.NewQueryClient(clientCtx))
	if err != nil {
		panic(err)
	}
}

// RegisterLegacyAminoCodec implements module.AppModuleBasic/RegisterLegacyAminoCodec.
func (am AppModule) RegisterLegacyAminoCodec(amino *codec.LegacyAmino) {
	v1.RegisterLegacyAminoCodec(amino)
}

// InitGenesis implements module.HasGenesis/InitGenesis.
func (am AppModule) InitGenesis(ctx sdk.Context, jsonCodec codec.JSONCodec, message json.RawMessage) {
	err := am.srv.InitGenesis(ctx, jsonCodec, message)
	if err != nil {
		panic(err)
	}
}

// ExportGenesis implements module.HasGenesis/ExportGenesis.
func (am AppModule) ExportGenesis(ctx sdk.Context, jsonCodec codec.JSONCodec) json.RawMessage {
	export, err := am.srv.ExportGenesis(ctx, jsonCodec)
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

// RegisterServices implements module.HasServices/RegisterServices.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	v1.RegisterMsgServer(cfg.MsgServer(), am.srv)
	v1.RegisterQueryServer(cfg.QueryServer(), am.srv)
}
