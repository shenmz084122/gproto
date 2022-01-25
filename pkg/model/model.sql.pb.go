// Code generated by protoc-gen-gosql. DO NOT EDIT.
// versions:
// 		protoc-gen-gosql v0.1
// source: proto/model.proto

package model

import (
	driver "database/sql/driver"
	json "encoding/json"
	_ "github.com/DataWorkbench/gproto/pkg/datasourcepb"
	_ "github.com/DataWorkbench/gproto/pkg/flinkpb"
	_ "github.com/yu31/proto-go-plugin/pkg/pb/gosqlpb"
	_ "github.com/yu31/proto-go-plugin/pkg/pb/validatorpb"
)

// Scan for implements sql.Scanner (- database/sql).
func (t *StreamJobCode) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), t)
}

// Value for implements driver.Valuer (- database/sql/driver)
func (t *StreamJobCode) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return json.Marshal(t)
}

// Scan for implements sql.Scanner (- database/sql).
func (t *StreamJobArgs) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), t)
}

// Value for implements driver.Valuer (- database/sql/driver)
func (t *StreamJobArgs) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return json.Marshal(t)
}

// Scan for implements sql.Scanner (- database/sql).
func (t *StreamJobSchedule) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), t)
}

// Value for implements driver.Valuer (- database/sql/driver)
func (t *StreamJobSchedule) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return json.Marshal(t)
}

// Scan for implements sql.Scanner (- database/sql).
func (t *DataSource_URL) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), t)
}

// Value for implements driver.Valuer (- database/sql/driver)
func (t *DataSource_URL) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return json.Marshal(t)
}

// Scan for implements sql.Scanner (- database/sql).
func (t *HostAliases) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), t)
}

// Value for implements driver.Valuer (- database/sql/driver)
func (t *HostAliases) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return json.Marshal(t)
}
