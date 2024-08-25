// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chora/admin/v1/msg.proto

package v1

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// MsgUpdate is the Msg/Update request type.
type MsgUpdate struct {
	// admin is the address of the admin account.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// new_admin is the address of the new admin account.
	NewAdmin string `protobuf:"bytes,2,opt,name=new_admin,json=newAdmin,proto3" json:"new_admin,omitempty"`
}

func (m *MsgUpdate) Reset()         { *m = MsgUpdate{} }
func (m *MsgUpdate) String() string { return proto.CompactTextString(m) }
func (*MsgUpdate) ProtoMessage()    {}
func (*MsgUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9548170b8495a83, []int{0}
}
func (m *MsgUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdate.Merge(m, src)
}
func (m *MsgUpdate) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdate proto.InternalMessageInfo

func (m *MsgUpdate) GetAdmin() string {
	if m != nil {
		return m.Admin
	}
	return ""
}

func (m *MsgUpdate) GetNewAdmin() string {
	if m != nil {
		return m.NewAdmin
	}
	return ""
}

// MsgUpdateResponse is the Msg/Update response type.
type MsgUpdateResponse struct {
	// admin is the address of the admin account.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// new_admin is the address of the new admin account.
	NewAdmin string `protobuf:"bytes,2,opt,name=new_admin,json=newAdmin,proto3" json:"new_admin,omitempty"`
}

func (m *MsgUpdateResponse) Reset()         { *m = MsgUpdateResponse{} }
func (m *MsgUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateResponse) ProtoMessage()    {}
func (*MsgUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9548170b8495a83, []int{1}
}
func (m *MsgUpdateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateResponse.Merge(m, src)
}
func (m *MsgUpdateResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateResponse proto.InternalMessageInfo

func (m *MsgUpdateResponse) GetAdmin() string {
	if m != nil {
		return m.Admin
	}
	return ""
}

func (m *MsgUpdateResponse) GetNewAdmin() string {
	if m != nil {
		return m.NewAdmin
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgUpdate)(nil), "chora.admin.v1.MsgUpdate")
	proto.RegisterType((*MsgUpdateResponse)(nil), "chora.admin.v1.MsgUpdateResponse")
}

func init() { proto.RegisterFile("chora/admin/v1/msg.proto", fileDescriptor_e9548170b8495a83) }

var fileDescriptor_e9548170b8495a83 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xce, 0xc8, 0x2f,
	0x4a, 0xd4, 0x4f, 0x4c, 0xc9, 0xcd, 0xcc, 0xd3, 0x2f, 0x33, 0xd4, 0xcf, 0x2d, 0x4e, 0xd7, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x03, 0xcb, 0xe8, 0x81, 0x65, 0xf4, 0xca, 0x0c, 0xa5, 0xc4,
	0x93, 0xf3, 0x8b, 0x73, 0xf3, 0x8b, 0x41, 0x2a, 0x50, 0x14, 0x2a, 0x79, 0x71, 0x71, 0xfa, 0x16,
	0xa7, 0x87, 0x16, 0xa4, 0x24, 0x96, 0xa4, 0x0a, 0x89, 0x70, 0xb1, 0x82, 0x75, 0x48, 0x30, 0x2a,
	0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x42, 0xd2, 0x5c, 0x9c, 0x79, 0xa9, 0xe5, 0xf1, 0x10, 0x19,
	0x26, 0xb0, 0x0c, 0x47, 0x5e, 0x6a, 0xb9, 0x23, 0x88, 0x6f, 0xc5, 0xd5, 0xf4, 0x7c, 0x83, 0x16,
	0x44, 0xa1, 0x92, 0x1b, 0x97, 0x20, 0xdc, 0xac, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0x62, 0x72,
	0xcc, 0x34, 0xf2, 0xe5, 0x62, 0xf6, 0x2d, 0x4e, 0x17, 0x72, 0xe3, 0x62, 0x83, 0xba, 0x4b, 0x52,
	0x0f, 0xd5, 0x3b, 0x7a, 0x70, 0x6b, 0xa4, 0x14, 0x71, 0x4a, 0xc1, 0x5c, 0xe0, 0xe4, 0x78, 0xe2,
	0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70,
	0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xea, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49,
	0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x60, 0x63, 0x74, 0x33, 0xf3, 0xf5, 0x73, 0xf3, 0x53, 0x8a, 0xa1,
	0x41, 0x5a, 0x52, 0x59, 0x90, 0x5a, 0xac, 0x5f, 0x66, 0x98, 0xc4, 0x06, 0x0e, 0x2c, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xa2, 0xe6, 0xba, 0x4e, 0x71, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// Update updates the admin account.
	Update(ctx context.Context, in *MsgUpdate, opts ...grpc.CallOption) (*MsgUpdateResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Update(ctx context.Context, in *MsgUpdate, opts ...grpc.CallOption) (*MsgUpdateResponse, error) {
	out := new(MsgUpdateResponse)
	err := c.cc.Invoke(ctx, "/chora.admin.v1.Msg/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Update updates the admin account.
	Update(context.Context, *MsgUpdate) (*MsgUpdateResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Update(ctx context.Context, req *MsgUpdate) (*MsgUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chora.admin.v1.Msg/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Update(ctx, req.(*MsgUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chora.admin.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _Msg_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chora/admin/v1/msg.proto",
}

func (m *MsgUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NewAdmin) > 0 {
		i -= len(m.NewAdmin)
		copy(dAtA[i:], m.NewAdmin)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.NewAdmin)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Admin) > 0 {
		i -= len(m.Admin)
		copy(dAtA[i:], m.Admin)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Admin)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NewAdmin) > 0 {
		i -= len(m.NewAdmin)
		copy(dAtA[i:], m.NewAdmin)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.NewAdmin)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Admin) > 0 {
		i -= len(m.Admin)
		copy(dAtA[i:], m.Admin)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Admin)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Admin)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = len(m.NewAdmin)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	return n
}

func (m *MsgUpdateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Admin)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = len(m.NewAdmin)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	return n
}

func sovMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsg(x uint64) (n int) {
	return sovMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Admin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewAdmin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewAdmin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *MsgUpdateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgUpdateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Admin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewAdmin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewAdmin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func skipMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
				return 0, ErrInvalidLengthMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsg = fmt.Errorf("proto: unexpected end of group")
)