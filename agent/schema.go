package agent

import (
	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"

	agentv1 "github.com/chora-io/mods/agent/api/v1"
)

// ModuleSchema is the schema of the module.
var ModuleSchema = ormv1alpha1.ModuleSchemaDescriptor{
	SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: agentv1.File_chora_agent_v1_state_proto.Path()},
	},
}
