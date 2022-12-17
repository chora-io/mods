package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/choraio/mods/example"
)

// RegisterInvariants registers the ecocredit module invariants.
func (s Server) RegisterInvariants(registry sdk.InvariantRegistry) {
	registry.RegisterRoute(example.ModuleName, "content", s.contentInvariant())
}

func (s Server) contentInvariant() sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		return "", false
	}
}
