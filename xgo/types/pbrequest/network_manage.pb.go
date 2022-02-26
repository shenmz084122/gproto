// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: proto/types/request/network_manage.proto

package pbrequest

import (
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbdefaults"
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

type ListNetworks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The workspace id in HTTP Request-URI. Is Required.
	SpaceId string `protobuf:"bytes,1,opt,name=space_id,json=spaceId,proto3" json:"space_id" uri:"space_id" swaggerignore:"true"`
	// Limit the maximum number of entries returned this time.
	// Not required, Max 100, default 100.
	Limit int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit" form:"limit"`
	// The offset position. Not required, default 0.
	Offset int32 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset" form:"offset"`
	// The field list used to sorted query results.
	// Optional values: {id, name, created, updated}.
	// Multiple fields are separated by commas(","), eg: sort_by="updated".
	// Not required, default: id.
	SortBy string `protobuf:"bytes,4,opt,name=sort_by,json=sortBy,proto3" json:"sort_by" form:"sort_by"`
	// Reverse order results. Not required, default: false.
	Reverse bool `protobuf:"varint,5,opt,name=reverse,proto3" json:"reverse" form:"reverse"`
	// Search with workspace name; Not required.
	Search string `protobuf:"bytes,6,opt,name=search,proto3" json:"search" form:"search"`
	// Filter by network name; valid if `search` is empty; Not required.
	Name string `protobuf:"bytes,7,opt,name=name,proto3" json:"name" form:"name"`
}

func (x *ListNetworks) Reset() {
	*x = ListNetworks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_request_network_manage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNetworks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNetworks) ProtoMessage() {}

func (x *ListNetworks) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_request_network_manage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNetworks.ProtoReflect.Descriptor instead.
func (*ListNetworks) Descriptor() ([]byte, []int) {
	return file_proto_types_request_network_manage_proto_rawDescGZIP(), []int{0}
}

func (x *ListNetworks) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *ListNetworks) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListNetworks) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListNetworks) GetSortBy() string {
	if x != nil {
		return x.SortBy
	}
	return ""
}

func (x *ListNetworks) GetReverse() bool {
	if x != nil {
		return x.Reverse
	}
	return false
}

func (x *ListNetworks) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *ListNetworks) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteNetworks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The workspace id in HTTP Request-URI
	SpaceId string `protobuf:"bytes,1,opt,name=space_id,json=spaceId,proto3" json:"space_id" uri:"space_id" swaggerignore:"true"`
	// The list of network id. Is required, Min 1, Max 100.
	NetworkIds []string `protobuf:"bytes,2,rep,name=network_ids,json=networkIds,proto3" json:"network_ids"`
}

func (x *DeleteNetworks) Reset() {
	*x = DeleteNetworks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_request_network_manage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNetworks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNetworks) ProtoMessage() {}

func (x *DeleteNetworks) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_request_network_manage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNetworks.ProtoReflect.Descriptor instead.
func (*DeleteNetworks) Descriptor() ([]byte, []int) {
	return file_proto_types_request_network_manage_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteNetworks) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *DeleteNetworks) GetNetworkIds() []string {
	if x != nil {
		return x.NetworkIds
	}
	return nil
}

type CreateNetwork struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The workspace id in HTTP Request-URI
	SpaceId string `protobuf:"bytes,1,opt,name=space_id,json=spaceId,proto3" json:"space_id" uri:"space_id" swaggerignore:"true"`
	// The user-id of created this cluster. It fixed to request user id.
	CreatedBy string `protobuf:"bytes,2,opt,name=created_by,json=createdBy,proto3" json:"created_by" swaggerignore:"true"`
	// The owner of workspace, only used to check quota. Set by APIServer.
	SpaceOwner string `protobuf:"bytes,3,opt,name=space_owner,json=spaceOwner,proto3" json:"space_owner" swaggerignore:"true"`
	// Network Name. Is required.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`
	// VPC's route id, Is required.
	RouterId string `protobuf:"bytes,5,opt,name=router_id,json=routerId,proto3" json:"router_id"`
	// vxnet id. Is required.
	VxnetId string `protobuf:"bytes,6,opt,name=vxnet_id,json=vxnetId,proto3" json:"vxnet_id"`
}

func (x *CreateNetwork) Reset() {
	*x = CreateNetwork{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_request_network_manage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNetwork) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNetwork) ProtoMessage() {}

func (x *CreateNetwork) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_request_network_manage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNetwork.ProtoReflect.Descriptor instead.
func (*CreateNetwork) Descriptor() ([]byte, []int) {
	return file_proto_types_request_network_manage_proto_rawDescGZIP(), []int{2}
}

func (x *CreateNetwork) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *CreateNetwork) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *CreateNetwork) GetSpaceOwner() string {
	if x != nil {
		return x.SpaceOwner
	}
	return ""
}

func (x *CreateNetwork) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateNetwork) GetRouterId() string {
	if x != nil {
		return x.RouterId
	}
	return ""
}

func (x *CreateNetwork) GetVxnetId() string {
	if x != nil {
		return x.VxnetId
	}
	return ""
}

type UpdateNetwork struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The workspace id in HTTP Request-URI
	SpaceId string `protobuf:"bytes,1,opt,name=space_id,json=spaceId,proto3" json:"space_id" uri:"space_id" swaggerignore:"true"`
	// The flink cluster id in HTTP Request-URI
	NetworkId string `protobuf:"bytes,2,opt,name=network_id,json=networkId,proto3" json:"network_id" uri:"network_id" swaggerignore:"true"`
	// Network Name. Is required.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	// VPC's route id, Is required.
	RouterId string `protobuf:"bytes,4,opt,name=router_id,json=routerId,proto3" json:"router_id"`
	// vxnet id. Is required.
	VxnetId string `protobuf:"bytes,5,opt,name=vxnet_id,json=vxnetId,proto3" json:"vxnet_id"`
}

func (x *UpdateNetwork) Reset() {
	*x = UpdateNetwork{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_request_network_manage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNetwork) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNetwork) ProtoMessage() {}

func (x *UpdateNetwork) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_request_network_manage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNetwork.ProtoReflect.Descriptor instead.
func (*UpdateNetwork) Descriptor() ([]byte, []int) {
	return file_proto_types_request_network_manage_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateNetwork) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *UpdateNetwork) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

func (x *UpdateNetwork) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateNetwork) GetRouterId() string {
	if x != nil {
		return x.RouterId
	}
	return ""
}

func (x *UpdateNetwork) GetVxnetId() string {
	if x != nil {
		return x.VxnetId
	}
	return ""
}

type DescribeNetwork struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The workspace id in HTTP Request-URI
	NetworkId string `protobuf:"bytes,1,opt,name=network_id,json=networkId,proto3" json:"network_id" uri:"network_id" swaggerignore:"true"`
}

func (x *DescribeNetwork) Reset() {
	*x = DescribeNetwork{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_request_network_manage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeNetwork) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeNetwork) ProtoMessage() {}

func (x *DescribeNetwork) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_request_network_manage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeNetwork.ProtoReflect.Descriptor instead.
func (*DescribeNetwork) Descriptor() ([]byte, []int) {
	return file_proto_types_request_network_manage_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeNetwork) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

var File_proto_types_request_network_manage_proto protoreflect.FileDescriptor

var file_proto_types_request_network_manage_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x79, 0x75, 0x33, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x75, 0x33, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x02, 0x0a,
	0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x12, 0x2e, 0x0a,
	0x08, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x13, 0xe2, 0xdf, 0x1f, 0x0f, 0x12, 0x0d, 0xc2, 0x01, 0x0a, 0xf0, 0x01, 0x14, 0xca, 0x02, 0x04,
	0x77, 0x6b, 0x73, 0x2d, 0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2d, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x17, 0xa2, 0xa1,
	0x1f, 0x06, 0xaa, 0x06, 0x03, 0x31, 0x30, 0x30, 0xe2, 0xdf, 0x1f, 0x09, 0x12, 0x07, 0xb2, 0x01,
	0x04, 0x30, 0x00, 0x38, 0x64, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x23, 0x0a, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0b, 0xe2, 0xdf,
	0x1f, 0x07, 0x12, 0x05, 0xb2, 0x01, 0x02, 0x40, 0x00, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x40, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x27, 0xe2, 0xdf, 0x1f, 0x23, 0x12, 0x21, 0xc2, 0x01, 0x1e, 0x4a, 0x00, 0x4a,
	0x02, 0x69, 0x64, 0x4a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4a, 0x07, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x4a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x72,
	0x74, 0x42, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x7c, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x12, 0x2e, 0x0a, 0x08, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0xe2,
	0xdf, 0x1f, 0x0f, 0x12, 0x0d, 0xc2, 0x01, 0x0a, 0xf0, 0x01, 0x14, 0xca, 0x02, 0x04, 0x77, 0x6b,
	0x73, 0x2d, 0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0b, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x42, 0x19, 0xe2, 0xdf, 0x1f, 0x15, 0x12, 0x13, 0xea, 0x01, 0x10, 0x30, 0x00, 0x38, 0x64, 0x5a,
	0x0a, 0xc2, 0x01, 0x07, 0xca, 0x02, 0x04, 0x6e, 0x65, 0x74, 0x2d, 0x52, 0x0a, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x73, 0x22, 0xa1, 0x02, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0xe2, 0xdf, 0x1f,
	0x0f, 0x12, 0x0d, 0xc2, 0x01, 0x0a, 0xf0, 0x01, 0x14, 0xca, 0x02, 0x04, 0x77, 0x6b, 0x73, 0x2d,
	0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xe2,
	0xdf, 0x1f, 0x0b, 0x12, 0x09, 0xc2, 0x01, 0x06, 0x80, 0x02, 0x00, 0x88, 0x02, 0x41, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x2c, 0x0a, 0x0b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b,
	0xe2, 0xdf, 0x1f, 0x07, 0x12, 0x05, 0xc2, 0x01, 0x02, 0x22, 0x00, 0x52, 0x0a, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xe2, 0xdf, 0x1f, 0x0c, 0x12, 0x0a, 0xc2, 0x01, 0x07,
	0x80, 0x02, 0x01, 0x98, 0x02, 0x80, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a,
	0x09, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x10, 0xe2, 0xdf, 0x1f, 0x0c, 0x12, 0x0a, 0xc2, 0x01, 0x07, 0xca, 0x02, 0x04, 0x72, 0x74,
	0x72, 0x2d, 0x52, 0x08, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x08,
	0x76, 0x78, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x12,
	0xe2, 0xdf, 0x1f, 0x0e, 0x12, 0x0c, 0xc2, 0x01, 0x09, 0xca, 0x02, 0x06, 0x76, 0x78, 0x6e, 0x65,
	0x74, 0x2d, 0x52, 0x07, 0x76, 0x78, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x22, 0xf7, 0x01, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x2e, 0x0a,
	0x08, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x13, 0xe2, 0xdf, 0x1f, 0x0f, 0x12, 0x0d, 0xc2, 0x01, 0x0a, 0xf0, 0x01, 0x14, 0xca, 0x02, 0x04,
	0x77, 0x6b, 0x73, 0x2d, 0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x32, 0x0a,
	0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x13, 0xe2, 0xdf, 0x1f, 0x0f, 0x12, 0x0d, 0xc2, 0x01, 0x0a, 0xf0, 0x01, 0x14, 0xca,
	0x02, 0x04, 0x6e, 0x65, 0x74, 0x2d, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49,
	0x64, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x10, 0xe2, 0xdf, 0x1f, 0x0c, 0x12, 0x0a, 0xc2, 0x01, 0x07, 0x80, 0x02, 0x01, 0x98, 0x02, 0x80,
	0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xe2, 0xdf, 0x1f, 0x0c,
	0x12, 0x0a, 0xc2, 0x01, 0x07, 0xca, 0x02, 0x04, 0x72, 0x74, 0x72, 0x2d, 0x52, 0x08, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x08, 0x76, 0x78, 0x6e, 0x65, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x12, 0xe2, 0xdf, 0x1f, 0x0e, 0x12, 0x0c,
	0xc2, 0x01, 0x09, 0xca, 0x02, 0x06, 0x76, 0x78, 0x6e, 0x65, 0x74, 0x2d, 0x52, 0x07, 0x76, 0x78,
	0x6e, 0x65, 0x74, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x32, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0xe2, 0xdf,
	0x1f, 0x0f, 0x12, 0x0d, 0xc2, 0x01, 0x0a, 0xf0, 0x01, 0x14, 0xca, 0x02, 0x04, 0x6e, 0x65, 0x74,
	0x2d, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x42, 0x75, 0x0a, 0x24,
	0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6f, 0x6d, 0x6e, 0x69, 0x73, 0x2e, 0x67, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x42, 0x16, 0x50, 0x42, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x50, 0x00, 0x5a, 0x33,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x57,
	0x6f, 0x72, 0x6b, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2f, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x78, 0x67, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_types_request_network_manage_proto_rawDescOnce sync.Once
	file_proto_types_request_network_manage_proto_rawDescData = file_proto_types_request_network_manage_proto_rawDesc
)

func file_proto_types_request_network_manage_proto_rawDescGZIP() []byte {
	file_proto_types_request_network_manage_proto_rawDescOnce.Do(func() {
		file_proto_types_request_network_manage_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_request_network_manage_proto_rawDescData)
	})
	return file_proto_types_request_network_manage_proto_rawDescData
}

var file_proto_types_request_network_manage_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_types_request_network_manage_proto_goTypes = []interface{}{
	(*ListNetworks)(nil),    // 0: request.ListNetworks
	(*DeleteNetworks)(nil),  // 1: request.DeleteNetworks
	(*CreateNetwork)(nil),   // 2: request.CreateNetwork
	(*UpdateNetwork)(nil),   // 3: request.UpdateNetwork
	(*DescribeNetwork)(nil), // 4: request.DescribeNetwork
}
var file_proto_types_request_network_manage_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_types_request_network_manage_proto_init() }
func file_proto_types_request_network_manage_proto_init() {
	if File_proto_types_request_network_manage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_types_request_network_manage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNetworks); i {
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
		file_proto_types_request_network_manage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNetworks); i {
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
		file_proto_types_request_network_manage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNetwork); i {
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
		file_proto_types_request_network_manage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNetwork); i {
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
		file_proto_types_request_network_manage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeNetwork); i {
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
			RawDescriptor: file_proto_types_request_network_manage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_request_network_manage_proto_goTypes,
		DependencyIndexes: file_proto_types_request_network_manage_proto_depIdxs,
		MessageInfos:      file_proto_types_request_network_manage_proto_msgTypes,
	}.Build()
	File_proto_types_request_network_manage_proto = out.File
	file_proto_types_request_network_manage_proto_rawDesc = nil
	file_proto_types_request_network_manage_proto_goTypes = nil
	file_proto_types_request_network_manage_proto_depIdxs = nil
}
