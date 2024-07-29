package module

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	authorityv1 "github.com/chora-io/mods/authority/api/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: authorityv1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod:      "Authority",
					Use:            "authority",
					Short:          "query authority",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              authorityv1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: false, // use custom commands until v0.51
			RpcCommandOptions:    []*autocliv1.RpcCommandOptions{
				//{
				//	RpcMethod: "Update",
				//	Use:       "update [new-authority]",
				//	Short:     "submit a transaction to update the authority account",
				//	PositionalArgs: []*autocliv1.PositionalArgDescriptor{
				//		{ProtoField: "id"},
				//		{ProtoField: "new_authority"},
				//	},
				//},
			},
		},
	}
}
