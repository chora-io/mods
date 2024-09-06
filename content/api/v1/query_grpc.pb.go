// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: chora/content/v1/query.proto

package contentv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Query_Content_FullMethodName           = "/chora.content.v1.Query/Content"
	Query_Contents_FullMethodName          = "/chora.content.v1.Query/Contents"
	Query_ContentsByCurator_FullMethodName = "/chora.content.v1.Query/ContentsByCurator"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Query is the Query service.
type QueryClient interface {
	// Content queries content by hash.
	Content(ctx context.Context, in *QueryContentRequest, opts ...grpc.CallOption) (*QueryContentResponse, error)
	// Contents queries all content.
	Contents(ctx context.Context, in *QueryContentsRequest, opts ...grpc.CallOption) (*QueryContentsResponse, error)
	// ContentsByCurator queries content by curator.
	ContentsByCurator(ctx context.Context, in *QueryContentsByCuratorRequest, opts ...grpc.CallOption) (*QueryContentsByCuratorResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Content(ctx context.Context, in *QueryContentRequest, opts ...grpc.CallOption) (*QueryContentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryContentResponse)
	err := c.cc.Invoke(ctx, Query_Content_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Contents(ctx context.Context, in *QueryContentsRequest, opts ...grpc.CallOption) (*QueryContentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryContentsResponse)
	err := c.cc.Invoke(ctx, Query_Contents_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ContentsByCurator(ctx context.Context, in *QueryContentsByCuratorRequest, opts ...grpc.CallOption) (*QueryContentsByCuratorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryContentsByCuratorResponse)
	err := c.cc.Invoke(ctx, Query_ContentsByCurator_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility.
//
// Query is the Query service.
type QueryServer interface {
	// Content queries content by hash.
	Content(context.Context, *QueryContentRequest) (*QueryContentResponse, error)
	// Contents queries all content.
	Contents(context.Context, *QueryContentsRequest) (*QueryContentsResponse, error)
	// ContentsByCurator queries content by curator.
	ContentsByCurator(context.Context, *QueryContentsByCuratorRequest) (*QueryContentsByCuratorResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedQueryServer struct{}

func (UnimplementedQueryServer) Content(context.Context, *QueryContentRequest) (*QueryContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Content not implemented")
}
func (UnimplementedQueryServer) Contents(context.Context, *QueryContentsRequest) (*QueryContentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Contents not implemented")
}
func (UnimplementedQueryServer) ContentsByCurator(context.Context, *QueryContentsByCuratorRequest) (*QueryContentsByCuratorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContentsByCurator not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}
func (UnimplementedQueryServer) testEmbeddedByValue()               {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	// If the following call pancis, it indicates UnimplementedQueryServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Content_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Content(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Content_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Content(ctx, req.(*QueryContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Contents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryContentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Contents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Contents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Contents(ctx, req.(*QueryContentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ContentsByCurator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryContentsByCuratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ContentsByCurator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ContentsByCurator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ContentsByCurator(ctx, req.(*QueryContentsByCuratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chora.content.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Content",
			Handler:    _Query_Content_Handler,
		},
		{
			MethodName: "Contents",
			Handler:    _Query_Contents_Handler,
		},
		{
			MethodName: "ContentsByCurator",
			Handler:    _Query_ContentsByCurator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chora/content/v1/query.proto",
}
