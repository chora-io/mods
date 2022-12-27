package geonode

import (
	ormv1alpha1 "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"

	geonodev1 "github.com/choraio/mods/geonode/api/v1"
)

// ModuleSchema is the schema of the module.
var ModuleSchema = ormv1alpha1.ModuleSchemaDescriptor{
	SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: geonodev1.File_v1_state_proto.Path()},
	},
}
