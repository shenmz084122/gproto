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

type UploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	SpaceID     string `protobuf:"bytes,2,opt,name=SpaceID,proto3" json:"SpaceID,omitempty"`
	Data        []byte `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
	FileName    string `protobuf:"bytes,4,opt,name=FileName,proto3" json:"FileName,omitempty"`
	FileDir     string `protobuf:"bytes,5,opt,name=FileDir,proto3" json:"FileDir,omitempty"`
	FileType    int32  `protobuf:"varint,6,opt,name=FileType,proto3" json:"FileType,omitempty"`
	HdfsAddress string `protobuf:"bytes,7,opt,name=HdfsAddress,proto3" json:"HdfsAddress,omitempty"`
	UserID      string `protobuf:"bytes,8,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filemanager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UploadRequest.ProtoReflect.Descriptor instead.
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return file_proto_filemanager_proto_rawDescGZIP(), []int{0}
}

func (x *UploadRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UploadRequest) GetSpaceID() string {
	if x != nil {
		return x.SpaceID
	}
	return ""
}

func (x *UploadRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UploadRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UploadRequest) GetFileDir() string {
	if x != nil {
		return x.FileDir
	}
	return ""
}

func (x *UploadRequest) GetFileType() int32 {
	if x != nil {
		return x.FileType
	}
	return 0
}

func (x *UploadRequest) GetHdfsAddress() string {
	if x != nil {
		return x.HdfsAddress
	}
	return ""
}

func (x *UploadRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type DirRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parent   string  `protobuf:"bytes,1,opt,name=Parent,proto3" json:"Parent,omitempty"`
	SpaceID  string  `protobuf:"bytes,2,opt,name=SpaceID,proto3" json:"SpaceID,omitempty"`
	Address  string  `protobuf:"bytes,3,opt,name=Address,proto3" json:"Address,omitempty"`
	FileType []int32 `protobuf:"varint,4,rep,packed,name=FileType,proto3" json:"FileType,omitempty"`
}

func (x *DirRequest) Reset() {
	*x = DirRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filemanager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirRequest) ProtoMessage() {}

func (x *DirRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DirRequest.ProtoReflect.Descriptor instead.
func (*DirRequest) Descriptor() ([]byte, []int) {
	return file_proto_filemanager_proto_rawDescGZIP(), []int{1}
}

func (x *DirRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *DirRequest) GetSpaceID() string {
	if x != nil {
		return x.SpaceID
	}
	return ""
}

func (x *DirRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *DirRequest) GetFileType() []int32 {
	if x != nil {
		return x.FileType
	}
	return nil
}

type DownloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DownloadRequest) Reset() {
	*x = DownloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filemanager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadRequest) ProtoMessage() {}

func (x *DownloadRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DownloadRequest.ProtoReflect.Descriptor instead.
func (*DownloadRequest) Descriptor() ([]byte, []int) {
	return file_proto_filemanager_proto_rawDescGZIP(), []int{2}
}

func (x *DownloadRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filemanager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_filemanager_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type IdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *IdRequest) Reset() {
	*x = IdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filemanager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdRequest) ProtoMessage() {}

func (x *IdRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use IdRequest.ProtoReflect.Descriptor instead.
func (*IdRequest) Descriptor() ([]byte, []int) {
	return file_proto_filemanager_proto_rawDescGZIP(), []int{4}
}

func (x *IdRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type ListFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address  string `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	SpaceID  string `protobuf:"bytes,2,opt,name=SpaceID,proto3" json:"SpaceID,omitempty"`
	URI      string `protobuf:"bytes,3,opt,name=URI,proto3" json:"URI,omitempty"`
	FileType int32  `protobuf:"varint,4,opt,name=FileType,proto3" json:"FileType,omitempty"`
	FileName string `protobuf:"bytes,5,opt,name=FileName,proto3" json:"FileName,omitempty"`
}

func (x *ListFileRequest) Reset() {
	*x = ListFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filemanager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFileRequest) ProtoMessage() {}

func (x *ListFileRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListFileRequest.ProtoReflect.Descriptor instead.
func (*ListFileRequest) Descriptor() ([]byte, []int) {
	return file_proto_filemanager_proto_rawDescGZIP(), []int{5}
}

func (x *ListFileRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ListFileRequest) GetSpaceID() string {
	if x != nil {
		return x.SpaceID
	}
	return ""
}

func (x *ListFileRequest) GetURI() string {
	if x != nil {
		return x.URI
	}
	return ""
}

func (x *ListFileRequest) GetFileType() int32 {
	if x != nil {
		return x.FileType
	}
	return 0
}

func (x *ListFileRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type ListDirReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileInfo []*ListDirReply_ListInfo `protobuf:"bytes,1,rep,name=FileInfo,proto3" json:"FileInfo,omitempty"`
}

func (x *ListDirReply) Reset() {
	*x = ListDirReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filemanager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDirReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDirReply) ProtoMessage() {}

func (x *ListDirReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListDirReply.ProtoReflect.Descriptor instead.
func (*ListDirReply) Descriptor() ([]byte, []int) {
	return file_proto_filemanager_proto_rawDescGZIP(), []int{6}
}

func (x *ListDirReply) GetFileInfo() []*ListDirReply_ListInfo {
	if x != nil {
		return x.FileInfo
	}
	return nil
}

type DownloadReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *DownloadReply) Reset() {
	*x = DownloadReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filemanager_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadReply) ProtoMessage() {}

func (x *DownloadReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DownloadReply.ProtoReflect.Descriptor instead.
func (*DownloadReply) Descriptor() ([]byte, []int) {
	return file_proto_filemanager_proto_rawDescGZIP(), []int{7}
}

func (x *DownloadReply) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type FileInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	FileType int32  `protobuf:"varint,2,opt,name=FileType,proto3" json:"FileType,omitempty"`
	URI      string `protobuf:"bytes,4,opt,name=URI,proto3" json:"URI,omitempty"`
}

func (x *FileInfoReply) Reset() {
	*x = FileInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filemanager_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfoReply) ProtoMessage() {}

func (x *FileInfoReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use FileInfoReply.ProtoReflect.Descriptor instead.
func (*FileInfoReply) Descriptor() ([]byte, []int) {
	return file_proto_filemanager_proto_rawDescGZIP(), []int{8}
}

func (x *FileInfoReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileInfoReply) GetFileType() int32 {
	if x != nil {
		return x.FileType
	}
	return 0
}

func (x *FileInfoReply) GetURI() string {
	if x != nil {
		return x.URI
	}
	return ""
}

type FileListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total        uint32           `protobuf:"varint,1,opt,name=Total,proto3" json:"Total,omitempty"`
	FileInfoList []*FileInfoReply `protobuf:"bytes,2,rep,name=FileInfoList,proto3" json:"FileInfoList,omitempty"`
}

func (x *FileListReply) Reset() {
	*x = FileListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filemanager_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileListReply) ProtoMessage() {}

func (x *FileListReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use FileListReply.ProtoReflect.Descriptor instead.
func (*FileListReply) Descriptor() ([]byte, []int) {
	return file_proto_filemanager_proto_rawDescGZIP(), []int{9}
}

func (x *FileListReply) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *FileListReply) GetFileInfoList() []*FileInfoReply {
	if x != nil {
		return x.FileInfoList
	}
	return nil
}

type ListDirReply_ListInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Path  string `protobuf:"bytes,3,opt,name=Path,proto3" json:"Path,omitempty"`
	IsDir bool   `protobuf:"varint,4,opt,name=IsDir,proto3" json:"IsDir,omitempty"`
}

func (x *ListDirReply_ListInfo) Reset() {
	*x = ListDirReply_ListInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_filemanager_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDirReply_ListInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDirReply_ListInfo) ProtoMessage() {}

func (x *ListDirReply_ListInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_filemanager_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDirReply_ListInfo.ProtoReflect.Descriptor instead.
func (*ListDirReply_ListInfo) Descriptor() ([]byte, []int) {
	return file_proto_filemanager_proto_rawDescGZIP(), []int{6, 0}
}

func (x *ListDirReply_ListInfo) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *ListDirReply_ListInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListDirReply_ListInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ListDirReply_ListInfo) GetIsDir() bool {
	if x != nil {
		return x.IsDir
	}
	return false
}

var File_proto_filemanager_proto protoreflect.FileDescriptor

var file_proto_filemanager_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x6d, 0x70, 0x62, 0x1a,
	0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74,
	0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x02, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x78, 0x19, 0x52, 0x02, 0x49, 0x44, 0x12, 0x21,
	0x0a, 0x07, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01, 0x14, 0x52, 0x07, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49,
	0x44, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52,
	0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x07, 0x46, 0x69, 0x6c,
	0x65, 0x44, 0x69, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02,
	0x20, 0x01, 0x52, 0x07, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x69, 0x72, 0x12, 0x22, 0x0a, 0x08, 0x46,
	0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x42, 0x06, 0xe2,
	0xdf, 0x1f, 0x02, 0x18, 0x04, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x28, 0x0a, 0x0b, 0x48, 0x64, 0x66, 0x73, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x0b, 0x48, 0x64,
	0x66, 0x73, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20,
	0x00, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x8d, 0x01, 0x0a, 0x0a, 0x44, 0x69,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x78, 0x19,
	0x52, 0x06, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x07, 0x53, 0x70, 0x61, 0x63,
	0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x80,
	0x01, 0x14, 0x52, 0x07, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf,
	0x1f, 0x02, 0x20, 0x01, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x29, 0x0a, 0x0f, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x78, 0x19,
	0x52, 0x02, 0x49, 0x44, 0x22, 0x27, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x78, 0x19, 0x52, 0x02, 0x49, 0x44, 0x22, 0x23, 0x0a,
	0x09, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x78, 0x19, 0x52, 0x02,
	0x49, 0x44, 0x22, 0x8f, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52,
	0x49, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x49, 0x12, 0x1a, 0x0a, 0x08,
	0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0xa1, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x37, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66, 0x6d, 0x70, 0x62, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x44, 0x69, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x58,
	0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x73, 0x44, 0x69, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x49, 0x73, 0x44, 0x69, 0x72, 0x22, 0x23, 0x0a, 0x0d, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x51, 0x0a,
	0x0d, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x55, 0x52, 0x49, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x49,
	0x22, 0x5e, 0x0a, 0x0d, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x37, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x66, 0x6d, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x52, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x32, 0xe4, 0x03, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x12, 0x33, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x13, 0x2e, 0x66, 0x6d, 0x70,
	0x62, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x13, 0x2e, 0x66, 0x6d, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00,
	0x28, 0x01, 0x12, 0x40, 0x0a, 0x0e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x15, 0x2e, 0x66, 0x6d, 0x70, 0x62, 0x2e, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x66, 0x6d,
	0x70, 0x62, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x34, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x44, 0x69, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x10, 0x2e, 0x66, 0x6d, 0x70, 0x62, 0x2e, 0x44, 0x69, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x66, 0x6d, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x44, 0x69, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x13, 0x2e, 0x66,
	0x6d, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x44, 0x69, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x13, 0x2e, 0x66, 0x6d, 0x70, 0x62, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x42, 0x79,
	0x49, 0x64, 0x12, 0x0f, 0x2e, 0x66, 0x6d, 0x70, 0x62, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x66, 0x6d, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x66, 0x6d, 0x70, 0x62,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x66, 0x6d, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x57, 0x6f, 0x72, 0x6b, 0x62, 0x65,
	0x6e, 0x63, 0x68, 0x2f, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66,
	0x6d, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_proto_filemanager_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_filemanager_proto_goTypes = []interface{}{
	(*UploadRequest)(nil),         // 0: fmpb.UploadRequest
	(*DirRequest)(nil),            // 1: fmpb.DirRequest
	(*DownloadRequest)(nil),       // 2: fmpb.DownloadRequest
	(*DeleteRequest)(nil),         // 3: fmpb.DeleteRequest
	(*IdRequest)(nil),             // 4: fmpb.IdRequest
	(*ListFileRequest)(nil),       // 5: fmpb.ListFileRequest
	(*ListDirReply)(nil),          // 6: fmpb.ListDirReply
	(*DownloadReply)(nil),         // 7: fmpb.DownloadReply
	(*FileInfoReply)(nil),         // 8: fmpb.FileInfoReply
	(*FileListReply)(nil),         // 9: fmpb.FileListReply
	(*ListDirReply_ListInfo)(nil), // 10: fmpb.ListDirReply.ListInfo
	(*model.EmptyStruct)(nil),     // 11: model.EmptyStruct
}
var file_proto_filemanager_proto_depIdxs = []int32{
	10, // 0: fmpb.ListDirReply.FileInfo:type_name -> fmpb.ListDirReply.ListInfo
	8,  // 1: fmpb.FileListReply.FileInfoList:type_name -> fmpb.FileInfoReply
	0,  // 2: fmpb.FileManager.Upload:input_type -> fmpb.UploadRequest
	0,  // 3: fmpb.FileManager.UploadStream:input_type -> fmpb.UploadRequest
	2,  // 4: fmpb.FileManager.DownloadStream:input_type -> fmpb.DownloadRequest
	1,  // 5: fmpb.FileManager.GetDirList:input_type -> fmpb.DirRequest
	3,  // 6: fmpb.FileManager.DeleteFileById:input_type -> fmpb.DeleteRequest
	3,  // 7: fmpb.FileManager.DeleteDirById:input_type -> fmpb.DeleteRequest
	4,  // 8: fmpb.FileManager.GetFileById:input_type -> fmpb.IdRequest
	5,  // 9: fmpb.FileManager.GetFileList:input_type -> fmpb.ListFileRequest
	11, // 10: fmpb.FileManager.Upload:output_type -> model.EmptyStruct
	11, // 11: fmpb.FileManager.UploadStream:output_type -> model.EmptyStruct
	7,  // 12: fmpb.FileManager.DownloadStream:output_type -> fmpb.DownloadReply
	6,  // 13: fmpb.FileManager.GetDirList:output_type -> fmpb.ListDirReply
	11, // 14: fmpb.FileManager.DeleteFileById:output_type -> model.EmptyStruct
	11, // 15: fmpb.FileManager.DeleteDirById:output_type -> model.EmptyStruct
	8,  // 16: fmpb.FileManager.GetFileById:output_type -> fmpb.FileInfoReply
	9,  // 17: fmpb.FileManager.GetFileList:output_type -> fmpb.FileListReply
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_filemanager_proto_init() }
func file_proto_filemanager_proto_init() {
	if File_proto_filemanager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_filemanager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadRequest); i {
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
			switch v := v.(*DirRequest); i {
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
		file_proto_filemanager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_filemanager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdRequest); i {
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
			switch v := v.(*ListFileRequest); i {
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
			switch v := v.(*ListDirReply); i {
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
			switch v := v.(*DownloadReply); i {
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
			switch v := v.(*FileInfoReply); i {
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
			switch v := v.(*FileListReply); i {
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
		file_proto_filemanager_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDirReply_ListInfo); i {
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
			NumMessages:   11,
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
