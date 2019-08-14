/*
Notice: This file has been modified for Hyperledger Fabric SDK Go usage.
Please review third_party pinning scripts and patches for more details.
*/
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/collection.proto

package common

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// CollectionConfigPackage represents an array of CollectionConfig
// messages; the extra struct is required because repeated oneof is
// forbidden by the protobuf syntax
type CollectionConfigPackage struct {
	Config               []*CollectionConfig `protobuf:"bytes,1,rep,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CollectionConfigPackage) Reset()         { *m = CollectionConfigPackage{} }
func (m *CollectionConfigPackage) String() string { return proto.CompactTextString(m) }
func (*CollectionConfigPackage) ProtoMessage()    {}
func (*CollectionConfigPackage) Descriptor() ([]byte, []int) {
	return fileDescriptor_89f245fc544906c7, []int{0}
}

func (m *CollectionConfigPackage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionConfigPackage.Unmarshal(m, b)
}
func (m *CollectionConfigPackage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionConfigPackage.Marshal(b, m, deterministic)
}
func (m *CollectionConfigPackage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionConfigPackage.Merge(m, src)
}
func (m *CollectionConfigPackage) XXX_Size() int {
	return xxx_messageInfo_CollectionConfigPackage.Size(m)
}
func (m *CollectionConfigPackage) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionConfigPackage.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionConfigPackage proto.InternalMessageInfo

func (m *CollectionConfigPackage) GetConfig() []*CollectionConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

// CollectionConfig defines the configuration of a collection object;
// it currently contains a single, static type.
// Dynamic collections are deferred.
type CollectionConfig struct {
	// Types that are valid to be assigned to Payload:
	//	*CollectionConfig_StaticCollectionConfig
	Payload              isCollectionConfig_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *CollectionConfig) Reset()         { *m = CollectionConfig{} }
func (m *CollectionConfig) String() string { return proto.CompactTextString(m) }
func (*CollectionConfig) ProtoMessage()    {}
func (*CollectionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_89f245fc544906c7, []int{1}
}

func (m *CollectionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionConfig.Unmarshal(m, b)
}
func (m *CollectionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionConfig.Marshal(b, m, deterministic)
}
func (m *CollectionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionConfig.Merge(m, src)
}
func (m *CollectionConfig) XXX_Size() int {
	return xxx_messageInfo_CollectionConfig.Size(m)
}
func (m *CollectionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionConfig proto.InternalMessageInfo

type isCollectionConfig_Payload interface {
	isCollectionConfig_Payload()
}

type CollectionConfig_StaticCollectionConfig struct {
	StaticCollectionConfig *StaticCollectionConfig `protobuf:"bytes,1,opt,name=static_collection_config,json=staticCollectionConfig,proto3,oneof"`
}

func (*CollectionConfig_StaticCollectionConfig) isCollectionConfig_Payload() {}

func (m *CollectionConfig) GetPayload() isCollectionConfig_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *CollectionConfig) GetStaticCollectionConfig() *StaticCollectionConfig {
	if x, ok := m.GetPayload().(*CollectionConfig_StaticCollectionConfig); ok {
		return x.StaticCollectionConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CollectionConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CollectionConfig_StaticCollectionConfig)(nil),
	}
}

// StaticCollectionConfig constitutes the configuration parameters of a
// static collection object. Static collections are collections that are
// known at chaincode instantiation time, and that cannot be changed.
// Dynamic collections are deferred.
type StaticCollectionConfig struct {
	// the name of the collection inside the denoted chaincode
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// a reference to a policy residing / managed in the config block
	// to define which orgs have access to this collection’s private data
	MemberOrgsPolicy *CollectionPolicyConfig `protobuf:"bytes,2,opt,name=member_orgs_policy,json=memberOrgsPolicy,proto3" json:"member_orgs_policy,omitempty"`
	// The minimum number of peers private data will be sent to upon
	// endorsement. The endorsement would fail if dissemination to at least
	// this number of peers is not achieved.
	RequiredPeerCount int32 `protobuf:"varint,3,opt,name=required_peer_count,json=requiredPeerCount,proto3" json:"required_peer_count,omitempty"`
	// The maximum number of peers that private data will be sent to
	// upon endorsement. This number has to be bigger than required_peer_count.
	MaximumPeerCount int32 `protobuf:"varint,4,opt,name=maximum_peer_count,json=maximumPeerCount,proto3" json:"maximum_peer_count,omitempty"`
	// The number of blocks after which the collection data expires.
	// For instance if the value is set to 10, a key last modified by block number 100
	// will be purged at block number 111. A zero value is treated same as MaxUint64
	BlockToLive uint64 `protobuf:"varint,5,opt,name=block_to_live,json=blockToLive,proto3" json:"block_to_live,omitempty"`
	// The member only read access denotes whether only collection member clients
	// can read the private data (if set to true), or even non members can
	// read the data (if set to false, for example if you want to implement more granular
	// access logic in the chaincode)
	MemberOnlyRead bool `protobuf:"varint,6,opt,name=member_only_read,json=memberOnlyRead,proto3" json:"member_only_read,omitempty"`
	// The member only write access denotes whether only collection member clients
	// can write the private data (if set to true), or even non members can
	// write the data (if set to false, for example if you want to implement more granular
	// access logic in the chaincode)
	MemberOnlyWrite      bool     `protobuf:"varint,7,opt,name=member_only_write,json=memberOnlyWrite,proto3" json:"member_only_write,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StaticCollectionConfig) Reset()         { *m = StaticCollectionConfig{} }
func (m *StaticCollectionConfig) String() string { return proto.CompactTextString(m) }
func (*StaticCollectionConfig) ProtoMessage()    {}
func (*StaticCollectionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_89f245fc544906c7, []int{2}
}

func (m *StaticCollectionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StaticCollectionConfig.Unmarshal(m, b)
}
func (m *StaticCollectionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StaticCollectionConfig.Marshal(b, m, deterministic)
}
func (m *StaticCollectionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StaticCollectionConfig.Merge(m, src)
}
func (m *StaticCollectionConfig) XXX_Size() int {
	return xxx_messageInfo_StaticCollectionConfig.Size(m)
}
func (m *StaticCollectionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_StaticCollectionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_StaticCollectionConfig proto.InternalMessageInfo

func (m *StaticCollectionConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StaticCollectionConfig) GetMemberOrgsPolicy() *CollectionPolicyConfig {
	if m != nil {
		return m.MemberOrgsPolicy
	}
	return nil
}

func (m *StaticCollectionConfig) GetRequiredPeerCount() int32 {
	if m != nil {
		return m.RequiredPeerCount
	}
	return 0
}

func (m *StaticCollectionConfig) GetMaximumPeerCount() int32 {
	if m != nil {
		return m.MaximumPeerCount
	}
	return 0
}

func (m *StaticCollectionConfig) GetBlockToLive() uint64 {
	if m != nil {
		return m.BlockToLive
	}
	return 0
}

func (m *StaticCollectionConfig) GetMemberOnlyRead() bool {
	if m != nil {
		return m.MemberOnlyRead
	}
	return false
}

func (m *StaticCollectionConfig) GetMemberOnlyWrite() bool {
	if m != nil {
		return m.MemberOnlyWrite
	}
	return false
}

// Collection policy configuration. Initially, the configuration can only
// contain a SignaturePolicy. In the future, the SignaturePolicy may be a
// more general Policy. Instead of containing the actual policy, the
// configuration may in the future contain a string reference to a policy.
type CollectionPolicyConfig struct {
	// Types that are valid to be assigned to Payload:
	//	*CollectionPolicyConfig_SignaturePolicy
	Payload              isCollectionPolicyConfig_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *CollectionPolicyConfig) Reset()         { *m = CollectionPolicyConfig{} }
func (m *CollectionPolicyConfig) String() string { return proto.CompactTextString(m) }
func (*CollectionPolicyConfig) ProtoMessage()    {}
func (*CollectionPolicyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_89f245fc544906c7, []int{3}
}

func (m *CollectionPolicyConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionPolicyConfig.Unmarshal(m, b)
}
func (m *CollectionPolicyConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionPolicyConfig.Marshal(b, m, deterministic)
}
func (m *CollectionPolicyConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionPolicyConfig.Merge(m, src)
}
func (m *CollectionPolicyConfig) XXX_Size() int {
	return xxx_messageInfo_CollectionPolicyConfig.Size(m)
}
func (m *CollectionPolicyConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionPolicyConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionPolicyConfig proto.InternalMessageInfo

type isCollectionPolicyConfig_Payload interface {
	isCollectionPolicyConfig_Payload()
}

type CollectionPolicyConfig_SignaturePolicy struct {
	SignaturePolicy *SignaturePolicyEnvelope `protobuf:"bytes,1,opt,name=signature_policy,json=signaturePolicy,proto3,oneof"`
}

func (*CollectionPolicyConfig_SignaturePolicy) isCollectionPolicyConfig_Payload() {}

func (m *CollectionPolicyConfig) GetPayload() isCollectionPolicyConfig_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *CollectionPolicyConfig) GetSignaturePolicy() *SignaturePolicyEnvelope {
	if x, ok := m.GetPayload().(*CollectionPolicyConfig_SignaturePolicy); ok {
		return x.SignaturePolicy
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CollectionPolicyConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CollectionPolicyConfig_SignaturePolicy)(nil),
	}
}

// CollectionCriteria defines an element of a private data that corresponds
// to a certain transaction and collection
type CollectionCriteria struct {
	Channel              string   `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	TxId                 string   `protobuf:"bytes,2,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	Collection           string   `protobuf:"bytes,3,opt,name=collection,proto3" json:"collection,omitempty"`
	Namespace            string   `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectionCriteria) Reset()         { *m = CollectionCriteria{} }
func (m *CollectionCriteria) String() string { return proto.CompactTextString(m) }
func (*CollectionCriteria) ProtoMessage()    {}
func (*CollectionCriteria) Descriptor() ([]byte, []int) {
	return fileDescriptor_89f245fc544906c7, []int{4}
}

func (m *CollectionCriteria) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionCriteria.Unmarshal(m, b)
}
func (m *CollectionCriteria) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionCriteria.Marshal(b, m, deterministic)
}
func (m *CollectionCriteria) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionCriteria.Merge(m, src)
}
func (m *CollectionCriteria) XXX_Size() int {
	return xxx_messageInfo_CollectionCriteria.Size(m)
}
func (m *CollectionCriteria) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionCriteria.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionCriteria proto.InternalMessageInfo

func (m *CollectionCriteria) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *CollectionCriteria) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *CollectionCriteria) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *CollectionCriteria) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func init() {
	proto.RegisterType((*CollectionConfigPackage)(nil), "sdk.common.CollectionConfigPackage")
	proto.RegisterType((*CollectionConfig)(nil), "sdk.common.CollectionConfig")
	proto.RegisterType((*StaticCollectionConfig)(nil), "sdk.common.StaticCollectionConfig")
	proto.RegisterType((*CollectionPolicyConfig)(nil), "sdk.common.CollectionPolicyConfig")
	proto.RegisterType((*CollectionCriteria)(nil), "sdk.common.CollectionCriteria")
}

func init() { proto.RegisterFile("common/collection.proto", fileDescriptor_89f245fc544906c7) }

var fileDescriptor_89f245fc544906c7 = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0x5f, 0x6b, 0xdb, 0x3c,
	0x14, 0xc6, 0xeb, 0x36, 0x7f, 0x5e, 0x9f, 0xf0, 0xae, 0xa9, 0xca, 0x52, 0x33, 0x46, 0x17, 0xc2,
	0x2e, 0xcc, 0x36, 0x9c, 0xd1, 0x7d, 0x83, 0x86, 0x41, 0xc7, 0x02, 0x0b, 0xee, 0x60, 0xd0, 0x1b,
	0xa3, 0xc8, 0xa7, 0x8e, 0xa8, 0x2d, 0xb9, 0xb2, 0x92, 0xc5, 0x97, 0xfb, 0x8a, 0xfb, 0x44, 0x23,
	0x92, 0x1d, 0xbb, 0x21, 0x77, 0xd1, 0xf3, 0xfc, 0xce, 0xc9, 0xd1, 0x79, 0x64, 0xb8, 0x62, 0x32,
	0xcb, 0xa4, 0x98, 0x32, 0x99, 0xa6, 0xc8, 0x34, 0x97, 0x22, 0xc8, 0x95, 0xd4, 0x92, 0xf4, 0xac,
	0xf1, 0xe6, 0x75, 0x05, 0xe4, 0x32, 0xe5, 0x8c, 0x63, 0x61, 0xed, 0xc9, 0x77, 0xb8, 0x9a, 0xed,
	0x4b, 0x66, 0x52, 0x3c, 0xf2, 0x64, 0x41, 0xd9, 0x13, 0x4d, 0x90, 0x7c, 0x86, 0x1e, 0x33, 0x82,
	0xe7, 0x8c, 0xcf, 0xfc, 0xc1, 0x8d, 0x17, 0xd8, 0x16, 0xc1, 0x61, 0x41, 0x58, 0x71, 0x93, 0x12,
	0x86, 0x87, 0x1e, 0x79, 0x00, 0xaf, 0xd0, 0x54, 0x73, 0x16, 0x35, 0xa3, 0x45, 0xfb, 0xbe, 0x8e,
	0x3f, 0xb8, 0xb9, 0xae, 0xfb, 0xde, 0x1b, 0xee, 0xb0, 0xc3, 0xdd, 0x49, 0x38, 0x2a, 0x8e, 0x3a,
	0xb7, 0x2e, 0xf4, 0x73, 0x5a, 0xa6, 0x92, 0xc6, 0x93, 0xbf, 0xa7, 0x30, 0x3a, 0x5e, 0x4f, 0x08,
	0x74, 0x04, 0xcd, 0xd0, 0xfc, 0x9b, 0x1b, 0x9a, 0xdf, 0x64, 0x0e, 0x24, 0xc3, 0x6c, 0x89, 0x2a,
	0x92, 0x2a, 0x29, 0x22, 0xb3, 0x94, 0xd2, 0x3b, 0x7d, 0x39, 0x4f, 0xd3, 0x69, 0x61, 0xfc, 0xea,
	0xb6, 0x43, 0x5b, 0xf9, 0x43, 0x25, 0x85, 0xd5, 0x49, 0x00, 0x97, 0x0a, 0x9f, 0xd7, 0x5c, 0x61,
	0x1c, 0xe5, 0x88, 0x2a, 0x62, 0x72, 0x2d, 0xb4, 0x77, 0x36, 0x76, 0xfc, 0x6e, 0x78, 0x51, 0x5b,
	0x0b, 0x44, 0x35, 0xdb, 0x19, 0xe4, 0x13, 0x90, 0x8c, 0x6e, 0x79, 0xb6, 0xce, 0xda, 0x78, 0xc7,
	0xe0, 0xc3, 0xca, 0x69, 0xe8, 0x09, 0xfc, 0xbf, 0x4c, 0x25, 0x7b, 0x8a, 0xb4, 0x8c, 0x52, 0xbe,
	0x41, 0xaf, 0x3b, 0x76, 0xfc, 0x4e, 0x38, 0x30, 0xe2, 0x4f, 0x39, 0xe7, 0x1b, 0x24, 0x3e, 0x0c,
	0xeb, 0xfb, 0x88, 0xb4, 0x8c, 0x14, 0xd2, 0xd8, 0xeb, 0x8d, 0x1d, 0xff, 0xbf, 0xf0, 0x55, 0x35,
	0xad, 0x48, 0xcb, 0x10, 0x69, 0x4c, 0x3e, 0xc0, 0x45, 0x9b, 0xfc, 0xad, 0xb8, 0x46, 0xaf, 0x6f,
	0xd0, 0xf3, 0x06, 0xfd, 0xb5, 0x93, 0x27, 0xcf, 0x30, 0x3a, 0xbe, 0x03, 0x32, 0x87, 0x61, 0xc1,
	0x13, 0x41, 0xf5, 0x5a, 0x61, 0xbd, 0x3d, 0x9b, 0xe6, 0xbb, 0x7d, 0x9a, 0xb5, 0x6f, 0x0b, 0xbf,
	0x8a, 0x0d, 0xa6, 0x32, 0xc7, 0xbb, 0x93, 0xf0, 0xbc, 0x78, 0x69, 0xb5, 0x73, 0xfc, 0xe3, 0x00,
	0x69, 0x25, 0xb8, 0x1b, 0x43, 0x71, 0x4a, 0x3c, 0xe8, 0xb3, 0x15, 0x15, 0x02, 0xd3, 0x2a, 0xc6,
	0xfa, 0x48, 0x2e, 0xa1, 0xab, 0xb7, 0x11, 0x8f, 0x4d, 0x78, 0x6e, 0xd8, 0xd1, 0xdb, 0x6f, 0x31,
	0xb9, 0x06, 0x68, 0x5e, 0x9b, 0xc9, 0xc1, 0x0d, 0x5b, 0x0a, 0x79, 0x0b, 0xee, 0xee, 0x19, 0x14,
	0x39, 0x65, 0x68, 0xf6, 0xee, 0x86, 0x8d, 0x70, 0x7b, 0x0f, 0xef, 0xa5, 0x4a, 0x82, 0x55, 0x99,
	0xa3, 0x4a, 0x31, 0x4e, 0x50, 0x05, 0x8f, 0x74, 0xa9, 0x38, 0xb3, 0xdf, 0x4c, 0x51, 0xdd, 0xf0,
	0xe1, 0x63, 0xc2, 0xf5, 0x6a, 0xbd, 0xdc, 0x1d, 0xa7, 0x2d, 0x78, 0x6a, 0xe1, 0xa9, 0x85, 0xa7,
	0x16, 0x5e, 0xf6, 0xcc, 0xf1, 0xcb, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x4b, 0x91, 0x18,
	0xa9, 0x03, 0x00, 0x00,
}
