// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.24.3
// source: datatrails-common-api/filter/v1/filter/filter.proto

//
// Structured filter for internal list of OR'd condition
// to be part of filter: [ {or: [foo=bar, baz=fiz]} ]
//

package filter

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Filter struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Or            []string               `protobuf:"bytes,1,rep,name=or,proto3" json:"or,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Filter) Reset() {
	*x = Filter{}
	mi := &file_datatrails_common_api_filter_v1_filter_filter_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_datatrails_common_api_filter_v1_filter_filter_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_datatrails_common_api_filter_v1_filter_filter_proto_rawDescGZIP(), []int{0}
}

func (x *Filter) GetOr() []string {
	if x != nil {
		return x.Or
	}
	return nil
}

var File_datatrails_common_api_filter_v1_filter_filter_proto protoreflect.FileDescriptor

const file_datatrails_common_api_filter_v1_filter_filter_proto_rawDesc = "" +
	"\n" +
	"3datatrails-common-api/filter/v1/filter/filter.proto\x12\farchivist.v1\x1a.protoc-gen-openapiv2/options/annotations.proto\x1a\x17validate/validate.proto\"\x8a\x01\n" +
	"\x06Filter\x123\n" +
	"\x02or\x18\x01 \x03(\tB#\x92A\x112\x0fThe filter list\xfaB\f\x92\x01\t\x18\x01\"\x05r\x03\x18\x80\bR\x02or:K\x92AH\n" +
	"\b2\x06Filter2<{      \"or\": [\"group=maintainers\", \"group=supervisors\"]    }BLZJgithub.com/datatrails/go-datatrails-common-api-gen/filter/v1/filter;filterb\x06proto3"

var (
	file_datatrails_common_api_filter_v1_filter_filter_proto_rawDescOnce sync.Once
	file_datatrails_common_api_filter_v1_filter_filter_proto_rawDescData []byte
)

func file_datatrails_common_api_filter_v1_filter_filter_proto_rawDescGZIP() []byte {
	file_datatrails_common_api_filter_v1_filter_filter_proto_rawDescOnce.Do(func() {
		file_datatrails_common_api_filter_v1_filter_filter_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_datatrails_common_api_filter_v1_filter_filter_proto_rawDesc), len(file_datatrails_common_api_filter_v1_filter_filter_proto_rawDesc)))
	})
	return file_datatrails_common_api_filter_v1_filter_filter_proto_rawDescData
}

var file_datatrails_common_api_filter_v1_filter_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_datatrails_common_api_filter_v1_filter_filter_proto_goTypes = []any{
	(*Filter)(nil), // 0: archivist.v1.Filter
}
var file_datatrails_common_api_filter_v1_filter_filter_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_datatrails_common_api_filter_v1_filter_filter_proto_init() }
func file_datatrails_common_api_filter_v1_filter_filter_proto_init() {
	if File_datatrails_common_api_filter_v1_filter_filter_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_datatrails_common_api_filter_v1_filter_filter_proto_rawDesc), len(file_datatrails_common_api_filter_v1_filter_filter_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datatrails_common_api_filter_v1_filter_filter_proto_goTypes,
		DependencyIndexes: file_datatrails_common_api_filter_v1_filter_filter_proto_depIdxs,
		MessageInfos:      file_datatrails_common_api_filter_v1_filter_filter_proto_msgTypes,
	}.Build()
	File_datatrails_common_api_filter_v1_filter_filter_proto = out.File
	file_datatrails_common_api_filter_v1_filter_filter_proto_goTypes = nil
	file_datatrails_common_api_filter_v1_filter_filter_proto_depIdxs = nil
}
