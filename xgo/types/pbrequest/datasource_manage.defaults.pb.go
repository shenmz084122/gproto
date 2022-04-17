// Code generated by protoc-gen-godefaults. DO NOT EDIT.
// versions:
// 		protoc-gen-godefaults 0.0.2
// source: proto/types/request/datasource_manage.proto

package pbrequest

import (
	_ "github.com/DataWorkbench/gproto/xgo/types/pbmodel"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbdefaults"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
)

// Set default value for message request.ListDataSources
func (this *ListDataSources) SetDefaults() {
	if this == nil {
		return
	}
	if this.Limit == 0 {
		this.Limit = 100
	}
	return
}

// Set default value for message request.CreateDataSource
func (this *CreateDataSource) SetDefaults() {
	if this == nil {
		return
	}
	if this.Url != nil {
		if dt, ok := interface{}(this.Url).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.LastConnection != nil {
		if dt, ok := interface{}(this.LastConnection).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}

// Set default value for message request.UpdateDataSource
func (this *UpdateDataSource) SetDefaults() {
	if this == nil {
		return
	}
	if this.Url != nil {
		if dt, ok := interface{}(this.Url).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}

// Set default value for message request.DeleteDataSources
func (this *DeleteDataSources) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.DescribeDataSource
func (this *DescribeDataSource) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.EnableDataSources
func (this *EnableDataSources) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.DisableDataSources
func (this *DisableDataSources) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.ListDataSourceConnections
func (this *ListDataSourceConnections) SetDefaults() {
	if this == nil {
		return
	}
	if this.Limit == 0 {
		this.Limit = 100
	}
	return
}

// Set default value for message request.PingDataSourceConnection
func (this *PingDataSourceConnection) SetDefaults() {
	if this == nil {
		return
	}
	if this.Url != nil {
		if dt, ok := interface{}(this.Url).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}

// Set default value for message request.DescribeDataSourceTables
func (this *DescribeDataSourceTables) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message request.DescribeDataSourceTableSchema
func (this *DescribeDataSourceTableSchema) SetDefaults() {
	if this == nil {
		return
	}
	return
}
