package genesis

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/orm/model/ormdb"
	"cosmossdk.io/orm/model/ormtable"
	"cosmossdk.io/orm/testing/ormtest"
	"github.com/choraio/mods/validator"
	validatorv1 "github.com/choraio/mods/validator/api/v1"
)

func TestValidateGenesis(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name   string
		setup  func(ctx context.Context, ss validatorv1.StateStore)
		errMsg string
	}{
		{
			name: "valid",
			setup: func(ctx context.Context, ss validatorv1.StateStore) {
				require.NoError(t, ss.ValidatorTable().Insert(ctx, &validatorv1.Validator{
					Address:  "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
					Metadata: "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf",
				}))
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			bz := setup(t, tc.setup)
			err := ValidateGenesis(bz)
			if tc.errMsg != "" {
				require.Error(t, err)
				require.ErrorContains(t, err, tc.errMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func setup(t *testing.T, setup func(ctx context.Context, ss validatorv1.StateStore)) json.RawMessage {
	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())

	db, err := ormdb.NewModuleDB(&validator.ModuleSchema, ormdb.ModuleDBOptions{})
	require.NoError(t, err)

	ss, err := validatorv1.NewStateStore(db)
	require.NoError(t, err)

	setup(ormCtx, ss)

	//target := ormjson.NewRawMessageTarget()
	//require.NoError(t, db.ExportJSON(ormCtx, target))
	//
	//bz, err := target.JSON()
	//require.NoError(t, err)
	//
	//return bz

	return nil
}
