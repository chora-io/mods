package errors

import (
	"cosmossdk.io/errors"

	"github.com/chora-io/mods/admin"
)

var (
	ErrParse = errors.Register(admin.ModuleName, 2, "parse error")
)
