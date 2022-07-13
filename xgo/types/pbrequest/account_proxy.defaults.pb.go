// Code generated by protoc-gen-godefaults. DO NOT EDIT.
// versions:
// 		protoc-gen-godefaults 0.0.2
// source: proto/types/request/account_proxy.proto

package pbrequest

import (
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbdefaults"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
)

// Set default value for message request.DescribeAccessKeyByProxy
func (this *DescribeAccessKeyByProxy) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.ListUsersByProxy
func (this *ListUsersByProxy) SetDefaults() {
	if this == nil {
		return
	}
	if this.Limit == 0 {
		this.Limit = 100
	}
	return
}

// Set default value for message request.ListNotificationsByProxy
func (this *ListNotificationsByProxy) SetDefaults() {
	if this == nil {
		return
	}
	if this.Limit == 0 {
		this.Limit = 100
	}
	return
}