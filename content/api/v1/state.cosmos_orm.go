// Code generated by protoc-gen-go-cosmos-orm. DO NOT EDIT.

package contentv1

import (
	context "context"
	ormlist "cosmossdk.io/orm/model/ormlist"
	ormtable "cosmossdk.io/orm/model/ormtable"
	ormerrors "cosmossdk.io/orm/types/ormerrors"
)

type ContentTable interface {
	Insert(ctx context.Context, content *Content) error
	Update(ctx context.Context, content *Content) error
	Save(ctx context.Context, content *Content) error
	Delete(ctx context.Context, content *Content) error
	Has(ctx context.Context, hash string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, hash string) (*Content, error)
	List(ctx context.Context, prefixKey ContentIndexKey, opts ...ormlist.Option) (ContentIterator, error)
	ListRange(ctx context.Context, from, to ContentIndexKey, opts ...ormlist.Option) (ContentIterator, error)
	DeleteBy(ctx context.Context, prefixKey ContentIndexKey) error
	DeleteRange(ctx context.Context, from, to ContentIndexKey) error

	doNotImplement()
}

type ContentIterator struct {
	ormtable.Iterator
}

func (i ContentIterator) Value() (*Content, error) {
	var content Content
	err := i.UnmarshalMessage(&content)
	return &content, err
}

type ContentIndexKey interface {
	id() uint32
	values() []interface{}
	contentIndexKey()
}

// primary key starting index..
type ContentPrimaryKey = ContentHashIndexKey

type ContentHashIndexKey struct {
	vs []interface{}
}

func (x ContentHashIndexKey) id() uint32            { return 0 }
func (x ContentHashIndexKey) values() []interface{} { return x.vs }
func (x ContentHashIndexKey) contentIndexKey()      {}

func (this ContentHashIndexKey) WithHash(hash string) ContentHashIndexKey {
	this.vs = []interface{}{hash}
	return this
}

type ContentCuratorIndexKey struct {
	vs []interface{}
}

func (x ContentCuratorIndexKey) id() uint32            { return 1 }
func (x ContentCuratorIndexKey) values() []interface{} { return x.vs }
func (x ContentCuratorIndexKey) contentIndexKey()      {}

func (this ContentCuratorIndexKey) WithCurator(curator []byte) ContentCuratorIndexKey {
	this.vs = []interface{}{curator}
	return this
}

type contentTable struct {
	table ormtable.Table
}

func (this contentTable) Insert(ctx context.Context, content *Content) error {
	return this.table.Insert(ctx, content)
}

func (this contentTable) Update(ctx context.Context, content *Content) error {
	return this.table.Update(ctx, content)
}

func (this contentTable) Save(ctx context.Context, content *Content) error {
	return this.table.Save(ctx, content)
}

func (this contentTable) Delete(ctx context.Context, content *Content) error {
	return this.table.Delete(ctx, content)
}

func (this contentTable) Has(ctx context.Context, hash string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, hash)
}

func (this contentTable) Get(ctx context.Context, hash string) (*Content, error) {
	var content Content
	found, err := this.table.PrimaryKey().Get(ctx, &content, hash)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &content, nil
}

func (this contentTable) List(ctx context.Context, prefixKey ContentIndexKey, opts ...ormlist.Option) (ContentIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return ContentIterator{it}, err
}

func (this contentTable) ListRange(ctx context.Context, from, to ContentIndexKey, opts ...ormlist.Option) (ContentIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return ContentIterator{it}, err
}

func (this contentTable) DeleteBy(ctx context.Context, prefixKey ContentIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this contentTable) DeleteRange(ctx context.Context, from, to ContentIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this contentTable) doNotImplement() {}

var _ ContentTable = contentTable{}

func NewContentTable(db ormtable.Schema) (ContentTable, error) {
	table := db.GetTable(&Content{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Content{}).ProtoReflect().Descriptor().FullName()))
	}
	return contentTable{table}, nil
}

type StateStore interface {
	ContentTable() ContentTable

	doNotImplement()
}

type stateStore struct {
	content ContentTable
}

func (x stateStore) ContentTable() ContentTable {
	return x.content
}

func (stateStore) doNotImplement() {}

var _ StateStore = stateStore{}

func NewStateStore(db ormtable.Schema) (StateStore, error) {
	contentTable, err := NewContentTable(db)
	if err != nil {
		return nil, err
	}

	return stateStore{
		contentTable,
	}, nil
}
