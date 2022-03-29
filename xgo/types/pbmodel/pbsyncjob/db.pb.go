// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: proto/types/model/syncjob/db.proto

package pbsyncjob

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

type DBTarget_WriteMode int32

const (
	DBTarget_insert  DBTarget_WriteMode = 0
	DBTarget_replace DBTarget_WriteMode = 1
	DBTarget_update  DBTarget_WriteMode = 2
)

// Enum value maps for DBTarget_WriteMode.
var (
	DBTarget_WriteMode_name = map[int32]string{
		0: "insert",
		1: "replace",
		2: "update",
	}
	DBTarget_WriteMode_value = map[string]int32{
		"insert":  0,
		"replace": 1,
		"update":  2,
	}
)

func (x DBTarget_WriteMode) Enum() *DBTarget_WriteMode {
	p := new(DBTarget_WriteMode)
	*p = x
	return p
}

func (x DBTarget_WriteMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DBTarget_WriteMode) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_types_model_syncjob_db_proto_enumTypes[0].Descriptor()
}

func (DBTarget_WriteMode) Type() protoreflect.EnumType {
	return &file_proto_types_model_syncjob_db_proto_enumTypes[0]
}

func (x DBTarget_WriteMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DBTarget_WriteMode.Descriptor instead.
func (DBTarget_WriteMode) EnumDescriptor() ([]byte, []int) {
	return file_proto_types_model_syncjob_db_proto_rawDescGZIP(), []int{1, 0}
}

type DBTarget_Semantic int32

const (
	DBTarget_AtLeastOnce DBTarget_Semantic = 0
	DBTarget_ExactlyOnce DBTarget_Semantic = 1
)

// Enum value maps for DBTarget_Semantic.
var (
	DBTarget_Semantic_name = map[int32]string{
		0: "AtLeastOnce",
		1: "ExactlyOnce",
	}
	DBTarget_Semantic_value = map[string]int32{
		"AtLeastOnce": 0,
		"ExactlyOnce": 1,
	}
)

func (x DBTarget_Semantic) Enum() *DBTarget_Semantic {
	p := new(DBTarget_Semantic)
	*p = x
	return p
}

func (x DBTarget_Semantic) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DBTarget_Semantic) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_types_model_syncjob_db_proto_enumTypes[1].Descriptor()
}

func (DBTarget_Semantic) Type() protoreflect.EnumType {
	return &file_proto_types_model_syncjob_db_proto_enumTypes[1]
}

func (x DBTarget_Semantic) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DBTarget_Semantic.Descriptor instead.
func (DBTarget_Semantic) EnumDescriptor() ([]byte, []int) {
	return file_proto_types_model_syncjob_db_proto_rawDescGZIP(), []int{1, 1}
}

type DBSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// column
	Column []*Column `protobuf:"bytes,1,rep,name=column,proto3" json:"column"`
	// table
	Table []string `protobuf:"bytes,2,rep,name=table,proto3" json:"table"`
	// schema
	Schema string `protobuf:"bytes,3,opt,name=schema,proto3" json:"schema"`
	// where
	Where string `protobuf:"bytes,4,opt,name=where,proto3" json:"where"`
	// split_pk
	SplitPk string `protobuf:"bytes,5,opt,name=split_pk,json=splitPk,proto3" json:"split_pk"`
	// mapping type
	MappingType BaseEnum_ColumnMapping `protobuf:"varint,6,opt,name=mapping_type,json=mappingType,proto3,enum=model.BaseEnum_ColumnMapping" json:"mapping_type"`
}

func (x *DBSource) Reset() {
	*x = DBSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_model_syncjob_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBSource) ProtoMessage() {}

func (x *DBSource) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_model_syncjob_db_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBSource.ProtoReflect.Descriptor instead.
func (*DBSource) Descriptor() ([]byte, []int) {
	return file_proto_types_model_syncjob_db_proto_rawDescGZIP(), []int{0}
}

func (x *DBSource) GetColumn() []*Column {
	if x != nil {
		return x.Column
	}
	return nil
}

func (x *DBSource) GetTable() []string {
	if x != nil {
		return x.Table
	}
	return nil
}

func (x *DBSource) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *DBSource) GetWhere() string {
	if x != nil {
		return x.Where
	}
	return ""
}

func (x *DBSource) GetSplitPk() string {
	if x != nil {
		return x.SplitPk
	}
	return ""
}

func (x *DBSource) GetMappingType() BaseEnum_ColumnMapping {
	if x != nil {
		return x.MappingType
	}
	return BaseEnum_ColumnUnset
}

type DBTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// column
	Column []*Column `protobuf:"bytes,1,rep,name=column,proto3" json:"column"`
	// table
	Table []string `protobuf:"bytes,2,rep,name=table,proto3" json:"table"`
	// schema
	Schema string `protobuf:"bytes,3,opt,name=schema,proto3" json:"schema"`
	// pre sql
	PreSql []string `protobuf:"bytes,4,rep,name=pre_sql,json=preSql,proto3" json:"pre_sql"`
	// post sql
	PostSql []string `protobuf:"bytes,5,rep,name=post_sql,json=postSql,proto3" json:"post_sql"`
	// write mode
	WriteMode DBTarget_WriteMode `protobuf:"varint,6,opt,name=write_mode,json=writeMode,proto3,enum=model.DBTarget_WriteMode" json:"write_mode"`
	// batch size
	BatchSize int32 `protobuf:"varint,7,opt,name=batch_size,json=batchSize,proto3" json:"batch_size"`
	// update key
	UpdateKey []string `protobuf:"bytes,8,rep,name=update_key,json=updateKey,proto3" json:"update_key"`
	// semantic
	Semantic DBTarget_Semantic `protobuf:"varint,9,opt,name=semantic,proto3,enum=model.DBTarget_Semantic" json:"semantic"`
	// with no lock
	WithNoLock string `protobuf:"bytes,10,opt,name=with_no_lock,json=withNoLock,proto3" json:"with_no_lock"`
}

func (x *DBTarget) Reset() {
	*x = DBTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_model_syncjob_db_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBTarget) ProtoMessage() {}

func (x *DBTarget) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_model_syncjob_db_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBTarget.ProtoReflect.Descriptor instead.
func (*DBTarget) Descriptor() ([]byte, []int) {
	return file_proto_types_model_syncjob_db_proto_rawDescGZIP(), []int{1}
}

func (x *DBTarget) GetColumn() []*Column {
	if x != nil {
		return x.Column
	}
	return nil
}

func (x *DBTarget) GetTable() []string {
	if x != nil {
		return x.Table
	}
	return nil
}

func (x *DBTarget) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *DBTarget) GetPreSql() []string {
	if x != nil {
		return x.PreSql
	}
	return nil
}

func (x *DBTarget) GetPostSql() []string {
	if x != nil {
		return x.PostSql
	}
	return nil
}

func (x *DBTarget) GetWriteMode() DBTarget_WriteMode {
	if x != nil {
		return x.WriteMode
	}
	return DBTarget_insert
}

func (x *DBTarget) GetBatchSize() int32 {
	if x != nil {
		return x.BatchSize
	}
	return 0
}

func (x *DBTarget) GetUpdateKey() []string {
	if x != nil {
		return x.UpdateKey
	}
	return nil
}

func (x *DBTarget) GetSemantic() DBTarget_Semantic {
	if x != nil {
		return x.Semantic
	}
	return DBTarget_AtLeastOnce
}

func (x *DBTarget) GetWithNoLock() string {
	if x != nil {
		return x.WithNoLock
	}
	return ""
}

var File_proto_types_model_syncjob_db_proto protoreflect.FileDescriptor

var file_proto_types_model_syncjob_db_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x6a, 0x6f, 0x62, 0x2f, 0x64, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x33, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x75, 0x33, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x6a, 0x6f, 0x62, 0x2f, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x79, 0x6e, 0x63,
	0x6a, 0x6f, 0x62, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xd2, 0x01, 0x0a, 0x08, 0x44, 0x42, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x25, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x06,
	0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x68, 0x65, 0x72, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x77, 0x68, 0x65, 0x72, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x70,
	0x6c, 0x69, 0x74, 0x5f, 0x70, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70,
	0x6c, 0x69, 0x74, 0x50, 0x6b, 0x12, 0x40, 0x0a, 0x0c, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x43, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x0b, 0x6d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x22, 0xc3, 0x03, 0x0a, 0x08, 0x44, 0x42, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x12, 0x25, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x72, 0x65,
	0x5f, 0x73, 0x71, 0x6c, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x53,
	0x71, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x71, 0x6c, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x53, 0x71, 0x6c, 0x12, 0x38, 0x0a,
	0x0a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x19, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x42, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x08, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69,
	0x63, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x44, 0x42, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x53, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69,
	0x63, 0x52, 0x08, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x12, 0x20, 0x0a, 0x0c, 0x77,
	0x69, 0x74, 0x68, 0x5f, 0x6e, 0x6f, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x77, 0x69, 0x74, 0x68, 0x4e, 0x6f, 0x4c, 0x6f, 0x63, 0x6b, 0x22, 0x30, 0x0a,
	0x09, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x69, 0x6e,
	0x73, 0x65, 0x72, 0x74, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x10, 0x02, 0x22,
	0x2c, 0x0a, 0x08, 0x53, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x12, 0x0f, 0x0a, 0x0b, 0x41,
	0x74, 0x4c, 0x65, 0x61, 0x73, 0x74, 0x4f, 0x6e, 0x63, 0x65, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b,
	0x45, 0x78, 0x61, 0x63, 0x74, 0x6c, 0x79, 0x4f, 0x6e, 0x63, 0x65, 0x10, 0x01, 0x42, 0x7d, 0x0a,
	0x2c, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6f, 0x6d, 0x6e, 0x69, 0x73, 0x2e, 0x67,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x62, 0x73, 0x79, 0x6e, 0x63, 0x6a, 0x6f, 0x62, 0x42, 0x0e, 0x50,
	0x42, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x64, 0x62, 0x50, 0x00, 0x5a,
	0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61,
	0x57, 0x6f, 0x72, 0x6b, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2f, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x78, 0x67, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x70, 0x62, 0x73, 0x79, 0x6e, 0x63, 0x6a, 0x6f, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_types_model_syncjob_db_proto_rawDescOnce sync.Once
	file_proto_types_model_syncjob_db_proto_rawDescData = file_proto_types_model_syncjob_db_proto_rawDesc
)

func file_proto_types_model_syncjob_db_proto_rawDescGZIP() []byte {
	file_proto_types_model_syncjob_db_proto_rawDescOnce.Do(func() {
		file_proto_types_model_syncjob_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_model_syncjob_db_proto_rawDescData)
	})
	return file_proto_types_model_syncjob_db_proto_rawDescData
}

var file_proto_types_model_syncjob_db_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_types_model_syncjob_db_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_types_model_syncjob_db_proto_goTypes = []interface{}{
	(DBTarget_WriteMode)(0),     // 0: model.DBTarget.WriteMode
	(DBTarget_Semantic)(0),      // 1: model.DBTarget.Semantic
	(*DBSource)(nil),            // 2: model.DBSource
	(*DBTarget)(nil),            // 3: model.DBTarget
	(*Column)(nil),              // 4: model.Column
	(BaseEnum_ColumnMapping)(0), // 5: model.BaseEnum.ColumnMapping
}
var file_proto_types_model_syncjob_db_proto_depIdxs = []int32{
	4, // 0: model.DBSource.column:type_name -> model.Column
	5, // 1: model.DBSource.mapping_type:type_name -> model.BaseEnum.ColumnMapping
	4, // 2: model.DBTarget.column:type_name -> model.Column
	0, // 3: model.DBTarget.write_mode:type_name -> model.DBTarget.WriteMode
	1, // 4: model.DBTarget.semantic:type_name -> model.DBTarget.Semantic
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_types_model_syncjob_db_proto_init() }
func file_proto_types_model_syncjob_db_proto_init() {
	if File_proto_types_model_syncjob_db_proto != nil {
		return
	}
	file_proto_types_model_syncjob_column_proto_init()
	file_proto_types_model_syncjob_baseenum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_types_model_syncjob_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBSource); i {
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
		file_proto_types_model_syncjob_db_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBTarget); i {
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
			RawDescriptor: file_proto_types_model_syncjob_db_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_model_syncjob_db_proto_goTypes,
		DependencyIndexes: file_proto_types_model_syncjob_db_proto_depIdxs,
		EnumInfos:         file_proto_types_model_syncjob_db_proto_enumTypes,
		MessageInfos:      file_proto_types_model_syncjob_db_proto_msgTypes,
	}.Build()
	File_proto_types_model_syncjob_db_proto = out.File
	file_proto_types_model_syncjob_db_proto_rawDesc = nil
	file_proto_types_model_syncjob_db_proto_goTypes = nil
	file_proto_types_model_syncjob_db_proto_depIdxs = nil
}
