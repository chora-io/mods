package server

import (
	"encoding/json"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/types/v2/ormstore"

	"github.com/choraio/mods/example"
	api "github.com/choraio/mods/example/api/v1"
	types "github.com/choraio/mods/example/types/v1"
)

var (
	_ types.MsgServer   = &Server{}
	_ types.QueryServer = &Server{}
)

// Server is the server.
type Server struct {
	db ormdb.ModuleDB
	ss api.StateStore
}

// NewServer creates a new server.
func NewServer(key storetypes.StoreKey) Server {
	s := Server{}

	var err error
	s.db, err = ormstore.NewStoreKeyDB(&example.ModuleSchema, key, ormdb.ModuleDBOptions{})
	if err != nil {
		panic(err)
	}

	s.ss, err = api.NewStateStore(s.db)
	if err != nil {
		panic(err)
	}

	return s
}

// RegisterInvariants registers the invariants.
func (s Server) RegisterInvariants(_ sdk.InvariantRegistry) {
	return
}

// InitGenesis creates genesis state.
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

// ValidateGenesis validates genesis state.
func (s Server) ValidateGenesis(_ json.RawMessage) error {
	return nil
}
