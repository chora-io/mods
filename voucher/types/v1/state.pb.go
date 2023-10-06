// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chora/voucher/v1/state.proto

package v1

import (
	_ "cosmossdk.io/orm"
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/cosmos/gogoproto/types"
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

// Voucher defines the table and properties of a voucher.
type Voucher struct {
	// id is the unique identifier of the voucher.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// issuer is the address of the voucher issuer.
	Issuer []byte `protobuf:"bytes,2,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// metadata is the metadata of the voucher.
	Metadata string `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *Voucher) Reset()         { *m = Voucher{} }
func (m *Voucher) String() string { return proto.CompactTextString(m) }
func (*Voucher) ProtoMessage()    {}
func (*Voucher) Descriptor() ([]byte, []int) {
	return fileDescriptor_532d2718e66f4a8a, []int{0}
}
func (m *Voucher) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Voucher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Voucher.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Voucher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Voucher.Merge(m, src)
}
func (m *Voucher) XXX_Size() int {
	return m.Size()
}
func (m *Voucher) XXX_DiscardUnknown() {
	xxx_messageInfo_Voucher.DiscardUnknown(m)
}

var xxx_messageInfo_Voucher proto.InternalMessageInfo

func (m *Voucher) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Voucher) GetIssuer() []byte {
	if m != nil {
		return m.Issuer
	}
	return nil
}

func (m *Voucher) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

// Balance defines the table and properties of a voucher balance.
type Balance struct {
	// id is the unique identifier of the voucher.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// address is the address of the voucher owner.
	Address []byte `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// amount is the amount of vouchers the address owns.
	Amount string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	// expiration is the expiration of the vouchers.
	Expiration *types.Timestamp `protobuf:"bytes,4,opt,name=expiration,proto3" json:"expiration,omitempty"`
}

func (m *Balance) Reset()         { *m = Balance{} }
func (m *Balance) String() string { return proto.CompactTextString(m) }
func (*Balance) ProtoMessage()    {}
func (*Balance) Descriptor() ([]byte, []int) {
	return fileDescriptor_532d2718e66f4a8a, []int{1}
}
func (m *Balance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Balance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Balance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Balance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Balance.Merge(m, src)
}
func (m *Balance) XXX_Size() int {
	return m.Size()
}
func (m *Balance) XXX_DiscardUnknown() {
	xxx_messageInfo_Balance.DiscardUnknown(m)
}

var xxx_messageInfo_Balance proto.InternalMessageInfo

func (m *Balance) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Balance) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Balance) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *Balance) GetExpiration() *types.Timestamp {
	if m != nil {
		return m.Expiration
	}
	return nil
}

func init() {
	proto.RegisterType((*Voucher)(nil), "chora.voucher.v1.Voucher")
	proto.RegisterType((*Balance)(nil), "chora.voucher.v1.Balance")
}

func init() { proto.RegisterFile("chora/voucher/v1/state.proto", fileDescriptor_532d2718e66f4a8a) }

var fileDescriptor_532d2718e66f4a8a = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x4f, 0x4a, 0xc3, 0x40,
	0x18, 0xc5, 0x3b, 0x69, 0x69, 0x75, 0x2a, 0x25, 0x0c, 0xd8, 0x86, 0x50, 0x62, 0xe8, 0x2a, 0x8b,
	0x92, 0xa1, 0xba, 0x91, 0x2e, 0x5c, 0xf4, 0x08, 0x41, 0x5c, 0xb8, 0x9b, 0x26, 0x63, 0x3b, 0xd8,
	0xe9, 0x17, 0x66, 0x26, 0x41, 0x2f, 0x21, 0x9e, 0xc0, 0xeb, 0xe8, 0xb2, 0xe0, 0xc6, 0xa5, 0xb4,
	0x37, 0xf0, 0x04, 0xd2, 0xfc, 0xa9, 0x05, 0x97, 0x8f, 0xef, 0xf1, 0x7e, 0xef, 0xf1, 0xe1, 0x61,
	0xbc, 0x04, 0xc5, 0x68, 0x0e, 0x59, 0xbc, 0xe4, 0x8a, 0xe6, 0x13, 0xaa, 0x0d, 0x33, 0x3c, 0x4c,
	0x15, 0x18, 0x20, 0x76, 0x71, 0x0d, 0xab, 0x6b, 0x98, 0x4f, 0xdc, 0x41, 0x0c, 0x5a, 0x82, 0xa6,
	0xa0, 0xe4, 0xde, 0x0c, 0x4a, 0x96, 0x56, 0xf7, 0x62, 0x01, 0xb0, 0x58, 0x71, 0x5a, 0xa8, 0x79,
	0xf6, 0x40, 0x8d, 0x90, 0x5c, 0x1b, 0x26, 0xd3, 0xd2, 0x30, 0x7a, 0xc4, 0x9d, 0xbb, 0x32, 0x87,
	0xf4, 0xb0, 0x25, 0x12, 0x07, 0xf9, 0x28, 0x68, 0x45, 0x96, 0x48, 0x48, 0x1f, 0xb7, 0x85, 0xd6,
	0x19, 0x57, 0x8e, 0xe5, 0xa3, 0xe0, 0x2c, 0xaa, 0x14, 0x71, 0xf1, 0x89, 0xe4, 0x86, 0x25, 0xcc,
	0x30, 0xa7, 0xe9, 0xa3, 0xe0, 0x34, 0x3a, 0xe8, 0xe9, 0xf0, 0xe7, 0xed, 0xf3, 0xa5, 0xd9, 0xc7,
	0xed, 0x7d, 0x96, 0x8d, 0x08, 0xae, 0x33, 0x6c, 0xe4, 0xa0, 0xd1, 0x3b, 0xc2, 0x9d, 0x19, 0x5b,
	0xb1, 0x75, 0xcc, 0xff, 0xd1, 0x1c, 0xdc, 0x61, 0x49, 0xa2, 0xb8, 0xd6, 0x15, 0xae, 0x96, 0xfb,
	0x1e, 0x4c, 0x42, 0xb6, 0x36, 0x15, 0xad, 0x52, 0x64, 0x8a, 0x31, 0x7f, 0x4a, 0x85, 0x62, 0x46,
	0xc0, 0xda, 0x69, 0xf9, 0x28, 0xe8, 0x5e, 0xba, 0x61, 0x39, 0x38, 0xac, 0x07, 0x87, 0xb7, 0xf5,
	0xe0, 0xe8, 0xc8, 0x3d, 0xbd, 0x29, 0x7a, 0x5e, 0xe3, 0x01, 0x3e, 0x17, 0xc9, 0xb8, 0x22, 0x8d,
	0xff, 0x0c, 0xa4, 0x7b, 0xa8, 0x63, 0x23, 0xd2, 0x3b, 0x26, 0xd9, 0x96, 0x63, 0xcd, 0x66, 0x1f,
	0x5b, 0x0f, 0x6d, 0xb6, 0x1e, 0xfa, 0xde, 0x7a, 0xe8, 0x75, 0xe7, 0x35, 0x36, 0x3b, 0xaf, 0xf1,
	0xb5, 0xf3, 0x1a, 0xf7, 0xc1, 0x42, 0x98, 0x65, 0x36, 0x0f, 0x63, 0x90, 0xb4, 0xf8, 0x93, 0x00,
	0x2a, 0x21, 0xd1, 0x87, 0x67, 0x9a, 0xe7, 0x94, 0x6b, 0x9a, 0x4f, 0xe6, 0xed, 0xa2, 0xe3, 0xd5,
	0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x34, 0x73, 0x0c, 0xd6, 0xed, 0x01, 0x00, 0x00,
}

func (m *Voucher) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Voucher) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Voucher) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Metadata) > 0 {
		i -= len(m.Metadata)
		copy(dAtA[i:], m.Metadata)
		i = encodeVarintState(dAtA, i, uint64(len(m.Metadata)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Issuer) > 0 {
		i -= len(m.Issuer)
		copy(dAtA[i:], m.Issuer)
		i = encodeVarintState(dAtA, i, uint64(len(m.Issuer)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Balance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Balance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Balance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Expiration != nil {
		{
			size, err := m.Expiration.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintState(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintState(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintState(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintState(dAtA []byte, offset int, v uint64) int {
	offset -= sovState(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Voucher) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovState(uint64(m.Id))
	}
	l = len(m.Issuer)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.Metadata)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	return n
}

func (m *Balance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovState(uint64(m.Id))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.Expiration != nil {
		l = m.Expiration.Size()
		n += 1 + l + sovState(uint64(l))
	}
	return n
}

func sovState(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozState(x uint64) (n int) {
	return sovState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Voucher) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: Voucher: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Voucher: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issuer = append(m.Issuer[:0], dAtA[iNdEx:postIndex]...)
			if m.Issuer == nil {
				m.Issuer = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func (m *Balance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: Balance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Balance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Expiration == nil {
				m.Expiration = &types.Timestamp{}
			}
			if err := m.Expiration.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func skipState(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
				return 0, ErrInvalidLengthState
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupState
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthState
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthState        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowState          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupState = fmt.Errorf("proto: unexpected end of group")
)
