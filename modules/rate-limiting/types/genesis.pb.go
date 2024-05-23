// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ratelimit/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// GenesisState defines the ratelimit module's genesis state.
type GenesisState struct {
	Params                           Params                   `protobuf:"bytes,1,opt,name=params,proto3" json:"params" yaml:"params"`
	RateLimits                       []RateLimit              `protobuf:"bytes,2,rep,name=rate_limits,json=rateLimits,proto3" json:"rate_limits" yaml:"rate_limits"`
	WhitelistedAddressPairs          []WhitelistedAddressPair `protobuf:"bytes,3,rep,name=whitelisted_address_pairs,json=whitelistedAddressPairs,proto3" json:"whitelisted_address_pairs" yaml:"whitelisted_address_pairs"`
	BlacklistedDenoms                []string                 `protobuf:"bytes,4,rep,name=blacklisted_denoms,json=blacklistedDenoms,proto3" json:"blacklisted_denoms,omitempty"`
	PendingSendPacketSequenceNumbers []string                 `protobuf:"bytes,5,rep,name=pending_send_packet_sequence_numbers,json=pendingSendPacketSequenceNumbers,proto3" json:"pending_send_packet_sequence_numbers,omitempty"`
	HourEpoch                        HourEpoch                `protobuf:"bytes,6,opt,name=hour_epoch,json=hourEpoch,proto3" json:"hour_epoch" yaml:"hour_epoch"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbb08d6119688a03, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetRateLimits() []RateLimit {
	if m != nil {
		return m.RateLimits
	}
	return nil
}

func (m *GenesisState) GetWhitelistedAddressPairs() []WhitelistedAddressPair {
	if m != nil {
		return m.WhitelistedAddressPairs
	}
	return nil
}

func (m *GenesisState) GetBlacklistedDenoms() []string {
	if m != nil {
		return m.BlacklistedDenoms
	}
	return nil
}

func (m *GenesisState) GetPendingSendPacketSequenceNumbers() []string {
	if m != nil {
		return m.PendingSendPacketSequenceNumbers
	}
	return nil
}

func (m *GenesisState) GetHourEpoch() HourEpoch {
	if m != nil {
		return m.HourEpoch
	}
	return HourEpoch{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "ratelimit.v1.GenesisState")
}

func init() { proto.RegisterFile("ratelimit/v1/genesis.proto", fileDescriptor_fbb08d6119688a03) }

var fileDescriptor_fbb08d6119688a03 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xb1, 0x6e, 0xd3, 0x40,
	0x1c, 0xc6, 0x63, 0x52, 0x22, 0xf5, 0x52, 0x86, 0x5a, 0x45, 0x75, 0x2c, 0xe4, 0x5a, 0x56, 0x87,
	0x2c, 0xb1, 0xd5, 0xb2, 0x00, 0x1b, 0x01, 0x04, 0x03, 0xaa, 0x82, 0x83, 0x84, 0xc4, 0x62, 0x9d,
	0xed, 0xbf, 0xec, 0x53, 0x63, 0xdf, 0x71, 0xff, 0x73, 0xaa, 0xbe, 0x01, 0x62, 0xe2, 0xb1, 0x3a,
	0x76, 0x64, 0xaa, 0x50, 0xf2, 0x06, 0x3c, 0x01, 0xf2, 0x9d, 0x69, 0x12, 0x04, 0x9b, 0xad, 0xef,
	0xf7, 0xfb, 0x3e, 0x9d, 0x7d, 0xc4, 0x95, 0x54, 0xc1, 0x82, 0x55, 0x4c, 0x45, 0xcb, 0xb3, 0xa8,
	0x80, 0x1a, 0x90, 0x61, 0x28, 0x24, 0x57, 0xdc, 0x3e, 0xb8, 0xcf, 0xc2, 0xe5, 0x99, 0x7b, 0x54,
	0xf0, 0x82, 0xeb, 0x20, 0x6a, 0x9f, 0x0c, 0xe3, 0x8e, 0x76, 0x7c, 0x41, 0x25, 0xad, 0x3a, 0xdd,
	0x7d, 0xb2, 0x13, 0x6d, 0xba, 0x74, 0x1a, 0x7c, 0xdd, 0x23, 0x07, 0x6f, 0xcd, 0xdc, 0x5c, 0x51,
	0x05, 0xf6, 0x2b, 0x32, 0x30, 0xba, 0x63, 0xf9, 0xd6, 0x78, 0x78, 0x7e, 0x14, 0x6e, 0xcf, 0x87,
	0x33, 0x9d, 0x4d, 0x1f, 0xdf, 0xdc, 0x9d, 0xf4, 0x7e, 0xdd, 0x9d, 0x3c, 0xba, 0xa6, 0xd5, 0xe2,
	0x45, 0x60, 0x8c, 0x20, 0xee, 0x54, 0xfb, 0x23, 0x19, 0xb6, 0x56, 0xa2, 0x35, 0x74, 0x1e, 0xf8,
	0xfd, 0xf1, 0xf0, 0xfc, 0x78, 0xb7, 0x29, 0xa6, 0x0a, 0xde, 0xb7, 0x2f, 0x53, 0xb7, 0x2b, 0xb3,
	0x4d, 0xd9, 0x96, 0x19, 0xc4, 0x44, 0xfe, 0xc1, 0xd0, 0xfe, 0x66, 0x91, 0xd1, 0x55, 0xc9, 0xda,
	0x0e, 0x54, 0x90, 0x27, 0x34, 0xcf, 0x25, 0x20, 0x26, 0x82, 0x32, 0x89, 0x4e, 0x5f, 0x8f, 0x9c,
	0xee, 0x8e, 0x7c, 0xda, 0xe0, 0x2f, 0x0d, 0x3d, 0xa3, 0x4c, 0x4e, 0xc7, 0xdd, 0xa2, 0x6f, 0x16,
	0xff, 0x5b, 0x1a, 0xc4, 0xc7, 0x57, 0xff, 0x6c, 0x40, 0x7b, 0x42, 0xec, 0x74, 0x41, 0xb3, 0xcb,
	0x4e, 0xcb, 0xa1, 0xe6, 0x15, 0x3a, 0x7b, 0x7e, 0x7f, 0xbc, 0x1f, 0x1f, 0x6e, 0x25, 0xaf, 0x75,
	0x60, 0x5f, 0x90, 0x53, 0x01, 0x75, 0xce, 0xea, 0x22, 0x41, 0xa8, 0xf3, 0x44, 0xd0, 0xec, 0x12,
	0x54, 0x82, 0xf0, 0xa5, 0x81, 0x3a, 0x83, 0xa4, 0x6e, 0xaa, 0x14, 0x24, 0x3a, 0x0f, 0x75, 0x81,
	0xdf, 0xb1, 0x73, 0xa8, 0xf3, 0x99, 0x26, 0xe7, 0x1d, 0x78, 0x61, 0x38, 0xfb, 0x03, 0x21, 0x25,
	0x6f, 0x64, 0x02, 0x82, 0x67, 0xa5, 0x33, 0xd0, 0xbf, 0xea, 0xaf, 0x0f, 0xfc, 0x8e, 0x37, 0xf2,
	0x4d, 0x1b, 0x4f, 0x47, 0xdd, 0x71, 0x0f, 0xcd, 0x71, 0x37, 0x62, 0x10, 0xef, 0x97, 0xf7, 0xd4,
	0xfc, 0x66, 0xe5, 0x59, 0xb7, 0x2b, 0xcf, 0xfa, 0xb9, 0xf2, 0xac, 0xef, 0x6b, 0xaf, 0x77, 0xbb,
	0xf6, 0x7a, 0x3f, 0xd6, 0x5e, 0xef, 0xf3, 0xf3, 0x82, 0xa9, 0xb2, 0x49, 0xc3, 0x8c, 0x57, 0x51,
	0xc6, 0xb1, 0xe2, 0x18, 0xb1, 0x34, 0x9b, 0x50, 0x21, 0x30, 0xaa, 0x78, 0xde, 0x2c, 0x00, 0xf5,
	0xc5, 0x9a, 0xe8, 0x6d, 0x56, 0x17, 0xd1, 0xf2, 0x59, 0xa4, 0xae, 0x05, 0x60, 0x3a, 0xd0, 0xd7,
	0xec, 0xe9, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xea, 0x31, 0x41, 0xf9, 0xe1, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.HourEpoch.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.PendingSendPacketSequenceNumbers) > 0 {
		for iNdEx := len(m.PendingSendPacketSequenceNumbers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PendingSendPacketSequenceNumbers[iNdEx])
			copy(dAtA[i:], m.PendingSendPacketSequenceNumbers[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.PendingSendPacketSequenceNumbers[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.BlacklistedDenoms) > 0 {
		for iNdEx := len(m.BlacklistedDenoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.BlacklistedDenoms[iNdEx])
			copy(dAtA[i:], m.BlacklistedDenoms[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.BlacklistedDenoms[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.WhitelistedAddressPairs) > 0 {
		for iNdEx := len(m.WhitelistedAddressPairs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WhitelistedAddressPairs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.RateLimits) > 0 {
		for iNdEx := len(m.RateLimits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RateLimits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.RateLimits) > 0 {
		for _, e := range m.RateLimits {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.WhitelistedAddressPairs) > 0 {
		for _, e := range m.WhitelistedAddressPairs {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BlacklistedDenoms) > 0 {
		for _, s := range m.BlacklistedDenoms {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PendingSendPacketSequenceNumbers) > 0 {
		for _, s := range m.PendingSendPacketSequenceNumbers {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.HourEpoch.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RateLimits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RateLimits = append(m.RateLimits, RateLimit{})
			if err := m.RateLimits[len(m.RateLimits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WhitelistedAddressPairs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WhitelistedAddressPairs = append(m.WhitelistedAddressPairs, WhitelistedAddressPair{})
			if err := m.WhitelistedAddressPairs[len(m.WhitelistedAddressPairs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlacklistedDenoms", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlacklistedDenoms = append(m.BlacklistedDenoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingSendPacketSequenceNumbers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PendingSendPacketSequenceNumbers = append(m.PendingSendPacketSequenceNumbers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HourEpoch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.HourEpoch.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
