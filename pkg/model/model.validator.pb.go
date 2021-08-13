// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/model.proto

package model

import (
	fmt "fmt"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "github.com/yu31/proto-go-plugin/gosql/pb"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *EmptyStruct) Validate() error {
	return nil
}
func (this *Error) Validate() error {
	return nil
}
func (this *Workspace) Validate() error {
	if !(len(this.Id) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.Id))
	}
	if !(len(this.Owner) < 65) {
		return github_com_mwitkow_go_proto_validators.FieldError("Owner", fmt.Errorf(`value '%v' must have a length smaller than '65'`, this.Owner))
	}
	if !(len(this.Name) > 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length greater than '1'`, this.Name))
	}
	if !(len(this.Name) < 129) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length smaller than '129'`, this.Name))
	}
	if !(len(this.Desc) < 1025) {
		return github_com_mwitkow_go_proto_validators.FieldError("Desc", fmt.Errorf(`value '%v' must have a length smaller than '1025'`, this.Desc))
	}
	if !(this.Created > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Created", fmt.Errorf(`value '%v' must be greater than '0'`, this.Created))
	}
	if !(this.Updated > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Updated", fmt.Errorf(`value '%v' must be greater than '0'`, this.Updated))
	}
	return nil
}
func (this *StreamFlow) Validate() error {
	if !(len(this.SpaceId) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("SpaceId", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.SpaceId))
	}
	if !(len(this.Id) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.Id))
	}
	if !(len(this.Name) > 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length greater than '1'`, this.Name))
	}
	if !(len(this.Name) < 129) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length smaller than '129'`, this.Name))
	}
	if !(len(this.Desc) < 1025) {
		return github_com_mwitkow_go_proto_validators.FieldError("Desc", fmt.Errorf(`value '%v' must have a length smaller than '1025'`, this.Desc))
	}
	if !(this.Created > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Created", fmt.Errorf(`value '%v' must be greater than '0'`, this.Created))
	}
	if !(this.Updated > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Updated", fmt.Errorf(`value '%v' must be greater than '0'`, this.Updated))
	}
	return nil
}
func (this *StreamFlowEnv) Validate() error {
	if !(len(this.EngineId) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("EngineId", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.EngineId))
	}
	if !(this.Parallelism > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Parallelism", fmt.Errorf(`value '%v' must be greater than '0'`, this.Parallelism))
	}
	if !(this.JobCu > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("JobCu", fmt.Errorf(`value '%v' must be greater than '0'`, this.JobCu))
	}
	if !(this.TaskCu > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("TaskCu", fmt.Errorf(`value '%v' must be greater than '0'`, this.TaskCu))
	}
	if !(this.TaskNum > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("TaskNum", fmt.Errorf(`value '%v' must be greater than '0'`, this.TaskNum))
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *StreamFlowSchedule) Validate() error {
	if !(this.Started > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Started", fmt.Errorf(`value '%v' must be greater than '-1'`, this.Started))
	}
	if !(this.Ended > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Ended", fmt.Errorf(`value '%v' must be greater than '-1'`, this.Ended))
	}
	if !(this.RetryLimit > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("RetryLimit", fmt.Errorf(`value '%v' must be greater than '-1'`, this.RetryLimit))
	}
	if !(this.RetryLimit < 100) {
		return github_com_mwitkow_go_proto_validators.FieldError("RetryLimit", fmt.Errorf(`value '%v' must be less than '100'`, this.RetryLimit))
	}
	if !(this.RetryInterval > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("RetryInterval", fmt.Errorf(`value '%v' must be greater than '-1'`, this.RetryInterval))
	}
	if !(this.RetryInterval < 31) {
		return github_com_mwitkow_go_proto_validators.FieldError("RetryInterval", fmt.Errorf(`value '%v' must be less than '31'`, this.RetryInterval))
	}
	if !(this.Timeout > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Timeout", fmt.Errorf(`value '%v' must be greater than '-1'`, this.Timeout))
	}
	if !(this.Timeout < 100) {
		return github_com_mwitkow_go_proto_validators.FieldError("Timeout", fmt.Errorf(`value '%v' must be less than '100'`, this.Timeout))
	}
	return nil
}
func (this *StreamFlowProperty) Validate() error {
	if !(len(this.FlowId) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("FlowId", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.FlowId))
	}
	if !(this.Version > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Version", fmt.Errorf(`value '%v' must be greater than '0'`, this.Version))
	}
	if this.Env != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Env); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Env", err)
		}
	}
	if this.Schedule != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Schedule); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Schedule", err)
		}
	}
	return nil
}
func (this *StreamFlowRelease) Validate() error {
	if !(len(this.SpaceId) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("SpaceId", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.SpaceId))
	}
	if !(len(this.FlowId) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("FlowId", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.FlowId))
	}
	if !(this.Version > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Version", fmt.Errorf(`value '%v' must be greater than '0'`, this.Version))
	}
	if !(len(this.Name) > 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length greater than '1'`, this.Name))
	}
	if !(len(this.Name) < 129) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length smaller than '129'`, this.Name))
	}
	if !(len(this.Desc) < 1025) {
		return github_com_mwitkow_go_proto_validators.FieldError("Desc", fmt.Errorf(`value '%v' must have a length smaller than '1025'`, this.Desc))
	}
	if !(this.Created > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Created", fmt.Errorf(`value '%v' must be greater than '0'`, this.Created))
	}
	if !(this.Updated > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Updated", fmt.Errorf(`value '%v' must be greater than '0'`, this.Updated))
	}
	return nil
}
func (this *StreamFlowInst) Validate() error {
	if !(len(this.SpaceId) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("SpaceId", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.SpaceId))
	}
	if !(len(this.FlowId) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("FlowId", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.FlowId))
	}
	if !(this.Version > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Version", fmt.Errorf(`value '%v' must be greater than '0'`, this.Version))
	}
	if !(len(this.Id) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.Id))
	}
	if !(this.Created > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Created", fmt.Errorf(`value '%v' must be greater than '0'`, this.Created))
	}
	if !(this.Updated > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Updated", fmt.Errorf(`value '%v' must be greater than '0'`, this.Updated))
	}
	return nil
}
func (this *OpAudit) Validate() error {
	if !(len(this.UserId) < 65) {
		return github_com_mwitkow_go_proto_validators.FieldError("UserId", fmt.Errorf(`value '%v' must have a length smaller than '65'`, this.UserId))
	}
	if !(len(this.Action) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Action", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.Action))
	}
	if !(this.Created > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Created", fmt.Errorf(`value '%v' must be greater than '0'`, this.Created))
	}
	return nil
}
func (this *Role) Validate() error {
	if !(len(this.SpaceId) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("SpaceId", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.SpaceId))
	}
	if !(len(this.Id) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.Id))
	}
	if !(len(this.Code) < 64) {
		return github_com_mwitkow_go_proto_validators.FieldError("Code", fmt.Errorf(`value '%v' must have a length smaller than '64'`, this.Code))
	}
	if !(len(this.Name) < 128) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length smaller than '128'`, this.Name))
	}
	return nil
}
func (this *Member) Validate() error {
	if !(len(this.SpaceId) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("SpaceId", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.SpaceId))
	}
	if !(len(this.UserId) < 65) {
		return github_com_mwitkow_go_proto_validators.FieldError("UserId", fmt.Errorf(`value '%v' must have a length smaller than '65'`, this.UserId))
	}
	if !(len(this.RoleIds) < 256) {
		return github_com_mwitkow_go_proto_validators.FieldError("RoleIds", fmt.Errorf(`value '%v' must have a length smaller than '256'`, this.RoleIds))
	}
	if !(this.Created > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Created", fmt.Errorf(`value '%v' must be greater than '0'`, this.Created))
	}
	if !(this.Updated > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Updated", fmt.Errorf(`value '%v' must be greater than '0'`, this.Updated))
	}
	return nil
}
func (this *MonitorRule) Validate() error {
	if !(len(this.SpaceId) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("SpaceId", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.SpaceId))
	}
	if !(len(this.Id) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.Id))
	}
	if !(len(this.Name) > 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length greater than '1'`, this.Name))
	}
	if !(len(this.Name) < 129) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length smaller than '129'`, this.Name))
	}
	if !(this.Unit > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Unit", fmt.Errorf(`value '%v' must be greater than '0'`, this.Unit))
	}
	if !(this.Unit < 3) {
		return github_com_mwitkow_go_proto_validators.FieldError("Unit", fmt.Errorf(`value '%v' must be less than '3'`, this.Unit))
	}
	if !(len(this.Text) > 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Text", fmt.Errorf(`value '%v' must have a length greater than '1'`, this.Text))
	}
	if !(len(this.Text) < 1025) {
		return github_com_mwitkow_go_proto_validators.FieldError("Text", fmt.Errorf(`value '%v' must have a length smaller than '1025'`, this.Text))
	}
	if !(this.Trigger > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Trigger", fmt.Errorf(`value '%v' must be greater than '0'`, this.Trigger))
	}
	if !(this.AlarmTimes > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("AlarmTimes", fmt.Errorf(`value '%v' must be greater than '0'`, this.AlarmTimes))
	}
	if !(this.AlarmTimes < 100) {
		return github_com_mwitkow_go_proto_validators.FieldError("AlarmTimes", fmt.Errorf(`value '%v' must be less than '100'`, this.AlarmTimes))
	}
	if !(this.AlarmInterval > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("AlarmInterval", fmt.Errorf(`value '%v' must be greater than '0'`, this.AlarmInterval))
	}
	if !(this.AlarmInterval < 31) {
		return github_com_mwitkow_go_proto_validators.FieldError("AlarmInterval", fmt.Errorf(`value '%v' must be less than '31'`, this.AlarmInterval))
	}
	if !(len(this.AlarmType) > 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("AlarmType", fmt.Errorf(`value '%v' must have a length greater than '1'`, this.AlarmType))
	}
	if !(len(this.AlarmType) < 33) {
		return github_com_mwitkow_go_proto_validators.FieldError("AlarmType", fmt.Errorf(`value '%v' must have a length smaller than '33'`, this.AlarmType))
	}
	if !(len(this.FreeTime) > 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("FreeTime", fmt.Errorf(`value '%v' must have a length greater than '1'`, this.FreeTime))
	}
	if !(len(this.FreeTime) < 17) {
		return github_com_mwitkow_go_proto_validators.FieldError("FreeTime", fmt.Errorf(`value '%v' must have a length smaller than '17'`, this.FreeTime))
	}
	if !(len(this.Receiver) > 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Receiver", fmt.Errorf(`value '%v' must have a length greater than '1'`, this.Receiver))
	}
	if !(len(this.Receiver) < 257) {
		return github_com_mwitkow_go_proto_validators.FieldError("Receiver", fmt.Errorf(`value '%v' must have a length smaller than '257'`, this.Receiver))
	}
	return nil
}
func (this *QueueMessage) Validate() error {
	if this.Property != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Property); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Property", err)
		}
	}
	return nil
}
func (this *InstanceStatusStat) Validate() error {
	if !(this.State > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("State", fmt.Errorf(`value '%v' must be greater than '-1'`, this.State))
	}
	if !(this.Count > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Count", fmt.Errorf(`value '%v' must be greater than '-1'`, this.Count))
	}
	return nil
}
func (this *InstanceRuntimeRankInfo) Validate() error {
	if !(this.RunningTime > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("RunningTime", fmt.Errorf(`value '%v' must be greater than '-1'`, this.RunningTime))
	}
	return nil
}
func (this *InstanceErrorRankInfo) Validate() error {
	if !(this.ErrorCount > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("ErrorCount", fmt.Errorf(`value '%v' must be greater than '-1'`, this.ErrorCount))
	}
	return nil
}
func (this *DispatchTaskCountInfo) Validate() error {
	if !(this.FlowCount > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("FlowCount", fmt.Errorf(`value '%v' must be greater than '-1'`, this.FlowCount))
	}
	if !(this.InstanceCount > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("InstanceCount", fmt.Errorf(`value '%v' must be greater than '-1'`, this.InstanceCount))
	}
	if !(this.Updated > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Updated", fmt.Errorf(`value '%v' must be greater than '0'`, this.Updated))
	}
	return nil
}
func (this *InstanceTaskExecStat) Validate() error {
	if !(this.Hour > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Hour", fmt.Errorf(`value '%v' must be greater than '-1'`, this.Hour))
	}
	if !(this.Hour < 25) {
		return github_com_mwitkow_go_proto_validators.FieldError("Hour", fmt.Errorf(`value '%v' must be less than '25'`, this.Hour))
	}
	if !(this.InstanceCount > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("InstanceCount", fmt.Errorf(`value '%v' must be greater than '-1'`, this.InstanceCount))
	}
	return nil
}
