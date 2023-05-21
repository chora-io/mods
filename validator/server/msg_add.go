package server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	validatorv1 "github.com/choraio/mods/validator/api/v1"
	v1 "github.com/choraio/mods/validator/types/v1"
)

// Add implements Msg/Add.
func (s Server) Add(ctx context.Context, req *v1.MsgAdd) (*v1.MsgAddResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	if s.authority.String() != req.Authority {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"authority: expected %s: received %s", s.authority.String(), req.Authority,
		)
	}

	// insert validator into validator table
	err := s.ss.ValidatorTable().Insert(ctx, &validatorv1.Validator{
		Address:  req.Address,
		Metadata: req.Metadata,
	})
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventAdd{
		Address: req.Address,
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgAddResponse{
		Address: req.Address,
	}, nil
}
