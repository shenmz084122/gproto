// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: proto/types/response/dataservice_manage.proto

package pbresponse

import (
	pbmodel "github.com/DataWorkbench/gproto/xgo/types/pbmodel"
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

type ListDataServiceClusters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos   []*pbmodel.DataServiceCluster `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos"`
	HasMore bool                          `protobuf:"varint,2,opt,name=has_more,json=hasMore,proto3" json:"has_more"`
	Total   int64                         `protobuf:"varint,3,opt,name=total,proto3" json:"total"`
}

func (x *ListDataServiceClusters) Reset() {
	*x = ListDataServiceClusters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_response_dataservice_manage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDataServiceClusters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDataServiceClusters) ProtoMessage() {}

func (x *ListDataServiceClusters) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_response_dataservice_manage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDataServiceClusters.ProtoReflect.Descriptor instead.
func (*ListDataServiceClusters) Descriptor() ([]byte, []int) {
	return file_proto_types_response_dataservice_manage_proto_rawDescGZIP(), []int{0}
}

func (x *ListDataServiceClusters) GetInfos() []*pbmodel.DataServiceCluster {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *ListDataServiceClusters) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

func (x *ListDataServiceClusters) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type DescribeDataServiceCluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *pbmodel.DataServiceCluster `protobuf:"bytes,1,opt,name=info,proto3" json:"info"`
}

func (x *DescribeDataServiceCluster) Reset() {
	*x = DescribeDataServiceCluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_response_dataservice_manage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeDataServiceCluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeDataServiceCluster) ProtoMessage() {}

func (x *DescribeDataServiceCluster) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_response_dataservice_manage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeDataServiceCluster.ProtoReflect.Descriptor instead.
func (*DescribeDataServiceCluster) Descriptor() ([]byte, []int) {
	return file_proto_types_response_dataservice_manage_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeDataServiceCluster) GetInfo() *pbmodel.DataServiceCluster {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateDataServiceCluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *CreateDataServiceCluster) Reset() {
	*x = CreateDataServiceCluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_response_dataservice_manage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDataServiceCluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDataServiceCluster) ProtoMessage() {}

func (x *CreateDataServiceCluster) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_response_dataservice_manage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDataServiceCluster.ProtoReflect.Descriptor instead.
func (*CreateDataServiceCluster) Descriptor() ([]byte, []int) {
	return file_proto_types_response_dataservice_manage_proto_rawDescGZIP(), []int{2}
}

func (x *CreateDataServiceCluster) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListApiGroups struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos   []*pbmodel.ApiGroup `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos"`
	HasMore bool                `protobuf:"varint,2,opt,name=has_more,json=hasMore,proto3" json:"has_more"`
	Total   int64               `protobuf:"varint,3,opt,name=total,proto3" json:"total"`
}

func (x *ListApiGroups) Reset() {
	*x = ListApiGroups{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_response_dataservice_manage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApiGroups) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApiGroups) ProtoMessage() {}

func (x *ListApiGroups) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_response_dataservice_manage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApiGroups.ProtoReflect.Descriptor instead.
func (*ListApiGroups) Descriptor() ([]byte, []int) {
	return file_proto_types_response_dataservice_manage_proto_rawDescGZIP(), []int{3}
}

func (x *ListApiGroups) GetInfos() []*pbmodel.ApiGroup {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *ListApiGroups) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

func (x *ListApiGroups) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type CreateApiGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *CreateApiGroup) Reset() {
	*x = CreateApiGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_response_dataservice_manage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateApiGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApiGroup) ProtoMessage() {}

func (x *CreateApiGroup) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_response_dataservice_manage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApiGroup.ProtoReflect.Descriptor instead.
func (*CreateApiGroup) Descriptor() ([]byte, []int) {
	return file_proto_types_response_dataservice_manage_proto_rawDescGZIP(), []int{4}
}

func (x *CreateApiGroup) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListApiConfigs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos   []*pbmodel.ApiConfig `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos"`
	Total   int64                `protobuf:"varint,2,opt,name=total,proto3" json:"total"`
	HasMore bool                 `protobuf:"varint,3,opt,name=has_more,json=hasMore,proto3" json:"has_more"`
}

func (x *ListApiConfigs) Reset() {
	*x = ListApiConfigs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_response_dataservice_manage_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApiConfigs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApiConfigs) ProtoMessage() {}

func (x *ListApiConfigs) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_response_dataservice_manage_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApiConfigs.ProtoReflect.Descriptor instead.
func (*ListApiConfigs) Descriptor() ([]byte, []int) {
	return file_proto_types_response_dataservice_manage_proto_rawDescGZIP(), []int{5}
}

func (x *ListApiConfigs) GetInfos() []*pbmodel.ApiConfig {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *ListApiConfigs) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListApiConfigs) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

type DescribeApiConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiConfig *pbmodel.ApiConfig `protobuf:"bytes,1,opt,name=api_config,json=apiConfig,proto3" json:"api_config"`
}

func (x *DescribeApiConfig) Reset() {
	*x = DescribeApiConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_response_dataservice_manage_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeApiConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeApiConfig) ProtoMessage() {}

func (x *DescribeApiConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_response_dataservice_manage_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeApiConfig.ProtoReflect.Descriptor instead.
func (*DescribeApiConfig) Descriptor() ([]byte, []int) {
	return file_proto_types_response_dataservice_manage_proto_rawDescGZIP(), []int{6}
}

func (x *DescribeApiConfig) GetApiConfig() *pbmodel.ApiConfig {
	if x != nil {
		return x.ApiConfig
	}
	return nil
}

type CreateApiConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *CreateApiConfig) Reset() {
	*x = CreateApiConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_response_dataservice_manage_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateApiConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApiConfig) ProtoMessage() {}

func (x *CreateApiConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_response_dataservice_manage_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApiConfig.ProtoReflect.Descriptor instead.
func (*CreateApiConfig) Descriptor() ([]byte, []int) {
	return file_proto_types_response_dataservice_manage_proto_rawDescGZIP(), []int{7}
}

func (x *CreateApiConfig) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DescribeDataServiceApiVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiVersion *pbmodel.ApiVersion `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version"`
}

func (x *DescribeDataServiceApiVersion) Reset() {
	*x = DescribeDataServiceApiVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_response_dataservice_manage_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeDataServiceApiVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeDataServiceApiVersion) ProtoMessage() {}

func (x *DescribeDataServiceApiVersion) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_response_dataservice_manage_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeDataServiceApiVersion.ProtoReflect.Descriptor instead.
func (*DescribeDataServiceApiVersion) Descriptor() ([]byte, []int) {
	return file_proto_types_response_dataservice_manage_proto_rawDescGZIP(), []int{8}
}

func (x *DescribeDataServiceApiVersion) GetApiVersion() *pbmodel.ApiVersion {
	if x != nil {
		return x.ApiVersion
	}
	return nil
}

type ListDataServiceApiVersions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos   []*pbmodel.ApiVersion `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos"`
	Total   int64                 `protobuf:"varint,2,opt,name=total,proto3" json:"total"`
	HasMore bool                  `protobuf:"varint,3,opt,name=has_more,json=hasMore,proto3" json:"has_more"`
}

func (x *ListDataServiceApiVersions) Reset() {
	*x = ListDataServiceApiVersions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_response_dataservice_manage_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDataServiceApiVersions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDataServiceApiVersions) ProtoMessage() {}

func (x *ListDataServiceApiVersions) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_response_dataservice_manage_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDataServiceApiVersions.ProtoReflect.Descriptor instead.
func (*ListDataServiceApiVersions) Descriptor() ([]byte, []int) {
	return file_proto_types_response_dataservice_manage_proto_rawDescGZIP(), []int{9}
}

func (x *ListDataServiceApiVersions) GetInfos() []*pbmodel.ApiVersion {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *ListDataServiceApiVersions) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListDataServiceApiVersions) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

type ListPublishedApiVersionsByClusterId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos   []*pbmodel.ApiVersion `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos"`
	Total   int64                 `protobuf:"varint,2,opt,name=total,proto3" json:"total"`
	HasMore bool                  `protobuf:"varint,3,opt,name=has_more,json=hasMore,proto3" json:"has_more"`
}

func (x *ListPublishedApiVersionsByClusterId) Reset() {
	*x = ListPublishedApiVersionsByClusterId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_response_dataservice_manage_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPublishedApiVersionsByClusterId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPublishedApiVersionsByClusterId) ProtoMessage() {}

func (x *ListPublishedApiVersionsByClusterId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_response_dataservice_manage_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPublishedApiVersionsByClusterId.ProtoReflect.Descriptor instead.
func (*ListPublishedApiVersionsByClusterId) Descriptor() ([]byte, []int) {
	return file_proto_types_response_dataservice_manage_proto_rawDescGZIP(), []int{10}
}

func (x *ListPublishedApiVersionsByClusterId) GetInfos() []*pbmodel.ApiVersion {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *ListPublishedApiVersionsByClusterId) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListPublishedApiVersionsByClusterId) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

type TestDataServiceApi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logs            string `protobuf:"bytes,1,opt,name=logs,proto3" json:"logs"`
	ResponseContent string `protobuf:"bytes,2,opt,name=response_content,json=responseContent,proto3" json:"response_content"`
}

func (x *TestDataServiceApi) Reset() {
	*x = TestDataServiceApi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_response_dataservice_manage_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestDataServiceApi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestDataServiceApi) ProtoMessage() {}

func (x *TestDataServiceApi) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_response_dataservice_manage_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestDataServiceApi.ProtoReflect.Descriptor instead.
func (*TestDataServiceApi) Descriptor() ([]byte, []int) {
	return file_proto_types_response_dataservice_manage_proto_rawDescGZIP(), []int{11}
}

func (x *TestDataServiceApi) GetLogs() string {
	if x != nil {
		return x.Logs
	}
	return ""
}

func (x *TestDataServiceApi) GetResponseContent() string {
	if x != nil {
		return x.ResponseContent
	}
	return ""
}

// SourceKind used as reply parameters in RPC or response body in HTTP.
type DescribeServiceDataSourceKinds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kinds []*pbmodel.ServiceDataSourceKind `protobuf:"bytes,1,rep,name=Kinds,proto3" json:"kinds"`
}

func (x *DescribeServiceDataSourceKinds) Reset() {
	*x = DescribeServiceDataSourceKinds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_response_dataservice_manage_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeServiceDataSourceKinds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeServiceDataSourceKinds) ProtoMessage() {}

func (x *DescribeServiceDataSourceKinds) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_response_dataservice_manage_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeServiceDataSourceKinds.ProtoReflect.Descriptor instead.
func (*DescribeServiceDataSourceKinds) Descriptor() ([]byte, []int) {
	return file_proto_types_response_dataservice_manage_proto_rawDescGZIP(), []int{12}
}

func (x *DescribeServiceDataSourceKinds) GetKinds() []*pbmodel.ServiceDataSourceKind {
	if x != nil {
		return x.Kinds
	}
	return nil
}

var File_proto_types_response_dataservice_manage_proto protoreflect.FileDescriptor

var file_proto_types_response_dataservice_manage_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b,
	0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x6e, 0x66,
	0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x61,
	0x73, 0x5f, 0x6d, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x61,
	0x73, 0x4d, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x4b, 0x0a, 0x1a, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x2a, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x67, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x69, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x25, 0x0a, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x19, 0x0a, 0x08,
	0x68, 0x61, 0x73, 0x5f, 0x6d, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x68, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x20, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x69, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x73, 0x12, 0x26, 0x0a, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x19, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x22, 0x44, 0x0a, 0x11, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x2f, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70, 0x69, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x09, 0x61, 0x70, 0x69, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x22, 0x21, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x53, 0x0a, 0x1d, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x61, 0x70,
	0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x76, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x70, 0x69, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x70,
	0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x6f, 0x72,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65,
	0x22, 0x7f, 0x0a, 0x23, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41,
	0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x6f,
	0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x72,
	0x65, 0x22, 0x53, 0x0a, 0x12, 0x54, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x41, 0x70, 0x69, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x54, 0x0a, 0x1e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x4b, 0x69, 0x6e, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x05, 0x4b, 0x69, 0x6e, 0x64, 0x73, 0x42, 0x7c, 0x0a, 0x25,
	0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6f, 0x6d, 0x6e, 0x69, 0x73, 0x2e, 0x67, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1b, 0x50, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x50, 0x00, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x44, 0x61, 0x74, 0x61, 0x57, 0x6f, 0x72, 0x6b, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2f, 0x67,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x78, 0x67, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x70, 0x62, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_types_response_dataservice_manage_proto_rawDescOnce sync.Once
	file_proto_types_response_dataservice_manage_proto_rawDescData = file_proto_types_response_dataservice_manage_proto_rawDesc
)

func file_proto_types_response_dataservice_manage_proto_rawDescGZIP() []byte {
	file_proto_types_response_dataservice_manage_proto_rawDescOnce.Do(func() {
		file_proto_types_response_dataservice_manage_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_response_dataservice_manage_proto_rawDescData)
	})
	return file_proto_types_response_dataservice_manage_proto_rawDescData
}

var file_proto_types_response_dataservice_manage_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_proto_types_response_dataservice_manage_proto_goTypes = []interface{}{
	(*ListDataServiceClusters)(nil),             // 0: response.ListDataServiceClusters
	(*DescribeDataServiceCluster)(nil),          // 1: response.DescribeDataServiceCluster
	(*CreateDataServiceCluster)(nil),            // 2: response.CreateDataServiceCluster
	(*ListApiGroups)(nil),                       // 3: response.ListApiGroups
	(*CreateApiGroup)(nil),                      // 4: response.CreateApiGroup
	(*ListApiConfigs)(nil),                      // 5: response.ListApiConfigs
	(*DescribeApiConfig)(nil),                   // 6: response.DescribeApiConfig
	(*CreateApiConfig)(nil),                     // 7: response.CreateApiConfig
	(*DescribeDataServiceApiVersion)(nil),       // 8: response.DescribeDataServiceApiVersion
	(*ListDataServiceApiVersions)(nil),          // 9: response.ListDataServiceApiVersions
	(*ListPublishedApiVersionsByClusterId)(nil), // 10: response.ListPublishedApiVersionsByClusterId
	(*TestDataServiceApi)(nil),                  // 11: response.TestDataServiceApi
	(*DescribeServiceDataSourceKinds)(nil),      // 12: response.DescribeServiceDataSourceKinds
	(*pbmodel.DataServiceCluster)(nil),          // 13: model.DataServiceCluster
	(*pbmodel.ApiGroup)(nil),                    // 14: model.ApiGroup
	(*pbmodel.ApiConfig)(nil),                   // 15: model.ApiConfig
	(*pbmodel.ApiVersion)(nil),                  // 16: model.ApiVersion
	(*pbmodel.ServiceDataSourceKind)(nil),       // 17: model.ServiceDataSourceKind
}
var file_proto_types_response_dataservice_manage_proto_depIdxs = []int32{
	13, // 0: response.ListDataServiceClusters.infos:type_name -> model.DataServiceCluster
	13, // 1: response.DescribeDataServiceCluster.info:type_name -> model.DataServiceCluster
	14, // 2: response.ListApiGroups.infos:type_name -> model.ApiGroup
	15, // 3: response.ListApiConfigs.infos:type_name -> model.ApiConfig
	15, // 4: response.DescribeApiConfig.api_config:type_name -> model.ApiConfig
	16, // 5: response.DescribeDataServiceApiVersion.api_version:type_name -> model.ApiVersion
	16, // 6: response.ListDataServiceApiVersions.infos:type_name -> model.ApiVersion
	16, // 7: response.ListPublishedApiVersionsByClusterId.infos:type_name -> model.ApiVersion
	17, // 8: response.DescribeServiceDataSourceKinds.Kinds:type_name -> model.ServiceDataSourceKind
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_proto_types_response_dataservice_manage_proto_init() }
func file_proto_types_response_dataservice_manage_proto_init() {
	if File_proto_types_response_dataservice_manage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_types_response_dataservice_manage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDataServiceClusters); i {
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
		file_proto_types_response_dataservice_manage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeDataServiceCluster); i {
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
		file_proto_types_response_dataservice_manage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDataServiceCluster); i {
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
		file_proto_types_response_dataservice_manage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListApiGroups); i {
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
		file_proto_types_response_dataservice_manage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateApiGroup); i {
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
		file_proto_types_response_dataservice_manage_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListApiConfigs); i {
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
		file_proto_types_response_dataservice_manage_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeApiConfig); i {
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
		file_proto_types_response_dataservice_manage_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateApiConfig); i {
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
		file_proto_types_response_dataservice_manage_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeDataServiceApiVersion); i {
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
		file_proto_types_response_dataservice_manage_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDataServiceApiVersions); i {
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
		file_proto_types_response_dataservice_manage_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPublishedApiVersionsByClusterId); i {
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
		file_proto_types_response_dataservice_manage_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestDataServiceApi); i {
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
		file_proto_types_response_dataservice_manage_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeServiceDataSourceKinds); i {
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
			RawDescriptor: file_proto_types_response_dataservice_manage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_response_dataservice_manage_proto_goTypes,
		DependencyIndexes: file_proto_types_response_dataservice_manage_proto_depIdxs,
		MessageInfos:      file_proto_types_response_dataservice_manage_proto_msgTypes,
	}.Build()
	File_proto_types_response_dataservice_manage_proto = out.File
	file_proto_types_response_dataservice_manage_proto_rawDesc = nil
	file_proto_types_response_dataservice_manage_proto_goTypes = nil
	file_proto_types_response_dataservice_manage_proto_depIdxs = nil
}
