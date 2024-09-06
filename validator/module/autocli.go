package module

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	validatorv1 "github.com/chora-io/mods/validator/api/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: validatorv1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Policy",
					Use:       "policy",
					Short:     "query policy",
				},
				{
					RpcMethod: "Validator",
					Use:       "validator [address]",
					Short:     "query validator by address",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "Validators",
					Use:       "validators",
					Short:     "query all validators",
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              validatorv1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: false, // use custom commands until v0.51
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				//{
				//	RpcMethod: "CreateValidator",
				//	Use:       "create-validator [metadata]",
				//	Short:     "submit transaction to create validator",
				//	PositionalArgs: []*autocliv1.PositionalArgDescriptor{
				//		{ProtoField: "metadata"},
				//	},
				//},
				//{
				//	RpcMethod: "RemoveValidator",
				//	Use:       "remove-validator [address]",
				//	Short:     "submit transaction to remove a validator",
				//	PositionalArgs: []*autocliv1.PositionalArgDescriptor{
				//		{ProtoField: "address"},
				//	},
				//},
				{
					RpcMethod: "UpdatePolicy",
					Use:       "update-policy [signed-blocks-window] [min-signed-per-window]",
					Short:     "submit transaction to update a validator",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "signed_blocks_window"},
						{ProtoField: "min_signed_per_window"},
					},
				},
				//{
				//	RpcMethod: "UpdateValidator",
				//	Use:       "update-validator [address] [new-metadata]",
				//	Short:     "submit transaction to update validator",
				//	PositionalArgs: []*autocliv1.PositionalArgDescriptor{
				//		{ProtoField: "address"},
				//		{ProtoField: "new_metadata"},
				//	},
				//},
			},
		},
	}
}
