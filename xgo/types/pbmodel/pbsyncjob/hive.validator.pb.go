// Code generated by protoc-gen-govalidator. DO NOT EDIT.
// versions:
// 		protoc-gen-govalidator 0.0.1
// source: proto/types/model/syncjob/hive.proto

package pbsyncjob

import (
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbdefaults"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbgosql"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
	protovalidator "github.com/yu31/protoc-plugin/xgo/pkg/protovalidator"
)

var _xxx_xxx_Validator_HiveTarget_InEnums_PartitionType = map[HiveTarget_PartitionType]bool{0: true, 1: true, 2: true}

func (this *HiveTarget) _xxx_xxx_Validator_Validate_partition_type() error {
	if !(this.PartitionType >= 0) {
		return protovalidator.FieldError1("HiveTarget", "the value of field 'partition_type' must be greater than or equal to '0'", protovalidator.Int32ToString(int32(this.PartitionType)))
	}
	if !(_xxx_xxx_Validator_HiveTarget_InEnums_PartitionType[this.PartitionType]) {
		return protovalidator.FieldError1("HiveTarget", "the value of field 'partition_type' must in enums of '[0 1 2]'", protovalidator.Int32ToString(int32(this.PartitionType)))
	}
	return nil
}

var _xxx_xxx_Validator_HiveTarget_InEnums_WriteMode = map[HiveTarget_WriteMode]bool{0: true, 1: true}

func (this *HiveTarget) _xxx_xxx_Validator_Validate_write_mode() error {
	if !(this.WriteMode >= 0) {
		return protovalidator.FieldError1("HiveTarget", "the value of field 'write_mode' must be greater than or equal to '0'", protovalidator.Int32ToString(int32(this.WriteMode)))
	}
	if !(_xxx_xxx_Validator_HiveTarget_InEnums_WriteMode[this.WriteMode]) {
		return protovalidator.FieldError1("HiveTarget", "the value of field 'write_mode' must in enums of '[0 1]'", protovalidator.Int32ToString(int32(this.WriteMode)))
	}
	return nil
}

var _xxx_xxx_Validator_HiveTarget_InEnums_FileType = map[HiveTarget_FileType]bool{0: true, 1: true, 2: true}

func (this *HiveTarget) _xxx_xxx_Validator_Validate_file_type() error {
	if !(this.FileType >= 0) {
		return protovalidator.FieldError1("HiveTarget", "the value of field 'file_type' must be greater than or equal to '0'", protovalidator.Int32ToString(int32(this.FileType)))
	}
	if !(_xxx_xxx_Validator_HiveTarget_InEnums_FileType[this.FileType]) {
		return protovalidator.FieldError1("HiveTarget", "the value of field 'file_type' must in enums of '[0 1 2]'", protovalidator.Int32ToString(int32(this.FileType)))
	}
	return nil
}

var _xxx_xxx_Validator_HiveTarget_InEnums_Compress = map[HiveTarget_CompressType]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true}

func (this *HiveTarget) _xxx_xxx_Validator_Validate_compress() error {
	if !(this.Compress >= 0) {
		return protovalidator.FieldError1("HiveTarget", "the value of field 'compress' must be greater than or equal to '0'", protovalidator.Int32ToString(int32(this.Compress)))
	}
	if !(_xxx_xxx_Validator_HiveTarget_InEnums_Compress[this.Compress]) {
		return protovalidator.FieldError1("HiveTarget", "the value of field 'compress' must in enums of '[0 1 2 3 4 5]'", protovalidator.Int32ToString(int32(this.Compress)))
	}
	return nil
}

func (this *HiveTarget) _xxx_xxx_Validator_CheckIf_field_delimiter() bool {
	if !(this.FileType == 0) {
		return false
	}
	return true
}

func (this *HiveTarget) _xxx_xxx_Validator_Validate_field_delimiter() error {
	if !this._xxx_xxx_Validator_CheckIf_field_delimiter() {
		return nil
	}
	if !(len(this.FieldDelimiter) >= 1) {
		return protovalidator.FieldError1("HiveTarget", "the byte length of field 'field_delimiter' must be greater than or equal to '1'", protovalidator.StringByteLenToString(this.FieldDelimiter))
	}
	return nil
}

var _xxx_xxx_Validator_HiveTarget_InEnums_Encoding = map[HiveTarget_Encoding]bool{0: true, 1: true}

func (this *HiveTarget) _xxx_xxx_Validator_Validate_encoding() error {
	if !(this.Encoding >= 0) {
		return protovalidator.FieldError1("HiveTarget", "the value of field 'encoding' must be greater than or equal to '0'", protovalidator.Int32ToString(int32(this.Encoding)))
	}
	if !(_xxx_xxx_Validator_HiveTarget_InEnums_Encoding[this.Encoding]) {
		return protovalidator.FieldError1("HiveTarget", "the value of field 'encoding' must in enums of '[0 1]'", protovalidator.Int32ToString(int32(this.Encoding)))
	}
	return nil
}

// Set default value for message model.HiveTarget
func (this *HiveTarget) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_partition_type(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_write_mode(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_file_type(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_compress(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_field_delimiter(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_encoding(); err != nil {
		return err
	}
	return nil
}
