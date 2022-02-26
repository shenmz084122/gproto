// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: proto/types/model/udf.proto

package pbmodel

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

type UDF_Type int32

const (
	UDF_TypeUnset UDF_Type = 0
	UDF_UDF       UDF_Type = 1
	UDF_UDTF      UDF_Type = 2
	UDF_UDTTF     UDF_Type = 3
)

// Enum value maps for UDF_Type.
var (
	UDF_Type_name = map[int32]string{
		0: "TypeUnset",
		1: "UDF",
		2: "UDTF",
		3: "UDTTF",
	}
	UDF_Type_value = map[string]int32{
		"TypeUnset": 0,
		"UDF":       1,
		"UDTF":      2,
		"UDTTF":     3,
	}
)

func (x UDF_Type) Enum() *UDF_Type {
	p := new(UDF_Type)
	*p = x
	return p
}

func (x UDF_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UDF_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_types_model_udf_proto_enumTypes[0].Descriptor()
}

func (UDF_Type) Type() protoreflect.EnumType {
	return &file_proto_types_model_udf_proto_enumTypes[0]
}

func (x UDF_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UDF_Type.Descriptor instead.
func (UDF_Type) EnumDescriptor() ([]byte, []int) {
	return file_proto_types_model_udf_proto_rawDescGZIP(), []int{0, 0}
}

type UDF_Language int32

const (
	UDF_LanguageUnset UDF_Language = 0
	UDF_Scala         UDF_Language = 1
	UDF_Java          UDF_Language = 2
	UDF_Python        UDF_Language = 3
)

// Enum value maps for UDF_Language.
var (
	UDF_Language_name = map[int32]string{
		0: "LanguageUnset",
		1: "Scala",
		2: "Java",
		3: "Python",
	}
	UDF_Language_value = map[string]int32{
		"LanguageUnset": 0,
		"Scala":         1,
		"Java":          2,
		"Python":        3,
	}
)

func (x UDF_Language) Enum() *UDF_Language {
	p := new(UDF_Language)
	*p = x
	return p
}

func (x UDF_Language) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UDF_Language) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_types_model_udf_proto_enumTypes[1].Descriptor()
}

func (UDF_Language) Type() protoreflect.EnumType {
	return &file_proto_types_model_udf_proto_enumTypes[1]
}

func (x UDF_Language) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UDF_Language.Descriptor instead.
func (UDF_Language) EnumDescriptor() ([]byte, []int) {
	return file_proto_types_model_udf_proto_rawDescGZIP(), []int{0, 1}
}

type UDF_Status int32

const (
	UDF_StatusUnset UDF_Status = 0
	UDF_Deleted     UDF_Status = 1
	UDF_Enabled     UDF_Status = 2
)

// Enum value maps for UDF_Status.
var (
	UDF_Status_name = map[int32]string{
		0: "StatusUnset",
		1: "Deleted",
		2: "Enabled",
	}
	UDF_Status_value = map[string]int32{
		"StatusUnset": 0,
		"Deleted":     1,
		"Enabled":     2,
	}
)

func (x UDF_Status) Enum() *UDF_Status {
	p := new(UDF_Status)
	*p = x
	return p
}

func (x UDF_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UDF_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_types_model_udf_proto_enumTypes[2].Descriptor()
}

func (UDF_Status) Type() protoreflect.EnumType {
	return &file_proto_types_model_udf_proto_enumTypes[2]
}

func (x UDF_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UDF_Status.Descriptor instead.
func (UDF_Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_types_model_udf_proto_rawDescGZIP(), []int{0, 2}
}

// UDF Schema
type UDF struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Space ID, workspace ID.
	SpaceId string `protobuf:"bytes,1,opt,name=space_id,json=spaceId,proto3" json:"space_id" gorm:"column:space_id;"`
	// ID, unique within a region.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id" gorm:"column:id;primaryKey;"`
	// UDF Name, unique within a space.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name" gorm:"column:name;"`
	// Desc is description this UDF.
	Desc string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc" gorm:"column:desc;"`
	// UDF status, Optional Values: 1 => "deleted", 2 => "enabled"
	Status UDF_Status `protobuf:"varint,5,opt,name=status,proto3,enum=model.UDF_Status" json:"status" gorm:"column:status;"`
	// UDF Type; Optional Values: 1=>UDF, 2=>UDTF 3=> UDTTF.
	Type UDF_Type `protobuf:"varint,6,opt,name=type,proto3,enum=model.UDF_Type" json:"type" gorm:"column:type;"`
	// language of UDF; Optional Values: 1 => Scala 2=> Java 3=> Python
	Language UDF_Language `protobuf:"varint,7,opt,name=language,proto3,enum=model.UDF_Language" json:"language" gorm:"column:language;"`
	// The id of resource. Used with language of JAVA.
	// Is required if language == 2
	FileId string `protobuf:"bytes,8,opt,name=file_id,json=fileId,proto3" json:"file_id" gorm:"column:file_id;"`
	// The code. Used with language of Python and Scala.
	// Is required if language == (1 or 3). Min Charset Length: 1, Max Charset Length: 20000.
	Code string `protobuf:"bytes,9,opt,name=code,proto3" json:"code" gorm:"column:code;"`
	// usage example for this udf
	UsageSample string `protobuf:"bytes,10,opt,name=usage_sample,json=usageSample,proto3" json:"usage_sample" gorm:"column:usage_sample;"`
	// Who created this udf.
	CreatedBy string `protobuf:"bytes,11,opt,name=created_by,json=createdBy,proto3" json:"created_by" gorm:"column:created_by"`
	// Timestamp of create time.
	Created int64 `protobuf:"varint,12,opt,name=created,proto3" json:"created" gorm:"column:created;autoCreateTime;"`
	// Timestamp of update time.
	Updated int64 `protobuf:"varint,13,opt,name=updated,proto3" json:"updated" gorm:"column:updated;autoUpdateTime;"`
}

func (x *UDF) Reset() {
	*x = UDF{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_model_udf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UDF) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UDF) ProtoMessage() {}

func (x *UDF) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_model_udf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UDF.ProtoReflect.Descriptor instead.
func (*UDF) Descriptor() ([]byte, []int) {
	return file_proto_types_model_udf_proto_rawDescGZIP(), []int{0}
}

func (x *UDF) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *UDF) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UDF) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UDF) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *UDF) GetStatus() UDF_Status {
	if x != nil {
		return x.Status
	}
	return UDF_StatusUnset
}

func (x *UDF) GetType() UDF_Type {
	if x != nil {
		return x.Type
	}
	return UDF_TypeUnset
}

func (x *UDF) GetLanguage() UDF_Language {
	if x != nil {
		return x.Language
	}
	return UDF_LanguageUnset
}

func (x *UDF) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *UDF) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UDF) GetUsageSample() string {
	if x != nil {
		return x.UsageSample
	}
	return ""
}

func (x *UDF) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *UDF) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *UDF) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

var File_proto_types_model_udf_proto protoreflect.FileDescriptor

var file_proto_types_model_udf_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x75, 0x64, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x79, 0x75, 0x33, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x06, 0x0a, 0x03, 0x55, 0x44,
	0x46, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x13, 0xe2, 0xdf, 0x1f, 0x0f, 0x12, 0x0d, 0xc2, 0x01, 0x0a, 0xf0, 0x01,
	0x14, 0xca, 0x02, 0x04, 0x77, 0x6b, 0x73, 0x2d, 0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x23, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0xe2,
	0xdf, 0x1f, 0x0f, 0x12, 0x0d, 0xc2, 0x01, 0x0a, 0xf0, 0x01, 0x14, 0xca, 0x02, 0x04, 0x75, 0x64,
	0x66, 0x2d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xe2, 0xdf, 0x1f, 0x0b, 0x12, 0x09, 0xc2, 0x01, 0x06, 0x90,
	0x02, 0x02, 0x98, 0x02, 0x40, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x64,
	0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xe2, 0xdf, 0x1f, 0x09, 0x12,
	0x07, 0xc2, 0x01, 0x04, 0xc8, 0x01, 0x80, 0x02, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x38,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55, 0x44, 0x46, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x42, 0x0d, 0xe2, 0xdf, 0x1f, 0x09, 0x12, 0x07, 0xda, 0x01, 0x04, 0x30, 0x00, 0x58, 0x01,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55,
	0x44, 0x46, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0d, 0xe2, 0xdf, 0x1f, 0x09, 0x12, 0x07, 0xda,
	0x01, 0x04, 0x30, 0x00, 0x58, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3e, 0x0a, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55, 0x44, 0x46, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x42, 0x0d, 0xe2, 0xdf, 0x1f, 0x09, 0x12, 0x07, 0xda, 0x01, 0x04, 0x30, 0x00,
	0x58, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x43, 0x0a, 0x07,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2a, 0xe2,
	0xdf, 0x1f, 0x13, 0x0a, 0x11, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12,
	0x05, 0xda, 0x01, 0x02, 0x18, 0x02, 0xe2, 0xdf, 0x1f, 0x0f, 0x12, 0x0d, 0xc2, 0x01, 0x0a, 0xf0,
	0x01, 0x14, 0xca, 0x02, 0x04, 0x72, 0x65, 0x73, 0x2d, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49,
	0x64, 0x12, 0x3e, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x2a, 0xe2, 0xdf, 0x1f, 0x15, 0x0a, 0x13, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x12, 0x07, 0xda, 0x01, 0x04, 0x4a, 0x02, 0x01, 0x03, 0xe2, 0xdf, 0x1f, 0x0d, 0x12, 0x0b,
	0xc2, 0x01, 0x08, 0xb0, 0x01, 0x00, 0xc8, 0x01, 0xa0, 0x9c, 0x01, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x30, 0x0a, 0x0c, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xe2, 0xdf, 0x1f, 0x09, 0x12, 0x07, 0xc2,
	0x01, 0x04, 0xc8, 0x01, 0xd0, 0x0f, 0x52, 0x0b, 0x75, 0x73, 0x61, 0x67, 0x65, 0x53, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62,
	0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xe2, 0xdf, 0x1f, 0x0b, 0x12, 0x09, 0xc2,
	0x01, 0x06, 0x80, 0x02, 0x00, 0x88, 0x02, 0x41, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x25, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x0b, 0xe2, 0xdf, 0x1f, 0x07, 0x12, 0x05, 0xb2, 0x01, 0x02, 0x30,
	0x00, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x07, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0b, 0xe2, 0xdf, 0x1f,
	0x07, 0x12, 0x05, 0xb2, 0x01, 0x02, 0x30, 0x00, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x22, 0x33, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x79, 0x70,
	0x65, 0x55, 0x6e, 0x73, 0x65, 0x74, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x55, 0x44, 0x46, 0x10,
	0x01, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x44, 0x54, 0x46, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x55,
	0x44, 0x54, 0x54, 0x46, 0x10, 0x03, 0x22, 0x3e, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x55, 0x6e,
	0x73, 0x65, 0x74, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x4a, 0x61, 0x76, 0x61, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x79,
	0x74, 0x68, 0x6f, 0x6e, 0x10, 0x03, 0x22, 0x33, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x6e, 0x73, 0x65, 0x74, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x10, 0x02, 0x42, 0x65, 0x0a, 0x22, 0x63,
	0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6f, 0x6d, 0x6e, 0x69, 0x73, 0x2e, 0x67, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x42, 0x0a, 0x50, 0x42, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x55, 0x44, 0x46, 0x50, 0x00, 0x5a,
	0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61,
	0x57, 0x6f, 0x72, 0x6b, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2f, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x78, 0x67, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_types_model_udf_proto_rawDescOnce sync.Once
	file_proto_types_model_udf_proto_rawDescData = file_proto_types_model_udf_proto_rawDesc
)

func file_proto_types_model_udf_proto_rawDescGZIP() []byte {
	file_proto_types_model_udf_proto_rawDescOnce.Do(func() {
		file_proto_types_model_udf_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_model_udf_proto_rawDescData)
	})
	return file_proto_types_model_udf_proto_rawDescData
}

var file_proto_types_model_udf_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_proto_types_model_udf_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_types_model_udf_proto_goTypes = []interface{}{
	(UDF_Type)(0),     // 0: model.UDF.Type
	(UDF_Language)(0), // 1: model.UDF.Language
	(UDF_Status)(0),   // 2: model.UDF.Status
	(*UDF)(nil),       // 3: model.UDF
}
var file_proto_types_model_udf_proto_depIdxs = []int32{
	2, // 0: model.UDF.status:type_name -> model.UDF.Status
	0, // 1: model.UDF.type:type_name -> model.UDF.Type
	1, // 2: model.UDF.language:type_name -> model.UDF.Language
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_types_model_udf_proto_init() }
func file_proto_types_model_udf_proto_init() {
	if File_proto_types_model_udf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_types_model_udf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UDF); i {
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
			RawDescriptor: file_proto_types_model_udf_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_model_udf_proto_goTypes,
		DependencyIndexes: file_proto_types_model_udf_proto_depIdxs,
		EnumInfos:         file_proto_types_model_udf_proto_enumTypes,
		MessageInfos:      file_proto_types_model_udf_proto_msgTypes,
	}.Build()
	File_proto_types_model_udf_proto = out.File
	file_proto_types_model_udf_proto_rawDesc = nil
	file_proto_types_model_udf_proto_goTypes = nil
	file_proto_types_model_udf_proto_depIdxs = nil
}
