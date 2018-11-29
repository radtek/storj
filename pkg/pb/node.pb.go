// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: node.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// NodeType is an enum of possible node types
type NodeType int32

const (
	NodeType_ADMIN   NodeType = 0
	NodeType_STORAGE NodeType = 1
)

var NodeType_name = map[int32]string{
	0: "ADMIN",
	1: "STORAGE",
}
var NodeType_value = map[string]int32{
	"ADMIN":   0,
	"STORAGE": 1,
}

func (x NodeType) String() string {
	return proto.EnumName(NodeType_name, int32(x))
}
func (NodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_node_2ea4a85792199846, []int{0}
}

// NodeTransport is an enum of possible transports for the overlay network
type NodeTransport int32

const (
	NodeTransport_TCP_TLS_GRPC NodeTransport = 0
)

var NodeTransport_name = map[int32]string{
	0: "TCP_TLS_GRPC",
}
var NodeTransport_value = map[string]int32{
	"TCP_TLS_GRPC": 0,
}

func (x NodeTransport) String() string {
	return proto.EnumName(NodeTransport_name, int32(x))
}
func (NodeTransport) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_node_2ea4a85792199846, []int{1}
}

//  NodeRestrictions contains all relevant data about a nodes ability to store data
type NodeRestrictions struct {
	FreeBandwidth        int64    `protobuf:"varint,1,opt,name=freeBandwidth,proto3" json:"freeBandwidth,omitempty"`
	FreeDisk             int64    `protobuf:"varint,2,opt,name=freeDisk,proto3" json:"freeDisk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeRestrictions) Reset()         { *m = NodeRestrictions{} }
func (m *NodeRestrictions) String() string { return proto.CompactTextString(m) }
func (*NodeRestrictions) ProtoMessage()    {}
func (*NodeRestrictions) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_2ea4a85792199846, []int{0}
}
func (m *NodeRestrictions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeRestrictions.Unmarshal(m, b)
}
func (m *NodeRestrictions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeRestrictions.Marshal(b, m, deterministic)
}
func (dst *NodeRestrictions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeRestrictions.Merge(dst, src)
}
func (m *NodeRestrictions) XXX_Size() int {
	return xxx_messageInfo_NodeRestrictions.Size(m)
}
func (m *NodeRestrictions) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeRestrictions.DiscardUnknown(m)
}

var xxx_messageInfo_NodeRestrictions proto.InternalMessageInfo

func (m *NodeRestrictions) GetFreeBandwidth() int64 {
	if m != nil {
		return m.FreeBandwidth
	}
	return 0
}

func (m *NodeRestrictions) GetFreeDisk() int64 {
	if m != nil {
		return m.FreeDisk
	}
	return 0
}

// Node represents a node in the overlay network
// Node is info for a updating a single storagenode, used in the Update rpc calls
type Node struct {
	Id                   NodeID            `protobuf:"bytes,1,opt,name=id,proto3,customtype=NodeID" json:"id"`
	Address              *NodeAddress      `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Type                 NodeType          `protobuf:"varint,3,opt,name=type,proto3,enum=node.NodeType" json:"type,omitempty"`
	Restrictions         *NodeRestrictions `protobuf:"bytes,4,opt,name=restrictions" json:"restrictions,omitempty"`
	Metadata             *NodeMetadata     `protobuf:"bytes,5,opt,name=metadata" json:"metadata,omitempty"`
	LatencyList          []int64           `protobuf:"varint,6,rep,packed,name=latency_list,json=latencyList" json:"latency_list,omitempty"`
	AuditSuccess         bool              `protobuf:"varint,7,opt,name=audit_success,json=auditSuccess,proto3" json:"audit_success,omitempty"`
	IsUp                 bool              `protobuf:"varint,8,opt,name=is_up,json=isUp,proto3" json:"is_up,omitempty"`
	UpdateLatency        bool              `protobuf:"varint,9,opt,name=update_latency,json=updateLatency,proto3" json:"update_latency,omitempty"`
	UpdateAuditSuccess   bool              `protobuf:"varint,10,opt,name=update_audit_success,json=updateAuditSuccess,proto3" json:"update_audit_success,omitempty"`
	UpdateUptime         bool              `protobuf:"varint,11,opt,name=update_uptime,json=updateUptime,proto3" json:"update_uptime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_2ea4a85792199846, []int{1}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (dst *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(dst, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetAddress() *NodeAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Node) GetType() NodeType {
	if m != nil {
		return m.Type
	}
	return NodeType_ADMIN
}

func (m *Node) GetRestrictions() *NodeRestrictions {
	if m != nil {
		return m.Restrictions
	}
	return nil
}

func (m *Node) GetMetadata() *NodeMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Node) GetLatencyList() []int64 {
	if m != nil {
		return m.LatencyList
	}
	return nil
}

func (m *Node) GetAuditSuccess() bool {
	if m != nil {
		return m.AuditSuccess
	}
	return false
}

func (m *Node) GetIsUp() bool {
	if m != nil {
		return m.IsUp
	}
	return false
}

func (m *Node) GetUpdateLatency() bool {
	if m != nil {
		return m.UpdateLatency
	}
	return false
}

func (m *Node) GetUpdateAuditSuccess() bool {
	if m != nil {
		return m.UpdateAuditSuccess
	}
	return false
}

func (m *Node) GetUpdateUptime() bool {
	if m != nil {
		return m.UpdateUptime
	}
	return false
}

// NodeAddress contains the information needed to communicate with a node on the network
type NodeAddress struct {
	Transport            NodeTransport `protobuf:"varint,1,opt,name=transport,proto3,enum=node.NodeTransport" json:"transport,omitempty"`
	Address              string        `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *NodeAddress) Reset()         { *m = NodeAddress{} }
func (m *NodeAddress) String() string { return proto.CompactTextString(m) }
func (*NodeAddress) ProtoMessage()    {}
func (*NodeAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_2ea4a85792199846, []int{2}
}
func (m *NodeAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeAddress.Unmarshal(m, b)
}
func (m *NodeAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeAddress.Marshal(b, m, deterministic)
}
func (dst *NodeAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeAddress.Merge(dst, src)
}
func (m *NodeAddress) XXX_Size() int {
	return xxx_messageInfo_NodeAddress.Size(m)
}
func (m *NodeAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeAddress.DiscardUnknown(m)
}

var xxx_messageInfo_NodeAddress proto.InternalMessageInfo

func (m *NodeAddress) GetTransport() NodeTransport {
	if m != nil {
		return m.Transport
	}
	return NodeTransport_TCP_TLS_GRPC
}

func (m *NodeAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// NodeStats is info about a single storagenode stored in the stats db
type NodeStats struct {
	NodeId               NodeID   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3,customtype=NodeID" json:"node_id"`
	Latency_90           int64    `protobuf:"varint,2,opt,name=latency_90,json=latency90,proto3" json:"latency_90,omitempty"`
	AuditSuccessRatio    float64  `protobuf:"fixed64,3,opt,name=audit_success_ratio,json=auditSuccessRatio,proto3" json:"audit_success_ratio,omitempty"`
	UptimeRatio          float64  `protobuf:"fixed64,4,opt,name=uptime_ratio,json=uptimeRatio,proto3" json:"uptime_ratio,omitempty"`
	AuditCount           int64    `protobuf:"varint,5,opt,name=audit_count,json=auditCount,proto3" json:"audit_count,omitempty"`
	AuditSuccessCount    int64    `protobuf:"varint,6,opt,name=audit_success_count,json=auditSuccessCount,proto3" json:"audit_success_count,omitempty"`
	UptimeCount          int64    `protobuf:"varint,7,opt,name=uptime_count,json=uptimeCount,proto3" json:"uptime_count,omitempty"`
	UptimeSuccessCount   int64    `protobuf:"varint,8,opt,name=uptime_success_count,json=uptimeSuccessCount,proto3" json:"uptime_success_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeStats) Reset()         { *m = NodeStats{} }
func (m *NodeStats) String() string { return proto.CompactTextString(m) }
func (*NodeStats) ProtoMessage()    {}
func (*NodeStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_2ea4a85792199846, []int{3}
}
func (m *NodeStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeStats.Unmarshal(m, b)
}
func (m *NodeStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeStats.Marshal(b, m, deterministic)
}
func (dst *NodeStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeStats.Merge(dst, src)
}
func (m *NodeStats) XXX_Size() int {
	return xxx_messageInfo_NodeStats.Size(m)
}
func (m *NodeStats) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeStats.DiscardUnknown(m)
}

var xxx_messageInfo_NodeStats proto.InternalMessageInfo

func (m *NodeStats) GetLatency_90() int64 {
	if m != nil {
		return m.Latency_90
	}
	return 0
}

func (m *NodeStats) GetAuditSuccessRatio() float64 {
	if m != nil {
		return m.AuditSuccessRatio
	}
	return 0
}

func (m *NodeStats) GetUptimeRatio() float64 {
	if m != nil {
		return m.UptimeRatio
	}
	return 0
}

func (m *NodeStats) GetAuditCount() int64 {
	if m != nil {
		return m.AuditCount
	}
	return 0
}

func (m *NodeStats) GetAuditSuccessCount() int64 {
	if m != nil {
		return m.AuditSuccessCount
	}
	return 0
}

func (m *NodeStats) GetUptimeCount() int64 {
	if m != nil {
		return m.UptimeCount
	}
	return 0
}

func (m *NodeStats) GetUptimeSuccessCount() int64 {
	if m != nil {
		return m.UptimeSuccessCount
	}
	return 0
}

// TODO: combine with `NodeStats`
// NodeRep is the reputation characteristics of a node
type NodeRep struct {
	MinUptime            float32  `protobuf:"fixed32,1,opt,name=min_uptime,json=minUptime,proto3" json:"min_uptime,omitempty"`
	MinAuditSuccess      float32  `protobuf:"fixed32,2,opt,name=min_audit_success,json=minAuditSuccess,proto3" json:"min_audit_success,omitempty"`
	MinAuditCount        int64    `protobuf:"varint,3,opt,name=min_audit_count,json=minAuditCount,proto3" json:"min_audit_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeRep) Reset()         { *m = NodeRep{} }
func (m *NodeRep) String() string { return proto.CompactTextString(m) }
func (*NodeRep) ProtoMessage()    {}
func (*NodeRep) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_2ea4a85792199846, []int{4}
}
func (m *NodeRep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeRep.Unmarshal(m, b)
}
func (m *NodeRep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeRep.Marshal(b, m, deterministic)
}
func (dst *NodeRep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeRep.Merge(dst, src)
}
func (m *NodeRep) XXX_Size() int {
	return xxx_messageInfo_NodeRep.Size(m)
}
func (m *NodeRep) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeRep.DiscardUnknown(m)
}

var xxx_messageInfo_NodeRep proto.InternalMessageInfo

func (m *NodeRep) GetMinUptime() float32 {
	if m != nil {
		return m.MinUptime
	}
	return 0
}

func (m *NodeRep) GetMinAuditSuccess() float32 {
	if m != nil {
		return m.MinAuditSuccess
	}
	return 0
}

func (m *NodeRep) GetMinAuditCount() int64 {
	if m != nil {
		return m.MinAuditCount
	}
	return 0
}

type NodeMetadata struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Wallet               string   `protobuf:"bytes,2,opt,name=wallet,proto3" json:"wallet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeMetadata) Reset()         { *m = NodeMetadata{} }
func (m *NodeMetadata) String() string { return proto.CompactTextString(m) }
func (*NodeMetadata) ProtoMessage()    {}
func (*NodeMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_2ea4a85792199846, []int{5}
}
func (m *NodeMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeMetadata.Unmarshal(m, b)
}
func (m *NodeMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeMetadata.Marshal(b, m, deterministic)
}
func (dst *NodeMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeMetadata.Merge(dst, src)
}
func (m *NodeMetadata) XXX_Size() int {
	return xxx_messageInfo_NodeMetadata.Size(m)
}
func (m *NodeMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_NodeMetadata proto.InternalMessageInfo

func (m *NodeMetadata) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *NodeMetadata) GetWallet() string {
	if m != nil {
		return m.Wallet
	}
	return ""
}

func init() {
	proto.RegisterType((*NodeRestrictions)(nil), "node.NodeRestrictions")
	proto.RegisterType((*Node)(nil), "node.Node")
	proto.RegisterType((*NodeAddress)(nil), "node.NodeAddress")
	proto.RegisterType((*NodeStats)(nil), "node.NodeStats")
	proto.RegisterType((*NodeRep)(nil), "node.NodeRep")
	proto.RegisterType((*NodeMetadata)(nil), "node.NodeMetadata")
	proto.RegisterEnum("node.NodeType", NodeType_name, NodeType_value)
	proto.RegisterEnum("node.NodeTransport", NodeTransport_name, NodeTransport_value)
}

func init() { proto.RegisterFile("node.proto", fileDescriptor_node_2ea4a85792199846) }

var fileDescriptor_node_2ea4a85792199846 = []byte{
	// 646 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0x4f, 0x4f, 0xdb, 0x4c,
	0x10, 0xc6, 0x89, 0xed, 0xfc, 0xf1, 0x38, 0x09, 0x61, 0x40, 0xc8, 0x42, 0x7a, 0x5f, 0x82, 0x79,
	0xdf, 0x36, 0xa2, 0x52, 0x44, 0xe9, 0x89, 0xaa, 0x97, 0x00, 0x15, 0x42, 0x02, 0x8a, 0x36, 0xe1,
	0xc2, 0xc5, 0x5a, 0xe2, 0x2d, 0x5d, 0x35, 0xb1, 0x2d, 0x7b, 0x23, 0x84, 0xd4, 0xcf, 0xd6, 0x73,
	0x0f, 0xfd, 0x04, 0x3d, 0xf0, 0x59, 0xaa, 0x9d, 0x75, 0x88, 0xad, 0xaa, 0x37, 0xef, 0xf3, 0xfc,
	0x76, 0x66, 0x77, 0x66, 0xbc, 0x00, 0x71, 0x12, 0x89, 0x61, 0x9a, 0x25, 0x2a, 0x41, 0x47, 0x7f,
	0xef, 0xc0, 0x43, 0xf2, 0x90, 0x18, 0x25, 0x98, 0x40, 0xef, 0x3a, 0x89, 0x04, 0x13, 0xb9, 0xca,
	0xe4, 0x54, 0xc9, 0x24, 0xce, 0xf1, 0x3f, 0xe8, 0x7c, 0xce, 0x84, 0x38, 0xe1, 0x71, 0xf4, 0x28,
	0x23, 0xf5, 0xc5, 0xaf, 0xf5, 0x6b, 0x03, 0x9b, 0x55, 0x45, 0xdc, 0x81, 0x96, 0x16, 0xce, 0x64,
	0xfe, 0xd5, 0xb7, 0x08, 0x78, 0x59, 0x07, 0xdf, 0x6d, 0x70, 0x74, 0x58, 0xfc, 0x17, 0x2c, 0x19,
	0xd1, 0xfe, 0xf6, 0x49, 0xf7, 0xc7, 0xf3, 0xee, 0xda, 0xaf, 0xe7, 0xdd, 0x86, 0x76, 0x2e, 0xce,
	0x98, 0x25, 0x23, 0x7c, 0x03, 0x4d, 0x1e, 0x45, 0x99, 0xc8, 0x73, 0x8a, 0xe1, 0x1d, 0x6d, 0x0c,
	0xe9, 0xb8, 0x1a, 0x19, 0x19, 0x83, 0x2d, 0x09, 0x0c, 0xc0, 0x51, 0x4f, 0xa9, 0xf0, 0xed, 0x7e,
	0x6d, 0xd0, 0x3d, 0xea, 0xae, 0xc8, 0xc9, 0x53, 0x2a, 0x18, 0x79, 0xf8, 0x1e, 0xda, 0x59, 0xe9,
	0x2e, 0xbe, 0x43, 0x51, 0xb7, 0x57, 0x6c, 0xf9, 0xa6, 0xac, 0xc2, 0xe2, 0x10, 0x5a, 0x73, 0xa1,
	0x78, 0xc4, 0x15, 0xf7, 0xeb, 0xb4, 0x0f, 0x57, 0xfb, 0xae, 0x0a, 0x87, 0xbd, 0x30, 0xb8, 0x07,
	0xed, 0x19, 0x57, 0x22, 0x9e, 0x3e, 0x85, 0x33, 0x99, 0x2b, 0xbf, 0xd1, 0xb7, 0x07, 0x36, 0xf3,
	0x0a, 0xed, 0x52, 0xe6, 0x0a, 0xf7, 0xa1, 0xc3, 0x17, 0x91, 0x54, 0x61, 0xbe, 0x98, 0x4e, 0xf5,
	0x2d, 0x9b, 0xfd, 0xda, 0xa0, 0xc5, 0xda, 0x24, 0x8e, 0x8d, 0x86, 0x9b, 0x50, 0x97, 0x79, 0xb8,
	0x48, 0xfd, 0x16, 0x99, 0x8e, 0xcc, 0x6f, 0x53, 0xfc, 0x1f, 0xba, 0x8b, 0x34, 0xe2, 0x4a, 0x84,
	0x45, 0x3c, 0xdf, 0x25, 0xb7, 0x63, 0xd4, 0x4b, 0x23, 0xe2, 0x21, 0x6c, 0x15, 0x58, 0x35, 0x0f,
	0x10, 0x8c, 0xc6, 0x1b, 0x95, 0xb3, 0xed, 0x43, 0x11, 0x22, 0x5c, 0xa4, 0x4a, 0xce, 0x85, 0xef,
	0x99, 0x23, 0x19, 0xf1, 0x96, 0xb4, 0xe0, 0x0e, 0xbc, 0x52, 0x0b, 0xf0, 0x2d, 0xb8, 0x2a, 0xe3,
	0x71, 0x9e, 0x26, 0x99, 0xa2, 0x6e, 0x76, 0x8f, 0x36, 0x4b, 0xe5, 0x5f, 0x5a, 0x6c, 0x45, 0xa1,
	0x5f, 0xed, 0xac, 0xfb, 0xd2, 0xc6, 0xe0, 0xa7, 0x05, 0xae, 0xde, 0x36, 0x56, 0x5c, 0xe5, 0xf8,
	0x1a, 0x9a, 0x3a, 0x50, 0xf8, 0xd7, 0x31, 0x69, 0x68, 0xfb, 0x22, 0xc2, 0x7f, 0x00, 0x96, 0xd5,
	0x3e, 0x3e, 0x2c, 0x26, 0xce, 0x2d, 0x94, 0xe3, 0x43, 0x1c, 0xc2, 0x66, 0xa5, 0x02, 0x61, 0xc6,
	0x95, 0x4c, 0x68, 0x56, 0x6a, 0x6c, 0xa3, 0x5c, 0x6f, 0xa6, 0x0d, 0xdd, 0x3c, 0x73, 0xff, 0x02,
	0x74, 0x08, 0xf4, 0x8c, 0x66, 0x90, 0x5d, 0xf0, 0x4c, 0xc8, 0x69, 0xb2, 0x88, 0x15, 0x8d, 0x84,
	0xcd, 0x80, 0xa4, 0x53, 0xad, 0xfc, 0x99, 0xd3, 0x80, 0x0d, 0x02, 0x2b, 0x39, 0x0d, 0xbf, 0xca,
	0x69, 0xc0, 0x26, 0x81, 0x45, 0x4e, 0x83, 0x50, 0x3f, 0x09, 0xa9, 0xc6, 0x6c, 0x11, 0x8a, 0xc6,
	0x2b, 0x07, 0x0d, 0xbe, 0x41, 0xd3, 0xcc, 0x75, 0xaa, 0x4b, 0x34, 0x97, 0xf1, 0xb2, 0xaf, 0xba,
	0x9c, 0x16, 0x73, 0xe7, 0x32, 0x36, 0x4d, 0xc5, 0x03, 0xd8, 0xd0, 0x76, 0x75, 0x50, 0x2c, 0xa2,
	0xd6, 0xe7, 0x32, 0xae, 0x4c, 0xc9, 0x2b, 0x58, 0x5f, 0xb1, 0xe6, 0x08, 0xb6, 0x79, 0x05, 0x96,
	0xa4, 0xc9, 0xfe, 0x01, 0xda, 0xe5, 0xbf, 0x03, 0xb7, 0xa0, 0x2e, 0xe6, 0x5c, 0xce, 0x28, 0xbb,
	0xcb, 0xcc, 0x02, 0xb7, 0xa1, 0xf1, 0xc8, 0x67, 0x33, 0xa1, 0x8a, 0x59, 0x28, 0x56, 0x07, 0x01,
	0xb4, 0x96, 0xff, 0x2f, 0xba, 0x50, 0x1f, 0x9d, 0x5d, 0x5d, 0x5c, 0xf7, 0xd6, 0xd0, 0x83, 0xe6,
	0x78, 0xf2, 0x89, 0x8d, 0xce, 0x3f, 0xf6, 0x6a, 0x07, 0x7b, 0xd0, 0xa9, 0x0c, 0x19, 0xf6, 0xa0,
	0x3d, 0x39, 0xbd, 0x09, 0x27, 0x97, 0xe3, 0xf0, 0x9c, 0xdd, 0x9c, 0xf6, 0xd6, 0x4e, 0x9c, 0x3b,
	0x2b, 0xbd, 0xbf, 0x6f, 0xd0, 0x8b, 0xf6, 0xee, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x4d,
	0xc8, 0xfa, 0xf1, 0x04, 0x00, 0x00,
}
