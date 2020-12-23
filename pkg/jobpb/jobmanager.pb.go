// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/jobmanager.proto

package jobpb

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type JobReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *JobReply) Reset() {
	*x = JobReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_jobmanager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobReply) ProtoMessage() {}

func (x *JobReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_jobmanager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobReply.ProtoReflect.Descriptor instead.
func (*JobReply) Descriptor() ([]byte, []int) {
	return file_proto_jobmanager_proto_rawDescGZIP(), []int{0}
}

func (x *JobReply) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *JobReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// EmptyReply represents no value returned
type EmptyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyReply) Reset() {
	*x = EmptyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_jobmanager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyReply) ProtoMessage() {}

func (x *EmptyReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_jobmanager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyReply.ProtoReflect.Descriptor instead.
func (*EmptyReply) Descriptor() ([]byte, []int) {
	return file_proto_jobmanager_proto_rawDescGZIP(), []int{1}
}

type RunJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	WorkspaceID string `protobuf:"bytes,2,opt,name=WorkspaceID,proto3" json:"WorkspaceID,omitempty"`
	NodeType    string `protobuf:"bytes,3,opt,name=NodeType,proto3" json:"NodeType,omitempty"`
	Depends     string `protobuf:"bytes,4,opt,name=Depends,proto3" json:"Depends,omitempty"`
	MainRun     string `protobuf:"bytes,5,opt,name=MainRun,proto3" json:"MainRun,omitempty"`
}

func (x *RunJobRequest) Reset() {
	*x = RunJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_jobmanager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunJobRequest) ProtoMessage() {}

func (x *RunJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_jobmanager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunJobRequest.ProtoReflect.Descriptor instead.
func (*RunJobRequest) Descriptor() ([]byte, []int) {
	return file_proto_jobmanager_proto_rawDescGZIP(), []int{2}
}

func (x *RunJobRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *RunJobRequest) GetWorkspaceID() string {
	if x != nil {
		return x.WorkspaceID
	}
	return ""
}

func (x *RunJobRequest) GetNodeType() string {
	if x != nil {
		return x.NodeType
	}
	return ""
}

func (x *RunJobRequest) GetDepends() string {
	if x != nil {
		return x.Depends
	}
	return ""
}

func (x *RunJobRequest) GetMainRun() string {
	if x != nil {
		return x.MainRun
	}
	return ""
}

type GetJobStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetJobStatusRequest) Reset() {
	*x = GetJobStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_jobmanager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJobStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobStatusRequest) ProtoMessage() {}

func (x *GetJobStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_jobmanager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobStatusRequest.ProtoReflect.Descriptor instead.
func (*GetJobStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_jobmanager_proto_rawDescGZIP(), []int{3}
}

func (x *GetJobStatusRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type CancelJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *CancelJobRequest) Reset() {
	*x = CancelJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_jobmanager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelJobRequest) ProtoMessage() {}

func (x *CancelJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_jobmanager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelJobRequest.ProtoReflect.Descriptor instead.
func (*CancelJobRequest) Descriptor() ([]byte, []int) {
	return file_proto_jobmanager_proto_rawDescGZIP(), []int{4}
}

func (x *CancelJobRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

var File_proto_jobmanager_proto protoreflect.FileDescriptor

var file_proto_jobmanager_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a, 0x6f, 0x62, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6a, 0x6f, 0x62, 0x70, 0x62, 0x1a,
	0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74,
	0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x08, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x20, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x08, 0xe2, 0xdf, 0x1f, 0x04, 0x70, 0x00, 0x78, 0x15, 0x52, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x78, 0xd0, 0x0f, 0x52,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x0c, 0x0a, 0x0a, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0xbf, 0x01, 0x0a, 0x0d, 0x52, 0x75, 0x6e, 0x4a, 0x6f,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01, 0x14, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x29, 0x0a, 0x0b, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01, 0x14, 0x52,
	0x0b, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x08,
	0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xe2, 0xdf, 0x1f, 0x04, 0x70, 0x00, 0x78, 0x15, 0x52, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x21, 0x0a, 0x07, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x78, 0xa0, 0x1f, 0x52, 0x07, 0x44, 0x65,
	0x70, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x21, 0x0a, 0x07, 0x4d, 0x61, 0x69, 0x6e, 0x52, 0x75, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x78, 0xc0, 0x3e, 0x52,
	0x07, 0x4d, 0x61, 0x69, 0x6e, 0x52, 0x75, 0x6e, 0x22, 0x2e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4a,
	0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f,
	0x03, 0x80, 0x01, 0x14, 0x52, 0x02, 0x49, 0x44, 0x22, 0x2b, 0x0a, 0x10, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01,
	0x14, 0x52, 0x02, 0x49, 0x44, 0x32, 0xb9, 0x01, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x06, 0x52, 0x75, 0x6e, 0x4a, 0x6f, 0x62, 0x12, 0x14,
	0x2e, 0x6a, 0x6f, 0x62, 0x70, 0x62, 0x2e, 0x52, 0x75, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x6a, 0x6f, 0x62, 0x70, 0x62, 0x2e, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4a, 0x6f,
	0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x2e, 0x6a, 0x6f, 0x62, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x6a, 0x6f, 0x62, 0x70, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x4a, 0x6f, 0x62, 0x12, 0x17, 0x2e, 0x6a, 0x6f, 0x62, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x6a,
	0x6f, 0x62, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x44, 0x61, 0x74, 0x61, 0x57, 0x6f, 0x72, 0x6b, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2f, 0x67, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6a, 0x6f, 0x62, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_jobmanager_proto_rawDescOnce sync.Once
	file_proto_jobmanager_proto_rawDescData = file_proto_jobmanager_proto_rawDesc
)

func file_proto_jobmanager_proto_rawDescGZIP() []byte {
	file_proto_jobmanager_proto_rawDescOnce.Do(func() {
		file_proto_jobmanager_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_jobmanager_proto_rawDescData)
	})
	return file_proto_jobmanager_proto_rawDescData
}

var file_proto_jobmanager_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_jobmanager_proto_goTypes = []interface{}{
	(*JobReply)(nil),            // 0: jobpb.JobReply
	(*EmptyReply)(nil),          // 1: jobpb.EmptyReply
	(*RunJobRequest)(nil),       // 2: jobpb.RunJobRequest
	(*GetJobStatusRequest)(nil), // 3: jobpb.GetJobStatusRequest
	(*CancelJobRequest)(nil),    // 4: jobpb.CancelJobRequest
}
var file_proto_jobmanager_proto_depIdxs = []int32{
	2, // 0: jobpb.Jobmanager.RunJob:input_type -> jobpb.RunJobRequest
	3, // 1: jobpb.Jobmanager.GetJobStatus:input_type -> jobpb.GetJobStatusRequest
	4, // 2: jobpb.Jobmanager.CancelJob:input_type -> jobpb.CancelJobRequest
	0, // 3: jobpb.Jobmanager.RunJob:output_type -> jobpb.JobReply
	0, // 4: jobpb.Jobmanager.GetJobStatus:output_type -> jobpb.JobReply
	1, // 5: jobpb.Jobmanager.CancelJob:output_type -> jobpb.EmptyReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_jobmanager_proto_init() }
func file_proto_jobmanager_proto_init() {
	if File_proto_jobmanager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_jobmanager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobReply); i {
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
		file_proto_jobmanager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyReply); i {
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
		file_proto_jobmanager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunJobRequest); i {
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
		file_proto_jobmanager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJobStatusRequest); i {
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
		file_proto_jobmanager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelJobRequest); i {
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
			RawDescriptor: file_proto_jobmanager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_jobmanager_proto_goTypes,
		DependencyIndexes: file_proto_jobmanager_proto_depIdxs,
		MessageInfos:      file_proto_jobmanager_proto_msgTypes,
	}.Build()
	File_proto_jobmanager_proto = out.File
	file_proto_jobmanager_proto_rawDesc = nil
	file_proto_jobmanager_proto_goTypes = nil
	file_proto_jobmanager_proto_depIdxs = nil
}
