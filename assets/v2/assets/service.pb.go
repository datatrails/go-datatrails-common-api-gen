// Maintainers, please refer to the style guide here:
//     https://developers.google.com/protocol-buffers/docs/style

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.24.3
// source: datatrails-common-api/assets/v2/assets/service.proto

package assets

import (
	caps "github.com/datatrails/go-datatrails-common-api-gen/caps/v1/caps"
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

var File_datatrails_common_api_assets_v2_assets_service_proto protoreflect.FileDescriptor

const file_datatrails_common_api_assets_v2_assets_service_proto_rawDesc = "" +
	"\n" +
	"4datatrails-common-api/assets/v2/assets/service.proto\x12\farchivist.v2\x1a\x1cgoogle/api/annotations.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\x1a-datatrails-common-api/caps/v1/caps/caps.proto\x1a8datatrails-common-api/assets/v2/assets/createasset.proto\x1a:datatrails-common-api/assets/v2/assets/assetresponse.proto\x1a5datatrails-common-api/assets/v2/assets/getasset.proto\x1a7datatrails-common-api/assets/v2/assets/listassets.proto\x1a8datatrails-common-api/assets/v2/assets/createevent.proto\x1a:datatrails-common-api/assets/v2/assets/eventresponse.proto\x1a5datatrails-common-api/assets/v2/assets/getevent.proto\x1a7datatrails-common-api/assets/v2/assets/listevents.proto\x1a9datatrails-common-api/assets/v2/assets/miscmessages.proto2\xcb$\n" +
	"\x06Assets\x12\x83\a\n" +
	"\n" +
	"ListAssets\x12\x1f.archivist.v2.ListAssetsRequest\x1a .archivist.v2.ListAssetsResponse\"\xb1\x06\x92A\x91\x06\n" +
	"\x11Assets and Events\x12\vList Assets\x1a\x1aRetrieves a list of AssetsJ\xf5\x03\n" +
	"\x03206\x12\xed\x03\n" +
	"\xea\x03The number of assets exceeds the servers limit. The approximate number of matching results is provided by the x-total-count header if the 'x-request-total-count' header on the request is set to 'true'. The exact limit is available in the content-range header. The value format is 'items 0-LIMIT/TOTAL'. Note that x-total-count is always present for 200 and 206 responses. It is the servers best available approximation. Similarly, in any result set, you may get a few more than LIMIT items.JC\n" +
	"\x03401\x12<\n" +
	":Returned when the user is not authenticated to the system.JA\n" +
	"\x03403\x12:\n" +
	"8Returned when the user is not authorized to list Assets.JS\n" +
	"\x03429\x12L\n" +
	"JReturned when a user exceeds their subscription's rate limit for requests.\x82\xd3\xe4\x93\x02\x16\x12\x14/archivist/v2/assets\x12\xd4\x03\n" +
	"\vCreateAsset\x12 .archivist.v2.CreateAssetRequest\x1a\x1b.archivist.v2.AssetResponse\"\x85\x03\x92A\xe2\x02\n" +
	"\x11Assets and Events\x12\x0fCreate an Asset\x1a\x10Creates an AssetJC\n" +
	"\x03401\x12<\n" +
	":Returned when the user is not authenticated to the system.JI\n" +
	"\x03402\x12B\n" +
	"@Returned when the number of assets would exceed the user's quotaJE\n" +
	"\x03403\x12>\n" +
	"<Returned when the user is not authorized to create an Asset.JS\n" +
	"\x03429\x12L\n" +
	"JReturned when a user exceeds their subscription's rate limit for requests.\x82\xd3\xe4\x93\x02\x19:\x01*\"\x14/archivist/v2/assets\x12\xd8\x03\n" +
	"\bGetAsset\x12\x1d.archivist.v2.GetAssetRequest\x1a\x1b.archivist.v2.AssetResponse\"\x8f\x03\x92A\xe8\x02\n" +
	"\x11Assets and Events\x12\x1aRetrieves a specific Asset\x1a\x1aRetrieves a specific AssetJC\n" +
	"\x03401\x12<\n" +
	":Returned when the user is not authenticated to the system.JC\n" +
	"\x03403\x12<\n" +
	":Returned when the user is not authorized to view an Asset.J<\n" +
	"\x03404\x125\n" +
	"3Returned when the asset with the id does not exist.JS\n" +
	"\x03429\x12L\n" +
	"JReturned when a user exceeds their subscription's rate limit for requests.\x82\xd3\xe4\x93\x02\x1d\x12\x1b/archivist/v2/assets/{uuid}\x12\xa5\x04\n" +
	"\x11GetAssetPublicURL\x12&.archivist.v2.GetAssetPublicURLRequest\x1a$.archivist.v2.PublicAssetURLResponse\"\xc1\x03\x92A\x90\x03\n" +
	"\x11Assets and Events\x12.Retrieves the public url for a specific Asset.\x1a.Retrieves the public url for a specific Asset.JC\n" +
	"\x03401\x12<\n" +
	":Returned when the user is not authenticated to the system.JC\n" +
	"\x03403\x12<\n" +
	":Returned when the user is not authorized to view an Asset.J<\n" +
	"\x03404\x125\n" +
	"3Returned when the asset with the id does not exist.JS\n" +
	"\x03429\x12L\n" +
	"JReturned when a user exceeds their subscription's rate limit for requests.\x82\xd3\xe4\x93\x02'\x12%/archivist/v2/assets/{uuid}:publicurl\x12\xd2\x03\n" +
	"\bGetEvent\x12\x1d.archivist.v2.GetEventRequest\x1a\x1b.archivist.v2.EventResponse\"\x89\x03\x92A\xce\x02\n" +
	"\x11Assets and Events\x12\x0fRetrieves Event\x1a\x1aRetrieves a specific EventJC\n" +
	"\x03401\x12<\n" +
	":Returned when the user is not authenticated to the system.J@\n" +
	"\x03403\x129\n" +
	"7Returned when the user is not authorized to view Event.J0\n" +
	"\x03404\x12)\n" +
	"'Returned when the event does not exist.JS\n" +
	"\x03429\x12L\n" +
	"JReturned when a user exceeds their subscription's rate limit for requests.\x82\xd3\xe4\x93\x021\x12//archivist/v2/assets/{asset_uuid}/events/{uuid}\x12\xb0\x04\n" +
	"\x11GetEventPublicURL\x12\x1d.archivist.v2.GetEventRequest\x1a$.archivist.v2.PublicEventURLResponse\"\xd5\x03\x92A\x90\x03\n" +
	"\x11Assets and Events\x12.Retrieves the public url for a specific Event.\x1a.Retrieves the public url for a specific Event.JC\n" +
	"\x03401\x12<\n" +
	":Returned when the user is not authenticated to the system.JC\n" +
	"\x03403\x12<\n" +
	":Returned when the user is not authorized to view an Asset.J<\n" +
	"\x03404\x125\n" +
	"3Returned when the asset with the id does not exist.JS\n" +
	"\x03429\x12L\n" +
	"JReturned when a user exceeds their subscription's rate limit for requests.\x82\xd3\xe4\x93\x02;\x129/archivist/v2/assets/{asset_uuid}/events/{uuid}:publicurl\x12\xbe\x06\n" +
	"\n" +
	"ListEvents\x12\x1f.archivist.v2.ListEventsRequest\x1a .archivist.v2.ListEventsResponse\"\xec\x05\x92A\xbe\x05\n" +
	"\x11Assets and Events\x12\vList Events\x1a\fLists EventsJ\xb0\x03\n" +
	"\x03206\x12\xa8\x03\n" +
	"\xa5\x03The number of events exceeds the servers limit. The approximate number of matching results is provided by the x-total-count header, the exact limit is available in the content-range header. The value format is 'items 0-LIMIT/TOTAL'.  Note that x-total-count is always present for 200 and 206 responses. It is the servers best available approximation. Similarly, in any result set, you may get a few more than LIMIT items.JC\n" +
	"\x03401\x12<\n" +
	":Returned when the user is not authenticated to the system.JA\n" +
	"\x03403\x12:\n" +
	"8Returned when the user is not authorized to list Events.JS\n" +
	"\x03429\x12L\n" +
	"JReturned when a user exceeds their subscription's rate limit for requests.\x82\xd3\xe4\x93\x02$\x12\"/archivist/v2/assets/{uuid}/events\x12\x97\x03\n" +
	"\aGetCaps\x12\x1c.archivist.v1.GetCapsRequest\x1a\x12.archivist.v1.Caps\"\xd9\x02\x92A\xb4\x02\n" +
	"\x11Assets and Events\n" +
	"\vUnsupported\x12)Get remaining capped resources for Assets\x1aMNot stable or officially supported. Get remaining capped resources for AssetsJC\n" +
	"\x03401\x12<\n" +
	":Returned when the user is not authenticated to the system.JS\n" +
	"\x03429\x12L\n" +
	"JReturned when a user exceeds their subscription's rate limit for requests.\x82\xd3\xe4\x93\x02\x1b\x12\x19/archivist/v2/assets:caps2\x9c\x03\n" +
	"\x06Events\x12\x91\x03\n" +
	"\x06Create\x12 .archivist.v2.CreateEventRequest\x1a\x1b.archivist.v2.EventResponse\"\xc7\x02\x92A\x96\x02\n" +
	"\x11Assets and Events\x12\x10Creates an Event\x1a\x10Creates an EventJC\n" +
	"\x03401\x12<\n" +
	":Returned when the user is not authenticated to the system.JC\n" +
	"\x03402\x12<\n" +
	":Returned when the user's quota of Events has been reached.JS\n" +
	"\x03429\x12L\n" +
	"JReturned when a user exceeds their subscription's rate limit for requests.\x82\xd3\xe4\x93\x02':\x01*\"\"/archivist/v2/assets/{uuid}/eventsB\xa5\x02\x92A\xd5\x01\x12_\n" +
	"\n" +
	"Assets API\x12#API for asset and event management.\"'\n" +
	"\n" +
	"DataTrails\x12\x19https://www.datatrails.ai2\x032.0\"\x14/archivist/v2/assets*\x01\x022\x10application/json:\x10application/jsonj5\n" +
	"\x11Assets and Events\x12 Primary API for Asset managementZJgithub.com/datatrails/go-datatrails-common-api-gen/assets/v2/assets;assetsb\x06proto3"

var file_datatrails_common_api_assets_v2_assets_service_proto_goTypes = []any{
	(*ListAssetsRequest)(nil),        // 0: archivist.v2.ListAssetsRequest
	(*CreateAssetRequest)(nil),       // 1: archivist.v2.CreateAssetRequest
	(*GetAssetRequest)(nil),          // 2: archivist.v2.GetAssetRequest
	(*GetAssetPublicURLRequest)(nil), // 3: archivist.v2.GetAssetPublicURLRequest
	(*GetEventRequest)(nil),          // 4: archivist.v2.GetEventRequest
	(*ListEventsRequest)(nil),        // 5: archivist.v2.ListEventsRequest
	(*caps.GetCapsRequest)(nil),      // 6: archivist.v1.GetCapsRequest
	(*CreateEventRequest)(nil),       // 7: archivist.v2.CreateEventRequest
	(*ListAssetsResponse)(nil),       // 8: archivist.v2.ListAssetsResponse
	(*AssetResponse)(nil),            // 9: archivist.v2.AssetResponse
	(*PublicAssetURLResponse)(nil),   // 10: archivist.v2.PublicAssetURLResponse
	(*EventResponse)(nil),            // 11: archivist.v2.EventResponse
	(*PublicEventURLResponse)(nil),   // 12: archivist.v2.PublicEventURLResponse
	(*ListEventsResponse)(nil),       // 13: archivist.v2.ListEventsResponse
	(*caps.Caps)(nil),                // 14: archivist.v1.Caps
}
var file_datatrails_common_api_assets_v2_assets_service_proto_depIdxs = []int32{
	0,  // 0: archivist.v2.Assets.ListAssets:input_type -> archivist.v2.ListAssetsRequest
	1,  // 1: archivist.v2.Assets.CreateAsset:input_type -> archivist.v2.CreateAssetRequest
	2,  // 2: archivist.v2.Assets.GetAsset:input_type -> archivist.v2.GetAssetRequest
	3,  // 3: archivist.v2.Assets.GetAssetPublicURL:input_type -> archivist.v2.GetAssetPublicURLRequest
	4,  // 4: archivist.v2.Assets.GetEvent:input_type -> archivist.v2.GetEventRequest
	4,  // 5: archivist.v2.Assets.GetEventPublicURL:input_type -> archivist.v2.GetEventRequest
	5,  // 6: archivist.v2.Assets.ListEvents:input_type -> archivist.v2.ListEventsRequest
	6,  // 7: archivist.v2.Assets.GetCaps:input_type -> archivist.v1.GetCapsRequest
	7,  // 8: archivist.v2.Events.Create:input_type -> archivist.v2.CreateEventRequest
	8,  // 9: archivist.v2.Assets.ListAssets:output_type -> archivist.v2.ListAssetsResponse
	9,  // 10: archivist.v2.Assets.CreateAsset:output_type -> archivist.v2.AssetResponse
	9,  // 11: archivist.v2.Assets.GetAsset:output_type -> archivist.v2.AssetResponse
	10, // 12: archivist.v2.Assets.GetAssetPublicURL:output_type -> archivist.v2.PublicAssetURLResponse
	11, // 13: archivist.v2.Assets.GetEvent:output_type -> archivist.v2.EventResponse
	12, // 14: archivist.v2.Assets.GetEventPublicURL:output_type -> archivist.v2.PublicEventURLResponse
	13, // 15: archivist.v2.Assets.ListEvents:output_type -> archivist.v2.ListEventsResponse
	14, // 16: archivist.v2.Assets.GetCaps:output_type -> archivist.v1.Caps
	11, // 17: archivist.v2.Events.Create:output_type -> archivist.v2.EventResponse
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_datatrails_common_api_assets_v2_assets_service_proto_init() }
func file_datatrails_common_api_assets_v2_assets_service_proto_init() {
	if File_datatrails_common_api_assets_v2_assets_service_proto != nil {
		return
	}
	file_datatrails_common_api_assets_v2_assets_createasset_proto_init()
	file_datatrails_common_api_assets_v2_assets_assetresponse_proto_init()
	file_datatrails_common_api_assets_v2_assets_getasset_proto_init()
	file_datatrails_common_api_assets_v2_assets_listassets_proto_init()
	file_datatrails_common_api_assets_v2_assets_createevent_proto_init()
	file_datatrails_common_api_assets_v2_assets_eventresponse_proto_init()
	file_datatrails_common_api_assets_v2_assets_getevent_proto_init()
	file_datatrails_common_api_assets_v2_assets_listevents_proto_init()
	file_datatrails_common_api_assets_v2_assets_miscmessages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_datatrails_common_api_assets_v2_assets_service_proto_rawDesc), len(file_datatrails_common_api_assets_v2_assets_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_datatrails_common_api_assets_v2_assets_service_proto_goTypes,
		DependencyIndexes: file_datatrails_common_api_assets_v2_assets_service_proto_depIdxs,
	}.Build()
	File_datatrails_common_api_assets_v2_assets_service_proto = out.File
	file_datatrails_common_api_assets_v2_assets_service_proto_goTypes = nil
	file_datatrails_common_api_assets_v2_assets_service_proto_depIdxs = nil
}
