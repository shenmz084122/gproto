// Code generated by protoc-gen-godefaults. DO NOT EDIT.
// versions:
// 		protoc-gen-godefaults v0.1
// source: proto/flink.proto

package flinkpb

import (
	_ "github.com/yu31/proto-go-plugin/pkg/pb/defaultspb"
	_ "github.com/yu31/proto-go-plugin/pkg/pb/gosqlpb"
	_ "github.com/yu31/proto-go-plugin/pkg/pb/validatorpb"
)

// Set default value for message flink.ColumnAs
func (this *ColumnAs) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.DestOperator
func (this *DestOperator) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.ValuesOperator
func (this *ValuesOperator) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.ValuesOperator.ValuesType
func (this *ValuesOperator_ValuesType) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.OrderByOperator
func (this *OrderByOperator) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.OrderByOperator.OrderByColumn
func (this *OrderByOperator_OrderByColumn) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.LimitOperator
func (this *LimitOperator) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.OffsetOperator
func (this *OffsetOperator) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.FetchOperator
func (this *FetchOperator) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.UnionOperator
func (this *UnionOperator) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.GroupByOperator
func (this *GroupByOperator) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.HavingOperator
func (this *HavingOperator) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.ConstOperator
func (this *ConstOperator) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.SourceOperator
func (this *SourceOperator) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.DimensionOperator
func (this *DimensionOperator) SetDefaults() {
	if this == nil {
		return
	}
	if this.TimeColumn != nil {
		if dt, ok := interface{}(this.TimeColumn).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}

// Set default value for message flink.ExceptOperator
func (this *ExceptOperator) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.IntersectOperator
func (this *IntersectOperator) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.FilterOperator
func (this *FilterOperator) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.JoinOperator
func (this *JoinOperator) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.UDTFOperator
func (this *UDTFOperator) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.UDTTFOperator
func (this *UDTTFOperator) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.OperatorProperty
func (this *OperatorProperty) SetDefaults() {
	if this == nil {
		return
	}
	if this.Source != nil {
		if dt, ok := interface{}(this.Source).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Dest != nil {
		if dt, ok := interface{}(this.Dest).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Values != nil {
		if dt, ok := interface{}(this.Values).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.OrderBy != nil {
		if dt, ok := interface{}(this.OrderBy).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Limit != nil {
		if dt, ok := interface{}(this.Limit).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Offset != nil {
		if dt, ok := interface{}(this.Offset).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Fetch != nil {
		if dt, ok := interface{}(this.Fetch).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Union != nil {
		if dt, ok := interface{}(this.Union).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.GroupBy != nil {
		if dt, ok := interface{}(this.GroupBy).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Having != nil {
		if dt, ok := interface{}(this.Having).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Const != nil {
		if dt, ok := interface{}(this.Const).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Dimension != nil {
		if dt, ok := interface{}(this.Dimension).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Except != nil {
		if dt, ok := interface{}(this.Except).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Intersect != nil {
		if dt, ok := interface{}(this.Intersect).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Filter != nil {
		if dt, ok := interface{}(this.Filter).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Udtf != nil {
		if dt, ok := interface{}(this.Udtf).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Udttf != nil {
		if dt, ok := interface{}(this.Udttf).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Join != nil {
		if dt, ok := interface{}(this.Join).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}

// Set default value for message flink.FlinkOperator
func (this *FlinkOperator) SetDefaults() {
	if this == nil {
		return
	}
	if this.Property != nil {
		if dt, ok := interface{}(this.Property).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}

// Set default value for message flink.FlinkJar
func (this *FlinkJar) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.FlinkScala
func (this *FlinkScala) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.FlinkPython
func (this *FlinkPython) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.FlinkSQL
func (this *FlinkSQL) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.SqlColumnType
func (this *SqlColumnType) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.SqlTimeColumnType
func (this *SqlTimeColumnType) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.ConnectorOption
func (this *ConnectorOption) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.MySQLTable
func (this *MySQLTable) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.PostgreSQLTable
func (this *PostgreSQLTable) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.ClickHouseTable
func (this *ClickHouseTable) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.HBaseTable
func (this *HBaseTable) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.HDFSTable
func (this *HDFSTable) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.S3Table
func (this *S3Table) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.FtpTable
func (this *FtpTable) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.KafkaTable
func (this *KafkaTable) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.TableSchema
func (this *TableSchema) SetDefaults() {
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

// Set default value for message flink.FlinkConfig
func (this *FlinkConfig) SetDefaults() {
	if this == nil {
		return
	}
	if this.RestartStrategy != nil {
		if dt, ok := interface{}(this.RestartStrategy).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	if this.Logger != nil {
		if dt, ok := interface{}(this.Logger).(interface{ SetDefaults() }); ok {
			dt.SetDefaults()
		}
	}
	return
}

// Set default value for message flink.FlinkConfig.Item
func (this *FlinkConfig_Item) SetDefaults() {
	if this == nil {
		return
	}
	return
}

// Set default value for message flink.FlinkConfig.RestartStrategy
func (this *FlinkConfig_RestartStrategy) SetDefaults() {
	if this == nil {
		return
	}
	if this.FixedDelayAttempts == 0 {
		this.FixedDelayAttempts = 1
	}
	if this.FixedDelayDelay == 0 {
		this.FixedDelayDelay = 1
	}
	if this.FailureRateMaxFailuresPerInterval == 0 {
		this.FailureRateMaxFailuresPerInterval = 1
	}
	if this.FailureRateDelay == 0 {
		this.FailureRateDelay = 1
	}
	if this.FailureRateFailureRateInterval == 0 {
		this.FailureRateFailureRateInterval = 3
	}
	return
}

// Set default value for message flink.FlinkConfig.Logger
func (this *FlinkConfig_Logger) SetDefaults() {
	if this == nil {
		return
	}
	if this.RootLogLevel == "" {
		this.RootLogLevel = "INFO"
	}
	return
}