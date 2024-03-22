// Maintainers, please refer to the style guide here:
//     https://developers.google.com/protocol-buffers/docs/style
//
// This file provides a type that can marshal the event data a customer sees
// from our apis.  It is typically used for demo and api code in go-lang that
// needs to work with our api responses.
//
// NOTE: It MUST be kept up to date with the EventResponse message definition

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: datatrails-common-api/assets/v2/assets/eventresponsejsonapi.proto

package assets

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// EventResponseJSONAPI represents how the consumer of the events api sees the event data.
type EventResponseJSONAPI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Relative Resource Name for the operation event
	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	// relative resource name for associated asset ( asset the operation is performed on  - has to have specific behaviour enabled)
	AssetIdentity string `protobuf:"bytes,2,opt,name=asset_identity,json=assetIdentity,proto3" json:"asset_identity,omitempty"`
	// map of event attributes. Specific behaviours define required and optional event attributes for each supported operation.
	EventAttributes map[string]string `protobuf:"bytes,16,rep,name=event_attributes,json=eventAttributes,proto3" json:"event_attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// map of asset attributes. Specific behaviours define required and optional asset attributes. These attributes cause the corresponding attributes on the asset to be updated.
	AssetAttributes map[string]string `protobuf:"bytes,17,rep,name=asset_attributes,json=assetAttributes,proto3" json:"asset_attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// name of operation on this behviour
	Operation string `protobuf:"bytes,4,opt,name=operation,proto3" json:"operation,omitempty"`
	// name of this behaviour
	Behaviour string `protobuf:"bytes,14,opt,name=behaviour,proto3" json:"behaviour,omitempty"`
	// timestamp when operation was actually performed - if not provided will be set to timestamp_accepted
	TimestampDeclared string `protobuf:"bytes,5,opt,name=timestamp_declared,json=timestampDeclared,proto3" json:"timestamp_declared,omitempty"`
	// timestamp when system received operation request
	TimestampAccepted string `protobuf:"bytes,6,opt,name=timestamp_accepted,json=timestampAccepted,proto3" json:"timestamp_accepted,omitempty"`
	// timestamp for when the event was committed to a verifiable log
	TimestampCommitted string `protobuf:"bytes,7,opt,name=timestamp_committed,json=timestampCommitted,proto3" json:"timestamp_committed,omitempty"`
	// principal information associated with event - if not provided will be set to principal_accepted
	PrincipalDeclared *Principal `protobuf:"bytes,8,opt,name=principal_declared,json=principalDeclared,proto3" json:"principal_declared,omitempty"`
	// principal logged into the system that performed the operation
	PrincipalAccepted *Principal `protobuf:"bytes,9,opt,name=principal_accepted,json=principalAccepted,proto3" json:"principal_accepted,omitempty"`
	// indicated if operation has been committed to the blockchain
	ConfirmationStatus string `protobuf:"bytes,10,opt,name=confirmation_status,json=confirmationStatus,proto3" json:"confirmation_status,omitempty"`
	// hash of transaction committing this operation on blockchain
	TransactionId string `protobuf:"bytes,11,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	// block number of committing transaction
	BlockNumber uint64 `protobuf:"varint,12,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	// transaction index of committing transaction
	TransactionIndex uint64 `protobuf:"varint,13,opt,name=transaction_index,json=transactionIndex,proto3" json:"transaction_index,omitempty"`
	// wallet address for the creator of this event
	From           string `protobuf:"bytes,15,opt,name=from,proto3" json:"from,omitempty"`
	TenantIdentity string `protobuf:"bytes,18,opt,name=tenant_identity,json=tenantIdentity,proto3" json:"tenant_identity,omitempty"`
	// proof details for proof_mechanism MERKLE_LOG
	MerklelogEntry *MerkleLogEntry `protobuf:"bytes,19,opt,name=merklelog_entry,json=merklelogEntry,proto3" json:"merklelog_entry,omitempty"`
}

func (x *EventResponseJSONAPI) Reset() {
	*x = EventResponseJSONAPI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventResponseJSONAPI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventResponseJSONAPI) ProtoMessage() {}

func (x *EventResponseJSONAPI) ProtoReflect() protoreflect.Message {
	mi := &file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventResponseJSONAPI.ProtoReflect.Descriptor instead.
func (*EventResponseJSONAPI) Descriptor() ([]byte, []int) {
	return file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_rawDescGZIP(), []int{0}
}

func (x *EventResponseJSONAPI) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *EventResponseJSONAPI) GetAssetIdentity() string {
	if x != nil {
		return x.AssetIdentity
	}
	return ""
}

func (x *EventResponseJSONAPI) GetEventAttributes() map[string]string {
	if x != nil {
		return x.EventAttributes
	}
	return nil
}

func (x *EventResponseJSONAPI) GetAssetAttributes() map[string]string {
	if x != nil {
		return x.AssetAttributes
	}
	return nil
}

func (x *EventResponseJSONAPI) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *EventResponseJSONAPI) GetBehaviour() string {
	if x != nil {
		return x.Behaviour
	}
	return ""
}

func (x *EventResponseJSONAPI) GetTimestampDeclared() string {
	if x != nil {
		return x.TimestampDeclared
	}
	return ""
}

func (x *EventResponseJSONAPI) GetTimestampAccepted() string {
	if x != nil {
		return x.TimestampAccepted
	}
	return ""
}

func (x *EventResponseJSONAPI) GetTimestampCommitted() string {
	if x != nil {
		return x.TimestampCommitted
	}
	return ""
}

func (x *EventResponseJSONAPI) GetPrincipalDeclared() *Principal {
	if x != nil {
		return x.PrincipalDeclared
	}
	return nil
}

func (x *EventResponseJSONAPI) GetPrincipalAccepted() *Principal {
	if x != nil {
		return x.PrincipalAccepted
	}
	return nil
}

func (x *EventResponseJSONAPI) GetConfirmationStatus() string {
	if x != nil {
		return x.ConfirmationStatus
	}
	return ""
}

func (x *EventResponseJSONAPI) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *EventResponseJSONAPI) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *EventResponseJSONAPI) GetTransactionIndex() uint64 {
	if x != nil {
		return x.TransactionIndex
	}
	return 0
}

func (x *EventResponseJSONAPI) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *EventResponseJSONAPI) GetTenantIdentity() string {
	if x != nil {
		return x.TenantIdentity
	}
	return ""
}

func (x *EventResponseJSONAPI) GetMerklelogEntry() *MerkleLogEntry {
	if x != nil {
		return x.MerklelogEntry
	}
	return nil
}

var File_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto protoreflect.FileDescriptor

var file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_rawDesc = []byte{
	0x0a, 0x41, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x73, 0x2d, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76,
	0x32, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x6a, 0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x2e, 0x76,
	0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x32, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x73, 0x2d, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f,
	0x76, 0x32, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x69, 0x6c,
	0x73, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x70, 0x72,
	0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x64,
	0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x73, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x6c, 0x6f, 0x67, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x17, 0x0a, 0x14, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4a, 0x53, 0x4f, 0x4e,
	0x41, 0x50, 0x49, 0x12, 0x3f, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0x92, 0x41, 0x20, 0x32, 0x1c, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x20,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x40, 0x01, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x80, 0x01, 0x0a, 0x0e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x59, 0x92,
	0x41, 0x56, 0x32, 0x52, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x20, 0x6f, 0x66, 0x20,
	0x61, 0x20, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x20, 0x61, 0x73, 0x73, 0x65, 0x74, 0x20,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x60, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73,
	0x2f, 0x31, 0x31, 0x62, 0x66, 0x35, 0x62, 0x33, 0x37, 0x2d, 0x65, 0x30, 0x62, 0x38, 0x2d, 0x34,
	0x32, 0x65, 0x30, 0x2d, 0x38, 0x64, 0x63, 0x66, 0x2d, 0x64, 0x63, 0x38, 0x63, 0x34, 0x61, 0x65,
	0x66, 0x63, 0x30, 0x30, 0x30, 0x60, 0x40, 0x01, 0x52, 0x0d, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x8e, 0x01, 0x0a, 0x10, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x10, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x37, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x2e, 0x76,
	0x32, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4a,
	0x53, 0x4f, 0x4e, 0x41, 0x50, 0x49, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x2a, 0x92, 0x41, 0x27,
	0x32, 0x25, 0x6b, 0x65, 0x79, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x6d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x20, 0x6f, 0x66, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x20, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x8e, 0x01, 0x0a, 0x10, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x11, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x2e,
	0x76, 0x32, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4a, 0x53, 0x4f, 0x4e, 0x41, 0x50, 0x49, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x2a, 0x92, 0x41,
	0x27, 0x32, 0x25, 0x6b, 0x65, 0x79, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x6d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x73, 0x73, 0x65, 0x74, 0x20, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x0f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x58, 0x0a, 0x09, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3a, 0x92, 0x41,
	0x37, 0x32, 0x30, 0x54, 0x68, 0x65, 0x20, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x20, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x20, 0x60, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x60, 0x40, 0x01, 0x78, 0x80, 0x20, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x5c, 0x0a, 0x09, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x75, 0x72,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3e, 0x92, 0x41, 0x3b, 0x32, 0x34, 0x54, 0x68, 0x65,
	0x20, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20,
	0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x20, 0x60, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65,
	0x60, 0x40, 0x01, 0x78, 0x80, 0x20, 0x52, 0x09, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x75,
	0x72, 0x12, 0x64, 0x0a, 0x12, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x64,
	0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x35, 0x92,
	0x41, 0x32, 0x32, 0x2e, 0x52, 0x46, 0x43, 0x20, 0x33, 0x33, 0x33, 0x39, 0x20, 0x74, 0x69, 0x6d,
	0x65, 0x20, 0x6f, 0x66, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x20, 0x61, 0x73, 0x20, 0x64, 0x65,
	0x63, 0x6c, 0x61, 0x72, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73,
	0x65, 0x72, 0x40, 0x01, 0x52, 0x11, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x44,
	0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x64, 0x12, 0x66, 0x0a, 0x12, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x37, 0x92, 0x41, 0x34, 0x32, 0x30, 0x52, 0x46, 0x43, 0x20, 0x33, 0x33,
	0x33, 0x39, 0x20, 0x74, 0x69, 0x6d, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x20, 0x61, 0x73, 0x20, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x40, 0x01, 0x52, 0x11, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x12,
	0x70, 0x0a, 0x13, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3f, 0x92, 0x41,
	0x3c, 0x32, 0x38, 0x52, 0x46, 0x43, 0x20, 0x33, 0x33, 0x33, 0x39, 0x20, 0x74, 0x69, 0x6d, 0x65,
	0x20, 0x6f, 0x66, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x20, 0x61, 0x73, 0x20, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x65, 0x64, 0x20, 0x69, 0x6e, 0x20, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x20, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x40, 0x01, 0x52, 0x12, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65,
	0x64, 0x12, 0x71, 0x0a, 0x12, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x5f, 0x64,
	0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x69,
	0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x42, 0x29, 0x92, 0x41, 0x26, 0x32, 0x1e, 0x70, 0x72, 0x69,
	0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x20,
	0x62, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x40, 0x01, 0x9a, 0x02, 0x01,
	0x06, 0x52, 0x11, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x44, 0x65, 0x63, 0x6c,
	0x61, 0x72, 0x65, 0x64, 0x12, 0x73, 0x0a, 0x12, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61,
	0x6c, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x2e,
	0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x42, 0x2b, 0x92, 0x41, 0x28, 0x32, 0x20,
	0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x20, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x40, 0x01, 0x9a, 0x02, 0x01, 0x06, 0x52, 0x11, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61,
	0x6c, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x12, 0x83, 0x01, 0x0a, 0x13, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x52, 0x92, 0x41, 0x4f, 0x32, 0x47, 0x69, 0x6e,
	0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x20, 0x69, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x20, 0x68, 0x61, 0x73, 0x20, 0x62, 0x65, 0x65, 0x6e, 0x20, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x74, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x40, 0x01, 0x9a, 0x02, 0x01, 0x07, 0x52, 0x12, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x7b, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x54, 0x92, 0x41, 0x51, 0x32, 0x4c, 0x68, 0x61,
	0x73, 0x68, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x61, 0x73, 0x20, 0x61, 0x20, 0x68, 0x65, 0x78, 0x20, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x60, 0x30, 0x78, 0x31, 0x31, 0x62, 0x66, 0x35, 0x62, 0x33,
	0x37, 0x65, 0x30, 0x62, 0x38, 0x34, 0x32, 0x65, 0x30, 0x38, 0x64, 0x63, 0x66, 0x64, 0x63, 0x38,
	0x63, 0x34, 0x61, 0x65, 0x66, 0x63, 0x30, 0x30, 0x30, 0x60, 0x78, 0x80, 0x20, 0x52, 0x0d, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x4f, 0x0a, 0x0c,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x2c, 0x92, 0x41, 0x29, 0x32, 0x25, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x20,
	0x6f, 0x66, 0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x20, 0x77,
	0x61, 0x73, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x20, 0x6f, 0x6e, 0x40, 0x01,
	0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x58, 0x0a,
	0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x42, 0x2b, 0x92, 0x41, 0x28, 0x32, 0x24, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x20, 0x6f, 0x66, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x20, 0x77, 0x69,
	0x74, 0x68, 0x69, 0x6e, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x20, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x40, 0x01, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x47, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33, 0x92, 0x41, 0x30, 0x32, 0x2c, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x20, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68,
	0x69, 0x73, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x40, 0x01, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x12, 0x63, 0x0a, 0x0f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3a, 0x92, 0x41, 0x37, 0x32, 0x32,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x78, 0x80, 0x08, 0x52, 0x0e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x76, 0x0a, 0x0f, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x6c,
	0x6f, 0x67, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x4d, 0x65,
	0x72, 0x6b, 0x6c, 0x65, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x2f, 0x92, 0x41,
	0x2c, 0x32, 0x27, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x6d, 0x65,
	0x72, 0x6b, 0x6c, 0x65, 0x20, 0x6d, 0x6d, 0x72, 0x20, 0x6c, 0x6f, 0x67, 0x20, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x78, 0x80, 0x08, 0x52, 0x0e, 0x6d,
	0x65, 0x72, 0x6b, 0x6c, 0x65, 0x6c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x1a, 0x42, 0x0a,
	0x14, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x42, 0x0a, 0x14, 0x41, 0x73, 0x73, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x80, 0x07, 0x92, 0x41, 0xfc, 0x06, 0x0a, 0x1a, 0x32, 0x18,
	0x54, 0x68, 0x69, 0x73, 0x20, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x73, 0x20, 0x61,
	0x6e, 0x20, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x32, 0xdd, 0x06, 0x7b, 0x20, 0x22, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x3a, 0x20, 0x22, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73,
	0x2f, 0x61, 0x64, 0x64, 0x33, 0x30, 0x32, 0x33, 0x35, 0x2d, 0x31, 0x34, 0x32, 0x34, 0x2d, 0x34,
	0x66, 0x64, 0x61, 0x2d, 0x38, 0x34, 0x30, 0x61, 0x2d, 0x64, 0x35, 0x65, 0x66, 0x38, 0x32, 0x63,
	0x34, 0x63, 0x39, 0x36, 0x66, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x31, 0x31, 0x62,
	0x66, 0x35, 0x62, 0x33, 0x37, 0x2d, 0x65, 0x30, 0x62, 0x38, 0x2d, 0x34, 0x32, 0x65, 0x30, 0x2d,
	0x38, 0x64, 0x63, 0x66, 0x2d, 0x64, 0x63, 0x38, 0x63, 0x34, 0x61, 0x65, 0x66, 0x63, 0x30, 0x30,
	0x30, 0x22, 0x2c, 0x20, 0x22, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x22, 0x3a, 0x20, 0x22, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x61, 0x64,
	0x64, 0x33, 0x30, 0x32, 0x33, 0x35, 0x2d, 0x31, 0x34, 0x32, 0x34, 0x2d, 0x34, 0x66, 0x64, 0x61,
	0x2d, 0x38, 0x34, 0x30, 0x61, 0x2d, 0x64, 0x35, 0x65, 0x66, 0x38, 0x32, 0x63, 0x34, 0x63, 0x39,
	0x36, 0x66, 0x22, 0x2c, 0x20, 0x22, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x3a, 0x20, 0x22, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x2c, 0x20, 0x22, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x22, 0x3a, 0x20, 0x22, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x2c, 0x20, 0x22, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0x3a, 0x20, 0x7b,
	0x20, 0x20, 0x20, 0x20, 0x22, 0x61, 0x72, 0x63, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x50, 0x61, 0x69, 0x6e, 0x74, 0x22, 0x2c,
	0x20, 0x20, 0x20, 0x20, 0x22, 0x61, 0x72, 0x63, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x20, 0x22, 0x50, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x64, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x20, 0x7d, 0x2c, 0x20, 0x22, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22,
	0x3a, 0x20, 0x7b, 0x20, 0x20, 0x20, 0x20, 0x22, 0x63, 0x6f, 0x6c, 0x6f, 0x75, 0x72, 0x22, 0x3a,
	0x20, 0x22, 0x4d, 0x69, 0x64, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x20, 0x42, 0x6c, 0x75, 0x65, 0x22,
	0x20, 0x7d, 0x2c, 0x20, 0x22, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x61,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x32, 0x30, 0x31, 0x39, 0x2d,
	0x31, 0x31, 0x2d, 0x32, 0x37, 0x54, 0x31, 0x34, 0x3a, 0x34, 0x34, 0x3a, 0x31, 0x39, 0x5a, 0x22,
	0x2c, 0x20, 0x22, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x64, 0x65, 0x63,
	0x6c, 0x61, 0x72, 0x65, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x32, 0x30, 0x31, 0x39, 0x2d, 0x31, 0x31,
	0x2d, 0x32, 0x37, 0x54, 0x31, 0x34, 0x3a, 0x34, 0x34, 0x3a, 0x31, 0x39, 0x5a, 0x22, 0x2c, 0x20,
	0x22, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x74, 0x65, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x32, 0x30, 0x31, 0x39, 0x2d, 0x31, 0x31, 0x2d,
	0x32, 0x37, 0x54, 0x31, 0x34, 0x3a, 0x34, 0x34, 0x3a, 0x31, 0x39, 0x5a, 0x22, 0x2c, 0x20, 0x22,
	0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x5f, 0x64, 0x65, 0x63, 0x6c, 0x61, 0x72,
	0x65, 0x64, 0x22, 0x3a, 0x20, 0x7b, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x72, 0x22, 0x3a, 0x20, 0x22, 0x6a, 0x6f, 0x62, 0x2e, 0x69, 0x64, 0x70, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x31, 0x32, 0x33, 0x34, 0x22, 0x2c, 0x20, 0x22, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x22, 0x3a, 0x22, 0x62, 0x6f, 0x62, 0x40, 0x6a, 0x6f, 0x62, 0x22, 0x20,
	0x20, 0x7d, 0x2c, 0x20, 0x20, 0x22, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x5f,
	0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x22, 0x3a, 0x20, 0x7b, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x22, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x22, 0x3a, 0x20, 0x22, 0x6a, 0x6f, 0x62, 0x2e,
	0x69, 0x64, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x31, 0x32, 0x33, 0x34, 0x22,
	0x2c, 0x20, 0x22, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x3a, 0x22, 0x62, 0x6f, 0x62,
	0x40, 0x6a, 0x6f, 0x62, 0x22, 0x20, 0x7d, 0x2c, 0x20, 0x22, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3a, 0x20,
	0x22, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x45, 0x44, 0x22, 0x2c, 0x20, 0x22, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x3a, 0x20, 0x31, 0x32, 0x2c,
	0x20, 0x22, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x22, 0x3a, 0x20, 0x35, 0x2c, 0x20, 0x22, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x30, 0x78, 0x30, 0x37,
	0x35, 0x36, 0x39, 0x22, 0x2c, 0x20, 0x22, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x3a, 0x20, 0x22, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x2f, 0x38, 0x65, 0x30, 0x62, 0x36, 0x30, 0x30, 0x63, 0x2d, 0x38, 0x32, 0x33, 0x34, 0x2d, 0x34,
	0x33, 0x65, 0x34, 0x2d, 0x38, 0x36, 0x30, 0x63, 0x2d, 0x65, 0x39, 0x35, 0x62, 0x64, 0x63, 0x64,
	0x36, 0x39, 0x35, 0x61, 0x39, 0x22, 0x20, 0x7d, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x69, 0x6c,
	0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x73, 0x2d,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x3b,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_rawDescOnce sync.Once
	file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_rawDescData = file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_rawDesc
)

func file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_rawDescGZIP() []byte {
	file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_rawDescOnce.Do(func() {
		file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_rawDescData)
	})
	return file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_rawDescData
}

var file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_goTypes = []interface{}{
	(*EventResponseJSONAPI)(nil), // 0: archivist.v2.EventResponseJSONAPI
	nil,                          // 1: archivist.v2.EventResponseJSONAPI.EventAttributesEntry
	nil,                          // 2: archivist.v2.EventResponseJSONAPI.AssetAttributesEntry
	(*Principal)(nil),            // 3: archivist.v2.Principal
	(*MerkleLogEntry)(nil),       // 4: archivist.v2.MerkleLogEntry
}
var file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_depIdxs = []int32{
	1, // 0: archivist.v2.EventResponseJSONAPI.event_attributes:type_name -> archivist.v2.EventResponseJSONAPI.EventAttributesEntry
	2, // 1: archivist.v2.EventResponseJSONAPI.asset_attributes:type_name -> archivist.v2.EventResponseJSONAPI.AssetAttributesEntry
	3, // 2: archivist.v2.EventResponseJSONAPI.principal_declared:type_name -> archivist.v2.Principal
	3, // 3: archivist.v2.EventResponseJSONAPI.principal_accepted:type_name -> archivist.v2.Principal
	4, // 4: archivist.v2.EventResponseJSONAPI.merklelog_entry:type_name -> archivist.v2.MerkleLogEntry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_init() }
func file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_init() {
	if File_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto != nil {
		return
	}
	file_datatrails_common_api_assets_v2_assets_enums_proto_init()
	file_datatrails_common_api_assets_v2_assets_principal_proto_init()
	file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventResponseJSONAPI); i {
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
			RawDescriptor: file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_goTypes,
		DependencyIndexes: file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_depIdxs,
		MessageInfos:      file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_msgTypes,
	}.Build()
	File_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto = out.File
	file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_rawDesc = nil
	file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_goTypes = nil
	file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_depIdxs = nil
}
