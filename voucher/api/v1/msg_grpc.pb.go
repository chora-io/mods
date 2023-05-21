// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: chora/voucher/v1/msg.proto

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
	Msg_Create_FullMethodName         = "/chora.voucher.v1.Msg/Create"
	Msg_Issue_FullMethodName          = "/chora.voucher.v1.Msg/Issue"
	Msg_UpdateIssuer_FullMethodName   = "/chora.voucher.v1.Msg/UpdateIssuer"
	Msg_UpdateMetadata_FullMethodName = "/chora.voucher.v1.Msg/UpdateMetadata"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// Create creates a voucher.
	Create(ctx context.Context, in *MsgCreate, opts ...grpc.CallOption) (*MsgCreateResponse, error)
	// Issue issues vouchers to a recipient.
	Issue(ctx context.Context, in *MsgIssue, opts ...grpc.CallOption) (*MsgIssueResponse, error)
	// UpdateIssuer updates the issuer of a voucher.
	UpdateIssuer(ctx context.Context, in *MsgUpdateIssuer, opts ...grpc.CallOption) (*MsgUpdateIssuerResponse, error)
	// UpdateMetadata updates the metadata of a voucher.
	UpdateMetadata(ctx context.Context, in *MsgUpdateMetadata, opts ...grpc.CallOption) (*MsgUpdateMetadataResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Create(ctx context.Context, in *MsgCreate, opts ...grpc.CallOption) (*MsgCreateResponse, error) {
	out := new(MsgCreateResponse)
	err := c.cc.Invoke(ctx, Msg_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Issue(ctx context.Context, in *MsgIssue, opts ...grpc.CallOption) (*MsgIssueResponse, error) {
	out := new(MsgIssueResponse)
	err := c.cc.Invoke(ctx, Msg_Issue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateIssuer(ctx context.Context, in *MsgUpdateIssuer, opts ...grpc.CallOption) (*MsgUpdateIssuerResponse, error) {
	out := new(MsgUpdateIssuerResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateIssuer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateMetadata(ctx context.Context, in *MsgUpdateMetadata, opts ...grpc.CallOption) (*MsgUpdateMetadataResponse, error) {
	out := new(MsgUpdateMetadataResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateMetadata_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// Create creates a voucher.
	Create(context.Context, *MsgCreate) (*MsgCreateResponse, error)
	// Issue issues vouchers to a recipient.
	Issue(context.Context, *MsgIssue) (*MsgIssueResponse, error)
	// UpdateIssuer updates the issuer of a voucher.
	UpdateIssuer(context.Context, *MsgUpdateIssuer) (*MsgUpdateIssuerResponse, error)
	// UpdateMetadata updates the metadata of a voucher.
	UpdateMetadata(context.Context, *MsgUpdateMetadata) (*MsgUpdateMetadataResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) Create(context.Context, *MsgCreate) (*MsgCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedMsgServer) Issue(context.Context, *MsgIssue) (*MsgIssueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Issue not implemented")
}
func (UnimplementedMsgServer) UpdateIssuer(context.Context, *MsgUpdateIssuer) (*MsgUpdateIssuerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIssuer not implemented")
}
func (UnimplementedMsgServer) UpdateMetadata(context.Context, *MsgUpdateMetadata) (*MsgUpdateMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMetadata not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Create(ctx, req.(*MsgCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Issue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgIssue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Issue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Issue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Issue(ctx, req.(*MsgIssue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateIssuer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateIssuer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateIssuer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateIssuer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateIssuer(ctx, req.(*MsgUpdateIssuer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateMetadata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateMetadata(ctx, req.(*MsgUpdateMetadata))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chora.voucher.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Msg_Create_Handler,
		},
		{
			MethodName: "Issue",
			Handler:    _Msg_Issue_Handler,
		},
		{
			MethodName: "UpdateIssuer",
			Handler:    _Msg_UpdateIssuer_Handler,
		},
		{
			MethodName: "UpdateMetadata",
			Handler:    _Msg_UpdateMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chora/voucher/v1/msg.proto",
}
