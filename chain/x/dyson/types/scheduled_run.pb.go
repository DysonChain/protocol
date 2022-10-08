// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dyson/scheduled_run.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type ScheduledRun struct {
	Index    string          `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Creator  string          `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Height   uint64          `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Gas      uint64          `protobuf:"varint,4,opt,name=gas,proto3" json:"gas,omitempty"`
	Msg      *MsgRun         `protobuf:"bytes,5,opt,name=msg,proto3" json:"msg,omitempty"`
	Resp     *MsgRunResponse `protobuf:"bytes,6,opt,name=resp,proto3" json:"resp,omitempty"`
	Error    string          `protobuf:"bytes,7,opt,name=error,proto3" json:"error,omitempty"`
	Gasprice *types.DecCoin  `protobuf:"bytes,8,opt,name=gasprice,proto3" json:"gasprice,omitempty"`
	Fee      *types.Coin     `protobuf:"bytes,9,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (m *ScheduledRun) Reset()         { *m = ScheduledRun{} }
func (m *ScheduledRun) String() string { return proto.CompactTextString(m) }
func (*ScheduledRun) ProtoMessage()    {}
func (*ScheduledRun) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb6987718177d5c8, []int{0}
}
func (m *ScheduledRun) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScheduledRun) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ScheduledRun.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ScheduledRun) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduledRun.Merge(m, src)
}
func (m *ScheduledRun) XXX_Size() int {
	return m.Size()
}
func (m *ScheduledRun) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduledRun.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduledRun proto.InternalMessageInfo

func (m *ScheduledRun) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *ScheduledRun) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ScheduledRun) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ScheduledRun) GetGas() uint64 {
	if m != nil {
		return m.Gas
	}
	return 0
}

func (m *ScheduledRun) GetMsg() *MsgRun {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *ScheduledRun) GetResp() *MsgRunResponse {
	if m != nil {
		return m.Resp
	}
	return nil
}

func (m *ScheduledRun) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ScheduledRun) GetGasprice() *types.DecCoin {
	if m != nil {
		return m.Gasprice
	}
	return nil
}

func (m *ScheduledRun) GetFee() *types.Coin {
	if m != nil {
		return m.Fee
	}
	return nil
}

func init() {
	proto.RegisterType((*ScheduledRun)(nil), "dyson.ScheduledRun")
}

func init() { proto.RegisterFile("dyson/scheduled_run.proto", fileDescriptor_bb6987718177d5c8) }

var fileDescriptor_bb6987718177d5c8 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x6d, 0x9a, 0x7e, 0xae, 0x0a, 0xb2, 0x54, 0xd9, 0x16, 0x89, 0xa5, 0xa7, 0x8a, 0xb0, 0xa1,
	0x7a, 0xf1, 0xe0, 0x49, 0xbd, 0x7a, 0x89, 0x37, 0x2f, 0x92, 0x8f, 0x71, 0x13, 0x30, 0x99, 0xb0,
	0x93, 0x48, 0xfb, 0x2f, 0xfc, 0x2b, 0xfe, 0x0b, 0x8f, 0x3d, 0x7a, 0x94, 0xf6, 0x8f, 0x48, 0x36,
	0x69, 0x41, 0xf0, 0xb4, 0xf3, 0x66, 0xde, 0x7b, 0xfb, 0x78, 0x6c, 0x1c, 0xad, 0x08, 0x33, 0x97,
	0xc2, 0x18, 0xa2, 0xf2, 0x0d, 0xa2, 0x17, 0x5d, 0x66, 0x32, 0xd7, 0x58, 0x20, 0xef, 0x9a, 0xd3,
	0x64, 0xa4, 0x50, 0xa1, 0xd9, 0xb8, 0xd5, 0x54, 0x1f, 0x27, 0x4e, 0x88, 0x94, 0x22, 0xb9, 0x81,
	0x4f, 0xe0, 0xbe, 0x2f, 0x02, 0x28, 0xfc, 0x85, 0x1b, 0x62, 0xd2, 0x88, 0x27, 0xbc, 0xf6, 0x4d,
	0x49, 0xed, 0x0d, 0x67, 0x9f, 0x6d, 0x76, 0xf8, 0xb4, 0xfb, 0xc8, 0x2b, 0x33, 0x3e, 0x62, 0xdd,
	0x24, 0x8b, 0x60, 0x29, 0xac, 0xa9, 0x35, 0x1f, 0x7a, 0x35, 0xe0, 0x82, 0xf5, 0x43, 0x0d, 0x7e,
	0x81, 0x5a, 0xb4, 0xcd, 0x7e, 0x07, 0xf9, 0x29, 0xeb, 0xc5, 0x90, 0xa8, 0xb8, 0x10, 0xf6, 0xd4,
	0x9a, 0x77, 0xbc, 0x06, 0xf1, 0x63, 0x66, 0x2b, 0x9f, 0x44, 0xc7, 0x2c, 0xab, 0x91, 0x9f, 0x33,
	0x3b, 0x25, 0x25, 0xba, 0x53, 0x6b, 0x7e, 0x70, 0x75, 0x24, 0x4d, 0x18, 0xf9, 0x48, 0xca, 0x2b,
	0x33, 0xaf, 0xba, 0xf0, 0x0b, 0xd6, 0xd1, 0x40, 0xb9, 0xe8, 0x19, 0xc6, 0xc9, 0x5f, 0x06, 0x50,
	0x8e, 0x19, 0x81, 0x67, 0x28, 0x55, 0x4a, 0xd0, 0x1a, 0xb5, 0xe8, 0xd7, 0x29, 0x0d, 0xe0, 0x37,
	0x6c, 0xa0, 0x7c, 0xca, 0x75, 0x12, 0x82, 0x18, 0x18, 0x93, 0x33, 0x59, 0x77, 0x22, 0xab, 0x4e,
	0x64, 0xd3, 0x89, 0x7c, 0x80, 0xf0, 0x1e, 0x93, 0xcc, 0xdb, 0xb3, 0xf9, 0x25, 0xb3, 0x5f, 0x01,
	0xc4, 0xd0, 0x88, 0xc6, 0xff, 0x8a, 0x8c, 0xa2, 0x62, 0xdd, 0xdd, 0x7e, 0x6d, 0x1c, 0x6b, 0xbd,
	0x71, 0xac, 0x9f, 0x8d, 0x63, 0x7d, 0x6c, 0x9d, 0xd6, 0x7a, 0xeb, 0xb4, 0xbe, 0xb7, 0x4e, 0xeb,
	0x79, 0xa6, 0x92, 0x22, 0x2e, 0x03, 0x19, 0x62, 0xea, 0xa2, 0x56, 0x6e, 0x5d, 0xf8, 0xb2, 0x79,
	0x8b, 0x55, 0x0e, 0x14, 0xf4, 0x4c, 0xf1, 0xd7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x49,
	0x13, 0x5a, 0xe6, 0x01, 0x00, 0x00,
}

func (m *ScheduledRun) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScheduledRun) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ScheduledRun) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Fee != nil {
		{
			size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintScheduledRun(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if m.Gasprice != nil {
		{
			size, err := m.Gasprice.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintScheduledRun(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintScheduledRun(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Resp != nil {
		{
			size, err := m.Resp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintScheduledRun(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.Msg != nil {
		{
			size, err := m.Msg.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintScheduledRun(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.Gas != 0 {
		i = encodeVarintScheduledRun(dAtA, i, uint64(m.Gas))
		i--
		dAtA[i] = 0x20
	}
	if m.Height != 0 {
		i = encodeVarintScheduledRun(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintScheduledRun(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintScheduledRun(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintScheduledRun(dAtA []byte, offset int, v uint64) int {
	offset -= sovScheduledRun(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ScheduledRun) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovScheduledRun(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovScheduledRun(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovScheduledRun(uint64(m.Height))
	}
	if m.Gas != 0 {
		n += 1 + sovScheduledRun(uint64(m.Gas))
	}
	if m.Msg != nil {
		l = m.Msg.Size()
		n += 1 + l + sovScheduledRun(uint64(l))
	}
	if m.Resp != nil {
		l = m.Resp.Size()
		n += 1 + l + sovScheduledRun(uint64(l))
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovScheduledRun(uint64(l))
	}
	if m.Gasprice != nil {
		l = m.Gasprice.Size()
		n += 1 + l + sovScheduledRun(uint64(l))
	}
	if m.Fee != nil {
		l = m.Fee.Size()
		n += 1 + l + sovScheduledRun(uint64(l))
	}
	return n
}

func sovScheduledRun(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozScheduledRun(x uint64) (n int) {
	return sovScheduledRun(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ScheduledRun) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowScheduledRun
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
			return fmt.Errorf("proto: ScheduledRun: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScheduledRun: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheduledRun
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
				return ErrInvalidLengthScheduledRun
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScheduledRun
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheduledRun
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
				return ErrInvalidLengthScheduledRun
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScheduledRun
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheduledRun
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gas", wireType)
			}
			m.Gas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheduledRun
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Gas |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheduledRun
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
				return ErrInvalidLengthScheduledRun
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthScheduledRun
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Msg == nil {
				m.Msg = &MsgRun{}
			}
			if err := m.Msg.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheduledRun
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
				return ErrInvalidLengthScheduledRun
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthScheduledRun
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Resp == nil {
				m.Resp = &MsgRunResponse{}
			}
			if err := m.Resp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheduledRun
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
				return ErrInvalidLengthScheduledRun
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScheduledRun
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gasprice", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheduledRun
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
				return ErrInvalidLengthScheduledRun
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthScheduledRun
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Gasprice == nil {
				m.Gasprice = &types.DecCoin{}
			}
			if err := m.Gasprice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheduledRun
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
				return ErrInvalidLengthScheduledRun
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthScheduledRun
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Fee == nil {
				m.Fee = &types.Coin{}
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipScheduledRun(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthScheduledRun
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
func skipScheduledRun(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowScheduledRun
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
					return 0, ErrIntOverflowScheduledRun
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
					return 0, ErrIntOverflowScheduledRun
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
				return 0, ErrInvalidLengthScheduledRun
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupScheduledRun
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthScheduledRun
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthScheduledRun        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowScheduledRun          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupScheduledRun = fmt.Errorf("proto: unexpected end of group")
)