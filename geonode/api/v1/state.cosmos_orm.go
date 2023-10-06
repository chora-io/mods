// Code generated by protoc-gen-go-cosmos-orm. DO NOT EDIT.

package geonodev1

import (
	context "context"
	ormlist "cosmossdk.io/orm/model/ormlist"
	ormtable "cosmossdk.io/orm/model/ormtable"
	ormerrors "cosmossdk.io/orm/types/ormerrors"
)

type NodeTable interface {
	Insert(ctx context.Context, node *Node) error
	InsertReturningId(ctx context.Context, node *Node) (uint64, error)
	LastInsertedSequence(ctx context.Context) (uint64, error)
	Update(ctx context.Context, node *Node) error
	Save(ctx context.Context, node *Node) error
	Delete(ctx context.Context, node *Node) error
	Has(ctx context.Context, id uint64) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, id uint64) (*Node, error)
	List(ctx context.Context, prefixKey NodeIndexKey, opts ...ormlist.Option) (NodeIterator, error)
	ListRange(ctx context.Context, from, to NodeIndexKey, opts ...ormlist.Option) (NodeIterator, error)
	DeleteBy(ctx context.Context, prefixKey NodeIndexKey) error
	DeleteRange(ctx context.Context, from, to NodeIndexKey) error

	doNotImplement()
}

type NodeIterator struct {
	ormtable.Iterator
}

func (i NodeIterator) Value() (*Node, error) {
	var node Node
	err := i.UnmarshalMessage(&node)
	return &node, err
}

type NodeIndexKey interface {
	id() uint32
	values() []interface{}
	nodeIndexKey()
}

// primary key starting index..
type NodePrimaryKey = NodeIdIndexKey

type NodeIdIndexKey struct {
	vs []interface{}
}

func (x NodeIdIndexKey) id() uint32            { return 0 }
func (x NodeIdIndexKey) values() []interface{} { return x.vs }
func (x NodeIdIndexKey) nodeIndexKey()         {}

func (this NodeIdIndexKey) WithId(id uint64) NodeIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type NodeCuratorIndexKey struct {
	vs []interface{}
}

func (x NodeCuratorIndexKey) id() uint32            { return 1 }
func (x NodeCuratorIndexKey) values() []interface{} { return x.vs }
func (x NodeCuratorIndexKey) nodeIndexKey()         {}

func (this NodeCuratorIndexKey) WithCurator(curator []byte) NodeCuratorIndexKey {
	this.vs = []interface{}{curator}
	return this
}

type nodeTable struct {
	table ormtable.AutoIncrementTable
}

func (this nodeTable) Insert(ctx context.Context, node *Node) error {
	return this.table.Insert(ctx, node)
}

func (this nodeTable) Update(ctx context.Context, node *Node) error {
	return this.table.Update(ctx, node)
}

func (this nodeTable) Save(ctx context.Context, node *Node) error {
	return this.table.Save(ctx, node)
}

func (this nodeTable) Delete(ctx context.Context, node *Node) error {
	return this.table.Delete(ctx, node)
}

func (this nodeTable) InsertReturningId(ctx context.Context, node *Node) (uint64, error) {
	return this.table.InsertReturningPKey(ctx, node)
}

func (this nodeTable) LastInsertedSequence(ctx context.Context) (uint64, error) {
	return this.table.LastInsertedSequence(ctx)
}

func (this nodeTable) Has(ctx context.Context, id uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this nodeTable) Get(ctx context.Context, id uint64) (*Node, error) {
	var node Node
	found, err := this.table.PrimaryKey().Get(ctx, &node, id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &node, nil
}

func (this nodeTable) List(ctx context.Context, prefixKey NodeIndexKey, opts ...ormlist.Option) (NodeIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return NodeIterator{it}, err
}

func (this nodeTable) ListRange(ctx context.Context, from, to NodeIndexKey, opts ...ormlist.Option) (NodeIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return NodeIterator{it}, err
}

func (this nodeTable) DeleteBy(ctx context.Context, prefixKey NodeIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this nodeTable) DeleteRange(ctx context.Context, from, to NodeIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this nodeTable) doNotImplement() {}

var _ NodeTable = nodeTable{}

func NewNodeTable(db ormtable.Schema) (NodeTable, error) {
	table := db.GetTable(&Node{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Node{}).ProtoReflect().Descriptor().FullName()))
	}
	return nodeTable{table.(ormtable.AutoIncrementTable)}, nil
}

type StateStore interface {
	NodeTable() NodeTable

	doNotImplement()
}

type stateStore struct {
	node NodeTable
}

func (x stateStore) NodeTable() NodeTable {
	return x.node
}

func (stateStore) doNotImplement() {}

var _ StateStore = stateStore{}

func NewStateStore(db ormtable.Schema) (StateStore, error) {
	nodeTable, err := NewNodeTable(db)
	if err != nil {
		return nil, err
	}

	return stateStore{
		nodeTable,
	}, nil
}
