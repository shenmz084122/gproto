// Code generated by protoc-gen-godefaults. DO NOT EDIT.
// versions:
// 		protoc-gen-godefaults 0.0.2
// source: proto/types/model/member.proto

package pbmodel

import (
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
)

// Set default value for message model.Member
func (this *Member) SetDefaults() {
	if this == nil {
		return
	}
	if this.UserInfo != nil {
		if dt, ok := interface{}(this.UserInfo).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}
