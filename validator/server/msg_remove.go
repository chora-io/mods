package server

import (
	"context"

	v1 "github.com/choraio/mods/validator/types/v1"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Remove implements Msg/Remove.
func (s Server) Remove(ctx context.Context, req *v1.MsgRemove) (*v1.MsgRemoveResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	if s.authority.String() != req.Authority {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"authority: expected %s: received %s", s.authority.String(), req.Authority,
		)
	}

	// get validator from validator table
	validator, err := s.ss.ValidatorTable().Get(ctx, req.Address)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf(
				"validator with address %s: %s", req.Address, err,
			)
		}
		return nil, err // internal error
	}

	// delete validator from validator table
	err = s.ss.ValidatorTable().Delete(ctx, validator)
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventRemove{
		Address: req.Address,
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgRemoveResponse{
		Address: req.Address,
	}, nil
}
