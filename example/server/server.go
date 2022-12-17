package server

import (
	"github.com/choraio/mods/example"
	api "github.com/choraio/mods/example/api/v1"
	types "github.com/choraio/mods/example/types/v1"
	"github.com/choraio/mods/example/utils/ormstore"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

var (
	_ types.MsgServer   = &Server{}
	_ types.QueryServer = &Server{}
)

type Server struct {
	db ormdb.ModuleDB
	ss api.StateStore
}

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
