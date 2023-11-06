package errors

import (
	"cosmossdk.io/errors"

	"github.com/chora-io/mods/geonode"
)

var (
	ErrParse = errors.Register(geonode.ModuleName, 2, "parse error")
)
