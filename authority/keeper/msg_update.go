package keeper

import (
	"context"

	"cosmossdk.io/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/chora-io/mods/authority/types/v1"
)

// Update implements the Msg/Update method.
func (k Keeper) Update(ctx context.Context, req *v1.MsgUpdate) (*v1.MsgUpdateResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get account from authority address
	msgSigner, err := sdk.AccAddressFromBech32(req.Authority)
	if err != nil {
		return nil, err // internal error
	}

	// get authority from authority table
	authority, err := k.ss.AuthorityTable().Get(ctx)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("authority: %s", err)
		}
		return nil, err // internal error
	}

	// get account from account bytes
	authorityAuthority := sdk.AccAddress(authority.Address)

	// verify authority is authority account
	if !authorityAuthority.Equals(msgSigner) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"authority %s: authority account %s", msgSigner, authorityAuthority.String(),
		)
	}

	// get account from new authority address
	newAuthority, err := sdk.AccAddressFromBech32(req.NewAuthority)
	if err != nil {
		return nil, err // internal error
	}

	// set new authority
	authority.Address = newAuthority

	// update authority in authority table
	err = k.ss.AuthorityTable().Save(ctx, authority)
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventUpdate{
		Authority:    msgSigner.String(),
		NewAuthority: newAuthority.String(),
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgUpdateResponse{
		Authority:    msgSigner.String(),
		NewAuthority: newAuthority.String(),
	}, nil
}
