// Code generated by protoc-gen-godefaults. DO NOT EDIT.
// versions:
// 		protoc-gen-godefaults 0.0.2
// source: proto/types/request/udf_manage.proto

package pbrequest

import (
	_ "github.com/DataWorkbench/gproto/xgo/types/pbmodel"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbdefaults"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
)

// Set default value for message request.ListUDFs
func (this *ListUDFs) SetDefaults() {
	if this == nil {
		return
	}
	if this.Limit == 0 {
		this.Limit = 100
	}
	return
}

// Set default value for message request.DeleteUDFs
func (this *DeleteUDFs) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.CreateUDF
func (this *CreateUDF) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.UpdateUDF
func (this *UpdateUDF) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.DescribeUDF
func (this *DescribeUDF) SetDefaults() {
	if this == nil {
		return
	}
	return
}
