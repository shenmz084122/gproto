// Code generated by protoc-gen-godefaults. DO NOT EDIT.
// versions:
// 		protoc-gen-godefaults 0.0.2
// source: proto/types/model/account.proto

package pbmodel

import (
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
)

// Set default value for message model.User
func (this *User) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message model.Role
func (this *Role) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message model.AdminModule
func (this *AdminModule) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message model.AdminAPI
func (this *AdminAPI) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message model.AdminAPI.Permission
func (this *AdminAPI_Permission) SetDefaults() {
	if this == nil {
		return
	}
	if this.Role != nil {
		if dt, ok := interface{}(this.Role).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}
