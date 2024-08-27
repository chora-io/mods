package module

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	governorv1 "github.com/chora-io/mods/governor/api/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: governorv1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Policy",
					Use:       "policy",
					Short:     "query policy",
				},
				{
					RpcMethod: "Governor",
					Use:       "governor [address]",
					Short:     "query governor by address",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "Governors",
					Use:       "governors",
					Short:     "query all governors",
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              governorv1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: false, // use custom commands until v0.51
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				//{
				//	RpcMethod: "CreateGovernor",
				//	Use:       "create-governor [address] [metadata]",
				//	Short:     "submit a transaction to add a governor",
				//	PositionalArgs: []*autocliv1.PositionalArgDescriptor{
				//		{ProtoField: "address"},
				//		{ProtoField: "metadata"},
				//	},
				//},
				//{
				//	RpcMethod: "RemoveGovernor",
				//	Use:       "remove-governor [address]",
				//	Short:     "submit a transaction to remove a governor",
				//	PositionalArgs: []*autocliv1.PositionalArgDescriptor{
				//		{ProtoField: "address"},
				//	},
				//},
				{
					RpcMethod: "UpdatePolicy",
					Use:       "update-policy [admin] [signed-blocks-window] [min-signed-per-window]",
					Short:     "submit a transaction to update a governor",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "admin"},
						{ProtoField: "signed_blocks_window"},
						{ProtoField: "min_signed_per_window"},
					},
				},
				//{
				//	RpcMethod: "UpdateGovernor",
				//	Use:       "update-governor [address] [metadata]",
				//	Short:     "submit a transaction to update a governor",
				//	PositionalArgs: []*autocliv1.PositionalArgDescriptor{
				//		{ProtoField: "address"},
				//		{ProtoField: "metadata"},
				//	},
				//},
			},
		},
	}
}
