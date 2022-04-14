// Code generated by protoc-gen-gosql. DO NOT EDIT.
// versions:
// 		protoc-gen-gosql 0.0.1
// source: proto/types/model/syncjob/ftp.proto

package pbsyncjob

import (
	driver "database/sql/driver"
	json "encoding/json"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbgosql"
)

// Scan for implements sql.Scanner (- database/sql).
func (t *FtpSource) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), t)
}

// Value for implements driver.Valuer (- database/sql/driver)
func (t *FtpSource) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return json.Marshal(t)
}

// Scan for implements sql.Scanner (- database/sql).
func (t *FtpTarget) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), t)
}

// Value for implements driver.Valuer (- database/sql/driver)
func (t *FtpTarget) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return json.Marshal(t)
}
