package utils

import (
	"context"

	ormv1alpha1 "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/types/kv"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewStoreKeyDB creates an ormdb.ModuleDB from an ormdb.ModuleDB and a StoreKey.
// It is an interim solution for using the ORM in existing Cosmos SDK modules
// before fuller integration has been done.
func NewStoreKeyDB(desc *ormv1alpha1.ModuleSchemaDescriptor, key storetypes.StoreKey, options ormdb.ModuleDBOptions) (ormdb.ModuleDB, error) {
	backEndResolver := func(_ ormv1alpha1.StorageType) (ormtable.BackendResolver, error) {
		getBackend := func(ctx context.Context) (ormtable.ReadBackend, error) {
			sdkCtx := sdk.UnwrapSDKContext(ctx)
			store := sdkCtx.KVStore(key)
			wrapper := storeWrapper{store}
			return ormtable.NewBackend(ormtable.BackendOptions{
				CommitmentStore: wrapper,
				IndexStore:      wrapper,
			}), nil
		}
		return getBackend, nil
	}
	options.GetBackendResolver = backEndResolver
	return ormdb.NewModuleDB(desc, options)
}

type storeWrapper struct {
	store storetypes.KVStore
}

func (k storeWrapper) Set(key, value []byte) error {
	k.store.Set(key, value)
	return nil
}

func (k storeWrapper) Delete(key []byte) error {
	k.store.Delete(key)
	return nil
}

func (k storeWrapper) Get(key []byte) ([]byte, error) {
	x := k.store.Get(key)
	return x, nil
}

func (k storeWrapper) Has(key []byte) (bool, error) {
	x := k.store.Has(key)
	return x, nil
}

func (k storeWrapper) Iterator(start, end []byte) (kv.Iterator, error) {
	x := k.store.Iterator(start, end)
	return x, nil
}

func (k storeWrapper) ReverseIterator(start, end []byte) (kv.Iterator, error) {
	x := k.store.ReverseIterator(start, end)
	return x, nil
}

var _ kv.Store = &storeWrapper{}
