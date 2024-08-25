package errors

import (
	"cosmossdk.io/errors"

	"github.com/chora-io/mods/governor"
)

var (
	ErrParse = errors.Register(governor.ModuleName, 2, "parse error")
)
