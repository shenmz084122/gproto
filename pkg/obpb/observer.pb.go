// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/observer.proto

package obpb

import (
	model "github.com/DataWorkbench/gproto/pkg/model"
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

type PTasksStatusStatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpaceId   string `protobuf:"bytes,1,opt,name=space_id,json=spaceId,proto3" json:"space_id,omitempty"`
	StartTime int64  `protobuf:"varint,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   int64  `protobuf:"varint,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *PTasksStatusStatRequest) Reset() {
	*x = PTasksStatusStatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_observer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PTasksStatusStatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PTasksStatusStatRequest) ProtoMessage() {}

func (x *PTasksStatusStatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_observer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PTasksStatusStatRequest.ProtoReflect.Descriptor instead.
func (*PTasksStatusStatRequest) Descriptor() ([]byte, []int) {
	return file_proto_observer_proto_rawDescGZIP(), []int{0}
}

func (x *PTasksStatusStatRequest) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *PTasksStatusStatRequest) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *PTasksStatusStatRequest) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

type PTasksStatusStatReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*model.InstanceStatusStat `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos,omitempty"`
}

func (x *PTasksStatusStatReply) Reset() {
	*x = PTasksStatusStatReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_observer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PTasksStatusStatReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PTasksStatusStatReply) ProtoMessage() {}

func (x *PTasksStatusStatReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_observer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PTasksStatusStatReply.ProtoReflect.Descriptor instead.
func (*PTasksStatusStatReply) Descriptor() ([]byte, []int) {
	return file_proto_observer_proto_rawDescGZIP(), []int{1}
}

func (x *PTasksStatusStatReply) GetInfos() []*model.InstanceStatusStat {
	if x != nil {
		return x.Infos
	}
	return nil
}

type PTasksExecStatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpaceId string `protobuf:"bytes,1,opt,name=space_id,json=spaceId,proto3" json:"space_id,omitempty"`
	State   int32  `protobuf:"varint,2,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *PTasksExecStatRequest) Reset() {
	*x = PTasksExecStatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_observer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PTasksExecStatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PTasksExecStatRequest) ProtoMessage() {}

func (x *PTasksExecStatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_observer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PTasksExecStatRequest.ProtoReflect.Descriptor instead.
func (*PTasksExecStatRequest) Descriptor() ([]byte, []int) {
	return file_proto_observer_proto_rawDescGZIP(), []int{2}
}

func (x *PTasksExecStatRequest) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *PTasksExecStatRequest) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

type PTasksExecStatReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Today     []*model.InstanceTaskExecStat `protobuf:"bytes,1,rep,name=today,proto3" json:"today,omitempty"`
	Yesterday []*model.InstanceTaskExecStat `protobuf:"bytes,2,rep,name=yesterday,proto3" json:"yesterday,omitempty"`
	History   []*model.InstanceTaskExecStat `protobuf:"bytes,3,rep,name=history,proto3" json:"history,omitempty"`
}

func (x *PTasksExecStatReply) Reset() {
	*x = PTasksExecStatReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_observer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PTasksExecStatReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PTasksExecStatReply) ProtoMessage() {}

func (x *PTasksExecStatReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_observer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PTasksExecStatReply.ProtoReflect.Descriptor instead.
func (*PTasksExecStatReply) Descriptor() ([]byte, []int) {
	return file_proto_observer_proto_rawDescGZIP(), []int{3}
}

func (x *PTasksExecStatReply) GetToday() []*model.InstanceTaskExecStat {
	if x != nil {
		return x.Today
	}
	return nil
}

func (x *PTasksExecStatReply) GetYesterday() []*model.InstanceTaskExecStat {
	if x != nil {
		return x.Yesterday
	}
	return nil
}

func (x *PTasksExecStatReply) GetHistory() []*model.InstanceTaskExecStat {
	if x != nil {
		return x.History
	}
	return nil
}

type PTaskDispatchCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpaceId   string `protobuf:"bytes,1,opt,name=space_id,json=spaceId,proto3" json:"space_id,omitempty"`
	StartTime int64  `protobuf:"varint,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   int64  `protobuf:"varint,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *PTaskDispatchCountRequest) Reset() {
	*x = PTaskDispatchCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_observer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PTaskDispatchCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PTaskDispatchCountRequest) ProtoMessage() {}

func (x *PTaskDispatchCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_observer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PTaskDispatchCountRequest.ProtoReflect.Descriptor instead.
func (*PTaskDispatchCountRequest) Descriptor() ([]byte, []int) {
	return file_proto_observer_proto_rawDescGZIP(), []int{4}
}

func (x *PTaskDispatchCountRequest) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *PTaskDispatchCountRequest) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *PTaskDispatchCountRequest) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

type PTaskDispatchCountReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*model.DispatchTaskCountInfo `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos,omitempty"`
}

func (x *PTaskDispatchCountReply) Reset() {
	*x = PTaskDispatchCountReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_observer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PTaskDispatchCountReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PTaskDispatchCountReply) ProtoMessage() {}

func (x *PTaskDispatchCountReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_observer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PTaskDispatchCountReply.ProtoReflect.Descriptor instead.
func (*PTaskDispatchCountReply) Descriptor() ([]byte, []int) {
	return file_proto_observer_proto_rawDescGZIP(), []int{5}
}

func (x *PTaskDispatchCountReply) GetInfos() []*model.DispatchTaskCountInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

type PTaskRuntimeRankingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpaceId string `protobuf:"bytes,1,opt,name=space_id,json=spaceId,proto3" json:"space_id,omitempty"`
	Limit   int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset  int32  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *PTaskRuntimeRankingRequest) Reset() {
	*x = PTaskRuntimeRankingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_observer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PTaskRuntimeRankingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PTaskRuntimeRankingRequest) ProtoMessage() {}

func (x *PTaskRuntimeRankingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_observer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PTaskRuntimeRankingRequest.ProtoReflect.Descriptor instead.
func (*PTaskRuntimeRankingRequest) Descriptor() ([]byte, []int) {
	return file_proto_observer_proto_rawDescGZIP(), []int{6}
}

func (x *PTaskRuntimeRankingRequest) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *PTaskRuntimeRankingRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PTaskRuntimeRankingRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type PTaskRuntimeRankingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*model.InstanceRuntimeRankInfo `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos,omitempty"`
}

func (x *PTaskRuntimeRankingReply) Reset() {
	*x = PTaskRuntimeRankingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_observer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PTaskRuntimeRankingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PTaskRuntimeRankingReply) ProtoMessage() {}

func (x *PTaskRuntimeRankingReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_observer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PTaskRuntimeRankingReply.ProtoReflect.Descriptor instead.
func (*PTaskRuntimeRankingReply) Descriptor() ([]byte, []int) {
	return file_proto_observer_proto_rawDescGZIP(), []int{7}
}

func (x *PTaskRuntimeRankingReply) GetInfos() []*model.InstanceRuntimeRankInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

type PTaskErrorRankingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpaceId string `protobuf:"bytes,1,opt,name=space_id,json=spaceId,proto3" json:"space_id,omitempty"`
	Limit   int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset  int32  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *PTaskErrorRankingRequest) Reset() {
	*x = PTaskErrorRankingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_observer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PTaskErrorRankingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PTaskErrorRankingRequest) ProtoMessage() {}

func (x *PTaskErrorRankingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_observer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PTaskErrorRankingRequest.ProtoReflect.Descriptor instead.
func (*PTaskErrorRankingRequest) Descriptor() ([]byte, []int) {
	return file_proto_observer_proto_rawDescGZIP(), []int{8}
}

func (x *PTaskErrorRankingRequest) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *PTaskErrorRankingRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PTaskErrorRankingRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type PTaskErrorRankingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*model.InstanceErrorRankInfo `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos,omitempty"`
}

func (x *PTaskErrorRankingReply) Reset() {
	*x = PTaskErrorRankingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_observer_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PTaskErrorRankingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PTaskErrorRankingReply) ProtoMessage() {}

func (x *PTaskErrorRankingReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_observer_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PTaskErrorRankingReply.ProtoReflect.Descriptor instead.
func (*PTaskErrorRankingReply) Descriptor() ([]byte, []int) {
	return file_proto_observer_proto_rawDescGZIP(), []int{9}
}

func (x *PTaskErrorRankingReply) GetInfos() []*model.InstanceErrorRankInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

var File_proto_observer_proto protoreflect.FileDescriptor

var file_proto_observer_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x17,
	0x50, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x08, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58,
	0x01, 0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0f,
	0xe2, 0xdf, 0x1f, 0x0b, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x65, 0x6e,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0f, 0xe2, 0xdf,
	0x1f, 0x0b, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x4e, 0x0a, 0x15, 0x50, 0x54, 0x61, 0x73, 0x6b, 0x73,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x35, 0x0a, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x42, 0x04, 0xe2, 0xdf, 0x1f, 0x00, 0x52,
	0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x61, 0x0a, 0x15, 0x50, 0x54, 0x61, 0x73, 0x6b, 0x73,
	0x45, 0x78, 0x65, 0x63, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x08, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x25, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x0f, 0xe2, 0xdf, 0x1f, 0x0b, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0x01, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0xcc, 0x01, 0x0a, 0x13, 0x50, 0x54,
	0x61, 0x73, 0x6b, 0x73, 0x45, 0x78, 0x65, 0x63, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x37, 0x0a, 0x05, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x65, 0x63, 0x53, 0x74, 0x61, 0x74, 0x42, 0x04, 0xe2,
	0xdf, 0x1f, 0x00, 0x52, 0x05, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x12, 0x3f, 0x0a, 0x09, 0x79, 0x65,
	0x73, 0x74, 0x65, 0x72, 0x64, 0x61, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x45, 0x78, 0x65, 0x63, 0x53, 0x74, 0x61, 0x74, 0x42, 0x04, 0xe2, 0xdf, 0x1f, 0x00,
	0x52, 0x09, 0x79, 0x65, 0x73, 0x74, 0x65, 0x72, 0x64, 0x61, 0x79, 0x12, 0x3b, 0x0a, 0x07, 0x68,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x45, 0x78, 0x65, 0x63, 0x53, 0x74, 0x61, 0x74, 0x42, 0x04, 0xe2, 0xdf, 0x1f, 0x00, 0x52,
	0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x9a, 0x01, 0x0a, 0x19, 0x50, 0x54, 0x61,
	0x73, 0x6b, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x08, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01,
	0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0f, 0xe2,
	0xdf, 0x1f, 0x0b, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0f, 0xe2, 0xdf, 0x1f,
	0x0b, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x52, 0x07, 0x65, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x53, 0x0a, 0x17, 0x50, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x69,
	0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x38, 0x0a, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x04, 0xe2,
	0xdf, 0x1f, 0x00, 0x52, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x89, 0x01, 0x0a, 0x1a, 0x50,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x08, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f,
	0x03, 0x80, 0x01, 0x14, 0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0xe2, 0xdf,
	0x1f, 0x04, 0x10, 0x00, 0x18, 0x65, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x27, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0f, 0xe2,
	0xdf, 0x1f, 0x0b, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x52, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x56, 0x0a, 0x18, 0x50, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x3a, 0x0a, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x6e, 0x66,
	0x6f, 0x42, 0x04, 0xe2, 0xdf, 0x1f, 0x00, 0x52, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x87,
	0x01, 0x0a, 0x18, 0x50, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x61, 0x6e,
	0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x08, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2,
	0xdf, 0x1f, 0x03, 0x80, 0x01, 0x14, 0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08,
	0xe2, 0xdf, 0x1f, 0x04, 0x10, 0x00, 0x18, 0x65, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x27, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x0f, 0xe2, 0xdf, 0x1f, 0x0b, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x52, 0x0a, 0x16, 0x50, 0x54, 0x61, 0x73,
	0x6b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x38, 0x0a, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x42,
	0x04, 0xe2, 0xdf, 0x1f, 0x00, 0x52, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x32, 0xc4, 0x03, 0x0a,
	0x08, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x56, 0x0a, 0x20, 0x47, 0x65, 0x74,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x69, 0x63, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x18, 0x2e,
	0x50, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x50, 0x54, 0x61, 0x73, 0x6b, 0x73,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x55, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x69, 0x63,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x16, 0x2e, 0x50, 0x54, 0x61, 0x73, 0x6b,
	0x73, 0x45, 0x78, 0x65, 0x63, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x50, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x45, 0x78, 0x65, 0x63, 0x53, 0x74, 0x61,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x50,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x69, 0x63, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x44, 0x69, 0x73, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x50, 0x54, 0x61, 0x73,
	0x6b, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x50, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x69, 0x73,
	0x70, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x5a, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x69, 0x63,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x6b,
	0x69, 0x6e, 0x67, 0x12, 0x1b, 0x2e, 0x50, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x50, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x52,
	0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x54, 0x0a,
	0x1c, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x69, 0x63, 0x54, 0x61, 0x73, 0x6b,
	0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x19, 0x2e,
	0x50, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x50, 0x54, 0x61, 0x73, 0x6b,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x57, 0x6f, 0x72, 0x6b, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2f,
	0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6f, 0x62, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_observer_proto_rawDescOnce sync.Once
	file_proto_observer_proto_rawDescData = file_proto_observer_proto_rawDesc
)

func file_proto_observer_proto_rawDescGZIP() []byte {
	file_proto_observer_proto_rawDescOnce.Do(func() {
		file_proto_observer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_observer_proto_rawDescData)
	})
	return file_proto_observer_proto_rawDescData
}

var file_proto_observer_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_observer_proto_goTypes = []interface{}{
	(*PTasksStatusStatRequest)(nil),       // 0: PTasksStatusStatRequest
	(*PTasksStatusStatReply)(nil),         // 1: PTasksStatusStatReply
	(*PTasksExecStatRequest)(nil),         // 2: PTasksExecStatRequest
	(*PTasksExecStatReply)(nil),           // 3: PTasksExecStatReply
	(*PTaskDispatchCountRequest)(nil),     // 4: PTaskDispatchCountRequest
	(*PTaskDispatchCountReply)(nil),       // 5: PTaskDispatchCountReply
	(*PTaskRuntimeRankingRequest)(nil),    // 6: PTaskRuntimeRankingRequest
	(*PTaskRuntimeRankingReply)(nil),      // 7: PTaskRuntimeRankingReply
	(*PTaskErrorRankingRequest)(nil),      // 8: PTaskErrorRankingRequest
	(*PTaskErrorRankingReply)(nil),        // 9: PTaskErrorRankingReply
	(*model.InstanceStatusStat)(nil),      // 10: model.InstanceStatusStat
	(*model.InstanceTaskExecStat)(nil),    // 11: model.InstanceTaskExecStat
	(*model.DispatchTaskCountInfo)(nil),   // 12: model.DispatchTaskCountInfo
	(*model.InstanceRuntimeRankInfo)(nil), // 13: model.InstanceRuntimeRankInfo
	(*model.InstanceErrorRankInfo)(nil),   // 14: model.InstanceErrorRankInfo
}
var file_proto_observer_proto_depIdxs = []int32{
	10, // 0: PTasksStatusStatReply.infos:type_name -> model.InstanceStatusStat
	11, // 1: PTasksExecStatReply.today:type_name -> model.InstanceTaskExecStat
	11, // 2: PTasksExecStatReply.yesterday:type_name -> model.InstanceTaskExecStat
	11, // 3: PTasksExecStatReply.history:type_name -> model.InstanceTaskExecStat
	12, // 4: PTaskDispatchCountReply.infos:type_name -> model.DispatchTaskCountInfo
	13, // 5: PTaskRuntimeRankingReply.infos:type_name -> model.InstanceRuntimeRankInfo
	14, // 6: PTaskErrorRankingReply.infos:type_name -> model.InstanceErrorRankInfo
	0,  // 7: Observer.GetPeriodicTasksStatusStatistics:input_type -> PTasksStatusStatRequest
	2,  // 8: Observer.GetPeriodicTasksExecutingStatistics:input_type -> PTasksExecStatRequest
	4,  // 9: Observer.GetPeriodicTasksDispatchCount:input_type -> PTaskDispatchCountRequest
	6,  // 10: Observer.GetPeriodicTasksRuntimeRanking:input_type -> PTaskRuntimeRankingRequest
	8,  // 11: Observer.GetPeriodicTasksErrorRanking:input_type -> PTaskErrorRankingRequest
	1,  // 12: Observer.GetPeriodicTasksStatusStatistics:output_type -> PTasksStatusStatReply
	3,  // 13: Observer.GetPeriodicTasksExecutingStatistics:output_type -> PTasksExecStatReply
	5,  // 14: Observer.GetPeriodicTasksDispatchCount:output_type -> PTaskDispatchCountReply
	7,  // 15: Observer.GetPeriodicTasksRuntimeRanking:output_type -> PTaskRuntimeRankingReply
	9,  // 16: Observer.GetPeriodicTasksErrorRanking:output_type -> PTaskErrorRankingReply
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_observer_proto_init() }
func file_proto_observer_proto_init() {
	if File_proto_observer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_observer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PTasksStatusStatRequest); i {
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
		file_proto_observer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PTasksStatusStatReply); i {
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
		file_proto_observer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PTasksExecStatRequest); i {
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
		file_proto_observer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PTasksExecStatReply); i {
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
		file_proto_observer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PTaskDispatchCountRequest); i {
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
		file_proto_observer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PTaskDispatchCountReply); i {
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
		file_proto_observer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PTaskRuntimeRankingRequest); i {
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
		file_proto_observer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PTaskRuntimeRankingReply); i {
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
		file_proto_observer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PTaskErrorRankingRequest); i {
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
		file_proto_observer_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PTaskErrorRankingReply); i {
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
			RawDescriptor: file_proto_observer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_observer_proto_goTypes,
		DependencyIndexes: file_proto_observer_proto_depIdxs,
		MessageInfos:      file_proto_observer_proto_msgTypes,
	}.Build()
	File_proto_observer_proto = out.File
	file_proto_observer_proto_rawDesc = nil
	file_proto_observer_proto_goTypes = nil
	file_proto_observer_proto_depIdxs = nil
}
