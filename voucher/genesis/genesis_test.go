package genesis

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"

	"github.com/choraio/mods/voucher"
	voucherv1 "github.com/choraio/mods/voucher/api/v1"
)

func TestValidateGenesis(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name   string
		setup  func(ctx context.Context, ss voucherv1.StateStore)
		errMsg string
	}{
		{
			name: "valid",
			setup: func(ctx context.Context, ss voucherv1.StateStore) {
				require.NoError(t, ss.VoucherTable().Insert(ctx, &voucherv1.Voucher{
					Issuer:   []byte("BTZfSbi0JKqguZ/tIAPUIhdAa7Y="),
					Metadata: "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf",
				}))
				require.NoError(t, ss.BalanceTable().Insert(ctx, &voucherv1.Balance{
					Id:         1,
					Address:    []byte("BTZfSbi0JKqguZ/tIAPUIhdAa7Y="),
					Amount:     "1.25",
					Expiration: timestamppb.Now(),
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

func setup(t *testing.T, setup func(ctx context.Context, ss voucherv1.StateStore)) json.RawMessage {
	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())

	db, err := ormdb.NewModuleDB(&voucher.ModuleSchema, ormdb.ModuleDBOptions{})
	require.NoError(t, err)

	ss, err := voucherv1.NewStateStore(db)
	require.NoError(t, err)

	setup(ormCtx, ss)

	target := ormjson.NewRawMessageTarget()
	require.NoError(t, db.ExportJSON(ormCtx, target))

	bz, err := target.JSON()
	require.NoError(t, err)

	return bz
}
