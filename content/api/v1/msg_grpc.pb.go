// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: chora/content/v1/msg.proto

package contentv1

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
	Msg_Create_FullMethodName         = "/chora.content.v1.Msg/Create"
	Msg_Delete_FullMethodName         = "/chora.content.v1.Msg/Delete"
	Msg_UpdateCurator_FullMethodName  = "/chora.content.v1.Msg/UpdateCurator"
	Msg_UpdateMetadata_FullMethodName = "/chora.content.v1.Msg/UpdateMetadata"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// Create creates content.
	Create(ctx context.Context, in *MsgCreate, opts ...grpc.CallOption) (*MsgCreateResponse, error)
	// Delete deletes content.
	Delete(ctx context.Context, in *MsgDelete, opts ...grpc.CallOption) (*MsgDeleteResponse, error)
	// UpdateCurator updates the curator of content.
	UpdateCurator(ctx context.Context, in *MsgUpdateCurator, opts ...grpc.CallOption) (*MsgUpdateCuratorResponse, error)
	// UpdateMetadata updates the metadata of content.
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

func (c *msgClient) Delete(ctx context.Context, in *MsgDelete, opts ...grpc.CallOption) (*MsgDeleteResponse, error) {
	out := new(MsgDeleteResponse)
	err := c.cc.Invoke(ctx, Msg_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateCurator(ctx context.Context, in *MsgUpdateCurator, opts ...grpc.CallOption) (*MsgUpdateCuratorResponse, error) {
	out := new(MsgUpdateCuratorResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateCurator_FullMethodName, in, out, opts...)
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
	// Create creates content.
	Create(context.Context, *MsgCreate) (*MsgCreateResponse, error)
	// Delete deletes content.
	Delete(context.Context, *MsgDelete) (*MsgDeleteResponse, error)
	// UpdateCurator updates the curator of content.
	UpdateCurator(context.Context, *MsgUpdateCurator) (*MsgUpdateCuratorResponse, error)
	// UpdateMetadata updates the metadata of content.
	UpdateMetadata(context.Context, *MsgUpdateMetadata) (*MsgUpdateMetadataResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) Create(context.Context, *MsgCreate) (*MsgCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedMsgServer) Delete(context.Context, *MsgDelete) (*MsgDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedMsgServer) UpdateCurator(context.Context, *MsgUpdateCurator) (*MsgUpdateCuratorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCurator not implemented")
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

func _Msg_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDelete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Delete(ctx, req.(*MsgDelete))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateCurator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateCurator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateCurator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateCurator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateCurator(ctx, req.(*MsgUpdateCurator))
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
	ServiceName: "chora.content.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Msg_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Msg_Delete_Handler,
		},
		{
			MethodName: "UpdateCurator",
			Handler:    _Msg_UpdateCurator_Handler,
		},
		{
			MethodName: "UpdateMetadata",
			Handler:    _Msg_UpdateMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chora/content/v1/msg.proto",
}
