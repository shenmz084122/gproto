// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: proto/service/apidashboard/apidashboard.proto

package pbsvcapidashboard

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

var File_proto_service_apidashboard_apidashboard_proto protoreflect.FileDescriptor

var file_proto_service_apidashboard_apidashboard_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x61, 0x70, 0x69,
	0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x61, 0x70, 0x69, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x1a, 0x26, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xec, 0x0c,
	0x0a, 0x0c, 0x41, 0x70, 0x69, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x3c,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x14, 0x2e,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x1a, 0x15, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0b,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x13, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x1a, 0x14, 0x2e, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x22, 0x00, 0x12, 0x39, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x12, 0x14, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x4b, 0x0a,
	0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x42, 0x79, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x42, 0x79, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x17, 0x2e, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22,
	0x00, 0x12, 0x3f, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x70, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x17, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x1a, 0x12, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x70, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x17, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x1a, 0x12, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x70, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x73, 0x12, 0x16, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x1a, 0x17, 0x2e, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x70, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x53, 0x4c, 0x12, 0x12, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x53, 0x4c, 0x1a, 0x13, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x53, 0x4c, 0x22, 0x00, 0x12,
	0x35, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x53, 0x4c, 0x12, 0x12, 0x2e, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x53, 0x4c,
	0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x53,
	0x4c, 0x73, 0x12, 0x11, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x53, 0x4c, 0x73, 0x1a, 0x12, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x53, 0x4c, 0x73, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x10, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x19, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1a, 0x2e, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x19, 0x2e, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x69,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x43, 0x0a,
	0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x19, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x12, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x22, 0x00, 0x12, 0x48, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a,
	0x19, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0e,
	0x41, 0x64, 0x64, 0x53, 0x76, 0x63, 0x52, 0x65, 0x71, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x17,
	0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x76, 0x63, 0x52,
	0x65, 0x71, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x45, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x53, 0x76, 0x63, 0x52, 0x65, 0x71, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x17, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x76, 0x63,
	0x52, 0x65, 0x71, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x76, 0x63, 0x52, 0x65, 0x71, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12,
	0x42, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79,
	0x12, 0x16, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x1a, 0x17, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65,
	0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74,
	0x68, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x1a, 0x12, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68,
	0x4b, 0x65, 0x79, 0x12, 0x16, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x1a, 0x12, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22,
	0x00, 0x12, 0x3f, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79,
	0x73, 0x12, 0x15, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x73, 0x1a, 0x16, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x73,
	0x22, 0x00, 0x12, 0x39, 0x0a, 0x0b, 0x42, 0x69, 0x6e, 0x64, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65,
	0x79, 0x12, 0x14, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x69, 0x6e, 0x64,
	0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x3d, 0x0a,
	0x0d, 0x55, 0x6e, 0x62, 0x69, 0x6e, 0x64, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x16,
	0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x6e, 0x62, 0x69, 0x6e, 0x64, 0x41,
	0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x42, 0x84, 0x01, 0x0a,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6f, 0x6d, 0x6e, 0x69, 0x73, 0x2e, 0x67,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x62,
	0x73, 0x76, 0x63, 0x61, 0x70, 0x69, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x42,
	0x11, 0x50, 0x42, 0x53, 0x76, 0x63, 0x41, 0x70, 0x69, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x50, 0x00, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x44, 0x61, 0x74, 0x61, 0x57, 0x6f, 0x72, 0x6b, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2f, 0x67,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x78, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x70, 0x62, 0x73, 0x76, 0x63, 0x61, 0x70, 0x69, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_service_apidashboard_apidashboard_proto_goTypes = []interface{}{
	(*pbrequest.CreateRoute)(nil),          // 0: request.CreateRoute
	(*pbrequest.DeleteRoute)(nil),          // 1: request.DeleteRoute
	(*pbrequest.ListRoutes)(nil),           // 2: request.ListRoutes
	(*pbrequest.UpdateRoute)(nil),          // 3: request.UpdateRoute
	(*pbrequest.DeleteRouteByVersion)(nil), // 4: request.DeleteRouteByVersion
	(*pbrequest.CreateUpstream)(nil),       // 5: request.CreateUpstream
	(*pbrequest.DeleteUpstream)(nil),       // 6: request.DeleteUpstream
	(*pbrequest.UpdateUpstream)(nil),       // 7: request.UpdateUpstream
	(*pbrequest.ListUpstreams)(nil),        // 8: request.ListUpstreams
	(*pbrequest.CreateSSL)(nil),            // 9: request.CreateSSL
	(*pbrequest.DeleteSSL)(nil),            // 10: request.DeleteSSL
	(*pbrequest.ListSSLs)(nil),             // 11: request.ListSSLs
	(*pbrequest.CreateApiService)(nil),     // 12: request.CreateApiService
	(*pbrequest.DeleteApiService)(nil),     // 13: request.DeleteApiService
	(*pbrequest.UpdateApiService)(nil),     // 14: request.UpdateApiService
	(*pbrequest.ListApiServices)(nil),      // 15: request.ListApiServices
	(*pbrequest.AddSvcReqCount)(nil),       // 16: request.AddSvcReqCount
	(*pbrequest.GetSvcReqCount)(nil),       // 17: request.GetSvcReqCount
	(*pbrequest.DeleteProjectRoutes)(nil),  // 18: request.DeleteProjectRoutes
	(*pbrequest.CreateAuthKey)(nil),        // 19: request.CreateAuthKey
	(*pbrequest.DeleteAuthKey)(nil),        // 20: request.DeleteAuthKey
	(*pbrequest.UpdateAuthKey)(nil),        // 21: request.UpdateAuthKey
	(*pbrequest.ListAuthKeys)(nil),         // 22: request.ListAuthKeys
	(*pbrequest.BindAuthKey)(nil),          // 23: request.BindAuthKey
	(*pbrequest.UnbindAuthKey)(nil),        // 24: request.UnbindAuthKey
	(*pbresponse.CreateRoute)(nil),         // 25: response.CreateRoute
	(*pbmodel.EmptyStruct)(nil),            // 26: model.EmptyStruct
	(*pbresponse.ListRoutes)(nil),          // 27: response.ListRoutes
	(*pbresponse.CreateUpstream)(nil),      // 28: response.CreateUpstream
	(*pbresponse.ListUpstreams)(nil),       // 29: response.ListUpstreams
	(*pbresponse.CreateSSL)(nil),           // 30: response.CreateSSL
	(*pbresponse.ListSSLs)(nil),            // 31: response.ListSSLs
	(*pbresponse.CreateApiService)(nil),    // 32: response.CreateApiService
	(*pbresponse.ListApiServices)(nil),     // 33: response.ListApiServices
	(*pbresponse.GetSvcReqCount)(nil),      // 34: response.GetSvcReqCount
	(*pbresponse.CreateAuthKey)(nil),       // 35: response.CreateAuthKey
	(*pbresponse.ListAuthKeys)(nil),        // 36: response.ListAuthKeys
}
var file_proto_service_apidashboard_apidashboard_proto_depIdxs = []int32{
	0,  // 0: apidashboard.ApiDashboard.CreateRoute:input_type -> request.CreateRoute
	1,  // 1: apidashboard.ApiDashboard.DeleteRoute:input_type -> request.DeleteRoute
	2,  // 2: apidashboard.ApiDashboard.ListRoutes:input_type -> request.ListRoutes
	3,  // 3: apidashboard.ApiDashboard.UpdateRoute:input_type -> request.UpdateRoute
	4,  // 4: apidashboard.ApiDashboard.DeleteRouteByVersion:input_type -> request.DeleteRouteByVersion
	5,  // 5: apidashboard.ApiDashboard.CreateUpstream:input_type -> request.CreateUpstream
	6,  // 6: apidashboard.ApiDashboard.DeleteUpstream:input_type -> request.DeleteUpstream
	7,  // 7: apidashboard.ApiDashboard.UpdateUpstream:input_type -> request.UpdateUpstream
	8,  // 8: apidashboard.ApiDashboard.ListUpstreams:input_type -> request.ListUpstreams
	9,  // 9: apidashboard.ApiDashboard.CreateSSL:input_type -> request.CreateSSL
	10, // 10: apidashboard.ApiDashboard.DeleteSSL:input_type -> request.DeleteSSL
	11, // 11: apidashboard.ApiDashboard.ListSSLs:input_type -> request.ListSSLs
	12, // 12: apidashboard.ApiDashboard.CreateApiService:input_type -> request.CreateApiService
	13, // 13: apidashboard.ApiDashboard.DeleteApiService:input_type -> request.DeleteApiService
	14, // 14: apidashboard.ApiDashboard.UpdateApiService:input_type -> request.UpdateApiService
	15, // 15: apidashboard.ApiDashboard.ListApiServices:input_type -> request.ListApiServices
	16, // 16: apidashboard.ApiDashboard.AddSvcReqCount:input_type -> request.AddSvcReqCount
	17, // 17: apidashboard.ApiDashboard.GetSvcReqCount:input_type -> request.GetSvcReqCount
	18, // 18: apidashboard.ApiDashboard.DeleteProjectRoutes:input_type -> request.DeleteProjectRoutes
	19, // 19: apidashboard.ApiDashboard.CreateAuthKey:input_type -> request.CreateAuthKey
	20, // 20: apidashboard.ApiDashboard.DeleteAuthKey:input_type -> request.DeleteAuthKey
	21, // 21: apidashboard.ApiDashboard.UpdateAuthKey:input_type -> request.UpdateAuthKey
	22, // 22: apidashboard.ApiDashboard.ListAuthKeys:input_type -> request.ListAuthKeys
	23, // 23: apidashboard.ApiDashboard.BindAuthKey:input_type -> request.BindAuthKey
	24, // 24: apidashboard.ApiDashboard.UnbindAuthKey:input_type -> request.UnbindAuthKey
	25, // 25: apidashboard.ApiDashboard.CreateRoute:output_type -> response.CreateRoute
	26, // 26: apidashboard.ApiDashboard.DeleteRoute:output_type -> model.EmptyStruct
	27, // 27: apidashboard.ApiDashboard.ListRoutes:output_type -> response.ListRoutes
	26, // 28: apidashboard.ApiDashboard.UpdateRoute:output_type -> model.EmptyStruct
	26, // 29: apidashboard.ApiDashboard.DeleteRouteByVersion:output_type -> model.EmptyStruct
	28, // 30: apidashboard.ApiDashboard.CreateUpstream:output_type -> response.CreateUpstream
	26, // 31: apidashboard.ApiDashboard.DeleteUpstream:output_type -> model.EmptyStruct
	26, // 32: apidashboard.ApiDashboard.UpdateUpstream:output_type -> model.EmptyStruct
	29, // 33: apidashboard.ApiDashboard.ListUpstreams:output_type -> response.ListUpstreams
	30, // 34: apidashboard.ApiDashboard.CreateSSL:output_type -> response.CreateSSL
	26, // 35: apidashboard.ApiDashboard.DeleteSSL:output_type -> model.EmptyStruct
	31, // 36: apidashboard.ApiDashboard.ListSSLs:output_type -> response.ListSSLs
	32, // 37: apidashboard.ApiDashboard.CreateApiService:output_type -> response.CreateApiService
	26, // 38: apidashboard.ApiDashboard.DeleteApiService:output_type -> model.EmptyStruct
	26, // 39: apidashboard.ApiDashboard.UpdateApiService:output_type -> model.EmptyStruct
	33, // 40: apidashboard.ApiDashboard.ListApiServices:output_type -> response.ListApiServices
	26, // 41: apidashboard.ApiDashboard.AddSvcReqCount:output_type -> model.EmptyStruct
	34, // 42: apidashboard.ApiDashboard.GetSvcReqCount:output_type -> response.GetSvcReqCount
	26, // 43: apidashboard.ApiDashboard.DeleteProjectRoutes:output_type -> model.EmptyStruct
	35, // 44: apidashboard.ApiDashboard.CreateAuthKey:output_type -> response.CreateAuthKey
	26, // 45: apidashboard.ApiDashboard.DeleteAuthKey:output_type -> model.EmptyStruct
	26, // 46: apidashboard.ApiDashboard.UpdateAuthKey:output_type -> model.EmptyStruct
	36, // 47: apidashboard.ApiDashboard.ListAuthKeys:output_type -> response.ListAuthKeys
	26, // 48: apidashboard.ApiDashboard.BindAuthKey:output_type -> model.EmptyStruct
	26, // 49: apidashboard.ApiDashboard.UnbindAuthKey:output_type -> model.EmptyStruct
	25, // [25:50] is the sub-list for method output_type
	0,  // [0:25] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_service_apidashboard_apidashboard_proto_init() }
func file_proto_service_apidashboard_apidashboard_proto_init() {
	if File_proto_service_apidashboard_apidashboard_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_service_apidashboard_apidashboard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service_apidashboard_apidashboard_proto_goTypes,
		DependencyIndexes: file_proto_service_apidashboard_apidashboard_proto_depIdxs,
	}.Build()
	File_proto_service_apidashboard_apidashboard_proto = out.File
	file_proto_service_apidashboard_apidashboard_proto_rawDesc = nil
	file_proto_service_apidashboard_apidashboard_proto_goTypes = nil
	file_proto_service_apidashboard_apidashboard_proto_depIdxs = nil
}
