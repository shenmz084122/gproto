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

type QueueMessage_Action int32

const (
	QueueMessage_ActionUnset QueueMessage_Action = 0
	QueueMessage_Create      QueueMessage_Action = 1
	QueueMessage_Init        QueueMessage_Action = 5
	QueueMessage_Submit      QueueMessage_Action = 2
	QueueMessage_Check       QueueMessage_Action = 4
	QueueMessage_Retry       QueueMessage_Action = 3
)

// Enum value maps for QueueMessage_Action.
var (
	QueueMessage_Action_name = map[int32]string{
		0: "ActionUnset",
		1: "Create",
		5: "Init",
		2: "Submit",
		4: "Check",
		3: "Retry",
	}
	QueueMessage_Action_value = map[string]int32{
		"ActionUnset": 0,
		"Create":      1,
		"Init":        5,
		"Submit":      2,
		"Check":       4,
		"Retry":       3,
	}
)

func (x QueueMessage_Action) Enum() *QueueMessage_Action {
	p := new(QueueMessage_Action)
	*p = x
	return p
}

func (x QueueMessage_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QueueMessage_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_types_model_scheduler_proto_enumTypes[0].Descriptor()
}

func (QueueMessage_Action) Type() protoreflect.EnumType {
	return &file_proto_types_model_scheduler_proto_enumTypes[0]
}

func (x QueueMessage_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QueueMessage_Action.Descriptor instead.
func (QueueMessage_Action) EnumDescriptor() ([]byte, []int) {
	return file_proto_types_model_scheduler_proto_rawDescGZIP(), []int{0, 0}
}

// The message in inst queue..
type QueueMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The property of stream job.
	Property *StreamJobProperty `protobuf:"bytes,1,opt,name=property,proto3" json:"property,omitempty"`
	// The number of times the instance task was executed.
	Retries int32 `protobuf:"varint,2,opt,name=retries,proto3" json:"retries,omitempty"`
	// The start execution time of the task instance.
	Started int64 `protobuf:"varint,3,opt,name=started,proto3" json:"started,omitempty"`
	// zeppelin notebook id
	NoteId string `protobuf:"bytes,4,opt,name=note_id,json=noteId,proto3" json:"note_id"`
	// zeppelin paragraphID
	ParagraphId string `protobuf:"bytes,5,opt,name=paragraph_id,json=paragraphId,proto3" json:"paragraph_id"`
	FlinkId string `protobuf:"bytes,6,opt,name=flink_id,json=flinkId,proto3" json:"flink_id"`
}

func (x *QueueMessage) Reset() {
	*x = QueueMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_model_scheduler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueMessage) ProtoMessage() {}

func (x *QueueMessage) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use QueueMessage.ProtoReflect.Descriptor instead.
func (*QueueMessage) Descriptor() ([]byte, []int) {
	return file_proto_types_model_scheduler_proto_rawDescGZIP(), []int{0}
}

func (x *QueueMessage) GetProperty() *StreamJobProperty {
	if x != nil {
		return x.Property
	}
	return nil
}

func (x *QueueMessage) GetRetries() int32 {
	if x != nil {
		return x.Retries
	}
	return 0
}

func (x *QueueMessage) GetStarted() int64 {
	if x != nil {
		return x.Started
	}
	return 0
}

func (x *QueueMessage) GetNoteId() string {
	if x != nil {
		return x.NoteId
	}
	return ""
}

func (x *QueueMessage) GetParagraphId() string {
	if x != nil {
		return x.ParagraphId
	}
	return ""
}

func (x *QueueMessage) GetFlinkId() string {
	if x != nil {
		return x.FlinkId
	}
	return ""
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
	0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x02, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x75, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x4a, 0x6f, 0x62, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x42, 0x0b, 0xe2, 0xdf, 0x1f, 0x07, 0x12, 0x05, 0xe2, 0x01, 0x02, 0x10, 0x01, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x04, 0xe2, 0xdf, 0x1f, 0x00, 0x52, 0x07,
	0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x04, 0xe2, 0xdf, 0x1f, 0x00, 0x52, 0x07,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x07, 0x6e, 0x6f, 0x74, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0xdf, 0x1f, 0x00, 0x52, 0x06,
	0x6e, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x61, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0xdf,
	0x1f, 0x00, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x49, 0x64, 0x12,
	0x24, 0x0a, 0x08, 0x66, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x09, 0xe2, 0xdf, 0x1f, 0x05, 0x12, 0x03, 0xc2, 0x01, 0x00, 0x52, 0x07, 0x66, 0x6c,
	0x69, 0x6e, 0x6b, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0f, 0x0a, 0x0b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x73, 0x65, 0x74, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x49, 0x6e, 0x69, 0x74, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x10, 0x04, 0x12, 0x09, 0x0a,
	0x05, 0x52, 0x65, 0x74, 0x72, 0x79, 0x10, 0x03, 0x42, 0x6b, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x6f, 0x6d, 0x6e, 0x69, 0x73, 0x2e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x10,
	0x50, 0x42, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x50, 0x00, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44,
	0x61, 0x74, 0x61, 0x57, 0x6f, 0x72, 0x6b, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2f, 0x67, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x78, 0x67, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x62,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_proto_types_model_scheduler_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_types_model_scheduler_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_types_model_scheduler_proto_goTypes = []interface{}{
	(QueueMessage_Action)(0),  // 0: model.QueueMessage.Action
	(*QueueMessage)(nil),      // 1: model.QueueMessage
	(*StreamJobProperty)(nil), // 2: model.StreamJobProperty
}
var file_proto_types_model_scheduler_proto_depIdxs = []int32{
	2, // 0: model.QueueMessage.property:type_name -> model.StreamJobProperty
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
			switch v := v.(*QueueMessage); i {
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
			NumEnums:      1,
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
