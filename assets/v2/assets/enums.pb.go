// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.24.3
// source: datatrails-common-api/assets/v2/assets/enums.proto

package assets

import (
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

type ConfirmationStatus int32

const (
	ConfirmationStatus_CONFIRMATION_STATUS_UNSPECIFIED ConfirmationStatus = 0
	ConfirmationStatus_PENDING                         ConfirmationStatus = 1 // Not yet stored
	ConfirmationStatus_CONFIRMED                       ConfirmationStatus = 2 // A tree root including the event has been signed by DataTrails
	ConfirmationStatus_FAILED                          ConfirmationStatus = 3 // Permanent failure
	ConfirmationStatus_STORED                          ConfirmationStatus = 4 // In database, awaiting verifiable commitment
	ConfirmationStatus_COMMITTED                       ConfirmationStatus = 5 // The stored event is verifiable
	ConfirmationStatus_UNEQUIVOCAL                     ConfirmationStatus = 6 // Provable independent of DataTrails
)

// Enum value maps for ConfirmationStatus.
var (
	ConfirmationStatus_name = map[int32]string{
		0: "CONFIRMATION_STATUS_UNSPECIFIED",
		1: "PENDING",
		2: "CONFIRMED",
		3: "FAILED",
		4: "STORED",
		5: "COMMITTED",
		6: "UNEQUIVOCAL",
	}
	ConfirmationStatus_value = map[string]int32{
		"CONFIRMATION_STATUS_UNSPECIFIED": 0,
		"PENDING":                         1,
		"CONFIRMED":                       2,
		"FAILED":                          3,
		"STORED":                          4,
		"COMMITTED":                       5,
		"UNEQUIVOCAL":                     6,
	}
)

func (x ConfirmationStatus) Enum() *ConfirmationStatus {
	p := new(ConfirmationStatus)
	*p = x
	return p
}

func (x ConfirmationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConfirmationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_datatrails_common_api_assets_v2_assets_enums_proto_enumTypes[0].Descriptor()
}

func (ConfirmationStatus) Type() protoreflect.EnumType {
	return &file_datatrails_common_api_assets_v2_assets_enums_proto_enumTypes[0]
}

func (x ConfirmationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConfirmationStatus.Descriptor instead.
func (ConfirmationStatus) EnumDescriptor() ([]byte, []int) {
	return file_datatrails_common_api_assets_v2_assets_enums_proto_rawDescGZIP(), []int{0}
}

type TrackedStatus int32

const (
	TrackedStatus_TRACKED_STATUS_UNSPECIFIED TrackedStatus = 0
	TrackedStatus_TRACKED                    TrackedStatus = 1
	TrackedStatus_NOT_TRACKED                TrackedStatus = 2
	TrackedStatus_ANY                        TrackedStatus = 3
)

// Enum value maps for TrackedStatus.
var (
	TrackedStatus_name = map[int32]string{
		0: "TRACKED_STATUS_UNSPECIFIED",
		1: "TRACKED",
		2: "NOT_TRACKED",
		3: "ANY",
	}
	TrackedStatus_value = map[string]int32{
		"TRACKED_STATUS_UNSPECIFIED": 0,
		"TRACKED":                    1,
		"NOT_TRACKED":                2,
		"ANY":                        3,
	}
)

func (x TrackedStatus) Enum() *TrackedStatus {
	p := new(TrackedStatus)
	*p = x
	return p
}

func (x TrackedStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TrackedStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_datatrails_common_api_assets_v2_assets_enums_proto_enumTypes[1].Descriptor()
}

func (TrackedStatus) Type() protoreflect.EnumType {
	return &file_datatrails_common_api_assets_v2_assets_enums_proto_enumTypes[1]
}

func (x TrackedStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TrackedStatus.Descriptor instead.
func (TrackedStatus) EnumDescriptor() ([]byte, []int) {
	return file_datatrails_common_api_assets_v2_assets_enums_proto_rawDescGZIP(), []int{1}
}

type ProofMechanism int32

const (
	ProofMechanism_PROOF_MECHANISM_UNSPECIFIED ProofMechanism = 0
	ProofMechanism_RESERVED1                   ProofMechanism = 1
	ProofMechanism_SIMPLE_HASH                 ProofMechanism = 2
	ProofMechanism_MERKLE_LOG                  ProofMechanism = 3
)

// Enum value maps for ProofMechanism.
var (
	ProofMechanism_name = map[int32]string{
		0: "PROOF_MECHANISM_UNSPECIFIED",
		1: "RESERVED1",
		2: "SIMPLE_HASH",
		3: "MERKLE_LOG",
	}
	ProofMechanism_value = map[string]int32{
		"PROOF_MECHANISM_UNSPECIFIED": 0,
		"RESERVED1":                   1,
		"SIMPLE_HASH":                 2,
		"MERKLE_LOG":                  3,
	}
)

func (x ProofMechanism) Enum() *ProofMechanism {
	p := new(ProofMechanism)
	*p = x
	return p
}

func (x ProofMechanism) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProofMechanism) Descriptor() protoreflect.EnumDescriptor {
	return file_datatrails_common_api_assets_v2_assets_enums_proto_enumTypes[2].Descriptor()
}

func (ProofMechanism) Type() protoreflect.EnumType {
	return &file_datatrails_common_api_assets_v2_assets_enums_proto_enumTypes[2]
}

func (x ProofMechanism) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProofMechanism.Descriptor instead.
func (ProofMechanism) EnumDescriptor() ([]byte, []int) {
	return file_datatrails_common_api_assets_v2_assets_enums_proto_rawDescGZIP(), []int{2}
}

type Privacy int32

const (
	Privacy_PRIVACY_UNSPECIFIED Privacy = 0
	Privacy_RESTRICTED          Privacy = 1
	Privacy_PUBLIC              Privacy = 2
)

// Enum value maps for Privacy.
var (
	Privacy_name = map[int32]string{
		0: "PRIVACY_UNSPECIFIED",
		1: "RESTRICTED",
		2: "PUBLIC",
	}
	Privacy_value = map[string]int32{
		"PRIVACY_UNSPECIFIED": 0,
		"RESTRICTED":          1,
		"PUBLIC":              2,
	}
)

func (x Privacy) Enum() *Privacy {
	p := new(Privacy)
	*p = x
	return p
}

func (x Privacy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Privacy) Descriptor() protoreflect.EnumDescriptor {
	return file_datatrails_common_api_assets_v2_assets_enums_proto_enumTypes[3].Descriptor()
}

func (Privacy) Type() protoreflect.EnumType {
	return &file_datatrails_common_api_assets_v2_assets_enums_proto_enumTypes[3]
}

func (x Privacy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Privacy.Descriptor instead.
func (Privacy) EnumDescriptor() ([]byte, []int) {
	return file_datatrails_common_api_assets_v2_assets_enums_proto_rawDescGZIP(), []int{3}
}

var File_datatrails_common_api_assets_v2_assets_enums_proto protoreflect.FileDescriptor

var file_datatrails_common_api_assets_v2_assets_enums_proto_rawDesc = []byte{
	0x0a, 0x32, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x73, 0x2d, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76,
	0x32, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x69, 0x73, 0x74, 0x2e,
	0x76, 0x32, 0x2a, 0x8d, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x1f, 0x43, 0x4f, 0x4e,
	0x46, 0x49, 0x52, 0x4d, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43,
	0x4f, 0x4e, 0x46, 0x49, 0x52, 0x4d, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41,
	0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x44,
	0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x54, 0x45, 0x44, 0x10,
	0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x45, 0x51, 0x55, 0x49, 0x56, 0x4f, 0x43, 0x41, 0x4c,
	0x10, 0x06, 0x2a, 0x56, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x1a, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x45, 0x44, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x01,
	0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x54, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x45, 0x44, 0x10,
	0x02, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4e, 0x59, 0x10, 0x03, 0x2a, 0x61, 0x0a, 0x0e, 0x50, 0x72,
	0x6f, 0x6f, 0x66, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x12, 0x1f, 0x0a, 0x1b,
	0x50, 0x52, 0x4f, 0x4f, 0x46, 0x5f, 0x4d, 0x45, 0x43, 0x48, 0x41, 0x4e, 0x49, 0x53, 0x4d, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a,
	0x09, 0x52, 0x45, 0x53, 0x45, 0x52, 0x56, 0x45, 0x44, 0x31, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b,
	0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x48, 0x41, 0x53, 0x48, 0x10, 0x02, 0x12, 0x0e, 0x0a,
	0x0a, 0x4d, 0x45, 0x52, 0x4b, 0x4c, 0x45, 0x5f, 0x4c, 0x4f, 0x47, 0x10, 0x03, 0x2a, 0x3e, 0x0a,
	0x07, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x52, 0x49, 0x56,
	0x41, 0x43, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x45, 0x53, 0x54, 0x52, 0x49, 0x43, 0x54, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x10, 0x02, 0x42, 0x4c, 0x5a,
	0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x74, 0x72, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72,
	0x61, 0x69, 0x6c, 0x73, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2d,
	0x67, 0x65, 0x6e, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x73, 0x3b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_datatrails_common_api_assets_v2_assets_enums_proto_rawDescOnce sync.Once
	file_datatrails_common_api_assets_v2_assets_enums_proto_rawDescData = file_datatrails_common_api_assets_v2_assets_enums_proto_rawDesc
)

func file_datatrails_common_api_assets_v2_assets_enums_proto_rawDescGZIP() []byte {
	file_datatrails_common_api_assets_v2_assets_enums_proto_rawDescOnce.Do(func() {
		file_datatrails_common_api_assets_v2_assets_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_datatrails_common_api_assets_v2_assets_enums_proto_rawDescData)
	})
	return file_datatrails_common_api_assets_v2_assets_enums_proto_rawDescData
}

var file_datatrails_common_api_assets_v2_assets_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_datatrails_common_api_assets_v2_assets_enums_proto_goTypes = []any{
	(ConfirmationStatus)(0), // 0: archivist.v2.ConfirmationStatus
	(TrackedStatus)(0),      // 1: archivist.v2.TrackedStatus
	(ProofMechanism)(0),     // 2: archivist.v2.ProofMechanism
	(Privacy)(0),            // 3: archivist.v2.Privacy
}
var file_datatrails_common_api_assets_v2_assets_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_datatrails_common_api_assets_v2_assets_enums_proto_init() }
func file_datatrails_common_api_assets_v2_assets_enums_proto_init() {
	if File_datatrails_common_api_assets_v2_assets_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_datatrails_common_api_assets_v2_assets_enums_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datatrails_common_api_assets_v2_assets_enums_proto_goTypes,
		DependencyIndexes: file_datatrails_common_api_assets_v2_assets_enums_proto_depIdxs,
		EnumInfos:         file_datatrails_common_api_assets_v2_assets_enums_proto_enumTypes,
	}.Build()
	File_datatrails_common_api_assets_v2_assets_enums_proto = out.File
	file_datatrails_common_api_assets_v2_assets_enums_proto_rawDesc = nil
	file_datatrails_common_api_assets_v2_assets_enums_proto_goTypes = nil
	file_datatrails_common_api_assets_v2_assets_enums_proto_depIdxs = nil
}
