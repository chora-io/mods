// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: v1/query.proto

package v1

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryContentRequest is the Query/Content request type.
type QueryContentRequest struct {
	// id is the unique identifier of the content.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryContentRequest) Reset()         { *m = QueryContentRequest{} }
func (m *QueryContentRequest) String() string { return proto.CompactTextString(m) }
func (*QueryContentRequest) ProtoMessage()    {}
func (*QueryContentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0801432bccbe1b86, []int{0}
}
func (m *QueryContentRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryContentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContentRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryContentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContentRequest.Merge(m, src)
}
func (m *QueryContentRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryContentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContentRequest proto.InternalMessageInfo

func (m *QueryContentRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// QueryContentResponse is the Query/Content response type.
type QueryContentResponse struct {
	// id is the unique identifier of the content.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// creator is the creator of the content.
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	// hash is the content hash of the content.
	Hash string `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *QueryContentResponse) Reset()         { *m = QueryContentResponse{} }
func (m *QueryContentResponse) String() string { return proto.CompactTextString(m) }
func (*QueryContentResponse) ProtoMessage()    {}
func (*QueryContentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0801432bccbe1b86, []int{1}
}
func (m *QueryContentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryContentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryContentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContentResponse.Merge(m, src)
}
func (m *QueryContentResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryContentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContentResponse proto.InternalMessageInfo

func (m *QueryContentResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *QueryContentResponse) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *QueryContentResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

// QueryContentByCreatorRequest is the Query/ContentByCreator request type.
type QueryContentByCreatorRequest struct {
	// creator is the address of the content creator.
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// pagination is the optional pagination in the request.
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryContentByCreatorRequest) Reset()         { *m = QueryContentByCreatorRequest{} }
func (m *QueryContentByCreatorRequest) String() string { return proto.CompactTextString(m) }
func (*QueryContentByCreatorRequest) ProtoMessage()    {}
func (*QueryContentByCreatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0801432bccbe1b86, []int{2}
}
func (m *QueryContentByCreatorRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryContentByCreatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContentByCreatorRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryContentByCreatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContentByCreatorRequest.Merge(m, src)
}
func (m *QueryContentByCreatorRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryContentByCreatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContentByCreatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContentByCreatorRequest proto.InternalMessageInfo

func (m *QueryContentByCreatorRequest) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *QueryContentByCreatorRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryContentByCreatorResponse is the Query/ContentByCreator response type.
type QueryContentByCreatorResponse struct {
	// creator is the address of the content creator.
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// content is the content created by the creator.
	Content []*QueryContentByCreatorResponse_Content `protobuf:"bytes,2,rep,name=content,proto3" json:"content,omitempty"`
	// pagination is the pagination in the response.
	Pagination *query.PageResponse `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryContentByCreatorResponse) Reset()         { *m = QueryContentByCreatorResponse{} }
func (m *QueryContentByCreatorResponse) String() string { return proto.CompactTextString(m) }
func (*QueryContentByCreatorResponse) ProtoMessage()    {}
func (*QueryContentByCreatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0801432bccbe1b86, []int{3}
}
func (m *QueryContentByCreatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryContentByCreatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContentByCreatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryContentByCreatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContentByCreatorResponse.Merge(m, src)
}
func (m *QueryContentByCreatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryContentByCreatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContentByCreatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContentByCreatorResponse proto.InternalMessageInfo

func (m *QueryContentByCreatorResponse) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *QueryContentByCreatorResponse) GetContent() []*QueryContentByCreatorResponse_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *QueryContentByCreatorResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// Content is the content properties.
type QueryContentByCreatorResponse_Content struct {
	// id is the unique identifier of the content.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// hash is the content hash of the content.
	Hash string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *QueryContentByCreatorResponse_Content) Reset()         { *m = QueryContentByCreatorResponse_Content{} }
func (m *QueryContentByCreatorResponse_Content) String() string { return proto.CompactTextString(m) }
func (*QueryContentByCreatorResponse_Content) ProtoMessage()    {}
func (*QueryContentByCreatorResponse_Content) Descriptor() ([]byte, []int) {
	return fileDescriptor_0801432bccbe1b86, []int{3, 0}
}
func (m *QueryContentByCreatorResponse_Content) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryContentByCreatorResponse_Content) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContentByCreatorResponse_Content.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryContentByCreatorResponse_Content) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContentByCreatorResponse_Content.Merge(m, src)
}
func (m *QueryContentByCreatorResponse_Content) XXX_Size() int {
	return m.Size()
}
func (m *QueryContentByCreatorResponse_Content) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContentByCreatorResponse_Content.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContentByCreatorResponse_Content proto.InternalMessageInfo

func (m *QueryContentByCreatorResponse_Content) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *QueryContentByCreatorResponse_Content) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryContentRequest)(nil), "chora.example.v1.QueryContentRequest")
	proto.RegisterType((*QueryContentResponse)(nil), "chora.example.v1.QueryContentResponse")
	proto.RegisterType((*QueryContentByCreatorRequest)(nil), "chora.example.v1.QueryContentByCreatorRequest")
	proto.RegisterType((*QueryContentByCreatorResponse)(nil), "chora.example.v1.QueryContentByCreatorResponse")
	proto.RegisterType((*QueryContentByCreatorResponse_Content)(nil), "chora.example.v1.QueryContentByCreatorResponse.Content")
}

func init() { proto.RegisterFile("v1/query.proto", fileDescriptor_0801432bccbe1b86) }

var fileDescriptor_0801432bccbe1b86 = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcd, 0x8a, 0x13, 0x31,
	0x00, 0x6e, 0xd2, 0xd5, 0x62, 0x16, 0x96, 0x25, 0x7a, 0x18, 0xca, 0x3a, 0x94, 0x81, 0xad, 0x45,
	0x68, 0xe2, 0x54, 0xd0, 0x7b, 0x17, 0xf4, 0xea, 0x16, 0x4f, 0xde, 0xd2, 0x69, 0x98, 0x09, 0xec,
	0x4c, 0x66, 0x27, 0xe9, 0xe0, 0xb0, 0x2c, 0x2c, 0xe2, 0x03, 0x08, 0x3e, 0x86, 0x2f, 0xe0, 0x23,
	0x78, 0x5c, 0xf0, 0xe2, 0x51, 0x5a, 0x1f, 0x44, 0x9a, 0x64, 0xec, 0xb4, 0xd2, 0xad, 0xde, 0xd2,
	0xe6, 0xfb, 0xbe, 0xf9, 0x7e, 0x08, 0x3a, 0x2a, 0x43, 0x7a, 0x39, 0xe7, 0x45, 0x45, 0xf2, 0x42,
	0x6a, 0x89, 0x8f, 0xa3, 0x44, 0x16, 0x8c, 0xf0, 0xf7, 0x2c, 0xcd, 0x2f, 0x38, 0x29, 0xc3, 0xee,
	0xd3, 0x48, 0xaa, 0x54, 0x2a, 0x3a, 0x65, 0x8a, 0x5b, 0x28, 0x2d, 0xc3, 0x29, 0xd7, 0x2c, 0xa4,
	0x39, 0x8b, 0x45, 0xc6, 0xb4, 0x90, 0x99, 0x65, 0x77, 0x4f, 0x62, 0x29, 0xe3, 0x0b, 0x4e, 0x59,
	0x2e, 0x28, 0xcb, 0x32, 0xa9, 0xcd, 0xa5, 0xb2, 0xb7, 0xc1, 0x29, 0x7a, 0x78, 0xbe, 0xe2, 0x9f,
	0xc9, 0x4c, 0xf3, 0x4c, 0x4f, 0xf8, 0xe5, 0x9c, 0x2b, 0x8d, 0x8f, 0x10, 0x14, 0x33, 0x0f, 0xf4,
	0xc0, 0xe0, 0x60, 0x02, 0xc5, 0x2c, 0x78, 0x8b, 0x1e, 0x6d, 0xc2, 0x54, 0x2e, 0x33, 0xc5, 0xb7,
	0x71, 0xd8, 0x43, 0x9d, 0xa8, 0xe0, 0x4c, 0xcb, 0xc2, 0x83, 0x3d, 0x30, 0x78, 0x30, 0xa9, 0x7f,
	0x62, 0x8c, 0x0e, 0x12, 0xa6, 0x12, 0xaf, 0x6d, 0xfe, 0x36, 0xe7, 0xe0, 0x06, 0xa0, 0x93, 0xa6,
	0xec, 0xb8, 0x3a, 0xb3, 0xe8, 0xda, 0x46, 0x43, 0x0e, 0x6c, 0xca, 0xbd, 0x42, 0x68, 0x9d, 0xd4,
	0x7c, 0xeb, 0x70, 0xd4, 0x27, 0xb6, 0x16, 0xb2, 0xaa, 0x85, 0xd8, 0x06, 0x5d, 0x2d, 0xe4, 0x0d,
	0x8b, 0xb9, 0x53, 0x9d, 0x34, 0x98, 0xc1, 0x47, 0x88, 0x1e, 0xef, 0xb0, 0xe0, 0x22, 0xee, 0xf6,
	0x70, 0x8e, 0x3a, 0x91, 0x65, 0x79, 0xb0, 0xd7, 0x1e, 0x1c, 0x8e, 0x5e, 0x92, 0xed, 0xa5, 0xc8,
	0x9d, 0xda, 0xa4, 0xae, 0xb3, 0xd6, 0xc1, 0xaf, 0x37, 0x62, 0xb5, 0x4d, 0xac, 0x27, 0x7b, 0x63,
	0x59, 0xb5, 0x66, 0xae, 0xee, 0x10, 0x75, 0x9c, 0xf8, 0x5f, 0x1b, 0xd5, 0x4b, 0xc0, 0xf5, 0x12,
	0xa3, 0xaf, 0x10, 0xdd, 0x33, 0x56, 0xf1, 0x0d, 0x58, 0x33, 0x4f, 0xef, 0xce, 0xe3, 0xfa, 0xec,
	0xf6, 0xf7, 0xc1, 0xac, 0xbf, 0xa0, 0xff, 0xe1, 0xfb, 0xaf, 0xcf, 0xb0, 0x87, 0x7d, 0x6a, 0xf0,
	0xd4, 0xe1, 0x69, 0x19, 0x52, 0x97, 0x9f, 0x5e, 0x89, 0xd9, 0x35, 0xfe, 0x02, 0xd0, 0xf1, 0x76,
	0x65, 0x98, 0xfc, 0x73, 0xb7, 0xd6, 0x14, 0xfd, 0xcf, 0x2d, 0x82, 0x17, 0xc6, 0xdd, 0x33, 0x4c,
	0x76, 0xba, 0x1b, 0x4e, 0xab, 0xa1, 0xdb, 0x9e, 0x5e, 0xb9, 0xc3, 0xf5, 0x78, 0xfc, 0x6d, 0xe1,
	0x83, 0xdb, 0x85, 0x0f, 0x7e, 0x2e, 0x7c, 0xf0, 0x69, 0xe9, 0xb7, 0x6e, 0x97, 0x7e, 0xeb, 0xc7,
	0xd2, 0x6f, 0xbd, 0x1b, 0xc4, 0x42, 0x27, 0xf3, 0x29, 0x89, 0x64, 0x6a, 0x35, 0x85, 0xa4, 0xa9,
	0x9c, 0xa9, 0x3f, 0xd2, 0xba, 0xca, 0xb9, 0x5a, 0xbd, 0xdc, 0xfb, 0xe6, 0x31, 0x3e, 0xff, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0x17, 0xc1, 0x5a, 0xee, 0xfa, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Content queries content by the unique identifier of the content.
	Content(ctx context.Context, in *QueryContentRequest, opts ...grpc.CallOption) (*QueryContentResponse, error)
	// ContentByCreator queries all content by the creator of the content.
	ContentByCreator(ctx context.Context, in *QueryContentByCreatorRequest, opts ...grpc.CallOption) (*QueryContentByCreatorResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Content(ctx context.Context, in *QueryContentRequest, opts ...grpc.CallOption) (*QueryContentResponse, error) {
	out := new(QueryContentResponse)
	err := c.cc.Invoke(ctx, "/chora.example.v1.Query/Content", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ContentByCreator(ctx context.Context, in *QueryContentByCreatorRequest, opts ...grpc.CallOption) (*QueryContentByCreatorResponse, error) {
	out := new(QueryContentByCreatorResponse)
	err := c.cc.Invoke(ctx, "/chora.example.v1.Query/ContentByCreator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Content queries content by the unique identifier of the content.
	Content(context.Context, *QueryContentRequest) (*QueryContentResponse, error)
	// ContentByCreator queries all content by the creator of the content.
	ContentByCreator(context.Context, *QueryContentByCreatorRequest) (*QueryContentByCreatorResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Content(ctx context.Context, req *QueryContentRequest) (*QueryContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Content not implemented")
}
func (*UnimplementedQueryServer) ContentByCreator(ctx context.Context, req *QueryContentByCreatorRequest) (*QueryContentByCreatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContentByCreator not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
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
		FullMethod: "/chora.example.v1.Query/Content",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Content(ctx, req.(*QueryContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ContentByCreator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryContentByCreatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ContentByCreator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chora.example.v1.Query/ContentByCreator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ContentByCreator(ctx, req.(*QueryContentByCreatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chora.example.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Content",
			Handler:    _Query_Content_Handler,
		},
		{
			MethodName: "ContentByCreator",
			Handler:    _Query_ContentByCreator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/query.proto",
}

func (m *QueryContentRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContentRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContentRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryContentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryContentByCreatorRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContentByCreatorRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContentByCreatorRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryContentByCreatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContentByCreatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContentByCreatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Content) > 0 {
		for iNdEx := len(m.Content) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Content[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryContentByCreatorResponse_Content) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContentByCreatorResponse_Content) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContentByCreatorResponse_Content) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryContentRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	return n
}

func (m *QueryContentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryContentByCreatorRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryContentByCreatorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if len(m.Content) > 0 {
		for _, e := range m.Content {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryContentByCreatorResponse_Content) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryContentRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryContentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryContentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryContentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryContentByCreatorRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryContentByCreatorRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContentByCreatorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryContentByCreatorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryContentByCreatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContentByCreatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = append(m.Content, &QueryContentByCreatorResponse_Content{})
			if err := m.Content[len(m.Content)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryContentByCreatorResponse_Content) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Content: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Content: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
