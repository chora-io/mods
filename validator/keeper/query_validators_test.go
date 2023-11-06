package keeper

import (
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	validatorv1 "github.com/chora-io/mods/validator/api/v1"
	v1 "github.com/chora-io/mods/validator/types/v1"
)

type queryValidators struct {
	*baseSuite
	res *v1.QueryValidatorsResponse
	err error
}

func TestQueryValidators(t *testing.T) {
	gocuke.NewRunner(t, &queryValidators{}).
		Path("./features/query_validators.feature").
		Run()
}

func (s *queryValidators) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *queryValidators) Validator(a gocuke.DocString) {
	var validator validatorv1.Validator
	err := jsonpb.UnmarshalString(a.Content, &validator)
	require.NoError(s.t, err)

	err = s.k.ss.ValidatorTable().Insert(s.sdkCtx, &validatorv1.Validator{
		Address:  validator.Address,
		Metadata: validator.Metadata,
	})
	require.NoError(s.t, err)
}

func (s *queryValidators) QueryValidators(a gocuke.DocString) {
	var req v1.QueryValidatorsRequest
	err := jsonpb.UnmarshalString(a.Content, &req)
	require.NoError(s.t, err)

	s.res, s.err = s.k.Validators(s.sdkCtx, &req)
}

func (s *queryValidators) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *queryValidators) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *queryValidators) ExpectResponse(a gocuke.DocString) {
	var expected v1.QueryValidatorsResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
