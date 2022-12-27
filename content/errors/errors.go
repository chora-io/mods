package errors

import (
	"cosmossdk.io/errors"

	"github.com/choraio/mods/content"
)

var (
	ErrParse = errors.Register(content.ModuleName, 2, "parse error")
)
