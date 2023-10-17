// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: tenancies/v1/tenancies/idtokenhint.proto

package tenancies

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

type InviteTokenHintRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	TenantId string `protobuf:"bytes,2,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *InviteTokenHintRequest) Reset() {
	*x = InviteTokenHintRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenancies_v1_tenancies_idtokenhint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteTokenHintRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteTokenHintRequest) ProtoMessage() {}

func (x *InviteTokenHintRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tenancies_v1_tenancies_idtokenhint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteTokenHintRequest.ProtoReflect.Descriptor instead.
func (*InviteTokenHintRequest) Descriptor() ([]byte, []int) {
	return file_tenancies_v1_tenancies_idtokenhint_proto_rawDescGZIP(), []int{0}
}

func (x *InviteTokenHintRequest) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *InviteTokenHintRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *InviteTokenHintRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type InviteTokenHintResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ExpiryTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expiry_time,json=expiryTime,proto3" json:"expiry_time,omitempty"`
}

func (x *InviteTokenHintResponse) Reset() {
	*x = InviteTokenHintResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenancies_v1_tenancies_idtokenhint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteTokenHintResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteTokenHintResponse) ProtoMessage() {}

func (x *InviteTokenHintResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tenancies_v1_tenancies_idtokenhint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteTokenHintResponse.ProtoReflect.Descriptor instead.
func (*InviteTokenHintResponse) Descriptor() ([]byte, []int) {
	return file_tenancies_v1_tenancies_idtokenhint_proto_rawDescGZIP(), []int{1}
}

func (x *InviteTokenHintResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *InviteTokenHintResponse) GetExpiryTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiryTime
	}
	return nil
}

type UserTenantIDTokenHintRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantId string `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

func (x *UserTenantIDTokenHintRequest) Reset() {
	*x = UserTenantIDTokenHintRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenancies_v1_tenancies_idtokenhint_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTenantIDTokenHintRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTenantIDTokenHintRequest) ProtoMessage() {}

func (x *UserTenantIDTokenHintRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tenancies_v1_tenancies_idtokenhint_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTenantIDTokenHintRequest.ProtoReflect.Descriptor instead.
func (*UserTenantIDTokenHintRequest) Descriptor() ([]byte, []int) {
	return file_tenancies_v1_tenancies_idtokenhint_proto_rawDescGZIP(), []int{2}
}

func (x *UserTenantIDTokenHintRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

type UserTenantIDTokenHintResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *UserTenantIDTokenHintResponse) Reset() {
	*x = UserTenantIDTokenHintResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenancies_v1_tenancies_idtokenhint_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTenantIDTokenHintResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTenantIDTokenHintResponse) ProtoMessage() {}

func (x *UserTenantIDTokenHintResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tenancies_v1_tenancies_idtokenhint_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTenantIDTokenHintResponse.ProtoReflect.Descriptor instead.
func (*UserTenantIDTokenHintResponse) Descriptor() ([]byte, []int) {
	return file_tenancies_v1_tenancies_idtokenhint_proto_rawDescGZIP(), []int{3}
}

func (x *UserTenantIDTokenHintResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_tenancies_v1_tenancies_idtokenhint_proto protoreflect.FileDescriptor

var file_tenancies_v1_tenancies_idtokenhint_proto_rawDesc = []byte{
	0x0a, 0x28, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x69, 0x64, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x68, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a, 0x16, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x48, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0x80, 0x08, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x25, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0x80, 0x08,
	0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03,
	0x18, 0x80, 0x08, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x87, 0x01, 0x0a, 0x17, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x48, 0x69, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x56, 0x0a, 0x0b,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x19, 0x92,
	0x41, 0x16, 0x32, 0x12, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x20, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x79, 0x20, 0x74, 0x69, 0x6d, 0x65, 0x40, 0x01, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x45, 0x0a, 0x1c, 0x55, 0x73, 0x65, 0x72, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x49, 0x44, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x48, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0x80,
	0x08, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x1d, 0x55,
	0x73, 0x65, 0x72, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x48, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x6b, 0x76, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2d, 0x72, 0x6b, 0x76, 0x73, 0x74, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x63, 0x69, 0x65, 0x73, 0x3b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tenancies_v1_tenancies_idtokenhint_proto_rawDescOnce sync.Once
	file_tenancies_v1_tenancies_idtokenhint_proto_rawDescData = file_tenancies_v1_tenancies_idtokenhint_proto_rawDesc
)

func file_tenancies_v1_tenancies_idtokenhint_proto_rawDescGZIP() []byte {
	file_tenancies_v1_tenancies_idtokenhint_proto_rawDescOnce.Do(func() {
		file_tenancies_v1_tenancies_idtokenhint_proto_rawDescData = protoimpl.X.CompressGZIP(file_tenancies_v1_tenancies_idtokenhint_proto_rawDescData)
	})
	return file_tenancies_v1_tenancies_idtokenhint_proto_rawDescData
}

var file_tenancies_v1_tenancies_idtokenhint_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tenancies_v1_tenancies_idtokenhint_proto_goTypes = []interface{}{
	(*InviteTokenHintRequest)(nil),        // 0: archivist.v1.InviteTokenHintRequest
	(*InviteTokenHintResponse)(nil),       // 1: archivist.v1.InviteTokenHintResponse
	(*UserTenantIDTokenHintRequest)(nil),  // 2: archivist.v1.UserTenantIDTokenHintRequest
	(*UserTenantIDTokenHintResponse)(nil), // 3: archivist.v1.UserTenantIDTokenHintResponse
	(*timestamppb.Timestamp)(nil),         // 4: google.protobuf.Timestamp
}
var file_tenancies_v1_tenancies_idtokenhint_proto_depIdxs = []int32{
	4, // 0: archivist.v1.InviteTokenHintResponse.expiry_time:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tenancies_v1_tenancies_idtokenhint_proto_init() }
func file_tenancies_v1_tenancies_idtokenhint_proto_init() {
	if File_tenancies_v1_tenancies_idtokenhint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tenancies_v1_tenancies_idtokenhint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteTokenHintRequest); i {
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
		file_tenancies_v1_tenancies_idtokenhint_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteTokenHintResponse); i {
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
		file_tenancies_v1_tenancies_idtokenhint_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTenantIDTokenHintRequest); i {
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
		file_tenancies_v1_tenancies_idtokenhint_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTenantIDTokenHintResponse); i {
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
			RawDescriptor: file_tenancies_v1_tenancies_idtokenhint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tenancies_v1_tenancies_idtokenhint_proto_goTypes,
		DependencyIndexes: file_tenancies_v1_tenancies_idtokenhint_proto_depIdxs,
		MessageInfos:      file_tenancies_v1_tenancies_idtokenhint_proto_msgTypes,
	}.Build()
	File_tenancies_v1_tenancies_idtokenhint_proto = out.File
	file_tenancies_v1_tenancies_idtokenhint_proto_rawDesc = nil
	file_tenancies_v1_tenancies_idtokenhint_proto_goTypes = nil
	file_tenancies_v1_tenancies_idtokenhint_proto_depIdxs = nil
}
