// Maintainers, please refer to the style guide here:
//     https://developers.google.com/protocol-buffers/docs/style

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.24.3
// source: datatrails-common-api/assets/v2/assets/principal.proto

package assets

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Principal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Issuer      string `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Subject     string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Email       string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Principal) Reset() {
	*x = Principal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datatrails_common_api_assets_v2_assets_principal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Principal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Principal) ProtoMessage() {}

func (x *Principal) ProtoReflect() protoreflect.Message {
	mi := &file_datatrails_common_api_assets_v2_assets_principal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Principal.ProtoReflect.Descriptor instead.
func (*Principal) Descriptor() ([]byte, []int) {
	return file_datatrails_common_api_assets_v2_assets_principal_proto_rawDescGZIP(), []int{0}
}

func (x *Principal) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *Principal) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Principal) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Principal) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_datatrails_common_api_assets_v2_assets_principal_proto protoreflect.FileDescriptor

var file_datatrails_common_api_assets_v2_assets_principal_proto_rawDesc = []byte{
	0x0a, 0x36, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x73, 0x2d, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76,
	0x32, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x69, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc3, 0x07, 0x0a, 0x09, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x12, 0x9c, 0x01,
	0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x83,
	0x01, 0x92, 0x41, 0x78, 0x32, 0x73, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x20, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x72, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x72, 0x69,
	0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x20, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x20, 0x57, 0x68, 0x65, 0x72, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x72, 0x20, 0x69, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x64, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x20, 0x69, 0x73,
	0x20, 0x74, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x20, 0x61, 0x73, 0x20, 0x61, 0x20, 0x66, 0x72,
	0x65, 0x65, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x78, 0x80, 0x08, 0xfa, 0x42, 0x05, 0x72,
	0x03, 0x18, 0x80, 0x08, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x64, 0x0a, 0x07,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x4a, 0x92,
	0x41, 0x3f, 0x32, 0x3a, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x20, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x72, 0x69,
	0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x20, 0x28, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x20, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x72, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x29, 0x78, 0x80,
	0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0x80, 0x08, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0xcc, 0x01, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0xa8, 0x01, 0x92, 0x41, 0x9c, 0x01,
	0x32, 0x96, 0x01, 0x54, 0x68, 0x65, 0x20, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x61, 0x62,
	0x6c, 0x65, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x65,
	0x6e, 0x64, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x20, 0x20, 0x54, 0x68, 0x65, 0x20, 0x6e, 0x61,
	0x6d, 0x65, 0x20, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x20, 0x69, 0x73, 0x20, 0x70, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x64, 0x2c, 0x20, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x20, 0x62,
	0x79, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x20, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x2c, 0x20,
	0x74, 0x68, 0x65, 0x6e, 0x20, 0x61, 0x20, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65,
	0x20, 0x6f, 0x66, 0x20, 0x20, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x2c,
	0x20, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x2c, 0x20, 0x66, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x78, 0x80, 0x08, 0xfa, 0x42, 0x05, 0x72,
	0x03, 0x18, 0x80, 0x08, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0xc2, 0x01, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0xab, 0x01, 0x92, 0x41, 0x9f, 0x01, 0x32, 0x99, 0x01, 0x54, 0x68, 0x65, 0x20, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x65, 0x6e, 0x64,
	0x2d, 0x75, 0x73, 0x65, 0x72, 0x20, 0x69, 0x66, 0x20, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x20, 0x49, 0x66, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x20, 0x69, 0x73, 0x20, 0x20, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x20, 0x69, 0x74, 0x20, 0x69, 0x73, 0x20, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x64, 0x2e, 0x20, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x20, 0x69, 0x66, 0x20, 0x6e, 0x65, 0x69, 0x74,
	0x68, 0x65, 0x72, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x20, 0x6f, 0x72, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x20, 0x20, 0x61, 0x72, 0x65,
	0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x49, 0x64, 0x50, 0x78, 0x80, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0x80, 0x08, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x3a, 0x9c, 0x02, 0x92, 0x41, 0x98, 0x02, 0x0a, 0xdb, 0x01,
	0x32, 0xd8, 0x01, 0x54, 0x68, 0x65, 0x20, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c,
	0x20, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x61, 0x73, 0x73,
	0x75, 0x72, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x20, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x20,
	0x20, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x20, 0x41, 0x6c, 0x6c, 0x20, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x20, 0x61, 0x72, 0x65, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x20, 0x74, 0x6f, 0x20, 0x4f, 0x49, 0x44, 0x43, 0x20, 0x69, 0x64, 0x20, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x20, 0x61, 0x6e, 0x64, 0x20,
	0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73,
	0x2e, 0x20, 0x20, 0x53, 0x65, 0x65, 0x20, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x6f,
	0x70, 0x65, 0x6e, 0x69, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f,
	0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x2d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2d, 0x63,
	0x6f, 0x72, 0x65, 0x2d, 0x31, 0x5f, 0x30, 0x2e, 0x68, 0x74, 0x6d, 0x6c, 0x23, 0x53, 0x74, 0x61,
	0x6e, 0x64, 0x61, 0x72, 0x64, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x32, 0x38, 0x7b, 0x20, 0x22,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x22, 0x3a, 0x20, 0x22, 0x6a, 0x6f, 0x62, 0x2e, 0x69, 0x64,
	0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x31, 0x32, 0x33, 0x34, 0x22, 0x2c, 0x20,
	0x22, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x3a, 0x20, 0x22, 0x62, 0x6f, 0x62, 0x40,
	0x6a, 0x6f, 0x62, 0x22, 0x7d, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x67,
	0x6f, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x73, 0x2d, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x3b, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_datatrails_common_api_assets_v2_assets_principal_proto_rawDescOnce sync.Once
	file_datatrails_common_api_assets_v2_assets_principal_proto_rawDescData = file_datatrails_common_api_assets_v2_assets_principal_proto_rawDesc
)

func file_datatrails_common_api_assets_v2_assets_principal_proto_rawDescGZIP() []byte {
	file_datatrails_common_api_assets_v2_assets_principal_proto_rawDescOnce.Do(func() {
		file_datatrails_common_api_assets_v2_assets_principal_proto_rawDescData = protoimpl.X.CompressGZIP(file_datatrails_common_api_assets_v2_assets_principal_proto_rawDescData)
	})
	return file_datatrails_common_api_assets_v2_assets_principal_proto_rawDescData
}

var file_datatrails_common_api_assets_v2_assets_principal_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_datatrails_common_api_assets_v2_assets_principal_proto_goTypes = []any{
	(*Principal)(nil), // 0: archivist.v2.Principal
}
var file_datatrails_common_api_assets_v2_assets_principal_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_datatrails_common_api_assets_v2_assets_principal_proto_init() }
func file_datatrails_common_api_assets_v2_assets_principal_proto_init() {
	if File_datatrails_common_api_assets_v2_assets_principal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datatrails_common_api_assets_v2_assets_principal_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Principal); i {
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
			RawDescriptor: file_datatrails_common_api_assets_v2_assets_principal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datatrails_common_api_assets_v2_assets_principal_proto_goTypes,
		DependencyIndexes: file_datatrails_common_api_assets_v2_assets_principal_proto_depIdxs,
		MessageInfos:      file_datatrails_common_api_assets_v2_assets_principal_proto_msgTypes,
	}.Build()
	File_datatrails_common_api_assets_v2_assets_principal_proto = out.File
	file_datatrails_common_api_assets_v2_assets_principal_proto_rawDesc = nil
	file_datatrails_common_api_assets_v2_assets_principal_proto_goTypes = nil
	file_datatrails_common_api_assets_v2_assets_principal_proto_depIdxs = nil
}
