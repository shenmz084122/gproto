// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: proto/types/model/syncjob/elasticsearch.proto

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

type ElasticSearchSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// column
	Column []*Column `protobuf:"bytes,1,rep,name=column,proto3" json:"column"`
	// version
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version"`
	// index
	Index string `protobuf:"bytes,3,opt,name=index,proto3" json:"index"`
	// batch size
	BatchSize int32 `protobuf:"varint,4,opt,name=batch_size,json=batchSize,proto3" json:"batch_size"`
}

func (x *ElasticSearchSource) Reset() {
	*x = ElasticSearchSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_model_syncjob_elasticsearch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElasticSearchSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElasticSearchSource) ProtoMessage() {}

func (x *ElasticSearchSource) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_model_syncjob_elasticsearch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElasticSearchSource.ProtoReflect.Descriptor instead.
func (*ElasticSearchSource) Descriptor() ([]byte, []int) {
	return file_proto_types_model_syncjob_elasticsearch_proto_rawDescGZIP(), []int{0}
}

func (x *ElasticSearchSource) GetColumn() []*Column {
	if x != nil {
		return x.Column
	}
	return nil
}

func (x *ElasticSearchSource) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ElasticSearchSource) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

func (x *ElasticSearchSource) GetBatchSize() int32 {
	if x != nil {
		return x.BatchSize
	}
	return 0
}

type ElasticSearchTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// column
	Column []*Column `protobuf:"bytes,1,rep,name=column,proto3" json:"column"`
	// version
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version"`
	// index
	Index string `protobuf:"bytes,3,opt,name=index,proto3" json:"index"`
	// batch size
	BatchSize int32 `protobuf:"varint,4,opt,name=batch_size,json=batchSize,proto3" json:"batch_size"`
	// key delimiter
	KeyDelimiter string `protobuf:"bytes,5,opt,name=key_delimiter,json=keyDelimiter,proto3" json:"key_delimiter"`
}

func (x *ElasticSearchTarget) Reset() {
	*x = ElasticSearchTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_model_syncjob_elasticsearch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElasticSearchTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElasticSearchTarget) ProtoMessage() {}

func (x *ElasticSearchTarget) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_model_syncjob_elasticsearch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElasticSearchTarget.ProtoReflect.Descriptor instead.
func (*ElasticSearchTarget) Descriptor() ([]byte, []int) {
	return file_proto_types_model_syncjob_elasticsearch_proto_rawDescGZIP(), []int{1}
}

func (x *ElasticSearchTarget) GetColumn() []*Column {
	if x != nil {
		return x.Column
	}
	return nil
}

func (x *ElasticSearchTarget) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ElasticSearchTarget) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

func (x *ElasticSearchTarget) GetBatchSize() int32 {
	if x != nil {
		return x.BatchSize
	}
	return 0
}

func (x *ElasticSearchTarget) GetKeyDelimiter() string {
	if x != nil {
		return x.KeyDelimiter
	}
	return ""
}

var File_proto_types_model_syncjob_elasticsearch_proto protoreflect.FileDescriptor

var file_proto_types_model_syncjob_elasticsearch_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x6a, 0x6f, 0x62, 0x2f, 0x65, 0x6c, 0x61, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x79, 0x75, 0x33, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73,
	0x79, 0x6e, 0x63, 0x6a, 0x6f, 0x62, 0x2f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x13, 0x45, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69, 0x7a,
	0x65, 0x22, 0xb0, 0x01, 0x0a, 0x13, 0x45, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x25, 0x0a, 0x06, 0x63, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x6b, 0x65, 0x79, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6b, 0x65, 0x79, 0x44, 0x65, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x65, 0x72, 0x42, 0x7e, 0x0a, 0x2c, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x6f, 0x6d, 0x6e, 0x69, 0x73, 0x2e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x62, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x62, 0x73, 0x79, 0x6e,
	0x63, 0x6a, 0x6f, 0x62, 0x42, 0x0f, 0x50, 0x42, 0x45, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x00, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x57, 0x6f, 0x72, 0x6b, 0x62, 0x65, 0x6e, 0x63,
	0x68, 0x2f, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x78, 0x67, 0x6f, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x70, 0x62, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x62, 0x73, 0x79, 0x6e,
	0x63, 0x6a, 0x6f, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_types_model_syncjob_elasticsearch_proto_rawDescOnce sync.Once
	file_proto_types_model_syncjob_elasticsearch_proto_rawDescData = file_proto_types_model_syncjob_elasticsearch_proto_rawDesc
)

func file_proto_types_model_syncjob_elasticsearch_proto_rawDescGZIP() []byte {
	file_proto_types_model_syncjob_elasticsearch_proto_rawDescOnce.Do(func() {
		file_proto_types_model_syncjob_elasticsearch_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_model_syncjob_elasticsearch_proto_rawDescData)
	})
	return file_proto_types_model_syncjob_elasticsearch_proto_rawDescData
}

var file_proto_types_model_syncjob_elasticsearch_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_types_model_syncjob_elasticsearch_proto_goTypes = []interface{}{
	(*ElasticSearchSource)(nil), // 0: model.ElasticSearchSource
	(*ElasticSearchTarget)(nil), // 1: model.ElasticSearchTarget
	(*Column)(nil),              // 2: model.Column
}
var file_proto_types_model_syncjob_elasticsearch_proto_depIdxs = []int32{
	2, // 0: model.ElasticSearchSource.column:type_name -> model.Column
	2, // 1: model.ElasticSearchTarget.column:type_name -> model.Column
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_types_model_syncjob_elasticsearch_proto_init() }
func file_proto_types_model_syncjob_elasticsearch_proto_init() {
	if File_proto_types_model_syncjob_elasticsearch_proto != nil {
		return
	}
	file_proto_types_model_syncjob_column_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_types_model_syncjob_elasticsearch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElasticSearchSource); i {
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
		file_proto_types_model_syncjob_elasticsearch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElasticSearchTarget); i {
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
			RawDescriptor: file_proto_types_model_syncjob_elasticsearch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_model_syncjob_elasticsearch_proto_goTypes,
		DependencyIndexes: file_proto_types_model_syncjob_elasticsearch_proto_depIdxs,
		MessageInfos:      file_proto_types_model_syncjob_elasticsearch_proto_msgTypes,
	}.Build()
	File_proto_types_model_syncjob_elasticsearch_proto = out.File
	file_proto_types_model_syncjob_elasticsearch_proto_rawDesc = nil
	file_proto_types_model_syncjob_elasticsearch_proto_goTypes = nil
	file_proto_types_model_syncjob_elasticsearch_proto_depIdxs = nil
}
