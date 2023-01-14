package server

import (
	"encoding/json"
	"time"

	abci "github.com/tendermint/tendermint/abci/types"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/choraio/mods/voucher"
	voucherv1 "github.com/choraio/mods/voucher/api/v1"
	v1 "github.com/choraio/mods/voucher/types/v1"
	"github.com/choraio/mods/voucher/utils"
)

var (
	_ v1.MsgServer   = &Server{}
	_ v1.QueryServer = &Server{}
)

// Server is the server.
type Server struct {
	db ormdb.ModuleDB
	ss voucherv1.StateStore
}

// NewServer creates a new server.
func NewServer(key storetypes.StoreKey) Server {
	s := Server{}

	var err error
	s.db, err = utils.NewStoreKeyDB(&voucher.ModuleSchema, key, ormdb.ModuleDBOptions{})
	if err != nil {
		panic(err)
	}

	s.ss, err = voucherv1.NewStateStore(s.db)
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

// PruneVouchers removes expired vouchers from state.
func (s Server) PruneVouchers(ctx sdk.Context) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	minTime := timestamppb.New(time.Time{})
	blockTime := timestamppb.New(sdkCtx.BlockTime())

	fromKey := voucherv1.BalanceExpirationIndexKey{}.WithExpiration(minTime)
	toKey := voucherv1.BalanceExpirationIndexKey{}.WithExpiration(blockTime)

	it, err := s.ss.BalanceTable().ListRange(ctx, fromKey, toKey)
	if err != nil {
		return err
	}

	for it.Next() {
		balance, err := it.Value()
		if err != nil {
			return err
		}

		err = s.ss.BalanceTable().Delete(ctx, balance)
		if err != nil {
			return err
		}
	}

	it.Close()

	return nil
}
