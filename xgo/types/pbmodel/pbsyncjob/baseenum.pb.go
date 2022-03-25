// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: proto/types/model/syncjob/baseenum.proto

package pbsyncjob

import (
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

type BaseEnum_PartitionType int32

const (
	BaseEnum_DAY    BaseEnum_PartitionType = 0
	BaseEnum_HOUR   BaseEnum_PartitionType = 1
	BaseEnum_MINUTE BaseEnum_PartitionType = 2
)

// Enum value maps for BaseEnum_PartitionType.
var (
	BaseEnum_PartitionType_name = map[int32]string{
		0: "DAY",
		1: "HOUR",
		2: "MINUTE",
	}
	BaseEnum_PartitionType_value = map[string]int32{
		"DAY":    0,
		"HOUR":   1,
		"MINUTE": 2,
	}
)

func (x BaseEnum_PartitionType) Enum() *BaseEnum_PartitionType {
	p := new(BaseEnum_PartitionType)
	*p = x
	return p
}

func (x BaseEnum_PartitionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BaseEnum_PartitionType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_types_model_syncjob_baseenum_proto_enumTypes[0].Descriptor()
}

func (BaseEnum_PartitionType) Type() protoreflect.EnumType {
	return &file_proto_types_model_syncjob_baseenum_proto_enumTypes[0]
}

func (x BaseEnum_PartitionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BaseEnum_PartitionType.Descriptor instead.
func (BaseEnum_PartitionType) EnumDescriptor() ([]byte, []int) {
	return file_proto_types_model_syncjob_baseenum_proto_rawDescGZIP(), []int{0, 0}
}

type BaseEnum_WriteMode int32

const (
	BaseEnum_append    BaseEnum_WriteMode = 0
	BaseEnum_overwrite BaseEnum_WriteMode = 1
)

// Enum value maps for BaseEnum_WriteMode.
var (
	BaseEnum_WriteMode_name = map[int32]string{
		0: "append",
		1: "overwrite",
	}
	BaseEnum_WriteMode_value = map[string]int32{
		"append":    0,
		"overwrite": 1,
	}
)

func (x BaseEnum_WriteMode) Enum() *BaseEnum_WriteMode {
	p := new(BaseEnum_WriteMode)
	*p = x
	return p
}

func (x BaseEnum_WriteMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BaseEnum_WriteMode) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_types_model_syncjob_baseenum_proto_enumTypes[1].Descriptor()
}

func (BaseEnum_WriteMode) Type() protoreflect.EnumType {
	return &file_proto_types_model_syncjob_baseenum_proto_enumTypes[1]
}

func (x BaseEnum_WriteMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BaseEnum_WriteMode.Descriptor instead.
func (BaseEnum_WriteMode) EnumDescriptor() ([]byte, []int) {
	return file_proto_types_model_syncjob_baseenum_proto_rawDescGZIP(), []int{0, 1}
}

type BaseEnum_FileType int32

const (
	BaseEnum_text    BaseEnum_FileType = 0
	BaseEnum_orc     BaseEnum_FileType = 1
	BaseEnum_parquet BaseEnum_FileType = 2
)

// Enum value maps for BaseEnum_FileType.
var (
	BaseEnum_FileType_name = map[int32]string{
		0: "text",
		1: "orc",
		2: "parquet",
	}
	BaseEnum_FileType_value = map[string]int32{
		"text":    0,
		"orc":     1,
		"parquet": 2,
	}
)

func (x BaseEnum_FileType) Enum() *BaseEnum_FileType {
	p := new(BaseEnum_FileType)
	*p = x
	return p
}

func (x BaseEnum_FileType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BaseEnum_FileType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_types_model_syncjob_baseenum_proto_enumTypes[2].Descriptor()
}

func (BaseEnum_FileType) Type() protoreflect.EnumType {
	return &file_proto_types_model_syncjob_baseenum_proto_enumTypes[2]
}

func (x BaseEnum_FileType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BaseEnum_FileType.Descriptor instead.
func (BaseEnum_FileType) EnumDescriptor() ([]byte, []int) {
	return file_proto_types_model_syncjob_baseenum_proto_rawDescGZIP(), []int{0, 2}
}

type BaseEnum_CompressType int32

const (
	BaseEnum_GZIP   BaseEnum_CompressType = 0
	BaseEnum_BZIP2  BaseEnum_CompressType = 1
	BaseEnum_SNAPPY BaseEnum_CompressType = 2
	BaseEnum_BZIP   BaseEnum_CompressType = 3
	BaseEnum_LZ4    BaseEnum_CompressType = 4
	BaseEnum_LZO    BaseEnum_CompressType = 5
)

// Enum value maps for BaseEnum_CompressType.
var (
	BaseEnum_CompressType_name = map[int32]string{
		0: "GZIP",
		1: "BZIP2",
		2: "SNAPPY",
		3: "BZIP",
		4: "LZ4",
		5: "LZO",
	}
	BaseEnum_CompressType_value = map[string]int32{
		"GZIP":   0,
		"BZIP2":  1,
		"SNAPPY": 2,
		"BZIP":   3,
		"LZ4":    4,
		"LZO":    5,
	}
)

func (x BaseEnum_CompressType) Enum() *BaseEnum_CompressType {
	p := new(BaseEnum_CompressType)
	*p = x
	return p
}

func (x BaseEnum_CompressType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BaseEnum_CompressType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_types_model_syncjob_baseenum_proto_enumTypes[3].Descriptor()
}

func (BaseEnum_CompressType) Type() protoreflect.EnumType {
	return &file_proto_types_model_syncjob_baseenum_proto_enumTypes[3]
}

func (x BaseEnum_CompressType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BaseEnum_CompressType.Descriptor instead.
func (BaseEnum_CompressType) EnumDescriptor() ([]byte, []int) {
	return file_proto_types_model_syncjob_baseenum_proto_rawDescGZIP(), []int{0, 3}
}

type BaseEnum_Encoding int32

const (
	BaseEnum_UTF8 BaseEnum_Encoding = 0
	BaseEnum_GBK  BaseEnum_Encoding = 1
)

// Enum value maps for BaseEnum_Encoding.
var (
	BaseEnum_Encoding_name = map[int32]string{
		0: "UTF8",
		1: "GBK",
	}
	BaseEnum_Encoding_value = map[string]int32{
		"UTF8": 0,
		"GBK":  1,
	}
)

func (x BaseEnum_Encoding) Enum() *BaseEnum_Encoding {
	p := new(BaseEnum_Encoding)
	*p = x
	return p
}

func (x BaseEnum_Encoding) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BaseEnum_Encoding) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_types_model_syncjob_baseenum_proto_enumTypes[4].Descriptor()
}

func (BaseEnum_Encoding) Type() protoreflect.EnumType {
	return &file_proto_types_model_syncjob_baseenum_proto_enumTypes[4]
}

func (x BaseEnum_Encoding) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BaseEnum_Encoding.Descriptor instead.
func (BaseEnum_Encoding) EnumDescriptor() ([]byte, []int) {
	return file_proto_types_model_syncjob_baseenum_proto_rawDescGZIP(), []int{0, 4}
}

type BaseEnum_ColumnMapping int32

const (
	BaseEnum_ColumnUnset BaseEnum_ColumnMapping = 0
	BaseEnum_Name        BaseEnum_ColumnMapping = 1
	BaseEnum_Row         BaseEnum_ColumnMapping = 2
	BaseEnum_Auto        BaseEnum_ColumnMapping = 3
)

// Enum value maps for BaseEnum_ColumnMapping.
var (
	BaseEnum_ColumnMapping_name = map[int32]string{
		0: "ColumnUnset",
		1: "Name",
		2: "Row",
		3: "Auto",
	}
	BaseEnum_ColumnMapping_value = map[string]int32{
		"ColumnUnset": 0,
		"Name":        1,
		"Row":         2,
		"Auto":        3,
	}
)

func (x BaseEnum_ColumnMapping) Enum() *BaseEnum_ColumnMapping {
	p := new(BaseEnum_ColumnMapping)
	*p = x
	return p
}

func (x BaseEnum_ColumnMapping) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BaseEnum_ColumnMapping) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_types_model_syncjob_baseenum_proto_enumTypes[5].Descriptor()
}

func (BaseEnum_ColumnMapping) Type() protoreflect.EnumType {
	return &file_proto_types_model_syncjob_baseenum_proto_enumTypes[5]
}

func (x BaseEnum_ColumnMapping) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BaseEnum_ColumnMapping.Descriptor instead.
func (BaseEnum_ColumnMapping) EnumDescriptor() ([]byte, []int) {
	return file_proto_types_model_syncjob_baseenum_proto_rawDescGZIP(), []int{0, 5}
}

type BaseEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BaseEnum) Reset() {
	*x = BaseEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_model_syncjob_baseenum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseEnum) ProtoMessage() {}

func (x *BaseEnum) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_model_syncjob_baseenum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseEnum.ProtoReflect.Descriptor instead.
func (*BaseEnum) Descriptor() ([]byte, []int) {
	return file_proto_types_model_syncjob_baseenum_proto_rawDescGZIP(), []int{0}
}

var File_proto_types_model_syncjob_baseenum_proto protoreflect.FileDescriptor

var file_proto_types_model_syncjob_baseenum_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x6a, 0x6f, 0x62, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x22, 0xb9, 0x02, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x2e,
	0x0a, 0x0d, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x07, 0x0a, 0x03, 0x44, 0x41, 0x59, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f, 0x55, 0x52,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x49, 0x4e, 0x55, 0x54, 0x45, 0x10, 0x02, 0x22, 0x26,
	0x0a, 0x09, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x61,
	0x70, 0x70, 0x65, 0x6e, 0x64, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x6f, 0x76, 0x65, 0x72, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x10, 0x01, 0x22, 0x2a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03,
	0x6f, 0x72, 0x63, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x71, 0x75, 0x65, 0x74,
	0x10, 0x02, 0x22, 0x4b, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x5a, 0x49, 0x50, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x42, 0x5a, 0x49, 0x50, 0x32, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x4e, 0x41, 0x50, 0x50,
	0x59, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x5a, 0x49, 0x50, 0x10, 0x03, 0x12, 0x07, 0x0a,
	0x03, 0x4c, 0x5a, 0x34, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x5a, 0x4f, 0x10, 0x05, 0x22,
	0x1d, 0x0a, 0x08, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x08, 0x0a, 0x04, 0x55,
	0x54, 0x46, 0x38, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x42, 0x4b, 0x10, 0x01, 0x22, 0x3d,
	0x0a, 0x0d, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12,
	0x0f, 0x0a, 0x0b, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x55, 0x6e, 0x73, 0x65, 0x74, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x6f,
	0x77, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x6f, 0x10, 0x03, 0x42, 0x79, 0x0a,
	0x2c, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6f, 0x6d, 0x6e, 0x69, 0x73, 0x2e, 0x67,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x62, 0x73, 0x79, 0x6e, 0x63, 0x6a, 0x6f, 0x62, 0x42, 0x0a, 0x50,
	0x42, 0x42, 0x61, 0x73, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x50, 0x00, 0x5a, 0x3b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x57, 0x6f, 0x72, 0x6b,
	0x62, 0x65, 0x6e, 0x63, 0x68, 0x2f, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x78, 0x67, 0x6f,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70,
	0x62, 0x73, 0x79, 0x6e, 0x63, 0x6a, 0x6f, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_types_model_syncjob_baseenum_proto_rawDescOnce sync.Once
	file_proto_types_model_syncjob_baseenum_proto_rawDescData = file_proto_types_model_syncjob_baseenum_proto_rawDesc
)

func file_proto_types_model_syncjob_baseenum_proto_rawDescGZIP() []byte {
	file_proto_types_model_syncjob_baseenum_proto_rawDescOnce.Do(func() {
		file_proto_types_model_syncjob_baseenum_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_model_syncjob_baseenum_proto_rawDescData)
	})
	return file_proto_types_model_syncjob_baseenum_proto_rawDescData
}

var file_proto_types_model_syncjob_baseenum_proto_enumTypes = make([]protoimpl.EnumInfo, 6)
var file_proto_types_model_syncjob_baseenum_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_types_model_syncjob_baseenum_proto_goTypes = []interface{}{
	(BaseEnum_PartitionType)(0), // 0: model.BaseEnum.PartitionType
	(BaseEnum_WriteMode)(0),     // 1: model.BaseEnum.WriteMode
	(BaseEnum_FileType)(0),      // 2: model.BaseEnum.FileType
	(BaseEnum_CompressType)(0),  // 3: model.BaseEnum.CompressType
	(BaseEnum_Encoding)(0),      // 4: model.BaseEnum.Encoding
	(BaseEnum_ColumnMapping)(0), // 5: model.BaseEnum.ColumnMapping
	(*BaseEnum)(nil),            // 6: model.BaseEnum
}
var file_proto_types_model_syncjob_baseenum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_types_model_syncjob_baseenum_proto_init() }
func file_proto_types_model_syncjob_baseenum_proto_init() {
	if File_proto_types_model_syncjob_baseenum_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_types_model_syncjob_baseenum_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseEnum); i {
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
			RawDescriptor: file_proto_types_model_syncjob_baseenum_proto_rawDesc,
			NumEnums:      6,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_model_syncjob_baseenum_proto_goTypes,
		DependencyIndexes: file_proto_types_model_syncjob_baseenum_proto_depIdxs,
		EnumInfos:         file_proto_types_model_syncjob_baseenum_proto_enumTypes,
		MessageInfos:      file_proto_types_model_syncjob_baseenum_proto_msgTypes,
	}.Build()
	File_proto_types_model_syncjob_baseenum_proto = out.File
	file_proto_types_model_syncjob_baseenum_proto_rawDesc = nil
	file_proto_types_model_syncjob_baseenum_proto_goTypes = nil
	file_proto_types_model_syncjob_baseenum_proto_depIdxs = nil
}
