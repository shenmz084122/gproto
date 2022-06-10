// Code generated by protoc-gen-govalidator. DO NOT EDIT.
// versions:
// 		protoc-gen-govalidator 0.0.1
// source: proto/types/request/flink_developer.proto

package pbrequest

import (
	_ "github.com/DataWorkbench/gproto/xgo/types/pbmodel"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbdefaults"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
	protovalidator "github.com/yu31/protoc-plugin/xgo/pkg/protovalidator"
)

func (this *SubmitFlinkJob) _xxx_xxx_Validator_Validate_flink_id() error {
	if !(len(this.FlinkId) == 32) {
		return protovalidator.FieldError1("SubmitFlinkJob", "the byte length of field 'flink_id' must be equal to '32'", protovalidator.StringByteLenToString(this.FlinkId))
	}
	return nil
}

func (this *SubmitFlinkJob) _xxx_xxx_Validator_Validate_flink_url() error {
	if !(this.FlinkUrl != "") {
		return protovalidator.FieldError1("SubmitFlinkJob", "the value of field 'flink_url' must be not equal to ''", this.FlinkUrl)
	}
	return nil
}

func (this *SubmitFlinkJob) _xxx_xxx_Validator_Validate_flink_version() error {
	if !(this.FlinkVersion != "") {
		return protovalidator.FieldError1("SubmitFlinkJob", "the value of field 'flink_version' must be not equal to ''", this.FlinkVersion)
	}
	return nil
}

func (this *SubmitFlinkJob) _xxx_xxx_Validator_Validate_args() error {
	if dt, ok := interface{}(this.Args).(interface{ Validate() error }); ok {
		if err := dt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (this *SubmitFlinkJob) _xxx_xxx_Validator_Validate_code() error {
	if dt, ok := interface{}(this.Code).(interface{ Validate() error }); ok {
		if err := dt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Set default value for message request.SubmitFlinkJob
func (this *SubmitFlinkJob) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_flink_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_flink_url(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_flink_version(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_args(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_code(); err != nil {
		return err
	}
	return nil
}

func (this *SubmitFlinkJobInteractive) _xxx_xxx_Validator_Validate_job() error {
	if dt, ok := interface{}(this.Job).(interface{ Validate() error }); ok {
		if err := dt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (this *SubmitFlinkJobInteractive) _xxx_xxx_Validator_Validate_max_count() error {
	if !(this.MaxCount > 0) {
		return protovalidator.FieldError1("SubmitFlinkJobInteractive", "the value of field 'max_count' must be greater than '0'", protovalidator.Int32ToString(this.MaxCount))
	}
	if !(this.MaxCount <= 100) {
		return protovalidator.FieldError1("SubmitFlinkJobInteractive", "the value of field 'max_count' must be less than or equal to '100'", protovalidator.Int32ToString(this.MaxCount))
	}
	return nil
}

func (this *SubmitFlinkJobInteractive) _xxx_xxx_Validator_Validate_refresh_interval() error {
	if !(this.RefreshInterval >= 1000) {
		return protovalidator.FieldError1("SubmitFlinkJobInteractive", "the value of field 'refresh_interval' must be greater than or equal to '1000'", protovalidator.Int32ToString(this.RefreshInterval))
	}
	if !(this.RefreshInterval <= 10000) {
		return protovalidator.FieldError1("SubmitFlinkJobInteractive", "the value of field 'refresh_interval' must be less than or equal to '10000'", protovalidator.Int32ToString(this.RefreshInterval))
	}
	return nil
}

// Set default value for message request.SubmitFlinkJobInteractive
func (this *SubmitFlinkJobInteractive) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_job(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_max_count(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_refresh_interval(); err != nil {
		return err
	}
	return nil
}

func (this *ValidateFlinkJob) _xxx_xxx_Validator_Validate_code() error {
	if dt, ok := interface{}(this.Code).(interface{ Validate() error }); ok {
		if err := dt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Set default value for message request.ValidateFlinkJob
func (this *ValidateFlinkJob) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_code(); err != nil {
		return err
	}
	return nil
}

func (this *ValidateFlinkJobV2) _xxx_xxx_Validator_Validate_flink_url() error {
	if !(this.FlinkUrl != "") {
		return protovalidator.FieldError1("ValidateFlinkJobV2", "the value of field 'flink_url' must be not equal to ''", this.FlinkUrl)
	}
	return nil
}

func (this *ValidateFlinkJobV2) _xxx_xxx_Validator_Validate_flink_version() error {
	if !(this.FlinkVersion != "") {
		return protovalidator.FieldError1("ValidateFlinkJobV2", "the value of field 'flink_version' must be not equal to ''", this.FlinkVersion)
	}
	return nil
}

func (this *ValidateFlinkJobV2) _xxx_xxx_Validator_Validate_args() error {
	if dt, ok := interface{}(this.Args).(interface{ Validate() error }); ok {
		if err := dt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (this *ValidateFlinkJobV2) _xxx_xxx_Validator_Validate_code() error {
	if dt, ok := interface{}(this.Code).(interface{ Validate() error }); ok {
		if err := dt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Set default value for message request.ValidateFlinkJob_v2
func (this *ValidateFlinkJobV2) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_flink_url(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_flink_version(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_args(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_code(); err != nil {
		return err
	}
	return nil
}
