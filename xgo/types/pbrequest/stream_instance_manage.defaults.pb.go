// Code generated by protoc-gen-godefaults. DO NOT EDIT.
// versions:
// 		protoc-gen-godefaults 0.0.2
// source: proto/types/request/stream_instance_manage.proto

package pbrequest

import (
	_ "github.com/DataWorkbench/gproto/xgo/types/pbmodel"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbdefaults"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
)

// Set default value for message request.ListStreamInstances
func (this *ListStreamInstances) SetDefaults() {
	if this == nil {
		return
	}
	if this.Limit == 0 {
		this.Limit = 100
	}
	return
}

// Set default value for message request.TerminateStreamInstances
func (this *TerminateStreamInstances) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.SuspendStreamInstances
func (this *SuspendStreamInstances) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.ResumeStreamInstances
func (this *ResumeStreamInstances) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.DescribeStreamInstance
func (this *DescribeStreamInstance) SetDefaults() {
	if this == nil {
		return
	}
	return
}
