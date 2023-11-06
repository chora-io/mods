package errors

import (
	"cosmossdk.io/errors"

	"github.com/chora-io/mods/content"
)

var (
	ErrParse = errors.Register(content.ModuleName, 2, "parse error")
)
