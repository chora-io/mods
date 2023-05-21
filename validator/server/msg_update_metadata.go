package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/choraio/mods/validator/types/v1"
)

// UpdateMetadata implements the Msg/UpdateMetadata method.
func (s Server) UpdateMetadata(ctx context.Context, req *v1.MsgUpdateMetadata) (*v1.MsgUpdateMetadataResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get validator from validator table
	validator, err := s.ss.ValidatorTable().Get(ctx, req.Address)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("validator with address %s: %s", req.Address, err)
		}
		return nil, err // internal error
	}

	// set new validator metadata
	validator.Metadata = req.NewMetadata

	// update validator in validator table
	err = s.ss.ValidatorTable().Update(ctx, validator)
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventUpdateMetadata{
		Address: req.Address,
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgUpdateMetadataResponse{
		Address: req.Address,
	}, nil
}
