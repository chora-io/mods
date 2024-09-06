package keeper

import (
	"context"
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/group"

	subjectv1 "github.com/chora-io/mods/subject/api/v1"
	v1 "github.com/chora-io/mods/subject/types/v1"
)

// CreateSubject implements Msg/CreateSubject.
func (k Keeper) CreateSubject(ctx context.Context, req *v1.MsgCreateSubject) (*v1.MsgCreateSubjectResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// get account from steward address
	steward, err := sdk.AccAddressFromBech32(req.Steward)
	if err != nil {
		return nil, err // internal error
	}

	// get sequence for key generation
	subjectSequence, err := k.ss.SubjectSequenceTable().Get(ctx)
	if err != nil {
		return nil, err // internal error
	}

	// generate account derivation key
	accountKey := make([]byte, 10)
	binary.LittleEndian.PutUint64(accountKey, subjectSequence.Sequence)

	// create module account using account derivation key
	subjectAccount, err := authtypes.NewModuleCredential(group.ModuleName, accountKey)
	if err != nil {
		return nil, err
	}

	// insert subject into subject table
	err = k.ss.SubjectTable().Insert(ctx, &subjectv1.Subject{
		Steward:  steward,
		Address:  subjectAccount.Address(),
		Metadata: req.Metadata,
	})
	if err != nil {
		return nil, err // internal error
	}

	// update subject sequence table
	err = k.ss.SubjectSequenceTable().Save(ctx, &subjectv1.SubjectSequence{
		Sequence: subjectSequence.Sequence + 1,
	})
	if err != nil {
		return nil, err // internal error
	}

	// get string address from subject account
	address := sdk.AccAddress(subjectAccount.Address()).String()

	// emit event
	if err = sdkCtx.EventManager().EmitTypedEvent(&v1.EventCreateSubject{
		Address: address,
	}); err != nil {
		return nil, err // internal error
	}

	// return response
	return &v1.MsgCreateSubjectResponse{
		Address: address,
	}, nil
}
