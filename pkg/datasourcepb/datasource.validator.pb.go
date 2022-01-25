// Code generated by protoc-gen-govalidator. DO NOT EDIT.
// versions:
// 		protoc-gen-govalidator v0.1
// source: proto/datasource.proto

package datasourcepb

import (
	_ "github.com/yu31/proto-go-plugin/pkg/pb/gosqlpb"
	_ "github.com/yu31/proto-go-plugin/pkg/pb/validatorpb"
	protovalidator "github.com/yu31/proto-go-plugin/pkg/protovalidator"
)

func (this *MySQLURL) _xxx_xxx_Validator_Validate_host() error {
	if !(len(this.Host) >= 1) {
		return protovalidator.FieldError1("MySQLURL", "the byte length of field 'host' must be greater than or equal to '1'", protovalidator.StringByteLenToString(this.Host))
	}
	if !(len(this.Host) <= 64) {
		return protovalidator.FieldError1("MySQLURL", "the byte length of field 'host' must be less than or equal to '64'", protovalidator.StringByteLenToString(this.Host))
	}
	return nil
}

func (this *MySQLURL) _xxx_xxx_Validator_Validate_port() error {
	if !(this.Port >= 0) {
		return protovalidator.FieldError1("MySQLURL", "the value of field 'port' must be greater than or equal to '0'", protovalidator.Int32ToString(this.Port))
	}
	if !(this.Port <= 65536) {
		return protovalidator.FieldError1("MySQLURL", "the value of field 'port' must be less than or equal to '65536'", protovalidator.Int32ToString(this.Port))
	}
	return nil
}

func (this *MySQLURL) _xxx_xxx_Validator_Validate_user() error {
	if !(len(this.User) >= 1) {
		return protovalidator.FieldError1("MySQLURL", "the byte length of field 'user' must be greater than or equal to '1'", protovalidator.StringByteLenToString(this.User))
	}
	if !(len(this.User) <= 64) {
		return protovalidator.FieldError1("MySQLURL", "the byte length of field 'user' must be less than or equal to '64'", protovalidator.StringByteLenToString(this.User))
	}
	return nil
}

func (this *MySQLURL) _xxx_xxx_Validator_Validate_password() error {
	if !(len(this.Password) >= 1) {
		return protovalidator.FieldError1("MySQLURL", "the byte length of field 'password' must be greater than or equal to '1'", protovalidator.StringByteLenToString(this.Password))
	}
	if !(len(this.Password) <= 64) {
		return protovalidator.FieldError1("MySQLURL", "the byte length of field 'password' must be less than or equal to '64'", protovalidator.StringByteLenToString(this.Password))
	}
	return nil
}

func (this *MySQLURL) _xxx_xxx_Validator_Validate_database() error {
	if !(len(this.Database) >= 1) {
		return protovalidator.FieldError1("MySQLURL", "the byte length of field 'database' must be greater than or equal to '1'", protovalidator.StringByteLenToString(this.Database))
	}
	if !(len(this.Database) <= 64) {
		return protovalidator.FieldError1("MySQLURL", "the byte length of field 'database' must be less than or equal to '64'", protovalidator.StringByteLenToString(this.Database))
	}
	return nil
}

// Set default value for message datasource.MySQLURL
func (this *MySQLURL) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_host(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_port(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_user(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_password(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_database(); err != nil {
		return err
	}
	return nil
}

func (this *PostgreSQLURL) _xxx_xxx_Validator_Validate_host() error {
	if !(len(this.Host) >= 1) {
		return protovalidator.FieldError1("PostgreSQLURL", "the byte length of field 'host' must be greater than or equal to '1'", protovalidator.StringByteLenToString(this.Host))
	}
	if !(len(this.Host) <= 64) {
		return protovalidator.FieldError1("PostgreSQLURL", "the byte length of field 'host' must be less than or equal to '64'", protovalidator.StringByteLenToString(this.Host))
	}
	return nil
}

func (this *PostgreSQLURL) _xxx_xxx_Validator_Validate_port() error {
	if !(this.Port >= 0) {
		return protovalidator.FieldError1("PostgreSQLURL", "the value of field 'port' must be greater than or equal to '0'", protovalidator.Int32ToString(this.Port))
	}
	if !(this.Port <= 65536) {
		return protovalidator.FieldError1("PostgreSQLURL", "the value of field 'port' must be less than or equal to '65536'", protovalidator.Int32ToString(this.Port))
	}
	return nil
}

func (this *PostgreSQLURL) _xxx_xxx_Validator_Validate_user() error {
	if !(len(this.User) >= 1) {
		return protovalidator.FieldError1("PostgreSQLURL", "the byte length of field 'user' must be greater than or equal to '1'", protovalidator.StringByteLenToString(this.User))
	}
	if !(len(this.User) <= 64) {
		return protovalidator.FieldError1("PostgreSQLURL", "the byte length of field 'user' must be less than or equal to '64'", protovalidator.StringByteLenToString(this.User))
	}
	return nil
}

func (this *PostgreSQLURL) _xxx_xxx_Validator_Validate_password() error {
	if !(len(this.Password) >= 1) {
		return protovalidator.FieldError1("PostgreSQLURL", "the byte length of field 'password' must be greater than or equal to '1'", protovalidator.StringByteLenToString(this.Password))
	}
	if !(len(this.Password) <= 64) {
		return protovalidator.FieldError1("PostgreSQLURL", "the byte length of field 'password' must be less than or equal to '64'", protovalidator.StringByteLenToString(this.Password))
	}
	return nil
}

func (this *PostgreSQLURL) _xxx_xxx_Validator_Validate_database() error {
	if !(len(this.Database) >= 1) {
		return protovalidator.FieldError1("PostgreSQLURL", "the byte length of field 'database' must be greater than or equal to '1'", protovalidator.StringByteLenToString(this.Database))
	}
	if !(len(this.Database) <= 64) {
		return protovalidator.FieldError1("PostgreSQLURL", "the byte length of field 'database' must be less than or equal to '64'", protovalidator.StringByteLenToString(this.Database))
	}
	return nil
}

// Set default value for message datasource.PostgreSQLURL
func (this *PostgreSQLURL) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_host(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_port(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_user(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_password(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_database(); err != nil {
		return err
	}
	return nil
}

func (this *ClickHouseURL) _xxx_xxx_Validator_Validate_host() error {
	if !(len(this.Host) >= 1) {
		return protovalidator.FieldError1("ClickHouseURL", "the byte length of field 'host' must be greater than or equal to '1'", protovalidator.StringByteLenToString(this.Host))
	}
	if !(len(this.Host) <= 64) {
		return protovalidator.FieldError1("ClickHouseURL", "the byte length of field 'host' must be less than or equal to '64'", protovalidator.StringByteLenToString(this.Host))
	}
	return nil
}

func (this *ClickHouseURL) _xxx_xxx_Validator_Validate_port() error {
	if !(this.Port >= 0) {
		return protovalidator.FieldError1("ClickHouseURL", "the value of field 'port' must be greater than or equal to '0'", protovalidator.Int32ToString(this.Port))
	}
	if !(this.Port <= 65536) {
		return protovalidator.FieldError1("ClickHouseURL", "the value of field 'port' must be less than or equal to '65536'", protovalidator.Int32ToString(this.Port))
	}
	return nil
}

func (this *ClickHouseURL) _xxx_xxx_Validator_Validate_user() error {
	if !(len(this.User) >= 1) {
		return protovalidator.FieldError1("ClickHouseURL", "the byte length of field 'user' must be greater than or equal to '1'", protovalidator.StringByteLenToString(this.User))
	}
	if !(len(this.User) <= 64) {
		return protovalidator.FieldError1("ClickHouseURL", "the byte length of field 'user' must be less than or equal to '64'", protovalidator.StringByteLenToString(this.User))
	}
	return nil
}

func (this *ClickHouseURL) _xxx_xxx_Validator_Validate_password() error {
	if !(len(this.Password) >= 1) {
		return protovalidator.FieldError1("ClickHouseURL", "the byte length of field 'password' must be greater than or equal to '1'", protovalidator.StringByteLenToString(this.Password))
	}
	if !(len(this.Password) <= 64) {
		return protovalidator.FieldError1("ClickHouseURL", "the byte length of field 'password' must be less than or equal to '64'", protovalidator.StringByteLenToString(this.Password))
	}
	return nil
}

func (this *ClickHouseURL) _xxx_xxx_Validator_Validate_database() error {
	if !(len(this.Database) >= 1) {
		return protovalidator.FieldError1("ClickHouseURL", "the byte length of field 'database' must be greater than or equal to '1'", protovalidator.StringByteLenToString(this.Database))
	}
	if !(len(this.Database) <= 64) {
		return protovalidator.FieldError1("ClickHouseURL", "the byte length of field 'database' must be less than or equal to '64'", protovalidator.StringByteLenToString(this.Database))
	}
	return nil
}

// Set default value for message datasource.ClickHouseURL
func (this *ClickHouseURL) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_host(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_port(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_user(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_password(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_database(); err != nil {
		return err
	}
	return nil
}

func (this *KafkaURL) _xxx_xxx_Validator_Validate_kafka_brokers() error {
	if !(len(this.KafkaBrokers) >= 1) {
		return protovalidator.FieldError1("KafkaURL", "the byte length of field 'kafka_brokers' must be greater than or equal to '1'", protovalidator.StringByteLenToString(this.KafkaBrokers))
	}
	if !(len(this.KafkaBrokers) <= 1024) {
		return protovalidator.FieldError1("KafkaURL", "the byte length of field 'kafka_brokers' must be less than or equal to '1024'", protovalidator.StringByteLenToString(this.KafkaBrokers))
	}
	return nil
}

// Set default value for message datasource.KafkaURL
func (this *KafkaURL) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_kafka_brokers(); err != nil {
		return err
	}
	return nil
}

// Set default value for message datasource.S3URL
func (this *S3URL) Validate() error {
	if this == nil {
		return nil
	}
	return nil
}

func (this *HBaseURL) _xxx_xxx_Validator_Validate_zk_hosts() error {
	if !(len(this.ZkHosts) >= 1) {
		return protovalidator.FieldError1("HBaseURL", "the byte length of field 'zk_hosts' must be greater than or equal to '1'", protovalidator.StringByteLenToString(this.ZkHosts))
	}
	if !(len(this.ZkHosts) <= 1024) {
		return protovalidator.FieldError1("HBaseURL", "the byte length of field 'zk_hosts' must be less than or equal to '1024'", protovalidator.StringByteLenToString(this.ZkHosts))
	}
	return nil
}

func (this *HBaseURL) _xxx_xxx_Validator_Validate_zk_path() error {
	if !(len(this.ZkPath) >= 1) {
		return protovalidator.FieldError1("HBaseURL", "the byte length of field 'zk_path' must be greater than or equal to '1'", protovalidator.StringByteLenToString(this.ZkPath))
	}
	if !(len(this.ZkPath) <= 1024) {
		return protovalidator.FieldError1("HBaseURL", "the byte length of field 'zk_path' must be less than or equal to '1024'", protovalidator.StringByteLenToString(this.ZkPath))
	}
	return nil
}

// Set default value for message datasource.HBaseURL
func (this *HBaseURL) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_zk_hosts(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_zk_path(); err != nil {
		return err
	}
	return nil
}

func (this *FtpURL) _xxx_xxx_Validator_Validate_host() error {
	if !(len(this.Host) >= 1) {
		return protovalidator.FieldError1("FtpURL", "the byte length of field 'host' must be greater than or equal to '1'", protovalidator.StringByteLenToString(this.Host))
	}
	if !(len(this.Host) <= 64) {
		return protovalidator.FieldError1("FtpURL", "the byte length of field 'host' must be less than or equal to '64'", protovalidator.StringByteLenToString(this.Host))
	}
	return nil
}

func (this *FtpURL) _xxx_xxx_Validator_Validate_port() error {
	if !(this.Port >= 0) {
		return protovalidator.FieldError1("FtpURL", "the value of field 'port' must be greater than or equal to '0'", protovalidator.Int32ToString(this.Port))
	}
	if !(this.Port <= 65536) {
		return protovalidator.FieldError1("FtpURL", "the value of field 'port' must be less than or equal to '65536'", protovalidator.Int32ToString(this.Port))
	}
	return nil
}

func (this *FtpURL) _xxx_xxx_Validator_Validate_user() error {
	if !(len(this.User) >= 1) {
		return protovalidator.FieldError1("FtpURL", "the byte length of field 'user' must be greater than or equal to '1'", protovalidator.StringByteLenToString(this.User))
	}
	if !(len(this.User) <= 64) {
		return protovalidator.FieldError1("FtpURL", "the byte length of field 'user' must be less than or equal to '64'", protovalidator.StringByteLenToString(this.User))
	}
	return nil
}

func (this *FtpURL) _xxx_xxx_Validator_Validate_password() error {
	if !(len(this.Password) >= 1) {
		return protovalidator.FieldError1("FtpURL", "the byte length of field 'password' must be greater than or equal to '1'", protovalidator.StringByteLenToString(this.Password))
	}
	if !(len(this.Password) <= 64) {
		return protovalidator.FieldError1("FtpURL", "the byte length of field 'password' must be less than or equal to '64'", protovalidator.StringByteLenToString(this.Password))
	}
	return nil
}

// Set default value for message datasource.FtpURL
func (this *FtpURL) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_host(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_port(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_user(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_password(); err != nil {
		return err
	}
	return nil
}

func (this *HDFSURL) _xxx_xxx_Validator_Validate_nodes() error {
	if !(this.Nodes != nil) {
		return protovalidator.FieldError2("HDFSURL", "the value of field 'nodes' cannot be null")
	}
	if dt, ok := interface{}(this.Nodes).(interface{ Validate() error }); ok {
		if err := dt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Set default value for message datasource.HDFSURL
func (this *HDFSURL) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_nodes(); err != nil {
		return err
	}
	return nil
}

func (this *HDFSURL_NodeURL) _xxx_xxx_Validator_Validate_name_node() error {
	if !(len(this.NameNode) >= 1) {
		return protovalidator.FieldError1("HDFSURL_NodeURL", "the byte length of field 'name_node' must be greater than or equal to '1'", protovalidator.StringByteLenToString(this.NameNode))
	}
	if !(len(this.NameNode) <= 64) {
		return protovalidator.FieldError1("HDFSURL_NodeURL", "the byte length of field 'name_node' must be less than or equal to '64'", protovalidator.StringByteLenToString(this.NameNode))
	}
	return nil
}

func (this *HDFSURL_NodeURL) _xxx_xxx_Validator_Validate_port() error {
	if !(this.Port >= 0) {
		return protovalidator.FieldError1("HDFSURL_NodeURL", "the value of field 'port' must be greater than or equal to '0'", protovalidator.Int32ToString(this.Port))
	}
	if !(this.Port <= 65536) {
		return protovalidator.FieldError1("HDFSURL_NodeURL", "the value of field 'port' must be less than or equal to '65536'", protovalidator.Int32ToString(this.Port))
	}
	return nil
}

// Set default value for message datasource.HDFSURL.NodeURL
func (this *HDFSURL_NodeURL) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_name_node(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_port(); err != nil {
		return err
	}
	return nil
}

// Set default value for message datasource.SourceKind
func (this *SourceKind) Validate() error {
	if this == nil {
		return nil
	}
	return nil
}
