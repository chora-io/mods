package errors

import (
	"cosmossdk.io/errors"

	"github.com/chora-io/mods/ecosystem"
)

var (
	ErrParse = errors.Register(ecosystem.ModuleName, 2, "parse error")
)
