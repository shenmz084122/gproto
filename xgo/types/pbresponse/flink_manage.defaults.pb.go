// Code generated by protoc-gen-godefaults. DO NOT EDIT.
// versions:
// 		protoc-gen-godefaults 0.0.2
// source: proto/types/response/flink_manage.proto

package pbresponse

import (
	_ "github.com/DataWorkbench/gproto/xgo/types/pbmodel"
)

// Set default value for message response.DescribeFlinkClusterAPI
func (this *DescribeFlinkClusterAPI) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message response.ListAvailableFlinkVersions
func (this *ListAvailableFlinkVersions) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message response.ListFlinkClusters
func (this *ListFlinkClusters) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message response.CreateFlinkCluster
func (this *CreateFlinkCluster) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message response.DescribeFlinkCluster
func (this *DescribeFlinkCluster) SetDefaults() {
	if this == nil {
		return
	}
	if this.Info != nil {
		if dt, ok := interface{}(this.Info).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}
