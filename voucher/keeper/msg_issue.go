package keeper

import (
	"context"

	"cosmossdk.io/math"
	"cosmossdk.io/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/protobuf/types/known/timestamppb"

	voucherv1 "github.com/chora-io/mods/voucher/api/v1"
	v1 "github.com/chora-io/mods/voucher/types/v1"
)

// Issue implements Msg/Issue.
func (k Keeper) Issue(ctx context.Context, req *v1.MsgIssue) (*v1.MsgIssueResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get account from issuer address
	issuer, err := sdk.AccAddressFromBech32(req.Issuer)
	if err != nil {
		return nil, err // internal error
	}

	// get voucher from voucher table
	voucher, err := k.ss.VoucherTable().Get(ctx, req.Id)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("voucher with id %d: %s", req.Id, err)
		}
		return nil, err // internal error
	}

	// get account from account bytes
	voucherIssuer := sdk.AccAddress(voucher.Issuer)

	// verify issuer is voucher issuer
	if !voucherIssuer.Equals(issuer) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf(
			"issuer %s: voucher issuer %s", issuer, voucherIssuer.String(),
		)
	}

	// verify expiration is in the future
	if !req.Expiration.After(sdkCtx.BlockTime()) {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf(
			"expiration must be in the future: received %s", req.Expiration,
		)
	}

	// get account from recipient address
	recipient, err := sdk.AccAddressFromBech32(req.Recipient)
	if err != nil {
		return nil, err // internal error
	}

	// convert expiration to proto expiration
	expiration := timestamppb.New(*req.Expiration)

	// get current balance from state
	balance, err := k.ss.BalanceTable().Get(ctx, req.Id, recipient, expiration)
	if err != nil {
		if !ormerrors.NotFound.Is(err) {
			return nil, err // internal error
		}
		balance = &voucherv1.Balance{
			Id:         req.Id,
			Address:    recipient,
			Amount:     "0",
			Expiration: expiration,
		}
	}

	// convert balance amount to decimal
	balanceAmount, err := math.LegacyNewDecFromStr(balance.Amount)
	if err != nil {
		return nil, err // internal error
	}

	// convert request amount to decimal
	requestAmount, err := math.LegacyNewDecFromStr(req.Amount)
	if err != nil {
		return nil, err // internal error
	}

	// add request amount to balance amount
	newBalance := balanceAmount.Add(requestAmount)

	// set new balance
	balance.Amount = newBalance.String()

	// save balance to balance table
	err = k.ss.BalanceTable().Save(ctx, balance)
	if err != nil {
		return nil, err // internal error
	}

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventIssue{
		Id:       req.Id,
		Metadata: req.Metadata,
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgIssueResponse{
		Id: req.Id,
	}, nil
}
