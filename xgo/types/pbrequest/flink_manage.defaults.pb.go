// Code generated by protoc-gen-godefaults. DO NOT EDIT.
// versions:
// 		protoc-gen-godefaults 0.0.2
// source: proto/types/request/flink_manage.proto

package pbrequest

import (
	_ "github.com/DataWorkbench/gproto/xgo/types/pbmodel"
	_ "github.com/DataWorkbench/gproto/xgo/types/pbmodel/pbflink"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbdefaults"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
)

// Set default value for message request.ListAvailableFlinkVersions
func (this *ListAvailableFlinkVersions) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.DescribeFlinkClusterAPI
func (this *DescribeFlinkClusterAPI) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.ListFlinkClusters
func (this *ListFlinkClusters) SetDefaults() {
	if this == nil {
		return
	}
	if this.Limit == 0 {
		this.Limit = 100
	}
	return
}

// Set default value for message request.DeleteFlinkClusters
func (this *DeleteFlinkClusters) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.StartFlinkClusters
func (this *StartFlinkClusters) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.StopFlinkClusters
func (this *StopFlinkClusters) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.RestartFlinkClusters
func (this *RestartFlinkClusters) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.CreateFlinkCluster
func (this *CreateFlinkCluster) SetDefaults() {
	if this == nil {
		return
	}
	if this.HostAliases != nil {
		if dt, ok := interface{}(this.HostAliases).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Config != nil {
		if dt, ok := interface{}(this.Config).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}

// Set default value for message request.UpdateFlinkCluster
func (this *UpdateFlinkCluster) SetDefaults() {
	if this == nil {
		return
	}
	if this.HostAliases != nil {
		if dt, ok := interface{}(this.HostAliases).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Config != nil {
		if dt, ok := interface{}(this.Config).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}

// Set default value for message request.DescribeFlinkCluster
func (this *DescribeFlinkCluster) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.UpdateFlinkClusterStatusByScheduler
func (this *UpdateFlinkClusterStatusByScheduler) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.StopFlinkClusterByScheduler
func (this *StopFlinkClusterByScheduler) SetDefaults() {
	if this == nil {
		return
	}
	return
}
