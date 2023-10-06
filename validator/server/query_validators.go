package server

import (
	"context"

	"cosmossdk.io/orm/model/ormlist"

	validatorv1 "github.com/choraio/mods/validator/api/v1"
	v1 "github.com/choraio/mods/validator/types/v1"
	"github.com/choraio/mods/validator/utils"
)

// Validators implements the Query/Validators method.
func (s Server) Validators(ctx context.Context, req *v1.QueryValidatorsRequest) (*v1.QueryValidatorsResponse, error) {

	// set the index for table lookup
	index := validatorv1.ValidatorAddressIndexKey{}

	// set the pagination for table lookup
	pgnReq, err := utils.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err // internal error
	}

	// get validators from validator table
	it, err := s.ss.ValidatorTable().List(ctx, index, ormlist.Paginate(pgnReq))
	if err != nil {
		return nil, err // internal error
	}

	// set validators for query response
	validators := make([]*v1.QueryValidatorsResponse_Validator, 0, 10)
	for it.Next() {
		v, err := it.Value()
		if err != nil {
			return nil, err // internal error
		}

		validator := v1.QueryValidatorsResponse_Validator{
			Address:  v.Address,
			Metadata: v.Metadata,
		}

		validators = append(validators, &validator)
	}

	// set the pagination for query response
	pgnRes, err := utils.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err // internal error
	}

	// return query response
	return &v1.QueryValidatorsResponse{
		Validators: validators,
		Pagination: pgnRes,
	}, nil
}
