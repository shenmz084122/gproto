// Code generated by protoc-gen-govalidator. DO NOT EDIT.
// versions:
// 		protoc-gen-govalidator 0.0.1
// source: proto/types/model/stream_job.proto

package pbmodel

import (
	_ "github.com/DataWorkbench/gproto/xgo/types/pbmodel/pbflink"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbdefaults"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbgosql"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
	protovalidator "github.com/yu31/protoc-plugin/xgo/pkg/protovalidator"
	strconv "strconv"
	strings "strings"
	utf8 "unicode/utf8"
)

func (this *StreamJob) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("StreamJob", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("StreamJob", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

func (this *StreamJob) _xxx_xxx_Validator_CheckIf_pid() bool {
	if !(this.Pid != "") {
		return false
	}
	return true
}

func (this *StreamJob) _xxx_xxx_Validator_Validate_pid() error {
	if !this._xxx_xxx_Validator_CheckIf_pid() {
		return nil
	}
	if !(len(this.Pid) == 20) {
		return protovalidator.FieldError1("StreamJob", "the byte length of field 'pid' must be equal to '20'", protovalidator.StringByteLenToString(this.Pid))
	}
	if !(strings.HasPrefix(this.Pid, "stj-")) {
		return protovalidator.FieldError1("StreamJob", "the value of field 'pid' must start with string 'stj-'", this.Pid)
	}
	return nil
}

func (this *StreamJob) _xxx_xxx_Validator_Validate_id() error {
	if !(len(this.Id) == 20) {
		return protovalidator.FieldError1("StreamJob", "the byte length of field 'id' must be equal to '20'", protovalidator.StringByteLenToString(this.Id))
	}
	if !(strings.HasPrefix(this.Id, "stj-")) {
		return protovalidator.FieldError1("StreamJob", "the value of field 'id' must start with string 'stj-'", this.Id)
	}
	return nil
}

func (this *StreamJob) _xxx_xxx_Validator_Validate_version() error {
	if !(len(this.Version) == 16) {
		return protovalidator.FieldError1("StreamJob", "the byte length of field 'version' must be equal to '16'", protovalidator.StringByteLenToString(this.Version))
	}
	return nil
}

func (this *StreamJob) _xxx_xxx_Validator_Validate_name() error {
	if !(len(this.Name) > 1) {
		return protovalidator.FieldError1("StreamJob", "the byte length of field 'name' must be greater than '1'", protovalidator.StringByteLenToString(this.Name))
	}
	if !(len(this.Name) <= 128) {
		return protovalidator.FieldError1("StreamJob", "the byte length of field 'name' must be less than or equal to '128'", protovalidator.StringByteLenToString(this.Name))
	}
	return nil
}

func (this *StreamJob) _xxx_xxx_Validator_Validate_desc() error {
	if !(utf8.RuneCountInString(this.Desc) <= 1024) {
		return protovalidator.FieldError1("StreamJob", "the character length of field 'desc' must be less than or equal to '1024'", protovalidator.StringCharsetLenToString(this.Desc))
	}
	return nil
}

func (this *StreamJob) _xxx_xxx_Validator_CheckIf_type() bool {
	if !(this.IsDirectory == false) {
		return false
	}
	return true
}

var _xxx_xxx_Validator_StreamJob_InEnums_Type = map[StreamJob_Type]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true}

func (this *StreamJob) _xxx_xxx_Validator_Validate_type() error {
	if !this._xxx_xxx_Validator_CheckIf_type() {
		return nil
	}
	if !(this.Type > 0) {
		return protovalidator.FieldError1("StreamJob", "the value of field 'type' must be greater than '0'", protovalidator.Int32ToString(int32(this.Type)))
	}
	if !(_xxx_xxx_Validator_StreamJob_InEnums_Type[this.Type]) {
		return protovalidator.FieldError1("StreamJob", "the value of field 'type' must in enums of '[0 1 2 3 4 5]'", protovalidator.Int32ToString(int32(this.Type)))
	}
	return nil
}

var _xxx_xxx_Validator_StreamJob_InEnums_Status = map[StreamJob_Status]bool{0: true, 1: true, 2: true}

func (this *StreamJob) _xxx_xxx_Validator_Validate_status() error {
	if !(this.Status > 0) {
		return protovalidator.FieldError1("StreamJob", "the value of field 'status' must be greater than '0'", protovalidator.Int32ToString(int32(this.Status)))
	}
	if !(_xxx_xxx_Validator_StreamJob_InEnums_Status[this.Status]) {
		return protovalidator.FieldError1("StreamJob", "the value of field 'status' must in enums of '[0 1 2]'", protovalidator.Int32ToString(int32(this.Status)))
	}
	return nil
}

func (this *StreamJob) _xxx_xxx_Validator_Validate_created_by() error {
	if !(len(this.CreatedBy) > 0) {
		return protovalidator.FieldError1("StreamJob", "the byte length of field 'created_by' must be greater than '0'", protovalidator.StringByteLenToString(this.CreatedBy))
	}
	if !(len(this.CreatedBy) < 65) {
		return protovalidator.FieldError1("StreamJob", "the byte length of field 'created_by' must be less than '65'", protovalidator.StringByteLenToString(this.CreatedBy))
	}
	return nil
}

func (this *StreamJob) _xxx_xxx_Validator_Validate_created() error {
	if !(this.Created > 0) {
		return protovalidator.FieldError1("StreamJob", "the value of field 'created' must be greater than '0'", protovalidator.Int64ToString(this.Created))
	}
	return nil
}

func (this *StreamJob) _xxx_xxx_Validator_Validate_updated() error {
	if !(this.Updated > 0) {
		return protovalidator.FieldError1("StreamJob", "the value of field 'updated' must be greater than '0'", protovalidator.Int64ToString(this.Updated))
	}
	return nil
}

// Set default value for message model.StreamJob
func (this *StreamJob) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_space_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_pid(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_version(); err != nil {
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
	if err := this._xxx_xxx_Validator_Validate_status(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_created_by(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_created(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_updated(); err != nil {
		return err
	}
	return nil
}

func (this *StreamJobProperty) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("StreamJobProperty", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("StreamJobProperty", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

func (this *StreamJobProperty) _xxx_xxx_Validator_Validate_id() error {
	if !(len(this.Id) == 20) {
		return protovalidator.FieldError1("StreamJobProperty", "the byte length of field 'id' must be equal to '20'", protovalidator.StringByteLenToString(this.Id))
	}
	if !(strings.HasPrefix(this.Id, "stj-")) {
		return protovalidator.FieldError1("StreamJobProperty", "the value of field 'id' must start with string 'stj-'", this.Id)
	}
	return nil
}

func (this *StreamJobProperty) _xxx_xxx_Validator_Validate_version() error {
	if !(len(this.Version) == 16) {
		return protovalidator.FieldError1("StreamJobProperty", "the byte length of field 'version' must be equal to '16'", protovalidator.StringByteLenToString(this.Version))
	}
	return nil
}

func (this *StreamJobProperty) _xxx_xxx_Validator_Validate_code() error {
	if dt, ok := interface{}(this.Code).(interface{ Validate() error }); ok {
		if err := dt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (this *StreamJobProperty) _xxx_xxx_Validator_Validate_args() error {
	if dt, ok := interface{}(this.Args).(interface{ Validate() error }); ok {
		if err := dt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (this *StreamJobProperty) _xxx_xxx_Validator_Validate_schedule() error {
	if dt, ok := interface{}(this.Schedule).(interface{ Validate() error }); ok {
		if err := dt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Set default value for message model.StreamJobProperty
func (this *StreamJobProperty) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_space_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_version(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_code(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_args(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_schedule(); err != nil {
		return err
	}
	return nil
}

var _xxx_xxx_Validator_StreamJobCode_InEnums_Type = map[StreamJob_Type]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true}

func (this *StreamJobCode) _xxx_xxx_Validator_Validate_type() error {
	if !(this.Type > 0) {
		return protovalidator.FieldError1("StreamJobCode", "the value of field 'type' must be greater than '0'", protovalidator.Int32ToString(int32(this.Type)))
	}
	if !(_xxx_xxx_Validator_StreamJobCode_InEnums_Type[this.Type]) {
		return protovalidator.FieldError1("StreamJobCode", "the value of field 'type' must in enums of '[0 1 2 3 4 5]'", protovalidator.Int32ToString(int32(this.Type)))
	}
	return nil
}

func (this *StreamJobCode) _xxx_xxx_Validator_CheckIf_operators() bool {
	if !(this.Type == 1) {
		return false
	}
	return true
}

func (this *StreamJobCode) _xxx_xxx_Validator_Validate_operators() error {
	if !this._xxx_xxx_Validator_CheckIf_operators() {
		return nil
	}
	if !(len(this.Operators) > 0) {
		return protovalidator.FieldError1("StreamJobCode", "the length of field 'operators' must be greater than '0'", strconv.Itoa(len(this.Operators)))
	}
	for _, item := range this.Operators {
		_ = item // To avoid unused panics.
		if dt, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := dt.Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (this *StreamJobCode) _xxx_xxx_Validator_CheckIf_sql() bool {
	if !(this.Type == 2) {
		return false
	}
	return true
}

func (this *StreamJobCode) _xxx_xxx_Validator_Validate_sql() error {
	if !this._xxx_xxx_Validator_CheckIf_sql() {
		return nil
	}
	if !(this.Sql != nil) {
		return protovalidator.FieldError2("StreamJobCode", "the value of field 'sql' cannot be null")
	}
	if dt, ok := interface{}(this.Sql).(interface{ Validate() error }); ok {
		if err := dt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (this *StreamJobCode) _xxx_xxx_Validator_CheckIf_jar() bool {
	if !(this.Type == 3) {
		return false
	}
	return true
}

func (this *StreamJobCode) _xxx_xxx_Validator_Validate_jar() error {
	if !this._xxx_xxx_Validator_CheckIf_jar() {
		return nil
	}
	if !(this.Jar != nil) {
		return protovalidator.FieldError2("StreamJobCode", "the value of field 'jar' cannot be null")
	}
	if dt, ok := interface{}(this.Jar).(interface{ Validate() error }); ok {
		if err := dt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (this *StreamJobCode) _xxx_xxx_Validator_CheckIf_python_code() bool {
	if !(this.Type == 4) {
		return false
	}
	return true
}

func (this *StreamJobCode) _xxx_xxx_Validator_Validate_python_code() error {
	if !this._xxx_xxx_Validator_CheckIf_python_code() {
		return nil
	}
	if !(this.PythonCode != nil) {
		return protovalidator.FieldError2("StreamJobCode", "the value of field 'python_code' cannot be null")
	}
	if dt, ok := interface{}(this.PythonCode).(interface{ Validate() error }); ok {
		if err := dt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (this *StreamJobCode) _xxx_xxx_Validator_CheckIf_python_file() bool {
	if !(this.Type == 5) {
		return false
	}
	return true
}

func (this *StreamJobCode) _xxx_xxx_Validator_Validate_python_file() error {
	if !this._xxx_xxx_Validator_CheckIf_python_file() {
		return nil
	}
	if !(this.PythonFile != nil) {
		return protovalidator.FieldError2("StreamJobCode", "the value of field 'python_file' cannot be null")
	}
	if dt, ok := interface{}(this.PythonFile).(interface{ Validate() error }); ok {
		if err := dt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Set default value for message model.StreamJobCode
func (this *StreamJobCode) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_type(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_operators(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_sql(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_jar(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_python_code(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_python_file(); err != nil {
		return err
	}
	return nil
}

func (this *StreamJobArgs) _xxx_xxx_Validator_CheckIf_cluster_id() bool {
	if !(this.ClusterId != "") {
		return false
	}
	return true
}

func (this *StreamJobArgs) _xxx_xxx_Validator_Validate_cluster_id() error {
	if !this._xxx_xxx_Validator_CheckIf_cluster_id() {
		return nil
	}
	if !(len(this.ClusterId) == 20) {
		return protovalidator.FieldError1("StreamJobArgs", "the byte length of field 'cluster_id' must be equal to '20'", protovalidator.StringByteLenToString(this.ClusterId))
	}
	if !(strings.HasPrefix(this.ClusterId, "cfi-")) {
		return protovalidator.FieldError1("StreamJobArgs", "the value of field 'cluster_id' must start with string 'cfi-'", this.ClusterId)
	}
	return nil
}

func (this *StreamJobArgs) _xxx_xxx_Validator_Validate_parallelism() error {
	if !(this.Parallelism >= 0) {
		return protovalidator.FieldError1("StreamJobArgs", "the value of field 'parallelism' must be greater than or equal to '0'", protovalidator.Int32ToString(this.Parallelism))
	}
	if !(this.Parallelism <= 100) {
		return protovalidator.FieldError1("StreamJobArgs", "the value of field 'parallelism' must be less than or equal to '100'", protovalidator.Int32ToString(this.Parallelism))
	}
	return nil
}

func (this *StreamJobArgs) _xxx_xxx_Validator_Validate_files() error {
	if !(len(this.Files) <= 100) {
		return protovalidator.FieldError1("StreamJobArgs", "the length of field 'files' must be less than or equal to '100'", strconv.Itoa(len(this.Files)))
	}
	for _, item := range this.Files {
		_ = item // To avoid unused panics.
		if !(strings.HasPrefix(item, "res-")) {
			return protovalidator.FieldError1("StreamJobArgs", "the value of array item where in field 'files' must start with string 'res-'", item)
		}
	}
	return nil
}

func (this *StreamJobArgs) _xxx_xxx_Validator_Validate_delete_files() error {
	for _, item := range this.DeleteFiles {
		_ = item // To avoid unused panics.
		if !(strings.HasPrefix(item, "res-")) {
			return protovalidator.FieldError1("StreamJobArgs", "the value of array item where in field 'delete_files' must start with string 'res-'", item)
		}
	}
	return nil
}

func (this *StreamJobArgs) _xxx_xxx_Validator_CheckIf_delete_cluster_id() bool {
	if !(this.DeleteClusterId != "") {
		return false
	}
	return true
}

func (this *StreamJobArgs) _xxx_xxx_Validator_Validate_delete_cluster_id() error {
	if !this._xxx_xxx_Validator_CheckIf_delete_cluster_id() {
		return nil
	}
	if !(len(this.DeleteClusterId) == 20) {
		return protovalidator.FieldError1("StreamJobArgs", "the byte length of field 'delete_cluster_id' must be equal to '20'", protovalidator.StringByteLenToString(this.DeleteClusterId))
	}
	if !(strings.HasPrefix(this.DeleteClusterId, "cfi-")) {
		return protovalidator.FieldError1("StreamJobArgs", "the value of field 'delete_cluster_id' must start with string 'cfi-'", this.DeleteClusterId)
	}
	return nil
}

func (this *StreamJobArgs) _xxx_xxx_Validator_Validate_py_archives() error {
	if !(len(this.PyArchives) <= 100) {
		return protovalidator.FieldError1("StreamJobArgs", "the length of field 'py_archives' must be less than or equal to '100'", strconv.Itoa(len(this.PyArchives)))
	}
	for _, item := range this.PyArchives {
		_ = item // To avoid unused panics.
		if !(strings.HasPrefix(item, "res-")) {
			return protovalidator.FieldError1("StreamJobArgs", "the value of array item where in field 'py_archives' must start with string 'res-'", item)
		}
	}
	return nil
}

func (this *StreamJobArgs) _xxx_xxx_Validator_Validate_py_files() error {
	if !(len(this.PyFiles) <= 100) {
		return protovalidator.FieldError1("StreamJobArgs", "the length of field 'py_files' must be less than or equal to '100'", strconv.Itoa(len(this.PyFiles)))
	}
	for _, item := range this.PyFiles {
		_ = item // To avoid unused panics.
		if !(strings.HasPrefix(item, "res-")) {
			return protovalidator.FieldError1("StreamJobArgs", "the value of array item where in field 'py_files' must start with string 'res-'", item)
		}
	}
	return nil
}

func (this *StreamJobArgs) _xxx_xxx_Validator_Validate_py_requirements() error {
	if !(len(this.PyRequirements) <= 100) {
		return protovalidator.FieldError1("StreamJobArgs", "the length of field 'py_requirements' must be less than or equal to '100'", strconv.Itoa(len(this.PyRequirements)))
	}
	for _, item := range this.PyRequirements {
		_ = item // To avoid unused panics.
		if !(strings.HasPrefix(item, "res-")) {
			return protovalidator.FieldError1("StreamJobArgs", "the value of array item where in field 'py_requirements' must start with string 'res-'", item)
		}
	}
	return nil
}

func (this *StreamJobArgs) _xxx_xxx_Validator_Validate_py_executable() error {
	if !(len(this.PyExecutable) <= 255) {
		return protovalidator.FieldError1("StreamJobArgs", "the byte length of field 'py_executable' must be less than or equal to '255'", protovalidator.StringByteLenToString(this.PyExecutable))
	}
	return nil
}

func (this *StreamJobArgs) _xxx_xxx_Validator_Validate_model() error {
	if !(this.Model >= 0) {
		return protovalidator.FieldError1("StreamJobArgs", "the value of field 'model' must be greater than or equal to '0'", protovalidator.Int32ToString(this.Model))
	}
	if !(this.Model <= 100) {
		return protovalidator.FieldError1("StreamJobArgs", "the value of field 'model' must be less than or equal to '100'", protovalidator.Int32ToString(this.Model))
	}
	return nil
}

// Set default value for message model.StreamJobArgs
func (this *StreamJobArgs) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_cluster_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_parallelism(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_files(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_delete_files(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_delete_cluster_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_py_archives(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_py_files(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_py_requirements(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_py_executable(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_model(); err != nil {
		return err
	}
	return nil
}

var _xxx_xxx_Validator_StreamJobSchedule_InEnums_SchedulePolicy = map[StreamJobSchedule_SchedulePolicy]bool{0: true, 1: true, 2: true, 3: true}

func (this *StreamJobSchedule) _xxx_xxx_Validator_Validate_schedule_policy() error {
	if !(this.SchedulePolicy > 0) {
		return protovalidator.FieldError1("StreamJobSchedule", "the value of field 'schedule_policy' must be greater than '0'", protovalidator.Int32ToString(int32(this.SchedulePolicy)))
	}
	if !(_xxx_xxx_Validator_StreamJobSchedule_InEnums_SchedulePolicy[this.SchedulePolicy]) {
		return protovalidator.FieldError1("StreamJobSchedule", "the value of field 'schedule_policy' must in enums of '[0 1 2 3]'", protovalidator.Int32ToString(int32(this.SchedulePolicy)))
	}
	return nil
}

func (this *StreamJobSchedule) _xxx_xxx_Validator_CheckIf_executed() bool {
	if !(this.SchedulePolicy == 2) {
		return false
	}
	return true
}

func (this *StreamJobSchedule) _xxx_xxx_Validator_Validate_executed() error {
	if !this._xxx_xxx_Validator_CheckIf_executed() {
		return nil
	}
	if !(this.Executed >= 31507200) {
		return protovalidator.FieldError1("StreamJobSchedule", "the value of field 'executed' must be greater than or equal to '31507200'", protovalidator.Int64ToString(this.Executed))
	}
	return nil
}

func (this *StreamJobSchedule) _xxx_xxx_Validator_Validate_started() error {
	if !(this.Started >= 0) {
		return protovalidator.FieldError1("StreamJobSchedule", "the value of field 'started' must be greater than or equal to '0'", protovalidator.Int64ToString(this.Started))
	}
	return nil
}

func (this *StreamJobSchedule) _xxx_xxx_Validator_Validate_ended() error {
	if !(this.Ended >= 0) {
		return protovalidator.FieldError1("StreamJobSchedule", "the value of field 'ended' must be greater than or equal to '0'", protovalidator.Int64ToString(this.Ended))
	}
	return nil
}

var _xxx_xxx_Validator_StreamJobSchedule_InEnums_ConcurrencyPolicy = map[StreamJobSchedule_ConcurrencyPolicy]bool{0: true, 1: true, 2: true, 3: true}

func (this *StreamJobSchedule) _xxx_xxx_Validator_Validate_concurrency_policy() error {
	if !(this.ConcurrencyPolicy > 0) {
		return protovalidator.FieldError1("StreamJobSchedule", "the value of field 'concurrency_policy' must be greater than '0'", protovalidator.Int32ToString(int32(this.ConcurrencyPolicy)))
	}
	if !(_xxx_xxx_Validator_StreamJobSchedule_InEnums_ConcurrencyPolicy[this.ConcurrencyPolicy]) {
		return protovalidator.FieldError1("StreamJobSchedule", "the value of field 'concurrency_policy' must in enums of '[0 1 2 3]'", protovalidator.Int32ToString(int32(this.ConcurrencyPolicy)))
	}
	return nil
}

func (this *StreamJobSchedule) _xxx_xxx_Validator_CheckIf_period_type() bool {
	if !(this.SchedulePolicy == 1) {
		return false
	}
	return true
}

var _xxx_xxx_Validator_StreamJobSchedule_In_PeriodType = map[string]bool{"minute": true, "hour": true, "day": true, "week": true, "month": true, "year": true}

func (this *StreamJobSchedule) _xxx_xxx_Validator_Validate_period_type() error {
	if !this._xxx_xxx_Validator_CheckIf_period_type() {
		return nil
	}
	if !(_xxx_xxx_Validator_StreamJobSchedule_In_PeriodType[this.PeriodType]) {
		return protovalidator.FieldError1("StreamJobSchedule", "the value of field 'period_type' must be one of '[minute hour day week month year]'", this.PeriodType)
	}
	return nil
}

func (this *StreamJobSchedule) _xxx_xxx_Validator_CheckIf_express() bool {
	if !(this.SchedulePolicy == 1) {
		return false
	}
	return true
}

func (this *StreamJobSchedule) _xxx_xxx_Validator_Validate_express() error {
	if !this._xxx_xxx_Validator_CheckIf_express() {
		return nil
	}
	if !(protovalidator.StringIsUnixCron(this.Express)) {
		return protovalidator.FieldError1("StreamJobSchedule", "the value of field 'express' must be a valid standard UNIX-Style crontab expression", this.Express)
	}
	return nil
}

func (this *StreamJobSchedule) _xxx_xxx_Validator_Validate_timeout() error {
	if !(this.Timeout >= 0) {
		return protovalidator.FieldError1("StreamJobSchedule", "the value of field 'timeout' must be greater than or equal to '0'", protovalidator.Int32ToString(this.Timeout))
	}
	if !(this.Timeout <= 100) {
		return protovalidator.FieldError1("StreamJobSchedule", "the value of field 'timeout' must be less than or equal to '100'", protovalidator.Int32ToString(this.Timeout))
	}
	return nil
}

var _xxx_xxx_Validator_StreamJobSchedule_InEnums_RetryPolicy = map[StreamJobSchedule_RetryPolicy]bool{0: true, 1: true, 2: true}

func (this *StreamJobSchedule) _xxx_xxx_Validator_Validate_retry_policy() error {
	if !(this.RetryPolicy > 0) {
		return protovalidator.FieldError1("StreamJobSchedule", "the value of field 'retry_policy' must be greater than '0'", protovalidator.Int32ToString(int32(this.RetryPolicy)))
	}
	if !(_xxx_xxx_Validator_StreamJobSchedule_InEnums_RetryPolicy[this.RetryPolicy]) {
		return protovalidator.FieldError1("StreamJobSchedule", "the value of field 'retry_policy' must in enums of '[0 1 2]'", protovalidator.Int32ToString(int32(this.RetryPolicy)))
	}
	return nil
}

func (this *StreamJobSchedule) _xxx_xxx_Validator_CheckIf_retry_limit() bool {
	if !(this.RetryPolicy == 2) {
		return false
	}
	return true
}

func (this *StreamJobSchedule) _xxx_xxx_Validator_Validate_retry_limit() error {
	if !this._xxx_xxx_Validator_CheckIf_retry_limit() {
		return nil
	}
	if !(this.RetryLimit >= 0) {
		return protovalidator.FieldError1("StreamJobSchedule", "the value of field 'retry_limit' must be greater than or equal to '0'", protovalidator.Int32ToString(this.RetryLimit))
	}
	if !(this.RetryLimit <= 100) {
		return protovalidator.FieldError1("StreamJobSchedule", "the value of field 'retry_limit' must be less than or equal to '100'", protovalidator.Int32ToString(this.RetryLimit))
	}
	return nil
}

func (this *StreamJobSchedule) _xxx_xxx_Validator_CheckIf_retry_interval() bool {
	if !(this.RetryPolicy == 2) {
		return false
	}
	return true
}

func (this *StreamJobSchedule) _xxx_xxx_Validator_Validate_retry_interval() error {
	if !this._xxx_xxx_Validator_CheckIf_retry_interval() {
		return nil
	}
	if !(this.RetryInterval >= 1) {
		return protovalidator.FieldError1("StreamJobSchedule", "the value of field 'retry_interval' must be greater than or equal to '1'", protovalidator.Int32ToString(this.RetryInterval))
	}
	if !(this.RetryInterval <= 30) {
		return protovalidator.FieldError1("StreamJobSchedule", "the value of field 'retry_interval' must be less than or equal to '30'", protovalidator.Int32ToString(this.RetryInterval))
	}
	return nil
}

// Set default value for message model.StreamJobSchedule
func (this *StreamJobSchedule) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_schedule_policy(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_executed(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_started(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_ended(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_concurrency_policy(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_period_type(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_express(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_timeout(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_retry_policy(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_retry_limit(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_retry_interval(); err != nil {
		return err
	}
	return nil
}

func (this *StreamJobRelease) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("StreamJobRelease", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("StreamJobRelease", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

func (this *StreamJobRelease) _xxx_xxx_Validator_Validate_id() error {
	if !(len(this.Id) == 20) {
		return protovalidator.FieldError1("StreamJobRelease", "the byte length of field 'id' must be equal to '20'", protovalidator.StringByteLenToString(this.Id))
	}
	if !(strings.HasPrefix(this.Id, "stj-")) {
		return protovalidator.FieldError1("StreamJobRelease", "the value of field 'id' must start with string 'stj-'", this.Id)
	}
	return nil
}

func (this *StreamJobRelease) _xxx_xxx_Validator_Validate_version() error {
	if !(len(this.Version) == 16) {
		return protovalidator.FieldError1("StreamJobRelease", "the byte length of field 'version' must be equal to '16'", protovalidator.StringByteLenToString(this.Version))
	}
	return nil
}

func (this *StreamJobRelease) _xxx_xxx_Validator_Validate_name() error {
	if !(len(this.Name) >= 2) {
		return protovalidator.FieldError1("StreamJobRelease", "the byte length of field 'name' must be greater than or equal to '2'", protovalidator.StringByteLenToString(this.Name))
	}
	if !(len(this.Name) <= 128) {
		return protovalidator.FieldError1("StreamJobRelease", "the byte length of field 'name' must be less than or equal to '128'", protovalidator.StringByteLenToString(this.Name))
	}
	return nil
}

var _xxx_xxx_Validator_StreamJobRelease_InEnums_Type = map[StreamJob_Type]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true}

func (this *StreamJobRelease) _xxx_xxx_Validator_Validate_type() error {
	if !(this.Type > 0) {
		return protovalidator.FieldError1("StreamJobRelease", "the value of field 'type' must be greater than '0'", protovalidator.Int32ToString(int32(this.Type)))
	}
	if !(_xxx_xxx_Validator_StreamJobRelease_InEnums_Type[this.Type]) {
		return protovalidator.FieldError1("StreamJobRelease", "the value of field 'type' must in enums of '[0 1 2 3 4 5]'", protovalidator.Int32ToString(int32(this.Type)))
	}
	return nil
}

func (this *StreamJobRelease) _xxx_xxx_Validator_Validate_created_by() error {
	if !(len(this.CreatedBy) > 0) {
		return protovalidator.FieldError1("StreamJobRelease", "the byte length of field 'created_by' must be greater than '0'", protovalidator.StringByteLenToString(this.CreatedBy))
	}
	if !(len(this.CreatedBy) < 65) {
		return protovalidator.FieldError1("StreamJobRelease", "the byte length of field 'created_by' must be less than '65'", protovalidator.StringByteLenToString(this.CreatedBy))
	}
	return nil
}

func (this *StreamJobRelease) _xxx_xxx_Validator_Validate_created() error {
	if !(this.Created > 0) {
		return protovalidator.FieldError1("StreamJobRelease", "the value of field 'created' must be greater than '0'", protovalidator.Int64ToString(this.Created))
	}
	return nil
}

func (this *StreamJobRelease) _xxx_xxx_Validator_Validate_updated() error {
	if !(this.Updated > 0) {
		return protovalidator.FieldError1("StreamJobRelease", "the value of field 'updated' must be greater than '0'", protovalidator.Int64ToString(this.Updated))
	}
	return nil
}

// Set default value for message model.StreamJobRelease
func (this *StreamJobRelease) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_space_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_version(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_name(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_type(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_created_by(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_created(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_updated(); err != nil {
		return err
	}
	return nil
}
