// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/filemanager.proto

package fmpb

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

type UploadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The file id for test. Not required.
	FileId string `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	// The workspace id in HTTP Request-URI
	SpaceId string `protobuf:"bytes,2,opt,name=space_id,json=spaceId,proto3" json:"space_id,omitempty" params:"space_id" uri:"space_id" binding:"len=20" swaggerignore:"true"`
	// The file name. required.
	FileName string `protobuf:"bytes,3,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty" params:"file_name" binding:"required,lte=500"`
	// The file type. required.
	FileType int32 `protobuf:"varint,4,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty" params:"file_type" binding:"required,oneof=1 2"`
	// The data for file. required.
	Data []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UploadFileRequest) Reset() {
	*x = UploadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filemanager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileRequest) ProtoMessage() {}

func (x *UploadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_filemanager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileRequest.ProtoReflect.Descriptor instead.
func (*UploadFileRequest) Descriptor() ([]byte, []int) {
	return file_proto_filemanager_proto_rawDescGZIP(), []int{0}
}

func (x *UploadFileRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *UploadFileRequest) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *UploadFileRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UploadFileRequest) GetFileType() int32 {
	if x != nil {
		return x.FileType
	}
	return 0
}

func (x *UploadFileRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type DownloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The file id in HTTP Request_URI.
	FileId string `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty" params:"file_id" uri:"file_id" binding:"len=20" swaggerignore:"true"`
}

func (x *DownloadRequest) Reset() {
	*x = DownloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filemanager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadRequest) ProtoMessage() {}

func (x *DownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_filemanager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadRequest.ProtoReflect.Descriptor instead.
func (*DownloadRequest) Descriptor() ([]byte, []int) {
	return file_proto_filemanager_proto_rawDescGZIP(), []int{1}
}

func (x *DownloadRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The workspace id in HTTP Request-URI
	SpaceId string `protobuf:"bytes,1,opt,name=space_id,json=spaceId,proto3" json:"space_id,omitempty" params:"space_id" uri:"space_id" binding:"len=20" swaggerignore:"true"`
	// Limit the maximum number of entries returned this time.
	// Not required, Max 100, default 100.
	Limit int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty" params:"limit" form:"limit" default:"100" binding:"gt=0,lte=100"`
	// The offset position. Not required, default 0.
	Offset int32 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty" params:"offset" form:"offset" default:"100" binding:"gt=0,lte=100"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filemanager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_filemanager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_proto_filemanager_proto_rawDescGZIP(), []int{2}
}

func (x *ListRequest) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *ListRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type UpdateFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The file id in HTTP Request_URI.
	FileId string `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty" params:"file_id" uri:"file_id" binding:"len=20" swaggerignore:"true"`
	// The file name. required.
	FileName string `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty" params:"file_name" binding:"required,lte=500"`
	// The file type. required.
	FileType int32 `protobuf:"varint,3,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty" params:"file_type" binding:"required,oneof=1 2"`
}

func (x *UpdateFileRequest) Reset() {
	*x = UpdateFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filemanager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFileRequest) ProtoMessage() {}

func (x *UpdateFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_filemanager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFileRequest.ProtoReflect.Descriptor instead.
func (*UpdateFileRequest) Descriptor() ([]byte, []int) {
	return file_proto_filemanager_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateFileRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *UpdateFileRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UpdateFileRequest) GetFileType() int32 {
	if x != nil {
		return x.FileType
	}
	return 0
}

type DeleteFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of file id. Is required, Min 1 Max 100.
	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty" params:"ids" binding:"gte=1,lte=100"`
}

func (x *DeleteFilesRequest) Reset() {
	*x = DeleteFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filemanager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFilesRequest) ProtoMessage() {}

func (x *DeleteFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_filemanager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFilesRequest.ProtoReflect.Descriptor instead.
func (*DeleteFilesRequest) Descriptor() ([]byte, []int) {
	return file_proto_filemanager_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteFilesRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type DeleteAllFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of space id. Is required, Min 1 Max 100
	SpaceIds []string `protobuf:"bytes,1,rep,name=space_ids,json=spaceIds,proto3" json:"space_ids,omitempty" param:"space_ids" binding:"gte=1,lte=100"`
}

func (x *DeleteAllFilesRequest) Reset() {
	*x = DeleteAllFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filemanager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAllFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAllFilesRequest) ProtoMessage() {}

func (x *DeleteAllFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_filemanager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAllFilesRequest.ProtoReflect.Descriptor instead.
func (*DeleteAllFilesRequest) Descriptor() ([]byte, []int) {
	return file_proto_filemanager_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteAllFilesRequest) GetSpaceIds() []string {
	if x != nil {
		return x.SpaceIds
	}
	return nil
}

type DescribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The file id in HTTP Request_URI.
	FileId string `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty" params:"file_id" uri:"file_id" binding:"len=20" swaggerignore:"true"`
}

func (x *DescribeRequest) Reset() {
	*x = DescribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filemanager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRequest) ProtoMessage() {}

func (x *DescribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_filemanager_proto_msgTypes[6]
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
	return file_proto_filemanager_proto_rawDescGZIP(), []int{6}
}

func (x *DescribeRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type DownloadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DownloadResponse) Reset() {
	*x = DownloadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filemanager_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadResponse) ProtoMessage() {}

func (x *DownloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_filemanager_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadResponse.ProtoReflect.Descriptor instead.
func (*DownloadResponse) Descriptor() ([]byte, []int) {
	return file_proto_filemanager_proto_rawDescGZIP(), []int{7}
}

func (x *DownloadResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type FileInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SpaceId  string `protobuf:"bytes,2,opt,name=space_id,json=spaceId,proto3" json:"space_id,omitempty"`
	FileName string `protobuf:"bytes,3,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FilePath string `protobuf:"bytes,4,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	FileType int32  `protobuf:"varint,5,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty"`
	Url      string `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *FileInfoResponse) Reset() {
	*x = FileInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filemanager_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfoResponse) ProtoMessage() {}

func (x *FileInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_filemanager_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfoResponse.ProtoReflect.Descriptor instead.
func (*FileInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_filemanager_proto_rawDescGZIP(), []int{8}
}

func (x *FileInfoResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FileInfoResponse) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *FileInfoResponse) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *FileInfoResponse) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *FileInfoResponse) GetFileType() int32 {
	if x != nil {
		return x.FileType
	}
	return 0
}

func (x *FileInfoResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos   []*FileInfoResponse `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos,omitempty"`
	HasMore bool                `protobuf:"varint,2,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
	Total   int64               `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filemanager_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_filemanager_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_proto_filemanager_proto_rawDescGZIP(), []int{9}
}

func (x *ListResponse) GetInfos() []*FileInfoResponse {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *ListResponse) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

func (x *ListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_proto_filemanager_proto protoreflect.FileDescriptor

var file_proto_filemanager_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x6d, 0x70, 0x62, 0x1a,
	0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74,
	0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c,
	0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x08, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01, 0x14, 0x52, 0x07,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02,
	0x20, 0x01, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x09,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x18, 0x03, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x42,
	0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x33, 0x0a,
	0x0f, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01, 0x14, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x64, 0x22, 0x7a, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x22, 0x0a, 0x08, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01, 0x14, 0x52, 0x07, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0xe2, 0xdf, 0x1f, 0x04, 0x10, 0x00, 0x18, 0x65, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0f, 0xe2, 0xdf, 0x1f, 0x0b, 0x10, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x7f,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01, 0x14, 0x52, 0x06, 0x66,
	0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x09, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x06, 0xe2,
	0xdf, 0x1f, 0x02, 0x18, 0x03, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x30, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x42, 0x08, 0xe2, 0xdf, 0x1f, 0x04, 0x70, 0x00, 0x78, 0x65, 0x52, 0x03, 0x69, 0x64,
	0x73, 0x22, 0x3e, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x09, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x08, 0xe2,
	0xdf, 0x1f, 0x04, 0x70, 0x00, 0x78, 0x65, 0x52, 0x08, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64,
	0x73, 0x22, 0x33, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01, 0x14, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x26, 0x0a, 0x10, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa6,
	0x01, 0x0a, 0x10, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x79, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x6e, 0x66, 0x6f, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x66, 0x6d, 0x70, 0x62, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05,
	0x69, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x1f, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x6f, 0x72,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x42, 0x04, 0xe2, 0xdf, 0x1f, 0x00, 0x52, 0x07, 0x68,
	0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x04, 0xe2, 0xdf, 0x1f, 0x00, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x32, 0xc7, 0x03, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x12, 0x3d, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x17, 0x2e, 0x66, 0x6d, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x28,
	0x01, 0x12, 0x41, 0x0a, 0x0c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x15, 0x2e, 0x66, 0x6d, 0x70, 0x62, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x66, 0x6d, 0x70, 0x62, 0x2e,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x3f, 0x0a, 0x0c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x15, 0x2e, 0x66, 0x6d, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x66, 0x6d,
	0x70, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c,
	0x65, 0x73, 0x12, 0x11, 0x2e, 0x66, 0x6d, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x66, 0x6d, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x66, 0x6d, 0x70, 0x62,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x66, 0x6d, 0x70, 0x62, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x1b, 0x2e, 0x66, 0x6d, 0x70, 0x62,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x57,
	0x6f, 0x72, 0x6b, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2f, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x66, 0x6d, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_filemanager_proto_rawDescOnce sync.Once
	file_proto_filemanager_proto_rawDescData = file_proto_filemanager_proto_rawDesc
)

func file_proto_filemanager_proto_rawDescGZIP() []byte {
	file_proto_filemanager_proto_rawDescOnce.Do(func() {
		file_proto_filemanager_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_filemanager_proto_rawDescData)
	})
	return file_proto_filemanager_proto_rawDescData
}

var file_proto_filemanager_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_filemanager_proto_goTypes = []interface{}{
	(*UploadFileRequest)(nil),     // 0: fmpb.UploadFileRequest
	(*DownloadRequest)(nil),       // 1: fmpb.DownloadRequest
	(*ListRequest)(nil),           // 2: fmpb.ListRequest
	(*UpdateFileRequest)(nil),     // 3: fmpb.UpdateFileRequest
	(*DeleteFilesRequest)(nil),    // 4: fmpb.DeleteFilesRequest
	(*DeleteAllFilesRequest)(nil), // 5: fmpb.DeleteAllFilesRequest
	(*DescribeRequest)(nil),       // 6: fmpb.DescribeRequest
	(*DownloadResponse)(nil),      // 7: fmpb.DownloadResponse
	(*FileInfoResponse)(nil),      // 8: fmpb.FileInfoResponse
	(*ListResponse)(nil),          // 9: fmpb.ListResponse
	(*model.EmptyStruct)(nil),     // 10: model.EmptyStruct
}
var file_proto_filemanager_proto_depIdxs = []int32{
	8,  // 0: fmpb.ListResponse.infos:type_name -> fmpb.FileInfoResponse
	0,  // 1: fmpb.FileManager.UploadFile:input_type -> fmpb.UploadFileRequest
	1,  // 2: fmpb.FileManager.DownloadFile:input_type -> fmpb.DownloadRequest
	6,  // 3: fmpb.FileManager.DescribeFile:input_type -> fmpb.DescribeRequest
	2,  // 4: fmpb.FileManager.ListFiles:input_type -> fmpb.ListRequest
	3,  // 5: fmpb.FileManager.UpdateFile:input_type -> fmpb.UpdateFileRequest
	4,  // 6: fmpb.FileManager.DeleteFiles:input_type -> fmpb.DeleteFilesRequest
	5,  // 7: fmpb.FileManager.DeleteAllFiles:input_type -> fmpb.DeleteAllFilesRequest
	10, // 8: fmpb.FileManager.UploadFile:output_type -> model.EmptyStruct
	7,  // 9: fmpb.FileManager.DownloadFile:output_type -> fmpb.DownloadResponse
	8,  // 10: fmpb.FileManager.DescribeFile:output_type -> fmpb.FileInfoResponse
	9,  // 11: fmpb.FileManager.ListFiles:output_type -> fmpb.ListResponse
	10, // 12: fmpb.FileManager.UpdateFile:output_type -> model.EmptyStruct
	10, // 13: fmpb.FileManager.DeleteFiles:output_type -> model.EmptyStruct
	10, // 14: fmpb.FileManager.DeleteAllFiles:output_type -> model.EmptyStruct
	8,  // [8:15] is the sub-list for method output_type
	1,  // [1:8] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_proto_filemanager_proto_init() }
func file_proto_filemanager_proto_init() {
	if File_proto_filemanager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_filemanager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileRequest); i {
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
		file_proto_filemanager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadRequest); i {
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
		file_proto_filemanager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_proto_filemanager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFileRequest); i {
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
		file_proto_filemanager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFilesRequest); i {
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
		file_proto_filemanager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAllFilesRequest); i {
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
		file_proto_filemanager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_filemanager_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadResponse); i {
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
		file_proto_filemanager_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfoResponse); i {
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
		file_proto_filemanager_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
			RawDescriptor: file_proto_filemanager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_filemanager_proto_goTypes,
		DependencyIndexes: file_proto_filemanager_proto_depIdxs,
		MessageInfos:      file_proto_filemanager_proto_msgTypes,
	}.Build()
	File_proto_filemanager_proto = out.File
	file_proto_filemanager_proto_rawDesc = nil
	file_proto_filemanager_proto_goTypes = nil
	file_proto_filemanager_proto_depIdxs = nil
}
