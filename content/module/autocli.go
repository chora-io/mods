package module

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	contentv1 "github.com/chora-io/mods/content/api/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: contentv1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Content",
					Use:       "content [id]",
					Short:     "query content by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "id"},
					},
				},
				{
					RpcMethod: "Contents",
					Use:       "contents",
					Short:     "query all content",
				},
				{
					RpcMethod: "ContentsByCurator",
					Use:       "contents-by-curator [curator]",
					Short:     "query content by curator",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "curator"},
					},
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              contentv1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: false, // use custom commands until v0.51
			RpcCommandOptions:    []*autocliv1.RpcCommandOptions{
				//{
				//	RpcMethod: "Create",
				//	Use:       "create [metadata]",
				//	Short:     "submit transaction to create content",
				//	PositionalArgs: []*autocliv1.PositionalArgDescriptor{
				//		{ProtoField: "metadata"},
				//	},
				//},
				//{
				//	RpcMethod: "Delete",
				//	Use:       "delete [id]",
				//	Short:     "submit transaction to delete content",
				//	PositionalArgs: []*autocliv1.PositionalArgDescriptor{
				//		{ProtoField: "id"},
				//	},
				//},
				//{
				//	RpcMethod: "UpdateCurator",
				//	Use:       "update-curator [id] [new-curator]",
				//	Short:     "submit transaction to update content curator",
				//	PositionalArgs: []*autocliv1.PositionalArgDescriptor{
				//		{ProtoField: "id"},
				//		{ProtoField: "new_curator"},
				//	},
				//},
				//{
				//	RpcMethod: "UpdateContentMetadata",
				//	Use:       "update-content-metadata [id] [new-metadata]",
				//	Short:     "submit transaction to update content metadata",
				//	PositionalArgs: []*autocliv1.PositionalArgDescriptor{
				//		{ProtoField: "id"},
				//		{ProtoField: "new_metadata"},
				//	},
				//},
			},
		},
	}
}
