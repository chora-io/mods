// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: chora/geonode/v1/query.proto

package geonodev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Query_Node_FullMethodName           = "/chora.geonode.v1.Query/Node"
	Query_Nodes_FullMethodName          = "/chora.geonode.v1.Query/Nodes"
	Query_NodesByCurator_FullMethodName = "/chora.geonode.v1.Query/NodesByCurator"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Node queries a node by id.
	Node(ctx context.Context, in *QueryNodeRequest, opts ...grpc.CallOption) (*QueryNodeResponse, error)
	// Nodes queries all nodes.
	Nodes(ctx context.Context, in *QueryNodesRequest, opts ...grpc.CallOption) (*QueryNodesResponse, error)
	// NodesByCurator queries nodes by curator.
	NodesByCurator(ctx context.Context, in *QueryNodesByCuratorRequest, opts ...grpc.CallOption) (*QueryNodesByCuratorResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Node(ctx context.Context, in *QueryNodeRequest, opts ...grpc.CallOption) (*QueryNodeResponse, error) {
	out := new(QueryNodeResponse)
	err := c.cc.Invoke(ctx, Query_Node_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Nodes(ctx context.Context, in *QueryNodesRequest, opts ...grpc.CallOption) (*QueryNodesResponse, error) {
	out := new(QueryNodesResponse)
	err := c.cc.Invoke(ctx, Query_Nodes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) NodesByCurator(ctx context.Context, in *QueryNodesByCuratorRequest, opts ...grpc.CallOption) (*QueryNodesByCuratorResponse, error) {
	out := new(QueryNodesByCuratorResponse)
	err := c.cc.Invoke(ctx, Query_NodesByCurator_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Node queries a node by id.
	Node(context.Context, *QueryNodeRequest) (*QueryNodeResponse, error)
	// Nodes queries all nodes.
	Nodes(context.Context, *QueryNodesRequest) (*QueryNodesResponse, error)
	// NodesByCurator queries nodes by curator.
	NodesByCurator(context.Context, *QueryNodesByCuratorRequest) (*QueryNodesByCuratorResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Node(context.Context, *QueryNodeRequest) (*QueryNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Node not implemented")
}
func (UnimplementedQueryServer) Nodes(context.Context, *QueryNodesRequest) (*QueryNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Nodes not implemented")
}
func (UnimplementedQueryServer) NodesByCurator(context.Context, *QueryNodesByCuratorRequest) (*QueryNodesByCuratorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodesByCurator not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Node_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Node(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Node_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Node(ctx, req.(*QueryNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Nodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Nodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Nodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Nodes(ctx, req.(*QueryNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_NodesByCurator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryNodesByCuratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).NodesByCurator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_NodesByCurator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).NodesByCurator(ctx, req.(*QueryNodesByCuratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chora.geonode.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Node",
			Handler:    _Query_Node_Handler,
		},
		{
			MethodName: "Nodes",
			Handler:    _Query_Nodes_Handler,
		},
		{
			MethodName: "NodesByCurator",
			Handler:    _Query_NodesByCurator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chora/geonode/v1/query.proto",
}
