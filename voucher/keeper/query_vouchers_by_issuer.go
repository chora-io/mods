package keeper

import (
	"context"

	"cosmossdk.io/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	voucherv1 "github.com/chora-io/mods/voucher/api/v1"
	v1 "github.com/chora-io/mods/voucher/types/v1"
	"github.com/chora-io/mods/voucher/utils"
)

// VouchersByIssuer implements the Query/VouchersByIssuer method.
func (k Keeper) VouchersByIssuer(ctx context.Context, req *v1.QueryVouchersByIssuerRequest) (*v1.QueryVouchersByIssuerResponse, error) {

	// get account from issuer address
	issuer, err := sdk.AccAddressFromBech32(req.Issuer)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("issuer: %s", err)
	}

	// set index for table lookup
	index := voucherv1.VoucherIssuerIndexKey{}.WithIssuer(issuer)

	// set pagination for table lookup
	pgnReq, err := utils.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err // internal error
	}

	// get vouchers by issuer from voucher table
	it, err := k.ss.VoucherTable().List(ctx, index, ormlist.Paginate(pgnReq))
	if err != nil {
		return nil, err // internal error
	}

	// set vouchers for query response
	vouchers := make([]*v1.QueryVouchersByIssuerResponse_Voucher, 0, 10)
	for it.Next() {
		v, err := it.Value()
		if err != nil {
			return nil, err // internal error
		}

		voucher := v1.QueryVouchersByIssuerResponse_Voucher{
			Id:       v.Id,
			Metadata: v.Metadata,
		}

		vouchers = append(vouchers, &voucher)
	}

	// set pagination for query response
	pgnRes, err := utils.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err // internal error
	}

	// return query response
	return &v1.QueryVouchersByIssuerResponse{
		Issuer:     issuer.String(),
		Vouchers:   vouchers,
		Pagination: pgnRes,
	}, nil
}
