package module

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	validatorv1 "github.com/chora-io/mods/validator/api/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service:              validatorv1.Query_ServiceDesc.ServiceName,
			EnhanceCustomCommand: false, // use custom commands until v0.51
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "MaxMissedBlocks",
					Use:       "max-missed-blocks",
					Short:     "query the maximum number of missed blocks",
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
				{
					RpcMethod: "AddValidator",
					Use:       "add-validator [address] [metadata]",
					Short:     "submit a transaction to add a validator",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "address"},
						{ProtoField: "metadata"},
					},
				},
				{
					RpcMethod: "RemoveValidator",
					Use:       "remove-validator [address]",
					Short:     "submit a transaction to remove a validator",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "UpdateMetadata",
					Use:       "update-metadata [address] [metadata]",
					Short:     "submit a transaction to update validator metadata",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "address"},
						{ProtoField: "metadata"},
					},
				},
			},
		},
	}
}
