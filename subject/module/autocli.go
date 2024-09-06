package module

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	subjectv1 "github.com/chora-io/mods/subject/api/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: subjectv1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Subject",
					Use:       "subject [address]",
					Short:     "query subject",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "address"},
					},
				},
				{
					RpcMethod: "Subjects",
					Use:       "subjects",
					Short:     "query subjects",
				},
				{
					RpcMethod: "SubjectsBySteward",
					Use:       "subjects-by-steward [steward]",
					Short:     "query subjects by steward",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "steward"},
					},
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              subjectv1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: false, // use custom commands until v0.51
			RpcCommandOptions:    []*autocliv1.RpcCommandOptions{
				//{
				//	RpcMethod: "CreateSubject",
				//	Use:       "create-subject [metadata]",
				//	Short:     "submit transaction to create subject",
				//	PositionalArgs: []*autocliv1.PositionalArgDescriptor{
				//		{ProtoField: "metadata"},
				//	},
				//},
				//{
				//	RpcMethod: "RemoveSubject",
				//	Use:       "remove-subject [address]",
				//	Short:     "submit transaction to remove subject",
				//	PositionalArgs: []*autocliv1.PositionalArgDescriptor{
				//		{ProtoField: "address"},
				//	},
				//},
				//{
				//	RpcMethod: "UpdateSubjectSteward",
				//	Use:       "update-subject-steward [address] [new-steward]",
				//	Short:     "submit transaction to update subject steward",
				//	PositionalArgs: []*autocliv1.PositionalArgDescriptor{
				//		{ProtoField: "address"},
				//		{ProtoField: "new_steward"},
				//	},
				//},
				//{
				//	RpcMethod: "UpdateSubjectMetadata",
				//	Use:       "update-subject-metadata [address] [new-metadata]",
				//	Short:     "submit transaction to update subject metadata",
				//	PositionalArgs: []*autocliv1.PositionalArgDescriptor{
				//		{ProtoField: "address"},
				//		{ProtoField: "new_metadata"},
				//	},
				//},
			},
		},
	}
}
