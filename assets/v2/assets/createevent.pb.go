// Maintainers, please refer to the style guide here:
//     https://developers.google.com/protocol-buffers/docs/style

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.24.3
// source: datatrails-common-api/assets/v2/assets/createevent.proto

package assets

import (
	attribute "github.com/datatrails/go-datatrails-common-api-gen/attribute/v2/attribute"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type CreateEventRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// relative resource name for associated asset ( asset the operation is performed on  - has to have specific behaviour enabled)
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// name of this behaviour
	Behaviour string `protobuf:"bytes,2,opt,name=behaviour,proto3" json:"behaviour,omitempty"`
	// name of operation on this behviour
	Operation string `protobuf:"bytes,3,opt,name=operation,proto3" json:"operation,omitempty"`
	// principal information associated with event - if not provided will be set to principal_accepted
	PrincipalDeclared *Principal `protobuf:"bytes,5,opt,name=principal_declared,json=principalDeclared,proto3" json:"principal_declared,omitempty"`
	// timestamp when operation was actually performed - if not provided will be set to timestamp_accepted
	TimestampDeclared *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=timestamp_declared,json=timestampDeclared,proto3" json:"timestamp_declared,omitempty"`
	// map of event attributes. Specific behaviours define required and optional event attributes for each supported operation.
	EventAttributes map[string]*attribute.Attribute `protobuf:"bytes,7,rep,name=event_attributes,json=eventAttributes,proto3" json:"event_attributes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// map of asset attributes. Specific behaviours define required and optional asset attributes. These attributes cause the corresponding attributes on the asset to be updated.
	AssetAttributes map[string]*attribute.Attribute `protobuf:"bytes,8,rep,name=asset_attributes,json=assetAttributes,proto3" json:"asset_attributes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *CreateEventRequest) Reset() {
	*x = CreateEventRequest{}
	mi := &file_datatrails_common_api_assets_v2_assets_createevent_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventRequest) ProtoMessage() {}

func (x *CreateEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_datatrails_common_api_assets_v2_assets_createevent_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventRequest.ProtoReflect.Descriptor instead.
func (*CreateEventRequest) Descriptor() ([]byte, []int) {
	return file_datatrails_common_api_assets_v2_assets_createevent_proto_rawDescGZIP(), []int{0}
}

func (x *CreateEventRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *CreateEventRequest) GetBehaviour() string {
	if x != nil {
		return x.Behaviour
	}
	return ""
}

func (x *CreateEventRequest) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *CreateEventRequest) GetPrincipalDeclared() *Principal {
	if x != nil {
		return x.PrincipalDeclared
	}
	return nil
}

func (x *CreateEventRequest) GetTimestampDeclared() *timestamppb.Timestamp {
	if x != nil {
		return x.TimestampDeclared
	}
	return nil
}

func (x *CreateEventRequest) GetEventAttributes() map[string]*attribute.Attribute {
	if x != nil {
		return x.EventAttributes
	}
	return nil
}

func (x *CreateEventRequest) GetAssetAttributes() map[string]*attribute.Attribute {
	if x != nil {
		return x.AssetAttributes
	}
	return nil
}

var File_datatrails_common_api_assets_v2_assets_createevent_proto protoreflect.FileDescriptor

const file_datatrails_common_api_assets_v2_assets_createevent_proto_rawDesc = "" +
	"\n" +
	"8datatrails-common-api/assets/v2/assets/createevent.proto\x12\farchivist.v2\x1a\x17validate/validate.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a6datatrails-common-api/assets/v2/assets/principal.proto\x1a<datatrails-common-api/attribute/v2/attribute/attribute.proto\"\x8b\f\n" +
	"\x12CreateEventRequest\x12\xd0\x01\n" +
	"\x04uuid\x18\x01 \x01(\tB\xbb\x01\x92A\xaf\x012\xaa\x01Specify the Asset UUID where `assets/{uuid}` is the Asset Identity e.g. `add30235-1424-4fda-840a-d5ef82c4c96f` from Identity `assets/add30235-1424-4fda-840a-d5ef82c4c96f`@\x01\xfaB\x05r\x03\xb0\x01\x01R\x04uuid\x12&\n" +
	"\tbehaviour\x18\x02 \x01(\tB\b\xfaB\x05r\x03\x18\x80\bR\tbehaviour\x12&\n" +
	"\toperation\x18\x03 \x01(\tB\b\xfaB\x05r\x03\x18\x80\bR\toperation\x12F\n" +
	"\x12principal_declared\x18\x05 \x01(\v2\x17.archivist.v2.PrincipalR\x11principalDeclared\x12w\n" +
	"\x12timestamp_declared\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampB,\x92A)2%time of event as declared by the user@\x01R\x11timestampDeclared\x12\xfa\x01\n" +
	"\x10event_attributes\x18\a \x03(\v25.archivist.v2.CreateEventRequest.EventAttributesEntryB\x97\x01\x92A'2%key value mapping of event attributes\xfaBj\x9a\x01g\"erc\x10\x01\x18\x80\b2X^[^[:cntrl:]]*?[^[[:cntrl:]]+?[^[:cntrl:]]$|^[^[:cntrl:]]$|^[^[:cntrl:]]*?[^][:cntrl:]]$\xba\x01\x01.R\x0feventAttributes\x12\xfa\x01\n" +
	"\x10asset_attributes\x18\b \x03(\v25.archivist.v2.CreateEventRequest.AssetAttributesEntryB\x97\x01\x92A'2%key value mapping of asset attributes\xfaBj\x9a\x01g\"erc\x10\x01\x18\x80\b2X^[^[:cntrl:]]*?[^[[:cntrl:]]+?[^[:cntrl:]]$|^[^[:cntrl:]]$|^[^[:cntrl:]]*?[^][:cntrl:]]$\xba\x01\x01.R\x0fassetAttributes\x1a[\n" +
	"\x14EventAttributesEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12-\n" +
	"\x05value\x18\x02 \x01(\v2\x17.archivist.v2.AttributeR\x05value:\x028\x01\x1a[\n" +
	"\x14AssetAttributesEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12-\n" +
	"\x05value\x18\x02 \x01(\v2\x17.archivist.v2.AttributeR\x05value:\x028\x01:\xdc\x02\x92A\xd8\x02\n" +
	"\x82\x012hRequest creating RecordEvidence event Specify `operation` to chose which RecordEvidence event to create \xd2\x01\tbehaviour\xd2\x01\toperation2\xd0\x01{ \"operation\": \"Record\", \"behaviour\": \"RecordEvidence\", \"event_attributes\": {    \"arc_display_type\": \"Paint\",    \"arc_description\": \"Painted the fence\"}, \"asset_attributes\": {    \"colour\": \"Midnight Blue\" } }BLZJgithub.com/datatrails/go-datatrails-common-api-gen/assets/v2/assets;assetsb\x06proto3"

var (
	file_datatrails_common_api_assets_v2_assets_createevent_proto_rawDescOnce sync.Once
	file_datatrails_common_api_assets_v2_assets_createevent_proto_rawDescData []byte
)

func file_datatrails_common_api_assets_v2_assets_createevent_proto_rawDescGZIP() []byte {
	file_datatrails_common_api_assets_v2_assets_createevent_proto_rawDescOnce.Do(func() {
		file_datatrails_common_api_assets_v2_assets_createevent_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_datatrails_common_api_assets_v2_assets_createevent_proto_rawDesc), len(file_datatrails_common_api_assets_v2_assets_createevent_proto_rawDesc)))
	})
	return file_datatrails_common_api_assets_v2_assets_createevent_proto_rawDescData
}

var file_datatrails_common_api_assets_v2_assets_createevent_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_datatrails_common_api_assets_v2_assets_createevent_proto_goTypes = []any{
	(*CreateEventRequest)(nil),    // 0: archivist.v2.CreateEventRequest
	nil,                           // 1: archivist.v2.CreateEventRequest.EventAttributesEntry
	nil,                           // 2: archivist.v2.CreateEventRequest.AssetAttributesEntry
	(*Principal)(nil),             // 3: archivist.v2.Principal
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*attribute.Attribute)(nil),   // 5: archivist.v2.Attribute
}
var file_datatrails_common_api_assets_v2_assets_createevent_proto_depIdxs = []int32{
	3, // 0: archivist.v2.CreateEventRequest.principal_declared:type_name -> archivist.v2.Principal
	4, // 1: archivist.v2.CreateEventRequest.timestamp_declared:type_name -> google.protobuf.Timestamp
	1, // 2: archivist.v2.CreateEventRequest.event_attributes:type_name -> archivist.v2.CreateEventRequest.EventAttributesEntry
	2, // 3: archivist.v2.CreateEventRequest.asset_attributes:type_name -> archivist.v2.CreateEventRequest.AssetAttributesEntry
	5, // 4: archivist.v2.CreateEventRequest.EventAttributesEntry.value:type_name -> archivist.v2.Attribute
	5, // 5: archivist.v2.CreateEventRequest.AssetAttributesEntry.value:type_name -> archivist.v2.Attribute
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_datatrails_common_api_assets_v2_assets_createevent_proto_init() }
func file_datatrails_common_api_assets_v2_assets_createevent_proto_init() {
	if File_datatrails_common_api_assets_v2_assets_createevent_proto != nil {
		return
	}
	file_datatrails_common_api_assets_v2_assets_principal_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_datatrails_common_api_assets_v2_assets_createevent_proto_rawDesc), len(file_datatrails_common_api_assets_v2_assets_createevent_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datatrails_common_api_assets_v2_assets_createevent_proto_goTypes,
		DependencyIndexes: file_datatrails_common_api_assets_v2_assets_createevent_proto_depIdxs,
		MessageInfos:      file_datatrails_common_api_assets_v2_assets_createevent_proto_msgTypes,
	}.Build()
	File_datatrails_common_api_assets_v2_assets_createevent_proto = out.File
	file_datatrails_common_api_assets_v2_assets_createevent_proto_goTypes = nil
	file_datatrails_common_api_assets_v2_assets_createevent_proto_depIdxs = nil
}
