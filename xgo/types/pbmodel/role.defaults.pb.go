// Code generated by protoc-gen-godefaults. DO NOT EDIT.
// versions:
// 		protoc-gen-godefaults 0.0.1
// source: proto/types/model/role.proto

package pbmodel

import (
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
)

// Set default value for message model.CustomRole
func (this *CustomRole) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message model.SystemRole
func (this *SystemRole) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message model.ProjectModule
func (this *ProjectModule) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message model.ProjectAPI
func (this *ProjectAPI) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message model.ProjectAPI.Permission
func (this *ProjectAPI_Permission) SetDefaults() {
	if this == nil {
		return
	}
	if this.SystemRole != nil {
		if dt, ok := interface{}(this.SystemRole).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}
