// Code generated by protoc-gen-govalidator. DO NOT EDIT.
// versions:
// 		protoc-gen-govalidator 0.0.1
// source: proto/types/request/sync_job_manage.proto

package pbrequest

import (
	pbmodel "github.com/DataWorkbench/gproto/xgo/types/pbmodel"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbdefaults"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbgosql"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
	protovalidator "github.com/yu31/protoc-plugin/xgo/pkg/protovalidator"
	strconv "strconv"
	strings "strings"
	utf8 "unicode/utf8"
)

func (this *ListSyncJobs) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("ListSyncJobs", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("ListSyncJobs", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

func (this *ListSyncJobs) _xxx_xxx_Validator_Validate_limit() error {
	if !(this.Limit > 0) {
		return protovalidator.FieldError1("ListSyncJobs", "the value of field 'limit' must be greater than '0'", protovalidator.Int32ToString(this.Limit))
	}
	if !(this.Limit <= 100) {
		return protovalidator.FieldError1("ListSyncJobs", "the value of field 'limit' must be less than or equal to '100'", protovalidator.Int32ToString(this.Limit))
	}
	return nil
}

func (this *ListSyncJobs) _xxx_xxx_Validator_Validate_offset() error {
	if !(this.Offset >= 0) {
		return protovalidator.FieldError1("ListSyncJobs", "the value of field 'offset' must be greater than or equal to '0'", protovalidator.Int32ToString(this.Offset))
	}
	return nil
}

var _xxx_xxx_Validator_ListSyncJobs_In_SortBy = map[string]bool{"": true, "id": true, "name": true, "created": true, "updated": true}

func (this *ListSyncJobs) _xxx_xxx_Validator_Validate_sort_by() error {
	if !(_xxx_xxx_Validator_ListSyncJobs_In_SortBy[this.SortBy]) {
		return protovalidator.FieldError1("ListSyncJobs", "the value of field 'sort_by' must be one of '[ id name created updated]'", this.SortBy)
	}
	return nil
}

// Set default value for message request.ListSyncJobs
func (this *ListSyncJobs) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_space_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_limit(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_offset(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_sort_by(); err != nil {
		return err
	}
	return nil
}

func (this *CreateSyncJob) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("CreateSyncJob", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("CreateSyncJob", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

func (this *CreateSyncJob) _xxx_xxx_Validator_Validate_created_by() error {
	if !(len(this.CreatedBy) < 65) {
		return protovalidator.FieldError1("CreateSyncJob", "the byte length of field 'created_by' must be less than '65'", protovalidator.StringByteLenToString(this.CreatedBy))
	}
	return nil
}

func (this *CreateSyncJob) _xxx_xxx_Validator_CheckIf_pid() bool {
	if !(this.Pid != "") {
		return false
	}
	return true
}

func (this *CreateSyncJob) _xxx_xxx_Validator_Validate_pid() error {
	if !this._xxx_xxx_Validator_CheckIf_pid() {
		return nil
	}
	if !(len(this.Pid) == 20) {
		return protovalidator.FieldError1("CreateSyncJob", "the byte length of field 'pid' must be equal to '20'", protovalidator.StringByteLenToString(this.Pid))
	}
	if !(strings.HasPrefix(this.Pid, "syj-")) {
		return protovalidator.FieldError1("CreateSyncJob", "the value of field 'pid' must start with string 'syj-'", this.Pid)
	}
	return nil
}

func (this *CreateSyncJob) _xxx_xxx_Validator_Validate_name() error {
	if !(utf8.RuneCountInString(this.Name) >= 2) {
		return protovalidator.FieldError1("CreateSyncJob", "the character length of field 'name' must be greater than or equal to '2'", protovalidator.StringCharsetLenToString(this.Name))
	}
	if !(utf8.RuneCountInString(this.Name) <= 128) {
		return protovalidator.FieldError1("CreateSyncJob", "the character length of field 'name' must be less than or equal to '128'", protovalidator.StringCharsetLenToString(this.Name))
	}
	return nil
}

func (this *CreateSyncJob) _xxx_xxx_Validator_Validate_desc() error {
	if !(utf8.RuneCountInString(this.Desc) <= 1020) {
		return protovalidator.FieldError1("CreateSyncJob", "the character length of field 'desc' must be less than or equal to '1020'", protovalidator.StringCharsetLenToString(this.Desc))
	}
	return nil
}

func (this *CreateSyncJob) _xxx_xxx_Validator_CheckIf_type() bool {
	if !(this.IsDirectory == false) {
		return false
	}
	return true
}

var _xxx_xxx_Validator_CreateSyncJob_InEnums_Type = map[pbmodel.SyncJob_Type]bool{0: true, 1: true, 2: true, 3: true}

func (this *CreateSyncJob) _xxx_xxx_Validator_Validate_type() error {
	if !this._xxx_xxx_Validator_CheckIf_type() {
		return nil
	}
	if !(this.Type >= 0) {
		return protovalidator.FieldError1("CreateSyncJob", "the value of field 'type' must be greater than or equal to '0'", protovalidator.Int32ToString(int32(this.Type)))
	}
	if !(_xxx_xxx_Validator_CreateSyncJob_InEnums_Type[this.Type]) {
		return protovalidator.FieldError1("CreateSyncJob", "the value of field 'type' must in enums of '[0 1 2 3]'", protovalidator.Int32ToString(int32(this.Type)))
	}
	return nil
}

func (this *CreateSyncJob) _xxx_xxx_Validator_Validate_space_owner() error {
	if !(this.SpaceOwner != "") {
		return protovalidator.FieldError1("CreateSyncJob", "the value of field 'space_owner' must be not equal to ''", this.SpaceOwner)
	}
	return nil
}

func (this *CreateSyncJob) _xxx_xxx_Validator_CheckIf_source_type() bool {
	if !(this.IsDirectory == false) {
		return false
	}
	return true
}

var _xxx_xxx_Validator_CreateSyncJob_InEnums_SourceType = map[pbmodel.DataSource_Type]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true}

func (this *CreateSyncJob) _xxx_xxx_Validator_Validate_source_type() error {
	if !this._xxx_xxx_Validator_CheckIf_source_type() {
		return nil
	}
	if !(this.SourceType > 0) {
		return protovalidator.FieldError1("CreateSyncJob", "the value of field 'source_type' must be greater than '0'", protovalidator.Int32ToString(int32(this.SourceType)))
	}
	if !(_xxx_xxx_Validator_CreateSyncJob_InEnums_SourceType[this.SourceType]) {
		return protovalidator.FieldError1("CreateSyncJob", "the value of field 'source_type' must in enums of '[0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16]'", protovalidator.Int32ToString(int32(this.SourceType)))
	}
	return nil
}

func (this *CreateSyncJob) _xxx_xxx_Validator_CheckIf_target_type() bool {
	if !(this.IsDirectory == false) {
		return false
	}
	return true
}

var _xxx_xxx_Validator_CreateSyncJob_InEnums_TargetType = map[pbmodel.DataSource_Type]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true, 11: true, 12: true, 13: true, 14: true, 15: true, 16: true}

func (this *CreateSyncJob) _xxx_xxx_Validator_Validate_target_type() error {
	if !this._xxx_xxx_Validator_CheckIf_target_type() {
		return nil
	}
	if !(this.TargetType > 0) {
		return protovalidator.FieldError1("CreateSyncJob", "the value of field 'target_type' must be greater than '0'", protovalidator.Int32ToString(int32(this.TargetType)))
	}
	if !(_xxx_xxx_Validator_CreateSyncJob_InEnums_TargetType[this.TargetType]) {
		return protovalidator.FieldError1("CreateSyncJob", "the value of field 'target_type' must in enums of '[0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16]'", protovalidator.Int32ToString(int32(this.TargetType)))
	}
	return nil
}

// Set default value for message request.CreateSyncJob
func (this *CreateSyncJob) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_space_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_created_by(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_pid(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_name(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_desc(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_type(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_space_owner(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_source_type(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_target_type(); err != nil {
		return err
	}
	return nil
}

func (this *DeleteSyncJobs) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("DeleteSyncJobs", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("DeleteSyncJobs", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

func (this *DeleteSyncJobs) _xxx_xxx_Validator_Validate_job_ids() error {
	if !(len(this.JobIds) > 0) {
		return protovalidator.FieldError1("DeleteSyncJobs", "the length of field 'job_ids' must be greater than '0'", strconv.Itoa(len(this.JobIds)))
	}
	if !(len(this.JobIds) <= 100) {
		return protovalidator.FieldError1("DeleteSyncJobs", "the length of field 'job_ids' must be less than or equal to '100'", strconv.Itoa(len(this.JobIds)))
	}
	for _, item := range this.JobIds {
		_ = item // To avoid unused panics.
		if !(strings.HasPrefix(item, "syj-")) {
			return protovalidator.FieldError1("DeleteSyncJobs", "the value of array item where in field 'job_ids' must start with string 'syj-'", item)
		}
	}
	return nil
}

// Set default value for message request.DeleteSyncJobs
func (this *DeleteSyncJobs) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_space_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_job_ids(); err != nil {
		return err
	}
	return nil
}

func (this *MoveSyncJobs) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("MoveSyncJobs", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("MoveSyncJobs", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

func (this *MoveSyncJobs) _xxx_xxx_Validator_Validate_job_ids() error {
	if !(len(this.JobIds) > 0) {
		return protovalidator.FieldError1("MoveSyncJobs", "the length of field 'job_ids' must be greater than '0'", strconv.Itoa(len(this.JobIds)))
	}
	if !(len(this.JobIds) <= 100) {
		return protovalidator.FieldError1("MoveSyncJobs", "the length of field 'job_ids' must be less than or equal to '100'", strconv.Itoa(len(this.JobIds)))
	}
	for _, item := range this.JobIds {
		_ = item // To avoid unused panics.
		if !(strings.HasPrefix(item, "syj-")) {
			return protovalidator.FieldError1("MoveSyncJobs", "the value of array item where in field 'job_ids' must start with string 'syj-'", item)
		}
	}
	return nil
}

func (this *MoveSyncJobs) _xxx_xxx_Validator_CheckIf_target() bool {
	if !(this.Target != "") {
		return false
	}
	return true
}

func (this *MoveSyncJobs) _xxx_xxx_Validator_Validate_target() error {
	if !this._xxx_xxx_Validator_CheckIf_target() {
		return nil
	}
	if !(len(this.Target) == 20) {
		return protovalidator.FieldError1("MoveSyncJobs", "the byte length of field 'target' must be equal to '20'", protovalidator.StringByteLenToString(this.Target))
	}
	if !(strings.HasPrefix(this.Target, "syj-")) {
		return protovalidator.FieldError1("MoveSyncJobs", "the value of field 'target' must start with string 'syj-'", this.Target)
	}
	return nil
}

// Set default value for message request.MoveSyncJobs
func (this *MoveSyncJobs) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_space_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_job_ids(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_target(); err != nil {
		return err
	}
	return nil
}

func (this *UpdateSyncJob) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("UpdateSyncJob", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("UpdateSyncJob", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

func (this *UpdateSyncJob) _xxx_xxx_Validator_Validate_job_id() error {
	if !(len(this.JobId) == 20) {
		return protovalidator.FieldError1("UpdateSyncJob", "the byte length of field 'job_id' must be equal to '20'", protovalidator.StringByteLenToString(this.JobId))
	}
	if !(strings.HasPrefix(this.JobId, "syj-")) {
		return protovalidator.FieldError1("UpdateSyncJob", "the value of field 'job_id' must start with string 'syj-'", this.JobId)
	}
	return nil
}

func (this *UpdateSyncJob) _xxx_xxx_Validator_Validate_name() error {
	if !(utf8.RuneCountInString(this.Name) >= 2) {
		return protovalidator.FieldError1("UpdateSyncJob", "the character length of field 'name' must be greater than or equal to '2'", protovalidator.StringCharsetLenToString(this.Name))
	}
	if !(utf8.RuneCountInString(this.Name) <= 128) {
		return protovalidator.FieldError1("UpdateSyncJob", "the character length of field 'name' must be less than or equal to '128'", protovalidator.StringCharsetLenToString(this.Name))
	}
	return nil
}

func (this *UpdateSyncJob) _xxx_xxx_Validator_Validate_desc() error {
	if !(utf8.RuneCountInString(this.Desc) <= 1024) {
		return protovalidator.FieldError1("UpdateSyncJob", "the character length of field 'desc' must be less than or equal to '1024'", protovalidator.StringCharsetLenToString(this.Desc))
	}
	return nil
}

// Set default value for message request.UpdateSyncJob
func (this *UpdateSyncJob) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_space_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_job_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_name(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_desc(); err != nil {
		return err
	}
	return nil
}

func (this *DescribeSyncJob) _xxx_xxx_Validator_Validate_job_id() error {
	if !(len(this.JobId) == 20) {
		return protovalidator.FieldError1("DescribeSyncJob", "the byte length of field 'job_id' must be equal to '20'", protovalidator.StringByteLenToString(this.JobId))
	}
	if !(strings.HasPrefix(this.JobId, "syj-")) {
		return protovalidator.FieldError1("DescribeSyncJob", "the value of field 'job_id' must start with string 'syj-'", this.JobId)
	}
	return nil
}

// Set default value for message request.DescribeSyncJob
func (this *DescribeSyncJob) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_job_id(); err != nil {
		return err
	}
	return nil
}

func (this *SetSyncJobSchedule) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("SetSyncJobSchedule", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("SetSyncJobSchedule", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

func (this *SetSyncJobSchedule) _xxx_xxx_Validator_Validate_job_id() error {
	if !(len(this.JobId) == 20) {
		return protovalidator.FieldError1("SetSyncJobSchedule", "the byte length of field 'job_id' must be equal to '20'", protovalidator.StringByteLenToString(this.JobId))
	}
	if !(strings.HasPrefix(this.JobId, "syj-")) {
		return protovalidator.FieldError1("SetSyncJobSchedule", "the value of field 'job_id' must start with string 'syj-'", this.JobId)
	}
	return nil
}

func (this *SetSyncJobSchedule) _xxx_xxx_Validator_Validate_schedule() error {
	if !(this.Schedule != nil) {
		return protovalidator.FieldError2("SetSyncJobSchedule", "the value of field 'schedule' cannot be null")
	}
	if dt, ok := interface{}(this.Schedule).(interface{ Validate() error }); ok {
		if err := dt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Set default value for message request.SetSyncJobSchedule
func (this *SetSyncJobSchedule) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_space_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_job_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_schedule(); err != nil {
		return err
	}
	return nil
}

func (this *SetSyncJobConf) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("SetSyncJobConf", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("SetSyncJobConf", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

func (this *SetSyncJobConf) _xxx_xxx_Validator_Validate_job_id() error {
	if !(len(this.JobId) == 20) {
		return protovalidator.FieldError1("SetSyncJobConf", "the byte length of field 'job_id' must be equal to '20'", protovalidator.StringByteLenToString(this.JobId))
	}
	if !(strings.HasPrefix(this.JobId, "syj-")) {
		return protovalidator.FieldError1("SetSyncJobConf", "the value of field 'job_id' must start with string 'syj-'", this.JobId)
	}
	return nil
}

func (this *SetSyncJobConf) _xxx_xxx_Validator_Validate_conf() error {
	if !(this.Conf != nil) {
		return protovalidator.FieldError2("SetSyncJobConf", "the value of field 'conf' cannot be null")
	}
	if dt, ok := interface{}(this.Conf).(interface{ Validate() error }); ok {
		if err := dt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Set default value for message request.SetSyncJobConf
func (this *SetSyncJobConf) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_space_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_job_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_conf(); err != nil {
		return err
	}
	return nil
}

func (this *GetSyncJobSchedule) _xxx_xxx_Validator_Validate_job_id() error {
	if !(len(this.JobId) == 20) {
		return protovalidator.FieldError1("GetSyncJobSchedule", "the byte length of field 'job_id' must be equal to '20'", protovalidator.StringByteLenToString(this.JobId))
	}
	if !(strings.HasPrefix(this.JobId, "syj-")) {
		return protovalidator.FieldError1("GetSyncJobSchedule", "the value of field 'job_id' must start with string 'syj-'", this.JobId)
	}
	return nil
}

// Set default value for message request.GetSyncJobSchedule
func (this *GetSyncJobSchedule) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_job_id(); err != nil {
		return err
	}
	return nil
}

func (this *GetSyncJobConf) _xxx_xxx_Validator_Validate_job_id() error {
	if !(len(this.JobId) == 20) {
		return protovalidator.FieldError1("GetSyncJobConf", "the byte length of field 'job_id' must be equal to '20'", protovalidator.StringByteLenToString(this.JobId))
	}
	if !(strings.HasPrefix(this.JobId, "syj-")) {
		return protovalidator.FieldError1("GetSyncJobConf", "the value of field 'job_id' must start with string 'syj-'", this.JobId)
	}
	return nil
}

// Set default value for message request.GetSyncJobConf
func (this *GetSyncJobConf) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_job_id(); err != nil {
		return err
	}
	return nil
}

func (this *ListReleaseSyncJobs) _xxx_xxx_Validator_Validate_limit() error {
	if !(this.Limit > 0) {
		return protovalidator.FieldError1("ListReleaseSyncJobs", "the value of field 'limit' must be greater than '0'", protovalidator.Int32ToString(this.Limit))
	}
	if !(this.Limit <= 100) {
		return protovalidator.FieldError1("ListReleaseSyncJobs", "the value of field 'limit' must be less than or equal to '100'", protovalidator.Int32ToString(this.Limit))
	}
	return nil
}

func (this *ListReleaseSyncJobs) _xxx_xxx_Validator_Validate_offset() error {
	if !(this.Offset >= 0) {
		return protovalidator.FieldError1("ListReleaseSyncJobs", "the value of field 'offset' must be greater than or equal to '0'", protovalidator.Int32ToString(this.Offset))
	}
	return nil
}

func (this *ListReleaseSyncJobs) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("ListReleaseSyncJobs", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("ListReleaseSyncJobs", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

var _xxx_xxx_Validator_ListReleaseSyncJobs_In_SortBy = map[string]bool{"": true, "id": true, "name": true, "created": true, "updated": true}

func (this *ListReleaseSyncJobs) _xxx_xxx_Validator_Validate_sort_by() error {
	if !(_xxx_xxx_Validator_ListReleaseSyncJobs_In_SortBy[this.SortBy]) {
		return protovalidator.FieldError1("ListReleaseSyncJobs", "the value of field 'sort_by' must be one of '[ id name created updated]'", this.SortBy)
	}
	return nil
}

var _xxx_xxx_Validator_ListReleaseSyncJobs_InEnums_Status = map[pbmodel.SyncJobRelease_Status]bool{0: true, 1: true, 2: true, 3: true, 4: true}

func (this *ListReleaseSyncJobs) _xxx_xxx_Validator_Validate_status() error {
	if !(_xxx_xxx_Validator_ListReleaseSyncJobs_InEnums_Status[this.Status]) {
		return protovalidator.FieldError1("ListReleaseSyncJobs", "the value of field 'status' must in enums of '[0 1 2 3 4]'", protovalidator.Int32ToString(int32(this.Status)))
	}
	return nil
}

// Set default value for message request.ListReleaseSyncJobs
func (this *ListReleaseSyncJobs) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_limit(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_offset(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_space_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_sort_by(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_status(); err != nil {
		return err
	}
	return nil
}

func (this *UpdateReleaseSyncJobStatus) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("UpdateReleaseSyncJobStatus", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("UpdateReleaseSyncJobStatus", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

func (this *UpdateReleaseSyncJobStatus) _xxx_xxx_Validator_Validate_job_id() error {
	if !(len(this.JobId) == 20) {
		return protovalidator.FieldError1("UpdateReleaseSyncJobStatus", "the byte length of field 'job_id' must be equal to '20'", protovalidator.StringByteLenToString(this.JobId))
	}
	if !(strings.HasPrefix(this.JobId, "syj-")) {
		return protovalidator.FieldError1("UpdateReleaseSyncJobStatus", "the value of field 'job_id' must start with string 'syj-'", this.JobId)
	}
	return nil
}

func (this *UpdateReleaseSyncJobStatus) _xxx_xxx_Validator_Validate_version() error {
	if !(len(this.Version) == 16) {
		return protovalidator.FieldError1("UpdateReleaseSyncJobStatus", "the byte length of field 'version' must be equal to '16'", protovalidator.StringByteLenToString(this.Version))
	}
	return nil
}

var _xxx_xxx_Validator_UpdateReleaseSyncJobStatus_In_Status = map[pbmodel.SyncJobRelease_Status]bool{4: true}

func (this *UpdateReleaseSyncJobStatus) _xxx_xxx_Validator_Validate_status() error {
	if !(_xxx_xxx_Validator_UpdateReleaseSyncJobStatus_In_Status[this.Status]) {
		return protovalidator.FieldError1("UpdateReleaseSyncJobStatus", "the value of field 'status' must be one of '[4]'", protovalidator.Int32ToString(int32(this.Status)))
	}
	return nil
}

// Set default value for message request.UpdateReleaseSyncJobStatus
func (this *UpdateReleaseSyncJobStatus) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_space_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_job_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_version(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_status(); err != nil {
		return err
	}
	return nil
}

func (this *ReleaseSyncJob) _xxx_xxx_Validator_Validate_job_id() error {
	if !(len(this.JobId) == 20) {
		return protovalidator.FieldError1("ReleaseSyncJob", "the byte length of field 'job_id' must be equal to '20'", protovalidator.StringByteLenToString(this.JobId))
	}
	if !(strings.HasPrefix(this.JobId, "syj-")) {
		return protovalidator.FieldError1("ReleaseSyncJob", "the value of field 'job_id' must start with string 'syj-'", this.JobId)
	}
	return nil
}

func (this *ReleaseSyncJob) _xxx_xxx_Validator_Validate_created_by() error {
	if !(len(this.CreatedBy) < 65) {
		return protovalidator.FieldError1("ReleaseSyncJob", "the byte length of field 'created_by' must be less than '65'", protovalidator.StringByteLenToString(this.CreatedBy))
	}
	return nil
}

// Set default value for message request.ReleaseSyncJob
func (this *ReleaseSyncJob) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_job_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_created_by(); err != nil {
		return err
	}
	return nil
}

func (this *OfflineReleaseSyncJob) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("OfflineReleaseSyncJob", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("OfflineReleaseSyncJob", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

func (this *OfflineReleaseSyncJob) _xxx_xxx_Validator_Validate_job_id() error {
	if !(len(this.JobId) == 20) {
		return protovalidator.FieldError1("OfflineReleaseSyncJob", "the byte length of field 'job_id' must be equal to '20'", protovalidator.StringByteLenToString(this.JobId))
	}
	if !(strings.HasPrefix(this.JobId, "syj-")) {
		return protovalidator.FieldError1("OfflineReleaseSyncJob", "the value of field 'job_id' must start with string 'syj-'", this.JobId)
	}
	return nil
}

// Set default value for message request.OfflineReleaseSyncJob
func (this *OfflineReleaseSyncJob) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_space_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_job_id(); err != nil {
		return err
	}
	return nil
}

func (this *ReopenReleaseSyncJob) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("ReopenReleaseSyncJob", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("ReopenReleaseSyncJob", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

func (this *ReopenReleaseSyncJob) _xxx_xxx_Validator_Validate_job_id() error {
	if !(len(this.JobId) == 20) {
		return protovalidator.FieldError1("ReopenReleaseSyncJob", "the byte length of field 'job_id' must be equal to '20'", protovalidator.StringByteLenToString(this.JobId))
	}
	if !(strings.HasPrefix(this.JobId, "syj-")) {
		return protovalidator.FieldError1("ReopenReleaseSyncJob", "the value of field 'job_id' must start with string 'syj-'", this.JobId)
	}
	return nil
}

// Set default value for message request.ReopenReleaseSyncJob
func (this *ReopenReleaseSyncJob) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_space_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_job_id(); err != nil {
		return err
	}
	return nil
}

func (this *ConvertSyncJobMode) _xxx_xxx_Validator_Validate_job_id() error {
	if !(len(this.JobId) == 20) {
		return protovalidator.FieldError1("ConvertSyncJobMode", "the byte length of field 'job_id' must be equal to '20'", protovalidator.StringByteLenToString(this.JobId))
	}
	if !(strings.HasPrefix(this.JobId, "syj-")) {
		return protovalidator.FieldError1("ConvertSyncJobMode", "the value of field 'job_id' must start with string 'syj-'", this.JobId)
	}
	return nil
}

func (this *ConvertSyncJobMode) _xxx_xxx_Validator_Validate_conf() error {
	if dt, ok := interface{}(this.Conf).(interface{ Validate() error }); ok {
		if err := dt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Set default value for message request.ConvertSyncJobMode
func (this *ConvertSyncJobMode) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_job_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_conf(); err != nil {
		return err
	}
	return nil
}

func (this *ListSyncJobVersions) _xxx_xxx_Validator_Validate_limit() error {
	if !(this.Limit > 0) {
		return protovalidator.FieldError1("ListSyncJobVersions", "the value of field 'limit' must be greater than '0'", protovalidator.Int32ToString(this.Limit))
	}
	if !(this.Limit <= 100) {
		return protovalidator.FieldError1("ListSyncJobVersions", "the value of field 'limit' must be less than or equal to '100'", protovalidator.Int32ToString(this.Limit))
	}
	return nil
}

func (this *ListSyncJobVersions) _xxx_xxx_Validator_Validate_offset() error {
	if !(this.Offset >= 0) {
		return protovalidator.FieldError1("ListSyncJobVersions", "the value of field 'offset' must be greater than or equal to '0'", protovalidator.Int32ToString(this.Offset))
	}
	return nil
}

var _xxx_xxx_Validator_ListSyncJobVersions_In_SortBy = map[string]bool{"": true, "version": true, "created": true, "updated": true}

func (this *ListSyncJobVersions) _xxx_xxx_Validator_Validate_sort_by() error {
	if !(_xxx_xxx_Validator_ListSyncJobVersions_In_SortBy[this.SortBy]) {
		return protovalidator.FieldError1("ListSyncJobVersions", "the value of field 'sort_by' must be one of '[ version created updated]'", this.SortBy)
	}
	return nil
}

func (this *ListSyncJobVersions) _xxx_xxx_Validator_Validate_job_id() error {
	if !(len(this.JobId) == 20) {
		return protovalidator.FieldError1("ListSyncJobVersions", "the byte length of field 'job_id' must be equal to '20'", protovalidator.StringByteLenToString(this.JobId))
	}
	if !(strings.HasPrefix(this.JobId, "syj-")) {
		return protovalidator.FieldError1("ListSyncJobVersions", "the value of field 'job_id' must start with string 'syj-'", this.JobId)
	}
	return nil
}

// Set default value for message request.ListSyncJobVersions
func (this *ListSyncJobVersions) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_limit(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_offset(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_sort_by(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_job_id(); err != nil {
		return err
	}
	return nil
}

func (this *GenerateJobJson) _xxx_xxx_Validator_Validate_job_id() error {
	if !(len(this.JobId) == 20) {
		return protovalidator.FieldError1("GenerateJobJson", "the byte length of field 'job_id' must be equal to '20'", protovalidator.StringByteLenToString(this.JobId))
	}
	if !(strings.HasPrefix(this.JobId, "syj-")) {
		return protovalidator.FieldError1("GenerateJobJson", "the value of field 'job_id' must start with string 'syj-'", this.JobId)
	}
	return nil
}

// Set default value for message request.GenerateJobJson
func (this *GenerateJobJson) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_job_id(); err != nil {
		return err
	}
	return nil
}
