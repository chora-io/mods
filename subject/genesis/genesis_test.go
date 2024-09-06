package genesis

import (
	"context"
	"encoding/json"
	"testing"

	"cosmossdk.io/orm/model/ormdb"
	"cosmossdk.io/orm/model/ormtable"
	"cosmossdk.io/orm/testing/ormtest"
	"github.com/stretchr/testify/require"

	"github.com/chora-io/mods/subject"
	subjectv1 "github.com/chora-io/mods/subject/api/v1"
)

func TestValidateGenesis(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name   string
		setup  func(ctx context.Context, ss subjectv1.StateStore)
		errMsg string
	}{
		{
			name: "valid",
			setup: func(ctx context.Context, ss subjectv1.StateStore) {
				require.NoError(t, ss.SubjectTable().Insert(ctx, &subjectv1.Subject{
					Steward:  []byte("BTZfSbi0JKqguZ/tIAPUIhdAa7Y="),
					Address:  []byte("G+ksLYTNBuzyqdTij+Xkx1ztGDzOMACTUcjF6iEkiH0="),
					Metadata: "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf",
				}))
			},
		},
		{
			name: "valid",
			setup: func(ctx context.Context, ss subjectv1.StateStore) {
				require.NoError(t, ss.SubjectSequenceTable().Save(ctx, &subjectv1.SubjectSequence{
					Sequence: 0,
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

func setup(t *testing.T, setup func(ctx context.Context, ss subjectv1.StateStore)) json.RawMessage {
	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())

	db, err := ormdb.NewModuleDB(&subject.ModuleSchema, ormdb.ModuleDBOptions{})
	require.NoError(t, err)

	ss, err := subjectv1.NewStateStore(db)
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
