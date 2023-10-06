package server

import (
	"context"

	"cosmossdk.io/orm/types/ormerrors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/choraio/mods/validator/types/v1"
)

// Validator implements the Query/Validator method.
func (s Server) Validator(ctx context.Context, req *v1.QueryValidatorRequest) (*v1.QueryValidatorResponse, error) {

	// get validator from validator table
	validator, err := s.ss.ValidatorTable().Get(ctx, req.Address)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("validator with address %s", req.Address)
		}
		return nil, err // internal error
	}

	// return query response
	return &v1.QueryValidatorResponse{
		Address:  validator.Address,
		Metadata: validator.Metadata,
	}, nil
}
