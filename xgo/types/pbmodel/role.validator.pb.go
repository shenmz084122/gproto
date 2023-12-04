// Code generated by protoc-gen-govalidator. DO NOT EDIT.
// versions:
// 		protoc-gen-govalidator 0.0.1
// source: proto/types/model/role.proto

package pbmodel

import (
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
	protovalidator "github.com/yu31/protoc-plugin/xgo/pkg/protovalidator"
	strconv "strconv"
	strings "strings"
	utf8 "unicode/utf8"
)

func (this *CustomRole) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("CustomRole", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("CustomRole", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

func (this *CustomRole) _xxx_xxx_Validator_Validate_id() error {
	if !(len(this.Id) == 20) {
		return protovalidator.FieldError1("CustomRole", "the byte length of field 'id' must be equal to '20'", protovalidator.StringByteLenToString(this.Id))
	}
	return nil
}

func (this *CustomRole) _xxx_xxx_Validator_Validate_name() error {
	if !(len(this.Name) >= 1) {
		return protovalidator.FieldError1("CustomRole", "the byte length of field 'name' must be greater than or equal to '1'", protovalidator.StringByteLenToString(this.Name))
	}
	if !(len(this.Name) <= 128) {
		return protovalidator.FieldError1("CustomRole", "the byte length of field 'name' must be less than or equal to '128'", protovalidator.StringByteLenToString(this.Name))
	}
	return nil
}

func (this *CustomRole) _xxx_xxx_Validator_Validate_desc() error {
	if !(utf8.RuneCountInString(this.Desc) <= 1024) {
		return protovalidator.FieldError1("CustomRole", "the character length of field 'desc' must be less than or equal to '1024'", protovalidator.StringCharsetLenToString(this.Desc))
	}
	return nil
}

var _xxx_xxx_Validator_CustomRole_InEnums_Status = map[CustomRole_Status]bool{0: true, 1: true, 2: true}

func (this *CustomRole) _xxx_xxx_Validator_Validate_status() error {
	if !(this.Status > 0) {
		return protovalidator.FieldError1("CustomRole", "the value of field 'status' must be greater than '0'", protovalidator.Int32ToString(int32(this.Status)))
	}
	if !(_xxx_xxx_Validator_CustomRole_InEnums_Status[this.Status]) {
		return protovalidator.FieldError1("CustomRole", "the value of field 'status' must in enums of '[0 1 2]'", protovalidator.Int32ToString(int32(this.Status)))
	}
	return nil
}

func (this *CustomRole) _xxx_xxx_Validator_Validate_module_ids() error {
	if !(len(this.ModuleIds) <= 1024) {
		return protovalidator.FieldError1("CustomRole", "the byte length of field 'module_ids' must be less than or equal to '1024'", protovalidator.StringByteLenToString(this.ModuleIds))
	}
	return nil
}

func (this *CustomRole) _xxx_xxx_Validator_CheckIf_perm_type() bool {
	if !(this.ModuleIds != "") {
		return false
	}
	return true
}

var _xxx_xxx_Validator_CustomRole_InEnums_PermType = map[ProjectAPI_PermType]bool{0: true, 1: true, 2: true}

func (this *CustomRole) _xxx_xxx_Validator_Validate_perm_type() error {
	if !this._xxx_xxx_Validator_CheckIf_perm_type() {
		return nil
	}
	if !(this.PermType > 0) {
		return protovalidator.FieldError1("CustomRole", "the value of field 'perm_type' must be greater than '0'", protovalidator.Int32ToString(int32(this.PermType)))
	}
	if !(_xxx_xxx_Validator_CustomRole_InEnums_PermType[this.PermType]) {
		return protovalidator.FieldError1("CustomRole", "the value of field 'perm_type' must in enums of '[0 1 2]'", protovalidator.Int32ToString(int32(this.PermType)))
	}
	return nil
}

// Set default value for message model.CustomRole
func (this *CustomRole) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_space_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_name(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_desc(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_status(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_module_ids(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_perm_type(); err != nil {
		return err
	}
	return nil
}

func (this *SystemRole) _xxx_xxx_Validator_Validate_id() error {
	if !(len(this.Id) == 20) {
		return protovalidator.FieldError1("SystemRole", "the byte length of field 'id' must be equal to '20'", protovalidator.StringByteLenToString(this.Id))
	}
	if !(strings.HasPrefix(this.Id, "ros-")) {
		return protovalidator.FieldError1("SystemRole", "the value of field 'id' must start with string 'ros-'", this.Id)
	}
	return nil
}

var _xxx_xxx_Validator_SystemRole_InEnums_Type = map[SystemRole_Type]bool{0: true, 1: true, 2: true, 3: true, 4: true}

func (this *SystemRole) _xxx_xxx_Validator_Validate_type() error {
	if !(this.Type > 0) {
		return protovalidator.FieldError1("SystemRole", "the value of field 'type' must be greater than '0'", protovalidator.Int32ToString(int32(this.Type)))
	}
	if !(_xxx_xxx_Validator_SystemRole_InEnums_Type[this.Type]) {
		return protovalidator.FieldError1("SystemRole", "the value of field 'type' must in enums of '[0 1 2 3 4]'", protovalidator.Int32ToString(int32(this.Type)))
	}
	return nil
}

func (this *SystemRole) _xxx_xxx_Validator_Validate_name() error {
	if !(this.Name != "") {
		return protovalidator.FieldError1("SystemRole", "the value of field 'name' must be not equal to ''", this.Name)
	}
	return nil
}

// Set default value for message model.SystemRole
func (this *SystemRole) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_type(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_name(); err != nil {
		return err
	}
	return nil
}

func (this *ProjectModule) _xxx_xxx_Validator_Validate_id() error {
	if !(len(this.Id) == 20) {
		return protovalidator.FieldError1("ProjectModule", "the byte length of field 'id' must be equal to '20'", protovalidator.StringByteLenToString(this.Id))
	}
	if !(strings.HasPrefix(this.Id, "pmo-")) {
		return protovalidator.FieldError1("ProjectModule", "the value of field 'id' must start with string 'pmo-'", this.Id)
	}
	return nil
}

var _xxx_xxx_Validator_ProjectModule_InEnums_Classify = map[ProjectModule_Classify]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 101: true, 102: true, 103: true, 104: true, 105: true, 106: true, 107: true, 108: true, 109: true, 110: true, 111: true, 112: true, 113: true, 114: true, 115: true, 116: true, 117: true, 118: true, 119: true, 120: true, 220: true, 221: true, 222: true}

func (this *ProjectModule) _xxx_xxx_Validator_Validate_classify() error {
	if !(this.Classify > 0) {
		return protovalidator.FieldError1("ProjectModule", "the value of field 'classify' must be greater than '0'", protovalidator.Int32ToString(int32(this.Classify)))
	}
	if !(_xxx_xxx_Validator_ProjectModule_InEnums_Classify[this.Classify]) {
		return protovalidator.FieldError1("ProjectModule", "the value of field 'classify' must in enums of '[0 1 2 3 4 5 101 102 103 104 105 106 107 108 109 110 111 112 113 114 115 116 117 118 119 120 220 221 222]'", protovalidator.Int32ToString(int32(this.Classify)))
	}
	return nil
}

func (this *ProjectModule) _xxx_xxx_Validator_Validate_name() error {
	if !(this.Name != "") {
		return protovalidator.FieldError1("ProjectModule", "the value of field 'name' must be not equal to ''", this.Name)
	}
	return nil
}

func (this *ProjectModule) _xxx_xxx_Validator_Validate_api_lists() error {
	for _, item := range this.ApiLists {
		_ = item // To avoid unused panics.
		if dt, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := dt.Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}

// Set default value for message model.ProjectModule
func (this *ProjectModule) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_classify(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_name(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_api_lists(); err != nil {
		return err
	}
	return nil
}

func (this *ProjectAPI) _xxx_xxx_Validator_Validate_api_name() error {
	if !(this.ApiName != "") {
		return protovalidator.FieldError1("ProjectAPI", "the value of field 'api_name' must be not equal to ''", this.ApiName)
	}
	return nil
}

func (this *ProjectAPI) _xxx_xxx_Validator_Validate_display_name() error {
	if !(this.DisplayName != "") {
		return protovalidator.FieldError1("ProjectAPI", "the value of field 'display_name' must be not equal to ''", this.DisplayName)
	}
	return nil
}

var _xxx_xxx_Validator_ProjectAPI_InEnums_PermType = map[ProjectAPI_PermType]bool{0: true, 1: true, 2: true}

func (this *ProjectAPI) _xxx_xxx_Validator_Validate_perm_type() error {
	if !(this.PermType > 0) {
		return protovalidator.FieldError1("ProjectAPI", "the value of field 'perm_type' must be greater than '0'", protovalidator.Int32ToString(int32(this.PermType)))
	}
	if !(_xxx_xxx_Validator_ProjectAPI_InEnums_PermType[this.PermType]) {
		return protovalidator.FieldError1("ProjectAPI", "the value of field 'perm_type' must in enums of '[0 1 2]'", protovalidator.Int32ToString(int32(this.PermType)))
	}
	return nil
}

func (this *ProjectAPI) _xxx_xxx_Validator_Validate_permissions() error {
	if !(len(this.Permissions) > 0) {
		return protovalidator.FieldError1("ProjectAPI", "the length of field 'permissions' must be greater than '0'", strconv.Itoa(len(this.Permissions)))
	}
	for _, item := range this.Permissions {
		_ = item // To avoid unused panics.
		if dt, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := dt.Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (this *ProjectAPI) _xxx_xxx_Validator_Validate_system_roles() error {
	if !(len(this.SystemRoles) > 0) {
		return protovalidator.FieldError1("ProjectAPI", "the length of field 'system_roles' must be greater than '0'", strconv.Itoa(len(this.SystemRoles)))
	}
	return nil
}

// Set default value for message model.ProjectAPI
func (this *ProjectAPI) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_api_name(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_display_name(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_perm_type(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_permissions(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_system_roles(); err != nil {
		return err
	}
	return nil
}

func (this *ProjectAPI_Permission) _xxx_xxx_Validator_Validate_system_role() error {
	if dt, ok := interface{}(this.SystemRole).(interface{ Validate() error }); ok {
		if err := dt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Set default value for message model.ProjectAPI.Permission
func (this *ProjectAPI_Permission) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_system_role(); err != nil {
		return err
	}
	return nil
}
