package server

import (
	"context"

	"cosmossdk.io/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "github.com/choraio/mods/voucher/types/v1"
)

// Voucher implements the Query/Voucher method.
func (s Server) Voucher(ctx context.Context, req *v1.QueryVoucherRequest) (*v1.QueryVoucherResponse, error) {

	// get voucher from voucher table
	voucher, err := s.ss.VoucherTable().Get(ctx, req.Id)
	if err != nil {
		if ormerrors.NotFound.Is(err) {
			return nil, sdkerrors.ErrNotFound.Wrapf("voucher with id %d", req.Id)
		}
		return nil, err // internal error
	}

	// get account from account bytes
	issuer := sdk.AccAddress(voucher.Issuer)

	// return query response
	return &v1.QueryVoucherResponse{
		Id:       voucher.Id,
		Issuer:   issuer.String(),
		Metadata: voucher.Metadata,
	}, nil
}
