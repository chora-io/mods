package server

import (
	"context"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
	tmdb "github.com/tendermint/tm-db"

	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/choraio/mods/validator"
)

type baseSuite struct {
	t         gocuke.TestingT
	ctx       context.Context
	sdkCtx    sdk.Context
	authority sdk.AccAddress
	srv       Server
}

func setupBase(t gocuke.TestingT) *baseSuite {
	s := &baseSuite{t: t}

	// create in-memory database
	mdb := tmdb.NewMemDB()

	// create commit multi-store
	cms := store.NewCommitMultiStore(mdb)

	// create store key
	key := sdk.NewKVStoreKey(validator.ModuleName)

	// mount store with key and in-memory database
	cms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, mdb)

	// load commit multi-store
	require.NoError(t, cms.LoadLatestVersion())

	// create in-memory orm backend and set orm context
	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())

	// create and set sdk context from commit multi-store with orm context
	s.sdkCtx = sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger()).WithContext(ormCtx)

	// create and set context with sdk context
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	var err error

	// authority test account
	s.authority, err = sdk.AccAddressFromBech32("chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38")
	require.NoError(t, err)

	// create and set server
	s.srv = NewServer(key, s.authority)

	return s
}
