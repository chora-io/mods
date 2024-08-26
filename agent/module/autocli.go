package module

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	agentv1 "github.com/chora-io/mods/agent/api/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: agentv1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Agent",
					Use:       "agent [address]",
					Short:     "query agent by address",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "Agents",
					Use:       "agents",
					Short:     "query all agents",
				},
				{
					RpcMethod: "AgentsByAdmin",
					Use:       "agents-by-admin [admin]",
					Short:     "query agents by admin",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "admin"},
					},
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              agentv1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: false, // use custom commands until v0.51
			RpcCommandOptions:    []*autocliv1.RpcCommandOptions{
				//{
				//	RpcMethod: "Create",
				//	Use:       "create [metadata]",
				//	Short:     "submit a transaction to create agent",
				//	PositionalArgs: []*autocliv1.PositionalArgDescriptor{
				//		{ProtoField: "metadata"},
				//	},
				//},
				//{
				//	RpcMethod: "UpdateAgentAdmin",
				//	Use:       "update-agent-admin [address] [new-admin]",
				//	Short:     "submit a transaction to update agent admin",
				//	PositionalArgs: []*autocliv1.PositionalArgDescriptor{
				//		{ProtoField: "address"},
				//		{ProtoField: "new_admin"},
				//	},
				//},
				//{
				//	RpcMethod: "UpdateAgentMetadata",
				//	Use:       "update-metadata [address] [new-metadata]",
				//	Short:     "submit a transaction to update agent metadata",
				//	PositionalArgs: []*autocliv1.PositionalArgDescriptor{
				//		{ProtoField: "address"},
				//		{ProtoField: "new_metadata"},
				//	},
				//},
			},
		},
	}
}
