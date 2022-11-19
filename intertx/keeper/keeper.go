package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ctypes "github.com/cosmos/cosmos-sdk/x/capability/types"

	"github.com/choraio/mods/intertx"
)

type Keeper struct {
	cdc codec.BinaryCodec

	scopedKeeper        CapabilityKeeper
	icaControllerKeeper ICAControllerKeeper
}

func NewKeeper(cdc codec.BinaryCodec, iaKeeper ICAControllerKeeper, scopedKeeper CapabilityKeeper) Keeper {
	return Keeper{
		cdc:                 cdc,
		scopedKeeper:        scopedKeeper,
		icaControllerKeeper: iaKeeper,
	}
}

// Logger returns the application logger, scoped to the associated module
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", intertx.ModuleName))
}

// ClaimCapability claims the channel capability passed via the OnOpenChanInit callback
func (k Keeper) ClaimCapability(ctx sdk.Context, cap *ctypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}
