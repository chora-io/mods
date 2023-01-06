package server

import (
	"encoding/json"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/choraio/mods/geonode"
	geonodev1 "github.com/choraio/mods/geonode/api/v1"
	v1 "github.com/choraio/mods/geonode/types/v1"
	"github.com/choraio/mods/geonode/utils"
)

var (
	_ v1.MsgServer   = &Server{}
	_ v1.QueryServer = &Server{}
)

// Server is the server.
type Server struct {
	db ormdb.ModuleDB
	ss geonodev1.StateStore
}

// NewServer creates a new server.
func NewServer(key storetypes.StoreKey) Server {
	s := Server{}

	var err error
	s.db, err = utils.NewStoreKeyDB(&geonode.ModuleSchema, key, ormdb.ModuleDBOptions{})
	if err != nil {
		panic(err)
	}

	s.ss, err = geonodev1.NewStateStore(s.db)
	if err != nil {
		panic(err)
	}

	return s
}

// InitGenesis initializes genesis state.
func (s Server) InitGenesis(ctx sdk.Context, _ codec.JSONCodec, data json.RawMessage) ([]abci.ValidatorUpdate, error) {
	source, err := ormjson.NewRawMessageSource(data)
	if err != nil {
		return nil, err
	}

	err = s.db.ImportJSON(sdk.WrapSDKContext(ctx), source)
	if err != nil {
		return nil, err
	}

	return []abci.ValidatorUpdate{}, nil
}

// ExportGenesis exports genesis state.
func (s Server) ExportGenesis(ctx sdk.Context, _ codec.JSONCodec) (json.RawMessage, error) {
	target := ormjson.NewRawMessageTarget()

	err := s.db.ExportJSON(sdk.WrapSDKContext(ctx), target)
	if err != nil {
		return nil, err
	}

	return target.JSON()
}
