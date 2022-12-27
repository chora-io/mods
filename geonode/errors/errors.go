package errors

import (
	"cosmossdk.io/errors"

	"github.com/choraio/mods/geonode"
)

var (
	ErrParse = errors.Register(geonode.ModuleName, 2, "parse error")
)
