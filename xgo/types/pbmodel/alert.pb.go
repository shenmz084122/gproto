// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: proto/types/model/alert.proto

package pbmodel

import (
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

type AlertPolicy_Status int32

const (
	AlertPolicy_StatusUnset AlertPolicy_Status = 0
	AlertPolicy_Deleted     AlertPolicy_Status = 1
	AlertPolicy_Enabled     AlertPolicy_Status = 2
	AlertPolicy_Disabled    AlertPolicy_Status = 3
)

// Enum value maps for AlertPolicy_Status.
var (
	AlertPolicy_Status_name = map[int32]string{
		0: "StatusUnset",
		1: "Deleted",
		2: "Enabled",
		3: "Disabled",
	}
	AlertPolicy_Status_value = map[string]int32{
		"StatusUnset": 0,
		"Deleted":     1,
		"Enabled":     2,
		"Disabled":    3,
	}
)

func (x AlertPolicy_Status) Enum() *AlertPolicy_Status {
	p := new(AlertPolicy_Status)
	*p = x
	return p
}

func (x AlertPolicy_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AlertPolicy_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_types_model_alert_proto_enumTypes[0].Descriptor()
}

func (AlertPolicy_Status) Type() protoreflect.EnumType {
	return &file_proto_types_model_alert_proto_enumTypes[0]
}

func (x AlertPolicy_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AlertPolicy_Status.Descriptor instead.
func (AlertPolicy_Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_types_model_alert_proto_rawDescGZIP(), []int{0, 0}
}

type AlertPolicy_TriggerRule int32

const (
	AlertPolicy_TriggerRuleUnset   AlertPolicy_TriggerRule = 0
	AlertPolicy_TriggerRuleAnyItem AlertPolicy_TriggerRule = 1
)

// Enum value maps for AlertPolicy_TriggerRule.
var (
	AlertPolicy_TriggerRule_name = map[int32]string{
		0: "TriggerRuleUnset",
		1: "TriggerRuleAnyItem",
	}
	AlertPolicy_TriggerRule_value = map[string]int32{
		"TriggerRuleUnset":   0,
		"TriggerRuleAnyItem": 1,
	}
)

func (x AlertPolicy_TriggerRule) Enum() *AlertPolicy_TriggerRule {
	p := new(AlertPolicy_TriggerRule)
	*p = x
	return p
}

func (x AlertPolicy_TriggerRule) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AlertPolicy_TriggerRule) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_types_model_alert_proto_enumTypes[1].Descriptor()
}

func (AlertPolicy_TriggerRule) Type() protoreflect.EnumType {
	return &file_proto_types_model_alert_proto_enumTypes[1]
}

func (x AlertPolicy_TriggerRule) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AlertPolicy_TriggerRule.Descriptor instead.
func (AlertPolicy_TriggerRule) EnumDescriptor() ([]byte, []int) {
	return file_proto_types_model_alert_proto_rawDescGZIP(), []int{0, 1}
}

type AlertPolicy_TriggerAction int32

const (
	AlertPolicy_TriggerActionUnset            AlertPolicy_TriggerAction = 0
	AlertPolicy_TriggerActionSendNotification AlertPolicy_TriggerAction = 1
)

// Enum value maps for AlertPolicy_TriggerAction.
var (
	AlertPolicy_TriggerAction_name = map[int32]string{
		0: "TriggerActionUnset",
		1: "TriggerActionSendNotification",
	}
	AlertPolicy_TriggerAction_value = map[string]int32{
		"TriggerActionUnset":            0,
		"TriggerActionSendNotification": 1,
	}
)

func (x AlertPolicy_TriggerAction) Enum() *AlertPolicy_TriggerAction {
	p := new(AlertPolicy_TriggerAction)
	*p = x
	return p
}

func (x AlertPolicy_TriggerAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AlertPolicy_TriggerAction) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_types_model_alert_proto_enumTypes[2].Descriptor()
}

func (AlertPolicy_TriggerAction) Type() protoreflect.EnumType {
	return &file_proto_types_model_alert_proto_enumTypes[2]
}

func (x AlertPolicy_TriggerAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AlertPolicy_TriggerAction.Descriptor instead.
func (AlertPolicy_TriggerAction) EnumDescriptor() ([]byte, []int) {
	return file_proto_types_model_alert_proto_rawDescGZIP(), []int{0, 2}
}

type AlertPolicy_MonitorObject int32

const (
	AlertPolicy_ObjectUnset AlertPolicy_MonitorObject = 0
	AlertPolicy_StreamJob   AlertPolicy_MonitorObject = 1
	AlertPolicy_SyncJob     AlertPolicy_MonitorObject = 2
)

// Enum value maps for AlertPolicy_MonitorObject.
var (
	AlertPolicy_MonitorObject_name = map[int32]string{
		0: "ObjectUnset",
		1: "StreamJob",
		2: "SyncJob",
	}
	AlertPolicy_MonitorObject_value = map[string]int32{
		"ObjectUnset": 0,
		"StreamJob":   1,
		"SyncJob":     2,
	}
)

func (x AlertPolicy_MonitorObject) Enum() *AlertPolicy_MonitorObject {
	p := new(AlertPolicy_MonitorObject)
	*p = x
	return p
}

func (x AlertPolicy_MonitorObject) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AlertPolicy_MonitorObject) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_types_model_alert_proto_enumTypes[3].Descriptor()
}

func (AlertPolicy_MonitorObject) Type() protoreflect.EnumType {
	return &file_proto_types_model_alert_proto_enumTypes[3]
}

func (x AlertPolicy_MonitorObject) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AlertPolicy_MonitorObject.Descriptor instead.
func (AlertPolicy_MonitorObject) EnumDescriptor() ([]byte, []int) {
	return file_proto_types_model_alert_proto_rawDescGZIP(), []int{0, 3}
}

// The alert policy info
type AlertPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Workspace ID it belongs to.
	SpaceId string `protobuf:"bytes,1,opt,name=space_id,json=spaceId,proto3" json:"space_id" gorm:"column:space_id;"`
	// Alert ID, unique within a region.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id" gorm:"column:id;primarykey;"`
	// Name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name" gorm:"column:name;"`
	// Rule status, 1 => "Deleted", 2 => "enabled", 3 => "disabled".
	Status AlertPolicy_Status `protobuf:"varint,4,opt,name=status,proto3,enum=model.AlertPolicy_Status" json:"status" gorm:"column:status;"`
	// The monitor object , 1 => "StreamJob" 2 => "SyncJob'
	MonitorObject AlertPolicy_MonitorObject `protobuf:"varint,5,opt,name=monitor_object,json=monitorObject,proto3,enum=model.AlertPolicy_MonitorObject" json:"monitor_object" gorm:"column:monitor_object;"`
	// The alert item.
	MonitorItem *AlertPolicy_MonitorItem `protobuf:"bytes,6,opt,name=monitor_item,json=monitorItem,proto3" json:"monitor_item" gorm:"column:monitor_item;"`
	// The trigger rule, 1 => "TriggerRuleAnyItem",
	TriggerRule AlertPolicy_TriggerRule `protobuf:"varint,7,opt,name=trigger_rule,json=triggerRule,proto3,enum=model.AlertPolicy_TriggerRule" json:"trigger_rule" gorm:"column:trigger_rule;"`
	// The trigger action, 1 => "TriggerActionSendNotification",
	TriggerAction AlertPolicy_TriggerAction `protobuf:"varint,8,opt,name=trigger_action,json=triggerAction,proto3,enum=model.AlertPolicy_TriggerAction" json:"trigger_action" gorm:"column:trigger_action;"`
	// The notification list id of in IaaS. Multiple ids are separated by commas. e.g: "nl-mj9wzsa9,nl-pix7u2uo"
	NotificationIds string `protobuf:"bytes,9,opt,name=notification_ids,json=notificationIds,proto3" json:"notification_ids" gorm:"column:notification_ids;"`
	// AlertPolicy owner.
	CreatedBy string `protobuf:"bytes,10,opt,name=created_by,json=createdBy,proto3" json:"created_by" gorm:"column:created_by"`
	// Timestamp of create time.
	Created int64 `protobuf:"varint,11,opt,name=created,proto3" json:"created" gorm:"column:created;autoCreateTime;"`
	// Timestamp of update time.
	Updated int64 `protobuf:"varint,12,opt,name=updated,proto3" json:"updated" gorm:"column:updated;autoUpdateTime;"`
}

func (x *AlertPolicy) Reset() {
	*x = AlertPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_model_alert_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertPolicy) ProtoMessage() {}

func (x *AlertPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_model_alert_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertPolicy.ProtoReflect.Descriptor instead.
func (*AlertPolicy) Descriptor() ([]byte, []int) {
	return file_proto_types_model_alert_proto_rawDescGZIP(), []int{0}
}

func (x *AlertPolicy) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *AlertPolicy) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AlertPolicy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AlertPolicy) GetStatus() AlertPolicy_Status {
	if x != nil {
		return x.Status
	}
	return AlertPolicy_StatusUnset
}

func (x *AlertPolicy) GetMonitorObject() AlertPolicy_MonitorObject {
	if x != nil {
		return x.MonitorObject
	}
	return AlertPolicy_ObjectUnset
}

func (x *AlertPolicy) GetMonitorItem() *AlertPolicy_MonitorItem {
	if x != nil {
		return x.MonitorItem
	}
	return nil
}

func (x *AlertPolicy) GetTriggerRule() AlertPolicy_TriggerRule {
	if x != nil {
		return x.TriggerRule
	}
	return AlertPolicy_TriggerRuleUnset
}

func (x *AlertPolicy) GetTriggerAction() AlertPolicy_TriggerAction {
	if x != nil {
		return x.TriggerAction
	}
	return AlertPolicy_TriggerActionUnset
}

func (x *AlertPolicy) GetNotificationIds() string {
	if x != nil {
		return x.NotificationIds
	}
	return ""
}

func (x *AlertPolicy) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *AlertPolicy) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *AlertPolicy) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

// AlertPolicyBinding for record the relationships between alert policy and monitor object.
type AlertPolicyBinding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of alert policy.
	AlertId string `protobuf:"bytes,1,opt,name=alert_id,json=alertId,proto3" json:"alert_id" gorm:"column:alert_id;"`
	// The id of monitor object. StreamJob(stj-xxxxxxxxxxxx) or SyncJob(syj-xxxxxxxxxxxx).
	JobId string `protobuf:"bytes,2,opt,name=job_id,json=jobId,proto3" json:"job_id" gorm:"column:job_id;"`
}

func (x *AlertPolicyBinding) Reset() {
	*x = AlertPolicyBinding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_model_alert_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertPolicyBinding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertPolicyBinding) ProtoMessage() {}

func (x *AlertPolicyBinding) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_model_alert_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertPolicyBinding.ProtoReflect.Descriptor instead.
func (*AlertPolicyBinding) Descriptor() ([]byte, []int) {
	return file_proto_types_model_alert_proto_rawDescGZIP(), []int{1}
}

func (x *AlertPolicyBinding) GetAlertId() string {
	if x != nil {
		return x.AlertId
	}
	return ""
}

func (x *AlertPolicyBinding) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type AlertPolicy_MonitorStreamJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether send alerts when job instance run failed.
	InstanceRunFailed bool `protobuf:"varint,1,opt,name=instance_run_failed,json=instanceRunFailed,proto3" json:"instance_run_failed"`
	// Whether send alerts when job instance run timeout.
	InstanceRunTimeout bool `protobuf:"varint,2,opt,name=instance_run_timeout,json=instanceRunTimeout,proto3" json:"instance_run_timeout"`
	// Timeout for the job instance to run. Is valid only when `instance_run_timeout` is true.
	InstanceTimeout int32 `protobuf:"varint,3,opt,name=instance_timeout,json=instanceTimeout,proto3" json:"instance_timeout"`
}

func (x *AlertPolicy_MonitorStreamJob) Reset() {
	*x = AlertPolicy_MonitorStreamJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_model_alert_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertPolicy_MonitorStreamJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertPolicy_MonitorStreamJob) ProtoMessage() {}

func (x *AlertPolicy_MonitorStreamJob) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_model_alert_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertPolicy_MonitorStreamJob.ProtoReflect.Descriptor instead.
func (*AlertPolicy_MonitorStreamJob) Descriptor() ([]byte, []int) {
	return file_proto_types_model_alert_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AlertPolicy_MonitorStreamJob) GetInstanceRunFailed() bool {
	if x != nil {
		return x.InstanceRunFailed
	}
	return false
}

func (x *AlertPolicy_MonitorStreamJob) GetInstanceRunTimeout() bool {
	if x != nil {
		return x.InstanceRunTimeout
	}
	return false
}

func (x *AlertPolicy_MonitorStreamJob) GetInstanceTimeout() int32 {
	if x != nil {
		return x.InstanceTimeout
	}
	return 0
}

type AlertPolicy_MonitorSyncJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether send alerts when job instance run failed.
	InstanceRunFailed bool `protobuf:"varint,1,opt,name=instance_run_failed,json=instanceRunFailed,proto3" json:"instance_run_failed"`
	// Whether send alerts when job instance run timeout.
	InstanceRunTimeout bool `protobuf:"varint,2,opt,name=instance_run_timeout,json=instanceRunTimeout,proto3" json:"instance_run_timeout"`
	// Timeout for the job instance to run. Is valid only when `instance_run_timeout` is true.
	InstanceTimeout int32 `protobuf:"varint,3,opt,name=instance_timeout,json=instanceTimeout,proto3" json:"instance_timeout"`
}

func (x *AlertPolicy_MonitorSyncJob) Reset() {
	*x = AlertPolicy_MonitorSyncJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_model_alert_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertPolicy_MonitorSyncJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertPolicy_MonitorSyncJob) ProtoMessage() {}

func (x *AlertPolicy_MonitorSyncJob) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_model_alert_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertPolicy_MonitorSyncJob.ProtoReflect.Descriptor instead.
func (*AlertPolicy_MonitorSyncJob) Descriptor() ([]byte, []int) {
	return file_proto_types_model_alert_proto_rawDescGZIP(), []int{0, 1}
}

func (x *AlertPolicy_MonitorSyncJob) GetInstanceRunFailed() bool {
	if x != nil {
		return x.InstanceRunFailed
	}
	return false
}

func (x *AlertPolicy_MonitorSyncJob) GetInstanceRunTimeout() bool {
	if x != nil {
		return x.InstanceRunTimeout
	}
	return false
}

func (x *AlertPolicy_MonitorSyncJob) GetInstanceTimeout() int32 {
	if x != nil {
		return x.InstanceTimeout
	}
	return 0
}

type AlertPolicy_MonitorItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamJob *AlertPolicy_MonitorStreamJob `protobuf:"bytes,1,opt,name=stream_job,json=streamJob,proto3" json:"stream_job"`
	SyncJob   *AlertPolicy_MonitorSyncJob   `protobuf:"bytes,2,opt,name=sync_job,json=syncJob,proto3" json:"sync_job"`
}

func (x *AlertPolicy_MonitorItem) Reset() {
	*x = AlertPolicy_MonitorItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_model_alert_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertPolicy_MonitorItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertPolicy_MonitorItem) ProtoMessage() {}

func (x *AlertPolicy_MonitorItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_model_alert_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertPolicy_MonitorItem.ProtoReflect.Descriptor instead.
func (*AlertPolicy_MonitorItem) Descriptor() ([]byte, []int) {
	return file_proto_types_model_alert_proto_rawDescGZIP(), []int{0, 2}
}

func (x *AlertPolicy_MonitorItem) GetStreamJob() *AlertPolicy_MonitorStreamJob {
	if x != nil {
		return x.StreamJob
	}
	return nil
}

func (x *AlertPolicy_MonitorItem) GetSyncJob() *AlertPolicy_MonitorSyncJob {
	if x != nil {
		return x.SyncJob
	}
	return nil
}

var File_proto_types_model_alert_proto protoreflect.FileDescriptor

var file_proto_types_model_alert_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x79, 0x75, 0x33, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x75, 0x33, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x73, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x0b, 0x0a,
	0x0b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2e, 0x0a, 0x08,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13,
	0xe2, 0xdf, 0x1f, 0x0f, 0x12, 0x0d, 0xc2, 0x01, 0x0a, 0xf0, 0x01, 0x14, 0xca, 0x02, 0x04, 0x77,
	0x6b, 0x73, 0x2d, 0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0xe2, 0xdf, 0x1f, 0x0f, 0x12, 0x0d,
	0xc2, 0x01, 0x0a, 0xf0, 0x01, 0x14, 0xca, 0x02, 0x04, 0x61, 0x6c, 0x74, 0x2d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x10, 0xe2, 0xdf, 0x1f, 0x0c, 0x12, 0x0a, 0xc2, 0x01, 0x07, 0x90, 0x02, 0x02, 0x98, 0x02, 0x80,
	0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x42, 0x0d, 0xe2, 0xdf, 0x1f, 0x09, 0x12, 0x07, 0xda, 0x01, 0x04, 0x30, 0x00, 0x58,
	0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x56, 0x0a, 0x0e, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x20, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x42, 0x0d, 0xe2, 0xdf, 0x1f, 0x09, 0x12, 0x07, 0xda, 0x01, 0x04, 0x30, 0x00,
	0x58, 0x01, 0x52, 0x0d, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x4e, 0x0a, 0x0c, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x4d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x0b, 0xe2, 0xdf, 0x1f, 0x07, 0x12, 0x05, 0xe2,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x50, 0x0a, 0x0c, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x54, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x42, 0x0d, 0xe2, 0xdf, 0x1f, 0x09, 0x12, 0x07, 0xda,
	0x01, 0x04, 0x30, 0x00, 0x58, 0x01, 0x52, 0x0b, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52,
	0x75, 0x6c, 0x65, 0x12, 0x56, 0x0a, 0x0e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0xe2,
	0xdf, 0x1f, 0x09, 0x12, 0x07, 0xda, 0x01, 0x04, 0x30, 0x00, 0x58, 0x01, 0x52, 0x0d, 0x74, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x10, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xe2, 0xdf, 0x1f, 0x07, 0x12, 0x05, 0xc2, 0x01, 0x02,
	0x22, 0x00, 0x52, 0x0f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x73, 0x12, 0x2e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62,
	0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xe2, 0xdf, 0x1f, 0x0b, 0x12, 0x09, 0xc2,
	0x01, 0x06, 0x80, 0x02, 0x00, 0x88, 0x02, 0x41, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x25, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x0b, 0xe2, 0xdf, 0x1f, 0x07, 0x12, 0x05, 0xb2, 0x01, 0x02, 0x30,
	0x00, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x07, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0b, 0xe2, 0xdf, 0x1f,
	0x07, 0x12, 0x05, 0xb2, 0x01, 0x02, 0x30, 0x00, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x1a, 0x9f, 0x01, 0x0a, 0x10, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x4a, 0x6f, 0x62, 0x12, 0x2e, 0x0a, 0x13, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x11, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x75, 0x6e,
	0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x75,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x1a, 0x9d, 0x01, 0x0a, 0x0e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x53,
	0x79, 0x6e, 0x63, 0x4a, 0x6f, 0x62, 0x12, 0x2e, 0x0a, 0x13, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x11, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x75, 0x6e,
	0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x75,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x1a, 0x97, 0x01, 0x0a, 0x0b, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x42, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6a, 0x6f,
	0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x4d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4a, 0x6f, 0x62, 0x52, 0x09, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x4a, 0x6f, 0x62, 0x12, 0x3c, 0x0a, 0x08, 0x73, 0x79, 0x6e, 0x63, 0x5f,
	0x6a, 0x6f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x4d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x53, 0x79, 0x6e, 0x63, 0x4a, 0x6f, 0x62, 0x52, 0x07, 0x73, 0x79,
	0x6e, 0x63, 0x4a, 0x6f, 0x62, 0x3a, 0x06, 0xca, 0xb2, 0x04, 0x02, 0x0a, 0x00, 0x22, 0x41, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x55, 0x6e, 0x73, 0x65, 0x74, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x10, 0x03,
	0x22, 0x3b, 0x0a, 0x0b, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x12,
	0x14, 0x0a, 0x10, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x55, 0x6e,
	0x73, 0x65, 0x74, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x52, 0x75, 0x6c, 0x65, 0x41, 0x6e, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x10, 0x01, 0x22, 0x4a, 0x0a,
	0x0d, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16,
	0x0a, 0x12, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55,
	0x6e, 0x73, 0x65, 0x74, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x01, 0x22, 0x3c, 0x0a, 0x0d, 0x4d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x55, 0x6e, 0x73, 0x65, 0x74, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x4a, 0x6f, 0x62, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x79,
	0x6e, 0x63, 0x4a, 0x6f, 0x62, 0x10, 0x02, 0x22, 0x69, 0x0a, 0x12, 0x41, 0x6c, 0x65, 0x72, 0x74,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x2e, 0x0a,
	0x08, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x13, 0xe2, 0xdf, 0x1f, 0x0f, 0x12, 0x0d, 0xc2, 0x01, 0x0a, 0xf0, 0x01, 0x14, 0xca, 0x02, 0x04,
	0x61, 0x6c, 0x74, 0x2d, 0x52, 0x07, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a,
	0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe2,
	0xdf, 0x1f, 0x08, 0x12, 0x06, 0xc2, 0x01, 0x03, 0xf0, 0x01, 0x14, 0x52, 0x05, 0x6a, 0x6f, 0x62,
	0x49, 0x64, 0x42, 0x67, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6f, 0x6d,
	0x6e, 0x69, 0x73, 0x2e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x62, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x0c, 0x50, 0x42, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x50, 0x00, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x57, 0x6f, 0x72, 0x6b, 0x62, 0x65, 0x6e,
	0x63, 0x68, 0x2f, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x78, 0x67, 0x6f, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_types_model_alert_proto_rawDescOnce sync.Once
	file_proto_types_model_alert_proto_rawDescData = file_proto_types_model_alert_proto_rawDesc
)

func file_proto_types_model_alert_proto_rawDescGZIP() []byte {
	file_proto_types_model_alert_proto_rawDescOnce.Do(func() {
		file_proto_types_model_alert_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_model_alert_proto_rawDescData)
	})
	return file_proto_types_model_alert_proto_rawDescData
}

var file_proto_types_model_alert_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_proto_types_model_alert_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_types_model_alert_proto_goTypes = []interface{}{
	(AlertPolicy_Status)(0),              // 0: model.AlertPolicy.Status
	(AlertPolicy_TriggerRule)(0),         // 1: model.AlertPolicy.TriggerRule
	(AlertPolicy_TriggerAction)(0),       // 2: model.AlertPolicy.TriggerAction
	(AlertPolicy_MonitorObject)(0),       // 3: model.AlertPolicy.MonitorObject
	(*AlertPolicy)(nil),                  // 4: model.AlertPolicy
	(*AlertPolicyBinding)(nil),           // 5: model.AlertPolicyBinding
	(*AlertPolicy_MonitorStreamJob)(nil), // 6: model.AlertPolicy.MonitorStreamJob
	(*AlertPolicy_MonitorSyncJob)(nil),   // 7: model.AlertPolicy.MonitorSyncJob
	(*AlertPolicy_MonitorItem)(nil),      // 8: model.AlertPolicy.MonitorItem
}
var file_proto_types_model_alert_proto_depIdxs = []int32{
	0, // 0: model.AlertPolicy.status:type_name -> model.AlertPolicy.Status
	3, // 1: model.AlertPolicy.monitor_object:type_name -> model.AlertPolicy.MonitorObject
	8, // 2: model.AlertPolicy.monitor_item:type_name -> model.AlertPolicy.MonitorItem
	1, // 3: model.AlertPolicy.trigger_rule:type_name -> model.AlertPolicy.TriggerRule
	2, // 4: model.AlertPolicy.trigger_action:type_name -> model.AlertPolicy.TriggerAction
	6, // 5: model.AlertPolicy.MonitorItem.stream_job:type_name -> model.AlertPolicy.MonitorStreamJob
	7, // 6: model.AlertPolicy.MonitorItem.sync_job:type_name -> model.AlertPolicy.MonitorSyncJob
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_types_model_alert_proto_init() }
func file_proto_types_model_alert_proto_init() {
	if File_proto_types_model_alert_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_types_model_alert_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertPolicy); i {
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
		file_proto_types_model_alert_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertPolicyBinding); i {
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
		file_proto_types_model_alert_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertPolicy_MonitorStreamJob); i {
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
		file_proto_types_model_alert_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertPolicy_MonitorSyncJob); i {
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
		file_proto_types_model_alert_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertPolicy_MonitorItem); i {
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
			RawDescriptor: file_proto_types_model_alert_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_model_alert_proto_goTypes,
		DependencyIndexes: file_proto_types_model_alert_proto_depIdxs,
		EnumInfos:         file_proto_types_model_alert_proto_enumTypes,
		MessageInfos:      file_proto_types_model_alert_proto_msgTypes,
	}.Build()
	File_proto_types_model_alert_proto = out.File
	file_proto_types_model_alert_proto_rawDesc = nil
	file_proto_types_model_alert_proto_goTypes = nil
	file_proto_types_model_alert_proto_depIdxs = nil
}
