// Code generated by protoc-gen-go. DO NOT EDIT.
// source: worldcomm.proto

package worldcomm

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MessageType int32

const (
	MessageType_UNKNOWN_MESSAGE_TYPE            MessageType = 0
	MessageType_POSITION                        MessageType = 2
	MessageType_CHAT                            MessageType = 3
	MessageType_CLIENT_DISCONNECTED_FROM_SERVER MessageType = 4
	MessageType_PROFILE                         MessageType = 5
	MessageType_CLOCK_SKEW_DETECTED             MessageType = 6
	MessageType_FLOW_STATUS                     MessageType = 7
	MessageType_PING                            MessageType = 8
)

var MessageType_name = map[int32]string{
	0: "UNKNOWN_MESSAGE_TYPE",
	2: "POSITION",
	3: "CHAT",
	4: "CLIENT_DISCONNECTED_FROM_SERVER",
	5: "PROFILE",
	6: "CLOCK_SKEW_DETECTED",
	7: "FLOW_STATUS",
	8: "PING",
}

var MessageType_value = map[string]int32{
	"UNKNOWN_MESSAGE_TYPE":            0,
	"POSITION":                        2,
	"CHAT":                            3,
	"CLIENT_DISCONNECTED_FROM_SERVER": 4,
	"PROFILE":                         5,
	"CLOCK_SKEW_DETECTED":             6,
	"FLOW_STATUS":                     7,
	"PING":                            8,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}

func (MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_966603bf5abd3632, []int{0}
}

type FlowStatus int32

const (
	FlowStatus_UNKNOWN_STATUS FlowStatus = 0
	FlowStatus_OPEN           FlowStatus = 1
	FlowStatus_CLOSE          FlowStatus = 2
)

var FlowStatus_name = map[int32]string{
	0: "UNKNOWN_STATUS",
	1: "OPEN",
	2: "CLOSE",
}

var FlowStatus_value = map[string]int32{
	"UNKNOWN_STATUS": 0,
	"OPEN":           1,
	"CLOSE":          2,
}

func (x FlowStatus) String() string {
	return proto.EnumName(FlowStatus_name, int32(x))
}

func (FlowStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_966603bf5abd3632, []int{1}
}

type GenericMessage struct {
	Type                 MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=MessageType" json:"type,omitempty"`
	Time                 float64     `protobuf:"fixed64,2,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GenericMessage) Reset()         { *m = GenericMessage{} }
func (m *GenericMessage) String() string { return proto.CompactTextString(m) }
func (*GenericMessage) ProtoMessage()    {}
func (*GenericMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_966603bf5abd3632, []int{0}
}

func (m *GenericMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenericMessage.Unmarshal(m, b)
}
func (m *GenericMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenericMessage.Marshal(b, m, deterministic)
}
func (m *GenericMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericMessage.Merge(m, src)
}
func (m *GenericMessage) XXX_Size() int {
	return xxx_messageInfo_GenericMessage.Size(m)
}
func (m *GenericMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericMessage.DiscardUnknown(m)
}

var xxx_messageInfo_GenericMessage proto.InternalMessageInfo

func (m *GenericMessage) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_UNKNOWN_MESSAGE_TYPE
}

func (m *GenericMessage) GetTime() float64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type PositionMessage struct {
	Type                 MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=MessageType" json:"type,omitempty"`
	Time                 float64     `protobuf:"fixed64,2,opt,name=time,proto3" json:"time,omitempty"`
	PositionX            float32     `protobuf:"fixed32,3,opt,name=position_x,json=positionX,proto3" json:"position_x,omitempty"`
	PositionY            float32     `protobuf:"fixed32,4,opt,name=position_y,json=positionY,proto3" json:"position_y,omitempty"`
	PositionZ            float32     `protobuf:"fixed32,5,opt,name=position_z,json=positionZ,proto3" json:"position_z,omitempty"`
	RotationX            float32     `protobuf:"fixed32,6,opt,name=rotation_x,json=rotationX,proto3" json:"rotation_x,omitempty"`
	RotationY            float32     `protobuf:"fixed32,7,opt,name=rotation_y,json=rotationY,proto3" json:"rotation_y,omitempty"`
	RotationZ            float32     `protobuf:"fixed32,8,opt,name=rotation_z,json=rotationZ,proto3" json:"rotation_z,omitempty"`
	RotationW            float32     `protobuf:"fixed32,9,opt,name=rotation_w,json=rotationW,proto3" json:"rotation_w,omitempty"`
	Alias                uint32      `protobuf:"varint,10,opt,name=alias,proto3" json:"alias,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PositionMessage) Reset()         { *m = PositionMessage{} }
func (m *PositionMessage) String() string { return proto.CompactTextString(m) }
func (*PositionMessage) ProtoMessage()    {}
func (*PositionMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_966603bf5abd3632, []int{1}
}

func (m *PositionMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PositionMessage.Unmarshal(m, b)
}
func (m *PositionMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PositionMessage.Marshal(b, m, deterministic)
}
func (m *PositionMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PositionMessage.Merge(m, src)
}
func (m *PositionMessage) XXX_Size() int {
	return xxx_messageInfo_PositionMessage.Size(m)
}
func (m *PositionMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PositionMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PositionMessage proto.InternalMessageInfo

func (m *PositionMessage) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_UNKNOWN_MESSAGE_TYPE
}

func (m *PositionMessage) GetTime() float64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *PositionMessage) GetPositionX() float32 {
	if m != nil {
		return m.PositionX
	}
	return 0
}

func (m *PositionMessage) GetPositionY() float32 {
	if m != nil {
		return m.PositionY
	}
	return 0
}

func (m *PositionMessage) GetPositionZ() float32 {
	if m != nil {
		return m.PositionZ
	}
	return 0
}

func (m *PositionMessage) GetRotationX() float32 {
	if m != nil {
		return m.RotationX
	}
	return 0
}

func (m *PositionMessage) GetRotationY() float32 {
	if m != nil {
		return m.RotationY
	}
	return 0
}

func (m *PositionMessage) GetRotationZ() float32 {
	if m != nil {
		return m.RotationZ
	}
	return 0
}

func (m *PositionMessage) GetRotationW() float32 {
	if m != nil {
		return m.RotationW
	}
	return 0
}

func (m *PositionMessage) GetAlias() uint32 {
	if m != nil {
		return m.Alias
	}
	return 0
}

type ProfileMessage struct {
	Type                 MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=MessageType" json:"type,omitempty"`
	Time                 float64     `protobuf:"fixed64,2,opt,name=time,proto3" json:"time,omitempty"`
	PositionX            float32     `protobuf:"fixed32,3,opt,name=position_x,json=positionX,proto3" json:"position_x,omitempty"`
	PositionZ            float32     `protobuf:"fixed32,4,opt,name=position_z,json=positionZ,proto3" json:"position_z,omitempty"`
	AvatarType           string      `protobuf:"bytes,5,opt,name=avatar_type,json=avatarType,proto3" json:"avatar_type,omitempty"`
	DisplayName          string      `protobuf:"bytes,6,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	PeerId               string      `protobuf:"bytes,7,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Alias                uint32      `protobuf:"varint,8,opt,name=alias,proto3" json:"alias,omitempty"`
	PublicKey            string      `protobuf:"bytes,9,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ProfileMessage) Reset()         { *m = ProfileMessage{} }
func (m *ProfileMessage) String() string { return proto.CompactTextString(m) }
func (*ProfileMessage) ProtoMessage()    {}
func (*ProfileMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_966603bf5abd3632, []int{2}
}

func (m *ProfileMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileMessage.Unmarshal(m, b)
}
func (m *ProfileMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileMessage.Marshal(b, m, deterministic)
}
func (m *ProfileMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileMessage.Merge(m, src)
}
func (m *ProfileMessage) XXX_Size() int {
	return xxx_messageInfo_ProfileMessage.Size(m)
}
func (m *ProfileMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileMessage proto.InternalMessageInfo

func (m *ProfileMessage) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_UNKNOWN_MESSAGE_TYPE
}

func (m *ProfileMessage) GetTime() float64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *ProfileMessage) GetPositionX() float32 {
	if m != nil {
		return m.PositionX
	}
	return 0
}

func (m *ProfileMessage) GetPositionZ() float32 {
	if m != nil {
		return m.PositionZ
	}
	return 0
}

func (m *ProfileMessage) GetAvatarType() string {
	if m != nil {
		return m.AvatarType
	}
	return ""
}

func (m *ProfileMessage) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *ProfileMessage) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *ProfileMessage) GetAlias() uint32 {
	if m != nil {
		return m.Alias
	}
	return 0
}

func (m *ProfileMessage) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

type ChatMessage struct {
	Type                 MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=MessageType" json:"type,omitempty"`
	Time                 float64     `protobuf:"fixed64,2,opt,name=time,proto3" json:"time,omitempty"`
	MessageId            string      `protobuf:"bytes,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	PositionX            float32     `protobuf:"fixed32,4,opt,name=position_x,json=positionX,proto3" json:"position_x,omitempty"`
	PositionZ            float32     `protobuf:"fixed32,5,opt,name=position_z,json=positionZ,proto3" json:"position_z,omitempty"`
	Text                 string      `protobuf:"bytes,6,opt,name=text,proto3" json:"text,omitempty"`
	Alias                uint32      `protobuf:"varint,7,opt,name=alias,proto3" json:"alias,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ChatMessage) Reset()         { *m = ChatMessage{} }
func (m *ChatMessage) String() string { return proto.CompactTextString(m) }
func (*ChatMessage) ProtoMessage()    {}
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_966603bf5abd3632, []int{3}
}

func (m *ChatMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatMessage.Unmarshal(m, b)
}
func (m *ChatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatMessage.Marshal(b, m, deterministic)
}
func (m *ChatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatMessage.Merge(m, src)
}
func (m *ChatMessage) XXX_Size() int {
	return xxx_messageInfo_ChatMessage.Size(m)
}
func (m *ChatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ChatMessage proto.InternalMessageInfo

func (m *ChatMessage) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_UNKNOWN_MESSAGE_TYPE
}

func (m *ChatMessage) GetTime() float64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *ChatMessage) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *ChatMessage) GetPositionX() float32 {
	if m != nil {
		return m.PositionX
	}
	return 0
}

func (m *ChatMessage) GetPositionZ() float32 {
	if m != nil {
		return m.PositionZ
	}
	return 0
}

func (m *ChatMessage) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *ChatMessage) GetAlias() uint32 {
	if m != nil {
		return m.Alias
	}
	return 0
}

type ClientDisconnectedFromServerMessage struct {
	Type                 MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=MessageType" json:"type,omitempty"`
	Time                 float64     `protobuf:"fixed64,2,opt,name=time,proto3" json:"time,omitempty"`
	Alias                uint32      `protobuf:"varint,3,opt,name=alias,proto3" json:"alias,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ClientDisconnectedFromServerMessage) Reset()         { *m = ClientDisconnectedFromServerMessage{} }
func (m *ClientDisconnectedFromServerMessage) String() string { return proto.CompactTextString(m) }
func (*ClientDisconnectedFromServerMessage) ProtoMessage()    {}
func (*ClientDisconnectedFromServerMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_966603bf5abd3632, []int{4}
}

func (m *ClientDisconnectedFromServerMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientDisconnectedFromServerMessage.Unmarshal(m, b)
}
func (m *ClientDisconnectedFromServerMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientDisconnectedFromServerMessage.Marshal(b, m, deterministic)
}
func (m *ClientDisconnectedFromServerMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientDisconnectedFromServerMessage.Merge(m, src)
}
func (m *ClientDisconnectedFromServerMessage) XXX_Size() int {
	return xxx_messageInfo_ClientDisconnectedFromServerMessage.Size(m)
}
func (m *ClientDisconnectedFromServerMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientDisconnectedFromServerMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ClientDisconnectedFromServerMessage proto.InternalMessageInfo

func (m *ClientDisconnectedFromServerMessage) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_UNKNOWN_MESSAGE_TYPE
}

func (m *ClientDisconnectedFromServerMessage) GetTime() float64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *ClientDisconnectedFromServerMessage) GetAlias() uint32 {
	if m != nil {
		return m.Alias
	}
	return 0
}

type ClockSkewMessage struct {
	Type                 MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=MessageType" json:"type,omitempty"`
	Time                 float64     `protobuf:"fixed64,2,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ClockSkewMessage) Reset()         { *m = ClockSkewMessage{} }
func (m *ClockSkewMessage) String() string { return proto.CompactTextString(m) }
func (*ClockSkewMessage) ProtoMessage()    {}
func (*ClockSkewMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_966603bf5abd3632, []int{5}
}

func (m *ClockSkewMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClockSkewMessage.Unmarshal(m, b)
}
func (m *ClockSkewMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClockSkewMessage.Marshal(b, m, deterministic)
}
func (m *ClockSkewMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClockSkewMessage.Merge(m, src)
}
func (m *ClockSkewMessage) XXX_Size() int {
	return xxx_messageInfo_ClockSkewMessage.Size(m)
}
func (m *ClockSkewMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ClockSkewMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ClockSkewMessage proto.InternalMessageInfo

func (m *ClockSkewMessage) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_UNKNOWN_MESSAGE_TYPE
}

func (m *ClockSkewMessage) GetTime() float64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type FlowStatusMessage struct {
	Type                 MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=MessageType" json:"type,omitempty"`
	Time                 float64     `protobuf:"fixed64,2,opt,name=time,proto3" json:"time,omitempty"`
	FlowStatus           FlowStatus  `protobuf:"varint,3,opt,name=flow_status,json=flowStatus,proto3,enum=FlowStatus" json:"flow_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *FlowStatusMessage) Reset()         { *m = FlowStatusMessage{} }
func (m *FlowStatusMessage) String() string { return proto.CompactTextString(m) }
func (*FlowStatusMessage) ProtoMessage()    {}
func (*FlowStatusMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_966603bf5abd3632, []int{6}
}

func (m *FlowStatusMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowStatusMessage.Unmarshal(m, b)
}
func (m *FlowStatusMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowStatusMessage.Marshal(b, m, deterministic)
}
func (m *FlowStatusMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowStatusMessage.Merge(m, src)
}
func (m *FlowStatusMessage) XXX_Size() int {
	return xxx_messageInfo_FlowStatusMessage.Size(m)
}
func (m *FlowStatusMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowStatusMessage.DiscardUnknown(m)
}

var xxx_messageInfo_FlowStatusMessage proto.InternalMessageInfo

func (m *FlowStatusMessage) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_UNKNOWN_MESSAGE_TYPE
}

func (m *FlowStatusMessage) GetTime() float64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *FlowStatusMessage) GetFlowStatus() FlowStatus {
	if m != nil {
		return m.FlowStatus
	}
	return FlowStatus_UNKNOWN_STATUS
}

type PingMessage struct {
	Type                 MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=MessageType" json:"type,omitempty"`
	Time                 float64     `protobuf:"fixed64,2,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PingMessage) Reset()         { *m = PingMessage{} }
func (m *PingMessage) String() string { return proto.CompactTextString(m) }
func (*PingMessage) ProtoMessage()    {}
func (*PingMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_966603bf5abd3632, []int{7}
}

func (m *PingMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingMessage.Unmarshal(m, b)
}
func (m *PingMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingMessage.Marshal(b, m, deterministic)
}
func (m *PingMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingMessage.Merge(m, src)
}
func (m *PingMessage) XXX_Size() int {
	return xxx_messageInfo_PingMessage.Size(m)
}
func (m *PingMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PingMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PingMessage proto.InternalMessageInfo

func (m *PingMessage) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_UNKNOWN_MESSAGE_TYPE
}

func (m *PingMessage) GetTime() float64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func init() {
	proto.RegisterEnum("MessageType", MessageType_name, MessageType_value)
	proto.RegisterEnum("FlowStatus", FlowStatus_name, FlowStatus_value)
	proto.RegisterType((*GenericMessage)(nil), "GenericMessage")
	proto.RegisterType((*PositionMessage)(nil), "PositionMessage")
	proto.RegisterType((*ProfileMessage)(nil), "ProfileMessage")
	proto.RegisterType((*ChatMessage)(nil), "ChatMessage")
	proto.RegisterType((*ClientDisconnectedFromServerMessage)(nil), "ClientDisconnectedFromServerMessage")
	proto.RegisterType((*ClockSkewMessage)(nil), "ClockSkewMessage")
	proto.RegisterType((*FlowStatusMessage)(nil), "FlowStatusMessage")
	proto.RegisterType((*PingMessage)(nil), "PingMessage")
}

func init() { proto.RegisterFile("worldcomm.proto", fileDescriptor_966603bf5abd3632) }

var fileDescriptor_966603bf5abd3632 = []byte{
	// 615 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0x5f, 0x4e, 0xdb, 0x4c,
	0x14, 0xc5, 0x71, 0xfe, 0xfb, 0x9a, 0x2f, 0xf8, 0x9b, 0x22, 0xe1, 0x97, 0x8a, 0x34, 0xbc, 0x44,
	0xa8, 0xe2, 0x81, 0xaa, 0x0b, 0x40, 0xce, 0x04, 0xac, 0x04, 0xdb, 0x1a, 0x9b, 0x06, 0x78, 0x19,
	0x99, 0x64, 0xa0, 0x16, 0xb6, 0x27, 0xb5, 0x0d, 0xc1, 0x2c, 0xa4, 0x3b, 0xe8, 0x16, 0xba, 0x88,
	0xae, 0xaa, 0xf2, 0xc4, 0x90, 0x98, 0x4a, 0x7d, 0xb1, 0xd4, 0xb7, 0x99, 0xf3, 0xbb, 0xb9, 0xe7,
	0xde, 0x13, 0xdb, 0xb0, 0xb3, 0xe4, 0x71, 0x30, 0x9f, 0xf1, 0x30, 0x3c, 0x5a, 0xc4, 0x3c, 0xe5,
	0xfd, 0x11, 0x74, 0x4f, 0x59, 0xc4, 0x62, 0x7f, 0x76, 0xce, 0x92, 0xc4, 0xbb, 0x63, 0xa8, 0x07,
	0x8d, 0x34, 0x5b, 0x30, 0x4d, 0xea, 0x49, 0x83, 0xee, 0xf1, 0xf6, 0x51, 0xa1, 0xbb, 0xd9, 0x82,
	0x11, 0x41, 0x10, 0x82, 0x46, 0xea, 0x87, 0x4c, 0xab, 0xf5, 0xa4, 0x81, 0x44, 0xc4, 0xb9, 0xff,
	0xb3, 0x06, 0x3b, 0x36, 0x4f, 0xfc, 0xd4, 0xe7, 0x51, 0xa5, 0x4e, 0xe8, 0x3d, 0xc0, 0xa2, 0x68,
	0x44, 0x9f, 0xb4, 0x7a, 0x4f, 0x1a, 0xd4, 0x88, 0xfc, 0xa2, 0x5c, 0x96, 0x70, 0xa6, 0x35, 0xca,
	0xf8, 0xaa, 0x84, 0x9f, 0xb5, 0x66, 0x19, 0x5f, 0xe7, 0x38, 0xe6, 0xa9, 0x57, 0x34, 0x6f, 0xad,
	0xf0, 0x8b, 0x72, 0x59, 0xc2, 0x99, 0xd6, 0x2e, 0xe3, 0xab, 0x12, 0x7e, 0xd6, 0x3a, 0x65, 0x5c,
	0x6e, 0xbe, 0xd4, 0xe4, 0x32, 0x9e, 0xa2, 0x5d, 0x68, 0x7a, 0x81, 0xef, 0x25, 0x1a, 0xf4, 0xa4,
	0xc1, 0x7f, 0x64, 0x75, 0xe9, 0x7f, 0xaf, 0x41, 0xd7, 0x8e, 0xf9, 0xad, 0x1f, 0xb0, 0x7f, 0x96,
	0xdb, 0xf3, 0xdb, 0xdc, 0xae, 0xd1, 0x3e, 0x28, 0xde, 0xa3, 0x97, 0x7a, 0x31, 0x15, 0xd6, 0x79,
	0x70, 0x32, 0x81, 0x95, 0x94, 0x1b, 0xa3, 0x0f, 0xb0, 0x3d, 0xf7, 0x93, 0x45, 0xe0, 0x65, 0x34,
	0xf2, 0x42, 0x26, 0xb2, 0x93, 0x89, 0x52, 0x68, 0xa6, 0x17, 0x32, 0xb4, 0x07, 0xed, 0x05, 0x63,
	0x31, 0xf5, 0xe7, 0x22, 0x3a, 0x99, 0xb4, 0xf2, 0xab, 0x31, 0x5f, 0x6f, 0xde, 0xd9, 0xd8, 0x5c,
	0x4c, 0xf4, 0x70, 0x13, 0xf8, 0x33, 0x7a, 0xcf, 0x32, 0x11, 0x97, 0x4c, 0xe4, 0x95, 0x32, 0x66,
	0x59, 0xff, 0x97, 0x04, 0x8a, 0xfe, 0xd5, 0x4b, 0x2b, 0xa7, 0x12, 0xae, 0x0a, 0xf3, 0xb1, 0xea,
	0x2b, 0x93, 0x42, 0x31, 0xe6, 0x6f, 0x42, 0x6b, 0xfc, 0x3d, 0xb4, 0x3f, 0x9e, 0xa6, 0xdc, 0x90,
	0x3d, 0xa5, 0x45, 0x16, 0xe2, 0xbc, 0xde, 0xb5, 0xbd, 0xf9, 0x2f, 0x7f, 0x83, 0x03, 0x3d, 0xf0,
	0x59, 0x94, 0x0e, 0xfd, 0x64, 0xc6, 0xa3, 0x88, 0xcd, 0x52, 0x36, 0x1f, 0xc5, 0x3c, 0x74, 0x58,
	0xfc, 0xc8, 0xe2, 0x6a, 0x3b, 0xbe, 0x5a, 0xd6, 0x37, 0x2d, 0xcf, 0x40, 0xd5, 0x03, 0x3e, 0xbb,
	0x77, 0xee, 0xd9, 0xb2, 0xda, 0xbb, 0xbd, 0x84, 0xff, 0x47, 0x01, 0x5f, 0x3a, 0xa9, 0x97, 0x3e,
	0x24, 0xd5, 0x46, 0xfd, 0x08, 0xca, 0x6d, 0xc0, 0x97, 0x34, 0x11, 0xbd, 0xc4, 0xc0, 0xdd, 0x63,
	0xe5, 0x68, 0xdd, 0x9e, 0xc0, 0xed, 0xeb, 0xb9, 0xaf, 0x83, 0x62, 0xfb, 0xd1, 0x5d, 0x25, 0xcb,
	0xc3, 0x1f, 0x12, 0x28, 0x1b, 0x95, 0x48, 0x83, 0xdd, 0x0b, 0x73, 0x6c, 0x5a, 0x53, 0x93, 0x9e,
	0x63, 0xc7, 0x39, 0x39, 0xc5, 0xd4, 0xbd, 0xb2, 0xb1, 0xba, 0x85, 0xb6, 0xa1, 0x63, 0x5b, 0x8e,
	0xe1, 0x1a, 0x96, 0xa9, 0xd6, 0x50, 0x07, 0x1a, 0xfa, 0xd9, 0x89, 0xab, 0xd6, 0xd1, 0x01, 0xec,
	0xeb, 0x13, 0x03, 0x9b, 0x2e, 0x1d, 0x1a, 0x8e, 0x6e, 0x99, 0x26, 0xd6, 0x5d, 0x3c, 0xa4, 0x23,
	0x62, 0x9d, 0x53, 0x07, 0x93, 0x2f, 0x98, 0xa8, 0x0d, 0xa4, 0x40, 0xdb, 0x26, 0xd6, 0xc8, 0x98,
	0x60, 0xb5, 0x89, 0xf6, 0xe0, 0x9d, 0x3e, 0xb1, 0xf4, 0x31, 0x75, 0xc6, 0x78, 0x4a, 0x87, 0xd8,
	0x15, 0xbf, 0x50, 0x5b, 0x68, 0x07, 0x94, 0xd1, 0xc4, 0x9a, 0x52, 0xc7, 0x3d, 0x71, 0x2f, 0x1c,
	0xb5, 0x9d, 0xbb, 0xd8, 0x86, 0x79, 0xaa, 0x76, 0x0e, 0x3f, 0x03, 0xac, 0x63, 0x40, 0x08, 0xba,
	0x2f, 0x53, 0x16, 0xb5, 0x5b, 0x79, 0xad, 0x65, 0x63, 0x53, 0x95, 0x90, 0x0c, 0x4d, 0x7d, 0x62,
	0x39, 0x58, 0xad, 0xdd, 0xb4, 0xc4, 0x77, 0xfc, 0xd3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x04,
	0x44, 0xdb, 0x77, 0xda, 0x05, 0x00, 0x00,
}
