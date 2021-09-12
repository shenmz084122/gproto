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
	if !(len(this.Name) > 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length greater than '1'`, this.Name))
	}
	if !(len(this.Name) < 65) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length smaller than '65'`, this.Name))
	}
	if !(len(this.Image) > 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Image", fmt.Errorf(`value '%v' must have a length greater than '1'`, this.Image))
	}
	if !(len(this.Image) < 10485761) {
		return github_com_mwitkow_go_proto_validators.FieldError("Image", fmt.Errorf(`value '%v' must have a length smaller than '10485761'`, this.Image))
	}
	if !(len(this.Desc) > 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Desc", fmt.Errorf(`value '%v' must have a length greater than '1'`, this.Desc))
	}
	if !(len(this.Desc) < 513) {
		return github_com_mwitkow_go_proto_validators.FieldError("Desc", fmt.Errorf(`value '%v' must have a length smaller than '513'`, this.Desc))
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
