// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: proto/service/spacemanager/udf_manage.proto

package pbsvcspace

import (
	pbmodel "github.com/DataWorkbench/gproto/xgo/types/pbmodel"
	pbrequest "github.com/DataWorkbench/gproto/xgo/types/pbrequest"
	pbresponse "github.com/DataWorkbench/gproto/xgo/types/pbresponse"
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

var File_proto_service_spacemanager_udf_manage_proto protoreflect.FileDescriptor

var file_proto_service_spacemanager_udf_manage_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x75, 0x64, 0x66,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x1a, 0x1d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f,
	0x75, 0x64, 0x66, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x75, 0x64, 0x66, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa6, 0x02, 0x0a, 0x09, 0x55, 0x44, 0x46, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x44, 0x46,
	0x73, 0x12, 0x11, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x44, 0x46, 0x73, 0x1a, 0x12, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x55, 0x44, 0x46, 0x73, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x44, 0x46, 0x73, 0x12, 0x13, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x44, 0x46, 0x73, 0x1a, 0x12, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x44, 0x46,
	0x12, 0x12, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x44, 0x46, 0x1a, 0x13, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x44, 0x46, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x44, 0x46, 0x12, 0x12, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x44, 0x46, 0x1a, 0x12, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x55, 0x44,
	0x46, 0x12, 0x14, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x55, 0x44, 0x46, 0x1a, 0x15, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x55, 0x44, 0x46, 0x22, 0x00,
	0x42, 0x73, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6f, 0x6d, 0x6e, 0x69,
	0x73, 0x2e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x62, 0x73, 0x76, 0x63, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x0e, 0x50, 0x42, 0x53,
	0x76, 0x63, 0x55, 0x44, 0x46, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x50, 0x00, 0x5a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x57, 0x6f,
	0x72, 0x6b, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2f, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x78,
	0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x73, 0x76, 0x63,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_service_spacemanager_udf_manage_proto_goTypes = []interface{}{
	(*pbrequest.ListUDFs)(nil),     // 0: request.ListUDFs
	(*pbrequest.DeleteUDFs)(nil),   // 1: request.DeleteUDFs
	(*pbrequest.CreateUDF)(nil),    // 2: request.CreateUDF
	(*pbrequest.UpdateUDF)(nil),    // 3: request.UpdateUDF
	(*pbrequest.DescribeUDF)(nil),  // 4: request.DescribeUDF
	(*pbresponse.ListUDFs)(nil),    // 5: response.ListUDFs
	(*pbmodel.EmptyStruct)(nil),    // 6: model.EmptyStruct
	(*pbresponse.CreateUDF)(nil),   // 7: response.CreateUDF
	(*pbresponse.DescribeUDF)(nil), // 8: response.DescribeUDF
}
var file_proto_service_spacemanager_udf_manage_proto_depIdxs = []int32{
	0, // 0: spacemanager.UDFManage.ListUDFs:input_type -> request.ListUDFs
	1, // 1: spacemanager.UDFManage.DeleteUDFs:input_type -> request.DeleteUDFs
	2, // 2: spacemanager.UDFManage.CreateUDF:input_type -> request.CreateUDF
	3, // 3: spacemanager.UDFManage.UpdateUDF:input_type -> request.UpdateUDF
	4, // 4: spacemanager.UDFManage.DescribeUDF:input_type -> request.DescribeUDF
	5, // 5: spacemanager.UDFManage.ListUDFs:output_type -> response.ListUDFs
	6, // 6: spacemanager.UDFManage.DeleteUDFs:output_type -> model.EmptyStruct
	7, // 7: spacemanager.UDFManage.CreateUDF:output_type -> response.CreateUDF
	6, // 8: spacemanager.UDFManage.UpdateUDF:output_type -> model.EmptyStruct
	8, // 9: spacemanager.UDFManage.DescribeUDF:output_type -> response.DescribeUDF
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_service_spacemanager_udf_manage_proto_init() }
func file_proto_service_spacemanager_udf_manage_proto_init() {
	if File_proto_service_spacemanager_udf_manage_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_service_spacemanager_udf_manage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service_spacemanager_udf_manage_proto_goTypes,
		DependencyIndexes: file_proto_service_spacemanager_udf_manage_proto_depIdxs,
	}.Build()
	File_proto_service_spacemanager_udf_manage_proto = out.File
	file_proto_service_spacemanager_udf_manage_proto_rawDesc = nil
	file_proto_service_spacemanager_udf_manage_proto_goTypes = nil
	file_proto_service_spacemanager_udf_manage_proto_depIdxs = nil
}
