package errors

import (
	"cosmossdk.io/errors"

	"github.com/chora-io/mods/subject"
)

var (
	ErrParse = errors.Register(subject.ModuleName, 2, "parse error")
)
