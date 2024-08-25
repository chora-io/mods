package keeper

import (
	"encoding/json"

	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/chora-io/mods/ecosystem"
	ecosystemv1 "github.com/chora-io/mods/ecosystem/api/v1"
	v1 "github.com/chora-io/mods/ecosystem/types/v1"
)

var (
	_ v1.MsgServer   = &Keeper{}
	_ v1.QueryServer = &Keeper{}
)

// Keeper is the keeper.
type Keeper struct {
	db ormdb.ModuleDB         // module database
	ss ecosystemv1.StateStore // module state store
}

// NewKeeper creates a new keeper.
func NewKeeper(storeService store.KVStoreService) Keeper {
	k := Keeper{}

	var err error
	k.db, err = ormdb.NewModuleDB(&ecosystem.ModuleSchema, ormdb.ModuleDBOptions{
		KVStoreService: storeService,
	})
	if err != nil {
		panic(err)
	}

	k.ss, err = ecosystemv1.NewStateStore(k.db)
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
	//err = s.db.ImportJSON(sdk.WrapSDKContext(ctx), source)
	//if err != nil {
	//	return err
	//}

	return nil
}

// ExportGenesis exports genesis state.
func (k Keeper) ExportGenesis(ctx sdk.Context, _ codec.JSONCodec) (json.RawMessage, error) {
	//target := ormjson.NewRawMessageTarget()
	//
	//err := s.db.ExportJSON(sdk.WrapSDKContext(ctx), target)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return target.JSON()

	return nil, nil
}
