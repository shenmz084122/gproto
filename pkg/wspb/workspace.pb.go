// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: proto/workspace.proto

package wspb

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

// The Workspace Info
type Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Owner       string `protobuf:"bytes,2,opt,name=Owner,proto3" json:"Owner,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	Status      int32  `protobuf:"varint,5,opt,name=Status,proto3" json:"Status,omitempty"`
	Created     int64  `protobuf:"varint,6,opt,name=Created,proto3" json:"Created,omitempty"`
	Updated     int64  `protobuf:"varint,7,opt,name=Updated,proto3" json:"Updated,omitempty"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_workspace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_proto_workspace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Info.ProtoReflect.Descriptor instead.
func (*Info) Descriptor() ([]byte, []int) {
	return file_proto_workspace_proto_rawDescGZIP(), []int{0}
}

func (x *Info) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Info) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Info) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Info) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Info) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Info) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Info) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

//
// Request parameters used to ListWorkspaces
type ListsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner  string `protobuf:"bytes,1,opt,name=Owner,proto3" json:"Owner,omitempty"`
	Limit  int32  `protobuf:"varint,2,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset int32  `protobuf:"varint,3,opt,name=Offset,proto3" json:"Offset,omitempty"`
}

func (x *ListsRequest) Reset() {
	*x = ListsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_workspace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListsRequest) ProtoMessage() {}

func (x *ListsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_workspace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListsRequest.ProtoReflect.Descriptor instead.
func (*ListsRequest) Descriptor() ([]byte, []int) {
	return file_proto_workspace_proto_rawDescGZIP(), []int{1}
}

func (x *ListsRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *ListsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

// Reply parameters used to ListWorkspaces
type ListsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Info `protobuf:"bytes,1,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *ListsReply) Reset() {
	*x = ListsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_workspace_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListsReply) ProtoMessage() {}

func (x *ListsReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_workspace_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListsReply.ProtoReflect.Descriptor instead.
func (*ListsReply) Descriptor() ([]byte, []int) {
	return file_proto_workspace_proto_rawDescGZIP(), []int{2}
}

func (x *ListsReply) GetInfos() []*Info {
	if x != nil {
		return x.Infos
	}
	return nil
}

//
// Request parameters used to CreateWorkspace
type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner       string `protobuf:"bytes,1,opt,name=Owner,proto3" json:"Owner,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_workspace_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_workspace_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_workspace_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *CreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

//
// Request parameters used to DeleteWorkspace
type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner string `protobuf:"bytes,1,opt,name=Owner,proto3" json:"Owner,omitempty"`
	ID    string `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_workspace_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_workspace_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_workspace_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *DeleteRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

//
// Request parameters used to UpdateWorkspace
type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner       string `protobuf:"bytes,1,opt,name=Owner,proto3" json:"Owner,omitempty"`
	ID          string `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_workspace_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_workspace_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_workspace_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *UpdateRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

//
// Request parameters used to DescribeWorkspace
type DescribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner string `protobuf:"bytes,1,opt,name=Owner,proto3" json:"Owner,omitempty"`
	ID    string `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DescribeRequest) Reset() {
	*x = DescribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_workspace_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRequest) ProtoMessage() {}

func (x *DescribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_workspace_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeRequest.ProtoReflect.Descriptor instead.
func (*DescribeRequest) Descriptor() ([]byte, []int) {
	return file_proto_workspace_proto_rawDescGZIP(), []int{6}
}

func (x *DescribeRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *DescribeRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

// Reply parameters used to DescribeWorkspace
type DescribeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Info `protobuf:"bytes,1,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DescribeReply) Reset() {
	*x = DescribeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_workspace_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeReply) ProtoMessage() {}

func (x *DescribeReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_workspace_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeReply.ProtoReflect.Descriptor instead.
func (*DescribeReply) Descriptor() ([]byte, []int) {
	return file_proto_workspace_proto_rawDescGZIP(), []int{7}
}

func (x *DescribeReply) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

//
// Request parameters used to DisableWorkspace
type DisableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner string `protobuf:"bytes,1,opt,name=Owner,proto3" json:"Owner,omitempty"`
	ID    string `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DisableRequest) Reset() {
	*x = DisableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_workspace_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisableRequest) ProtoMessage() {}

func (x *DisableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_workspace_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisableRequest.ProtoReflect.Descriptor instead.
func (*DisableRequest) Descriptor() ([]byte, []int) {
	return file_proto_workspace_proto_rawDescGZIP(), []int{8}
}

func (x *DisableRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *DisableRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

//
// Request parameters used to EnableWorkspace
type EnableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner string `protobuf:"bytes,1,opt,name=Owner,proto3" json:"Owner,omitempty"`
	ID    string `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *EnableRequest) Reset() {
	*x = EnableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_workspace_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnableRequest) ProtoMessage() {}

func (x *EnableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_workspace_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnableRequest.ProtoReflect.Descriptor instead.
func (*EnableRequest) Descriptor() ([]byte, []int) {
	return file_proto_workspace_proto_rawDescGZIP(), []int{9}
}

func (x *EnableRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *EnableRequest) GetID() string {
	if x != nil {
		return x.ID
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
		mi := &file_proto_workspace_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyReply) ProtoMessage() {}

func (x *EmptyReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_workspace_proto_msgTypes[10]
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
	return file_proto_workspace_proto_rawDescGZIP(), []int{10}
}

var File_proto_workspace_proto protoreflect.FileDescriptor

var file_proto_workspace_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x77, 0x73, 0x70, 0x62, 0x1a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f,
	0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x01, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03,
	0x80, 0x01, 0x14, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x78, 0x41, 0x52, 0x05,
	0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x09, 0xe2, 0xdf, 0x1f, 0x05, 0x70, 0x01, 0x78, 0x81, 0x01, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x78,
	0x81, 0x08, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x20, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x08, 0xe2, 0xdf, 0x1f, 0x04, 0x10, 0x00, 0x18, 0x03, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x20, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x07, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x07, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x07, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x6c, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x05, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x08, 0xe2, 0xdf, 0x1f, 0x04, 0x10, 0x00, 0x18, 0x65, 0x52, 0x05, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x1e, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x06, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x22, 0x36, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x28, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x77, 0x73, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x06, 0xe2, 0xdf,
	0x1f, 0x02, 0x20, 0x00, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x79, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x05,
	0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xe2, 0xdf, 0x1f,
	0x04, 0x70, 0x0a, 0x78, 0x41, 0x52, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xe2, 0xdf, 0x1f, 0x05,
	0x70, 0x01, 0x78, 0x81, 0x01, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x0b, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x78, 0x81, 0x08, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x48, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xe2, 0xdf, 0x1f, 0x04, 0x70, 0x0a, 0x78, 0x41,
	0x52, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01, 0x14, 0x52, 0x02, 0x49, 0x44,
	0x22, 0x90, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x08, 0xe2, 0xdf, 0x1f, 0x04, 0x70, 0x0a, 0x78, 0x41, 0x52, 0x05, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x12, 0x17, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01, 0x14, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1b, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x78,
	0x81, 0x01, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2,
	0xdf, 0x1f, 0x03, 0x78, 0x81, 0x08, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x4a, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xe2, 0xdf, 0x1f, 0x04, 0x70, 0x0a, 0x78, 0x41, 0x52,
	0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01, 0x14, 0x52, 0x02, 0x49, 0x44, 0x22,
	0x2f, 0x0a, 0x0d, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x1e, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x77, 0x73, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x49, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x08, 0xe2, 0xdf, 0x1f, 0x04, 0x70, 0x0a, 0x78, 0x41, 0x52, 0x05, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x12, 0x17, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01, 0x14, 0x52, 0x02, 0x49, 0x44, 0x22, 0x48, 0x0a, 0x0d, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x05,
	0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xe2, 0xdf, 0x1f,
	0x04, 0x70, 0x0a, 0x78, 0x41, 0x52, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01,
	0x14, 0x52, 0x02, 0x49, 0x44, 0x22, 0x0c, 0x0a, 0x0a, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x32, 0xf7, 0x02, 0x0a, 0x09, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x2f, 0x0a, 0x05, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x12, 0x2e, 0x77, 0x73, 0x70,
	0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x77, 0x73, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x31, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x77,
	0x73, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x77, 0x73, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x13, 0x2e, 0x77, 0x73, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x77, 0x73, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x13, 0x2e, 0x77, 0x73, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x77, 0x73, 0x70, 0x62, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x08, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x15, 0x2e, 0x77, 0x73, 0x70, 0x62, 0x2e, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x77, 0x73, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x07, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x14, 0x2e, 0x77, 0x73, 0x70, 0x62, 0x2e, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x77, 0x73, 0x70, 0x62, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x06, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x13, 0x2e, 0x77, 0x73, 0x70, 0x62, 0x2e, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x77, 0x73, 0x70, 0x62,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0a, 0x5a,
	0x08, 0x70, 0x6b, 0x67, 0x2f, 0x77, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_workspace_proto_rawDescOnce sync.Once
	file_proto_workspace_proto_rawDescData = file_proto_workspace_proto_rawDesc
)

func file_proto_workspace_proto_rawDescGZIP() []byte {
	file_proto_workspace_proto_rawDescOnce.Do(func() {
		file_proto_workspace_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_workspace_proto_rawDescData)
	})
	return file_proto_workspace_proto_rawDescData
}

var file_proto_workspace_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_workspace_proto_goTypes = []interface{}{
	(*Info)(nil),            // 0: wspb.Info
	(*ListsRequest)(nil),    // 1: wspb.ListsRequest
	(*ListsReply)(nil),      // 2: wspb.ListsReply
	(*CreateRequest)(nil),   // 3: wspb.CreateRequest
	(*DeleteRequest)(nil),   // 4: wspb.DeleteRequest
	(*UpdateRequest)(nil),   // 5: wspb.UpdateRequest
	(*DescribeRequest)(nil), // 6: wspb.DescribeRequest
	(*DescribeReply)(nil),   // 7: wspb.DescribeReply
	(*DisableRequest)(nil),  // 8: wspb.DisableRequest
	(*EnableRequest)(nil),   // 9: wspb.EnableRequest
	(*EmptyReply)(nil),      // 10: wspb.EmptyReply
}
var file_proto_workspace_proto_depIdxs = []int32{
	0,  // 0: wspb.ListsReply.Infos:type_name -> wspb.Info
	0,  // 1: wspb.DescribeReply.Info:type_name -> wspb.Info
	1,  // 2: wspb.Workspace.Lists:input_type -> wspb.ListsRequest
	3,  // 3: wspb.Workspace.Create:input_type -> wspb.CreateRequest
	4,  // 4: wspb.Workspace.Delete:input_type -> wspb.DeleteRequest
	5,  // 5: wspb.Workspace.Update:input_type -> wspb.UpdateRequest
	6,  // 6: wspb.Workspace.Describe:input_type -> wspb.DescribeRequest
	8,  // 7: wspb.Workspace.Disable:input_type -> wspb.DisableRequest
	9,  // 8: wspb.Workspace.Enable:input_type -> wspb.EnableRequest
	2,  // 9: wspb.Workspace.Lists:output_type -> wspb.ListsReply
	10, // 10: wspb.Workspace.Create:output_type -> wspb.EmptyReply
	10, // 11: wspb.Workspace.Delete:output_type -> wspb.EmptyReply
	10, // 12: wspb.Workspace.Update:output_type -> wspb.EmptyReply
	7,  // 13: wspb.Workspace.Describe:output_type -> wspb.DescribeReply
	10, // 14: wspb.Workspace.Disable:output_type -> wspb.EmptyReply
	10, // 15: wspb.Workspace.Enable:output_type -> wspb.EmptyReply
	9,  // [9:16] is the sub-list for method output_type
	2,  // [2:9] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_workspace_proto_init() }
func file_proto_workspace_proto_init() {
	if File_proto_workspace_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_workspace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Info); i {
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
		file_proto_workspace_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListsRequest); i {
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
		file_proto_workspace_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListsReply); i {
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
		file_proto_workspace_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_proto_workspace_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_proto_workspace_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_proto_workspace_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeRequest); i {
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
		file_proto_workspace_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeReply); i {
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
		file_proto_workspace_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisableRequest); i {
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
		file_proto_workspace_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnableRequest); i {
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
		file_proto_workspace_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_workspace_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_workspace_proto_goTypes,
		DependencyIndexes: file_proto_workspace_proto_depIdxs,
		MessageInfos:      file_proto_workspace_proto_msgTypes,
	}.Build()
	File_proto_workspace_proto = out.File
	file_proto_workspace_proto_rawDesc = nil
	file_proto_workspace_proto_goTypes = nil
	file_proto_workspace_proto_depIdxs = nil
}
