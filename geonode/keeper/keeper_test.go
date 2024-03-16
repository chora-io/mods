package keeper

import (
	db "github.com/cosmos/cosmos-db"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"cosmossdk.io/log"
	"cosmossdk.io/orm/model/ormtable"
	"cosmossdk.io/orm/testing/ormtest"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/chora-io/mods/geonode"
)

type baseSuite struct {
	t      gocuke.TestingT
	sdkCtx sdk.Context
	k      Keeper
}

func setupBase(t gocuke.TestingT) *baseSuite {
	s := &baseSuite{t: t}

	// create in-memory database
	mdb := db.NewMemDB()

	// create commit multi-store
	cms := store.NewCommitMultiStore(mdb, log.NewNopLogger(), metrics.NewNoOpMetrics())

	// create store key
	key := storetypes.NewKVStoreKey(geonode.ModuleName)

	// mount store with key and in-memory database
	cms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, mdb)

	// load commit multi-store
	require.NoError(t, cms.LoadLatestVersion())

	// create in-memory orm backend and set orm context
	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())

	// create and set sdk context from commit multi-store with orm context
	s.sdkCtx = sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger()).WithContext(ormCtx)

	// create store service
	service := runtime.NewKVStoreService(key)

	// create and set keeper
	s.k = NewKeeper(service)

	return s
}
