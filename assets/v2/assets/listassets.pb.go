// Maintainers, please refer to the style guide here:
//     https://developers.google.com/protocol-buffers/docs/style

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: datatrails-common-api/assets/v2/assets/listassets.proto

package assets

import (
	filter "github.com/datatrails/go-datatrails-common-api-gen/filter/v1/filter"
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

// Supported sort orders
type ListAssetsRequest_OrderBy int32

const (
	ListAssetsRequest_DEFAULT      ListAssetsRequest_OrderBy = 0
	ListAssetsRequest_SIMPLEHASHV1 ListAssetsRequest_OrderBy = 1
	ListAssetsRequest_SIMPLEHASHV2 ListAssetsRequest_OrderBy = 2
)

// Enum value maps for ListAssetsRequest_OrderBy.
var (
	ListAssetsRequest_OrderBy_name = map[int32]string{
		0: "DEFAULT",
		1: "SIMPLEHASHV1",
		2: "SIMPLEHASHV2",
	}
	ListAssetsRequest_OrderBy_value = map[string]int32{
		"DEFAULT":      0,
		"SIMPLEHASHV1": 1,
		"SIMPLEHASHV2": 2,
	}
)

func (x ListAssetsRequest_OrderBy) Enum() *ListAssetsRequest_OrderBy {
	p := new(ListAssetsRequest_OrderBy)
	*p = x
	return p
}

func (x ListAssetsRequest_OrderBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListAssetsRequest_OrderBy) Descriptor() protoreflect.EnumDescriptor {
	return file_datatrails_common_api_assets_v2_assets_listassets_proto_enumTypes[0].Descriptor()
}

func (ListAssetsRequest_OrderBy) Type() protoreflect.EnumType {
	return &file_datatrails_common_api_assets_v2_assets_listassets_proto_enumTypes[0]
}

func (x ListAssetsRequest_OrderBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListAssetsRequest_OrderBy.Descriptor instead.
func (ListAssetsRequest_OrderBy) EnumDescriptor() ([]byte, []int) {
	return file_datatrails_common_api_assets_v2_assets_listassets_proto_rawDescGZIP(), []int{0, 0}
}

type ListAssetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Maximum results per page.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token returned from a previous list request if any.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Specify the sort order for the results.
	OrderBy ListAssetsRequest_OrderBy `protobuf:"varint,3,opt,name=order_by,json=orderBy,proto3,enum=archivist.v2.ListAssetsRequest_OrderBy" json:"order_by,omitempty"`
	// indicates whether asset is still being tracked in the system
	Tracked TrackedStatus `protobuf:"varint,4,opt,name=tracked,proto3,enum=archivist.v2.TrackedStatus" json:"tracked,omitempty"`
	// indicates if the asset has been succesfully committed to the blockchain
	ConfirmationStatus ConfirmationStatus `protobuf:"varint,5,opt,name=confirmation_status,json=confirmationStatus,proto3,enum=archivist.v2.ConfirmationStatus" json:"confirmation_status,omitempty"`
	// user defined attributes key max_length: 1024 value max_length: 4096
	Attributes map[string]string `protobuf:"bytes,6,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// proof mechanism of the asset (and all its events)
	ProofMechanism ProofMechanism `protobuf:"varint,8,opt,name=proof_mechanism,json=proofMechanism,proto3,enum=archivist.v2.ProofMechanism" json:"proof_mechanism,omitempty"`
	ChainId        string         `protobuf:"bytes,9,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// Outer list AND, inner list INCLUSIVE OR, e.g. [[A, B], [A, B]] "(A OR B) in first position AND (A OR B) in second position (and anything after)"
	// The behaviour of these filters matches the filters available for
	// selecting assets in access policies:
	// On the Access Policy editor page in the UX wire frames, we see locations and a asset_types as two of the items to filter on. A policy that expresses “a policy for all door access readers in Basingstoke or Cambridge”
	// [
	//
	//	{ "or": ["location=basingstoke", "location=cambridge"] },
	//	{ "or": ["asset_type=door_access_reader"] }
	//
	// ]
	Filters []*filter.Filter `protobuf:"bytes,10,rep,name=filters,proto3" json:"filters,omitempty"`
	// privacy filter of the asset (and all its events)
	Privacy Privacy `protobuf:"varint,12,opt,name=privacy,proto3,enum=archivist.v2.Privacy" json:"privacy,omitempty"`
}

func (x *ListAssetsRequest) Reset() {
	*x = ListAssetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datatrails_common_api_assets_v2_assets_listassets_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAssetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAssetsRequest) ProtoMessage() {}

func (x *ListAssetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_datatrails_common_api_assets_v2_assets_listassets_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAssetsRequest.ProtoReflect.Descriptor instead.
func (*ListAssetsRequest) Descriptor() ([]byte, []int) {
	return file_datatrails_common_api_assets_v2_assets_listassets_proto_rawDescGZIP(), []int{0}
}

func (x *ListAssetsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListAssetsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListAssetsRequest) GetOrderBy() ListAssetsRequest_OrderBy {
	if x != nil {
		return x.OrderBy
	}
	return ListAssetsRequest_DEFAULT
}

func (x *ListAssetsRequest) GetTracked() TrackedStatus {
	if x != nil {
		return x.Tracked
	}
	return TrackedStatus_TRACKED_STATUS_UNSPECIFIED
}

func (x *ListAssetsRequest) GetConfirmationStatus() ConfirmationStatus {
	if x != nil {
		return x.ConfirmationStatus
	}
	return ConfirmationStatus_CONFIRMATION_STATUS_UNSPECIFIED
}

func (x *ListAssetsRequest) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *ListAssetsRequest) GetProofMechanism() ProofMechanism {
	if x != nil {
		return x.ProofMechanism
	}
	return ProofMechanism_PROOF_MECHANISM_UNSPECIFIED
}

func (x *ListAssetsRequest) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

func (x *ListAssetsRequest) GetFilters() []*filter.Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *ListAssetsRequest) GetPrivacy() Privacy {
	if x != nil {
		return x.Privacy
	}
	return Privacy_PRIVACY_UNSPECIFIED
}

type ListAssetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Assets []*AssetResponse `protobuf:"bytes,1,rep,name=assets,proto3" json:"assets,omitempty"`
	// Token to retrieve the next page of results or empty if there are none.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListAssetsResponse) Reset() {
	*x = ListAssetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datatrails_common_api_assets_v2_assets_listassets_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAssetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAssetsResponse) ProtoMessage() {}

func (x *ListAssetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_datatrails_common_api_assets_v2_assets_listassets_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAssetsResponse.ProtoReflect.Descriptor instead.
func (*ListAssetsResponse) Descriptor() ([]byte, []int) {
	return file_datatrails_common_api_assets_v2_assets_listassets_proto_rawDescGZIP(), []int{1}
}

func (x *ListAssetsResponse) GetAssets() []*AssetResponse {
	if x != nil {
		return x.Assets
	}
	return nil
}

func (x *ListAssetsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_datatrails_common_api_assets_v2_assets_listassets_proto protoreflect.FileDescriptor

var file_datatrails_common_api_assets_v2_assets_listassets_proto_rawDesc = []byte{
	0x0a, 0x37, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x73, 0x2d, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76,
	0x32, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x33, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x73, 0x2d, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x69, 0x6c,
	0x73, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x64, 0x61, 0x74, 0x61, 0x74,
	0x72, 0x61, 0x69, 0x6c, 0x73, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x07, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x1a, 0x02, 0x28, 0x00, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x42, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x27, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x2e, 0x76,
	0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x42, 0x79, 0x12, 0x35, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73,
	0x74, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x51, 0x0a, 0x13, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x8d,
	0x01, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x2e,
	0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x42, 0x3c, 0x92, 0x41, 0x19, 0x32, 0x17, 0x75, 0x73, 0x65, 0x72, 0x20,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x20, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0xfa, 0x42, 0x1d, 0x9a, 0x01, 0x1a, 0x22, 0x18, 0x72, 0x16, 0x10, 0x01, 0x18, 0x80,
	0x08, 0x32, 0x0f, 0x5e, 0x5b, 0x5e, 0x5b, 0x3a, 0x63, 0x6e, 0x74, 0x72, 0x6c, 0x3a, 0x5d, 0x5d,
	0x2b, 0x24, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x7a,
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73,
	0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x69, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x4d, 0x65, 0x63, 0x68,
	0x61, 0x6e, 0x69, 0x73, 0x6d, 0x42, 0x33, 0x92, 0x41, 0x30, 0x32, 0x2e, 0x74, 0x68, 0x65, 0x20,
	0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x74,
	0x6f, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x20, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x20, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x6f,
	0x66, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x12, 0x57, 0x0a, 0x08, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3c, 0x92, 0x41,
	0x39, 0x32, 0x35, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x20, 0x69, 0x64, 0x20, 0x6f, 0x66, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x20, 0x61, 0x73,
	0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x74, 0x68,
	0x69, 0x73, 0x20, 0x61, 0x73, 0x73, 0x65, 0x74, 0x40, 0x01, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x0c, 0x92, 0x41, 0x09, 0x32,
	0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x12, 0x5d, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x15, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x2e, 0x76,
	0x32, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x42, 0x2c, 0x92, 0x41, 0x21, 0x32, 0x1f,
	0x74, 0x68, 0x65, 0x20, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x20, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x61, 0x73, 0x73, 0x65, 0x74, 0xfa,
	0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x1a, 0x3d, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x3a, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45,
	0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x49, 0x4d, 0x50, 0x4c,
	0x45, 0x48, 0x41, 0x53, 0x48, 0x56, 0x31, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x49, 0x4d,
	0x50, 0x4c, 0x45, 0x48, 0x41, 0x53, 0x48, 0x56, 0x32, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x07, 0x10,
	0x08, 0x4a, 0x04, 0x08, 0x0b, 0x10, 0x0c, 0x52, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x69, 0x74, 0x79, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x22, 0xfb, 0x0b, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x12, 0x26,
	0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x3a, 0x87, 0x0b, 0x92, 0x41, 0x83, 0x0b, 0x0a, 0x35, 0x32,
	0x33, 0x41, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x73, 0x20, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x20, 0x74, 0x6f, 0x20, 0x60, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x60, 0x32, 0xc9, 0x0a, 0x7b, 0x20, 0x20, 0x20, 0x22, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x73, 0x22, 0x3a, 0x20, 0x5b, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7b, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x22, 0x3a, 0x20, 0x22, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x61, 0x64, 0x64,
	0x33, 0x30, 0x32, 0x33, 0x35, 0x2d, 0x31, 0x34, 0x32, 0x34, 0x2d, 0x34, 0x66, 0x64, 0x61, 0x2d,
	0x38, 0x34, 0x30, 0x61, 0x2d, 0x64, 0x35, 0x65, 0x66, 0x38, 0x32, 0x63, 0x34, 0x63, 0x39, 0x36,
	0x66, 0x22, 0x2c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x73, 0x22, 0x3a, 0x20, 0x5b, 0x22, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x5d, 0x2c,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0x3a, 0x20, 0x7b, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x61, 0x72, 0x63, 0x5f,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x20, 0x22,
	0x47, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x20, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x2c, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x61,
	0x72, 0x63, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x3a, 0x20, 0x22, 0x4d, 0x79, 0x20, 0x47, 0x61, 0x72, 0x64, 0x65, 0x6e, 0x20, 0x46, 0x65, 0x6e,
	0x63, 0x65, 0x22, 0x2c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x22, 0x63, 0x6f, 0x6c, 0x6f, 0x75, 0x72, 0x22, 0x3a, 0x20, 0x22, 0x50,
	0x6c, 0x61, 0x69, 0x6e, 0x20, 0x77, 0x6f, 0x6f, 0x64, 0x22, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x2c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x22, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3a, 0x20, 0x22, 0x50, 0x45, 0x4e, 0x44, 0x49,
	0x4e, 0x47, 0x22, 0x2c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x22, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x54, 0x52, 0x41, 0x43,
	0x4b, 0x45, 0x44, 0x22, 0x2c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x22, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x3a, 0x20, 0x22, 0x30, 0x78, 0x36, 0x30, 0x31,
	0x66, 0x35, 0x41, 0x37, 0x44, 0x33, 0x65, 0x36, 0x64, 0x63, 0x42, 0x35, 0x35, 0x65, 0x38, 0x37,
	0x62, 0x66, 0x32, 0x46, 0x31, 0x37, 0x62, 0x43, 0x38, 0x41, 0x32, 0x37, 0x41, 0x61, 0x43, 0x44,
	0x33, 0x35, 0x31, 0x31, 0x22, 0x2c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x22, 0x61, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x32, 0x30, 0x31,
	0x39, 0x2d, 0x31, 0x31, 0x2d, 0x32, 0x37, 0x54, 0x31, 0x34, 0x3a, 0x34, 0x34, 0x3a, 0x31, 0x39,
	0x5a, 0x22, 0x2c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22,
	0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x22,
	0x3a, 0x20, 0x22, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x48, 0x41, 0x53, 0x48, 0x22, 0x2c,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x22, 0x3a, 0x20, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x2c, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x3a, 0x20, 0x22, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x2f, 0x38, 0x65, 0x30, 0x62, 0x36, 0x30, 0x30, 0x63, 0x2d, 0x38, 0x32, 0x33, 0x34, 0x2d, 0x34,
	0x33, 0x65, 0x34, 0x2d, 0x38, 0x36, 0x30, 0x63, 0x2d, 0x65, 0x39, 0x35, 0x62, 0x64, 0x63, 0x64,
	0x36, 0x39, 0x35, 0x61, 0x39, 0x22, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x2c,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7b, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x22, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x3a, 0x20, 0x22,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x63, 0x65, 0x66, 0x36, 0x31, 0x33, 0x34, 0x36, 0x2d,
	0x32, 0x34, 0x35, 0x33, 0x2d, 0x35, 0x61, 0x65, 0x62, 0x2d, 0x39, 0x32, 0x31, 0x63, 0x2d, 0x65,
	0x36, 0x66, 0x61, 0x39, 0x33, 0x64, 0x35, 0x62, 0x30, 0x33, 0x32, 0x22, 0x2c, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x75, 0x72, 0x73, 0x22, 0x3a, 0x20, 0x5b, 0x22, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x45,
	0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x5d, 0x2c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x22, 0x3a, 0x20, 0x7b, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x61, 0x72, 0x63, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x49, 0x6f, 0x54, 0x20, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x22, 0x2c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x61, 0x72, 0x63, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x4d, 0x79, 0x20, 0x49, 0x6f,
	0x54, 0x20, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x2c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x61, 0x72, 0x63, 0x5f, 0x66,
	0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x3a, 0x20, 0x22, 0x33, 0x2e, 0x32, 0x2e, 0x31, 0x22, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x7d, 0x2c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x22, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3a, 0x20, 0x22, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e,
	0x47, 0x22, 0x2c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x54, 0x52, 0x41, 0x43, 0x4b,
	0x45, 0x44, 0x22, 0x2c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x22, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x3a, 0x20, 0x22, 0x30, 0x78, 0x36, 0x30, 0x31, 0x66,
	0x35, 0x41, 0x37, 0x44, 0x33, 0x65, 0x36, 0x64, 0x63, 0x42, 0x35, 0x35, 0x65, 0x38, 0x37, 0x62,
	0x66, 0x32, 0x46, 0x31, 0x37, 0x62, 0x43, 0x38, 0x41, 0x32, 0x37, 0x41, 0x61, 0x43, 0x44, 0x33,
	0x35, 0x31, 0x31, 0x22, 0x2c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x22, 0x61, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x32, 0x30, 0x31, 0x39,
	0x2d, 0x31, 0x31, 0x2d, 0x32, 0x37, 0x54, 0x31, 0x34, 0x3a, 0x34, 0x34, 0x3a, 0x31, 0x39, 0x5a,
	0x22, 0x2c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x70,
	0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x22, 0x3a,
	0x20, 0x22, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x48, 0x41, 0x53, 0x48, 0x22, 0x2c, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x22, 0x3a, 0x20, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x2c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x3a, 0x20, 0x22, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f,
	0x38, 0x65, 0x30, 0x62, 0x36, 0x30, 0x30, 0x63, 0x2d, 0x38, 0x32, 0x33, 0x34, 0x2d, 0x34, 0x33,
	0x65, 0x34, 0x2d, 0x38, 0x36, 0x30, 0x63, 0x2d, 0x65, 0x39, 0x35, 0x62, 0x64, 0x63, 0x64, 0x36,
	0x39, 0x35, 0x61, 0x39, 0x22, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x20, 0x20,
	0x20, 0x5d, 0x2c, 0x20, 0x20, 0x20, 0x22, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3a, 0x20, 0x22, 0x61, 0x62, 0x63, 0x64, 0x22, 0x7d,
	0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x64, 0x61, 0x74,
	0x61, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x73, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x61,
	0x70, 0x69, 0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x32,
	0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x3b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_datatrails_common_api_assets_v2_assets_listassets_proto_rawDescOnce sync.Once
	file_datatrails_common_api_assets_v2_assets_listassets_proto_rawDescData = file_datatrails_common_api_assets_v2_assets_listassets_proto_rawDesc
)

func file_datatrails_common_api_assets_v2_assets_listassets_proto_rawDescGZIP() []byte {
	file_datatrails_common_api_assets_v2_assets_listassets_proto_rawDescOnce.Do(func() {
		file_datatrails_common_api_assets_v2_assets_listassets_proto_rawDescData = protoimpl.X.CompressGZIP(file_datatrails_common_api_assets_v2_assets_listassets_proto_rawDescData)
	})
	return file_datatrails_common_api_assets_v2_assets_listassets_proto_rawDescData
}

var file_datatrails_common_api_assets_v2_assets_listassets_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_datatrails_common_api_assets_v2_assets_listassets_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_datatrails_common_api_assets_v2_assets_listassets_proto_goTypes = []interface{}{
	(ListAssetsRequest_OrderBy)(0), // 0: archivist.v2.ListAssetsRequest.OrderBy
	(*ListAssetsRequest)(nil),      // 1: archivist.v2.ListAssetsRequest
	(*ListAssetsResponse)(nil),     // 2: archivist.v2.ListAssetsResponse
	nil,                            // 3: archivist.v2.ListAssetsRequest.AttributesEntry
	(TrackedStatus)(0),             // 4: archivist.v2.TrackedStatus
	(ConfirmationStatus)(0),        // 5: archivist.v2.ConfirmationStatus
	(ProofMechanism)(0),            // 6: archivist.v2.ProofMechanism
	(*filter.Filter)(nil),          // 7: archivist.v1.Filter
	(Privacy)(0),                   // 8: archivist.v2.Privacy
	(*AssetResponse)(nil),          // 9: archivist.v2.AssetResponse
}
var file_datatrails_common_api_assets_v2_assets_listassets_proto_depIdxs = []int32{
	0, // 0: archivist.v2.ListAssetsRequest.order_by:type_name -> archivist.v2.ListAssetsRequest.OrderBy
	4, // 1: archivist.v2.ListAssetsRequest.tracked:type_name -> archivist.v2.TrackedStatus
	5, // 2: archivist.v2.ListAssetsRequest.confirmation_status:type_name -> archivist.v2.ConfirmationStatus
	3, // 3: archivist.v2.ListAssetsRequest.attributes:type_name -> archivist.v2.ListAssetsRequest.AttributesEntry
	6, // 4: archivist.v2.ListAssetsRequest.proof_mechanism:type_name -> archivist.v2.ProofMechanism
	7, // 5: archivist.v2.ListAssetsRequest.filters:type_name -> archivist.v1.Filter
	8, // 6: archivist.v2.ListAssetsRequest.privacy:type_name -> archivist.v2.Privacy
	9, // 7: archivist.v2.ListAssetsResponse.assets:type_name -> archivist.v2.AssetResponse
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_datatrails_common_api_assets_v2_assets_listassets_proto_init() }
func file_datatrails_common_api_assets_v2_assets_listassets_proto_init() {
	if File_datatrails_common_api_assets_v2_assets_listassets_proto != nil {
		return
	}
	file_datatrails_common_api_assets_v2_assets_enums_proto_init()
	file_datatrails_common_api_assets_v2_assets_assetresponse_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_datatrails_common_api_assets_v2_assets_listassets_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAssetsRequest); i {
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
		file_datatrails_common_api_assets_v2_assets_listassets_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAssetsResponse); i {
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
			RawDescriptor: file_datatrails_common_api_assets_v2_assets_listassets_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datatrails_common_api_assets_v2_assets_listassets_proto_goTypes,
		DependencyIndexes: file_datatrails_common_api_assets_v2_assets_listassets_proto_depIdxs,
		EnumInfos:         file_datatrails_common_api_assets_v2_assets_listassets_proto_enumTypes,
		MessageInfos:      file_datatrails_common_api_assets_v2_assets_listassets_proto_msgTypes,
	}.Build()
	File_datatrails_common_api_assets_v2_assets_listassets_proto = out.File
	file_datatrails_common_api_assets_v2_assets_listassets_proto_rawDesc = nil
	file_datatrails_common_api_assets_v2_assets_listassets_proto_goTypes = nil
	file_datatrails_common_api_assets_v2_assets_listassets_proto_depIdxs = nil
}
