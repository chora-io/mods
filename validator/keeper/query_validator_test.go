package keeper

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	validatorv1 "github.com/chora-io/mods/validator/api/v1"
	v1 "github.com/chora-io/mods/validator/types/v1"
)

type queryValidator struct {
	*baseSuite
	res *v1.QueryValidatorResponse
	err error
}

func TestQueryValidator(t *testing.T) {
	gocuke.NewRunner(t, &queryValidator{}).
		Path("./features/query_validator.feature").
		Run()
}

func (s *queryValidator) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *queryValidator) Validator(a gocuke.DocString) {
	var validator validatorv1.Validator
	err := jsonpb.UnmarshalString(a.Content, &validator)
	require.NoError(s.t, err)

	err = s.k.ss.ValidatorTable().Insert(s.sdkCtx, &validatorv1.Validator{
		Address:  validator.Address,
		Metadata: validator.Metadata,
	})
	require.NoError(s.t, err)
}

func (s *queryValidator) QueryValidator(a gocuke.DocString) {
	var req v1.QueryValidatorRequest
	err := jsonpb.UnmarshalString(a.Content, &req)
	require.NoError(s.t, err)

	s.res, s.err = s.k.Validator(s.sdkCtx, &req)
}

func (s *queryValidator) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *queryValidator) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *queryValidator) ExpectResponse(a gocuke.DocString) {
	var expected v1.QueryValidatorResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
