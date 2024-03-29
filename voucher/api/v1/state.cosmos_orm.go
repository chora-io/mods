// Code generated by protoc-gen-go-cosmos-orm. DO NOT EDIT.

package voucherv1

import (
	context "context"
	ormlist "cosmossdk.io/orm/model/ormlist"
	ormtable "cosmossdk.io/orm/model/ormtable"
	ormerrors "cosmossdk.io/orm/types/ormerrors"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type VoucherTable interface {
	Insert(ctx context.Context, voucher *Voucher) error
	InsertReturningId(ctx context.Context, voucher *Voucher) (uint64, error)
	LastInsertedSequence(ctx context.Context) (uint64, error)
	Update(ctx context.Context, voucher *Voucher) error
	Save(ctx context.Context, voucher *Voucher) error
	Delete(ctx context.Context, voucher *Voucher) error
	Has(ctx context.Context, id uint64) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, id uint64) (*Voucher, error)
	List(ctx context.Context, prefixKey VoucherIndexKey, opts ...ormlist.Option) (VoucherIterator, error)
	ListRange(ctx context.Context, from, to VoucherIndexKey, opts ...ormlist.Option) (VoucherIterator, error)
	DeleteBy(ctx context.Context, prefixKey VoucherIndexKey) error
	DeleteRange(ctx context.Context, from, to VoucherIndexKey) error

	doNotImplement()
}

type VoucherIterator struct {
	ormtable.Iterator
}

func (i VoucherIterator) Value() (*Voucher, error) {
	var voucher Voucher
	err := i.UnmarshalMessage(&voucher)
	return &voucher, err
}

type VoucherIndexKey interface {
	id() uint32
	values() []interface{}
	voucherIndexKey()
}

// primary key starting index..
type VoucherPrimaryKey = VoucherIdIndexKey

type VoucherIdIndexKey struct {
	vs []interface{}
}

func (x VoucherIdIndexKey) id() uint32            { return 0 }
func (x VoucherIdIndexKey) values() []interface{} { return x.vs }
func (x VoucherIdIndexKey) voucherIndexKey()      {}

func (this VoucherIdIndexKey) WithId(id uint64) VoucherIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type VoucherIssuerIndexKey struct {
	vs []interface{}
}

func (x VoucherIssuerIndexKey) id() uint32            { return 1 }
func (x VoucherIssuerIndexKey) values() []interface{} { return x.vs }
func (x VoucherIssuerIndexKey) voucherIndexKey()      {}

func (this VoucherIssuerIndexKey) WithIssuer(issuer []byte) VoucherIssuerIndexKey {
	this.vs = []interface{}{issuer}
	return this
}

type voucherTable struct {
	table ormtable.AutoIncrementTable
}

func (this voucherTable) Insert(ctx context.Context, voucher *Voucher) error {
	return this.table.Insert(ctx, voucher)
}

func (this voucherTable) Update(ctx context.Context, voucher *Voucher) error {
	return this.table.Update(ctx, voucher)
}

func (this voucherTable) Save(ctx context.Context, voucher *Voucher) error {
	return this.table.Save(ctx, voucher)
}

func (this voucherTable) Delete(ctx context.Context, voucher *Voucher) error {
	return this.table.Delete(ctx, voucher)
}

func (this voucherTable) InsertReturningId(ctx context.Context, voucher *Voucher) (uint64, error) {
	return this.table.InsertReturningPKey(ctx, voucher)
}

func (this voucherTable) LastInsertedSequence(ctx context.Context) (uint64, error) {
	return this.table.LastInsertedSequence(ctx)
}

func (this voucherTable) Has(ctx context.Context, id uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this voucherTable) Get(ctx context.Context, id uint64) (*Voucher, error) {
	var voucher Voucher
	found, err := this.table.PrimaryKey().Get(ctx, &voucher, id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &voucher, nil
}

func (this voucherTable) List(ctx context.Context, prefixKey VoucherIndexKey, opts ...ormlist.Option) (VoucherIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return VoucherIterator{it}, err
}

func (this voucherTable) ListRange(ctx context.Context, from, to VoucherIndexKey, opts ...ormlist.Option) (VoucherIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return VoucherIterator{it}, err
}

func (this voucherTable) DeleteBy(ctx context.Context, prefixKey VoucherIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this voucherTable) DeleteRange(ctx context.Context, from, to VoucherIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this voucherTable) doNotImplement() {}

var _ VoucherTable = voucherTable{}

func NewVoucherTable(db ormtable.Schema) (VoucherTable, error) {
	table := db.GetTable(&Voucher{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Voucher{}).ProtoReflect().Descriptor().FullName()))
	}
	return voucherTable{table.(ormtable.AutoIncrementTable)}, nil
}

type BalanceTable interface {
	Insert(ctx context.Context, balance *Balance) error
	Update(ctx context.Context, balance *Balance) error
	Save(ctx context.Context, balance *Balance) error
	Delete(ctx context.Context, balance *Balance) error
	Has(ctx context.Context, id uint64, address []byte, expiration *timestamppb.Timestamp) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, id uint64, address []byte, expiration *timestamppb.Timestamp) (*Balance, error)
	List(ctx context.Context, prefixKey BalanceIndexKey, opts ...ormlist.Option) (BalanceIterator, error)
	ListRange(ctx context.Context, from, to BalanceIndexKey, opts ...ormlist.Option) (BalanceIterator, error)
	DeleteBy(ctx context.Context, prefixKey BalanceIndexKey) error
	DeleteRange(ctx context.Context, from, to BalanceIndexKey) error

	doNotImplement()
}

type BalanceIterator struct {
	ormtable.Iterator
}

func (i BalanceIterator) Value() (*Balance, error) {
	var balance Balance
	err := i.UnmarshalMessage(&balance)
	return &balance, err
}

type BalanceIndexKey interface {
	id() uint32
	values() []interface{}
	balanceIndexKey()
}

// primary key starting index..
type BalancePrimaryKey = BalanceIdAddressExpirationIndexKey

type BalanceIdAddressExpirationIndexKey struct {
	vs []interface{}
}

func (x BalanceIdAddressExpirationIndexKey) id() uint32            { return 0 }
func (x BalanceIdAddressExpirationIndexKey) values() []interface{} { return x.vs }
func (x BalanceIdAddressExpirationIndexKey) balanceIndexKey()      {}

func (this BalanceIdAddressExpirationIndexKey) WithId(id uint64) BalanceIdAddressExpirationIndexKey {
	this.vs = []interface{}{id}
	return this
}

func (this BalanceIdAddressExpirationIndexKey) WithIdAddress(id uint64, address []byte) BalanceIdAddressExpirationIndexKey {
	this.vs = []interface{}{id, address}
	return this
}

func (this BalanceIdAddressExpirationIndexKey) WithIdAddressExpiration(id uint64, address []byte, expiration *timestamppb.Timestamp) BalanceIdAddressExpirationIndexKey {
	this.vs = []interface{}{id, address, expiration}
	return this
}

type BalanceAddressIndexKey struct {
	vs []interface{}
}

func (x BalanceAddressIndexKey) id() uint32            { return 1 }
func (x BalanceAddressIndexKey) values() []interface{} { return x.vs }
func (x BalanceAddressIndexKey) balanceIndexKey()      {}

func (this BalanceAddressIndexKey) WithAddress(address []byte) BalanceAddressIndexKey {
	this.vs = []interface{}{address}
	return this
}

type BalanceExpirationIndexKey struct {
	vs []interface{}
}

func (x BalanceExpirationIndexKey) id() uint32            { return 2 }
func (x BalanceExpirationIndexKey) values() []interface{} { return x.vs }
func (x BalanceExpirationIndexKey) balanceIndexKey()      {}

func (this BalanceExpirationIndexKey) WithExpiration(expiration *timestamppb.Timestamp) BalanceExpirationIndexKey {
	this.vs = []interface{}{expiration}
	return this
}

type balanceTable struct {
	table ormtable.Table
}

func (this balanceTable) Insert(ctx context.Context, balance *Balance) error {
	return this.table.Insert(ctx, balance)
}

func (this balanceTable) Update(ctx context.Context, balance *Balance) error {
	return this.table.Update(ctx, balance)
}

func (this balanceTable) Save(ctx context.Context, balance *Balance) error {
	return this.table.Save(ctx, balance)
}

func (this balanceTable) Delete(ctx context.Context, balance *Balance) error {
	return this.table.Delete(ctx, balance)
}

func (this balanceTable) Has(ctx context.Context, id uint64, address []byte, expiration *timestamppb.Timestamp) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id, address, expiration)
}

func (this balanceTable) Get(ctx context.Context, id uint64, address []byte, expiration *timestamppb.Timestamp) (*Balance, error) {
	var balance Balance
	found, err := this.table.PrimaryKey().Get(ctx, &balance, id, address, expiration)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &balance, nil
}

func (this balanceTable) List(ctx context.Context, prefixKey BalanceIndexKey, opts ...ormlist.Option) (BalanceIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return BalanceIterator{it}, err
}

func (this balanceTable) ListRange(ctx context.Context, from, to BalanceIndexKey, opts ...ormlist.Option) (BalanceIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return BalanceIterator{it}, err
}

func (this balanceTable) DeleteBy(ctx context.Context, prefixKey BalanceIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this balanceTable) DeleteRange(ctx context.Context, from, to BalanceIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this balanceTable) doNotImplement() {}

var _ BalanceTable = balanceTable{}

func NewBalanceTable(db ormtable.Schema) (BalanceTable, error) {
	table := db.GetTable(&Balance{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Balance{}).ProtoReflect().Descriptor().FullName()))
	}
	return balanceTable{table}, nil
}

type StateStore interface {
	VoucherTable() VoucherTable
	BalanceTable() BalanceTable

	doNotImplement()
}

type stateStore struct {
	voucher VoucherTable
	balance BalanceTable
}

func (x stateStore) VoucherTable() VoucherTable {
	return x.voucher
}

func (x stateStore) BalanceTable() BalanceTable {
	return x.balance
}

func (stateStore) doNotImplement() {}

var _ StateStore = stateStore{}

func NewStateStore(db ormtable.Schema) (StateStore, error) {
	voucherTable, err := NewVoucherTable(db)
	if err != nil {
		return nil, err
	}

	balanceTable, err := NewBalanceTable(db)
	if err != nil {
		return nil, err
	}

	return stateStore{
		voucherTable,
		balanceTable,
	}, nil
}
