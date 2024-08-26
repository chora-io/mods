package errors

import (
	"cosmossdk.io/errors"

	"github.com/chora-io/mods/agent"
)

var (
	ErrParse = errors.Register(agent.ModuleName, 2, "parse error")
)
