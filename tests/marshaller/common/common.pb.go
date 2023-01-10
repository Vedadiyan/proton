// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: common/proton.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_common_proton_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         51234,
		Name:          "proton.select",
		Tag:           "bytes,51234,opt,name=select",
		Filename:      "common/proton.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         51235,
		Name:          "proton.query",
		Tag:           "bytes,51235,opt,name=query",
		Filename:      "common/proton.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         51236,
		Name:          "proton.not_mapped",
		Tag:           "bytes,51236,opt,name=not_mapped",
		Filename:      "common/proton.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         51237,
		Name:          "proton.bind_to",
		Tag:           "bytes,51237,opt,name=bind_to",
		Filename:      "common/proton.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         51238,
		Name:          "proton.type",
		Tag:           "bytes,51238,opt,name=type",
		Filename:      "common/proton.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         51239,
		Name:          "proton.index",
		Tag:           "bytes,51239,opt,name=index",
		Filename:      "common/proton.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional string select = 51234;
	E_Select = &file_common_proton_proto_extTypes[0]
	// optional string query = 51235;
	E_Query = &file_common_proton_proto_extTypes[1]
	// optional string not_mapped = 51236;
	E_NotMapped = &file_common_proton_proto_extTypes[2]
	// optional string bind_to = 51237;
	E_BindTo = &file_common_proton_proto_extTypes[3]
	// optional string type = 51238;
	E_Type = &file_common_proton_proto_extTypes[4]
	// optional string index = 51239;
	E_Index = &file_common_proton_proto_extTypes[5]
)

var File_common_proton_proto protoreflect.FileDescriptor

var file_common_proton_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a,
	0x3a, 0x0a, 0x06, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa2, 0x90, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x88, 0x01, 0x01, 0x3a, 0x38, 0x0a, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xa3, 0x90, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x88, 0x01, 0x01, 0x3a, 0x41, 0x0a, 0x0a, 0x6e, 0x6f, 0x74, 0x5f, 0x6d, 0x61, 0x70,
	0x70, 0x65, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xa4, 0x90, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x6f, 0x74, 0x4d,
	0x61, 0x70, 0x70, 0x65, 0x64, 0x88, 0x01, 0x01, 0x3a, 0x3b, 0x0a, 0x07, 0x62, 0x69, 0x6e, 0x64,
	0x5f, 0x74, 0x6f, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xa5, 0x90, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x69, 0x6e, 0x64,
	0x54, 0x6f, 0x88, 0x01, 0x01, 0x3a, 0x36, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa6, 0x90, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x3a, 0x38, 0x0a,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa7, 0x90, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x88, 0x01, 0x01, 0x42, 0x10, 0x5a, 0x0e, 0x61, 0x75, 0x74, 0x6f, 0x67,
	0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_common_proton_proto_goTypes = []interface{}{
	(*descriptorpb.FieldOptions)(nil), // 0: google.protobuf.FieldOptions
}
var file_common_proton_proto_depIdxs = []int32{
	0, // 0: proton.select:extendee -> google.protobuf.FieldOptions
	0, // 1: proton.query:extendee -> google.protobuf.FieldOptions
	0, // 2: proton.not_mapped:extendee -> google.protobuf.FieldOptions
	0, // 3: proton.bind_to:extendee -> google.protobuf.FieldOptions
	0, // 4: proton.type:extendee -> google.protobuf.FieldOptions
	0, // 5: proton.index:extendee -> google.protobuf.FieldOptions
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	0, // [0:6] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_proton_proto_init() }
func file_common_proton_proto_init() {
	if File_common_proton_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_proton_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 6,
			NumServices:   0,
		},
		GoTypes:           file_common_proton_proto_goTypes,
		DependencyIndexes: file_common_proton_proto_depIdxs,
		ExtensionInfos:    file_common_proton_proto_extTypes,
	}.Build()
	File_common_proton_proto = out.File
	file_common_proton_proto_rawDesc = nil
	file_common_proton_proto_goTypes = nil
	file_common_proton_proto_depIdxs = nil
}
