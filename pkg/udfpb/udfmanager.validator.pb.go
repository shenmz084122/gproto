// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/udfmanager.proto

package udfpb

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

func (this *CreateRequest) Validate() error {
	if !(len(this.ID) < 21) {
		return github_com_mwitkow_go_proto_validators.FieldError("ID", fmt.Errorf(`value '%v' must have a length smaller than '21'`, this.ID))
	}
	if !(len(this.SpaceID) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("SpaceID", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.SpaceID))
	}
	if !(len(this.UdfType) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("UdfType", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.UdfType))
	}
	if !(len(this.UdfType) < 17) {
		return github_com_mwitkow_go_proto_validators.FieldError("UdfType", fmt.Errorf(`value '%v' must have a length smaller than '17'`, this.UdfType))
	}
	if !(len(this.Name) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.Name))
	}
	if !(len(this.Name) < 65) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length smaller than '65'`, this.Name))
	}
	if !(len(this.Comment) < 257) {
		return github_com_mwitkow_go_proto_validators.FieldError("Comment", fmt.Errorf(`value '%v' must have a length smaller than '257'`, this.Comment))
	}
	if !(len(this.Define) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Define", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.Define))
	}
	if !(len(this.UsageSample) > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("UsageSample", fmt.Errorf(`value '%v' must have a length greater than '-1'`, this.UsageSample))
	}
	return nil
}
func (this *UpdateRequest) Validate() error {
	if !(len(this.ID) < 21) {
		return github_com_mwitkow_go_proto_validators.FieldError("ID", fmt.Errorf(`value '%v' must have a length smaller than '21'`, this.ID))
	}
	if !(len(this.UdfType) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("UdfType", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.UdfType))
	}
	if !(len(this.UdfType) < 17) {
		return github_com_mwitkow_go_proto_validators.FieldError("UdfType", fmt.Errorf(`value '%v' must have a length smaller than '17'`, this.UdfType))
	}
	if !(len(this.Name) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.Name))
	}
	if !(len(this.Name) < 65) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length smaller than '65'`, this.Name))
	}
	if !(len(this.Comment) < 257) {
		return github_com_mwitkow_go_proto_validators.FieldError("Comment", fmt.Errorf(`value '%v' must have a length smaller than '257'`, this.Comment))
	}
	if !(len(this.Define) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Define", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.Define))
	}
	if !(len(this.UsageSample) > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("UsageSample", fmt.Errorf(`value '%v' must have a length greater than '-1'`, this.UsageSample))
	}
	return nil
}
func (this *DeleteRequest) Validate() error {
	if !(len(this.ID) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("ID", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.ID))
	}
	return nil
}
func (this *DeleteAllRequest) Validate() error {
	if !(len(this.SpaceID) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("SpaceID", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.SpaceID))
	}
	return nil
}
func (this *InfoReply) Validate() error {
	if !(len(this.ID) < 21) {
		return github_com_mwitkow_go_proto_validators.FieldError("ID", fmt.Errorf(`value '%v' must have a length smaller than '21'`, this.ID))
	}
	if !(len(this.SpaceID) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("SpaceID", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.SpaceID))
	}
	if !(len(this.UdfType) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("UdfType", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.UdfType))
	}
	if !(len(this.UdfType) < 17) {
		return github_com_mwitkow_go_proto_validators.FieldError("UdfType", fmt.Errorf(`value '%v' must have a length smaller than '17'`, this.UdfType))
	}
	if !(len(this.Name) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.Name))
	}
	if !(len(this.Name) < 65) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must have a length smaller than '65'`, this.Name))
	}
	if !(len(this.Comment) < 257) {
		return github_com_mwitkow_go_proto_validators.FieldError("Comment", fmt.Errorf(`value '%v' must have a length smaller than '257'`, this.Comment))
	}
	if !(len(this.Define) > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Define", fmt.Errorf(`value '%v' must have a length greater than '0'`, this.Define))
	}
	if !(len(this.CreateTime) < 65) {
		return github_com_mwitkow_go_proto_validators.FieldError("CreateTime", fmt.Errorf(`value '%v' must have a length smaller than '65'`, this.CreateTime))
	}
	if !(len(this.UpdateTime) < 65) {
		return github_com_mwitkow_go_proto_validators.FieldError("UpdateTime", fmt.Errorf(`value '%v' must have a length smaller than '65'`, this.UpdateTime))
	}
	if !(len(this.UsageSample) > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("UsageSample", fmt.Errorf(`value '%v' must have a length greater than '-1'`, this.UsageSample))
	}
	return nil
}
func (this *DescribeRequest) Validate() error {
	if !(len(this.ID) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("ID", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.ID))
	}
	return nil
}
func (this *ListsRequest) Validate() error {
	if !(this.Limit > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Limit", fmt.Errorf(`value '%v' must be greater than '0'`, this.Limit))
	}
	if !(this.Offset > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Offset", fmt.Errorf(`value '%v' must be greater than '-1'`, this.Offset))
	}
	if !(len(this.SpaceID) == 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("SpaceID", fmt.Errorf(`value '%v' must have a length equal to '20'`, this.SpaceID))
	}
	return nil
}
func (this *ListsReply) Validate() error {
	if !(this.Total > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Total", fmt.Errorf(`value '%v' must be greater than '-1'`, this.Total))
	}
	for _, item := range this.Infos {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Infos", err)
			}
		}
	}
	return nil
}
