// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: proto/types/request/engine_manage.proto

package pbrequest

import (
	pbmodel "github.com/DataWorkbench/gproto/xgo/types/pbmodel"
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

type CreateFlinkClusterInK8S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Cluster is the metadata info of flink cluster.
	Info *pbmodel.FlinkCluster `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *CreateFlinkClusterInK8S) Reset() {
	*x = CreateFlinkClusterInK8S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_request_engine_manage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFlinkClusterInK8S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFlinkClusterInK8S) ProtoMessage() {}

func (x *CreateFlinkClusterInK8S) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_request_engine_manage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFlinkClusterInK8S.ProtoReflect.Descriptor instead.
func (*CreateFlinkClusterInK8S) Descriptor() ([]byte, []int) {
	return file_proto_types_request_engine_manage_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFlinkClusterInK8S) GetInfo() *pbmodel.FlinkCluster {
	if x != nil {
		return x.Info
	}
	return nil
}

type DeleteFlinkClusterInK8S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Cluster is the metadata info of flink cluster.
	Info *pbmodel.FlinkCluster `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *DeleteFlinkClusterInK8S) Reset() {
	*x = DeleteFlinkClusterInK8S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_request_engine_manage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFlinkClusterInK8S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFlinkClusterInK8S) ProtoMessage() {}

func (x *DeleteFlinkClusterInK8S) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_request_engine_manage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFlinkClusterInK8S.ProtoReflect.Descriptor instead.
func (*DeleteFlinkClusterInK8S) Descriptor() ([]byte, []int) {
	return file_proto_types_request_engine_manage_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteFlinkClusterInK8S) GetInfo() *pbmodel.FlinkCluster {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateNetworkBrokerInK8S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Cluster is the metadata info of flink cluster.
	Info *pbmodel.Network `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *CreateNetworkBrokerInK8S) Reset() {
	*x = CreateNetworkBrokerInK8S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_request_engine_manage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNetworkBrokerInK8S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNetworkBrokerInK8S) ProtoMessage() {}

func (x *CreateNetworkBrokerInK8S) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_request_engine_manage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNetworkBrokerInK8S.ProtoReflect.Descriptor instead.
func (*CreateNetworkBrokerInK8S) Descriptor() ([]byte, []int) {
	return file_proto_types_request_engine_manage_proto_rawDescGZIP(), []int{2}
}

func (x *CreateNetworkBrokerInK8S) GetInfo() *pbmodel.Network {
	if x != nil {
		return x.Info
	}
	return nil
}

type DeleteNetworkBrokerInK8S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Cluster is the metadata info of flink cluster.
	Info *pbmodel.Network `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *DeleteNetworkBrokerInK8S) Reset() {
	*x = DeleteNetworkBrokerInK8S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_request_engine_manage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNetworkBrokerInK8S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNetworkBrokerInK8S) ProtoMessage() {}

func (x *DeleteNetworkBrokerInK8S) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_request_engine_manage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNetworkBrokerInK8S.ProtoReflect.Descriptor instead.
func (*DeleteNetworkBrokerInK8S) Descriptor() ([]byte, []int) {
	return file_proto_types_request_engine_manage_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteNetworkBrokerInK8S) GetInfo() *pbmodel.Network {
	if x != nil {
		return x.Info
	}
	return nil
}

type DeleteNamespacesInK8S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpaceIds []string `protobuf:"bytes,1,rep,name=space_ids,json=spaceIds,proto3" json:"space_ids"`
}

func (x *DeleteNamespacesInK8S) Reset() {
	*x = DeleteNamespacesInK8S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_request_engine_manage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNamespacesInK8S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNamespacesInK8S) ProtoMessage() {}

func (x *DeleteNamespacesInK8S) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_request_engine_manage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNamespacesInK8S.ProtoReflect.Descriptor instead.
func (*DeleteNamespacesInK8S) Descriptor() ([]byte, []int) {
	return file_proto_types_request_engine_manage_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteNamespacesInK8S) GetSpaceIds() []string {
	if x != nil {
		return x.SpaceIds
	}
	return nil
}

type CreateFlinkClusterInK8SV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Cluster is the metadata info of flink cluster.
	Info *pbmodel.FlinkCluster `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *CreateFlinkClusterInK8SV2) Reset() {
	*x = CreateFlinkClusterInK8SV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_request_engine_manage_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFlinkClusterInK8SV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFlinkClusterInK8SV2) ProtoMessage() {}

func (x *CreateFlinkClusterInK8SV2) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_request_engine_manage_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFlinkClusterInK8SV2.ProtoReflect.Descriptor instead.
func (*CreateFlinkClusterInK8SV2) Descriptor() ([]byte, []int) {
	return file_proto_types_request_engine_manage_proto_rawDescGZIP(), []int{5}
}

func (x *CreateFlinkClusterInK8SV2) GetInfo() *pbmodel.FlinkCluster {
	if x != nil {
		return x.Info
	}
	return nil
}

type DeleteFlinkClusterInK8SV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Workspace ID it belongs to.
	SpaceId   string `protobuf:"bytes,1,opt,name=space_id,json=spaceId,proto3" json:"space_id"`
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id"`
	// vxnet id. Is required.
	VxnetId string `protobuf:"bytes,3,opt,name=vxnet_id,json=vxnetId,proto3" json:"vxnet_id"`
}

func (x *DeleteFlinkClusterInK8SV2) Reset() {
	*x = DeleteFlinkClusterInK8SV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_request_engine_manage_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFlinkClusterInK8SV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFlinkClusterInK8SV2) ProtoMessage() {}

func (x *DeleteFlinkClusterInK8SV2) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_request_engine_manage_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFlinkClusterInK8SV2.ProtoReflect.Descriptor instead.
func (*DeleteFlinkClusterInK8SV2) Descriptor() ([]byte, []int) {
	return file_proto_types_request_engine_manage_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteFlinkClusterInK8SV2) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *DeleteFlinkClusterInK8SV2) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *DeleteFlinkClusterInK8SV2) GetVxnetId() string {
	if x != nil {
		return x.VxnetId
	}
	return ""
}

type CreateNetworkBrokerInK8SV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Workspace ID it belongs to.
	SpaceId string `protobuf:"bytes,1,opt,name=space_id,json=spaceId,proto3" json:"space_id"`
	// VPC's route id, Is required.
	RouterId string `protobuf:"bytes,2,opt,name=router_id,json=routerId,proto3" json:"router_id"`
	// vxnet id. Is required.
	VxnetId string `protobuf:"bytes,3,opt,name=vxnet_id,json=vxnetId,proto3" json:"vxnet_id"`
}

func (x *CreateNetworkBrokerInK8SV2) Reset() {
	*x = CreateNetworkBrokerInK8SV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_request_engine_manage_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNetworkBrokerInK8SV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNetworkBrokerInK8SV2) ProtoMessage() {}

func (x *CreateNetworkBrokerInK8SV2) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_request_engine_manage_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNetworkBrokerInK8SV2.ProtoReflect.Descriptor instead.
func (*CreateNetworkBrokerInK8SV2) Descriptor() ([]byte, []int) {
	return file_proto_types_request_engine_manage_proto_rawDescGZIP(), []int{7}
}

func (x *CreateNetworkBrokerInK8SV2) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *CreateNetworkBrokerInK8SV2) GetRouterId() string {
	if x != nil {
		return x.RouterId
	}
	return ""
}

func (x *CreateNetworkBrokerInK8SV2) GetVxnetId() string {
	if x != nil {
		return x.VxnetId
	}
	return ""
}

type DeleteNetworkBrokerInK8SV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Workspace ID it belongs to.
	SpaceId string `protobuf:"bytes,1,opt,name=space_id,json=spaceId,proto3" json:"space_id"`
	// VPC's route id, Is required.
	RouterId string `protobuf:"bytes,2,opt,name=router_id,json=routerId,proto3" json:"router_id"`
	// vxnet id. Is required.
	VxnetId string `protobuf:"bytes,3,opt,name=vxnet_id,json=vxnetId,proto3" json:"vxnet_id"`
}

func (x *DeleteNetworkBrokerInK8SV2) Reset() {
	*x = DeleteNetworkBrokerInK8SV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_request_engine_manage_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNetworkBrokerInK8SV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNetworkBrokerInK8SV2) ProtoMessage() {}

func (x *DeleteNetworkBrokerInK8SV2) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_request_engine_manage_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNetworkBrokerInK8SV2.ProtoReflect.Descriptor instead.
func (*DeleteNetworkBrokerInK8SV2) Descriptor() ([]byte, []int) {
	return file_proto_types_request_engine_manage_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteNetworkBrokerInK8SV2) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *DeleteNetworkBrokerInK8SV2) GetRouterId() string {
	if x != nil {
		return x.RouterId
	}
	return ""
}

func (x *DeleteNetworkBrokerInK8SV2) GetVxnetId() string {
	if x != nil {
		return x.VxnetId
	}
	return ""
}

var File_proto_types_request_engine_manage_proto protoreflect.FileDescriptor

var file_proto_types_request_engine_manage_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79,
	0x75, 0x33, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x17, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x6e, 0x4b, 0x38, 0x73, 0x12, 0x34, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x46, 0x6c, 0x69, 0x6e, 0x6b,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x42, 0x0b, 0xe2, 0xdf, 0x1f, 0x07, 0x12, 0x05, 0xe2,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x4f, 0x0a, 0x17, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x6e, 0x4b, 0x38, 0x73, 0x12, 0x34, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x46, 0x6c, 0x69, 0x6e,
	0x6b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x42, 0x0b, 0xe2, 0xdf, 0x1f, 0x07, 0x12, 0x05,
	0xe2, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x4b, 0x0a, 0x18, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x49, 0x6e, 0x4b, 0x38, 0x73, 0x12, 0x2f, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x0b, 0xe2, 0xdf, 0x1f, 0x07, 0x12, 0x05, 0xe2, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x4b, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x49,
	0x6e, 0x4b, 0x38, 0x73, 0x12, 0x2f, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x42, 0x0b, 0xe2, 0xdf, 0x1f, 0x07, 0x12, 0x05, 0xe2, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x50, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x49, 0x6e, 0x4b, 0x38, 0x73, 0x12, 0x37,
	0x0a, 0x09, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x42, 0x1a, 0xe2, 0xdf, 0x1f, 0x16, 0x12, 0x14, 0xea, 0x01, 0x11, 0x50, 0x01, 0x5a, 0x0d,
	0xc2, 0x01, 0x0a, 0xf0, 0x01, 0x14, 0xca, 0x02, 0x04, 0x77, 0x6b, 0x73, 0x2d, 0x52, 0x08, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x73, 0x22, 0x51, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x4b,
	0x38, 0x73, 0x56, 0x32, 0x12, 0x34, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x46, 0x6c, 0x69, 0x6e, 0x6b,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x42, 0x0b, 0xe2, 0xdf, 0x1f, 0x07, 0x12, 0x05, 0xe2,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0xc5, 0x01, 0x0a, 0x19, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x6e, 0x6b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x6e, 0x4b, 0x38, 0x73, 0x56, 0x32, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0xe2, 0xdf, 0x1f, 0x0f,
	0x12, 0x0d, 0xc2, 0x01, 0x0a, 0xf0, 0x01, 0x14, 0xca, 0x02, 0x04, 0x77, 0x6b, 0x73, 0x2d, 0x52,
	0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0xe2, 0xdf,
	0x1f, 0x0f, 0x12, 0x0d, 0xc2, 0x01, 0x0a, 0xf0, 0x01, 0x14, 0xca, 0x02, 0x04, 0x63, 0x66, 0x69,
	0x2d, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x08,
	0x76, 0x78, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29,
	0xe2, 0xdf, 0x1f, 0x13, 0x0a, 0x11, 0x0a, 0x08, 0x76, 0x78, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64,
	0x12, 0x05, 0xc2, 0x01, 0x02, 0x22, 0x00, 0xe2, 0xdf, 0x1f, 0x0e, 0x12, 0x0c, 0xc2, 0x01, 0x09,
	0xca, 0x02, 0x06, 0x76, 0x78, 0x6e, 0x65, 0x74, 0x2d, 0x52, 0x07, 0x76, 0x78, 0x6e, 0x65, 0x74,
	0x49, 0x64, 0x22, 0xaa, 0x01, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x49, 0x6e, 0x4b, 0x38, 0x73, 0x56,
	0x32, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x13, 0xe2, 0xdf, 0x1f, 0x0f, 0x12, 0x0d, 0xc2, 0x01, 0x0a, 0xf0, 0x01,
	0x14, 0xca, 0x02, 0x04, 0x77, 0x6b, 0x73, 0x2d, 0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x2d, 0x0a, 0x09, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0xe2, 0xdf, 0x1f, 0x0c, 0x12, 0x0a, 0xc2, 0x01, 0x07, 0xca,
	0x02, 0x04, 0x72, 0x74, 0x72, 0x2d, 0x52, 0x08, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x2d, 0x0a, 0x08, 0x76, 0x78, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x12, 0xe2, 0xdf, 0x1f, 0x0e, 0x12, 0x0c, 0xc2, 0x01, 0x09, 0xca, 0x02, 0x06,
	0x76, 0x78, 0x6e, 0x65, 0x74, 0x2d, 0x52, 0x07, 0x76, 0x78, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x22,
	0xaa, 0x01, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x49, 0x6e, 0x4b, 0x38, 0x73, 0x56, 0x32, 0x12, 0x2e,
	0x0a, 0x08, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x13, 0xe2, 0xdf, 0x1f, 0x0f, 0x12, 0x0d, 0xc2, 0x01, 0x0a, 0xf0, 0x01, 0x14, 0xca, 0x02,
	0x04, 0x77, 0x6b, 0x73, 0x2d, 0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2d,
	0x0a, 0x09, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x10, 0xe2, 0xdf, 0x1f, 0x0c, 0x12, 0x0a, 0xc2, 0x01, 0x07, 0xca, 0x02, 0x04, 0x72,
	0x74, 0x72, 0x2d, 0x52, 0x08, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2d, 0x0a,
	0x08, 0x76, 0x78, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x12, 0xe2, 0xdf, 0x1f, 0x0e, 0x12, 0x0c, 0xc2, 0x01, 0x09, 0xca, 0x02, 0x06, 0x76, 0x78, 0x6e,
	0x65, 0x74, 0x2d, 0x52, 0x07, 0x76, 0x78, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x42, 0x74, 0x0a, 0x24,
	0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6f, 0x6d, 0x6e, 0x69, 0x73, 0x2e, 0x67, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x42, 0x15, 0x50, 0x42, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x50, 0x00, 0x5a, 0x33, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x57, 0x6f,
	0x72, 0x6b, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2f, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x78,
	0x67, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_types_request_engine_manage_proto_rawDescOnce sync.Once
	file_proto_types_request_engine_manage_proto_rawDescData = file_proto_types_request_engine_manage_proto_rawDesc
)

func file_proto_types_request_engine_manage_proto_rawDescGZIP() []byte {
	file_proto_types_request_engine_manage_proto_rawDescOnce.Do(func() {
		file_proto_types_request_engine_manage_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_request_engine_manage_proto_rawDescData)
	})
	return file_proto_types_request_engine_manage_proto_rawDescData
}

var file_proto_types_request_engine_manage_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_types_request_engine_manage_proto_goTypes = []interface{}{
	(*CreateFlinkClusterInK8S)(nil),    // 0: request.CreateFlinkClusterInK8s
	(*DeleteFlinkClusterInK8S)(nil),    // 1: request.DeleteFlinkClusterInK8s
	(*CreateNetworkBrokerInK8S)(nil),   // 2: request.CreateNetworkBrokerInK8s
	(*DeleteNetworkBrokerInK8S)(nil),   // 3: request.DeleteNetworkBrokerInK8s
	(*DeleteNamespacesInK8S)(nil),      // 4: request.DeleteNamespacesInK8s
	(*CreateFlinkClusterInK8SV2)(nil),  // 5: request.CreateFlinkClusterInK8sV2
	(*DeleteFlinkClusterInK8SV2)(nil),  // 6: request.DeleteFlinkClusterInK8sV2
	(*CreateNetworkBrokerInK8SV2)(nil), // 7: request.CreateNetworkBrokerInK8sV2
	(*DeleteNetworkBrokerInK8SV2)(nil), // 8: request.DeleteNetworkBrokerInK8sV2
	(*pbmodel.FlinkCluster)(nil),       // 9: model.FlinkCluster
	(*pbmodel.Network)(nil),            // 10: model.Network
}
var file_proto_types_request_engine_manage_proto_depIdxs = []int32{
	9,  // 0: request.CreateFlinkClusterInK8s.info:type_name -> model.FlinkCluster
	9,  // 1: request.DeleteFlinkClusterInK8s.info:type_name -> model.FlinkCluster
	10, // 2: request.CreateNetworkBrokerInK8s.info:type_name -> model.Network
	10, // 3: request.DeleteNetworkBrokerInK8s.info:type_name -> model.Network
	9,  // 4: request.CreateFlinkClusterInK8sV2.info:type_name -> model.FlinkCluster
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_types_request_engine_manage_proto_init() }
func file_proto_types_request_engine_manage_proto_init() {
	if File_proto_types_request_engine_manage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_types_request_engine_manage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFlinkClusterInK8S); i {
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
		file_proto_types_request_engine_manage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFlinkClusterInK8S); i {
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
		file_proto_types_request_engine_manage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNetworkBrokerInK8S); i {
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
		file_proto_types_request_engine_manage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNetworkBrokerInK8S); i {
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
		file_proto_types_request_engine_manage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNamespacesInK8S); i {
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
		file_proto_types_request_engine_manage_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFlinkClusterInK8SV2); i {
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
		file_proto_types_request_engine_manage_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFlinkClusterInK8SV2); i {
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
		file_proto_types_request_engine_manage_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNetworkBrokerInK8SV2); i {
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
		file_proto_types_request_engine_manage_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNetworkBrokerInK8SV2); i {
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
			RawDescriptor: file_proto_types_request_engine_manage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_request_engine_manage_proto_goTypes,
		DependencyIndexes: file_proto_types_request_engine_manage_proto_depIdxs,
		MessageInfos:      file_proto_types_request_engine_manage_proto_msgTypes,
	}.Build()
	File_proto_types_request_engine_manage_proto = out.File
	file_proto_types_request_engine_manage_proto_rawDesc = nil
	file_proto_types_request_engine_manage_proto_goTypes = nil
	file_proto_types_request_engine_manage_proto_depIdxs = nil
}
