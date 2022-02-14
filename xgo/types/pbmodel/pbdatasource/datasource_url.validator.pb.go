// Code generated by protoc-gen-govalidator. DO NOT EDIT.
// versions:
// 		protoc-gen-govalidator 0.0.1
// source: proto/types/model/datasource/datasource_url.proto

package pbdatasource

import (
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbgosql"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
	protovalidator "github.com/yu31/protoc-plugin/xgo/pkg/protovalidator"
	strconv "strconv"
)

func (this *Host) _xxx_xxx_Validator_Validate_host() error {
	if !(len(this.Host) >= 1) {
		return protovalidator.FieldError1("Host", "the byte length of field 'host' must be greater than or equal to '1'", protovalidator.StringByteLenToString(this.Host))
	}
	if !(len(this.Host) <= 256) {
		return protovalidator.FieldError1("Host", "the byte length of field 'host' must be less than or equal to '256'", protovalidator.StringByteLenToString(this.Host))
	}
	return nil
}

func (this *Host) _xxx_xxx_Validator_Validate_port() error {
	if !(this.Port > 0) {
		return protovalidator.FieldError1("Host", "the value of field 'port' must be greater than '0'", protovalidator.Int32ToString(this.Port))
	}
	if !(this.Port < 65535) {
		return protovalidator.FieldError1("Host", "the value of field 'port' must be less than '65535'", protovalidator.Int32ToString(this.Port))
	}
	return nil
}

// Set default value for message datasource.Host
func (this *Host) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_host(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_port(); err != nil {
		return err
	}
	return nil
}

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

var _xxx_xxx_Validator_FtpURL_InEnums_Protocol = map[FtpURL_Protocol]bool{0: true, 1: true, 2: true}

func (this *FtpURL) _xxx_xxx_Validator_Validate_protocol() error {
	if !(this.Protocol > 0) {
		return protovalidator.FieldError1("FtpURL", "the value of field 'protocol' must be greater than '0'", protovalidator.Int32ToString(int32(this.Protocol)))
	}
	if !(_xxx_xxx_Validator_FtpURL_InEnums_Protocol[this.Protocol]) {
		return protovalidator.FieldError1("FtpURL", "the value of field 'protocol' must in enums of '[0 1 2]'", protovalidator.Int32ToString(int32(this.Protocol)))
	}
	return nil
}

var _xxx_xxx_Validator_FtpURL_InEnums_ConnectionMode = map[FtpURL_ConnectionMode]bool{0: true, 1: true, 2: true}

func (this *FtpURL) _xxx_xxx_Validator_Validate_connection_mode() error {
	if !(this.ConnectionMode > 0) {
		return protovalidator.FieldError1("FtpURL", "the value of field 'connection_mode' must be greater than '0'", protovalidator.Int32ToString(int32(this.ConnectionMode)))
	}
	if !(_xxx_xxx_Validator_FtpURL_InEnums_ConnectionMode[this.ConnectionMode]) {
		return protovalidator.FieldError1("FtpURL", "the value of field 'connection_mode' must in enums of '[0 1 2]'", protovalidator.Int32ToString(int32(this.ConnectionMode)))
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
	if err := this._xxx_xxx_Validator_Validate_protocol(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_connection_mode(); err != nil {
		return err
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

// Set default value for message datasource.S3URL
func (this *S3URL) Validate() error {
	if this == nil {
		return nil
	}
	return nil
}

func (this *KafkaURL) _xxx_xxx_Validator_Validate_kafka_brokers() error {
	if !(len(this.KafkaBrokers) >= 1) {
		return protovalidator.FieldError1("KafkaURL", "the length of field 'kafka_brokers' must be greater than or equal to '1'", strconv.Itoa(len(this.KafkaBrokers)))
	}
	if !(len(this.KafkaBrokers) <= 128) {
		return protovalidator.FieldError1("KafkaURL", "the length of field 'kafka_brokers' must be less than or equal to '128'", strconv.Itoa(len(this.KafkaBrokers)))
	}
	for _, item := range this.KafkaBrokers {
		_ = item // To avoid unused panics.
		if dt, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := dt.Validate(); err != nil {
				return err
			}
		}
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

func (this *HBaseURL) _xxx_xxx_Validator_Validate_config() error {
	if !(len(this.Config) >= 1) {
		return protovalidator.FieldError1("HBaseURL", "the byte length of field 'config' must be greater than or equal to '1'", protovalidator.StringByteLenToString(this.Config))
	}
	if !(len(this.Config) <= 16384) {
		return protovalidator.FieldError1("HBaseURL", "the byte length of field 'config' must be less than or equal to '16384'", protovalidator.StringByteLenToString(this.Config))
	}
	if !(protovalidator.StringIsJSON(this.Config)) {
		return protovalidator.FieldError1("HBaseURL", "the value of field 'config' must be a string in JSON format", this.Config)
	}
	return nil
}

// Set default value for message datasource.HBaseURL
func (this *HBaseURL) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_config(); err != nil {
		return err
	}
	return nil
}

func (this *HDFSURL) _xxx_xxx_Validator_Validate_name_node() error {
	if !(len(this.NameNode) >= 1) {
		return protovalidator.FieldError1("HDFSURL", "the byte length of field 'name_node' must be greater than or equal to '1'", protovalidator.StringByteLenToString(this.NameNode))
	}
	if !(len(this.NameNode) <= 64) {
		return protovalidator.FieldError1("HDFSURL", "the byte length of field 'name_node' must be less than or equal to '64'", protovalidator.StringByteLenToString(this.NameNode))
	}
	return nil
}

func (this *HDFSURL) _xxx_xxx_Validator_Validate_port() error {
	if !(this.Port >= 0) {
		return protovalidator.FieldError1("HDFSURL", "the value of field 'port' must be greater than or equal to '0'", protovalidator.Int32ToString(this.Port))
	}
	if !(this.Port <= 65536) {
		return protovalidator.FieldError1("HDFSURL", "the value of field 'port' must be less than or equal to '65536'", protovalidator.Int32ToString(this.Port))
	}
	return nil
}

func (this *HDFSURL) _xxx_xxx_Validator_CheckIf_config() bool {
	if !(this.Config != "") {
		return false
	}
	return true
}

func (this *HDFSURL) _xxx_xxx_Validator_Validate_config() error {
	if !this._xxx_xxx_Validator_CheckIf_config() {
		return nil
	}
	if !(len(this.Config) <= 16384) {
		return protovalidator.FieldError1("HDFSURL", "the byte length of field 'config' must be less than or equal to '16384'", protovalidator.StringByteLenToString(this.Config))
	}
	if !(protovalidator.StringIsJSON(this.Config)) {
		return protovalidator.FieldError1("HDFSURL", "the value of field 'config' must be a string in JSON format", this.Config)
	}
	return nil
}

// Set default value for message datasource.HDFSURL
func (this *HDFSURL) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_name_node(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_port(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_config(); err != nil {
		return err
	}
	return nil
}