package errors

import (
	"cosmossdk.io/errors"

	"github.com/chora-io/mods/validator"
)

var (
	ErrParse = errors.Register(validator.ModuleName, 2, "parse error")
)
