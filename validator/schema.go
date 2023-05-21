package validator

import (
	ormv1alpha1 "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"

	validatorv1 "github.com/choraio/mods/validator/api/v1"
)

// ModuleSchema is the schema of the module.
var ModuleSchema = ormv1alpha1.ModuleSchemaDescriptor{
	SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: validatorv1.File_chora_validator_v1_state_proto.Path()},
	},
}
