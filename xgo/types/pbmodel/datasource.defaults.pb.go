// Code generated by protoc-gen-godefaults. DO NOT EDIT.
// versions:
// 		protoc-gen-godefaults 0.0.1
// source: proto/types/model/datasource.proto

package pbmodel

import (
	_ "github.com/DataWorkbench/gproto/xgo/types/pbmodel/pbdatasource"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbgosql"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
)

// Set default value for message model.DataSource
func (this *DataSource) SetDefaults() {
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

// Set default value for message model.DataSource.URL
func (this *DataSource_URL) SetDefaults() {
	if this == nil {
		return
	}
	if this.Mysql != nil {
		if dt, ok := interface{}(this.Mysql).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Postgresql != nil {
		if dt, ok := interface{}(this.Postgresql).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Kafka != nil {
		if dt, ok := interface{}(this.Kafka).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.S3 != nil {
		if dt, ok := interface{}(this.S3).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Clickhouse != nil {
		if dt, ok := interface{}(this.Clickhouse).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Hbase != nil {
		if dt, ok := interface{}(this.Hbase).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Ftp != nil {
		if dt, ok := interface{}(this.Ftp).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Hdfs != nil {
		if dt, ok := interface{}(this.Hdfs).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}

// Set default value for message model.DataSourceConnection
func (this *DataSourceConnection) SetDefaults() {
	if this == nil {
		return
	}
	if this.NetworkInfo != nil {
		if dt, ok := interface{}(this.NetworkInfo).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}

// Set default value for message model.DataSourceKind
func (this *DataSourceKind) SetDefaults() {
	if this == nil {
		return
	}
	return
}