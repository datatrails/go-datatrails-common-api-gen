// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.24.3
// source: datatrails-common-api/assets/v2/assets/publicassets.proto

// Contains the generated protocol definitions for the registration protocol.
// Registrar clients, service proxies and the archivist service will all use
// these.
// Maintainers, please refer to the style guide here:
//     https://developers.google.com/protocol-buffers/docs/style

package assets

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_datatrails_common_api_assets_v2_assets_publicassets_proto protoreflect.FileDescriptor

const file_datatrails_common_api_assets_v2_assets_publicassets_proto_rawDesc = "" +
	"\n" +
	"9datatrails-common-api/assets/v2/assets/publicassets.proto\x12\farchivist.v2\x1a\x1cgoogle/api/annotations.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\x1a:datatrails-common-api/assets/v2/assets/assetresponse.proto\x1a5datatrails-common-api/assets/v2/assets/getasset.proto\x1a:datatrails-common-api/assets/v2/assets/eventresponse.proto\x1a5datatrails-common-api/assets/v2/assets/getevent.proto\x1a7datatrails-common-api/assets/v2/assets/listevents.proto2\xf0\t\n" +
	"\fPublicAssets\x12\xa7\x02\n" +
	"\x0eGetPublicEvent\x12\x1d.archivist.v2.GetEventRequest\x1a\x1b.archivist.v2.EventResponse\"\xd8\x01\x92A\x97\x01\x12*Retrieves Archivist event for public asset\x1a7Retrieves a specific Archivist event for a public assetJ0\n" +
	"\x03404\x12)\n" +
	"'Returned when the event does not exist.\x82\xd3\xe4\x93\x027\x125/archivist/v2/publicassets/{asset_uuid}/events/{uuid}\x12\x9e\x02\n" +
	"\x0eGetPublicAsset\x12\x1d.archivist.v2.GetAssetRequest\x1a\x1b.archivist.v2.AssetResponse\"\xcf\x01\x92A\xa2\x01\x12+Retrieves a specific public Archivist asset\x1a+Retrieves a specific public Archivist assetJF\n" +
	"\x03404\x12?\n" +
	"=Returned when the no asset with the provided id can be found.\x82\xd3\xe4\x93\x02#\x12!/archivist/v2/publicassets/{uuid}\x12\x94\x05\n" +
	"\x10ListPublicEvents\x12\x1f.archivist.v2.ListEventsRequest\x1a .archivist.v2.ListEventsResponse\"\xbc\x04\x92A\x88\x04\x12(List Archivist events for a public asset\x1a)Lists Archivist events for a public assetJ\xb0\x03\n" +
	"\x03206\x12\xa8\x03\n" +
	"\xa5\x03The number of events exceeds the servers limit. The approximate number of matching results is provided by the x-total-count header, the exact limit is available in the content-range header. The value format is 'items 0-LIMIT/TOTAL'.  Note that x-total-count is always present for 200 and 206 responses. It is the servers best available approximation. Similarly, in any result set, you may get a few more than LIMIT items.\x82\xd3\xe4\x93\x02*\x12(/archivist/v2/publicassets/{uuid}/eventsB\xd2\x02\x92A\x82\x02\x12l\n" +
	"\x10PublicAssets API\x12*API for public asset and event management.\"'\n" +
	"\n" +
	"DataTrails\x12\x19https://www.datatrails.ai2\x032.0\"\x1a/archivist/v2/publicassets*\x01\x022\x10application/json:\x10application/jsonjO\n" +
	"\fPublicAssets\x12?Retrieve Public Attestations and public Asset and Event recordsZJgithub.com/datatrails/go-datatrails-common-api-gen/assets/v2/assets;assetsb\x06proto3"

var file_datatrails_common_api_assets_v2_assets_publicassets_proto_goTypes = []any{
	(*GetEventRequest)(nil),    // 0: archivist.v2.GetEventRequest
	(*GetAssetRequest)(nil),    // 1: archivist.v2.GetAssetRequest
	(*ListEventsRequest)(nil),  // 2: archivist.v2.ListEventsRequest
	(*EventResponse)(nil),      // 3: archivist.v2.EventResponse
	(*AssetResponse)(nil),      // 4: archivist.v2.AssetResponse
	(*ListEventsResponse)(nil), // 5: archivist.v2.ListEventsResponse
}
var file_datatrails_common_api_assets_v2_assets_publicassets_proto_depIdxs = []int32{
	0, // 0: archivist.v2.PublicAssets.GetPublicEvent:input_type -> archivist.v2.GetEventRequest
	1, // 1: archivist.v2.PublicAssets.GetPublicAsset:input_type -> archivist.v2.GetAssetRequest
	2, // 2: archivist.v2.PublicAssets.ListPublicEvents:input_type -> archivist.v2.ListEventsRequest
	3, // 3: archivist.v2.PublicAssets.GetPublicEvent:output_type -> archivist.v2.EventResponse
	4, // 4: archivist.v2.PublicAssets.GetPublicAsset:output_type -> archivist.v2.AssetResponse
	5, // 5: archivist.v2.PublicAssets.ListPublicEvents:output_type -> archivist.v2.ListEventsResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_datatrails_common_api_assets_v2_assets_publicassets_proto_init() }
func file_datatrails_common_api_assets_v2_assets_publicassets_proto_init() {
	if File_datatrails_common_api_assets_v2_assets_publicassets_proto != nil {
		return
	}
	file_datatrails_common_api_assets_v2_assets_assetresponse_proto_init()
	file_datatrails_common_api_assets_v2_assets_getasset_proto_init()
	file_datatrails_common_api_assets_v2_assets_eventresponse_proto_init()
	file_datatrails_common_api_assets_v2_assets_getevent_proto_init()
	file_datatrails_common_api_assets_v2_assets_listevents_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_datatrails_common_api_assets_v2_assets_publicassets_proto_rawDesc), len(file_datatrails_common_api_assets_v2_assets_publicassets_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_datatrails_common_api_assets_v2_assets_publicassets_proto_goTypes,
		DependencyIndexes: file_datatrails_common_api_assets_v2_assets_publicassets_proto_depIdxs,
	}.Build()
	File_datatrails_common_api_assets_v2_assets_publicassets_proto = out.File
	file_datatrails_common_api_assets_v2_assets_publicassets_proto_goTypes = nil
	file_datatrails_common_api_assets_v2_assets_publicassets_proto_depIdxs = nil
}
