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
// 	protoc-gen-go v1.36.6
// 	protoc        v4.24.3
// source: datatrails-common-api/assets/v2/assets/eventresponsejsonapi.proto

package assets

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// EventResponseJSONAPI represents how the consumer of the events api sees the event data.
type EventResponseJSONAPI struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Relative Resource Name for the operation event
	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	// relative resource name for associated asset ( asset the operation is performed on  - has to have specific behaviour enabled)
	AssetIdentity string `protobuf:"bytes,2,opt,name=asset_identity,json=assetIdentity,proto3" json:"asset_identity,omitempty"`
	// map of event attributes. Specific behaviours define required and optional event attributes for each supported operation.
	EventAttributes map[string]string `protobuf:"bytes,16,rep,name=event_attributes,json=eventAttributes,proto3" json:"event_attributes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// map of asset attributes. Specific behaviours define required and optional asset attributes. These attributes cause the corresponding attributes on the asset to be updated.
	AssetAttributes map[string]string `protobuf:"bytes,17,rep,name=asset_attributes,json=assetAttributes,proto3" json:"asset_attributes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
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
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *EventResponseJSONAPI) Reset() {
	*x = EventResponseJSONAPI{}
	mi := &file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventResponseJSONAPI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventResponseJSONAPI) ProtoMessage() {}

func (x *EventResponseJSONAPI) ProtoReflect() protoreflect.Message {
	mi := &file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_msgTypes[0]
	if x != nil {
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

const file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_rawDesc = "" +
	"\n" +
	"Adatatrails-common-api/assets/v2/assets/eventresponsejsonapi.proto\x12\farchivist.v2\x1a.protoc-gen-openapiv2/options/annotations.proto\x1a6datatrails-common-api/assets/v2/assets/principal.proto\x1a;datatrails-common-api/assets/v2/assets/merklelogentry.proto\"\xbb\x17\n" +
	"\x14EventResponseJSONAPI\x12?\n" +
	"\bidentity\x18\x01 \x01(\tB#\x92A 2\x1cidentity of a event resource@\x01R\bidentity\x12\x80\x01\n" +
	"\x0easset_identity\x18\x02 \x01(\tBY\x92AV2Ridentity of a related asset resource `assets/11bf5b37-e0b8-42e0-8dcf-dc8c4aefc000`@\x01R\rassetIdentity\x12\x8e\x01\n" +
	"\x10event_attributes\x18\x10 \x03(\v27.archivist.v2.EventResponseJSONAPI.EventAttributesEntryB*\x92A'2%key value mapping of event attributesR\x0feventAttributes\x12\x8e\x01\n" +
	"\x10asset_attributes\x18\x11 \x03(\v27.archivist.v2.EventResponseJSONAPI.AssetAttributesEntryB*\x92A'2%key value mapping of asset attributesR\x0fassetAttributes\x12X\n" +
	"\toperation\x18\x04 \x01(\tB:\x92A720The operation represented by the event. `Record`@\x01x\x80 R\toperation\x12\\\n" +
	"\tbehaviour\x18\x0e \x01(\tB>\x92A;24The behaviour used to create event. `RecordEvidence`@\x01x\x80 R\tbehaviour\x12d\n" +
	"\x12timestamp_declared\x18\x05 \x01(\tB5\x92A22.RFC 3339 time of event as declared by the user@\x01R\x11timestampDeclared\x12f\n" +
	"\x12timestamp_accepted\x18\x06 \x01(\tB7\x92A420RFC 3339 time of event as recorded by the server@\x01R\x11timestampAccepted\x12p\n" +
	"\x13timestamp_committed\x18\a \x01(\tB?\x92A<28RFC 3339 time of event as recorded in verifiable storage@\x01R\x12timestampCommitted\x12q\n" +
	"\x12principal_declared\x18\b \x01(\v2\x17.archivist.v2.PrincipalB)\x92A&2\x1eprincipal provided by the user@\x01\x9a\x02\x01\x06R\x11principalDeclared\x12s\n" +
	"\x12principal_accepted\x18\t \x01(\v2\x17.archivist.v2.PrincipalB+\x92A(2 principal recorded by the server@\x01\x9a\x02\x01\x06R\x11principalAccepted\x12\x83\x01\n" +
	"\x13confirmation_status\x18\n" +
	" \x01(\tBR\x92AO2Gindicates if the event has been succesfully committed to the blockchain@\x01\x9a\x02\x01\aR\x12confirmationStatus\x12{\n" +
	"\x0etransaction_id\x18\v \x01(\tBT\x92AQ2Lhash of the transaction as a hex string `0x11bf5b37e0b842e08dcfdc8c4aefc000`x\x80 R\rtransactionId\x12O\n" +
	"\fblock_number\x18\f \x01(\x04B,\x92A)2%number of block event was commited on@\x01R\vblockNumber\x12X\n" +
	"\x11transaction_index\x18\r \x01(\x04B+\x92A(2$index of event within commited block@\x01R\x10transactionIndex\x12G\n" +
	"\x04from\x18\x0f \x01(\tB3\x92A02,wallet address for the creator of this event@\x01R\x04from\x12c\n" +
	"\x0ftenant_identity\x18\x12 \x01(\tB:\x92A722Identity of the tenant the that created this eventx\x80\bR\x0etenantIdentity\x12v\n" +
	"\x0fmerklelog_entry\x18\x13 \x01(\v2\x1c.archivist.v2.MerkleLogEntryB/\x92A,2'verifiable merkle mmr log entry detailsx\x80\bR\x0emerklelogEntry\x1aB\n" +
	"\x14EventAttributesEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\x1aB\n" +
	"\x14AssetAttributesEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01:\x80\a\x92A\xfc\x06\n" +
	"\x1a2\x18This describes an Event.2\xdd\x06{ \"identity\": \"assets/add30235-1424-4fda-840a-d5ef82c4c96f/events/11bf5b37-e0b8-42e0-8dcf-dc8c4aefc000\", \"asset_identity\": \"assets/add30235-1424-4fda-840a-d5ef82c4c96f\", \"operation\": \"Record\", \"behaviour\": \"RecordEvidence\", \"event_attributes\": {    \"arc_display_type\": \"Paint\",    \"arc_description\": \"Painted the fence\" }, \"asset_attributes\": {    \"colour\": \"Midnight Blue\" }, \"timestamp_accepted\": \"2019-11-27T14:44:19Z\", \"timestamp_declared\": \"2019-11-27T14:44:19Z\", \"timestamp_committed\": \"2019-11-27T14:44:19Z\", \"principal_declared\": {     \"issuer\": \"job.idp.server/1234\", \"subject\":\"bob@job\"  },  \"principal_accepted\": {     \"issuer\": \"job.idp.server/1234\", \"subject\":\"bob@job\" }, \"confirmation_status\": \"CONFIRMED\", \"block_number\": 12, \"transaction_index\": 5, \"transaction_id\": \"0x07569\", \"tenant_identity\": \"tenant/8e0b600c-8234-43e4-860c-e95bdcd695a9\" }BLZJgithub.com/datatrails/go-datatrails-common-api-gen/assets/v2/assets;assetsb\x06proto3"

var (
	file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_rawDescOnce sync.Once
	file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_rawDescData []byte
)

func file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_rawDescGZIP() []byte {
	file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_rawDescOnce.Do(func() {
		file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_rawDesc), len(file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_rawDesc)))
	})
	return file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_rawDescData
}

var file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_goTypes = []any{
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
	file_datatrails_common_api_assets_v2_assets_principal_proto_init()
	file_datatrails_common_api_assets_v2_assets_merklelogentry_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_rawDesc), len(file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_rawDesc)),
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
	file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_goTypes = nil
	file_datatrails_common_api_assets_v2_assets_eventresponsejsonapi_proto_depIdxs = nil
}
