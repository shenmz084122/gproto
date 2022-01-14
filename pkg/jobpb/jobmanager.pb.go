// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/jobmanager.proto

package jobpb

import (
	model "github.com/DataWorkbench/gproto/pkg/model"
	request "github.com/DataWorkbench/gproto/pkg/request"
	response "github.com/DataWorkbench/gproto/pkg/response"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_proto_jobmanager_proto protoreflect.FileDescriptor

var file_proto_jobmanager_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a, 0x6f, 0x62, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6a, 0x6f, 0x62, 0x70, 0x62, 0x1a,
	0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa0, 0x03, 0x0a, 0x0a,
	0x4a, 0x6f, 0x62, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0c, 0x46, 0x72,
	0x65, 0x65, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x4a, 0x6f, 0x62, 0x12, 0x15, 0x2e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x72, 0x65, 0x65, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x4a, 0x6f,
	0x62, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0c, 0x49, 0x6e, 0x69, 0x74, 0x46,
	0x6c, 0x69, 0x6e, 0x6b, 0x4a, 0x6f, 0x62, 0x12, 0x15, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x4a, 0x6f, 0x62, 0x1a, 0x16,
	0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x46, 0x6c,
	0x69, 0x6e, 0x6b, 0x4a, 0x6f, 0x62, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x4a, 0x6f, 0x62, 0x12, 0x17, 0x2e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x46, 0x6c, 0x69, 0x6e, 0x6b,
	0x4a, 0x6f, 0x62, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x4a, 0x6f, 0x62, 0x22, 0x00, 0x12,
	0x3c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x4a, 0x6f, 0x62, 0x12, 0x14,
	0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x69, 0x6e,
	0x6b, 0x4a, 0x6f, 0x62, 0x1a, 0x15, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x4a, 0x6f, 0x62, 0x22, 0x00, 0x12, 0x3f, 0x0a,
	0x0e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x4a, 0x6f, 0x62, 0x12,
	0x17, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x4a, 0x6f, 0x62, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x4e,
	0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x4a,
	0x6f, 0x62, 0x12, 0x19, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x4a, 0x6f, 0x62, 0x1a, 0x1d, 0x2e,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4a,
	0x6f, 0x62, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x22, 0x00, 0x42, 0x4f,
	0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x6f, 0x62, 0x70, 0x62, 0x42, 0x05, 0x4a, 0x6f, 0x62,
	0x70, 0x62, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x44, 0x61, 0x74, 0x61, 0x57, 0x6f, 0x72, 0x6b, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2f, 0x67,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6a, 0x6f, 0x62, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_jobmanager_proto_goTypes = []interface{}{
	(*request.FreeFlinkJob)(nil),         // 0: request.FreeFlinkJob
	(*request.InitFlinkJob)(nil),         // 1: request.InitFlinkJob
	(*request.SubmitFlinkJob)(nil),       // 2: request.SubmitFlinkJob
	(*request.GetFlinkJob)(nil),          // 3: request.GetFlinkJob
	(*request.CancelFlinkJob)(nil),       // 4: request.CancelFlinkJob
	(*request.ValidateFlinkJob)(nil),     // 5: request.ValidateFlinkJob
	(*model.EmptyStruct)(nil),            // 6: model.EmptyStruct
	(*response.InitFlinkJob)(nil),        // 7: response.InitFlinkJob
	(*response.SubmitFlinkJob)(nil),      // 8: response.SubmitFlinkJob
	(*response.GetFlinkJob)(nil),         // 9: response.GetFlinkJob
	(*response.StreamJobCodeSyntax)(nil), // 10: response.StreamJobCodeSyntax
}
var file_proto_jobmanager_proto_depIdxs = []int32{
	0,  // 0: jobpb.Jobmanager.FreeFlinkJob:input_type -> request.FreeFlinkJob
	1,  // 1: jobpb.Jobmanager.InitFlinkJob:input_type -> request.InitFlinkJob
	2,  // 2: jobpb.Jobmanager.SubmitFlinkJob:input_type -> request.SubmitFlinkJob
	3,  // 3: jobpb.Jobmanager.GetFlinkJob:input_type -> request.GetFlinkJob
	4,  // 4: jobpb.Jobmanager.CancelFlinkJob:input_type -> request.CancelFlinkJob
	5,  // 5: jobpb.Jobmanager.ValidateFlinkJob:input_type -> request.ValidateFlinkJob
	6,  // 6: jobpb.Jobmanager.FreeFlinkJob:output_type -> model.EmptyStruct
	7,  // 7: jobpb.Jobmanager.InitFlinkJob:output_type -> response.InitFlinkJob
	8,  // 8: jobpb.Jobmanager.SubmitFlinkJob:output_type -> response.SubmitFlinkJob
	9,  // 9: jobpb.Jobmanager.GetFlinkJob:output_type -> response.GetFlinkJob
	6,  // 10: jobpb.Jobmanager.CancelFlinkJob:output_type -> model.EmptyStruct
	10, // 11: jobpb.Jobmanager.ValidateFlinkJob:output_type -> response.StreamJobCodeSyntax
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_jobmanager_proto_init() }
func file_proto_jobmanager_proto_init() {
	if File_proto_jobmanager_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_jobmanager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_jobmanager_proto_goTypes,
		DependencyIndexes: file_proto_jobmanager_proto_depIdxs,
	}.Build()
	File_proto_jobmanager_proto = out.File
	file_proto_jobmanager_proto_rawDesc = nil
	file_proto_jobmanager_proto_goTypes = nil
	file_proto_jobmanager_proto_depIdxs = nil
}
