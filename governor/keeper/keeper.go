package keeper

import (
	"encoding/json"

	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	abci "github.com/cometbft/cometbft/abci/types"
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

// GetPolicy gets the governor signing policy.
func (k Keeper) GetPolicy(ctx sdk.Context) (*governorv1.Policy, error) {
	policy, err := k.ss.PolicyTable().Get(ctx)
	if err != nil {
		return nil, err
	}
	return policy, nil
}

// HandleSigningInfo tracks governor signing info and enforces the policy.
func (k Keeper) HandleSigningInfo(ctx sdk.Context, voteInfo abci.VoteInfo, policy *governorv1.Policy) error {
	height := ctx.BlockHeight()

	// set governor address
	address := sdk.ConsAddress(voteInfo.Governor.Address)

	// get governor signing info
	signingInfo, err := k.ss.GovernorSigningInfoTable().Get(ctx, address.String())
	if err != nil {
		return err // internal error
	}

	// Compute the relative index, so we count the blocks the governor *should*
	// have signed. We will use the 0-value default signing info if not present,
	// except for start height. The index is in the range [0, SignedBlocksWindow)
	// and is used to see if a governor signed a block at the given height
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

	// remove governor if governor missed blocks exceeds max missed blocks
	if height > minHeight && signingInfo.MissedBlocksCount > maxMissed {
		// get governor
		governor, err := k.ss.GovernorTable().Get(ctx, address.String())
		if err != nil {
			return err // internal error
		}

		// delete governor
		err = k.ss.GovernorTable().Delete(ctx, governor)
		if err != nil {
			return err // internal error
		}

		// delete governor signing info
		err = k.ss.GovernorSigningInfoTable().Delete(ctx, signingInfo)
		if err != nil {
			return err // internal error
		}

		// emit remove governor event
		if err = ctx.EventManager().EmitTypedEvent(&v1.EventRemoveGovernor{
			Address: address.String(),
		}); err != nil {
			return err // internal error
		}
	} else {
		// update governor signing info
		err = k.ss.GovernorSigningInfoTable().Update(ctx, signingInfo)
		if err != nil {
			return err // internal error
		}
	}

	return nil
}
