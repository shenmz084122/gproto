// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: proto/types/model/syncjob/relationaldb.proto

package pbsyncjob

import (
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbdefaults"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbgosql"
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

type RelationaldbSource struct {
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
}

func (x *RelationaldbSource) Reset() {
	*x = RelationaldbSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_model_syncjob_relationaldb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationaldbSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationaldbSource) ProtoMessage() {}

func (x *RelationaldbSource) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_model_syncjob_relationaldb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationaldbSource.ProtoReflect.Descriptor instead.
func (*RelationaldbSource) Descriptor() ([]byte, []int) {
	return file_proto_types_model_syncjob_relationaldb_proto_rawDescGZIP(), []int{0}
}

func (x *RelationaldbSource) GetColumn() []*Column {
	if x != nil {
		return x.Column
	}
	return nil
}

func (x *RelationaldbSource) GetTable() []string {
	if x != nil {
		return x.Table
	}
	return nil
}

func (x *RelationaldbSource) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *RelationaldbSource) GetWhere() string {
	if x != nil {
		return x.Where
	}
	return ""
}

func (x *RelationaldbSource) GetSplitPk() string {
	if x != nil {
		return x.SplitPk
	}
	return ""
}

type RelationaldbTarget struct {
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
	WriteMode string `protobuf:"bytes,6,opt,name=write_mode,json=writeMode,proto3" json:"write_mode"`
	// batch size
	BatchSize int32 `protobuf:"varint,7,opt,name=batch_size,json=batchSize,proto3" json:"batch_size"`
	// update key
	UpdateKey []string `protobuf:"bytes,8,rep,name=update_key,json=updateKey,proto3" json:"update_key"`
	// mode
	Mode string `protobuf:"bytes,9,opt,name=mode,proto3" json:"mode"`
	// semantic
	Semantic string `protobuf:"bytes,10,opt,name=semantic,proto3" json:"semantic"`
	// with no lock
	WithNoLock string `protobuf:"bytes,11,opt,name=with_no_lock,json=withNoLock,proto3" json:"with_no_lock"`
}

func (x *RelationaldbTarget) Reset() {
	*x = RelationaldbTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_model_syncjob_relationaldb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationaldbTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationaldbTarget) ProtoMessage() {}

func (x *RelationaldbTarget) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_model_syncjob_relationaldb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationaldbTarget.ProtoReflect.Descriptor instead.
func (*RelationaldbTarget) Descriptor() ([]byte, []int) {
	return file_proto_types_model_syncjob_relationaldb_proto_rawDescGZIP(), []int{1}
}

func (x *RelationaldbTarget) GetColumn() []*Column {
	if x != nil {
		return x.Column
	}
	return nil
}

func (x *RelationaldbTarget) GetTable() []string {
	if x != nil {
		return x.Table
	}
	return nil
}

func (x *RelationaldbTarget) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *RelationaldbTarget) GetPreSql() []string {
	if x != nil {
		return x.PreSql
	}
	return nil
}

func (x *RelationaldbTarget) GetPostSql() []string {
	if x != nil {
		return x.PostSql
	}
	return nil
}

func (x *RelationaldbTarget) GetWriteMode() string {
	if x != nil {
		return x.WriteMode
	}
	return ""
}

func (x *RelationaldbTarget) GetBatchSize() int32 {
	if x != nil {
		return x.BatchSize
	}
	return 0
}

func (x *RelationaldbTarget) GetUpdateKey() []string {
	if x != nil {
		return x.UpdateKey
	}
	return nil
}

func (x *RelationaldbTarget) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *RelationaldbTarget) GetSemantic() string {
	if x != nil {
		return x.Semantic
	}
	return ""
}

func (x *RelationaldbTarget) GetWithNoLock() string {
	if x != nil {
		return x.WithNoLock
	}
	return ""
}

var File_proto_types_model_syncjob_relationaldb_proto protoreflect.FileDescriptor

var file_proto_types_model_syncjob_relationaldb_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x6a, 0x6f, 0x62, 0x2f, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x79, 0x75, 0x33, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x75, 0x33, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x73, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x75, 0x33, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x6a, 0x6f, 0x62, 0x2f, 0x63, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x64, 0x62, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x25,
	0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x06, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x68, 0x65, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x77, 0x68, 0x65, 0x72, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x70, 0x6c,
	0x69, 0x74, 0x5f, 0x70, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x6c,
	0x69, 0x74, 0x50, 0x6b, 0x22, 0xcc, 0x02, 0x0a, 0x12, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x64, 0x62, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x25, 0x0a, 0x06, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x12, 0x17, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x5f, 0x73, 0x71, 0x6c, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x53, 0x71, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f, 0x73,
	0x74, 0x5f, 0x73, 0x71, 0x6c, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x73,
	0x74, 0x53, 0x71, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x72, 0x69, 0x74, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4b, 0x65,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69,
	0x63, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69,
	0x63, 0x12, 0x20, 0x0a, 0x0c, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x6e, 0x6f, 0x5f, 0x6c, 0x6f, 0x63,
	0x6b, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x69, 0x74, 0x68, 0x4e, 0x6f, 0x4c,
	0x6f, 0x63, 0x6b, 0x42, 0x7d, 0x0a, 0x2c, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6f,
	0x6d, 0x6e, 0x69, 0x73, 0x2e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x62, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x62, 0x73, 0x79, 0x6e, 0x63,
	0x6a, 0x6f, 0x62, 0x42, 0x0e, 0x50, 0x42, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x64, 0x62, 0x50, 0x00, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x57, 0x6f, 0x72, 0x6b, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2f,
	0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x78, 0x67, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x70, 0x62, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x62, 0x73, 0x79, 0x6e, 0x63, 0x6a,
	0x6f, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_types_model_syncjob_relationaldb_proto_rawDescOnce sync.Once
	file_proto_types_model_syncjob_relationaldb_proto_rawDescData = file_proto_types_model_syncjob_relationaldb_proto_rawDesc
)

func file_proto_types_model_syncjob_relationaldb_proto_rawDescGZIP() []byte {
	file_proto_types_model_syncjob_relationaldb_proto_rawDescOnce.Do(func() {
		file_proto_types_model_syncjob_relationaldb_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_model_syncjob_relationaldb_proto_rawDescData)
	})
	return file_proto_types_model_syncjob_relationaldb_proto_rawDescData
}

var file_proto_types_model_syncjob_relationaldb_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_types_model_syncjob_relationaldb_proto_goTypes = []interface{}{
	(*RelationaldbSource)(nil), // 0: model.RelationaldbSource
	(*RelationaldbTarget)(nil), // 1: model.RelationaldbTarget
	(*Column)(nil),             // 2: model.Column
}
var file_proto_types_model_syncjob_relationaldb_proto_depIdxs = []int32{
	2, // 0: model.RelationaldbSource.column:type_name -> model.Column
	2, // 1: model.RelationaldbTarget.column:type_name -> model.Column
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_types_model_syncjob_relationaldb_proto_init() }
func file_proto_types_model_syncjob_relationaldb_proto_init() {
	if File_proto_types_model_syncjob_relationaldb_proto != nil {
		return
	}
	file_proto_types_model_syncjob_column_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_types_model_syncjob_relationaldb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationaldbSource); i {
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
		file_proto_types_model_syncjob_relationaldb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationaldbTarget); i {
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
			RawDescriptor: file_proto_types_model_syncjob_relationaldb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_model_syncjob_relationaldb_proto_goTypes,
		DependencyIndexes: file_proto_types_model_syncjob_relationaldb_proto_depIdxs,
		MessageInfos:      file_proto_types_model_syncjob_relationaldb_proto_msgTypes,
	}.Build()
	File_proto_types_model_syncjob_relationaldb_proto = out.File
	file_proto_types_model_syncjob_relationaldb_proto_rawDesc = nil
	file_proto_types_model_syncjob_relationaldb_proto_goTypes = nil
	file_proto_types_model_syncjob_relationaldb_proto_depIdxs = nil
}