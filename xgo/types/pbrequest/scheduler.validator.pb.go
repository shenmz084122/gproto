// Code generated by protoc-gen-govalidator. DO NOT EDIT.
// versions:
// 		protoc-gen-govalidator 0.0.1
// source: proto/types/request/scheduler.proto

package pbrequest

import (
	_ "github.com/DataWorkbench/gproto/xgo/types/pbmodel"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
	protovalidator "github.com/yu31/protoc-plugin/xgo/pkg/protovalidator"
	strings "strings"
)

func (this *SubmitStreamJob) _xxx_xxx_Validator_Validate_property() error {
	if !(this.Property != nil) {
		return protovalidator.FieldError2("SubmitStreamJob", "the value of field 'property' cannot be null")
	}
	if dt, ok := interface{}(this.Property).(interface{ Validate() error }); ok {
		if err := dt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (this *SubmitStreamJob) _xxx_xxx_Validator_Validate_flink_url() error {
	if !(this.FlinkUrl != "") {
		return protovalidator.FieldError1("SubmitStreamJob", "the value of field 'flink_url' must be not equal to ''", this.FlinkUrl)
	}
	return nil
}

func (this *SubmitStreamJob) _xxx_xxx_Validator_Validate_flink_version() error {
	if !(this.FlinkVersion != "") {
		return protovalidator.FieldError1("SubmitStreamJob", "the value of field 'flink_version' must be not equal to ''", this.FlinkVersion)
	}
	return nil
}

// Set default value for message request.SubmitStreamJob
func (this *SubmitStreamJob) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_property(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_flink_url(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_flink_version(); err != nil {
		return err
	}
	return nil
}

func (this *StopStreamJob) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("StopStreamJob", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("StopStreamJob", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

func (this *StopStreamJob) _xxx_xxx_Validator_Validate_job_id() error {
	if !(len(this.JobId) == 20) {
		return protovalidator.FieldError1("StopStreamJob", "the byte length of field 'job_id' must be equal to '20'", protovalidator.StringByteLenToString(this.JobId))
	}
	if !(strings.HasPrefix(this.JobId, "stj-")) {
		return protovalidator.FieldError1("StopStreamJob", "the value of field 'job_id' must start with string 'stj-'", this.JobId)
	}
	return nil
}

// Set default value for message request.StopStreamJob
func (this *StopStreamJob) Validate() error {
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

func (this *DeleteStreamJobsBySpaceIds) _xxx_xxx_Validator_Validate_space_ids() error {
	for _, item := range this.SpaceIds {
		_ = item // To avoid unused panics.
		if !(len(item) == 20) {
			return protovalidator.FieldError1("DeleteStreamJobsBySpaceIds", "the byte length of array item where in field 'space_ids' must be equal to '20'", protovalidator.StringByteLenToString(item))
		}
		if !(strings.HasPrefix(item, "wks-")) {
			return protovalidator.FieldError1("DeleteStreamJobsBySpaceIds", "the value of array item where in field 'space_ids' must start with string 'wks-'", item)
		}
	}
	return nil
}

// Set default value for message request.DeleteStreamJobsBySpaceIds
func (this *DeleteStreamJobsBySpaceIds) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_space_ids(); err != nil {
		return err
	}
	return nil
}

func (this *DeleteStreamJobsByJobIds) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("DeleteStreamJobsByJobIds", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("DeleteStreamJobsByJobIds", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

func (this *DeleteStreamJobsByJobIds) _xxx_xxx_Validator_Validate_job_ids() error {
	for _, item := range this.JobIds {
		_ = item // To avoid unused panics.
		if !(len(item) == 20) {
			return protovalidator.FieldError1("DeleteStreamJobsByJobIds", "the byte length of array item where in field 'job_ids' must be equal to '20'", protovalidator.StringByteLenToString(item))
		}
		if !(strings.HasPrefix(item, "stj-")) {
			return protovalidator.FieldError1("DeleteStreamJobsByJobIds", "the value of array item where in field 'job_ids' must start with string 'stj-'", item)
		}
	}
	return nil
}

// Set default value for message request.DeleteStreamJobsByJobIds
func (this *DeleteStreamJobsByJobIds) Validate() error {
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

func (this *SubmitSyncJob) _xxx_xxx_Validator_Validate_property() error {
	if !(this.Property != nil) {
		return protovalidator.FieldError2("SubmitSyncJob", "the value of field 'property' cannot be null")
	}
	if dt, ok := interface{}(this.Property).(interface{ Validate() error }); ok {
		if err := dt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Set default value for message request.SubmitSyncJob
func (this *SubmitSyncJob) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_property(); err != nil {
		return err
	}
	return nil
}

func (this *StopSyncJob) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("StopSyncJob", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("StopSyncJob", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

func (this *StopSyncJob) _xxx_xxx_Validator_Validate_job_id() error {
	if !(len(this.JobId) == 20) {
		return protovalidator.FieldError1("StopSyncJob", "the byte length of field 'job_id' must be equal to '20'", protovalidator.StringByteLenToString(this.JobId))
	}
	if !(strings.HasPrefix(this.JobId, "syj-")) {
		return protovalidator.FieldError1("StopSyncJob", "the value of field 'job_id' must start with string 'syj-'", this.JobId)
	}
	return nil
}

// Set default value for message request.StopSyncJob
func (this *StopSyncJob) Validate() error {
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

func (this *DeleteSyncJobsBySpaceIds) _xxx_xxx_Validator_Validate_space_ids() error {
	for _, item := range this.SpaceIds {
		_ = item // To avoid unused panics.
		if !(len(item) == 20) {
			return protovalidator.FieldError1("DeleteSyncJobsBySpaceIds", "the byte length of array item where in field 'space_ids' must be equal to '20'", protovalidator.StringByteLenToString(item))
		}
		if !(strings.HasPrefix(item, "wks-")) {
			return protovalidator.FieldError1("DeleteSyncJobsBySpaceIds", "the value of array item where in field 'space_ids' must start with string 'wks-'", item)
		}
	}
	return nil
}

// Set default value for message request.DeleteSyncJobsBySpaceIds
func (this *DeleteSyncJobsBySpaceIds) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_space_ids(); err != nil {
		return err
	}
	return nil
}

func (this *DeleteSyncJobsByJobIds) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("DeleteSyncJobsByJobIds", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("DeleteSyncJobsByJobIds", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

func (this *DeleteSyncJobsByJobIds) _xxx_xxx_Validator_Validate_job_ids() error {
	for _, item := range this.JobIds {
		_ = item // To avoid unused panics.
		if !(len(item) == 20) {
			return protovalidator.FieldError1("DeleteSyncJobsByJobIds", "the byte length of array item where in field 'job_ids' must be equal to '20'", protovalidator.StringByteLenToString(item))
		}
		if !(strings.HasPrefix(item, "syj-")) {
			return protovalidator.FieldError1("DeleteSyncJobsByJobIds", "the value of array item where in field 'job_ids' must start with string 'syj-'", item)
		}
	}
	return nil
}

// Set default value for message request.DeleteSyncJobsByJobIds
func (this *DeleteSyncJobsByJobIds) Validate() error {
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
