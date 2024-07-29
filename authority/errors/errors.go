package errors

import (
	"cosmossdk.io/errors"

	"github.com/chora-io/mods/authority"
)

var (
	ErrParse = errors.Register(authority.ModuleName, 2, "parse error")
)
