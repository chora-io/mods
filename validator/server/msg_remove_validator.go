package server

import (
	"context"

	"cosmossdk.io/orm/types/ormerrors"
	v1 "github.com/choraio/mods/validator/types/v1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// RemoveValidator implements Msg/RemoveValidator.
func (s Server) RemoveValidator(ctx context.Context, req *v1.MsgRemoveValidator) (*v1.MsgRemoveValidatorResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	authorityAddress := s.authority.String()
	if authorityAddress != req.Authority {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"authority: expected %s: received %s", authorityAddress, req.Authority,
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

	// get signing info from validator signing info table
	signingInfo, err := s.ss.ValidatorSigningInfoTable().Get(ctx, req.Address)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf(
				"validator signing info with address %s", req.Address,
			)
		}
		return nil, err // internal error
	}

	// delete validator from validator table
	err = s.ss.ValidatorTable().Delete(ctx, validator)
	if err != nil {
		return nil, err // internal error
	}

	// delete validator from validator signing info table
	err = s.ss.ValidatorSigningInfoTable().Delete(ctx, signingInfo)
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventRemoveValidator{
		Address: req.Address,
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgRemoveValidatorResponse{
		Address: req.Address,
	}, nil
}
