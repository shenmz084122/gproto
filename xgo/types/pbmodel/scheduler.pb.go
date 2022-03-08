// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: proto/types/model/scheduler.proto

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

type StreamJobEvent_Type int32

const (
	StreamJobEvent_TypeUnset StreamJobEvent_Type = 0
	StreamJobEvent_Job       StreamJobEvent_Type = 1
	StreamJobEvent_Instance  StreamJobEvent_Type = 2
)

// Enum value maps for StreamJobEvent_Type.
var (
	StreamJobEvent_Type_name = map[int32]string{
		0: "TypeUnset",
		1: "Job",
		2: "Instance",
	}
	StreamJobEvent_Type_value = map[string]int32{
		"TypeUnset": 0,
		"Job":       1,
		"Instance":  2,
	}
)

func (x StreamJobEvent_Type) Enum() *StreamJobEvent_Type {
	p := new(StreamJobEvent_Type)
	*p = x
	return p
}

func (x StreamJobEvent_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StreamJobEvent_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_types_model_scheduler_proto_enumTypes[0].Descriptor()
}

func (StreamJobEvent_Type) Type() protoreflect.EnumType {
	return &file_proto_types_model_scheduler_proto_enumTypes[0]
}

func (x StreamJobEvent_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StreamJobEvent_Type.Descriptor instead.
func (StreamJobEvent_Type) EnumDescriptor() ([]byte, []int) {
	return file_proto_types_model_scheduler_proto_rawDescGZIP(), []int{0, 0}
}

type StreamJobEvent_Action int32

const (
	StreamJobEvent_ActionUnset StreamJobEvent_Action = 0
	StreamJobEvent_Create      StreamJobEvent_Action = 1
	StreamJobEvent_Submit      StreamJobEvent_Action = 2
	StreamJobEvent_Check       StreamJobEvent_Action = 4
	StreamJobEvent_Retry       StreamJobEvent_Action = 3
)

// Enum value maps for StreamJobEvent_Action.
var (
	StreamJobEvent_Action_name = map[int32]string{
		0: "ActionUnset",
		1: "Create",
		2: "Submit",
		4: "Check",
		3: "Retry",
	}
	StreamJobEvent_Action_value = map[string]int32{
		"ActionUnset": 0,
		"Create":      1,
		"Submit":      2,
		"Check":       4,
		"Retry":       3,
	}
)

func (x StreamJobEvent_Action) Enum() *StreamJobEvent_Action {
	p := new(StreamJobEvent_Action)
	*p = x
	return p
}

func (x StreamJobEvent_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StreamJobEvent_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_types_model_scheduler_proto_enumTypes[1].Descriptor()
}

func (StreamJobEvent_Action) Type() protoreflect.EnumType {
	return &file_proto_types_model_scheduler_proto_enumTypes[1]
}

func (x StreamJobEvent_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StreamJobEvent_Action.Descriptor instead.
func (StreamJobEvent_Action) EnumDescriptor() ([]byte, []int) {
	return file_proto_types_model_scheduler_proto_rawDescGZIP(), []int{0, 1}
}

// StreamJobEvent is the event message used in scheduler queue.
type StreamJobEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The property of stream job.
	Property *StreamJobProperty `protobuf:"bytes,1,opt,name=property,proto3" json:"property,omitempty"`
	// The internal access address of flink cluster.
	FlinkUrl string `protobuf:"bytes,2,opt,name=flink_url,json=flinkUrl,proto3" json:"flink_url,omitempty"`
	// The version of flink cluster.
	FlinkVersion string `protobuf:"bytes,3,opt,name=flink_version,json=flinkVersion,proto3" json:"flink_version,omitempty"`
	// The number of times the stream instance was executed.
	Retries int32 `protobuf:"varint,4,opt,name=retries,proto3" json:"retries,omitempty"`
	// The start execution time of the stream instance.
	Started int64 `protobuf:"varint,5,opt,name=started,proto3" json:"started,omitempty"`
}

func (x *StreamJobEvent) Reset() {
	*x = StreamJobEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_model_scheduler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamJobEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamJobEvent) ProtoMessage() {}

func (x *StreamJobEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_model_scheduler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamJobEvent.ProtoReflect.Descriptor instead.
func (*StreamJobEvent) Descriptor() ([]byte, []int) {
	return file_proto_types_model_scheduler_proto_rawDescGZIP(), []int{0}
}

func (x *StreamJobEvent) GetProperty() *StreamJobProperty {
	if x != nil {
		return x.Property
	}
	return nil
}

func (x *StreamJobEvent) GetFlinkUrl() string {
	if x != nil {
		return x.FlinkUrl
	}
	return ""
}

func (x *StreamJobEvent) GetFlinkVersion() string {
	if x != nil {
		return x.FlinkVersion
	}
	return ""
}

func (x *StreamJobEvent) GetRetries() int32 {
	if x != nil {
		return x.Retries
	}
	return 0
}

func (x *StreamJobEvent) GetStarted() int64 {
	if x != nil {
		return x.Started
	}
	return 0
}

var File_proto_types_model_scheduler_proto protoreflect.FileDescriptor

var file_proto_types_model_scheduler_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x33, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x75, 0x33, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x02, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4a, 0x6f,
	0x62, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x41, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4a, 0x6f, 0x62, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x42, 0x0b, 0xe2, 0xdf, 0x1f, 0x07, 0x12, 0x05, 0xe2, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x28, 0x0a, 0x09, 0x66, 0x6c, 0x69,
	0x6e, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xe2, 0xdf,
	0x1f, 0x07, 0x12, 0x05, 0xc2, 0x01, 0x02, 0x22, 0x00, 0x52, 0x08, 0x66, 0x6c, 0x69, 0x6e, 0x6b,
	0x55, 0x72, 0x6c, 0x12, 0x30, 0x0a, 0x0d, 0x66, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xe2, 0xdf, 0x1f, 0x07,
	0x12, 0x05, 0xc2, 0x01, 0x02, 0x22, 0x00, 0x52, 0x0c, 0x66, 0x6c, 0x69, 0x6e, 0x6b, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x04, 0xe2, 0xdf, 0x1f, 0x00, 0x52, 0x07, 0x72, 0x65,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x42, 0x04, 0xe2, 0xdf, 0x1f, 0x00, 0x52, 0x07, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x64, 0x22, 0x2c, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a,
	0x09, 0x54, 0x79, 0x70, 0x65, 0x55, 0x6e, 0x73, 0x65, 0x74, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03,
	0x4a, 0x6f, 0x62, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x10, 0x02, 0x22, 0x47, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x0a,
	0x0b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x73, 0x65, 0x74, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x10,
	0x04, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x65, 0x74, 0x72, 0x79, 0x10, 0x03, 0x42, 0x6b, 0x0a, 0x22,
	0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6f, 0x6d, 0x6e, 0x69, 0x73, 0x2e, 0x67, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x42, 0x10, 0x50, 0x42, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x50, 0x00, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x57, 0x6f, 0x72, 0x6b, 0x62, 0x65, 0x6e, 0x63, 0x68,
	0x2f, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x78, 0x67, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x70, 0x62, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_types_model_scheduler_proto_rawDescOnce sync.Once
	file_proto_types_model_scheduler_proto_rawDescData = file_proto_types_model_scheduler_proto_rawDesc
)

func file_proto_types_model_scheduler_proto_rawDescGZIP() []byte {
	file_proto_types_model_scheduler_proto_rawDescOnce.Do(func() {
		file_proto_types_model_scheduler_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_model_scheduler_proto_rawDescData)
	})
	return file_proto_types_model_scheduler_proto_rawDescData
}

var file_proto_types_model_scheduler_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_types_model_scheduler_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_types_model_scheduler_proto_goTypes = []interface{}{
	(StreamJobEvent_Type)(0),   // 0: model.StreamJobEvent.Type
	(StreamJobEvent_Action)(0), // 1: model.StreamJobEvent.Action
	(*StreamJobEvent)(nil),     // 2: model.StreamJobEvent
	(*StreamJobProperty)(nil),  // 3: model.StreamJobProperty
}
var file_proto_types_model_scheduler_proto_depIdxs = []int32{
	3, // 0: model.StreamJobEvent.property:type_name -> model.StreamJobProperty
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_types_model_scheduler_proto_init() }
func file_proto_types_model_scheduler_proto_init() {
	if File_proto_types_model_scheduler_proto != nil {
		return
	}
	file_proto_types_model_stream_job_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_types_model_scheduler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamJobEvent); i {
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
			RawDescriptor: file_proto_types_model_scheduler_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_model_scheduler_proto_goTypes,
		DependencyIndexes: file_proto_types_model_scheduler_proto_depIdxs,
		EnumInfos:         file_proto_types_model_scheduler_proto_enumTypes,
		MessageInfos:      file_proto_types_model_scheduler_proto_msgTypes,
	}.Build()
	File_proto_types_model_scheduler_proto = out.File
	file_proto_types_model_scheduler_proto_rawDesc = nil
	file_proto_types_model_scheduler_proto_goTypes = nil
	file_proto_types_model_scheduler_proto_depIdxs = nil
}
