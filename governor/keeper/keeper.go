package keeper

import (
	"encoding/json"

	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/chora-io/mods/governor"
	governorv1 "github.com/chora-io/mods/governor/api/v1"
	v1 "github.com/chora-io/mods/governor/types/v1"
)

var (
	_ v1.MsgServer   = &Keeper{}
	_ v1.QueryServer = &Keeper{}
)

// Keeper is the keeper.
type Keeper struct {
	admin sdk.AccAddress // admin account address

	db ormdb.ModuleDB        // module database
	ss governorv1.StateStore // module state store
}

// NewKeeper creates a new keeper.
func NewKeeper(storeService store.KVStoreService, admin sdk.AccAddress) Keeper {
	k := Keeper{admin: admin}

	var err error

	k.db, err = ormdb.NewModuleDB(&governor.ModuleSchema, ormdb.ModuleDBOptions{
		KVStoreService: storeService,
	})
	if err != nil {
		panic(err)
	}

	k.ss, err = governorv1.NewStateStore(k.db)
	if err != nil {
		panic(err)
	}

	return k
}

// InitGenesis initializes genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, _ codec.JSONCodec, data json.RawMessage) error {
	//source, err := ormjson.NewRawMessageSource(data)
	//if err != nil {
	//	return err
	//}
	//
	//err = k.db.ImportJSON(sdk.WrapSDKContext(ctx), source)
	//if err != nil {
	//	return err
	//}

	return nil
}

// ExportGenesis exports genesis state.
func (k Keeper) ExportGenesis(ctx sdk.Context, _ codec.JSONCodec) (json.RawMessage, error) {
	//target := ormjson.NewRawMessageTarget()
	//
	//err := k.db.ExportJSON(sdk.WrapSDKContext(ctx), target)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return target.JSON()

	return nil, nil
}
