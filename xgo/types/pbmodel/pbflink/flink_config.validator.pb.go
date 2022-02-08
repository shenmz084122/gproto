// Code generated by protoc-gen-govalidator. DO NOT EDIT.
// versions:
// 		protoc-gen-govalidator 0.0.1
// source: proto/types/model/flink/flink_config.proto

package pbflink

import (
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbdefaults"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbgosql"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
	protovalidator "github.com/yu31/protoc-plugin/xgo/pkg/protovalidator"
)

func (this *FlinkConfig) _xxx_xxx_Validator_Validate_custom() error {
	for _, item := range this.Custom {
		_ = item // To avoid unused panics.
		if dt, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := dt.Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (this *FlinkConfig) _xxx_xxx_Validator_Validate_restart_strategy() error {
	if dt, ok := interface{}(this.RestartStrategy).(interface{ Validate() error }); ok {
		if err := dt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (this *FlinkConfig) _xxx_xxx_Validator_Validate_logger() error {
	if dt, ok := interface{}(this.Logger).(interface{ Validate() error }); ok {
		if err := dt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Set default value for message flink.FlinkConfig
func (this *FlinkConfig) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_custom(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_restart_strategy(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_logger(); err != nil {
		return err
	}
	return nil
}

func (this *FlinkConfig_Item) _xxx_xxx_Validator_Validate_key() error {
	if !(len(this.Key) <= 1024) {
		return protovalidator.FieldError1("FlinkConfig_Item", "the byte length of field 'key' must be less than or equal to '1024'", protovalidator.StringByteLenToString(this.Key))
	}
	return nil
}

func (this *FlinkConfig_Item) _xxx_xxx_Validator_Validate_value() error {
	if !(len(this.Value) <= 1024) {
		return protovalidator.FieldError1("FlinkConfig_Item", "the byte length of field 'value' must be less than or equal to '1024'", protovalidator.StringByteLenToString(this.Value))
	}
	return nil
}

// Set default value for message flink.FlinkConfig.Item
func (this *FlinkConfig_Item) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_key(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_value(); err != nil {
		return err
	}
	return nil
}

var _xxx_xxx_Validator_FlinkConfig_RestartStrategy_In_RestartStrategy = map[string]bool{"none": true, "fixed-delay": true, "failure-rate": true}

func (this *FlinkConfig_RestartStrategy) _xxx_xxx_Validator_Validate_restart_strategy() error {
	if !(_xxx_xxx_Validator_FlinkConfig_RestartStrategy_In_RestartStrategy[this.RestartStrategy]) {
		return protovalidator.FieldError1("FlinkConfig_RestartStrategy", "the value of field 'restart_strategy' must be one of '[none fixed-delay failure-rate]'", this.RestartStrategy)
	}
	return nil
}

func (this *FlinkConfig_RestartStrategy) _xxx_xxx_Validator_Validate_fixed_delay_attempts() error {
	if !(this.FixedDelayAttempts >= 1) {
		return protovalidator.FieldError1("FlinkConfig_RestartStrategy", "the value of field 'fixed_delay_attempts' must be greater than or equal to '1'", protovalidator.Int32ToString(this.FixedDelayAttempts))
	}
	if !(this.FixedDelayAttempts <= 1000) {
		return protovalidator.FieldError1("FlinkConfig_RestartStrategy", "the value of field 'fixed_delay_attempts' must be less than or equal to '1000'", protovalidator.Int32ToString(this.FixedDelayAttempts))
	}
	return nil
}

func (this *FlinkConfig_RestartStrategy) _xxx_xxx_Validator_Validate_failure_rate_delay() error {
	if !(this.FailureRateDelay >= 1) {
		return protovalidator.FieldError1("FlinkConfig_RestartStrategy", "the value of field 'failure_rate_delay' must be greater than or equal to '1'", protovalidator.Int32ToString(this.FailureRateDelay))
	}
	if !(this.FailureRateDelay <= 86400) {
		return protovalidator.FieldError1("FlinkConfig_RestartStrategy", "the value of field 'failure_rate_delay' must be less than or equal to '86400'", protovalidator.Int32ToString(this.FailureRateDelay))
	}
	return nil
}

func (this *FlinkConfig_RestartStrategy) _xxx_xxx_Validator_Validate_failure_rate_failure_rate_interval() error {
	if !(this.FailureRateFailureRateInterval >= 1) {
		return protovalidator.FieldError1("FlinkConfig_RestartStrategy", "the value of field 'failure_rate_failure_rate_interval' must be greater than or equal to '1'", protovalidator.Int32ToString(this.FailureRateFailureRateInterval))
	}
	if !(this.FailureRateFailureRateInterval <= 86400) {
		return protovalidator.FieldError1("FlinkConfig_RestartStrategy", "the value of field 'failure_rate_failure_rate_interval' must be less than or equal to '86400'", protovalidator.Int32ToString(this.FailureRateFailureRateInterval))
	}
	return nil
}

// Set default value for message flink.FlinkConfig.RestartStrategy
func (this *FlinkConfig_RestartStrategy) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_restart_strategy(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_fixed_delay_attempts(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_failure_rate_delay(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_failure_rate_failure_rate_interval(); err != nil {
		return err
	}
	return nil
}

var _xxx_xxx_Validator_FlinkConfig_Logger_In_RootLogLevel = map[string]bool{"TRACE": true, "DEBUG": true, "INFO": true, "WARN": true, "ERROR": true}

func (this *FlinkConfig_Logger) _xxx_xxx_Validator_Validate_root_log_level() error {
	if !(_xxx_xxx_Validator_FlinkConfig_Logger_In_RootLogLevel[this.RootLogLevel]) {
		return protovalidator.FieldError1("FlinkConfig_Logger", "the value of field 'root_log_level' must be one of '[TRACE DEBUG INFO WARN ERROR]'", this.RootLogLevel)
	}
	return nil
}

// Set default value for message flink.FlinkConfig.Logger
func (this *FlinkConfig_Logger) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_root_log_level(); err != nil {
		return err
	}
	return nil
}
