package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	voucherv1 "github.com/choraio/mods/voucher/api/v1"
	v1 "github.com/choraio/mods/voucher/types/v1"
	"github.com/choraio/mods/voucher/utils"
)

// VouchersByIssuer implements the Query/VouchersByIssuer method.
func (s Server) VouchersByIssuer(ctx context.Context, req *v1.QueryVouchersByIssuerRequest) (*v1.QueryVouchersByIssuerResponse, error) {

	// get account from issuer address
	issuer, err := sdk.AccAddressFromBech32(req.Issuer)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("issuer: %s", err)
	}

	// set index for table lookup
	index := voucherv1.VoucherIssuerIndexKey{}.WithIssuer(issuer)

	// set pagination for table lookup
	pg, err := utils.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err // internal error
	}

	// get vouchers by issuer from voucher table
	it, err := s.ss.VoucherTable().List(ctx, index, ormlist.Paginate(pg))
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
		n := v1.QueryVouchersByIssuerResponse_Voucher{
			Id:       v.Id,
			Metadata: v.Metadata,
		}
		vouchers = append(vouchers, &n)
	}

	// set pagination for query response
	pr, err := utils.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err // internal error
	}

	// return query response
	return &v1.QueryVouchersByIssuerResponse{
		Issuer:     issuer.String(),
		Vouchers:   vouchers,
		Pagination: pr,
	}, nil
}
