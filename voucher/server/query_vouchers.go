package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

	voucherv1 "github.com/choraio/mods/voucher/api/v1"
	v1 "github.com/choraio/mods/voucher/types/v1"
	"github.com/choraio/mods/voucher/utils"
)

// Vouchers implements the Query/Vouchers method.
func (s Server) Vouchers(ctx context.Context, req *v1.QueryVouchersRequest) (*v1.QueryVouchersResponse, error) {

	// set index for table lookup
	index := voucherv1.VoucherIdIndexKey{}

	// set pagination for table lookup
	pgnReq, err := utils.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err // internal error
	}

	// get vouchers from voucher table
	it, err := s.ss.VoucherTable().List(ctx, index, ormlist.Paginate(pgnReq))
	if err != nil {
		return nil, err // internal error
	}

	// set vouchers for query response
	vouchers := make([]*v1.QueryVouchersResponse_Voucher, 0, 10)
	for it.Next() {
		v, err := it.Value()
		if err != nil {
			return nil, err // internal error
		}

		issuer := sdk.AccAddress(v.Issuer).String()

		voucher := v1.QueryVouchersResponse_Voucher{
			Id:       v.Id,
			Issuer:   issuer,
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
	return &v1.QueryVouchersResponse{
		Vouchers:   vouchers,
		Pagination: pgnRes,
	}, nil
}
