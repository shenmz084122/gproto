// Code generated by protoc-gen-godefaults. DO NOT EDIT.
// versions:
// 		protoc-gen-godefaults 0.0.2
// source: proto/types/request/member_manage.proto

package pbrequest

import (
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbdefaults"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
)

// Set default value for message request.ListMembers
func (this *ListMembers) SetDefaults() {
	if this == nil {
		return
	}
	if this.Limit == 0 {
		this.Limit = 100
	}
	return
}

// Set default value for message request.DeleteMembers
func (this *DeleteMembers) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.AddMembers
func (this *AddMembers) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.UpdateMember
func (this *UpdateMember) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.DescribeMember
func (this *DescribeMember) SetDefaults() {
	if this == nil {
		return
	}
	return
}
