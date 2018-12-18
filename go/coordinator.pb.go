// Code generated by protoc-gen-go. DO NOT EDIT.
// source: coordinator.proto

package coordinator

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

type CoordinatorMessageType int32

const (
	CoordinatorMessageType_UNKNOWN_MESSAGE_TYPE CoordinatorMessageType = 0
	CoordinatorMessageType_WELCOME_SERVER       CoordinatorMessageType = 1
	CoordinatorMessageType_WELCOME_CLIENT       CoordinatorMessageType = 2
	CoordinatorMessageType_CONNECT              CoordinatorMessageType = 3
	CoordinatorMessageType_WEBRTC_SUPPORTED     CoordinatorMessageType = 4
	CoordinatorMessageType_WEBRTC_OFFER         CoordinatorMessageType = 5
	CoordinatorMessageType_WEBRTC_ANSWER        CoordinatorMessageType = 6
	CoordinatorMessageType_WEBRTC_ICE_CANDIDATE CoordinatorMessageType = 7
)

var CoordinatorMessageType_name = map[int32]string{
	0: "UNKNOWN_MESSAGE_TYPE",
	1: "WELCOME_SERVER",
	2: "WELCOME_CLIENT",
	3: "CONNECT",
	4: "WEBRTC_SUPPORTED",
	5: "WEBRTC_OFFER",
	6: "WEBRTC_ANSWER",
	7: "WEBRTC_ICE_CANDIDATE",
}

var CoordinatorMessageType_value = map[string]int32{
	"UNKNOWN_MESSAGE_TYPE": 0,
	"WELCOME_SERVER":       1,
	"WELCOME_CLIENT":       2,
	"CONNECT":              3,
	"WEBRTC_SUPPORTED":     4,
	"WEBRTC_OFFER":         5,
	"WEBRTC_ANSWER":        6,
	"WEBRTC_ICE_CANDIDATE": 7,
}

func (x CoordinatorMessageType) String() string {
	return proto.EnumName(CoordinatorMessageType_name, int32(x))
}

func (CoordinatorMessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_99e779eb11ceee19, []int{0}
}

type CoordinatorMessage struct {
	Type                 CoordinatorMessageType `protobuf:"varint,1,opt,name=type,proto3,enum=CoordinatorMessageType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *CoordinatorMessage) Reset()         { *m = CoordinatorMessage{} }
func (m *CoordinatorMessage) String() string { return proto.CompactTextString(m) }
func (*CoordinatorMessage) ProtoMessage()    {}
func (*CoordinatorMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_99e779eb11ceee19, []int{0}
}

func (m *CoordinatorMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CoordinatorMessage.Unmarshal(m, b)
}
func (m *CoordinatorMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CoordinatorMessage.Marshal(b, m, deterministic)
}
func (m *CoordinatorMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoordinatorMessage.Merge(m, src)
}
func (m *CoordinatorMessage) XXX_Size() int {
	return xxx_messageInfo_CoordinatorMessage.Size(m)
}
func (m *CoordinatorMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CoordinatorMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CoordinatorMessage proto.InternalMessageInfo

func (m *CoordinatorMessage) GetType() CoordinatorMessageType {
	if m != nil {
		return m.Type
	}
	return CoordinatorMessageType_UNKNOWN_MESSAGE_TYPE
}

type WelcomeServerMessage struct {
	Type                 CoordinatorMessageType `protobuf:"varint,1,opt,name=type,proto3,enum=CoordinatorMessageType" json:"type,omitempty"`
	Alias                string                 `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	Peers                []string               `protobuf:"bytes,3,rep,name=peers,proto3" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *WelcomeServerMessage) Reset()         { *m = WelcomeServerMessage{} }
func (m *WelcomeServerMessage) String() string { return proto.CompactTextString(m) }
func (*WelcomeServerMessage) ProtoMessage()    {}
func (*WelcomeServerMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_99e779eb11ceee19, []int{1}
}

func (m *WelcomeServerMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WelcomeServerMessage.Unmarshal(m, b)
}
func (m *WelcomeServerMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WelcomeServerMessage.Marshal(b, m, deterministic)
}
func (m *WelcomeServerMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WelcomeServerMessage.Merge(m, src)
}
func (m *WelcomeServerMessage) XXX_Size() int {
	return xxx_messageInfo_WelcomeServerMessage.Size(m)
}
func (m *WelcomeServerMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_WelcomeServerMessage.DiscardUnknown(m)
}

var xxx_messageInfo_WelcomeServerMessage proto.InternalMessageInfo

func (m *WelcomeServerMessage) GetType() CoordinatorMessageType {
	if m != nil {
		return m.Type
	}
	return CoordinatorMessageType_UNKNOWN_MESSAGE_TYPE
}

func (m *WelcomeServerMessage) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *WelcomeServerMessage) GetPeers() []string {
	if m != nil {
		return m.Peers
	}
	return nil
}

type WelcomeClientMessage struct {
	Type                 CoordinatorMessageType `protobuf:"varint,1,opt,name=type,proto3,enum=CoordinatorMessageType" json:"type,omitempty"`
	Alias                string                 `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *WelcomeClientMessage) Reset()         { *m = WelcomeClientMessage{} }
func (m *WelcomeClientMessage) String() string { return proto.CompactTextString(m) }
func (*WelcomeClientMessage) ProtoMessage()    {}
func (*WelcomeClientMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_99e779eb11ceee19, []int{2}
}

func (m *WelcomeClientMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WelcomeClientMessage.Unmarshal(m, b)
}
func (m *WelcomeClientMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WelcomeClientMessage.Marshal(b, m, deterministic)
}
func (m *WelcomeClientMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WelcomeClientMessage.Merge(m, src)
}
func (m *WelcomeClientMessage) XXX_Size() int {
	return xxx_messageInfo_WelcomeClientMessage.Size(m)
}
func (m *WelcomeClientMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_WelcomeClientMessage.DiscardUnknown(m)
}

var xxx_messageInfo_WelcomeClientMessage proto.InternalMessageInfo

func (m *WelcomeClientMessage) GetType() CoordinatorMessageType {
	if m != nil {
		return m.Type
	}
	return CoordinatorMessageType_UNKNOWN_MESSAGE_TYPE
}

func (m *WelcomeClientMessage) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

type ConnectMessage struct {
	Type                 CoordinatorMessageType `protobuf:"varint,1,opt,name=type,proto3,enum=CoordinatorMessageType" json:"type,omitempty"`
	ToAlias              string                 `protobuf:"bytes,2,opt,name=toAlias,proto3" json:"toAlias,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ConnectMessage) Reset()         { *m = ConnectMessage{} }
func (m *ConnectMessage) String() string { return proto.CompactTextString(m) }
func (*ConnectMessage) ProtoMessage()    {}
func (*ConnectMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_99e779eb11ceee19, []int{3}
}

func (m *ConnectMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectMessage.Unmarshal(m, b)
}
func (m *ConnectMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectMessage.Marshal(b, m, deterministic)
}
func (m *ConnectMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectMessage.Merge(m, src)
}
func (m *ConnectMessage) XXX_Size() int {
	return xxx_messageInfo_ConnectMessage.Size(m)
}
func (m *ConnectMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectMessage proto.InternalMessageInfo

func (m *ConnectMessage) GetType() CoordinatorMessageType {
	if m != nil {
		return m.Type
	}
	return CoordinatorMessageType_UNKNOWN_MESSAGE_TYPE
}

func (m *ConnectMessage) GetToAlias() string {
	if m != nil {
		return m.ToAlias
	}
	return ""
}

type WebRtcMessage struct {
	Type                 CoordinatorMessageType `protobuf:"varint,1,opt,name=type,proto3,enum=CoordinatorMessageType" json:"type,omitempty"`
	Alias                string                 `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	ToAlias              string                 `protobuf:"bytes,3,opt,name=toAlias,proto3" json:"toAlias,omitempty"`
	Sdp                  string                 `protobuf:"bytes,4,opt,name=sdp,proto3" json:"sdp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *WebRtcMessage) Reset()         { *m = WebRtcMessage{} }
func (m *WebRtcMessage) String() string { return proto.CompactTextString(m) }
func (*WebRtcMessage) ProtoMessage()    {}
func (*WebRtcMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_99e779eb11ceee19, []int{4}
}

func (m *WebRtcMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebRtcMessage.Unmarshal(m, b)
}
func (m *WebRtcMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebRtcMessage.Marshal(b, m, deterministic)
}
func (m *WebRtcMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebRtcMessage.Merge(m, src)
}
func (m *WebRtcMessage) XXX_Size() int {
	return xxx_messageInfo_WebRtcMessage.Size(m)
}
func (m *WebRtcMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_WebRtcMessage.DiscardUnknown(m)
}

var xxx_messageInfo_WebRtcMessage proto.InternalMessageInfo

func (m *WebRtcMessage) GetType() CoordinatorMessageType {
	if m != nil {
		return m.Type
	}
	return CoordinatorMessageType_UNKNOWN_MESSAGE_TYPE
}

func (m *WebRtcMessage) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *WebRtcMessage) GetToAlias() string {
	if m != nil {
		return m.ToAlias
	}
	return ""
}

func (m *WebRtcMessage) GetSdp() string {
	if m != nil {
		return m.Sdp
	}
	return ""
}

func init() {
	proto.RegisterEnum("CoordinatorMessageType", CoordinatorMessageType_name, CoordinatorMessageType_value)
	proto.RegisterType((*CoordinatorMessage)(nil), "CoordinatorMessage")
	proto.RegisterType((*WelcomeServerMessage)(nil), "WelcomeServerMessage")
	proto.RegisterType((*WelcomeClientMessage)(nil), "WelcomeClientMessage")
	proto.RegisterType((*ConnectMessage)(nil), "ConnectMessage")
	proto.RegisterType((*WebRtcMessage)(nil), "WebRtcMessage")
}

func init() { proto.RegisterFile("coordinator.proto", fileDescriptor_99e779eb11ceee19) }

var fileDescriptor_99e779eb11ceee19 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x5f, 0x6b, 0xea, 0x30,
	0x18, 0xc6, 0x4f, 0xad, 0x7f, 0x38, 0xef, 0x39, 0x4a, 0x0c, 0xe5, 0x9c, 0x5e, 0x4a, 0xaf, 0x64,
	0x03, 0x2f, 0xb6, 0x4f, 0xd0, 0xc5, 0xd7, 0x21, 0xd3, 0x54, 0xd2, 0xb8, 0xe0, 0x55, 0xa9, 0x35,
	0x0c, 0xc1, 0x35, 0xa5, 0x2d, 0x03, 0x6f, 0xf6, 0xc9, 0xf6, 0xe1, 0x46, 0x3b, 0xb7, 0x39, 0xb6,
	0x2b, 0xf1, 0x2e, 0xcf, 0x8f, 0x87, 0xe7, 0x17, 0x42, 0xa0, 0x9f, 0x18, 0x93, 0x6f, 0xb6, 0x69,
	0x5c, 0x9a, 0x7c, 0x94, 0xe5, 0xa6, 0x34, 0x9e, 0x0f, 0x94, 0x7d, 0xc2, 0xb9, 0x2e, 0x8a, 0xf8,
	0x41, 0xd3, 0x4b, 0x68, 0x96, 0xfb, 0x4c, 0xbb, 0xd6, 0xc0, 0x1a, 0xf6, 0xae, 0xfe, 0x8f, 0xbe,
	0x57, 0xe4, 0x3e, 0xd3, 0xa2, 0x2e, 0x79, 0x06, 0x1c, 0xa5, 0x77, 0x89, 0x79, 0xd4, 0xa1, 0xce,
	0x9f, 0xf4, 0x49, 0x23, 0xd4, 0x81, 0x56, 0xbc, 0xdb, 0xc6, 0x85, 0xdb, 0x18, 0x58, 0xc3, 0xdf,
	0xe2, 0x2d, 0x54, 0x34, 0xd3, 0x3a, 0x2f, 0x5c, 0x7b, 0x60, 0x57, 0xb4, 0x0e, 0xde, 0xea, 0x43,
	0xc8, 0x76, 0x5b, 0x9d, 0x96, 0xe7, 0x13, 0x7a, 0x0a, 0x7a, 0xcc, 0xa4, 0xa9, 0x4e, 0x4e, 0x1b,
	0x75, 0xa1, 0x53, 0x1a, 0xff, 0x68, 0xf6, 0x3d, 0x7a, 0xcf, 0xd0, 0x55, 0x7a, 0x2d, 0xca, 0xe4,
	0x8c, 0xaf, 0x73, 0x64, 0xb3, 0xbf, 0xd8, 0x28, 0x01, 0xbb, 0xd8, 0x64, 0x6e, 0xb3, 0xa6, 0xd5,
	0xf1, 0xe2, 0xc5, 0x82, 0x7f, 0x3f, 0x2b, 0xa8, 0x0b, 0xce, 0x92, 0xdf, 0xf1, 0x40, 0xf1, 0x68,
	0x8e, 0x61, 0xe8, 0xdf, 0x62, 0x24, 0x57, 0x0b, 0x24, 0xbf, 0x28, 0x85, 0x9e, 0xc2, 0x19, 0x0b,
	0xe6, 0x18, 0x85, 0x28, 0xee, 0x51, 0x10, 0xeb, 0x98, 0xb1, 0xd9, 0x14, 0xb9, 0x24, 0x0d, 0xfa,
	0x07, 0x3a, 0x2c, 0xe0, 0x1c, 0x99, 0x24, 0x36, 0x75, 0x80, 0x28, 0xbc, 0x11, 0x92, 0x45, 0xe1,
	0x72, 0xb1, 0x08, 0x84, 0xc4, 0x31, 0x69, 0x52, 0x02, 0x7f, 0x0f, 0x34, 0x98, 0x4c, 0x50, 0x90,
	0x16, 0xed, 0x43, 0xf7, 0x40, 0x7c, 0x1e, 0x2a, 0x14, 0xa4, 0x5d, 0xdd, 0xe4, 0x80, 0xa6, 0x0c,
	0x23, 0xe6, 0xf3, 0xf1, 0x74, 0xec, 0x4b, 0x24, 0x9d, 0x75, 0xbb, 0xfe, 0xad, 0xd7, 0xaf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xf3, 0x8f, 0x9f, 0x8c, 0xc2, 0x02, 0x00, 0x00,
}
