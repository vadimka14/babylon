// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/btcstaking/v1/events.proto

package types

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

// EventNewBTCValidator is the event emitted when a BTC validator is created
type EventNewBTCValidator struct {
	BtcVal *BTCValidator `protobuf:"bytes,1,opt,name=btc_val,json=btcVal,proto3" json:"btc_val,omitempty"`
}

func (m *EventNewBTCValidator) Reset()         { *m = EventNewBTCValidator{} }
func (m *EventNewBTCValidator) String() string { return proto.CompactTextString(m) }
func (*EventNewBTCValidator) ProtoMessage()    {}
func (*EventNewBTCValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_74118427820fff75, []int{0}
}
func (m *EventNewBTCValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventNewBTCValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventNewBTCValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventNewBTCValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventNewBTCValidator.Merge(m, src)
}
func (m *EventNewBTCValidator) XXX_Size() int {
	return m.Size()
}
func (m *EventNewBTCValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_EventNewBTCValidator.DiscardUnknown(m)
}

var xxx_messageInfo_EventNewBTCValidator proto.InternalMessageInfo

func (m *EventNewBTCValidator) GetBtcVal() *BTCValidator {
	if m != nil {
		return m.BtcVal
	}
	return nil
}

// EventNewBTCDelegation is the event emitted when a BTC delegation is created
// NOTE: the BTC delegation is not active thus does not have voting power yet
// only after it receives a jury signature it becomes activated and has voting power
type EventNewBTCDelegation struct {
	BtcDel *BTCDelegation `protobuf:"bytes,1,opt,name=btc_del,json=btcDel,proto3" json:"btc_del,omitempty"`
}

func (m *EventNewBTCDelegation) Reset()         { *m = EventNewBTCDelegation{} }
func (m *EventNewBTCDelegation) String() string { return proto.CompactTextString(m) }
func (*EventNewBTCDelegation) ProtoMessage()    {}
func (*EventNewBTCDelegation) Descriptor() ([]byte, []int) {
	return fileDescriptor_74118427820fff75, []int{1}
}
func (m *EventNewBTCDelegation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventNewBTCDelegation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventNewBTCDelegation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventNewBTCDelegation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventNewBTCDelegation.Merge(m, src)
}
func (m *EventNewBTCDelegation) XXX_Size() int {
	return m.Size()
}
func (m *EventNewBTCDelegation) XXX_DiscardUnknown() {
	xxx_messageInfo_EventNewBTCDelegation.DiscardUnknown(m)
}

var xxx_messageInfo_EventNewBTCDelegation proto.InternalMessageInfo

func (m *EventNewBTCDelegation) GetBtcDel() *BTCDelegation {
	if m != nil {
		return m.BtcDel
	}
	return nil
}

// EventActivateBTCDelegation is the event emitted when jury activates a BTC delegation
// such that the BTC delegation starts to have voting power in its timelock period
type EventActivateBTCDelegation struct {
	BtcDel *BTCDelegation `protobuf:"bytes,1,opt,name=btc_del,json=btcDel,proto3" json:"btc_del,omitempty"`
}

func (m *EventActivateBTCDelegation) Reset()         { *m = EventActivateBTCDelegation{} }
func (m *EventActivateBTCDelegation) String() string { return proto.CompactTextString(m) }
func (*EventActivateBTCDelegation) ProtoMessage()    {}
func (*EventActivateBTCDelegation) Descriptor() ([]byte, []int) {
	return fileDescriptor_74118427820fff75, []int{2}
}
func (m *EventActivateBTCDelegation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventActivateBTCDelegation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventActivateBTCDelegation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventActivateBTCDelegation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventActivateBTCDelegation.Merge(m, src)
}
func (m *EventActivateBTCDelegation) XXX_Size() int {
	return m.Size()
}
func (m *EventActivateBTCDelegation) XXX_DiscardUnknown() {
	xxx_messageInfo_EventActivateBTCDelegation.DiscardUnknown(m)
}

var xxx_messageInfo_EventActivateBTCDelegation proto.InternalMessageInfo

func (m *EventActivateBTCDelegation) GetBtcDel() *BTCDelegation {
	if m != nil {
		return m.BtcDel
	}
	return nil
}

func init() {
	proto.RegisterType((*EventNewBTCValidator)(nil), "babylon.btcstaking.v1.EventNewBTCValidator")
	proto.RegisterType((*EventNewBTCDelegation)(nil), "babylon.btcstaking.v1.EventNewBTCDelegation")
	proto.RegisterType((*EventActivateBTCDelegation)(nil), "babylon.btcstaking.v1.EventActivateBTCDelegation")
}

func init() {
	proto.RegisterFile("babylon/btcstaking/v1/events.proto", fileDescriptor_74118427820fff75)
}

var fileDescriptor_74118427820fff75 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4a, 0x4a, 0x4c, 0xaa,
	0xcc, 0xc9, 0xcf, 0xd3, 0x4f, 0x2a, 0x49, 0x2e, 0x2e, 0x49, 0xcc, 0xce, 0xcc, 0x4b, 0xd7, 0x2f,
	0x33, 0xd4, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12,
	0x85, 0xaa, 0xd1, 0x43, 0xa8, 0xd1, 0x2b, 0x33, 0x94, 0x52, 0xc3, 0xae, 0x15, 0x49, 0x11, 0x58,
	0xbb, 0x52, 0x08, 0x97, 0x88, 0x2b, 0xc8, 0x38, 0xbf, 0xd4, 0x72, 0xa7, 0x10, 0xe7, 0xb0, 0xc4,
	0x9c, 0xcc, 0x94, 0xc4, 0x92, 0xfc, 0x22, 0x21, 0x1b, 0x2e, 0xf6, 0xa4, 0x92, 0xe4, 0xf8, 0xb2,
	0xc4, 0x1c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x65, 0x3d, 0xac, 0x16, 0xe9, 0x21, 0xeb,
	0x0a, 0x62, 0x4b, 0x2a, 0x49, 0x0e, 0x4b, 0xcc, 0x51, 0x0a, 0xe3, 0x12, 0x45, 0x32, 0xd5, 0x25,
	0x35, 0x27, 0x35, 0x3d, 0xb1, 0x24, 0x33, 0x3f, 0x4f, 0xc8, 0x16, 0x62, 0x6c, 0x4a, 0x2a, 0xcc,
	0x58, 0x15, 0xdc, 0xc6, 0x22, 0xb4, 0x81, 0xcd, 0x75, 0x49, 0xcd, 0x51, 0x8a, 0xe6, 0x92, 0x02,
	0x9b, 0xeb, 0x98, 0x5c, 0x92, 0x59, 0x96, 0x58, 0x92, 0x4a, 0x4d, 0xc3, 0x9d, 0x7c, 0x4e, 0x3c,
	0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e,
	0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x28, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49,
	0x2f, 0x39, 0x3f, 0x57, 0x1f, 0x6a, 0x62, 0x72, 0x46, 0x62, 0x66, 0x1e, 0x8c, 0xa3, 0x5f, 0x81,
	0x1c, 0xcc, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0xe0, 0xf0, 0x35, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x88, 0x0d, 0x1d, 0x71, 0xc4, 0x01, 0x00, 0x00,
}

func (m *EventNewBTCValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventNewBTCValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventNewBTCValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BtcVal != nil {
		{
			size, err := m.BtcVal.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventNewBTCDelegation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventNewBTCDelegation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventNewBTCDelegation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BtcDel != nil {
		{
			size, err := m.BtcDel.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventActivateBTCDelegation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventActivateBTCDelegation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventActivateBTCDelegation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BtcDel != nil {
		{
			size, err := m.BtcDel.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
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
func (m *EventNewBTCValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BtcVal != nil {
		l = m.BtcVal.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventNewBTCDelegation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BtcDel != nil {
		l = m.BtcDel.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventActivateBTCDelegation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BtcDel != nil {
		l = m.BtcDel.Size()
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
func (m *EventNewBTCValidator) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventNewBTCValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventNewBTCValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BtcVal", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BtcVal == nil {
				m.BtcVal = &BTCValidator{}
			}
			if err := m.BtcVal.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *EventNewBTCDelegation) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventNewBTCDelegation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventNewBTCDelegation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BtcDel", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BtcDel == nil {
				m.BtcDel = &BTCDelegation{}
			}
			if err := m.BtcDel.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *EventActivateBTCDelegation) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventActivateBTCDelegation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventActivateBTCDelegation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BtcDel", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BtcDel == nil {
				m.BtcDel = &BTCDelegation{}
			}
			if err := m.BtcDel.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
