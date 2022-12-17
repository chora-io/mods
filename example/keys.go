package example

import (
	ormapi "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"

	api "github.com/choraio/mods/example/api/v1"
)

const (
	ModuleName = "example"
)

var ModuleSchema = ormapi.ModuleSchemaDescriptor{
	SchemaFile: []*ormapi.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: api.File_v1_state_proto.Path()},
	},
}
