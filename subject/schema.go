package subject

import (
	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"

	subjectv1 "github.com/chora-io/mods/subject/api/v1"
)

// ModuleSchema is the schema of the module.
var ModuleSchema = ormv1alpha1.ModuleSchemaDescriptor{
	SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: subjectv1.File_chora_subject_v1_state_proto.Path()},
	},
}
