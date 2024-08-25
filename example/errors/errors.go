package errors

import (
	"cosmossdk.io/errors"

	"github.com/chora-io/mods/example"
)

var (
	ErrParse = errors.Register(example.ModuleName, 2, "parse error")
)
