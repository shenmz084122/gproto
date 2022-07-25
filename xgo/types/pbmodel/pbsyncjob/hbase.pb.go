// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: proto/types/model/syncjob/hbase.proto

package pbsyncjob

import (
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbgosql"
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

type HbaseSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// column
	Column []*Column `protobuf:"bytes,1,rep,name=column,proto3" json:"column"`
	// hbase table
	Table string `protobuf:"bytes,2,opt,name=table,proto3" json:"table"`
	// change_log
	ChangeLog string `protobuf:"bytes,3,opt,name=change_log,json=changeLog,proto3" json:"change_log"`
	// encoding
	Encoding string `protobuf:"bytes,4,opt,name=encoding,proto3" json:"encoding"`
	// scan cache size
	ScanCacheSize int32 `protobuf:"varint,5,opt,name=scan_cache_size,json=scanCacheSize,proto3" json:"scan_cache_size"`
	// hbase scan batch size
	ScanBatchSize int32  `protobuf:"varint,6,opt,name=scan_batch_size,json=scanBatchSize,proto3" json:"scan_batch_size"`
	StartRowKey   string `protobuf:"bytes,7,opt,name=start_row_key,json=startRowKey,proto3" json:"start_row_key"`
	// end row key
	EndRowKey string `protobuf:"bytes,8,opt,name=end_row_key,json=endRowKey,proto3" json:"end_row_key"`
	// is binary rowkey
	IsBinaryRowkey bool `protobuf:"varint,9,opt,name=is_binary_rowkey,json=isBinaryRowkey,proto3" json:"is_binary_rowkey"`
}

func (x *HbaseSource) Reset() {
	*x = HbaseSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_model_syncjob_hbase_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HbaseSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HbaseSource) ProtoMessage() {}

func (x *HbaseSource) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_model_syncjob_hbase_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HbaseSource.ProtoReflect.Descriptor instead.
func (*HbaseSource) Descriptor() ([]byte, []int) {
	return file_proto_types_model_syncjob_hbase_proto_rawDescGZIP(), []int{0}
}

func (x *HbaseSource) GetColumn() []*Column {
	if x != nil {
		return x.Column
	}
	return nil
}

func (x *HbaseSource) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *HbaseSource) GetChangeLog() string {
	if x != nil {
		return x.ChangeLog
	}
	return ""
}

func (x *HbaseSource) GetEncoding() string {
	if x != nil {
		return x.Encoding
	}
	return ""
}

func (x *HbaseSource) GetScanCacheSize() int32 {
	if x != nil {
		return x.ScanCacheSize
	}
	return 0
}

func (x *HbaseSource) GetScanBatchSize() int32 {
	if x != nil {
		return x.ScanBatchSize
	}
	return 0
}

func (x *HbaseSource) GetStartRowKey() string {
	if x != nil {
		return x.StartRowKey
	}
	return ""
}

func (x *HbaseSource) GetEndRowKey() string {
	if x != nil {
		return x.EndRowKey
	}
	return ""
}

func (x *HbaseSource) GetIsBinaryRowkey() bool {
	if x != nil {
		return x.IsBinaryRowkey
	}
	return false
}

type HbaseTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Column []*Column `protobuf:"bytes,1,rep,name=column,proto3" json:"column"`
	// hbase table
	Table string `protobuf:"bytes,2,opt,name=table,proto3" json:"table"`
	// null_mode
	NullMode BaseEnum_NullMode `protobuf:"varint,3,opt,name=null_mode,json=nullMode,proto3,enum=model.BaseEnum_NullMode" json:"null_mode"`
	// wal flag
	WalFlag bool `protobuf:"varint,4,opt,name=wal_flag,json=walFlag,proto3" json:"wal_flag"`
	// write buffer size
	WriteBufferSize int32 `protobuf:"varint,5,opt,name=write_buffer_size,json=writeBufferSize,proto3" json:"write_buffer_size"`
	// change_log
	ChangeLog string `protobuf:"bytes,6,opt,name=change_log,json=changeLog,proto3" json:"change_log"`
	// rowkey express
	RowkeyExpress string `protobuf:"bytes,7,opt,name=rowkey_express,json=rowkeyExpress,proto3" json:"rowkey_express"`
	// version column index
	VersionColumnIndex int32 `protobuf:"varint,8,opt,name=version_column_index,json=versionColumnIndex,proto3" json:"version_column_index"`
	// version column value
	VersionColumnValue string `protobuf:"bytes,9,opt,name=version_column_value,json=versionColumnValue,proto3" json:"version_column_value"`
	// scan cache size
	ScanCacheSize int32 `protobuf:"varint,10,opt,name=scan_cache_size,json=scanCacheSize,proto3" json:"scan_cache_size"`
	// hbase scan batch size
	ScanBatchSize int32 `protobuf:"varint,11,opt,name=scan_batch_size,json=scanBatchSize,proto3" json:"scan_batch_size"`
}

func (x *HbaseTarget) Reset() {
	*x = HbaseTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_model_syncjob_hbase_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HbaseTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HbaseTarget) ProtoMessage() {}

func (x *HbaseTarget) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_model_syncjob_hbase_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HbaseTarget.ProtoReflect.Descriptor instead.
func (*HbaseTarget) Descriptor() ([]byte, []int) {
	return file_proto_types_model_syncjob_hbase_proto_rawDescGZIP(), []int{1}
}

func (x *HbaseTarget) GetColumn() []*Column {
	if x != nil {
		return x.Column
	}
	return nil
}

func (x *HbaseTarget) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *HbaseTarget) GetNullMode() BaseEnum_NullMode {
	if x != nil {
		return x.NullMode
	}
	return BaseEnum_NullModeUnset
}

func (x *HbaseTarget) GetWalFlag() bool {
	if x != nil {
		return x.WalFlag
	}
	return false
}

func (x *HbaseTarget) GetWriteBufferSize() int32 {
	if x != nil {
		return x.WriteBufferSize
	}
	return 0
}

func (x *HbaseTarget) GetChangeLog() string {
	if x != nil {
		return x.ChangeLog
	}
	return ""
}

func (x *HbaseTarget) GetRowkeyExpress() string {
	if x != nil {
		return x.RowkeyExpress
	}
	return ""
}

func (x *HbaseTarget) GetVersionColumnIndex() int32 {
	if x != nil {
		return x.VersionColumnIndex
	}
	return 0
}

func (x *HbaseTarget) GetVersionColumnValue() string {
	if x != nil {
		return x.VersionColumnValue
	}
	return ""
}

func (x *HbaseTarget) GetScanCacheSize() int32 {
	if x != nil {
		return x.ScanCacheSize
	}
	return 0
}

func (x *HbaseTarget) GetScanBatchSize() int32 {
	if x != nil {
		return x.ScanBatchSize
	}
	return 0
}

var File_proto_types_model_syncjob_hbase_proto protoreflect.FileDescriptor

var file_proto_types_model_syncjob_hbase_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x6a, 0x6f, 0x62, 0x2f, 0x68, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x28,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x6a, 0x6f, 0x62, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x65, 0x6e,
	0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x79, 0x6e, 0x63,
	0x6a, 0x6f, 0x62, 0x2f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x75, 0x33,
	0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x73, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xcb, 0x02, 0x0a, 0x0b, 0x48, 0x62, 0x61, 0x73, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x25, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x1a, 0x0a,
	0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x63, 0x61,
	0x6e, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0d, 0x73, 0x63, 0x61, 0x6e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x63, 0x61, 0x6e,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x72, 0x6f, 0x77, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x52, 0x6f, 0x77, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a,
	0x0b, 0x65, 0x6e, 0x64, 0x5f, 0x72, 0x6f, 0x77, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x52, 0x6f, 0x77, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x0a,
	0x10, 0x69, 0x73, 0x5f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x72, 0x6f, 0x77, 0x6b, 0x65,
	0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x42, 0x69, 0x6e, 0x61, 0x72,
	0x79, 0x52, 0x6f, 0x77, 0x6b, 0x65, 0x79, 0x3a, 0x06, 0xca, 0xb2, 0x04, 0x02, 0x0a, 0x00, 0x22,
	0xca, 0x03, 0x0a, 0x0b, 0x48, 0x62, 0x61, 0x73, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12,
	0x25, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x06,
	0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x35, 0x0a, 0x09,
	0x6e, 0x75, 0x6c, 0x6c, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x45, 0x6e, 0x75, 0x6d,
	0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x08, 0x6e, 0x75, 0x6c, 0x6c, 0x4d,
	0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x77, 0x61, 0x6c, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x77, 0x61, 0x6c, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x2a,
	0x0a, 0x11, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x6f, 0x77,
	0x6b, 0x65, 0x79, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x72, 0x6f, 0x77, 0x6b, 0x65, 0x79, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x30, 0x0a, 0x14, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x30, 0x0a, 0x14, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f,
	0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73,
	0x63, 0x61, 0x6e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x26, 0x0a, 0x0f,
	0x73, 0x63, 0x61, 0x6e, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x63, 0x61, 0x6e, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x53, 0x69, 0x7a, 0x65, 0x3a, 0x06, 0xca, 0xb2, 0x04, 0x02, 0x0a, 0x00, 0x42, 0x76, 0x0a, 0x2c,
	0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6f, 0x6d, 0x6e, 0x69, 0x73, 0x2e, 0x67, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x62, 0x73, 0x79, 0x6e, 0x63, 0x6a, 0x6f, 0x62, 0x42, 0x07, 0x50, 0x42,
	0x48, 0x62, 0x61, 0x73, 0x65, 0x50, 0x00, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x57, 0x6f, 0x72, 0x6b, 0x62, 0x65, 0x6e, 0x63,
	0x68, 0x2f, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x78, 0x67, 0x6f, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x70, 0x62, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x62, 0x73, 0x79, 0x6e,
	0x63, 0x6a, 0x6f, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_types_model_syncjob_hbase_proto_rawDescOnce sync.Once
	file_proto_types_model_syncjob_hbase_proto_rawDescData = file_proto_types_model_syncjob_hbase_proto_rawDesc
)

func file_proto_types_model_syncjob_hbase_proto_rawDescGZIP() []byte {
	file_proto_types_model_syncjob_hbase_proto_rawDescOnce.Do(func() {
		file_proto_types_model_syncjob_hbase_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_model_syncjob_hbase_proto_rawDescData)
	})
	return file_proto_types_model_syncjob_hbase_proto_rawDescData
}

var file_proto_types_model_syncjob_hbase_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_types_model_syncjob_hbase_proto_goTypes = []interface{}{
	(*HbaseSource)(nil),    // 0: model.HbaseSource
	(*HbaseTarget)(nil),    // 1: model.HbaseTarget
	(*Column)(nil),         // 2: model.Column
	(BaseEnum_NullMode)(0), // 3: model.BaseEnum.NullMode
}
var file_proto_types_model_syncjob_hbase_proto_depIdxs = []int32{
	2, // 0: model.HbaseSource.column:type_name -> model.Column
	2, // 1: model.HbaseTarget.column:type_name -> model.Column
	3, // 2: model.HbaseTarget.null_mode:type_name -> model.BaseEnum.NullMode
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_types_model_syncjob_hbase_proto_init() }
func file_proto_types_model_syncjob_hbase_proto_init() {
	if File_proto_types_model_syncjob_hbase_proto != nil {
		return
	}
	file_proto_types_model_syncjob_baseenum_proto_init()
	file_proto_types_model_syncjob_column_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_types_model_syncjob_hbase_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HbaseSource); i {
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
		file_proto_types_model_syncjob_hbase_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HbaseTarget); i {
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
			RawDescriptor: file_proto_types_model_syncjob_hbase_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_model_syncjob_hbase_proto_goTypes,
		DependencyIndexes: file_proto_types_model_syncjob_hbase_proto_depIdxs,
		MessageInfos:      file_proto_types_model_syncjob_hbase_proto_msgTypes,
	}.Build()
	File_proto_types_model_syncjob_hbase_proto = out.File
	file_proto_types_model_syncjob_hbase_proto_rawDesc = nil
	file_proto_types_model_syncjob_hbase_proto_goTypes = nil
	file_proto_types_model_syncjob_hbase_proto_depIdxs = nil
}
