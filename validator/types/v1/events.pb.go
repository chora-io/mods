// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chora/validator/v1/events.proto

package v1

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

// EventAddValidator is an event emitted when a validator is added.
type EventAddValidator struct {
	// address is the address of the validator.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *EventAddValidator) Reset()         { *m = EventAddValidator{} }
func (m *EventAddValidator) String() string { return proto.CompactTextString(m) }
func (*EventAddValidator) ProtoMessage()    {}
func (*EventAddValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_7944b2035d0f825b, []int{0}
}
func (m *EventAddValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventAddValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventAddValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventAddValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAddValidator.Merge(m, src)
}
func (m *EventAddValidator) XXX_Size() int {
	return m.Size()
}
func (m *EventAddValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAddValidator.DiscardUnknown(m)
}

var xxx_messageInfo_EventAddValidator proto.InternalMessageInfo

func (m *EventAddValidator) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// EventMissedBlock is an event emitted when a validator missed a block.
type EventMissedBlock struct {
	// address is the address of the validator.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *EventMissedBlock) Reset()         { *m = EventMissedBlock{} }
func (m *EventMissedBlock) String() string { return proto.CompactTextString(m) }
func (*EventMissedBlock) ProtoMessage()    {}
func (*EventMissedBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_7944b2035d0f825b, []int{1}
}
func (m *EventMissedBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventMissedBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventMissedBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventMissedBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventMissedBlock.Merge(m, src)
}
func (m *EventMissedBlock) XXX_Size() int {
	return m.Size()
}
func (m *EventMissedBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_EventMissedBlock.DiscardUnknown(m)
}

var xxx_messageInfo_EventMissedBlock proto.InternalMessageInfo

func (m *EventMissedBlock) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// EventRemoveValidator is an event emitted when a validator is removed.
type EventRemoveValidator struct {
	// address is the address of the validator.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *EventRemoveValidator) Reset()         { *m = EventRemoveValidator{} }
func (m *EventRemoveValidator) String() string { return proto.CompactTextString(m) }
func (*EventRemoveValidator) ProtoMessage()    {}
func (*EventRemoveValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_7944b2035d0f825b, []int{2}
}
func (m *EventRemoveValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRemoveValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRemoveValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRemoveValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRemoveValidator.Merge(m, src)
}
func (m *EventRemoveValidator) XXX_Size() int {
	return m.Size()
}
func (m *EventRemoveValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRemoveValidator.DiscardUnknown(m)
}

var xxx_messageInfo_EventRemoveValidator proto.InternalMessageInfo

func (m *EventRemoveValidator) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// EventUpdatePolicy is an event emitted when the policy is updated.
type EventUpdatePolicy struct {
	// signed_blocks_window is the window within which a validator is expected to
	// sign a block.
	SignedBlocksWindow int64 `protobuf:"varint,1,opt,name=signed_blocks_window,json=signedBlocksWindow,proto3" json:"signed_blocks_window,omitempty"`
	// min_signed_per_window is the minimum number of signed blocks per signed
	// blocks window.
	MinSignedPerWindow int64 `protobuf:"varint,2,opt,name=min_signed_per_window,json=minSignedPerWindow,proto3" json:"min_signed_per_window,omitempty"`
}

func (m *EventUpdatePolicy) Reset()         { *m = EventUpdatePolicy{} }
func (m *EventUpdatePolicy) String() string { return proto.CompactTextString(m) }
func (*EventUpdatePolicy) ProtoMessage()    {}
func (*EventUpdatePolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_7944b2035d0f825b, []int{3}
}
func (m *EventUpdatePolicy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventUpdatePolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventUpdatePolicy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventUpdatePolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventUpdatePolicy.Merge(m, src)
}
func (m *EventUpdatePolicy) XXX_Size() int {
	return m.Size()
}
func (m *EventUpdatePolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_EventUpdatePolicy.DiscardUnknown(m)
}

var xxx_messageInfo_EventUpdatePolicy proto.InternalMessageInfo

func (m *EventUpdatePolicy) GetSignedBlocksWindow() int64 {
	if m != nil {
		return m.SignedBlocksWindow
	}
	return 0
}

func (m *EventUpdatePolicy) GetMinSignedPerWindow() int64 {
	if m != nil {
		return m.MinSignedPerWindow
	}
	return 0
}

// EventUpdateValidator is an event emitted when a validator is updated.
type EventUpdateValidator struct {
	// address is the address of the validator.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *EventUpdateValidator) Reset()         { *m = EventUpdateValidator{} }
func (m *EventUpdateValidator) String() string { return proto.CompactTextString(m) }
func (*EventUpdateValidator) ProtoMessage()    {}
func (*EventUpdateValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_7944b2035d0f825b, []int{4}
}
func (m *EventUpdateValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventUpdateValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventUpdateValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventUpdateValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventUpdateValidator.Merge(m, src)
}
func (m *EventUpdateValidator) XXX_Size() int {
	return m.Size()
}
func (m *EventUpdateValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_EventUpdateValidator.DiscardUnknown(m)
}

var xxx_messageInfo_EventUpdateValidator proto.InternalMessageInfo

func (m *EventUpdateValidator) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*EventAddValidator)(nil), "chora.validator.v1.EventAddValidator")
	proto.RegisterType((*EventMissedBlock)(nil), "chora.validator.v1.EventMissedBlock")
	proto.RegisterType((*EventRemoveValidator)(nil), "chora.validator.v1.EventRemoveValidator")
	proto.RegisterType((*EventUpdatePolicy)(nil), "chora.validator.v1.EventUpdatePolicy")
	proto.RegisterType((*EventUpdateValidator)(nil), "chora.validator.v1.EventUpdateValidator")
}

func init() { proto.RegisterFile("chora/validator/v1/events.proto", fileDescriptor_7944b2035d0f825b) }

var fileDescriptor_7944b2035d0f825b = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4b, 0xf3, 0x30,
	0x18, 0xc7, 0xdb, 0xf7, 0x05, 0xc5, 0x9c, 0x34, 0x4c, 0xd8, 0x29, 0xca, 0x4e, 0x82, 0xae, 0x59,
	0xf1, 0x13, 0x38, 0xd8, 0x51, 0x18, 0x15, 0x15, 0xbc, 0x94, 0xb6, 0x79, 0xd8, 0x82, 0x6d, 0x9e,
	0x92, 0xc4, 0xcc, 0x7d, 0x0b, 0x3f, 0x96, 0xc7, 0x1d, 0x3d, 0x4a, 0xfb, 0x45, 0x64, 0xe9, 0x3a,
	0x3c, 0xc9, 0x8e, 0xc9, 0xff, 0xf7, 0xcb, 0xf3, 0x27, 0x0f, 0xb9, 0x28, 0x96, 0xa8, 0x33, 0xee,
	0xb2, 0x52, 0x8a, 0xcc, 0xa2, 0xe6, 0x2e, 0xe6, 0xe0, 0x40, 0x59, 0x13, 0xd5, 0x1a, 0x2d, 0x52,
	0xea, 0x81, 0x68, 0x0f, 0x44, 0x2e, 0x1e, 0x8d, 0xc9, 0xd9, 0x6c, 0xcb, 0xdc, 0x09, 0xf1, 0xd4,
	0xdf, 0xd3, 0x21, 0x39, 0xce, 0x84, 0xd0, 0x60, 0xcc, 0x30, 0xbc, 0x0c, 0xaf, 0x4e, 0x92, 0xfe,
	0x38, 0xba, 0x21, 0xa7, 0x1e, 0xbf, 0x97, 0xc6, 0x80, 0x98, 0x96, 0x58, 0xbc, 0xfe, 0x41, 0x4f,
	0xc8, 0xc0, 0xd3, 0x09, 0x54, 0xe8, 0xe0, 0x90, 0xf7, 0xdf, 0x77, 0x75, 0x1e, 0x6b, 0x91, 0x59,
	0x98, 0x63, 0x29, 0x8b, 0x35, 0x9d, 0x90, 0x81, 0x91, 0x0b, 0x05, 0x22, 0xcd, 0xb7, 0x03, 0x4d,
	0xba, 0x92, 0x4a, 0xe0, 0xca, 0xbb, 0xff, 0x13, 0xda, 0x65, 0xbe, 0x8b, 0x79, 0xf6, 0x09, 0x8d,
	0xc9, 0x79, 0x25, 0x55, 0xba, 0xb3, 0x6a, 0xd0, 0xbd, 0xf2, 0xaf, 0x53, 0x2a, 0xa9, 0x1e, 0x7c,
	0x36, 0x07, 0xdd, 0x29, 0xfb, 0xae, 0xdd, 0xe4, 0x03, 0xba, 0x4e, 0x67, 0x9f, 0x0d, 0x0b, 0x37,
	0x0d, 0x0b, 0xbf, 0x1b, 0x16, 0x7e, 0xb4, 0x2c, 0xd8, 0xb4, 0x2c, 0xf8, 0x6a, 0x59, 0xf0, 0x72,
	0xbd, 0x90, 0x76, 0xf9, 0x96, 0x47, 0x05, 0x56, 0xdc, 0xff, 0xf9, 0x58, 0x22, 0xaf, 0x50, 0x98,
	0x5f, 0xcb, 0xb1, 0xeb, 0x1a, 0x0c, 0x77, 0x71, 0x7e, 0xe4, 0x97, 0x73, 0xfb, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x97, 0x45, 0x94, 0x2a, 0xbf, 0x01, 0x00, 0x00,
}

func (m *EventAddValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventAddValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventAddValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventMissedBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventMissedBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventMissedBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventRemoveValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRemoveValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRemoveValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventUpdatePolicy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventUpdatePolicy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventUpdatePolicy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MinSignedPerWindow != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.MinSignedPerWindow))
		i--
		dAtA[i] = 0x10
	}
	if m.SignedBlocksWindow != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.SignedBlocksWindow))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventUpdateValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventUpdateValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventUpdateValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventAddValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventMissedBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventRemoveValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventUpdatePolicy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SignedBlocksWindow != 0 {
		n += 1 + sovEvents(uint64(m.SignedBlocksWindow))
	}
	if m.MinSignedPerWindow != 0 {
		n += 1 + sovEvents(uint64(m.MinSignedPerWindow))
	}
	return n
}

func (m *EventUpdateValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventAddValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventAddValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventAddValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventMissedBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventMissedBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventMissedBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventRemoveValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventRemoveValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRemoveValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventUpdatePolicy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventUpdatePolicy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventUpdatePolicy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignedBlocksWindow", wireType)
			}
			m.SignedBlocksWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignedBlocksWindow |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinSignedPerWindow", wireType)
			}
			m.MinSignedPerWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinSignedPerWindow |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventUpdateValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventUpdateValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventUpdateValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
