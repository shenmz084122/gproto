// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: proto/types/model/flink/flink_job.proto

package pbflink

import (
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
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

// reference:
//
//	https://nightlies.apache.org/flink/flink-docs-master/api/java/org/apache/flink/api/common/JobStatus.html
type FlinkJob_Status int32

const (
	FlinkJob_StateUnset FlinkJob_Status = 0
	// Job has been cancelled.
	FlinkJob_CANCELED FlinkJob_Status = 1
	// Job is being cancelled.
	FlinkJob_CANCELLING FlinkJob_Status = 2
	// Job is newly created, no task has started to run.
	FlinkJob_CREATED FlinkJob_Status = 3
	// The job has failed with a non-recoverable task failure.
	FlinkJob_FAILED FlinkJob_Status = 4
	// The job has failed and is currently waiting for the cleanup to complete.
	FlinkJob_FAILING FlinkJob_Status = 5
	// All of the job's tasks have successfully finished.
	FlinkJob_FINISHED FlinkJob_Status = 6
	// The job has been received by the Dispatcher, and is waiting for the job manager to receive leadership and to be created.
	FlinkJob_INITIALIZING FlinkJob_Status = 7
	// The job is currently reconciling and waits for task execution report to recover state.
	FlinkJob_RECONCILING FlinkJob_Status = 8
	// The job is currently undergoing a reset and total restart.
	FlinkJob_RESTARTING FlinkJob_Status = 9
	// Some tasks are scheduled or running, some may be pending, some may be finished.
	FlinkJob_RUNNING FlinkJob_Status = 10
	// The job has been suspended which means that it has been stopped but not been removed from a potential HA job store.
	FlinkJob_SUSPENDED FlinkJob_Status = 11
)

// Enum value maps for FlinkJob_Status.
var (
	FlinkJob_Status_name = map[int32]string{
		0:  "StateUnset",
		1:  "CANCELED",
		2:  "CANCELLING",
		3:  "CREATED",
		4:  "FAILED",
		5:  "FAILING",
		6:  "FINISHED",
		7:  "INITIALIZING",
		8:  "RECONCILING",
		9:  "RESTARTING",
		10: "RUNNING",
		11: "SUSPENDED",
	}
	FlinkJob_Status_value = map[string]int32{
		"StateUnset":   0,
		"CANCELED":     1,
		"CANCELLING":   2,
		"CREATED":      3,
		"FAILED":       4,
		"FAILING":      5,
		"FINISHED":     6,
		"INITIALIZING": 7,
		"RECONCILING":  8,
		"RESTARTING":   9,
		"RUNNING":      10,
		"SUSPENDED":    11,
	}
)

func (x FlinkJob_Status) Enum() *FlinkJob_Status {
	p := new(FlinkJob_Status)
	*p = x
	return p
}

func (x FlinkJob_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FlinkJob_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_types_model_flink_flink_job_proto_enumTypes[0].Descriptor()
}

func (FlinkJob_Status) Type() protoreflect.EnumType {
	return &file_proto_types_model_flink_flink_job_proto_enumTypes[0]
}

func (x FlinkJob_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FlinkJob_Status.Descriptor instead.
func (FlinkJob_Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_types_model_flink_flink_job_proto_rawDescGZIP(), []int{0, 0}
}

type FlinkOperator_Type int32

const (
	FlinkOperator_Empty     FlinkOperator_Type = 0
	FlinkOperator_Values    FlinkOperator_Type = 1
	FlinkOperator_Const     FlinkOperator_Type = 2
	FlinkOperator_Source    FlinkOperator_Type = 3
	FlinkOperator_Dimension FlinkOperator_Type = 4
	FlinkOperator_Dest      FlinkOperator_Type = 5
	FlinkOperator_OrderBy   FlinkOperator_Type = 6
	FlinkOperator_Limit     FlinkOperator_Type = 7
	FlinkOperator_Offset    FlinkOperator_Type = 8
	FlinkOperator_Fetch     FlinkOperator_Type = 9
	FlinkOperator_Filter    FlinkOperator_Type = 10
	FlinkOperator_Union     FlinkOperator_Type = 11
	FlinkOperator_Except    FlinkOperator_Type = 12
	FlinkOperator_Intersect FlinkOperator_Type = 13
	FlinkOperator_GroupBy   FlinkOperator_Type = 14
	FlinkOperator_Having    FlinkOperator_Type = 15
	FlinkOperator_Join      FlinkOperator_Type = 16
	FlinkOperator_UDTF      FlinkOperator_Type = 17
	FlinkOperator_UDTTF     FlinkOperator_Type = 18
)

// Enum value maps for FlinkOperator_Type.
var (
	FlinkOperator_Type_name = map[int32]string{
		0:  "Empty",
		1:  "Values",
		2:  "Const",
		3:  "Source",
		4:  "Dimension",
		5:  "Dest",
		6:  "OrderBy",
		7:  "Limit",
		8:  "Offset",
		9:  "Fetch",
		10: "Filter",
		11: "Union",
		12: "Except",
		13: "Intersect",
		14: "GroupBy",
		15: "Having",
		16: "Join",
		17: "UDTF",
		18: "UDTTF",
	}
	FlinkOperator_Type_value = map[string]int32{
		"Empty":     0,
		"Values":    1,
		"Const":     2,
		"Source":    3,
		"Dimension": 4,
		"Dest":      5,
		"OrderBy":   6,
		"Limit":     7,
		"Offset":    8,
		"Fetch":     9,
		"Filter":    10,
		"Union":     11,
		"Except":    12,
		"Intersect": 13,
		"GroupBy":   14,
		"Having":    15,
		"Join":      16,
		"UDTF":      17,
		"UDTTF":     18,
	}
)

func (x FlinkOperator_Type) Enum() *FlinkOperator_Type {
	p := new(FlinkOperator_Type)
	*p = x
	return p
}

func (x FlinkOperator_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FlinkOperator_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_types_model_flink_flink_job_proto_enumTypes[1].Descriptor()
}

func (FlinkOperator_Type) Type() protoreflect.EnumType {
	return &file_proto_types_model_flink_flink_job_proto_enumTypes[1]
}

func (x FlinkOperator_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FlinkOperator_Type.Descriptor instead.
func (FlinkOperator_Type) EnumDescriptor() ([]byte, []int) {
	return file_proto_types_model_flink_flink_job_proto_rawDescGZIP(), []int{1, 0}
}

// FlinkJob for describes information of flink job.
type FlinkJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FlinkJob) Reset() {
	*x = FlinkJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_model_flink_flink_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlinkJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlinkJob) ProtoMessage() {}

func (x *FlinkJob) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_model_flink_flink_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlinkJob.ProtoReflect.Descriptor instead.
func (*FlinkJob) Descriptor() ([]byte, []int) {
	return file_proto_types_model_flink_flink_job_proto_rawDescGZIP(), []int{0}
}

// FlinkOperator
type FlinkOperator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OperatorType, Optional Value:
	// "1-Values" "2-Const" "3-Source" "4-Dimension" "5-Dest" "6-OrderBy"
	// "7-Limit" "8-Offset" "9-Fetch" "10-Filter" "11-Union" "12-Except"
	// "13-Intersect" "14-GroupBy" "15-Having" "16-Join" "17-UDTF" "18-UDTTF"
	Type FlinkOperator_Type `protobuf:"varint,1,opt,name=type,proto3,enum=flink.FlinkOperator_Type" json:"type"`
	// nodeid is unique in this flow.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// this node name
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	// the upstream node id
	Upstream string `protobuf:"bytes,4,opt,name=upstream,proto3" json:"upstream"`
	// the right upstream node id
	UpstreamRight string `protobuf:"bytes,5,opt,name=upstream_right,json=upstreamRight,proto3" json:"upstream_right"`
	// the downstream node id
	DownStream string `protobuf:"bytes,6,opt,name=down_stream,json=downStream,proto3" json:"down_stream"`
	// the PointX
	PointX int32 `protobuf:"varint,7,opt,name=point_x,json=pointX,proto3" json:"point_x"`
	// the PointY
	PointY int32 `protobuf:"varint,8,opt,name=point_y,json=pointY,proto3" json:"point_y"`
	// this operator's property
	Property *OperatorProperty `protobuf:"bytes,9,opt,name=property,proto3" json:"property"`
}

func (x *FlinkOperator) Reset() {
	*x = FlinkOperator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_model_flink_flink_job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlinkOperator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlinkOperator) ProtoMessage() {}

func (x *FlinkOperator) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_model_flink_flink_job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlinkOperator.ProtoReflect.Descriptor instead.
func (*FlinkOperator) Descriptor() ([]byte, []int) {
	return file_proto_types_model_flink_flink_job_proto_rawDescGZIP(), []int{1}
}

func (x *FlinkOperator) GetType() FlinkOperator_Type {
	if x != nil {
		return x.Type
	}
	return FlinkOperator_Empty
}

func (x *FlinkOperator) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FlinkOperator) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FlinkOperator) GetUpstream() string {
	if x != nil {
		return x.Upstream
	}
	return ""
}

func (x *FlinkOperator) GetUpstreamRight() string {
	if x != nil {
		return x.UpstreamRight
	}
	return ""
}

func (x *FlinkOperator) GetDownStream() string {
	if x != nil {
		return x.DownStream
	}
	return ""
}

func (x *FlinkOperator) GetPointX() int32 {
	if x != nil {
		return x.PointX
	}
	return 0
}

func (x *FlinkOperator) GetPointY() int32 {
	if x != nil {
		return x.PointY
	}
	return 0
}

func (x *FlinkOperator) GetProperty() *OperatorProperty {
	if x != nil {
		return x.Property
	}
	return nil
}

// FlinkJar
type FlinkJar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of resource file. Is Required.
	FileId string `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id"`
	// JarArgs
	JarArgs string `protobuf:"bytes,2,opt,name=jar_args,json=jarArgs,proto3" json:"jar_args"`
	// JarEntry
	JarEntry string `protobuf:"bytes,3,opt,name=jar_entry,json=jarEntry,proto3" json:"jar_entry"`
	// The resource file id that has been deleted.
	// Only used to GetStreamJobArgs and GetStreamJobVersionArgs.
	DeleteFileId string `protobuf:"bytes,4,opt,name=delete_file_id,json=deleteFileId,proto3" json:"delete_file_id"`
}

func (x *FlinkJar) Reset() {
	*x = FlinkJar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_model_flink_flink_job_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlinkJar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlinkJar) ProtoMessage() {}

func (x *FlinkJar) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_model_flink_flink_job_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlinkJar.ProtoReflect.Descriptor instead.
func (*FlinkJar) Descriptor() ([]byte, []int) {
	return file_proto_types_model_flink_flink_job_proto_rawDescGZIP(), []int{2}
}

func (x *FlinkJar) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *FlinkJar) GetJarArgs() string {
	if x != nil {
		return x.JarArgs
	}
	return ""
}

func (x *FlinkJar) GetJarEntry() string {
	if x != nil {
		return x.JarEntry
	}
	return ""
}

func (x *FlinkJar) GetDeleteFileId() string {
	if x != nil {
		return x.DeleteFileId
	}
	return ""
}

// PythonOperatorProperty
type FlinkPythonCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// code.
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code"`
}

func (x *FlinkPythonCode) Reset() {
	*x = FlinkPythonCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_model_flink_flink_job_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlinkPythonCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlinkPythonCode) ProtoMessage() {}

func (x *FlinkPythonCode) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_model_flink_flink_job_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlinkPythonCode.ProtoReflect.Descriptor instead.
func (*FlinkPythonCode) Descriptor() ([]byte, []int) {
	return file_proto_types_model_flink_flink_job_proto_rawDescGZIP(), []int{3}
}

func (x *FlinkPythonCode) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type FlinkPythonFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// python_path.
	PythonPath string `protobuf:"bytes,1,opt,name=python_path,json=pythonPath,proto3" json:"python_path"`
	// python_module.
	PythonModule string `protobuf:"bytes,2,opt,name=python_module,json=pythonModule,proto3" json:"python_module"`
	// python_args.
	PythonArgs string `protobuf:"bytes,3,opt,name=python_args,json=pythonArgs,proto3" json:"python_args"`
	// delete_file_id.
	// Only used to GetStreamJobArgs and GetStreamJobVersionArgs.
	DeleteFileId string `protobuf:"bytes,4,opt,name=delete_file_id,json=deleteFileId,proto3" json:"delete_file_id"`
}

func (x *FlinkPythonFile) Reset() {
	*x = FlinkPythonFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_model_flink_flink_job_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlinkPythonFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlinkPythonFile) ProtoMessage() {}

func (x *FlinkPythonFile) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_model_flink_flink_job_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlinkPythonFile.ProtoReflect.Descriptor instead.
func (*FlinkPythonFile) Descriptor() ([]byte, []int) {
	return file_proto_types_model_flink_flink_job_proto_rawDescGZIP(), []int{4}
}

func (x *FlinkPythonFile) GetPythonPath() string {
	if x != nil {
		return x.PythonPath
	}
	return ""
}

func (x *FlinkPythonFile) GetPythonModule() string {
	if x != nil {
		return x.PythonModule
	}
	return ""
}

func (x *FlinkPythonFile) GetPythonArgs() string {
	if x != nil {
		return x.PythonArgs
	}
	return ""
}

func (x *FlinkPythonFile) GetDeleteFileId() string {
	if x != nil {
		return x.DeleteFileId
	}
	return ""
}

// FlinkSQL
type FlinkSQL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// sql code.
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code"`
}

func (x *FlinkSQL) Reset() {
	*x = FlinkSQL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_model_flink_flink_job_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlinkSQL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlinkSQL) ProtoMessage() {}

func (x *FlinkSQL) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_model_flink_flink_job_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlinkSQL.ProtoReflect.Descriptor instead.
func (*FlinkSQL) Descriptor() ([]byte, []int) {
	return file_proto_types_model_flink_flink_job_proto_rawDescGZIP(), []int{5}
}

func (x *FlinkSQL) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_proto_types_model_flink_flink_job_proto protoreflect.FileDescriptor

var file_proto_types_model_flink_flink_job_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x66, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x66, 0x6c, 0x69, 0x6e, 0x6b, 0x5f,
	0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x66, 0x6c, 0x69, 0x6e, 0x6b,
	0x1a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x75, 0x33,
	0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x66, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x66,
	0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x08, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x4a, 0x6f, 0x62,
	0x22, 0xb9, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x0a, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x55, 0x6e, 0x73, 0x65, 0x74, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43,
	0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x41, 0x4e,
	0x43, 0x45, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45,
	0x41, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44,
	0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12,
	0x0c, 0x0a, 0x08, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x06, 0x12, 0x10, 0x0a,
	0x0c, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x49, 0x5a, 0x49, 0x4e, 0x47, 0x10, 0x07, 0x12,
	0x0f, 0x0a, 0x0b, 0x52, 0x45, 0x43, 0x4f, 0x4e, 0x43, 0x49, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x08,
	0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x45, 0x53, 0x54, 0x41, 0x52, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x09,
	0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x0a, 0x12, 0x0d, 0x0a,
	0x09, 0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x0b, 0x22, 0x91, 0x05, 0x0a,
	0x0d, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x3c,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x66,
	0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0d, 0xe2, 0xdf, 0x1f, 0x09, 0x12, 0x07, 0xda,
	0x01, 0x04, 0x30, 0x00, 0x58, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe2, 0xdf, 0x1f, 0x08, 0x12, 0x06,
	0xc2, 0x01, 0x03, 0xf0, 0x01, 0x14, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe2, 0xdf, 0x1f, 0x08, 0x12, 0x06,
	0xc2, 0x01, 0x03, 0x88, 0x02, 0x41, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x08,
	0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c,
	0xe2, 0xdf, 0x1f, 0x08, 0x12, 0x06, 0xc2, 0x01, 0x03, 0xf0, 0x01, 0x14, 0x52, 0x08, 0x75, 0x70,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x33, 0x0a, 0x0e, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x5f, 0x72, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c,
	0xe2, 0xdf, 0x1f, 0x08, 0x12, 0x06, 0xc2, 0x01, 0x03, 0xf0, 0x01, 0x14, 0x52, 0x0d, 0x75, 0x70,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x64,
	0x6f, 0x77, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0c, 0xe2, 0xdf, 0x1f, 0x08, 0x12, 0x06, 0xc2, 0x01, 0x03, 0xf0, 0x01, 0x15, 0x52, 0x0a,
	0x64, 0x6f, 0x77, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x27, 0x0a, 0x07, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x5f, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0e, 0xe2, 0xdf, 0x1f,
	0x0a, 0x12, 0x08, 0xb2, 0x01, 0x05, 0x38, 0xc8, 0x01, 0x40, 0x01, 0x52, 0x06, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x58, 0x12, 0x27, 0x0a, 0x07, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x79, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x0e, 0xe2, 0xdf, 0x1f, 0x0a, 0x12, 0x08, 0xb2, 0x01, 0x05, 0x38,
	0xc8, 0x01, 0x40, 0x01, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x59, 0x12, 0x39, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x66, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x50,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x42, 0x04, 0xe2, 0xdf, 0x1f, 0x00, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x22, 0xe6, 0x01, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x09, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x73, 0x74,
	0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x10, 0x03, 0x12, 0x0d,
	0x0a, 0x09, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x10, 0x04, 0x12, 0x08, 0x0a,
	0x04, 0x44, 0x65, 0x73, 0x74, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x42, 0x79, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x10, 0x07, 0x12,
	0x0a, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x10, 0x08, 0x12, 0x09, 0x0a, 0x05, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x10, 0x09, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x10, 0x0a, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x10, 0x0b, 0x12, 0x0a, 0x0a,
	0x06, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x10, 0x0c, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x73, 0x65, 0x63, 0x74, 0x10, 0x0d, 0x12, 0x0b, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x42, 0x79, 0x10, 0x0e, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x10,
	0x0f, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x10, 0x10, 0x12, 0x08, 0x0a, 0x04, 0x55,
	0x44, 0x54, 0x46, 0x10, 0x11, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x44, 0x54, 0x54, 0x46, 0x10, 0x12,
	0x22, 0xec, 0x01, 0x0a, 0x08, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x4a, 0x61, 0x72, 0x12, 0x2c, 0x0a,
	0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13,
	0xe2, 0xdf, 0x1f, 0x0f, 0x12, 0x0d, 0xc2, 0x01, 0x0a, 0xf0, 0x01, 0x14, 0xca, 0x02, 0x04, 0x72,
	0x65, 0x73, 0x2d, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x08, 0x6a,
	0x61, 0x72, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xe2,
	0xdf, 0x1f, 0x0c, 0x12, 0x0a, 0xc2, 0x01, 0x07, 0x98, 0x02, 0x80, 0x08, 0x88, 0x05, 0x01, 0x52,
	0x07, 0x6a, 0x61, 0x72, 0x41, 0x72, 0x67, 0x73, 0x12, 0x2d, 0x0a, 0x09, 0x6a, 0x61, 0x72, 0x5f,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xe2, 0xdf, 0x1f,
	0x0c, 0x12, 0x0a, 0xc2, 0x01, 0x07, 0x98, 0x02, 0x80, 0x08, 0x88, 0x05, 0x01, 0x52, 0x08, 0x6a,
	0x61, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x56, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x30, 0xe2, 0xdf, 0x1f, 0x19, 0x0a, 0x17, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x12, 0x05, 0xc2, 0x01, 0x02, 0x22, 0x00, 0xe2, 0xdf,
	0x1f, 0x0f, 0x12, 0x0d, 0xc2, 0x01, 0x0a, 0xf0, 0x01, 0x14, 0xca, 0x02, 0x04, 0x72, 0x65, 0x73,
	0x2d, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22,
	0x38, 0x0a, 0x0f, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x50, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x11, 0xe2, 0xdf, 0x1f, 0x0d, 0x12, 0x0b, 0xc2, 0x01, 0x08, 0x98, 0x02, 0xc0, 0xb8, 0x02,
	0x88, 0x05, 0x01, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x89, 0x02, 0x0a, 0x0f, 0x46, 0x6c,
	0x69, 0x6e, 0x6b, 0x50, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x34, 0x0a,
	0x0b, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x13, 0xe2, 0xdf, 0x1f, 0x0f, 0x12, 0x0d, 0xc2, 0x01, 0x0a, 0x98, 0x02, 0x14,
	0xca, 0x02, 0x04, 0x72, 0x65, 0x73, 0x2d, 0x52, 0x0a, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x35, 0x0a, 0x0d, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xe2, 0xdf, 0x1f, 0x0c,
	0x12, 0x0a, 0xc2, 0x01, 0x07, 0x98, 0x02, 0x80, 0x08, 0x88, 0x05, 0x01, 0x52, 0x0c, 0x70, 0x79,
	0x74, 0x68, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x70, 0x79,
	0x74, 0x68, 0x6f, 0x6e, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x10, 0xe2, 0xdf, 0x1f, 0x0c, 0x12, 0x0a, 0xc2, 0x01, 0x07, 0x98, 0x02, 0x80, 0x08, 0x88, 0x05,
	0x01, 0x52, 0x0a, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x41, 0x72, 0x67, 0x73, 0x12, 0x56, 0x0a,
	0x0e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0xe2, 0xdf, 0x1f, 0x19, 0x0a, 0x17, 0x0a, 0x0e, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x12, 0x05, 0xc2,
	0x01, 0x02, 0x22, 0x00, 0xe2, 0xdf, 0x1f, 0x0f, 0x12, 0x0d, 0xc2, 0x01, 0x0a, 0xf0, 0x01, 0x14,
	0xca, 0x02, 0x04, 0x72, 0x65, 0x73, 0x2d, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x08, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x53, 0x51,
	0x4c, 0x12, 0x25, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x11, 0xe2, 0xdf, 0x1f, 0x0d, 0x12, 0x0b, 0xc2, 0x01, 0x08, 0x98, 0x02, 0xc0, 0xb8, 0x02, 0x88,
	0x05, 0x01, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x75, 0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x6f, 0x6d, 0x6e, 0x69, 0x73, 0x2e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x62, 0x66, 0x6c, 0x69, 0x6e, 0x6b, 0x42, 0x0a, 0x50, 0x42, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x4a,
	0x6f, 0x62, 0x50, 0x00, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x44, 0x61, 0x74, 0x61, 0x57, 0x6f, 0x72, 0x6b, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2f, 0x67,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x78, 0x67, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x70, 0x62, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x62, 0x66, 0x6c, 0x69, 0x6e, 0x6b, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_types_model_flink_flink_job_proto_rawDescOnce sync.Once
	file_proto_types_model_flink_flink_job_proto_rawDescData = file_proto_types_model_flink_flink_job_proto_rawDesc
)

func file_proto_types_model_flink_flink_job_proto_rawDescGZIP() []byte {
	file_proto_types_model_flink_flink_job_proto_rawDescOnce.Do(func() {
		file_proto_types_model_flink_flink_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_model_flink_flink_job_proto_rawDescData)
	})
	return file_proto_types_model_flink_flink_job_proto_rawDescData
}

var file_proto_types_model_flink_flink_job_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_types_model_flink_flink_job_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_types_model_flink_flink_job_proto_goTypes = []interface{}{
	(FlinkJob_Status)(0),     // 0: flink.FlinkJob.Status
	(FlinkOperator_Type)(0),  // 1: flink.FlinkOperator.Type
	(*FlinkJob)(nil),         // 2: flink.FlinkJob
	(*FlinkOperator)(nil),    // 3: flink.FlinkOperator
	(*FlinkJar)(nil),         // 4: flink.FlinkJar
	(*FlinkPythonCode)(nil),  // 5: flink.FlinkPythonCode
	(*FlinkPythonFile)(nil),  // 6: flink.FlinkPythonFile
	(*FlinkSQL)(nil),         // 7: flink.FlinkSQL
	(*OperatorProperty)(nil), // 8: flink.OperatorProperty
}
var file_proto_types_model_flink_flink_job_proto_depIdxs = []int32{
	1, // 0: flink.FlinkOperator.type:type_name -> flink.FlinkOperator.Type
	8, // 1: flink.FlinkOperator.property:type_name -> flink.OperatorProperty
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_types_model_flink_flink_job_proto_init() }
func file_proto_types_model_flink_flink_job_proto_init() {
	if File_proto_types_model_flink_flink_job_proto != nil {
		return
	}
	file_proto_types_model_flink_flink_operator_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_types_model_flink_flink_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlinkJob); i {
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
		file_proto_types_model_flink_flink_job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlinkOperator); i {
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
		file_proto_types_model_flink_flink_job_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlinkJar); i {
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
		file_proto_types_model_flink_flink_job_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlinkPythonCode); i {
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
		file_proto_types_model_flink_flink_job_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlinkPythonFile); i {
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
		file_proto_types_model_flink_flink_job_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlinkSQL); i {
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
			RawDescriptor: file_proto_types_model_flink_flink_job_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_model_flink_flink_job_proto_goTypes,
		DependencyIndexes: file_proto_types_model_flink_flink_job_proto_depIdxs,
		EnumInfos:         file_proto_types_model_flink_flink_job_proto_enumTypes,
		MessageInfos:      file_proto_types_model_flink_flink_job_proto_msgTypes,
	}.Build()
	File_proto_types_model_flink_flink_job_proto = out.File
	file_proto_types_model_flink_flink_job_proto_rawDesc = nil
	file_proto_types_model_flink_flink_job_proto_goTypes = nil
	file_proto_types_model_flink_flink_job_proto_depIdxs = nil
}
