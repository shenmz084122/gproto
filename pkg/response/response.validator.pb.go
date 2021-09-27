// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/response.proto

package response

import (
	fmt "fmt"
	_ "github.com/DataWorkbench/gproto/pkg/model"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ListWorkspaces) Validate() error {
	for _, item := range this.Infos {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Infos", err)
			}
		}
	}
	return nil
}
func (this *ListAudits) Validate() error {
	for _, item := range this.Infos {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Infos", err)
			}
		}
	}
	return nil
}
func (this *CreateWorkspace) Validate() error {
	if !(len(this.Id) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.Id))
	}
	return nil
}
func (this *DescribeWorkspace) Validate() error {
	if this.Info != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Info); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Info", err)
		}
	}
	return nil
}
func (this *ListMembers) Validate() error {
	for _, item := range this.Infos {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Infos", err)
			}
		}
	}
	return nil
}
func (this *ListSystemRoles) Validate() error {
	for _, item := range this.Infos {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Infos", err)
			}
		}
	}
	return nil
}
func (this *ListStreamFlows) Validate() error {
	for _, item := range this.Infos {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Infos", err)
			}
		}
	}
	return nil
}
func (this *CreateStreamFlow) Validate() error {
	if !(len(this.Id) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.Id))
	}
	return nil
}
func (this *DescribeStreamFlow) Validate() error {
	if this.Info != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Info); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Info", err)
		}
	}
	return nil
}
func (this *GetStreamFlowNode) Validate() error {
	if this.Nodes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Nodes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Nodes", err)
		}
	}
	return nil
}
func (this *GetStreamFlowEnv) Validate() error {
	if this.Info != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Info); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Info", err)
		}
	}
	return nil
}
func (this *GetStreamFlowSchedule) Validate() error {
	if this.Info != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Info); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Info", err)
		}
	}
	return nil
}
func (this *ListReleaseStreamFlows) Validate() error {
	for _, item := range this.Infos {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Infos", err)
			}
		}
	}
	return nil
}
func (this *ListStreamFlowVersions) Validate() error {
	for _, item := range this.Infos {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Infos", err)
			}
		}
	}
	return nil
}
func (this *ListMonitorRules) Validate() error {
	for _, item := range this.Infos {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Infos", err)
			}
		}
	}
	return nil
}
func (this *DescribeMonitorRule) Validate() error {
	if nil == this.Info {
		return github_com_mwitkow_go_proto_validators.FieldError("Info", fmt.Errorf("message must exist"))
	}
	if this.Info != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Info); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Info", err)
		}
	}
	return nil
}
func (this *ListStreamInsts) Validate() error {
	for _, item := range this.Infos {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Infos", err)
			}
		}
	}
	return nil
}
func (this *DescribeSource) Validate() error {
	if this.Info != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Info); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Info", err)
		}
	}
	if !(len(this.Connected) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Connected", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.Connected))
	}
	if !(len(this.Connected) < 17) {
		return github_com_mwitkow_go_proto_validators.FieldError("Connected", fmt.Errorf(`value '%v' must have a length smaller than '17'`, this.Connected))
	}
	return nil
}
func (this *ListSource) Validate() error {
	for _, item := range this.Infos {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Infos", err)
			}
		}
	}
	return nil
}
func (this *SourceKind) Validate() error {
	for _, item := range this.Kinds {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Kinds", err)
			}
		}
	}
	return nil
}
func (this *SourceKind_Kind) Validate() error {
	if !(len(this.Name) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.Name))
	}
	if !(len(this.Name) < 65) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length smaller than '65'`, this.Name))
	}
	return nil
}
func (this *JsonList) Validate() error {
	for _, item := range this.JsonList {
		if !(len(item) > -1) {
			return github_com_mwitkow_go_proto_validators.FieldError("JsonList", fmt.Errorf(`value '%v' must have a length greater than '-1'`, item))
		}
		if !(len(item) < 10485761) {
			return github_com_mwitkow_go_proto_validators.FieldError("JsonList", fmt.Errorf(`value '%v' must have a length smaller than '10485761'`, item))
		}
	}
	return nil
}
func (this *TableColumns) Validate() error {
	for _, item := range this.Columns {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Columns", err)
			}
		}
	}
	return nil
}
func (this *TableColumns_Column) Validate() error {
	if !(len(this.Name) > 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length greater than '1'`, this.Name))
	}
	if !(len(this.Name) < 65) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length smaller than '65'`, this.Name))
	}
	if !(len(this.Type) > 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Type", fmt.Errorf(`value '%v' must have a length greater than '1'`, this.Type))
	}
	if !(len(this.Type) < 65) {
		return github_com_mwitkow_go_proto_validators.FieldError("Type", fmt.Errorf(`value '%v' must have a length smaller than '65'`, this.Type))
	}
	if !(len(this.Length) > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Length", fmt.Errorf(`value '%v' must have a length greater than '-1'`, this.Length))
	}
	if !(len(this.Length) < 65) {
		return github_com_mwitkow_go_proto_validators.FieldError("Length", fmt.Errorf(`value '%v' must have a length smaller than '65'`, this.Length))
	}
	if !(len(this.IsPrkey) < 8) {
		return github_com_mwitkow_go_proto_validators.FieldError("IsPrkey", fmt.Errorf(`value '%v' must have a length smaller than '8'`, this.IsPrkey))
	}
	return nil
}
func (this *DescribeTable) Validate() error {
	if this.Info != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Info); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Info", err)
		}
	}
	return nil
}
func (this *ListTable) Validate() error {
	for _, item := range this.Infos {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Infos", err)
			}
		}
	}
	return nil
}
func (this *ListResources) Validate() error {
	for _, item := range this.Infos {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Infos", err)
			}
		}
	}
	return nil
}
func (this *CreateDir) Validate() error {
	if !(len(this.Id) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.Id))
	}
	return nil
}
func (this *UploadFile) Validate() error {
	if !(len(this.Id) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.Id))
	}
	return nil
}
func (this *DownloadFile) Validate() error {
	return nil
}
func (this *ListUDF) Validate() error {
	for _, item := range this.Infos {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Infos", err)
			}
		}
	}
	return nil
}
func (this *DescribeUDF) Validate() error {
	if this.Info != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Info); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Info", err)
		}
	}
	return nil
}
func (this *JobState) Validate() error {
	return nil
}
func (this *NodeRelations) Validate() error {
	if !(len(this.Relations) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Relations", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.Relations))
	}
	if !(len(this.Relations) < 20000) {
		return github_com_mwitkow_go_proto_validators.FieldError("Relations", fmt.Errorf(`value '%v' must have a length smaller than '20000'`, this.Relations))
	}
	return nil
}
func (this *JobParser) Validate() error {
	if !(len(this.ZeppelinConf) > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("ZeppelinConf", fmt.Errorf(`value '%v' must have a length greater than '-1'`, this.ZeppelinConf))
	}
	if !(len(this.ZeppelinConf) < 40000) {
		return github_com_mwitkow_go_proto_validators.FieldError("ZeppelinConf", fmt.Errorf(`value '%v' must have a length smaller than '40000'`, this.ZeppelinConf))
	}
	if !(len(this.ZeppelinDepends) > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("ZeppelinDepends", fmt.Errorf(`value '%v' must have a length greater than '-1'`, this.ZeppelinDepends))
	}
	if !(len(this.ZeppelinDepends) < 40000) {
		return github_com_mwitkow_go_proto_validators.FieldError("ZeppelinDepends", fmt.Errorf(`value '%v' must have a length smaller than '40000'`, this.ZeppelinDepends))
	}
	if !(len(this.ZeppelinScalaUDF) > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("ZeppelinScalaUDF", fmt.Errorf(`value '%v' must have a length greater than '-1'`, this.ZeppelinScalaUDF))
	}
	if !(len(this.ZeppelinScalaUDF) < 40000) {
		return github_com_mwitkow_go_proto_validators.FieldError("ZeppelinScalaUDF", fmt.Errorf(`value '%v' must have a length smaller than '40000'`, this.ZeppelinScalaUDF))
	}
	if !(len(this.ZeppelinPythonUDF) > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("ZeppelinPythonUDF", fmt.Errorf(`value '%v' must have a length greater than '-1'`, this.ZeppelinPythonUDF))
	}
	if !(len(this.ZeppelinPythonUDF) < 40000) {
		return github_com_mwitkow_go_proto_validators.FieldError("ZeppelinPythonUDF", fmt.Errorf(`value '%v' must have a length smaller than '40000'`, this.ZeppelinPythonUDF))
	}
	if !(len(this.ZeppelinMainRun) > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("ZeppelinMainRun", fmt.Errorf(`value '%v' must have a length greater than '-1'`, this.ZeppelinMainRun))
	}
	if !(len(this.ZeppelinMainRun) < 40000) {
		return github_com_mwitkow_go_proto_validators.FieldError("ZeppelinMainRun", fmt.Errorf(`value '%v' must have a length smaller than '40000'`, this.ZeppelinMainRun))
	}
	if this.Resources != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Resources); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Resources", err)
		}
	}
	if this.Hbase != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Hbase); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Hbase", err)
		}
	}
	if this.S3 != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.S3); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("S3", err)
		}
	}
	return nil
}
func (this *JobFree) Validate() error {
	if !(len(this.ZeppelinDeleteJar) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ZeppelinDeleteJar", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.ZeppelinDeleteJar))
	}
	if !(len(this.ZeppelinDeleteJar) < 20000) {
		return github_com_mwitkow_go_proto_validators.FieldError("ZeppelinDeleteJar", fmt.Errorf(`value '%v' must have a length smaller than '20000'`, this.ZeppelinDeleteJar))
	}
	return nil
}
func (this *PTasksStatusStat) Validate() error {
	for _, item := range this.Infos {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Infos", err)
			}
		}
	}
	return nil
}
func (this *PTasksExecStat) Validate() error {
	for _, item := range this.Today {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Today", err)
			}
		}
	}
	for _, item := range this.Yesterday {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Yesterday", err)
			}
		}
	}
	for _, item := range this.History {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("History", err)
			}
		}
	}
	return nil
}
func (this *PTaskRuntimeRanking) Validate() error {
	for _, item := range this.Infos {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Infos", err)
			}
		}
	}
	if !(this.Total > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Total", fmt.Errorf(`value '%v' must be greater than '-1'`, this.Total))
	}
	return nil
}
func (this *PTaskErrorRanking) Validate() error {
	for _, item := range this.Infos {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Infos", err)
			}
		}
	}
	if !(this.Total > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Total", fmt.Errorf(`value '%v' must be greater than '-1'`, this.Total))
	}
	return nil
}
func (this *PTaskDispatchCount) Validate() error {
	for _, item := range this.Infos {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Infos", err)
			}
		}
	}
	return nil
}
func (this *ZeppelinAddress) Validate() error {
	if !(len(this.ServerAddress) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ServerAddress", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.ServerAddress))
	}
	if !(len(this.ServerAddress) < 1025) {
		return github_com_mwitkow_go_proto_validators.FieldError("ServerAddress", fmt.Errorf(`value '%v' must have a length smaller than '1025'`, this.ServerAddress))
	}
	return nil
}
