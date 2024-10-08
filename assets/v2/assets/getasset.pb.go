// Maintainers, please refer to the style guide here:
//     https://developers.google.com/protocol-buffers/docs/style

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.24.3
// source: datatrails-common-api/assets/v2/assets/getasset.proto

package assets

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type GetAssetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// asset identity
	Uuid   string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	AtTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=at_time,json=atTime,proto3" json:"at_time,omitempty"`
}

func (x *GetAssetRequest) Reset() {
	*x = GetAssetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datatrails_common_api_assets_v2_assets_getasset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssetRequest) ProtoMessage() {}

func (x *GetAssetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_datatrails_common_api_assets_v2_assets_getasset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssetRequest.ProtoReflect.Descriptor instead.
func (*GetAssetRequest) Descriptor() ([]byte, []int) {
	return file_datatrails_common_api_assets_v2_assets_getasset_proto_rawDescGZIP(), []int{0}
}

func (x *GetAssetRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *GetAssetRequest) GetAtTime() *timestamppb.Timestamp {
	if x != nil {
		return x.AtTime
	}
	return nil
}

type GetAssetPublicURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// asset identity
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GetAssetPublicURLRequest) Reset() {
	*x = GetAssetPublicURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datatrails_common_api_assets_v2_assets_getasset_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssetPublicURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssetPublicURLRequest) ProtoMessage() {}

func (x *GetAssetPublicURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_datatrails_common_api_assets_v2_assets_getasset_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssetPublicURLRequest.ProtoReflect.Descriptor instead.
func (*GetAssetPublicURLRequest) Descriptor() ([]byte, []int) {
	return file_datatrails_common_api_assets_v2_assets_getasset_proto_rawDescGZIP(), []int{1}
}

func (x *GetAssetPublicURLRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

var File_datatrails_common_api_assets_v2_assets_getasset_proto protoreflect.FileDescriptor

var file_datatrails_common_api_assets_v2_assets_getasset_proto_rawDesc = []byte{
	0x0a, 0x35, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x73, 0x2d, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76,
	0x32, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69,
	0x73, 0x74, 0x2e, 0x76, 0x32, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xea, 0x02, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0xd0, 0x01, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0xbb, 0x01, 0x92, 0x41, 0xaf, 0x01, 0x32, 0xaa, 0x01, 0x53, 0x70, 0x65, 0x63,
	0x69, 0x66, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20, 0x41, 0x73, 0x73, 0x65, 0x74, 0x20, 0x55, 0x55,
	0x49, 0x44, 0x20, 0x77, 0x68, 0x65, 0x72, 0x65, 0x20, 0x60, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73,
	0x2f, 0x7b, 0x75, 0x75, 0x69, 0x64, 0x7d, 0x60, 0x20, 0x69, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x20, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x20, 0x65,
	0x2e, 0x67, 0x2e, 0x20, 0x60, 0x61, 0x64, 0x64, 0x33, 0x30, 0x32, 0x33, 0x35, 0x2d, 0x31, 0x34,
	0x32, 0x34, 0x2d, 0x34, 0x66, 0x64, 0x61, 0x2d, 0x38, 0x34, 0x30, 0x61, 0x2d, 0x64, 0x35, 0x65,
	0x66, 0x38, 0x32, 0x63, 0x34, 0x63, 0x39, 0x36, 0x66, 0x60, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x20, 0x60, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73,
	0x2f, 0x61, 0x64, 0x64, 0x33, 0x30, 0x32, 0x33, 0x35, 0x2d, 0x31, 0x34, 0x32, 0x34, 0x2d, 0x34,
	0x66, 0x64, 0x61, 0x2d, 0x38, 0x34, 0x30, 0x61, 0x2d, 0x64, 0x35, 0x65, 0x66, 0x38, 0x32, 0x63,
	0x34, 0x63, 0x39, 0x36, 0x66, 0x60, 0x40, 0x01, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01,
	0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x83, 0x01, 0x0a, 0x07, 0x61, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x4e, 0x92, 0x41, 0x4b, 0x32, 0x47, 0x53, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x79, 0x20, 0x74, 0x69, 0x6d, 0x65, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70,
	0x61, 0x73, 0x74, 0x20, 0x74, 0x6f, 0x20, 0x73, 0x68, 0x6f, 0x77, 0x20, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x61, 0x73, 0x20, 0x69, 0x74, 0x20, 0x77, 0x61, 0x73,
	0x20, 0x61, 0x74, 0x20, 0x74, 0x69, 0x6d, 0x65, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x40, 0x01, 0x52, 0x06, 0x61, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xed, 0x01, 0x0a,
	0x18, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x55,
	0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0xd0, 0x01, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0xbb, 0x01, 0x92, 0x41, 0xaf, 0x01, 0x32,
	0xaa, 0x01, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x20, 0x55, 0x55, 0x49, 0x44, 0x20, 0x77, 0x68, 0x65, 0x72, 0x65, 0x20, 0x60,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x75, 0x75, 0x69, 0x64, 0x7d, 0x60, 0x20, 0x69,
	0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x41, 0x73, 0x73, 0x65, 0x74, 0x20, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x20, 0x65, 0x2e, 0x67, 0x2e, 0x20, 0x60, 0x61, 0x64, 0x64, 0x33, 0x30,
	0x32, 0x33, 0x35, 0x2d, 0x31, 0x34, 0x32, 0x34, 0x2d, 0x34, 0x66, 0x64, 0x61, 0x2d, 0x38, 0x34,
	0x30, 0x61, 0x2d, 0x64, 0x35, 0x65, 0x66, 0x38, 0x32, 0x63, 0x34, 0x63, 0x39, 0x36, 0x66, 0x60,
	0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x20, 0x60,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x61, 0x64, 0x64, 0x33, 0x30, 0x32, 0x33, 0x35, 0x2d,
	0x31, 0x34, 0x32, 0x34, 0x2d, 0x34, 0x66, 0x64, 0x61, 0x2d, 0x38, 0x34, 0x30, 0x61, 0x2d, 0x64,
	0x35, 0x65, 0x66, 0x38, 0x32, 0x63, 0x34, 0x63, 0x39, 0x36, 0x66, 0x60, 0x40, 0x01, 0xfa, 0x42,
	0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x42, 0x4c, 0x5a, 0x4a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x74,
	0x72, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61,
	0x69, 0x6c, 0x73, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67,
	0x65, 0x6e, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x73, 0x3b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_datatrails_common_api_assets_v2_assets_getasset_proto_rawDescOnce sync.Once
	file_datatrails_common_api_assets_v2_assets_getasset_proto_rawDescData = file_datatrails_common_api_assets_v2_assets_getasset_proto_rawDesc
)

func file_datatrails_common_api_assets_v2_assets_getasset_proto_rawDescGZIP() []byte {
	file_datatrails_common_api_assets_v2_assets_getasset_proto_rawDescOnce.Do(func() {
		file_datatrails_common_api_assets_v2_assets_getasset_proto_rawDescData = protoimpl.X.CompressGZIP(file_datatrails_common_api_assets_v2_assets_getasset_proto_rawDescData)
	})
	return file_datatrails_common_api_assets_v2_assets_getasset_proto_rawDescData
}

var file_datatrails_common_api_assets_v2_assets_getasset_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_datatrails_common_api_assets_v2_assets_getasset_proto_goTypes = []any{
	(*GetAssetRequest)(nil),          // 0: archivist.v2.GetAssetRequest
	(*GetAssetPublicURLRequest)(nil), // 1: archivist.v2.GetAssetPublicURLRequest
	(*timestamppb.Timestamp)(nil),    // 2: google.protobuf.Timestamp
}
var file_datatrails_common_api_assets_v2_assets_getasset_proto_depIdxs = []int32{
	2, // 0: archivist.v2.GetAssetRequest.at_time:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_datatrails_common_api_assets_v2_assets_getasset_proto_init() }
func file_datatrails_common_api_assets_v2_assets_getasset_proto_init() {
	if File_datatrails_common_api_assets_v2_assets_getasset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datatrails_common_api_assets_v2_assets_getasset_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetAssetRequest); i {
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
		file_datatrails_common_api_assets_v2_assets_getasset_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetAssetPublicURLRequest); i {
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
			RawDescriptor: file_datatrails_common_api_assets_v2_assets_getasset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datatrails_common_api_assets_v2_assets_getasset_proto_goTypes,
		DependencyIndexes: file_datatrails_common_api_assets_v2_assets_getasset_proto_depIdxs,
		MessageInfos:      file_datatrails_common_api_assets_v2_assets_getasset_proto_msgTypes,
	}.Build()
	File_datatrails_common_api_assets_v2_assets_getasset_proto = out.File
	file_datatrails_common_api_assets_v2_assets_getasset_proto_rawDesc = nil
	file_datatrails_common_api_assets_v2_assets_getasset_proto_goTypes = nil
	file_datatrails_common_api_assets_v2_assets_getasset_proto_depIdxs = nil
}
