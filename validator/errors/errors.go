package errors

import (
	"cosmossdk.io/errors"

	"github.com/choraio/mods/validator"
)

var (
	ErrParse = errors.Register(validator.ModuleName, 2, "parse error")
)
