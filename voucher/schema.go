package voucher

import (
	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"

	voucherv1 "github.com/chora-io/mods/voucher/api/v1"
)

// ModuleSchema is the schema of the module.
var ModuleSchema = ormv1alpha1.ModuleSchemaDescriptor{
	SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: voucherv1.File_chora_voucher_v1_state_proto.Path()},
	},
}
