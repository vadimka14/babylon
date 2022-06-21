// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/btclightclient/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// MsgInsertHeader defines the message for incoming header bytes
type MsgInsertHeader struct {
	Signer string          `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	Header *BTCHeaderBytes `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
}

func (m *MsgInsertHeader) Reset()         { *m = MsgInsertHeader{} }
func (m *MsgInsertHeader) String() string { return proto.CompactTextString(m) }
func (*MsgInsertHeader) ProtoMessage()    {}
func (*MsgInsertHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_84e67479ce863198, []int{0}
}
func (m *MsgInsertHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInsertHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInsertHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInsertHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInsertHeader.Merge(m, src)
}
func (m *MsgInsertHeader) XXX_Size() int {
	return m.Size()
}
func (m *MsgInsertHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInsertHeader.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInsertHeader proto.InternalMessageInfo

func (m *MsgInsertHeader) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *MsgInsertHeader) GetHeader() *BTCHeaderBytes {
	if m != nil {
		return m.Header
	}
	return nil
}

type MsgInsertHeaderResponse struct {
}

func (m *MsgInsertHeaderResponse) Reset()         { *m = MsgInsertHeaderResponse{} }
func (m *MsgInsertHeaderResponse) String() string { return proto.CompactTextString(m) }
func (*MsgInsertHeaderResponse) ProtoMessage()    {}
func (*MsgInsertHeaderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_84e67479ce863198, []int{1}
}
func (m *MsgInsertHeaderResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInsertHeaderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInsertHeaderResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInsertHeaderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInsertHeaderResponse.Merge(m, src)
}
func (m *MsgInsertHeaderResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgInsertHeaderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInsertHeaderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInsertHeaderResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgInsertHeader)(nil), "babylon.btclightclient.v1.MsgInsertHeader")
	proto.RegisterType((*MsgInsertHeaderResponse)(nil), "babylon.btclightclient.v1.MsgInsertHeaderResponse")
}

func init() { proto.RegisterFile("babylon/btclightclient/tx.proto", fileDescriptor_84e67479ce863198) }

var fileDescriptor_84e67479ce863198 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x4a, 0x4c, 0xaa,
	0xcc, 0xc9, 0xcf, 0xd3, 0x4f, 0x2a, 0x49, 0xce, 0xc9, 0x4c, 0xcf, 0x00, 0x91, 0xa9, 0x79, 0x25,
	0xfa, 0x25, 0x15, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x92, 0x50, 0x05, 0x7a, 0xa8, 0x0a,
	0xf4, 0xca, 0x0c, 0xa5, 0xb4, 0x71, 0xe8, 0x45, 0x53, 0x09, 0x36, 0x47, 0x29, 0x87, 0x8b, 0xdf,
	0xb7, 0x38, 0xdd, 0x33, 0xaf, 0x38, 0xb5, 0xa8, 0xc4, 0x23, 0x35, 0x31, 0x25, 0xb5, 0x48, 0x48,
	0x8c, 0x8b, 0xad, 0x38, 0x33, 0x3d, 0x2f, 0xb5, 0x48, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08,
	0xca, 0x13, 0x72, 0xe4, 0x62, 0xcb, 0x00, 0xab, 0x90, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x36, 0xd2,
	0xd4, 0xc3, 0xe9, 0x06, 0x3d, 0xa7, 0x10, 0x67, 0x88, 0x69, 0x4e, 0x95, 0x25, 0xa9, 0xc5, 0x41,
	0x50, 0x8d, 0x4a, 0x92, 0x5c, 0xe2, 0x68, 0xb6, 0x05, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7,
	0x1a, 0x95, 0x73, 0x31, 0xfb, 0x16, 0xa7, 0x0b, 0x15, 0x70, 0xf1, 0xa0, 0x38, 0x46, 0x0b, 0x8f,
	0x25, 0x68, 0x46, 0x49, 0x19, 0x11, 0xaf, 0x16, 0x66, 0xad, 0x12, 0x83, 0x53, 0xc0, 0x89, 0x47,
	0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85,
	0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x99, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9,
	0x25, 0xe7, 0xe7, 0xea, 0x43, 0x4d, 0x4e, 0xce, 0x48, 0xcc, 0xcc, 0x83, 0x71, 0xf4, 0x2b, 0x30,
	0xa2, 0xa7, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0x1c, 0xb4, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xfc, 0x68, 0x13, 0x8c, 0xc5, 0x01, 0x00, 0x00,
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
	InsertHeader(ctx context.Context, in *MsgInsertHeader, opts ...grpc.CallOption) (*MsgInsertHeaderResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) InsertHeader(ctx context.Context, in *MsgInsertHeader, opts ...grpc.CallOption) (*MsgInsertHeaderResponse, error) {
	out := new(MsgInsertHeaderResponse)
	err := c.cc.Invoke(ctx, "/babylon.btclightclient.v1.Msg/InsertHeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	InsertHeader(context.Context, *MsgInsertHeader) (*MsgInsertHeaderResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) InsertHeader(ctx context.Context, req *MsgInsertHeader) (*MsgInsertHeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertHeader not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_InsertHeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInsertHeader)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).InsertHeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/babylon.btclightclient.v1.Msg/InsertHeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).InsertHeader(ctx, req.(*MsgInsertHeader))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "babylon.btclightclient.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertHeader",
			Handler:    _Msg_InsertHeader_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "babylon/btclightclient/tx.proto",
}

func (m *MsgInsertHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInsertHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInsertHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgInsertHeaderResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInsertHeaderResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInsertHeaderResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgInsertHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgInsertHeaderResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgInsertHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgInsertHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInsertHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &BTCHeaderBytes{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgInsertHeaderResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgInsertHeaderResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInsertHeaderResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)