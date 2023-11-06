package keeper

import (
	"encoding/json"

	"cosmossdk.io/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/chora-io/mods/content"
	contentv1 "github.com/chora-io/mods/content/api/v1"
	v1 "github.com/chora-io/mods/content/types/v1"
)

var (
	_ v1.MsgServer   = &Keeper{}
	_ v1.QueryServer = &Keeper{}
)

// Keeper is the keeper.
type Keeper struct {
	db ormdb.ModuleDB       // module database
	ss contentv1.StateStore // module state store
}

// NewKeeper creates a new keeper.
func NewKeeper() Keeper {
	k := Keeper{}

	var err error
	k.db, err = ormdb.NewModuleDB(&content.ModuleSchema, ormdb.ModuleDBOptions{})
	if err != nil {
		panic(err)
	}

	k.ss, err = contentv1.NewStateStore(k.db)
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
