// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: chora/voucher/v1/query.proto

package voucherv1

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
	Query_Voucher_FullMethodName           = "/chora.voucher.v1.Query/Voucher"
	Query_Vouchers_FullMethodName          = "/chora.voucher.v1.Query/Vouchers"
	Query_VouchersByIssuer_FullMethodName  = "/chora.voucher.v1.Query/VouchersByIssuer"
	Query_Balance_FullMethodName           = "/chora.voucher.v1.Query/Balance"
	Query_BalancesByAddress_FullMethodName = "/chora.voucher.v1.Query/BalancesByAddress"
	Query_BalancesByVoucher_FullMethodName = "/chora.voucher.v1.Query/BalancesByVoucher"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Voucher queries a voucher by id.
	Voucher(ctx context.Context, in *QueryVoucherRequest, opts ...grpc.CallOption) (*QueryVoucherResponse, error)
	// Vouchers queries all vouchers.
	Vouchers(ctx context.Context, in *QueryVouchersRequest, opts ...grpc.CallOption) (*QueryVouchersResponse, error)
	// VouchersByIssuer queries vouchers by issuer.
	VouchersByIssuer(ctx context.Context, in *QueryVouchersByIssuerRequest, opts ...grpc.CallOption) (*QueryVouchersByIssuerResponse, error)
	// Balance queries the balance of a voucher and address.
	Balance(ctx context.Context, in *QueryBalanceRequest, opts ...grpc.CallOption) (*QueryBalanceResponse, error)
	// BalancesByAddress queries all balances of an address.
	BalancesByAddress(ctx context.Context, in *QueryBalancesByAddressRequest, opts ...grpc.CallOption) (*QueryBalancesByAddressResponse, error)
	// BalancesByVoucher queries all balances of a voucher.
	BalancesByVoucher(ctx context.Context, in *QueryBalancesByVoucherRequest, opts ...grpc.CallOption) (*QueryBalancesByVoucherResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Voucher(ctx context.Context, in *QueryVoucherRequest, opts ...grpc.CallOption) (*QueryVoucherResponse, error) {
	out := new(QueryVoucherResponse)
	err := c.cc.Invoke(ctx, Query_Voucher_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Vouchers(ctx context.Context, in *QueryVouchersRequest, opts ...grpc.CallOption) (*QueryVouchersResponse, error) {
	out := new(QueryVouchersResponse)
	err := c.cc.Invoke(ctx, Query_Vouchers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) VouchersByIssuer(ctx context.Context, in *QueryVouchersByIssuerRequest, opts ...grpc.CallOption) (*QueryVouchersByIssuerResponse, error) {
	out := new(QueryVouchersByIssuerResponse)
	err := c.cc.Invoke(ctx, Query_VouchersByIssuer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Balance(ctx context.Context, in *QueryBalanceRequest, opts ...grpc.CallOption) (*QueryBalanceResponse, error) {
	out := new(QueryBalanceResponse)
	err := c.cc.Invoke(ctx, Query_Balance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BalancesByAddress(ctx context.Context, in *QueryBalancesByAddressRequest, opts ...grpc.CallOption) (*QueryBalancesByAddressResponse, error) {
	out := new(QueryBalancesByAddressResponse)
	err := c.cc.Invoke(ctx, Query_BalancesByAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BalancesByVoucher(ctx context.Context, in *QueryBalancesByVoucherRequest, opts ...grpc.CallOption) (*QueryBalancesByVoucherResponse, error) {
	out := new(QueryBalancesByVoucherResponse)
	err := c.cc.Invoke(ctx, Query_BalancesByVoucher_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Voucher queries a voucher by id.
	Voucher(context.Context, *QueryVoucherRequest) (*QueryVoucherResponse, error)
	// Vouchers queries all vouchers.
	Vouchers(context.Context, *QueryVouchersRequest) (*QueryVouchersResponse, error)
	// VouchersByIssuer queries vouchers by issuer.
	VouchersByIssuer(context.Context, *QueryVouchersByIssuerRequest) (*QueryVouchersByIssuerResponse, error)
	// Balance queries the balance of a voucher and address.
	Balance(context.Context, *QueryBalanceRequest) (*QueryBalanceResponse, error)
	// BalancesByAddress queries all balances of an address.
	BalancesByAddress(context.Context, *QueryBalancesByAddressRequest) (*QueryBalancesByAddressResponse, error)
	// BalancesByVoucher queries all balances of a voucher.
	BalancesByVoucher(context.Context, *QueryBalancesByVoucherRequest) (*QueryBalancesByVoucherResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Voucher(context.Context, *QueryVoucherRequest) (*QueryVoucherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Voucher not implemented")
}
func (UnimplementedQueryServer) Vouchers(context.Context, *QueryVouchersRequest) (*QueryVouchersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Vouchers not implemented")
}
func (UnimplementedQueryServer) VouchersByIssuer(context.Context, *QueryVouchersByIssuerRequest) (*QueryVouchersByIssuerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VouchersByIssuer not implemented")
}
func (UnimplementedQueryServer) Balance(context.Context, *QueryBalanceRequest) (*QueryBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balance not implemented")
}
func (UnimplementedQueryServer) BalancesByAddress(context.Context, *QueryBalancesByAddressRequest) (*QueryBalancesByAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BalancesByAddress not implemented")
}
func (UnimplementedQueryServer) BalancesByVoucher(context.Context, *QueryBalancesByVoucherRequest) (*QueryBalancesByVoucherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BalancesByVoucher not implemented")
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

func _Query_Voucher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryVoucherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Voucher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Voucher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Voucher(ctx, req.(*QueryVoucherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Vouchers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryVouchersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Vouchers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Vouchers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Vouchers(ctx, req.(*QueryVouchersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_VouchersByIssuer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryVouchersByIssuerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).VouchersByIssuer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_VouchersByIssuer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).VouchersByIssuer(ctx, req.(*QueryVouchersByIssuerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Balance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Balance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Balance(ctx, req.(*QueryBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BalancesByAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBalancesByAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BalancesByAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_BalancesByAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BalancesByAddress(ctx, req.(*QueryBalancesByAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BalancesByVoucher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBalancesByVoucherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BalancesByVoucher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_BalancesByVoucher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BalancesByVoucher(ctx, req.(*QueryBalancesByVoucherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chora.voucher.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Voucher",
			Handler:    _Query_Voucher_Handler,
		},
		{
			MethodName: "Vouchers",
			Handler:    _Query_Vouchers_Handler,
		},
		{
			MethodName: "VouchersByIssuer",
			Handler:    _Query_VouchersByIssuer_Handler,
		},
		{
			MethodName: "Balance",
			Handler:    _Query_Balance_Handler,
		},
		{
			MethodName: "BalancesByAddress",
			Handler:    _Query_BalancesByAddress_Handler,
		},
		{
			MethodName: "BalancesByVoucher",
			Handler:    _Query_BalancesByVoucher_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chora/voucher/v1/query.proto",
}
