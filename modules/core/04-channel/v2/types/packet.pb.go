// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/core/channel/v2/packet.proto

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

// PacketStatus specifies the status of a RecvPacketResult.
type PacketStatus int32

const (
	// PACKET_STATUS_UNSPECIFIED indicates an unknown packet status.
	PacketStatus_NONE PacketStatus = 0
	// PACKET_STATUS_SUCCESS indicates a successful packet receipt.
	PacketStatus_Success PacketStatus = 1
	// PACKET_STATUS_FAILURE indicates a failed packet receipt.
	PacketStatus_Failure PacketStatus = 2
	// PACKET_STATUS_ASYNC indicates an async packet receipt.
	PacketStatus_Async PacketStatus = 3
)

var PacketStatus_name = map[int32]string{
	0: "PACKET_STATUS_UNSPECIFIED",
	1: "PACKET_STATUS_SUCCESS",
	2: "PACKET_STATUS_FAILURE",
	3: "PACKET_STATUS_ASYNC",
}

var PacketStatus_value = map[string]int32{
	"PACKET_STATUS_UNSPECIFIED": 0,
	"PACKET_STATUS_SUCCESS":     1,
	"PACKET_STATUS_FAILURE":     2,
	"PACKET_STATUS_ASYNC":       3,
}

func (x PacketStatus) String() string {
	return proto.EnumName(PacketStatus_name, int32(x))
}

func (PacketStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2f814aba9ca97169, []int{0}
}

// Packet defines a type that carries data across different chains through IBC
type Packet struct {
	// number corresponds to the order of sends and receives, where a Packet
	// with an earlier sequence number must be sent and received before a Packet
	// with a later sequence number.
	Sequence uint64 `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// identifies the sending chain.
	SourceChannel string `protobuf:"bytes,2,opt,name=source_channel,json=sourceChannel,proto3" json:"source_channel,omitempty"`
	// identifies the receiving chain.
	DestinationChannel string `protobuf:"bytes,3,opt,name=destination_channel,json=destinationChannel,proto3" json:"destination_channel,omitempty"`
	// timeout timestamp in seconds after which the packet times out.
	TimeoutTimestamp uint64 `protobuf:"varint,4,opt,name=timeout_timestamp,json=timeoutTimestamp,proto3" json:"timeout_timestamp,omitempty"`
	// a list of payloads, each one for a specific application.
	Payloads []Payload `protobuf:"bytes,5,rep,name=payloads,proto3" json:"payloads"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}
func (*Packet) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f814aba9ca97169, []int{0}
}
func (m *Packet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Packet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Packet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Packet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet.Merge(m, src)
}
func (m *Packet) XXX_Size() int {
	return m.Size()
}
func (m *Packet) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet.DiscardUnknown(m)
}

var xxx_messageInfo_Packet proto.InternalMessageInfo

func (m *Packet) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Packet) GetSourceChannel() string {
	if m != nil {
		return m.SourceChannel
	}
	return ""
}

func (m *Packet) GetDestinationChannel() string {
	if m != nil {
		return m.DestinationChannel
	}
	return ""
}

func (m *Packet) GetTimeoutTimestamp() uint64 {
	if m != nil {
		return m.TimeoutTimestamp
	}
	return 0
}

func (m *Packet) GetPayloads() []Payload {
	if m != nil {
		return m.Payloads
	}
	return nil
}

// Payload contains the source and destination ports and payload for the application (version, encoding, raw bytes)
type Payload struct {
	// specifies the source port of the packet.
	SourcePort string `protobuf:"bytes,1,opt,name=source_port,json=sourcePort,proto3" json:"source_port,omitempty"`
	// specifies the destination port of the packet.
	DestinationPort string `protobuf:"bytes,2,opt,name=destination_port,json=destinationPort,proto3" json:"destination_port,omitempty"`
	// version of the specified application.
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// the encoding used for the provided value.
	Encoding string `protobuf:"bytes,4,opt,name=encoding,proto3" json:"encoding,omitempty"`
	// the raw bytes for the payload.
	Value []byte `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f814aba9ca97169, []int{1}
}
func (m *Payload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Payload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payload.Merge(m, src)
}
func (m *Payload) XXX_Size() int {
	return m.Size()
}
func (m *Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Payload proto.InternalMessageInfo

func (m *Payload) GetSourcePort() string {
	if m != nil {
		return m.SourcePort
	}
	return ""
}

func (m *Payload) GetDestinationPort() string {
	if m != nil {
		return m.DestinationPort
	}
	return ""
}

func (m *Payload) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Payload) GetEncoding() string {
	if m != nil {
		return m.Encoding
	}
	return ""
}

func (m *Payload) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// Acknowledgement contains a list of all ack results associated with a single packet.
type Acknowledgement struct {
	RecvSuccess         bool     `protobuf:"varint,1,opt,name=recv_success,json=recvSuccess,proto3" json:"recv_success,omitempty"`
	AppAcknowledgements [][]byte `protobuf:"bytes,2,rep,name=app_acknowledgements,json=appAcknowledgements,proto3" json:"app_acknowledgements,omitempty"`
}

func (m *Acknowledgement) Reset()         { *m = Acknowledgement{} }
func (m *Acknowledgement) String() string { return proto.CompactTextString(m) }
func (*Acknowledgement) ProtoMessage()    {}
func (*Acknowledgement) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f814aba9ca97169, []int{2}
}
func (m *Acknowledgement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Acknowledgement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Acknowledgement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Acknowledgement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Acknowledgement.Merge(m, src)
}
func (m *Acknowledgement) XXX_Size() int {
	return m.Size()
}
func (m *Acknowledgement) XXX_DiscardUnknown() {
	xxx_messageInfo_Acknowledgement.DiscardUnknown(m)
}

var xxx_messageInfo_Acknowledgement proto.InternalMessageInfo

func (m *Acknowledgement) GetRecvSuccess() bool {
	if m != nil {
		return m.RecvSuccess
	}
	return false
}

func (m *Acknowledgement) GetAppAcknowledgements() [][]byte {
	if m != nil {
		return m.AppAcknowledgements
	}
	return nil
}

// RecvPacketResult speecifies the status of a packet as well as the acknowledgement bytes.
type RecvPacketResult struct {
	// status of the packet
	Status PacketStatus `protobuf:"varint,1,opt,name=status,proto3,enum=ibc.core.channel.v2.PacketStatus" json:"status,omitempty"`
	// acknowledgement of the packet
	Acknowledgement []byte `protobuf:"bytes,2,opt,name=acknowledgement,proto3" json:"acknowledgement,omitempty"`
}

func (m *RecvPacketResult) Reset()         { *m = RecvPacketResult{} }
func (m *RecvPacketResult) String() string { return proto.CompactTextString(m) }
func (*RecvPacketResult) ProtoMessage()    {}
func (*RecvPacketResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f814aba9ca97169, []int{3}
}
func (m *RecvPacketResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RecvPacketResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RecvPacketResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RecvPacketResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecvPacketResult.Merge(m, src)
}
func (m *RecvPacketResult) XXX_Size() int {
	return m.Size()
}
func (m *RecvPacketResult) XXX_DiscardUnknown() {
	xxx_messageInfo_RecvPacketResult.DiscardUnknown(m)
}

var xxx_messageInfo_RecvPacketResult proto.InternalMessageInfo

func (m *RecvPacketResult) GetStatus() PacketStatus {
	if m != nil {
		return m.Status
	}
	return PacketStatus_NONE
}

func (m *RecvPacketResult) GetAcknowledgement() []byte {
	if m != nil {
		return m.Acknowledgement
	}
	return nil
}

func init() {
	proto.RegisterEnum("ibc.core.channel.v2.PacketStatus", PacketStatus_name, PacketStatus_value)
	proto.RegisterType((*Packet)(nil), "ibc.core.channel.v2.Packet")
	proto.RegisterType((*Payload)(nil), "ibc.core.channel.v2.Payload")
	proto.RegisterType((*Acknowledgement)(nil), "ibc.core.channel.v2.Acknowledgement")
	proto.RegisterType((*RecvPacketResult)(nil), "ibc.core.channel.v2.RecvPacketResult")
}

func init() { proto.RegisterFile("ibc/core/channel/v2/packet.proto", fileDescriptor_2f814aba9ca97169) }

var fileDescriptor_2f814aba9ca97169 = []byte{
	// 605 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0x9b, 0xb5, 0xdd, 0x5a, 0xb7, 0x6c, 0xc1, 0x1d, 0x52, 0xa8, 0x50, 0x97, 0x55, 0x02,
	0x0a, 0x68, 0x09, 0x14, 0x2e, 0x93, 0x10, 0x52, 0x57, 0x3a, 0x69, 0x02, 0x95, 0xca, 0x69, 0x91,
	0xe0, 0x52, 0xb9, 0xae, 0x95, 0x45, 0x4b, 0xe3, 0x10, 0x3b, 0x99, 0xf6, 0x0a, 0x3b, 0xf1, 0x02,
	0x3b, 0x70, 0xe6, 0x45, 0x76, 0xdc, 0x91, 0x13, 0x42, 0xdb, 0x89, 0xb7, 0x40, 0xb1, 0xd3, 0xaa,
	0x2b, 0x70, 0x4a, 0xbe, 0xff, 0xf7, 0xfb, 0xdb, 0xf9, 0x7f, 0xd1, 0x07, 0x4c, 0x6f, 0x42, 0x6c,
	0xc2, 0x22, 0x6a, 0x93, 0x63, 0x1c, 0x04, 0xd4, 0xb7, 0x93, 0xb6, 0x1d, 0x62, 0x72, 0x42, 0x85,
	0x15, 0x46, 0x4c, 0x30, 0x58, 0xf3, 0x26, 0xc4, 0x4a, 0x09, 0x2b, 0x23, 0xac, 0xa4, 0x5d, 0xdf,
	0x76, 0x99, 0xcb, 0x64, 0xdf, 0x4e, 0xdf, 0x14, 0xda, 0xfc, 0xad, 0x81, 0xf5, 0x81, 0xf4, 0xc2,
	0x3a, 0x28, 0x71, 0xfa, 0x25, 0xa6, 0x01, 0xa1, 0x86, 0x66, 0x6a, 0xad, 0x02, 0x5a, 0xd4, 0xf0,
	0x21, 0xd8, 0xe4, 0x2c, 0x8e, 0x08, 0x1d, 0x67, 0x27, 0x1a, 0x6b, 0xa6, 0xd6, 0x2a, 0xa3, 0x3b,
	0x4a, 0xed, 0x2a, 0x11, 0xda, 0xa0, 0x36, 0xa5, 0x5c, 0x78, 0x01, 0x16, 0x1e, 0x0b, 0x16, 0x6c,
	0x5e, 0xb2, 0x70, 0xa9, 0x35, 0x37, 0x3c, 0x03, 0x77, 0x85, 0x37, 0xa3, 0x2c, 0x16, 0xe3, 0xf4,
	0xc9, 0x05, 0x9e, 0x85, 0x46, 0x41, 0x5e, 0xae, 0x67, 0x8d, 0xe1, 0x5c, 0x87, 0x6f, 0x40, 0x29,
	0xc4, 0x67, 0x3e, 0xc3, 0x53, 0x6e, 0x14, 0xcd, 0x7c, 0xab, 0xd2, 0x7e, 0x60, 0xfd, 0x23, 0xa9,
	0x35, 0x50, 0xd0, 0x41, 0xe1, 0xf2, 0xe7, 0x4e, 0x0e, 0x2d, 0x3c, 0xcd, 0x6f, 0x1a, 0xd8, 0xc8,
	0x7a, 0x70, 0x07, 0x54, 0xb2, 0x40, 0x21, 0x8b, 0x84, 0xcc, 0x5b, 0x46, 0x40, 0x49, 0x03, 0x16,
	0x09, 0xf8, 0x04, 0xe8, 0xcb, 0x51, 0x24, 0xa5, 0x32, 0x6f, 0x2d, 0xe9, 0x12, 0x35, 0xc0, 0x46,
	0x42, 0x23, 0xee, 0xb1, 0x20, 0x4b, 0x3a, 0x2f, 0xd3, 0x91, 0xd2, 0x80, 0xb0, 0xa9, 0x17, 0xb8,
	0x32, 0x55, 0x19, 0x2d, 0x6a, 0xb8, 0x0d, 0x8a, 0x09, 0xf6, 0x63, 0x6a, 0x14, 0x4d, 0xad, 0x55,
	0x45, 0xaa, 0x68, 0xba, 0x60, 0xab, 0x43, 0x4e, 0x02, 0x76, 0xea, 0xd3, 0xa9, 0x4b, 0x67, 0x34,
	0x10, 0x70, 0x17, 0x54, 0x23, 0x4a, 0x92, 0x31, 0x8f, 0x09, 0xa1, 0x9c, 0xcb, 0x6f, 0x2d, 0xa1,
	0x4a, 0xaa, 0x39, 0x4a, 0x82, 0x2f, 0xc0, 0x36, 0x0e, 0xc3, 0x31, 0xbe, 0xed, 0xe4, 0xc6, 0x9a,
	0x99, 0x6f, 0x55, 0x51, 0x0d, 0x87, 0xe1, 0xca, 0xa1, 0xbc, 0x79, 0x0a, 0x74, 0x44, 0x49, 0xa2,
	0xfe, 0x3d, 0xa2, 0x3c, 0xf6, 0x05, 0xdc, 0x07, 0xeb, 0x5c, 0x60, 0x11, 0xab, 0x3b, 0x36, 0xdb,
	0xbb, 0xff, 0x19, 0x6f, 0x6a, 0x71, 0x24, 0x88, 0x32, 0x03, 0x6c, 0x81, 0xad, 0x95, 0xdb, 0xe5,
	0xb4, 0xaa, 0x68, 0x55, 0x7e, 0xfa, 0x5d, 0x03, 0xd5, 0xe5, 0x23, 0xe0, 0x63, 0x70, 0x7f, 0xd0,
	0xe9, 0xbe, 0xeb, 0x0d, 0xc7, 0xce, 0xb0, 0x33, 0x1c, 0x39, 0xe3, 0x51, 0xdf, 0x19, 0xf4, 0xba,
	0x47, 0x87, 0x47, 0xbd, 0xb7, 0x7a, 0xae, 0x5e, 0x3a, 0xbf, 0x30, 0x0b, 0xfd, 0x0f, 0xfd, 0x1e,
	0x7c, 0x04, 0xee, 0xdd, 0x06, 0x9d, 0x51, 0xb7, 0xdb, 0x73, 0x1c, 0x5d, 0xab, 0x57, 0xce, 0x2f,
	0xcc, 0x8d, 0xf9, 0x34, 0xfe, 0xe2, 0x0e, 0x3b, 0x47, 0xef, 0x47, 0xa8, 0xa7, 0xaf, 0x29, 0xee,
	0x10, 0x7b, 0x7e, 0x1c, 0x51, 0xd8, 0x04, 0xb5, 0xdb, 0x5c, 0xc7, 0xf9, 0xd4, 0xef, 0xea, 0xf9,
	0x7a, 0xf9, 0xfc, 0xc2, 0x2c, 0x76, 0xf8, 0x59, 0x40, 0x0e, 0x3e, 0x5e, 0x5e, 0x37, 0xb4, 0xab,
	0xeb, 0x86, 0xf6, 0xeb, 0xba, 0xa1, 0x7d, 0xbd, 0x69, 0xe4, 0xae, 0x6e, 0x1a, 0xb9, 0x1f, 0x37,
	0x8d, 0xdc, 0xe7, 0xd7, 0xae, 0x27, 0x8e, 0xe3, 0x89, 0x45, 0xd8, 0xcc, 0x26, 0x8c, 0xcf, 0x18,
	0xb7, 0xbd, 0x09, 0xd9, 0x73, 0x99, 0x9d, 0xec, 0xdb, 0x33, 0x36, 0x8d, 0x7d, 0xca, 0xd5, 0x9a,
	0x3e, 0x7f, 0xb5, 0xb7, 0xb4, 0xa9, 0xe2, 0x2c, 0xa4, 0x7c, 0xb2, 0x2e, 0xd7, 0xef, 0xe5, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x96, 0x41, 0x25, 0xe6, 0xcd, 0x03, 0x00, 0x00,
}

func (m *Packet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Packet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Packet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Payloads) > 0 {
		for iNdEx := len(m.Payloads) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Payloads[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPacket(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x20
	}
	if len(m.DestinationChannel) > 0 {
		i -= len(m.DestinationChannel)
		copy(dAtA[i:], m.DestinationChannel)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.DestinationChannel)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SourceChannel) > 0 {
		i -= len(m.SourceChannel)
		copy(dAtA[i:], m.SourceChannel)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.SourceChannel)))
		i--
		dAtA[i] = 0x12
	}
	if m.Sequence != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Payload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Payload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Payload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Encoding) > 0 {
		i -= len(m.Encoding)
		copy(dAtA[i:], m.Encoding)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Encoding)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DestinationPort) > 0 {
		i -= len(m.DestinationPort)
		copy(dAtA[i:], m.DestinationPort)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.DestinationPort)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SourcePort) > 0 {
		i -= len(m.SourcePort)
		copy(dAtA[i:], m.SourcePort)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.SourcePort)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Acknowledgement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Acknowledgement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Acknowledgement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AppAcknowledgements) > 0 {
		for iNdEx := len(m.AppAcknowledgements) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AppAcknowledgements[iNdEx])
			copy(dAtA[i:], m.AppAcknowledgements[iNdEx])
			i = encodeVarintPacket(dAtA, i, uint64(len(m.AppAcknowledgements[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.RecvSuccess {
		i--
		if m.RecvSuccess {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RecvPacketResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RecvPacketResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RecvPacketResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Acknowledgement) > 0 {
		i -= len(m.Acknowledgement)
		copy(dAtA[i:], m.Acknowledgement)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Acknowledgement)))
		i--
		dAtA[i] = 0x12
	}
	if m.Status != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPacket(dAtA []byte, offset int, v uint64) int {
	offset -= sovPacket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Packet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sequence != 0 {
		n += 1 + sovPacket(uint64(m.Sequence))
	}
	l = len(m.SourceChannel)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.DestinationChannel)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovPacket(uint64(m.TimeoutTimestamp))
	}
	if len(m.Payloads) > 0 {
		for _, e := range m.Payloads {
			l = e.Size()
			n += 1 + l + sovPacket(uint64(l))
		}
	}
	return n
}

func (m *Payload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SourcePort)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.DestinationPort)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.Encoding)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func (m *Acknowledgement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RecvSuccess {
		n += 2
	}
	if len(m.AppAcknowledgements) > 0 {
		for _, b := range m.AppAcknowledgements {
			l = len(b)
			n += 1 + l + sovPacket(uint64(l))
		}
	}
	return n
}

func (m *RecvPacketResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovPacket(uint64(m.Status))
	}
	l = len(m.Acknowledgement)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func sovPacket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPacket(x uint64) (n int) {
	return sovPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Packet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: Packet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Packet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChannel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationChannel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payloads", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payloads = append(m.Payloads, Payload{})
			if err := m.Payloads[len(m.Payloads)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *Payload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: Payload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Payload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourcePort", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourcePort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationPort", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationPort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Encoding", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Encoding = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *Acknowledgement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: Acknowledgement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Acknowledgement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecvSuccess", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.RecvSuccess = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppAcknowledgements", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppAcknowledgements = append(m.AppAcknowledgements, make([]byte, postIndex-iNdEx))
			copy(m.AppAcknowledgements[len(m.AppAcknowledgements)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *RecvPacketResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: RecvPacketResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RecvPacketResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= PacketStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Acknowledgement", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Acknowledgement = append(m.Acknowledgement[:0], dAtA[iNdEx:postIndex]...)
			if m.Acknowledgement == nil {
				m.Acknowledgement = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func skipPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
				return 0, ErrInvalidLengthPacket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPacket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPacket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPacket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPacket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPacket = fmt.Errorf("proto: unexpected end of group")
)
