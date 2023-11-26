// Code generated by protoc-gen-gosql. DO NOT EDIT.
// versions:
// 		protoc-gen-gosql 0.0.1
// source: proto/types/model/datasource/datasource_url.proto

package pbdatasource

import (
	driver "database/sql/driver"
	json "encoding/json"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbgosql"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
)

// Scan for implements sql.Scanner (- database/sql).
func (t *MySQLURL) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), t)
}

// Value for implements driver.Valuer (- database/sql/driver)
func (t *MySQLURL) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return json.Marshal(t)
}

// Scan for implements sql.Scanner (- database/sql).
func (t *PostgreSQLURL) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), t)
}

// Value for implements driver.Valuer (- database/sql/driver)
func (t *PostgreSQLURL) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return json.Marshal(t)
}

// Scan for implements sql.Scanner (- database/sql).
func (t *ClickHouseURL) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), t)
}

// Value for implements driver.Valuer (- database/sql/driver)
func (t *ClickHouseURL) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return json.Marshal(t)
}

// Scan for implements sql.Scanner (- database/sql).
func (t *FtpURL) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), t)
}

// Value for implements driver.Valuer (- database/sql/driver)
func (t *FtpURL) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return json.Marshal(t)
}

// Scan for implements sql.Scanner (- database/sql).
func (t *S3URL) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), t)
}

// Value for implements driver.Valuer (- database/sql/driver)
func (t *S3URL) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return json.Marshal(t)
}

// Scan for implements sql.Scanner (- database/sql).
func (t *KafkaURL) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), t)
}

// Value for implements driver.Valuer (- database/sql/driver)
func (t *KafkaURL) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return json.Marshal(t)
}

// Scan for implements sql.Scanner (- database/sql).
func (t *HBaseURL) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), t)
}

// Value for implements driver.Valuer (- database/sql/driver)
func (t *HBaseURL) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return json.Marshal(t)
}

// Scan for implements sql.Scanner (- database/sql).
func (t *HDFSURL) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), t)
}

// Value for implements driver.Valuer (- database/sql/driver)
func (t *HDFSURL) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return json.Marshal(t)
}

// Scan for implements sql.Scanner (- database/sql).
func (t *SqlServerURL) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), t)
}

// Value for implements driver.Valuer (- database/sql/driver)
func (t *SqlServerURL) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return json.Marshal(t)
}

// Scan for implements sql.Scanner (- database/sql).
func (t *OracleURL) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), t)
}

// Value for implements driver.Valuer (- database/sql/driver)
func (t *OracleURL) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return json.Marshal(t)
}

// Scan for implements sql.Scanner (- database/sql).
func (t *DB2URL) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), t)
}

// Value for implements driver.Valuer (- database/sql/driver)
func (t *DB2URL) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return json.Marshal(t)
}

// Scan for implements sql.Scanner (- database/sql).
func (t *SapHanaURL) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), t)
}

// Value for implements driver.Valuer (- database/sql/driver)
func (t *SapHanaURL) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return json.Marshal(t)
}

// Scan for implements sql.Scanner (- database/sql).
func (t *HiveURL) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), t)
}

// Value for implements driver.Valuer (- database/sql/driver)
func (t *HiveURL) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return json.Marshal(t)
}

// Scan for implements sql.Scanner (- database/sql).
func (t *ElasticSearchURL) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), t)
}

// Value for implements driver.Valuer (- database/sql/driver)
func (t *ElasticSearchURL) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return json.Marshal(t)
}

// Scan for implements sql.Scanner (- database/sql).
func (t *MongoDbURL) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), t)
}

// Value for implements driver.Valuer (- database/sql/driver)
func (t *MongoDbURL) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return json.Marshal(t)
}

// Scan for implements sql.Scanner (- database/sql).
func (t *RedisURL) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), t)
}

// Value for implements driver.Valuer (- database/sql/driver)
func (t *RedisURL) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return json.Marshal(t)
}

// Scan for implements sql.Scanner (- database/sql).
func (t *OceanBaseURL) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), t)
}

// Value for implements driver.Valuer (- database/sql/driver)
func (t *OceanBaseURL) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	return json.Marshal(t)
}
