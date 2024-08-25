package governor

import (
	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"

	governorv1 "github.com/chora-io/mods/governor/api/v1"
)

// ModuleSchema is the schema of the module.
var ModuleSchema = ormv1alpha1.ModuleSchemaDescriptor{
	SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: governorv1.File_chora_governor_v1_state_proto.Path()},
	},
}
