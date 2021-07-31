// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.5.1
// source: api/helloworld/v1/error_reason.proto

package v1

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

type ErrorReason int32

const (
	ErrorReason_ERROR_REASON_UNSPECIFIED ErrorReason = 0
	ErrorReason_USER_NOT_FOUND           ErrorReason = 1
	ErrorReason_SERVICE_DISABLED         ErrorReason = 2
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0: "ERROR_REASON_UNSPECIFIED",
		1: "USER_NOT_FOUND",
		2: "SERVICE_DISABLED",
	}
	ErrorReason_value = map[string]int32{
		"ERROR_REASON_UNSPECIFIED": 0,
		"USER_NOT_FOUND":           1,
		"SERVICE_DISABLED":         2,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_api_helloworld_v1_error_reason_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_api_helloworld_v1_error_reason_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_api_helloworld_v1_error_reason_proto_rawDescGZIP(), []int{0}
}

var File_api_helloworld_v1_error_reason_proto protoreflect.FileDescriptor

var file_api_helloworld_v1_error_reason_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2a, 0x55, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45,
	0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46,
	0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43,
	0x45, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x42, 0x60, 0x0a, 0x23,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x50, 0x01, 0x5a, 0x1e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0xa2, 0x02, 0x16, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_helloworld_v1_error_reason_proto_rawDescOnce sync.Once
	file_api_helloworld_v1_error_reason_proto_rawDescData = file_api_helloworld_v1_error_reason_proto_rawDesc
)

func file_api_helloworld_v1_error_reason_proto_rawDescGZIP() []byte {
	file_api_helloworld_v1_error_reason_proto_rawDescOnce.Do(func() {
		file_api_helloworld_v1_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_helloworld_v1_error_reason_proto_rawDescData)
	})
	return file_api_helloworld_v1_error_reason_proto_rawDescData
}

var file_api_helloworld_v1_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_helloworld_v1_error_reason_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: helloworld.v1.ErrorReason
}
var file_api_helloworld_v1_error_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_helloworld_v1_error_reason_proto_init() }
func file_api_helloworld_v1_error_reason_proto_init() {
	if File_api_helloworld_v1_error_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_helloworld_v1_error_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_helloworld_v1_error_reason_proto_goTypes,
		DependencyIndexes: file_api_helloworld_v1_error_reason_proto_depIdxs,
		EnumInfos:         file_api_helloworld_v1_error_reason_proto_enumTypes,
	}.Build()
	File_api_helloworld_v1_error_reason_proto = out.File
	file_api_helloworld_v1_error_reason_proto_rawDesc = nil
	file_api_helloworld_v1_error_reason_proto_goTypes = nil
	file_api_helloworld_v1_error_reason_proto_depIdxs = nil
}
