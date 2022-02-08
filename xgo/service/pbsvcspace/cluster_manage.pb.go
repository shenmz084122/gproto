// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: proto/service/spacemanager/cluster_manage.proto

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

var File_proto_service_spacemanager_cluster_manage_proto protoreflect.FileDescriptor

var file_proto_service_spacemanager_cluster_manage_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x1a,
	0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xeb, 0x05, 0x0a, 0x0d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x12, 0x58, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x24, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x00, 0x12,
	0x60, 0x0a, 0x17, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x46, 0x6c, 0x69, 0x6e, 0x6b,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x50, 0x49, 0x12, 0x20, 0x2e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x46, 0x6c, 0x69,
	0x6e, 0x6b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x50, 0x49, 0x1a, 0x21, 0x2e, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x50, 0x49, 0x22,
	0x00, 0x12, 0x4e, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x1a, 0x1b, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x22,
	0x00, 0x12, 0x51, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x6e, 0x6b,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x1a, 0x1c, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x46,
	0x6c, 0x69, 0x6e, 0x6b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x1e, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x46,
	0x6c, 0x69, 0x6e, 0x6b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x00, 0x12, 0x47, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x2e,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6c,
	0x69, 0x6e, 0x6b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x12, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22,
	0x00, 0x12, 0x47, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x72, 0x74, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1b, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x11, 0x53, 0x74,
	0x6f, 0x70, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x1a, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x46, 0x6c,
	0x69, 0x6e, 0x6b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x12, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22,
	0x00, 0x42, 0x77, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6f, 0x6d, 0x6e,
	0x69, 0x73, 0x2e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x62, 0x73, 0x76, 0x63, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x12, 0x50, 0x42,
	0x53, 0x76, 0x63, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x50, 0x00, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44,
	0x61, 0x74, 0x61, 0x57, 0x6f, 0x72, 0x6b, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2f, 0x67, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x78, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x70, 0x62, 0x73, 0x76, 0x63, 0x73, 0x70, 0x61, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_proto_service_spacemanager_cluster_manage_proto_goTypes = []interface{}{
	(*pbmodel.EmptyStruct)(nil),                   // 0: model.EmptyStruct
	(*pbrequest.DescribeFlinkClusterAPI)(nil),     // 1: request.DescribeFlinkClusterAPI
	(*pbrequest.ListFlinkClusters)(nil),           // 2: request.ListFlinkClusters
	(*pbrequest.CreateFlinkCluster)(nil),          // 3: request.CreateFlinkCluster
	(*pbrequest.DescribeFlinkCluster)(nil),        // 4: request.DescribeFlinkCluster
	(*pbrequest.UpdateFlinkCluster)(nil),          // 5: request.UpdateFlinkCluster
	(*pbrequest.DeleteFlinkClusters)(nil),         // 6: request.DeleteFlinkClusters
	(*pbrequest.StartFlinkClusters)(nil),          // 7: request.StartFlinkClusters
	(*pbrequest.StopFlinkClusters)(nil),           // 8: request.StopFlinkClusters
	(*pbresponse.ListAvailableFlinkVersions)(nil), // 9: response.ListAvailableFlinkVersions
	(*pbresponse.DescribeFlinkClusterAPI)(nil),    // 10: response.DescribeFlinkClusterAPI
	(*pbresponse.ListFlinkClusters)(nil),          // 11: response.ListFlinkClusters
	(*pbresponse.CreateFlinkCluster)(nil),         // 12: response.CreateFlinkCluster
	(*pbresponse.DescribeFlinkCluster)(nil),       // 13: response.DescribeFlinkCluster
}
var file_proto_service_spacemanager_cluster_manage_proto_depIdxs = []int32{
	0,  // 0: spacemanager.ClusterManage.ListAvailableFlinkVersions:input_type -> model.EmptyStruct
	1,  // 1: spacemanager.ClusterManage.DescribeFlinkClusterAPI:input_type -> request.DescribeFlinkClusterAPI
	2,  // 2: spacemanager.ClusterManage.ListFlinkClusters:input_type -> request.ListFlinkClusters
	3,  // 3: spacemanager.ClusterManage.CreateFlinkCluster:input_type -> request.CreateFlinkCluster
	4,  // 4: spacemanager.ClusterManage.DescribeFlinkCluster:input_type -> request.DescribeFlinkCluster
	5,  // 5: spacemanager.ClusterManage.UpdateFlinkCluster:input_type -> request.UpdateFlinkCluster
	6,  // 6: spacemanager.ClusterManage.DeleteFlinkClusters:input_type -> request.DeleteFlinkClusters
	7,  // 7: spacemanager.ClusterManage.StartFlinkClusters:input_type -> request.StartFlinkClusters
	8,  // 8: spacemanager.ClusterManage.StopFlinkClusters:input_type -> request.StopFlinkClusters
	9,  // 9: spacemanager.ClusterManage.ListAvailableFlinkVersions:output_type -> response.ListAvailableFlinkVersions
	10, // 10: spacemanager.ClusterManage.DescribeFlinkClusterAPI:output_type -> response.DescribeFlinkClusterAPI
	11, // 11: spacemanager.ClusterManage.ListFlinkClusters:output_type -> response.ListFlinkClusters
	12, // 12: spacemanager.ClusterManage.CreateFlinkCluster:output_type -> response.CreateFlinkCluster
	13, // 13: spacemanager.ClusterManage.DescribeFlinkCluster:output_type -> response.DescribeFlinkCluster
	0,  // 14: spacemanager.ClusterManage.UpdateFlinkCluster:output_type -> model.EmptyStruct
	0,  // 15: spacemanager.ClusterManage.DeleteFlinkClusters:output_type -> model.EmptyStruct
	0,  // 16: spacemanager.ClusterManage.StartFlinkClusters:output_type -> model.EmptyStruct
	0,  // 17: spacemanager.ClusterManage.StopFlinkClusters:output_type -> model.EmptyStruct
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_service_spacemanager_cluster_manage_proto_init() }
func file_proto_service_spacemanager_cluster_manage_proto_init() {
	if File_proto_service_spacemanager_cluster_manage_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_service_spacemanager_cluster_manage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service_spacemanager_cluster_manage_proto_goTypes,
		DependencyIndexes: file_proto_service_spacemanager_cluster_manage_proto_depIdxs,
	}.Build()
	File_proto_service_spacemanager_cluster_manage_proto = out.File
	file_proto_service_spacemanager_cluster_manage_proto_rawDesc = nil
	file_proto_service_spacemanager_cluster_manage_proto_goTypes = nil
	file_proto_service_spacemanager_cluster_manage_proto_depIdxs = nil
}
