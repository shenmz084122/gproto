// Code generated by protoc-gen-govalidator. DO NOT EDIT.
// versions:
// 		protoc-gen-govalidator 0.0.1
// source: proto/types/request/space_manage.proto

package pbrequest

import (
	pbmodel "github.com/DataWorkbench/gproto/xgo/types/pbmodel"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbdefaults"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
	protovalidator "github.com/yu31/protoc-plugin/xgo/pkg/protovalidator"
	strconv "strconv"
	strings "strings"
	utf8 "unicode/utf8"
)

func (this *ListWorkspaces) _xxx_xxx_Validator_Validate_user_id() error {
	if !(len(this.UserId) > 1) {
		return protovalidator.FieldError1("ListWorkspaces", "the byte length of field 'user_id' must be greater than '1'", protovalidator.StringByteLenToString(this.UserId))
	}
	if !(len(this.UserId) < 65) {
		return protovalidator.FieldError1("ListWorkspaces", "the byte length of field 'user_id' must be less than '65'", protovalidator.StringByteLenToString(this.UserId))
	}
	return nil
}

func (this *ListWorkspaces) _xxx_xxx_Validator_Validate_limit() error {
	if !(this.Limit > 0) {
		return protovalidator.FieldError1("ListWorkspaces", "the value of field 'limit' must be greater than '0'", protovalidator.Int32ToString(this.Limit))
	}
	if !(this.Limit <= 100) {
		return protovalidator.FieldError1("ListWorkspaces", "the value of field 'limit' must be less than or equal to '100'", protovalidator.Int32ToString(this.Limit))
	}
	return nil
}

func (this *ListWorkspaces) _xxx_xxx_Validator_Validate_offset() error {
	if !(this.Offset >= 0) {
		return protovalidator.FieldError1("ListWorkspaces", "the value of field 'offset' must be greater than or equal to '0'", protovalidator.Int32ToString(this.Offset))
	}
	return nil
}

var _xxx_xxx_Validator_ListWorkspaces_In_SortBy = map[string]bool{"": true, "id": true, "name": true, "created": true, "updated": true}

func (this *ListWorkspaces) _xxx_xxx_Validator_Validate_sort_by() error {
	if !(_xxx_xxx_Validator_ListWorkspaces_In_SortBy[this.SortBy]) {
		return protovalidator.FieldError1("ListWorkspaces", "the value of field 'sort_by' must be one of '[ id name created updated]'", this.SortBy)
	}
	return nil
}

var _xxx_xxx_Validator_ListWorkspaces_InEnums_Status = map[pbmodel.Workspace_Status]bool{0: true, 1: true, 2: true, 3: true}

func (this *ListWorkspaces) _xxx_xxx_Validator_Validate_status() error {
	if !(_xxx_xxx_Validator_ListWorkspaces_InEnums_Status[this.Status]) {
		return protovalidator.FieldError1("ListWorkspaces", "the value of field 'status' must in enums of '[0 1 2 3]'", protovalidator.Int32ToString(int32(this.Status)))
	}
	return nil
}

// Set default value for message request.ListWorkspaces
func (this *ListWorkspaces) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_user_id(); err != nil {
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
	if err := this._xxx_xxx_Validator_Validate_status(); err != nil {
		return err
	}
	return nil
}

func (this *DeleteWorkspaces) _xxx_xxx_Validator_Validate_space_ids() error {
	if !(len(this.SpaceIds) > 0) {
		return protovalidator.FieldError1("DeleteWorkspaces", "the length of field 'space_ids' must be greater than '0'", strconv.Itoa(len(this.SpaceIds)))
	}
	if !(len(this.SpaceIds) <= 100) {
		return protovalidator.FieldError1("DeleteWorkspaces", "the length of field 'space_ids' must be less than or equal to '100'", strconv.Itoa(len(this.SpaceIds)))
	}
	for _, item := range this.SpaceIds {
		_ = item // To avoid unused panics.
		if !(strings.HasPrefix(item, "wks-")) {
			return protovalidator.FieldError1("DeleteWorkspaces", "the value of array item where in field 'space_ids' must start with string 'wks-'", item)
		}
	}
	return nil
}

func (this *DeleteWorkspaces) _xxx_xxx_Validator_Validate_req_user_id() error {
	if !(len(this.ReqUserId) > 1) {
		return protovalidator.FieldError1("DeleteWorkspaces", "the byte length of field 'req_user_id' must be greater than '1'", protovalidator.StringByteLenToString(this.ReqUserId))
	}
	if !(len(this.ReqUserId) < 65) {
		return protovalidator.FieldError1("DeleteWorkspaces", "the byte length of field 'req_user_id' must be less than '65'", protovalidator.StringByteLenToString(this.ReqUserId))
	}
	return nil
}

// Set default value for message request.DeleteWorkspaces
func (this *DeleteWorkspaces) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_space_ids(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_req_user_id(); err != nil {
		return err
	}
	return nil
}

func (this *DisableWorkspaces) _xxx_xxx_Validator_Validate_space_ids() error {
	if !(len(this.SpaceIds) > 0) {
		return protovalidator.FieldError1("DisableWorkspaces", "the length of field 'space_ids' must be greater than '0'", strconv.Itoa(len(this.SpaceIds)))
	}
	if !(len(this.SpaceIds) <= 100) {
		return protovalidator.FieldError1("DisableWorkspaces", "the length of field 'space_ids' must be less than or equal to '100'", strconv.Itoa(len(this.SpaceIds)))
	}
	for _, item := range this.SpaceIds {
		_ = item // To avoid unused panics.
		if !(strings.HasPrefix(item, "wks-")) {
			return protovalidator.FieldError1("DisableWorkspaces", "the value of array item where in field 'space_ids' must start with string 'wks-'", item)
		}
	}
	return nil
}

func (this *DisableWorkspaces) _xxx_xxx_Validator_Validate_req_user_id() error {
	if !(len(this.ReqUserId) > 1) {
		return protovalidator.FieldError1("DisableWorkspaces", "the byte length of field 'req_user_id' must be greater than '1'", protovalidator.StringByteLenToString(this.ReqUserId))
	}
	if !(len(this.ReqUserId) < 65) {
		return protovalidator.FieldError1("DisableWorkspaces", "the byte length of field 'req_user_id' must be less than '65'", protovalidator.StringByteLenToString(this.ReqUserId))
	}
	return nil
}

// Set default value for message request.DisableWorkspaces
func (this *DisableWorkspaces) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_space_ids(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_req_user_id(); err != nil {
		return err
	}
	return nil
}

func (this *EnableWorkspaces) _xxx_xxx_Validator_Validate_space_ids() error {
	if !(len(this.SpaceIds) > 0) {
		return protovalidator.FieldError1("EnableWorkspaces", "the length of field 'space_ids' must be greater than '0'", strconv.Itoa(len(this.SpaceIds)))
	}
	if !(len(this.SpaceIds) <= 100) {
		return protovalidator.FieldError1("EnableWorkspaces", "the length of field 'space_ids' must be less than or equal to '100'", strconv.Itoa(len(this.SpaceIds)))
	}
	for _, item := range this.SpaceIds {
		_ = item // To avoid unused panics.
		if !(strings.HasPrefix(item, "wks-")) {
			return protovalidator.FieldError1("EnableWorkspaces", "the value of array item where in field 'space_ids' must start with string 'wks-'", item)
		}
	}
	return nil
}

func (this *EnableWorkspaces) _xxx_xxx_Validator_Validate_req_user_id() error {
	if !(len(this.ReqUserId) > 1) {
		return protovalidator.FieldError1("EnableWorkspaces", "the byte length of field 'req_user_id' must be greater than '1'", protovalidator.StringByteLenToString(this.ReqUserId))
	}
	if !(len(this.ReqUserId) < 65) {
		return protovalidator.FieldError1("EnableWorkspaces", "the byte length of field 'req_user_id' must be less than '65'", protovalidator.StringByteLenToString(this.ReqUserId))
	}
	return nil
}

// Set default value for message request.EnableWorkspaces
func (this *EnableWorkspaces) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_space_ids(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_req_user_id(); err != nil {
		return err
	}
	return nil
}

func (this *CreateWorkspace) _xxx_xxx_Validator_Validate_owner() error {
	if !(len(this.Owner) > 1) {
		return protovalidator.FieldError1("CreateWorkspace", "the byte length of field 'owner' must be greater than '1'", protovalidator.StringByteLenToString(this.Owner))
	}
	if !(len(this.Owner) < 65) {
		return protovalidator.FieldError1("CreateWorkspace", "the byte length of field 'owner' must be less than '65'", protovalidator.StringByteLenToString(this.Owner))
	}
	return nil
}

func (this *CreateWorkspace) _xxx_xxx_Validator_Validate_name() error {
	if !(utf8.RuneCountInString(this.Name) >= 2) {
		return protovalidator.FieldError1("CreateWorkspace", "the character length of field 'name' must be greater than or equal to '2'", protovalidator.StringCharsetLenToString(this.Name))
	}
	if !(utf8.RuneCountInString(this.Name) <= 128) {
		return protovalidator.FieldError1("CreateWorkspace", "the character length of field 'name' must be less than or equal to '128'", protovalidator.StringCharsetLenToString(this.Name))
	}
	return nil
}

func (this *CreateWorkspace) _xxx_xxx_Validator_Validate_desc() error {
	if !(utf8.RuneCountInString(this.Desc) <= 1024) {
		return protovalidator.FieldError1("CreateWorkspace", "the character length of field 'desc' must be less than or equal to '1024'", protovalidator.StringCharsetLenToString(this.Desc))
	}
	return nil
}

// Set default value for message request.CreateWorkspace
func (this *CreateWorkspace) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_owner(); err != nil {
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

func (this *UpdateWorkspace) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("UpdateWorkspace", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("UpdateWorkspace", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

func (this *UpdateWorkspace) _xxx_xxx_Validator_Validate_name() error {
	if !(utf8.RuneCountInString(this.Name) >= 2) {
		return protovalidator.FieldError1("UpdateWorkspace", "the character length of field 'name' must be greater than or equal to '2'", protovalidator.StringCharsetLenToString(this.Name))
	}
	if !(utf8.RuneCountInString(this.Name) <= 128) {
		return protovalidator.FieldError1("UpdateWorkspace", "the character length of field 'name' must be less than or equal to '128'", protovalidator.StringCharsetLenToString(this.Name))
	}
	return nil
}

func (this *UpdateWorkspace) _xxx_xxx_Validator_Validate_desc() error {
	if !(utf8.RuneCountInString(this.Desc) <= 1024) {
		return protovalidator.FieldError1("UpdateWorkspace", "the character length of field 'desc' must be less than or equal to '1024'", protovalidator.StringCharsetLenToString(this.Desc))
	}
	return nil
}

// Set default value for message request.UpdateWorkspace
func (this *UpdateWorkspace) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_space_id(); err != nil {
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

func (this *DescribeWorkspace) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("DescribeWorkspace", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("DescribeWorkspace", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

// Set default value for message request.DescribeWorkspace
func (this *DescribeWorkspace) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_space_id(); err != nil {
		return err
	}
	return nil
}

func (this *CheckPermission) _xxx_xxx_Validator_Validate_req_user_id() error {
	if !(len(this.ReqUserId) > 1) {
		return protovalidator.FieldError1("CheckPermission", "the byte length of field 'req_user_id' must be greater than '1'", protovalidator.StringByteLenToString(this.ReqUserId))
	}
	if !(len(this.ReqUserId) < 65) {
		return protovalidator.FieldError1("CheckPermission", "the byte length of field 'req_user_id' must be less than '65'", protovalidator.StringByteLenToString(this.ReqUserId))
	}
	return nil
}

func (this *CheckPermission) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("CheckPermission", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("CheckPermission", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

func (this *CheckPermission) _xxx_xxx_Validator_Validate_module_id() error {
	if !(len(this.ModuleId) == 20) {
		return protovalidator.FieldError1("CheckPermission", "the byte length of field 'module_id' must be equal to '20'", protovalidator.StringByteLenToString(this.ModuleId))
	}
	if !(strings.HasPrefix(this.ModuleId, "pmo-")) {
		return protovalidator.FieldError1("CheckPermission", "the value of field 'module_id' must start with string 'pmo-'", this.ModuleId)
	}
	return nil
}

var _xxx_xxx_Validator_CheckPermission_InEnums_PermType = map[pbmodel.ProjectAPI_PermType]bool{0: true, 1: true, 2: true}

func (this *CheckPermission) _xxx_xxx_Validator_Validate_perm_type() error {
	if !(this.PermType > 0) {
		return protovalidator.FieldError1("CheckPermission", "the value of field 'perm_type' must be greater than '0'", protovalidator.Int32ToString(int32(this.PermType)))
	}
	if !(_xxx_xxx_Validator_CheckPermission_InEnums_PermType[this.PermType]) {
		return protovalidator.FieldError1("CheckPermission", "the value of field 'perm_type' must in enums of '[0 1 2]'", protovalidator.Int32ToString(int32(this.PermType)))
	}
	return nil
}

func (this *CheckPermission) _xxx_xxx_Validator_Validate_api_name() error {
	if !(len(this.ApiName) > 0) {
		return protovalidator.FieldError1("CheckPermission", "the byte length of field 'api_name' must be greater than '0'", protovalidator.StringByteLenToString(this.ApiName))
	}
	return nil
}

func (this *CheckPermission) _xxx_xxx_Validator_Validate_system_roles() error {
	if !(len(this.SystemRoles) > 0) {
		return protovalidator.FieldError1("CheckPermission", "the length of field 'system_roles' must be greater than '0'", strconv.Itoa(len(this.SystemRoles)))
	}
	return nil
}

// Set default value for message request.CheckPermission
func (this *CheckPermission) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_req_user_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_space_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_module_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_perm_type(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_api_name(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_system_roles(); err != nil {
		return err
	}
	return nil
}

func (this *DescribeWorkspaceConfig) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("DescribeWorkspaceConfig", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("DescribeWorkspaceConfig", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

func (this *DescribeWorkspaceConfig) _xxx_xxx_Validator_Validate_space_owner() error {
	if !(this.SpaceOwner != "") {
		return protovalidator.FieldError1("DescribeWorkspaceConfig", "the value of field 'space_owner' must be not equal to ''", this.SpaceOwner)
	}
	return nil
}

// Set default value for message request.DescribeWorkspaceConfig
func (this *DescribeWorkspaceConfig) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_space_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_space_owner(); err != nil {
		return err
	}
	return nil
}

func (this *DescribeNetworkConfig) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("DescribeNetworkConfig", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("DescribeNetworkConfig", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

// Set default value for message request.DescribeNetworkConfig
func (this *DescribeNetworkConfig) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_space_id(); err != nil {
		return err
	}
	return nil
}

func (this *AttachVPCToWorkspace) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("AttachVPCToWorkspace", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("AttachVPCToWorkspace", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

func (this *AttachVPCToWorkspace) _xxx_xxx_Validator_Validate_router_id() error {
	if !(strings.HasPrefix(this.RouterId, "rtr-")) {
		return protovalidator.FieldError1("AttachVPCToWorkspace", "the value of field 'router_id' must start with string 'rtr-'", this.RouterId)
	}
	return nil
}

func (this *AttachVPCToWorkspace) _xxx_xxx_Validator_Validate_vxnet_id() error {
	if !(strings.HasPrefix(this.VxnetId, "vxnet-")) {
		return protovalidator.FieldError1("AttachVPCToWorkspace", "the value of field 'vxnet_id' must start with string 'vxnet-'", this.VxnetId)
	}
	return nil
}

// Set default value for message request.AttachVPCToWorkspace
func (this *AttachVPCToWorkspace) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_space_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_router_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_vxnet_id(); err != nil {
		return err
	}
	return nil
}

func (this *DetachVPCFromWorkspace) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("DetachVPCFromWorkspace", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("DetachVPCFromWorkspace", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

// Set default value for message request.DetachVPCFromWorkspace
func (this *DetachVPCFromWorkspace) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_space_id(); err != nil {
		return err
	}
	return nil
}
