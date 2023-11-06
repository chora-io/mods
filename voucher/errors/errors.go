package errors

import (
	"cosmossdk.io/errors"

	"github.com/chora-io/mods/voucher"
)

var (
	ErrParse = errors.Register(voucher.ModuleName, 2, "parse error")
)
