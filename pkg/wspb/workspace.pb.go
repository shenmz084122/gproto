// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/workspace.proto

package wspb

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

// Request parameters used to ListWorkspaces
type ListWorkspacesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner  string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Limit  int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListWorkspacesRequest) Reset() {
	*x = ListWorkspacesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_workspace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWorkspacesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWorkspacesRequest) ProtoMessage() {}

func (x *ListWorkspacesRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListWorkspacesRequest.ProtoReflect.Descriptor instead.
func (*ListWorkspacesRequest) Descriptor() ([]byte, []int) {
	return file_proto_workspace_proto_rawDescGZIP(), []int{0}
}

func (x *ListWorkspacesRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *ListWorkspacesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListWorkspacesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

// Reply parameters used to ListWorkspaces
type ListWorkspacesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*model.SpaceInfo `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos,omitempty"`
}

func (x *ListWorkspacesReply) Reset() {
	*x = ListWorkspacesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_workspace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWorkspacesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWorkspacesReply) ProtoMessage() {}

func (x *ListWorkspacesReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListWorkspacesReply.ProtoReflect.Descriptor instead.
func (*ListWorkspacesReply) Descriptor() ([]byte, []int) {
	return file_proto_workspace_proto_rawDescGZIP(), []int{1}
}

func (x *ListWorkspacesReply) GetInfos() []*model.SpaceInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

// Request parameters used to CreateWorkspace
type CreateWorkspaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Desc  string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *CreateWorkspaceRequest) Reset() {
	*x = CreateWorkspaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_workspace_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWorkspaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWorkspaceRequest) ProtoMessage() {}

func (x *CreateWorkspaceRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateWorkspaceRequest.ProtoReflect.Descriptor instead.
func (*CreateWorkspaceRequest) Descriptor() ([]byte, []int) {
	return file_proto_workspace_proto_rawDescGZIP(), []int{2}
}

func (x *CreateWorkspaceRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *CreateWorkspaceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateWorkspaceRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

// Request parameters used to DeleteWorkspace
type DeleteWorkspaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteWorkspaceRequest) Reset() {
	*x = DeleteWorkspaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_workspace_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWorkspaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWorkspaceRequest) ProtoMessage() {}

func (x *DeleteWorkspaceRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeleteWorkspaceRequest.ProtoReflect.Descriptor instead.
func (*DeleteWorkspaceRequest) Descriptor() ([]byte, []int) {
	return file_proto_workspace_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteWorkspaceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Request parameters used to UpdateWorkspace
type UpdateWorkspaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Desc string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *UpdateWorkspaceRequest) Reset() {
	*x = UpdateWorkspaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_workspace_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateWorkspaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWorkspaceRequest) ProtoMessage() {}

func (x *UpdateWorkspaceRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UpdateWorkspaceRequest.ProtoReflect.Descriptor instead.
func (*UpdateWorkspaceRequest) Descriptor() ([]byte, []int) {
	return file_proto_workspace_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateWorkspaceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateWorkspaceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateWorkspaceRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

// Request parameters used to DescribeWorkspace
type DescribeWorkspaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DescribeWorkspaceRequest) Reset() {
	*x = DescribeWorkspaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_workspace_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeWorkspaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeWorkspaceRequest) ProtoMessage() {}

func (x *DescribeWorkspaceRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DescribeWorkspaceRequest.ProtoReflect.Descriptor instead.
func (*DescribeWorkspaceRequest) Descriptor() ([]byte, []int) {
	return file_proto_workspace_proto_rawDescGZIP(), []int{5}
}

func (x *DescribeWorkspaceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Reply parameters used to DescribeWorkspace
type DescribeWorkspaceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *model.SpaceInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *DescribeWorkspaceReply) Reset() {
	*x = DescribeWorkspaceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_workspace_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeWorkspaceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeWorkspaceReply) ProtoMessage() {}

func (x *DescribeWorkspaceReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DescribeWorkspaceReply.ProtoReflect.Descriptor instead.
func (*DescribeWorkspaceReply) Descriptor() ([]byte, []int) {
	return file_proto_workspace_proto_rawDescGZIP(), []int{6}
}

func (x *DescribeWorkspaceReply) GetInfo() *model.SpaceInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

// Request parameters used to DisableWorkspace
type DisableWorkspaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DisableWorkspaceRequest) Reset() {
	*x = DisableWorkspaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_workspace_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisableWorkspaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisableWorkspaceRequest) ProtoMessage() {}

func (x *DisableWorkspaceRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DisableWorkspaceRequest.ProtoReflect.Descriptor instead.
func (*DisableWorkspaceRequest) Descriptor() ([]byte, []int) {
	return file_proto_workspace_proto_rawDescGZIP(), []int{7}
}

func (x *DisableWorkspaceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Request parameters used to EnableWorkspace
type EnableWorkspaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EnableWorkspaceRequest) Reset() {
	*x = EnableWorkspaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_workspace_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnableWorkspaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnableWorkspaceRequest) ProtoMessage() {}

func (x *EnableWorkspaceRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use EnableWorkspaceRequest.ProtoReflect.Descriptor instead.
func (*EnableWorkspaceRequest) Descriptor() ([]byte, []int) {
	return file_proto_workspace_proto_rawDescGZIP(), []int{8}
}

func (x *EnableWorkspaceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Request parameters used to AddAudit
type AddAuditRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *model.AuditInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *AddAuditRequest) Reset() {
	*x = AddAuditRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_workspace_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAuditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAuditRequest) ProtoMessage() {}

func (x *AddAuditRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AddAuditRequest.ProtoReflect.Descriptor instead.
func (*AddAuditRequest) Descriptor() ([]byte, []int) {
	return file_proto_workspace_proto_rawDescGZIP(), []int{9}
}

func (x *AddAuditRequest) GetInfo() *model.AuditInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

// Request parameters used to ListAudits
type ListAuditsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpaceId string `protobuf:"bytes,1,opt,name=space_id,json=spaceId,proto3" json:"space_id,omitempty"`
	Limit   int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset  int32  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListAuditsRequest) Reset() {
	*x = ListAuditsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_workspace_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAuditsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuditsRequest) ProtoMessage() {}

func (x *ListAuditsRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListAuditsRequest.ProtoReflect.Descriptor instead.
func (*ListAuditsRequest) Descriptor() ([]byte, []int) {
	return file_proto_workspace_proto_rawDescGZIP(), []int{10}
}

func (x *ListAuditsRequest) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *ListAuditsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListAuditsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

// Reply parameters used to AddAudit
type ListAuditsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*model.AuditInfo `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos,omitempty"`
}

func (x *ListAuditsReply) Reset() {
	*x = ListAuditsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_workspace_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAuditsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuditsReply) ProtoMessage() {}

func (x *ListAuditsReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_workspace_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuditsReply.ProtoReflect.Descriptor instead.
func (*ListAuditsReply) Descriptor() ([]byte, []int) {
	return file_proto_workspace_proto_rawDescGZIP(), []int{11}
}

func (x *ListAuditsReply) GetInfos() []*model.AuditInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

var File_proto_workspace_proto protoreflect.FileDescriptor

var file_proto_workspace_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x77, 0x73, 0x70, 0x62, 0x1a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f,
	0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02,
	0x58, 0x01, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0xe2, 0xdf, 0x1f, 0x04, 0x10, 0x00,
	0x18, 0x65, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0f, 0xe2, 0xdf, 0x1f, 0x0b, 0x10,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x22, 0x45, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x6e, 0x66,
	0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02,
	0x20, 0x00, 0x52, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x74, 0x0a, 0x16, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0xe2, 0xdf, 0x1f, 0x04, 0x70, 0x0a, 0x78, 0x41, 0x52, 0x05, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x09, 0xe2, 0xdf, 0x1f, 0x05, 0x70, 0x01, 0x78, 0x81, 0x01, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x78, 0x81, 0x08, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x22,
	0x31, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01, 0x14, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x6b, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01,
	0x14, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x78, 0x81, 0x01, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x78, 0x81, 0x08, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x22,
	0x33, 0x0a, 0x18, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01, 0x14,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x16, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x24,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x22, 0x32, 0x0a, 0x17, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x57,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f,
	0x03, 0x80, 0x01, 0x14, 0x52, 0x02, 0x69, 0x64, 0x22, 0x31, 0x0a, 0x16, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01, 0x14, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x0f, 0x41,
	0x64, 0x64, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x06,
	0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x80, 0x01, 0x0a,
	0x11, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x22, 0x0a, 0x08, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01, 0x14, 0x52, 0x07, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0xe2, 0xdf, 0x1f, 0x04, 0x10, 0x00, 0x18, 0x65, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0f, 0xe2, 0xdf, 0x1f, 0x0b, 0x10, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22,
	0x41, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x00, 0x52, 0x05, 0x69, 0x6e, 0x66,
	0x6f, 0x73, 0x32, 0x8a, 0x05, 0x0a, 0x09, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x4a, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x12, 0x1b, 0x2e, 0x77, 0x73, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x77, 0x73, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x1c, 0x2e, 0x77, 0x73, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1c, 0x2e, 0x77, 0x73, 0x70, 0x62, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0f, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1c, 0x2e,
	0x77, 0x73, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22,
	0x00, 0x12, 0x53, 0x0a, 0x11, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1e, 0x2e, 0x77, 0x73, 0x70, 0x62, 0x2e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x77, 0x73, 0x70, 0x62, 0x2e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1d, 0x2e, 0x77, 0x73, 0x70,
	0x62, 0x2e, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12,
	0x45, 0x0a, 0x0f, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x1c, 0x2e, 0x77, 0x73, 0x70, 0x62, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x41, 0x75, 0x64,
	0x69, 0x74, 0x12, 0x15, 0x2e, 0x77, 0x73, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x75, 0x64,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12,
	0x3e, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x73, 0x12, 0x17, 0x2e,
	0x77, 0x73, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x77, 0x73, 0x70, 0x62, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42,
	0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61,
	0x74, 0x61, 0x57, 0x6f, 0x72, 0x6b, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2f, 0x67, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x77, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
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

var file_proto_workspace_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_workspace_proto_goTypes = []interface{}{
	(*ListWorkspacesRequest)(nil),    // 0: wspb.ListWorkspacesRequest
	(*ListWorkspacesReply)(nil),      // 1: wspb.ListWorkspacesReply
	(*CreateWorkspaceRequest)(nil),   // 2: wspb.CreateWorkspaceRequest
	(*DeleteWorkspaceRequest)(nil),   // 3: wspb.DeleteWorkspaceRequest
	(*UpdateWorkspaceRequest)(nil),   // 4: wspb.UpdateWorkspaceRequest
	(*DescribeWorkspaceRequest)(nil), // 5: wspb.DescribeWorkspaceRequest
	(*DescribeWorkspaceReply)(nil),   // 6: wspb.DescribeWorkspaceReply
	(*DisableWorkspaceRequest)(nil),  // 7: wspb.DisableWorkspaceRequest
	(*EnableWorkspaceRequest)(nil),   // 8: wspb.EnableWorkspaceRequest
	(*AddAuditRequest)(nil),          // 9: wspb.AddAuditRequest
	(*ListAuditsRequest)(nil),        // 10: wspb.ListAuditsRequest
	(*ListAuditsReply)(nil),          // 11: wspb.ListAuditsReply
	(*model.SpaceInfo)(nil),          // 12: model.SpaceInfo
	(*model.AuditInfo)(nil),          // 13: model.AuditInfo
	(*model.EmptyStruct)(nil),        // 14: model.EmptyStruct
}
var file_proto_workspace_proto_depIdxs = []int32{
	12, // 0: wspb.ListWorkspacesReply.infos:type_name -> model.SpaceInfo
	12, // 1: wspb.DescribeWorkspaceReply.info:type_name -> model.SpaceInfo
	13, // 2: wspb.AddAuditRequest.info:type_name -> model.AuditInfo
	13, // 3: wspb.ListAuditsReply.infos:type_name -> model.AuditInfo
	0,  // 4: wspb.Workspace.ListWorkspaces:input_type -> wspb.ListWorkspacesRequest
	2,  // 5: wspb.Workspace.CreateWorkspace:input_type -> wspb.CreateWorkspaceRequest
	3,  // 6: wspb.Workspace.DeleteWorkspace:input_type -> wspb.DeleteWorkspaceRequest
	4,  // 7: wspb.Workspace.UpdateWorkspace:input_type -> wspb.UpdateWorkspaceRequest
	5,  // 8: wspb.Workspace.DescribeWorkspace:input_type -> wspb.DescribeWorkspaceRequest
	7,  // 9: wspb.Workspace.DisableWorkspace:input_type -> wspb.DisableWorkspaceRequest
	8,  // 10: wspb.Workspace.EnableWorkspace:input_type -> wspb.EnableWorkspaceRequest
	9,  // 11: wspb.Workspace.AddAudit:input_type -> wspb.AddAuditRequest
	10, // 12: wspb.Workspace.ListAudits:input_type -> wspb.ListAuditsRequest
	1,  // 13: wspb.Workspace.ListWorkspaces:output_type -> wspb.ListWorkspacesReply
	14, // 14: wspb.Workspace.CreateWorkspace:output_type -> model.EmptyStruct
	14, // 15: wspb.Workspace.DeleteWorkspace:output_type -> model.EmptyStruct
	14, // 16: wspb.Workspace.UpdateWorkspace:output_type -> model.EmptyStruct
	6,  // 17: wspb.Workspace.DescribeWorkspace:output_type -> wspb.DescribeWorkspaceReply
	14, // 18: wspb.Workspace.DisableWorkspace:output_type -> model.EmptyStruct
	14, // 19: wspb.Workspace.EnableWorkspace:output_type -> model.EmptyStruct
	14, // 20: wspb.Workspace.AddAudit:output_type -> model.EmptyStruct
	11, // 21: wspb.Workspace.ListAudits:output_type -> wspb.ListAuditsReply
	13, // [13:22] is the sub-list for method output_type
	4,  // [4:13] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_workspace_proto_init() }
func file_proto_workspace_proto_init() {
	if File_proto_workspace_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_workspace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWorkspacesRequest); i {
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
			switch v := v.(*ListWorkspacesReply); i {
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
			switch v := v.(*CreateWorkspaceRequest); i {
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
			switch v := v.(*DeleteWorkspaceRequest); i {
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
			switch v := v.(*UpdateWorkspaceRequest); i {
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
			switch v := v.(*DescribeWorkspaceRequest); i {
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
			switch v := v.(*DescribeWorkspaceReply); i {
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
			switch v := v.(*DisableWorkspaceRequest); i {
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
			switch v := v.(*EnableWorkspaceRequest); i {
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
			switch v := v.(*AddAuditRequest); i {
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
			switch v := v.(*ListAuditsRequest); i {
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
		file_proto_workspace_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAuditsReply); i {
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
			NumMessages:   12,
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
