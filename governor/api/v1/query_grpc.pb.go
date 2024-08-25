// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: chora/governor/v1/query.proto

package governorv1

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
	Query_Governor_FullMethodName  = "/chora.governor.v1.Query/Governor"
	Query_Governors_FullMethodName = "/chora.governor.v1.Query/Governors"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Query is the Query service.
type QueryClient interface {
	// Governor queries a governor by address.
	Governor(ctx context.Context, in *QueryGovernorRequest, opts ...grpc.CallOption) (*QueryGovernorResponse, error)
	// Governors queries all governors.
	Governors(ctx context.Context, in *QueryGovernorsRequest, opts ...grpc.CallOption) (*QueryGovernorsResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Governor(ctx context.Context, in *QueryGovernorRequest, opts ...grpc.CallOption) (*QueryGovernorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryGovernorResponse)
	err := c.cc.Invoke(ctx, Query_Governor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Governors(ctx context.Context, in *QueryGovernorsRequest, opts ...grpc.CallOption) (*QueryGovernorsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryGovernorsResponse)
	err := c.cc.Invoke(ctx, Query_Governors_FullMethodName, in, out, cOpts...)
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
	// Governor queries a governor by address.
	Governor(context.Context, *QueryGovernorRequest) (*QueryGovernorResponse, error)
	// Governors queries all governors.
	Governors(context.Context, *QueryGovernorsRequest) (*QueryGovernorsResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedQueryServer struct{}

func (UnimplementedQueryServer) Governor(context.Context, *QueryGovernorRequest) (*QueryGovernorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Governor not implemented")
}
func (UnimplementedQueryServer) Governors(context.Context, *QueryGovernorsRequest) (*QueryGovernorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Governors not implemented")
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

func _Query_Governor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGovernorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Governor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Governor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Governor(ctx, req.(*QueryGovernorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Governors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGovernorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Governors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Governors_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Governors(ctx, req.(*QueryGovernorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chora.governor.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Governor",
			Handler:    _Query_Governor_Handler,
		},
		{
			MethodName: "Governors",
			Handler:    _Query_Governors_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chora/governor/v1/query.proto",
}
