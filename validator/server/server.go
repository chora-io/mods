package server

import (
	"encoding/json"

	"cosmossdk.io/orm/model/ormdb"
	storetypes "cosmossdk.io/store/types"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/choraio/mods/validator"
	validatorv1 "github.com/choraio/mods/validator/api/v1"
	v1 "github.com/choraio/mods/validator/types/v1"
)

var (
	_ v1.MsgServer   = &Server{}
	_ v1.QueryServer = &Server{}
)

// Server is the server.
type Server struct {
	db        ormdb.ModuleDB
	ss        validatorv1.StateStore
	authority sdk.AccAddress
}

// NewServer creates a new server.
func NewServer(key storetypes.StoreKey, authority sdk.AccAddress) Server {
	s := Server{}

	var err error
	s.db, err = ormdb.NewModuleDB(&validator.ModuleSchema, ormdb.ModuleDBOptions{})
	if err != nil {
		panic(err)
	}

	s.ss, err = validatorv1.NewStateStore(s.db)
	if err != nil {
		panic(err)
	}

	s.authority = authority

	return s
}

// InitGenesis initializes genesis state.
func (s Server) InitGenesis(ctx sdk.Context, _ codec.JSONCodec, data json.RawMessage) error {
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
func (s Server) ExportGenesis(ctx sdk.Context, _ codec.JSONCodec) (json.RawMessage, error) {
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

// GetPolicy gets the validator signing policy.
func (s Server) GetPolicy(ctx sdk.Context) (*validatorv1.Policy, error) {
	policy, err := s.ss.PolicyTable().Get(ctx)
	if err != nil {
		return nil, err
	}
	return policy, nil
}

// HandleSigningInfo tracks validator signing info and enforces the policy.
func (s Server) HandleSigningInfo(ctx sdk.Context, voteInfo abci.VoteInfo, policy *validatorv1.Policy) error {
	height := ctx.BlockHeight()

	// set validator address
	address := sdk.ConsAddress(voteInfo.Validator.Address)

	// get validator signing info
	signingInfo, err := s.ss.ValidatorSigningInfoTable().Get(ctx, address.String())
	if err != nil {
		return err // internal error
	}

	// Compute the relative index, so we count the blocks the validator *should*
	// have signed. We will use the 0-value default signing info if not present,
	// except for start height. The index is in the range [0, SignedBlocksWindow)
	// and is used to see if a validator signed a block at the given height
	index := signingInfo.IndexOffset % policy.SignedBlocksWindow

	// increment index offset
	signingInfo.IndexOffset++

	// missed and missed previous
	missed := signingInfo.MissedBlocks[height].Missed
	missedPrevious := signingInfo.MissedBlocks[index].Missed

	switch {
	case missed && !missedPrevious:
		signingInfo.MissedBlocks[index].Missed = true
		signingInfo.MissedBlocksCount++
	case !missed && missedPrevious:
		signingInfo.MissedBlocks[index].Missed = false
		signingInfo.MissedBlocksCount--
	default:
		// bitmap value at this index has not changed
	}

	minSignedPerWindow := policy.MinSignedPerWindow

	if missed {
		// emit missed block event
		if err = ctx.EventManager().EmitTypedEvent(&v1.EventMissedBlock{
			Address: address.String(),
		}); err != nil {
			return err // internal error
		}
	}

	minHeight := signingInfo.StartHeight + policy.SignedBlocksWindow
	maxMissed := policy.SignedBlocksWindow - minSignedPerWindow

	// remove validator if validator missed blocks exceeds max missed blocks
	if height > minHeight && signingInfo.MissedBlocksCount > maxMissed {
		// get validator
		validator, err := s.ss.ValidatorTable().Get(ctx, address.String())
		if err != nil {
			return err // internal error
		}

		// delete validator
		err = s.ss.ValidatorTable().Delete(ctx, validator)
		if err != nil {
			return err // internal error
		}

		// delete validator signing info
		err = s.ss.ValidatorSigningInfoTable().Delete(ctx, signingInfo)
		if err != nil {
			return err // internal error
		}

		// emit remove validator event
		if err = ctx.EventManager().EmitTypedEvent(&v1.EventRemoveValidator{
			Address: address.String(),
		}); err != nil {
			return err // internal error
		}
	} else {
		// update validator signing info
		err = s.ss.ValidatorSigningInfoTable().Update(ctx, signingInfo)
		if err != nil {
			return err // internal error
		}
	}

	return nil
}
