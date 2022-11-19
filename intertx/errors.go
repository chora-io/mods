package intertx

import (
	"cosmossdk.io/errors"
	"google.golang.org/grpc/codes"
)

var (
	ErrInvalidArgument = errors.RegisterWithGRPCCode(ModuleName, 3, codes.InvalidArgument, "invalid argument")
	ErrNotFound        = errors.RegisterWithGRPCCode(ModuleName, 5, codes.NotFound, "not found")
)
