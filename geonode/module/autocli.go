package module

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	geonodev1 "github.com/chora-io/mods/geonode/api/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service:              geonodev1.Query_ServiceDesc.ServiceName,
			EnhanceCustomCommand: false, // use custom commands until v0.51
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Node",
					Use:       "node [id]",
					Short:     "query node by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "id"},
					},
				},
				{
					RpcMethod: "Nodes",
					Use:       "nodes",
					Short:     "query all nodes",
				},
				{
					RpcMethod: "NodesByCurator",
					Use:       "nodes-by-curator [curator]",
					Short:     "query nodes by curator",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "curator"},
					},
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              geonodev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: false, // use custom commands until v0.51
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Create",
					Use:       "create [metadata]",
					Short:     "submit a transaction to create node",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "metadata"},
					},
				},
				{
					RpcMethod: "UpdateCurator",
					Use:       "update-curator [id] [new-curator]",
					Short:     "submit a transaction to update node curator",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "id"},
						{ProtoField: "new_curator"},
					},
				},
				{
					RpcMethod: "UpdateMetadata",
					Use:       "update-metadata [id] [new-metadata]",
					Short:     "submit a transaction to update node metadata",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "id"},
						{ProtoField: "new_metadata"},
					},
				},
			},
		},
	}
}
