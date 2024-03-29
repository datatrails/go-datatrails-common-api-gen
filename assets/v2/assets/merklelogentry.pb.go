// Maintainers, please refer to the style guide here:
//     https://developers.google.com/protocol-buffers/docs/style

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: datatrails-common-api/assets/v2/assets/merklelogentry.proto

package assets

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MerkeLogCommit provides the log entry details for a single mmr leaf.
type MerkleLogCommit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The mmr index
	Index uint64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	// time ordered and strictly unique per tenant. system wide unique with very
	// reasonable operational assumptions. prefixed with time epoch if len > 8
	// bytes (after conversion back from hex).
	Idtimestamp string `protobuf:"bytes,2,opt,name=idtimestamp,proto3" json:"idtimestamp,omitempty"`
}

func (x *MerkleLogCommit) Reset() {
	*x = MerkleLogCommit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerkleLogCommit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerkleLogCommit) ProtoMessage() {}

func (x *MerkleLogCommit) ProtoReflect() protoreflect.Message {
	mi := &file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerkleLogCommit.ProtoReflect.Descriptor instead.
func (*MerkleLogCommit) Descriptor() ([]byte, []int) {
	return file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_rawDescGZIP(), []int{0}
}

func (x *MerkleLogCommit) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *MerkleLogCommit) GetIdtimestamp() string {
	if x != nil {
		return x.Idtimestamp
	}
	return ""
}

type MerkleLogConfirm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The following correspond to mmrblobs.MMRState
	MmrSize uint64 `protobuf:"varint,1,opt,name=mmr_size,json=mmrSize,proto3" json:"mmr_size,omitempty"`
	Root    []byte `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
	// The regular unix time the root was signed
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The idtimestamp of the last leaf under mmr_size. prefixed with time epoch if len > 8 bytes (after conversion back from hex)
	Idtimestamp string `protobuf:"bytes,4,opt,name=idtimestamp,proto3" json:"idtimestamp,omitempty"`
	// The signed merkle tree head state at mmr_size. Contains COSE Sign1 formatted message.
	SignedTreeHead []byte `protobuf:"bytes,5,opt,name=signed_tree_head,json=signedTreeHead,proto3" json:"signed_tree_head,omitempty"`
}

func (x *MerkleLogConfirm) Reset() {
	*x = MerkleLogConfirm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerkleLogConfirm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerkleLogConfirm) ProtoMessage() {}

func (x *MerkleLogConfirm) ProtoReflect() protoreflect.Message {
	mi := &file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerkleLogConfirm.ProtoReflect.Descriptor instead.
func (*MerkleLogConfirm) Descriptor() ([]byte, []int) {
	return file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_rawDescGZIP(), []int{1}
}

func (x *MerkleLogConfirm) GetMmrSize() uint64 {
	if x != nil {
		return x.MmrSize
	}
	return 0
}

func (x *MerkleLogConfirm) GetRoot() []byte {
	if x != nil {
		return x.Root
	}
	return nil
}

func (x *MerkleLogConfirm) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *MerkleLogConfirm) GetIdtimestamp() string {
	if x != nil {
		return x.Idtimestamp
	}
	return ""
}

func (x *MerkleLogConfirm) GetSignedTreeHead() []byte {
	if x != nil {
		return x.SignedTreeHead
	}
	return nil
}

type MerkleLogUnequivocal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MerkleLogUnequivocal) Reset() {
	*x = MerkleLogUnequivocal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerkleLogUnequivocal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerkleLogUnequivocal) ProtoMessage() {}

func (x *MerkleLogUnequivocal) ProtoReflect() protoreflect.Message {
	mi := &file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerkleLogUnequivocal.ProtoReflect.Descriptor instead.
func (*MerkleLogUnequivocal) Descriptor() ([]byte, []int) {
	return file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_rawDescGZIP(), []int{2}
}

// The message sent from forestrie to avid notifying that the corresponding
// event is commited to the tenants log.
type MerkleLogCommitMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The tenant identity and the event identity for the committed event.
	TenantIdentity string `protobuf:"bytes,1,opt,name=tenant_identity,json=tenantIdentity,proto3" json:"tenant_identity,omitempty"`
	EventIdentity  string `protobuf:"bytes,2,opt,name=event_identity,json=eventIdentity,proto3" json:"event_identity,omitempty"`
	// The time portion of idtimestamp that contributed to the hash of the event
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Commit    *MerkleLogCommit       `protobuf:"bytes,4,opt,name=commit,proto3" json:"commit,omitempty"`
}

func (x *MerkleLogCommitMessage) Reset() {
	*x = MerkleLogCommitMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerkleLogCommitMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerkleLogCommitMessage) ProtoMessage() {}

func (x *MerkleLogCommitMessage) ProtoReflect() protoreflect.Message {
	mi := &file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerkleLogCommitMessage.ProtoReflect.Descriptor instead.
func (*MerkleLogCommitMessage) Descriptor() ([]byte, []int) {
	return file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_rawDescGZIP(), []int{3}
}

func (x *MerkleLogCommitMessage) GetTenantIdentity() string {
	if x != nil {
		return x.TenantIdentity
	}
	return ""
}

func (x *MerkleLogCommitMessage) GetEventIdentity() string {
	if x != nil {
		return x.EventIdentity
	}
	return ""
}

func (x *MerkleLogCommitMessage) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *MerkleLogCommitMessage) GetCommit() *MerkleLogCommit {
	if x != nil {
		return x.Commit
	}
	return nil
}

type MerkleLogConfirmMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantIdentity string            `protobuf:"bytes,1,opt,name=tenant_identity,json=tenantIdentity,proto3" json:"tenant_identity,omitempty"`
	Confirm        *MerkleLogConfirm `protobuf:"bytes,2,opt,name=confirm,proto3" json:"confirm,omitempty"`
}

func (x *MerkleLogConfirmMessage) Reset() {
	*x = MerkleLogConfirmMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerkleLogConfirmMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerkleLogConfirmMessage) ProtoMessage() {}

func (x *MerkleLogConfirmMessage) ProtoReflect() protoreflect.Message {
	mi := &file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerkleLogConfirmMessage.ProtoReflect.Descriptor instead.
func (*MerkleLogConfirmMessage) Descriptor() ([]byte, []int) {
	return file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_rawDescGZIP(), []int{4}
}

func (x *MerkleLogConfirmMessage) GetTenantIdentity() string {
	if x != nil {
		return x.TenantIdentity
	}
	return ""
}

func (x *MerkleLogConfirmMessage) GetConfirm() *MerkleLogConfirm {
	if x != nil {
		return x.Confirm
	}
	return nil
}

type MerkleLogUnequivocalMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantIdentity string                `protobuf:"bytes,1,opt,name=tenant_identity,json=tenantIdentity,proto3" json:"tenant_identity,omitempty"`
	Unequivocal    *MerkleLogUnequivocal `protobuf:"bytes,2,opt,name=unequivocal,proto3" json:"unequivocal,omitempty"`
}

func (x *MerkleLogUnequivocalMessage) Reset() {
	*x = MerkleLogUnequivocalMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerkleLogUnequivocalMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerkleLogUnequivocalMessage) ProtoMessage() {}

func (x *MerkleLogUnequivocalMessage) ProtoReflect() protoreflect.Message {
	mi := &file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerkleLogUnequivocalMessage.ProtoReflect.Descriptor instead.
func (*MerkleLogUnequivocalMessage) Descriptor() ([]byte, []int) {
	return file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_rawDescGZIP(), []int{5}
}

func (x *MerkleLogUnequivocalMessage) GetTenantIdentity() string {
	if x != nil {
		return x.TenantIdentity
	}
	return ""
}

func (x *MerkleLogUnequivocalMessage) GetUnequivocal() *MerkleLogUnequivocal {
	if x != nil {
		return x.Unequivocal
	}
	return nil
}

// The details stored in the SaaS db for a proof mech MERKLE_LOG commitment
type MerkleLogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Event trust level COMMITTED details
	Commit *MerkleLogCommit `protobuf:"bytes,1,opt,name=commit,proto3" json:"commit,omitempty"`
	// Event trust level CONFIRMED details
	Confirm *MerkleLogConfirm `protobuf:"bytes,2,opt,name=confirm,proto3" json:"confirm,omitempty"`
	// Event trust level UNEQUIVOCAL details
	Unequivocal *MerkleLogUnequivocal `protobuf:"bytes,3,opt,name=unequivocal,proto3" json:"unequivocal,omitempty"`
}

func (x *MerkleLogEntry) Reset() {
	*x = MerkleLogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerkleLogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerkleLogEntry) ProtoMessage() {}

func (x *MerkleLogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerkleLogEntry.ProtoReflect.Descriptor instead.
func (*MerkleLogEntry) Descriptor() ([]byte, []int) {
	return file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_rawDescGZIP(), []int{6}
}

func (x *MerkleLogEntry) GetCommit() *MerkleLogCommit {
	if x != nil {
		return x.Commit
	}
	return nil
}

func (x *MerkleLogEntry) GetConfirm() *MerkleLogConfirm {
	if x != nil {
		return x.Confirm
	}
	return nil
}

func (x *MerkleLogEntry) GetUnequivocal() *MerkleLogUnequivocal {
	if x != nil {
		return x.Unequivocal
	}
	return nil
}

var File_datatrails_common_api_assets_v2_assets_merklelogentry_proto protoreflect.FileDescriptor

var file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x73, 0x2d, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76,
	0x32, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x6c,
	0x6f, 0x67, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x0f,
	0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x64, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x64, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xab, 0x01, 0x0a, 0x10, 0x4d, 0x65, 0x72, 0x6b,
	0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x12, 0x19, 0x0a, 0x08,
	0x6d, 0x6d, 0x72, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x6d, 0x6d, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x64, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x69, 0x64, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x28, 0x0a, 0x10, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x74, 0x72, 0x65, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x72, 0x65,
	0x65, 0x48, 0x65, 0x61, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x4c,
	0x6f, 0x67, 0x55, 0x6e, 0x65, 0x71, 0x75, 0x69, 0x76, 0x6f, 0x63, 0x61, 0x6c, 0x22, 0xd9, 0x01,
	0x0a, 0x16, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x35, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x2e, 0x76,
	0x32, 0x2e, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x22, 0x7c, 0x0a, 0x17, 0x4d, 0x65, 0x72,
	0x6b, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x38, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x4d, 0x65,
	0x72, 0x6b, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x22, 0x8c, 0x01, 0x0a, 0x1b, 0x4d, 0x65, 0x72, 0x6b,
	0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x55, 0x6e, 0x65, 0x71, 0x75, 0x69, 0x76, 0x6f, 0x63, 0x61, 0x6c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x44, 0x0a, 0x0b, 0x75, 0x6e, 0x65, 0x71, 0x75, 0x69, 0x76, 0x6f, 0x63, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73,
	0x74, 0x2e, 0x76, 0x32, 0x2e, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x55, 0x6e,
	0x65, 0x71, 0x75, 0x69, 0x76, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x0b, 0x75, 0x6e, 0x65, 0x71, 0x75,
	0x69, 0x76, 0x6f, 0x63, 0x61, 0x6c, 0x22, 0xc7, 0x01, 0x0a, 0x0e, 0x4d, 0x65, 0x72, 0x6b, 0x6c,
	0x65, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x35, 0x0a, 0x06, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x4c,
	0x6f, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x12, 0x38, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x32,
	0x2e, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x12, 0x44, 0x0a, 0x0b, 0x75, 0x6e,
	0x65, 0x71, 0x75, 0x69, 0x76, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x4d,
	0x65, 0x72, 0x6b, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x55, 0x6e, 0x65, 0x71, 0x75, 0x69, 0x76, 0x6f,
	0x63, 0x61, 0x6c, 0x52, 0x0b, 0x75, 0x6e, 0x65, 0x71, 0x75, 0x69, 0x76, 0x6f, 0x63, 0x61, 0x6c,
	0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x64, 0x61, 0x74,
	0x61, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x73, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x61,
	0x70, 0x69, 0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x32,
	0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x3b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_rawDescOnce sync.Once
	file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_rawDescData = file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_rawDesc
)

func file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_rawDescGZIP() []byte {
	file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_rawDescOnce.Do(func() {
		file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_rawDescData = protoimpl.X.CompressGZIP(file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_rawDescData)
	})
	return file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_rawDescData
}

var file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_goTypes = []interface{}{
	(*MerkleLogCommit)(nil),             // 0: archivist.v2.MerkleLogCommit
	(*MerkleLogConfirm)(nil),            // 1: archivist.v2.MerkleLogConfirm
	(*MerkleLogUnequivocal)(nil),        // 2: archivist.v2.MerkleLogUnequivocal
	(*MerkleLogCommitMessage)(nil),      // 3: archivist.v2.MerkleLogCommitMessage
	(*MerkleLogConfirmMessage)(nil),     // 4: archivist.v2.MerkleLogConfirmMessage
	(*MerkleLogUnequivocalMessage)(nil), // 5: archivist.v2.MerkleLogUnequivocalMessage
	(*MerkleLogEntry)(nil),              // 6: archivist.v2.MerkleLogEntry
	(*timestamppb.Timestamp)(nil),       // 7: google.protobuf.Timestamp
}
var file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_depIdxs = []int32{
	7, // 0: archivist.v2.MerkleLogCommitMessage.timestamp:type_name -> google.protobuf.Timestamp
	0, // 1: archivist.v2.MerkleLogCommitMessage.commit:type_name -> archivist.v2.MerkleLogCommit
	1, // 2: archivist.v2.MerkleLogConfirmMessage.confirm:type_name -> archivist.v2.MerkleLogConfirm
	2, // 3: archivist.v2.MerkleLogUnequivocalMessage.unequivocal:type_name -> archivist.v2.MerkleLogUnequivocal
	0, // 4: archivist.v2.MerkleLogEntry.commit:type_name -> archivist.v2.MerkleLogCommit
	1, // 5: archivist.v2.MerkleLogEntry.confirm:type_name -> archivist.v2.MerkleLogConfirm
	2, // 6: archivist.v2.MerkleLogEntry.unequivocal:type_name -> archivist.v2.MerkleLogUnequivocal
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_init() }
func file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_init() {
	if File_datatrails_common_api_assets_v2_assets_merklelogentry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MerkleLogCommit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MerkleLogConfirm); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MerkleLogUnequivocal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MerkleLogCommitMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MerkleLogConfirmMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MerkleLogUnequivocalMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MerkleLogEntry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_goTypes,
		DependencyIndexes: file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_depIdxs,
		MessageInfos:      file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_msgTypes,
	}.Build()
	File_datatrails_common_api_assets_v2_assets_merklelogentry_proto = out.File
	file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_rawDesc = nil
	file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_goTypes = nil
	file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_depIdxs = nil
}
