// Code generated by protoc-gen-govalidator. DO NOT EDIT.
// versions:
// 		protoc-gen-govalidator 0.0.1
// source: proto/types/model/dataservice.proto

package pbmodel

import (
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbdefaults"
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
	protovalidator "github.com/yu31/protoc-plugin/xgo/pkg/protovalidator"
	_ "google.golang.org/protobuf/types/known/anypb"
	strings "strings"
)

func (this *DataServiceCluster) _xxx_xxx_Validator_Validate_id() error {
	if !(len(this.Id) == 20) {
		return protovalidator.FieldError1("DataServiceCluster", "the byte length of field 'id' must be equal to '20'", protovalidator.StringByteLenToString(this.Id))
	}
	if !(strings.HasPrefix(this.Id, "dsc-")) {
		return protovalidator.FieldError1("DataServiceCluster", "the value of field 'id' must start with string 'dsc-'", this.Id)
	}
	return nil
}

func (this *DataServiceCluster) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("DataServiceCluster", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("DataServiceCluster", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

func (this *DataServiceCluster) _xxx_xxx_Validator_Validate_name() error {
	if !(len(this.Name) >= 2) {
		return protovalidator.FieldError1("DataServiceCluster", "the byte length of field 'name' must be greater than or equal to '2'", protovalidator.StringByteLenToString(this.Name))
	}
	if !(len(this.Name) <= 128) {
		return protovalidator.FieldError1("DataServiceCluster", "the byte length of field 'name' must be less than or equal to '128'", protovalidator.StringByteLenToString(this.Name))
	}
	return nil
}

var _xxx_xxx_Validator_DataServiceCluster_InEnums_ResourceSpec = map[DataServiceCluster_ResourceSpec]bool{0: true, 1: true, 2: true, 3: true}

func (this *DataServiceCluster) _xxx_xxx_Validator_Validate_resource_spec() error {
	if !(this.ResourceSpec > 0) {
		return protovalidator.FieldError1("DataServiceCluster", "the value of field 'resource_spec' must be greater than '0'", protovalidator.Int32ToString(int32(this.ResourceSpec)))
	}
	if !(_xxx_xxx_Validator_DataServiceCluster_InEnums_ResourceSpec[this.ResourceSpec]) {
		return protovalidator.FieldError1("DataServiceCluster", "the value of field 'resource_spec' must in enums of '[0 1 2 3]'", protovalidator.Int32ToString(int32(this.ResourceSpec)))
	}
	return nil
}

var _xxx_xxx_Validator_DataServiceCluster_InEnums_Status = map[DataServiceCluster_Status]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true}

func (this *DataServiceCluster) _xxx_xxx_Validator_Validate_status() error {
	if !(this.Status > 0) {
		return protovalidator.FieldError1("DataServiceCluster", "the value of field 'status' must be greater than '0'", protovalidator.Int32ToString(int32(this.Status)))
	}
	if !(_xxx_xxx_Validator_DataServiceCluster_InEnums_Status[this.Status]) {
		return protovalidator.FieldError1("DataServiceCluster", "the value of field 'status' must in enums of '[0 1 2 3 4 5 6]'", protovalidator.Int32ToString(int32(this.Status)))
	}
	return nil
}

// Set default value for message model.DataServiceCluster
func (this *DataServiceCluster) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_space_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_name(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_resource_spec(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_status(); err != nil {
		return err
	}
	return nil
}

func (this *ApiGroup) _xxx_xxx_Validator_Validate_id() error {
	if !(len(this.Id) == 20) {
		return protovalidator.FieldError1("ApiGroup", "the byte length of field 'id' must be equal to '20'", protovalidator.StringByteLenToString(this.Id))
	}
	if !(strings.HasPrefix(this.Id, "dsg-")) {
		return protovalidator.FieldError1("ApiGroup", "the value of field 'id' must start with string 'dsg-'", this.Id)
	}
	return nil
}

func (this *ApiGroup) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("ApiGroup", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("ApiGroup", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

func (this *ApiGroup) _xxx_xxx_Validator_Validate_name() error {
	if !(len(this.Name) > 0) {
		return protovalidator.FieldError1("ApiGroup", "the byte length of field 'name' must be greater than '0'", protovalidator.StringByteLenToString(this.Name))
	}
	if !(len(this.Name) <= 64) {
		return protovalidator.FieldError1("ApiGroup", "the byte length of field 'name' must be less than or equal to '64'", protovalidator.StringByteLenToString(this.Name))
	}
	return nil
}

// Set default value for message model.ApiGroup
func (this *ApiGroup) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_space_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_name(); err != nil {
		return err
	}
	return nil
}

func (this *ApiConfig) _xxx_xxx_Validator_Validate_api_id() error {
	if !(len(this.ApiId) == 20) {
		return protovalidator.FieldError1("ApiConfig", "the byte length of field 'api_id' must be equal to '20'", protovalidator.StringByteLenToString(this.ApiId))
	}
	if !(strings.HasPrefix(this.ApiId, "dsa-")) {
		return protovalidator.FieldError1("ApiConfig", "the value of field 'api_id' must start with string 'dsa-'", this.ApiId)
	}
	return nil
}

func (this *ApiConfig) _xxx_xxx_Validator_Validate_group_id() error {
	if !(len(this.GroupId) == 20) {
		return protovalidator.FieldError1("ApiConfig", "the byte length of field 'group_id' must be equal to '20'", protovalidator.StringByteLenToString(this.GroupId))
	}
	if !(strings.HasPrefix(this.GroupId, "dsg-")) {
		return protovalidator.FieldError1("ApiConfig", "the value of field 'group_id' must start with string 'dsg-'", this.GroupId)
	}
	return nil
}

func (this *ApiConfig) _xxx_xxx_Validator_Validate_cluster_id() error {
	if !(len(this.ClusterId) == 20) {
		return protovalidator.FieldError1("ApiConfig", "the byte length of field 'cluster_id' must be equal to '20'", protovalidator.StringByteLenToString(this.ClusterId))
	}
	if !(strings.HasPrefix(this.ClusterId, "dsc-")) {
		return protovalidator.FieldError1("ApiConfig", "the value of field 'cluster_id' must start with string 'dsc-'", this.ClusterId)
	}
	return nil
}

func (this *ApiConfig) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("ApiConfig", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("ApiConfig", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

func (this *ApiConfig) _xxx_xxx_Validator_Validate_api_name() error {
	if !(len(this.ApiName) > 0) {
		return protovalidator.FieldError1("ApiConfig", "the byte length of field 'api_name' must be greater than '0'", protovalidator.StringByteLenToString(this.ApiName))
	}
	if !(len(this.ApiName) <= 64) {
		return protovalidator.FieldError1("ApiConfig", "the byte length of field 'api_name' must be less than or equal to '64'", protovalidator.StringByteLenToString(this.ApiName))
	}
	return nil
}

func (this *ApiConfig) _xxx_xxx_Validator_Validate_api_path() error {
	if !(len(this.ApiPath) > 0) {
		return protovalidator.FieldError1("ApiConfig", "the byte length of field 'api_path' must be greater than '0'", protovalidator.StringByteLenToString(this.ApiPath))
	}
	if !(len(this.ApiPath) <= 200) {
		return protovalidator.FieldError1("ApiConfig", "the byte length of field 'api_path' must be less than or equal to '200'", protovalidator.StringByteLenToString(this.ApiPath))
	}
	return nil
}

var _xxx_xxx_Validator_ApiConfig_InEnums_ApiMode = map[ApiConfig_ApiMode]bool{0: true, 1: true, 2: true}

func (this *ApiConfig) _xxx_xxx_Validator_Validate_api_mode() error {
	if !(this.ApiMode > 0) {
		return protovalidator.FieldError1("ApiConfig", "the value of field 'api_mode' must be greater than '0'", protovalidator.Int32ToString(int32(this.ApiMode)))
	}
	if !(_xxx_xxx_Validator_ApiConfig_InEnums_ApiMode[this.ApiMode]) {
		return protovalidator.FieldError1("ApiConfig", "the value of field 'api_mode' must in enums of '[0 1 2]'", protovalidator.Int32ToString(int32(this.ApiMode)))
	}
	return nil
}

func (this *ApiConfig) _xxx_xxx_Validator_Validate_api_description() error {
	if !(len(this.ApiDescription) <= 257) {
		return protovalidator.FieldError1("ApiConfig", "the byte length of field 'api_description' must be less than or equal to '257'", protovalidator.StringByteLenToString(this.ApiDescription))
	}
	return nil
}

var _xxx_xxx_Validator_ApiConfig_InEnums_RequestMethod = map[ApiConfig_RequestMethod]bool{0: true, 1: true, 2: true}

func (this *ApiConfig) _xxx_xxx_Validator_Validate_request_method() error {
	if !(this.RequestMethod > 0) {
		return protovalidator.FieldError1("ApiConfig", "the value of field 'request_method' must be greater than '0'", protovalidator.Int32ToString(int32(this.RequestMethod)))
	}
	if !(_xxx_xxx_Validator_ApiConfig_InEnums_RequestMethod[this.RequestMethod]) {
		return protovalidator.FieldError1("ApiConfig", "the value of field 'request_method' must in enums of '[0 1 2]'", protovalidator.Int32ToString(int32(this.RequestMethod)))
	}
	return nil
}

var _xxx_xxx_Validator_ApiConfig_InEnums_ResponseType = map[ApiConfig_ResponseType]bool{0: true, 1: true, 2: true}

func (this *ApiConfig) _xxx_xxx_Validator_Validate_response_type() error {
	if !(this.ResponseType > 0) {
		return protovalidator.FieldError1("ApiConfig", "the value of field 'response_type' must be greater than '0'", protovalidator.Int32ToString(int32(this.ResponseType)))
	}
	if !(_xxx_xxx_Validator_ApiConfig_InEnums_ResponseType[this.ResponseType]) {
		return protovalidator.FieldError1("ApiConfig", "the value of field 'response_type' must in enums of '[0 1 2]'", protovalidator.Int32ToString(int32(this.ResponseType)))
	}
	return nil
}

func (this *ApiConfig) _xxx_xxx_Validator_Validate_timeout() error {
	if !(this.Timeout >= 1) {
		return protovalidator.FieldError1("ApiConfig", "the value of field 'timeout' must be greater than or equal to '1'", protovalidator.Int64ToString(this.Timeout))
	}
	if !(this.Timeout <= 180) {
		return protovalidator.FieldError1("ApiConfig", "the value of field 'timeout' must be less than or equal to '180'", protovalidator.Int64ToString(this.Timeout))
	}
	return nil
}

func (this *ApiConfig) _xxx_xxx_Validator_Validate_datasource_id() error {
	if !(len(this.DatasourceId) == 20) {
		return protovalidator.FieldError1("ApiConfig", "the byte length of field 'datasource_id' must be equal to '20'", protovalidator.StringByteLenToString(this.DatasourceId))
	}
	if !(strings.HasPrefix(this.DatasourceId, "som-")) {
		return protovalidator.FieldError1("ApiConfig", "the value of field 'datasource_id' must start with string 'som-'", this.DatasourceId)
	}
	return nil
}

func (this *ApiConfig) _xxx_xxx_Validator_Validate_table_name() error {
	if !(len(this.TableName) > 0) {
		return protovalidator.FieldError1("ApiConfig", "the byte length of field 'table_name' must be greater than '0'", protovalidator.StringByteLenToString(this.TableName))
	}
	if !(len(this.TableName) <= 64) {
		return protovalidator.FieldError1("ApiConfig", "the byte length of field 'table_name' must be less than or equal to '64'", protovalidator.StringByteLenToString(this.TableName))
	}
	return nil
}

func (this *ApiConfig) _xxx_xxx_Validator_Validate_script() error {
	if !(len(this.Script) <= 20000) {
		return protovalidator.FieldError1("ApiConfig", "the byte length of field 'script' must be less than or equal to '20000'", protovalidator.StringByteLenToString(this.Script))
	}
	return nil
}

func (this *ApiConfig) _xxx_xxx_Validator_Validate_created() error {
	if !(this.Created > 0) {
		return protovalidator.FieldError1("ApiConfig", "the value of field 'created' must be greater than '0'", protovalidator.Int64ToString(this.Created))
	}
	return nil
}

func (this *ApiConfig) _xxx_xxx_Validator_Validate_updated() error {
	if !(this.Updated > 0) {
		return protovalidator.FieldError1("ApiConfig", "the value of field 'updated' must be greater than '0'", protovalidator.Int64ToString(this.Updated))
	}
	return nil
}

// Set default value for message model.ApiConfig
func (this *ApiConfig) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_api_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_group_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_cluster_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_space_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_api_name(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_api_path(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_api_mode(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_api_description(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_request_method(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_response_type(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_timeout(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_datasource_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_table_name(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_script(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_created(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_updated(); err != nil {
		return err
	}
	return nil
}

func (this *ApiVersion) _xxx_xxx_Validator_Validate_version_id() error {
	if !(len(this.VersionId) == 20) {
		return protovalidator.FieldError1("ApiVersion", "the byte length of field 'version_id' must be equal to '20'", protovalidator.StringByteLenToString(this.VersionId))
	}
	if !(strings.HasPrefix(this.VersionId, "dsv-")) {
		return protovalidator.FieldError1("ApiVersion", "the value of field 'version_id' must start with string 'dsv-'", this.VersionId)
	}
	return nil
}

var _xxx_xxx_Validator_ApiVersion_InEnums_PublishStatus = map[ApiVersion_Status]bool{0: true, 1: true, 2: true}

func (this *ApiVersion) _xxx_xxx_Validator_Validate_publish_status() error {
	if !(this.PublishStatus > 1) {
		return protovalidator.FieldError1("ApiVersion", "the value of field 'publish_status' must be greater than '1'", protovalidator.Int32ToString(int32(this.PublishStatus)))
	}
	if !(_xxx_xxx_Validator_ApiVersion_InEnums_PublishStatus[this.PublishStatus]) {
		return protovalidator.FieldError1("ApiVersion", "the value of field 'publish_status' must in enums of '[0 1 2]'", protovalidator.Int32ToString(int32(this.PublishStatus)))
	}
	return nil
}

func (this *ApiVersion) _xxx_xxx_Validator_Validate_publish_time() error {
	if !(this.PublishTime > 0) {
		return protovalidator.FieldError1("ApiVersion", "the value of field 'publish_time' must be greater than '0'", protovalidator.Int64ToString(this.PublishTime))
	}
	return nil
}

func (this *ApiVersion) _xxx_xxx_Validator_Validate_api_id() error {
	if !(len(this.ApiId) == 20) {
		return protovalidator.FieldError1("ApiVersion", "the byte length of field 'api_id' must be equal to '20'", protovalidator.StringByteLenToString(this.ApiId))
	}
	if !(strings.HasPrefix(this.ApiId, "dsa-")) {
		return protovalidator.FieldError1("ApiVersion", "the value of field 'api_id' must start with string 'dsa-'", this.ApiId)
	}
	return nil
}

func (this *ApiVersion) _xxx_xxx_Validator_Validate_group_id() error {
	if !(len(this.GroupId) == 20) {
		return protovalidator.FieldError1("ApiVersion", "the byte length of field 'group_id' must be equal to '20'", protovalidator.StringByteLenToString(this.GroupId))
	}
	if !(strings.HasPrefix(this.GroupId, "dsg-")) {
		return protovalidator.FieldError1("ApiVersion", "the value of field 'group_id' must start with string 'dsg-'", this.GroupId)
	}
	return nil
}

func (this *ApiVersion) _xxx_xxx_Validator_Validate_cluster_id() error {
	if !(len(this.ClusterId) == 20) {
		return protovalidator.FieldError1("ApiVersion", "the byte length of field 'cluster_id' must be equal to '20'", protovalidator.StringByteLenToString(this.ClusterId))
	}
	if !(strings.HasPrefix(this.ClusterId, "dsc-")) {
		return protovalidator.FieldError1("ApiVersion", "the value of field 'cluster_id' must start with string 'dsc-'", this.ClusterId)
	}
	return nil
}

func (this *ApiVersion) _xxx_xxx_Validator_Validate_space_id() error {
	if !(len(this.SpaceId) == 20) {
		return protovalidator.FieldError1("ApiVersion", "the byte length of field 'space_id' must be equal to '20'", protovalidator.StringByteLenToString(this.SpaceId))
	}
	if !(strings.HasPrefix(this.SpaceId, "wks-")) {
		return protovalidator.FieldError1("ApiVersion", "the value of field 'space_id' must start with string 'wks-'", this.SpaceId)
	}
	return nil
}

func (this *ApiVersion) _xxx_xxx_Validator_Validate_api_name() error {
	if !(len(this.ApiName) > 0) {
		return protovalidator.FieldError1("ApiVersion", "the byte length of field 'api_name' must be greater than '0'", protovalidator.StringByteLenToString(this.ApiName))
	}
	if !(len(this.ApiName) <= 64) {
		return protovalidator.FieldError1("ApiVersion", "the byte length of field 'api_name' must be less than or equal to '64'", protovalidator.StringByteLenToString(this.ApiName))
	}
	return nil
}

func (this *ApiVersion) _xxx_xxx_Validator_Validate_api_path() error {
	if !(len(this.ApiPath) > 0) {
		return protovalidator.FieldError1("ApiVersion", "the byte length of field 'api_path' must be greater than '0'", protovalidator.StringByteLenToString(this.ApiPath))
	}
	if !(len(this.ApiPath) <= 200) {
		return protovalidator.FieldError1("ApiVersion", "the byte length of field 'api_path' must be less than or equal to '200'", protovalidator.StringByteLenToString(this.ApiPath))
	}
	return nil
}

var _xxx_xxx_Validator_ApiVersion_InEnums_ApiMode = map[ApiVersion_ApiMode]bool{0: true, 1: true, 2: true}

func (this *ApiVersion) _xxx_xxx_Validator_Validate_api_mode() error {
	if !(this.ApiMode > 0) {
		return protovalidator.FieldError1("ApiVersion", "the value of field 'api_mode' must be greater than '0'", protovalidator.Int32ToString(int32(this.ApiMode)))
	}
	if !(_xxx_xxx_Validator_ApiVersion_InEnums_ApiMode[this.ApiMode]) {
		return protovalidator.FieldError1("ApiVersion", "the value of field 'api_mode' must in enums of '[0 1 2]'", protovalidator.Int32ToString(int32(this.ApiMode)))
	}
	return nil
}

func (this *ApiVersion) _xxx_xxx_Validator_Validate_api_description() error {
	if !(len(this.ApiDescription) <= 257) {
		return protovalidator.FieldError1("ApiVersion", "the byte length of field 'api_description' must be less than or equal to '257'", protovalidator.StringByteLenToString(this.ApiDescription))
	}
	return nil
}

var _xxx_xxx_Validator_ApiVersion_InEnums_RequestMethod = map[ApiVersion_RequestMethod]bool{0: true, 1: true, 2: true}

func (this *ApiVersion) _xxx_xxx_Validator_Validate_request_method() error {
	if !(this.RequestMethod > 0) {
		return protovalidator.FieldError1("ApiVersion", "the value of field 'request_method' must be greater than '0'", protovalidator.Int32ToString(int32(this.RequestMethod)))
	}
	if !(_xxx_xxx_Validator_ApiVersion_InEnums_RequestMethod[this.RequestMethod]) {
		return protovalidator.FieldError1("ApiVersion", "the value of field 'request_method' must in enums of '[0 1 2]'", protovalidator.Int32ToString(int32(this.RequestMethod)))
	}
	return nil
}

var _xxx_xxx_Validator_ApiVersion_InEnums_ResponseType = map[ApiVersion_ResponseType]bool{0: true, 1: true, 2: true}

func (this *ApiVersion) _xxx_xxx_Validator_Validate_response_type() error {
	if !(this.ResponseType > 0) {
		return protovalidator.FieldError1("ApiVersion", "the value of field 'response_type' must be greater than '0'", protovalidator.Int32ToString(int32(this.ResponseType)))
	}
	if !(_xxx_xxx_Validator_ApiVersion_InEnums_ResponseType[this.ResponseType]) {
		return protovalidator.FieldError1("ApiVersion", "the value of field 'response_type' must in enums of '[0 1 2]'", protovalidator.Int32ToString(int32(this.ResponseType)))
	}
	return nil
}

func (this *ApiVersion) _xxx_xxx_Validator_Validate_timeout() error {
	if !(this.Timeout >= 1) {
		return protovalidator.FieldError1("ApiVersion", "the value of field 'timeout' must be greater than or equal to '1'", protovalidator.Int64ToString(this.Timeout))
	}
	if !(this.Timeout <= 180) {
		return protovalidator.FieldError1("ApiVersion", "the value of field 'timeout' must be less than or equal to '180'", protovalidator.Int64ToString(this.Timeout))
	}
	return nil
}

func (this *ApiVersion) _xxx_xxx_Validator_Validate_datasource_id() error {
	if !(len(this.DatasourceId) == 20) {
		return protovalidator.FieldError1("ApiVersion", "the byte length of field 'datasource_id' must be equal to '20'", protovalidator.StringByteLenToString(this.DatasourceId))
	}
	if !(strings.HasPrefix(this.DatasourceId, "som-")) {
		return protovalidator.FieldError1("ApiVersion", "the value of field 'datasource_id' must start with string 'som-'", this.DatasourceId)
	}
	return nil
}

func (this *ApiVersion) _xxx_xxx_Validator_Validate_table_name() error {
	if !(len(this.TableName) > 0) {
		return protovalidator.FieldError1("ApiVersion", "the byte length of field 'table_name' must be greater than '0'", protovalidator.StringByteLenToString(this.TableName))
	}
	if !(len(this.TableName) <= 64) {
		return protovalidator.FieldError1("ApiVersion", "the byte length of field 'table_name' must be less than or equal to '64'", protovalidator.StringByteLenToString(this.TableName))
	}
	return nil
}

func (this *ApiVersion) _xxx_xxx_Validator_Validate_script() error {
	if !(len(this.Script) <= 20000) {
		return protovalidator.FieldError1("ApiVersion", "the byte length of field 'script' must be less than or equal to '20000'", protovalidator.StringByteLenToString(this.Script))
	}
	return nil
}

// Set default value for message model.ApiVersion
func (this *ApiVersion) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_version_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_publish_status(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_publish_time(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_api_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_group_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_cluster_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_space_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_api_name(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_api_path(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_api_mode(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_api_description(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_request_method(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_response_type(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_timeout(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_datasource_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_table_name(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_script(); err != nil {
		return err
	}
	return nil
}

func (this *ApiRequestParams) _xxx_xxx_Validator_Validate_param_id() error {
	if !(len(this.ParamId) == 20) {
		return protovalidator.FieldError1("ApiRequestParams", "the byte length of field 'param_id' must be equal to '20'", protovalidator.StringByteLenToString(this.ParamId))
	}
	if !(strings.HasPrefix(this.ParamId, "req-")) {
		return protovalidator.FieldError1("ApiRequestParams", "the value of field 'param_id' must start with string 'req-'", this.ParamId)
	}
	return nil
}

func (this *ApiRequestParams) _xxx_xxx_Validator_Validate_api_id() error {
	if !(len(this.ApiId) == 20) {
		return protovalidator.FieldError1("ApiRequestParams", "the byte length of field 'api_id' must be equal to '20'", protovalidator.StringByteLenToString(this.ApiId))
	}
	if !(strings.HasPrefix(this.ApiId, "dsa-")) {
		return protovalidator.FieldError1("ApiRequestParams", "the value of field 'api_id' must start with string 'dsa-'", this.ApiId)
	}
	return nil
}

func (this *ApiRequestParams) _xxx_xxx_Validator_Validate_column_name() error {
	if !(len(this.ColumnName) > 0) {
		return protovalidator.FieldError1("ApiRequestParams", "the byte length of field 'column_name' must be greater than '0'", protovalidator.StringByteLenToString(this.ColumnName))
	}
	if !(len(this.ColumnName) <= 64) {
		return protovalidator.FieldError1("ApiRequestParams", "the byte length of field 'column_name' must be less than or equal to '64'", protovalidator.StringByteLenToString(this.ColumnName))
	}
	return nil
}

func (this *ApiRequestParams) _xxx_xxx_Validator_Validate_default_value() error {
	if !(len(this.DefaultValue) <= 257) {
		return protovalidator.FieldError1("ApiRequestParams", "the byte length of field 'default_value' must be less than or equal to '257'", protovalidator.StringByteLenToString(this.DefaultValue))
	}
	return nil
}

func (this *ApiRequestParams) _xxx_xxx_Validator_Validate_example_value() error {
	if !(len(this.ExampleValue) <= 257) {
		return protovalidator.FieldError1("ApiRequestParams", "the byte length of field 'example_value' must be less than or equal to '257'", protovalidator.StringByteLenToString(this.ExampleValue))
	}
	return nil
}

var _xxx_xxx_Validator_ApiRequestParams_InEnums_DataType = map[ParameterDataType]bool{0: true, 1: true, 2: true, 3: true, 4: true}

func (this *ApiRequestParams) _xxx_xxx_Validator_Validate_data_type() error {
	if !(this.DataType > 0) {
		return protovalidator.FieldError1("ApiRequestParams", "the value of field 'data_type' must be greater than '0'", protovalidator.Int32ToString(int32(this.DataType)))
	}
	if !(_xxx_xxx_Validator_ApiRequestParams_InEnums_DataType[this.DataType]) {
		return protovalidator.FieldError1("ApiRequestParams", "the value of field 'data_type' must in enums of '[0 1 2 3 4]'", protovalidator.Int32ToString(int32(this.DataType)))
	}
	return nil
}

func (this *ApiRequestParams) _xxx_xxx_Validator_Validate_param_description() error {
	if !(len(this.ParamDescription) <= 257) {
		return protovalidator.FieldError1("ApiRequestParams", "the byte length of field 'param_description' must be less than or equal to '257'", protovalidator.StringByteLenToString(this.ParamDescription))
	}
	return nil
}

func (this *ApiRequestParams) _xxx_xxx_Validator_Validate_param_name() error {
	if !(len(this.ParamName) > 0) {
		return protovalidator.FieldError1("ApiRequestParams", "the byte length of field 'param_name' must be greater than '0'", protovalidator.StringByteLenToString(this.ParamName))
	}
	if !(len(this.ParamName) <= 64) {
		return protovalidator.FieldError1("ApiRequestParams", "the byte length of field 'param_name' must be less than or equal to '64'", protovalidator.StringByteLenToString(this.ParamName))
	}
	return nil
}

var _xxx_xxx_Validator_ApiRequestParams_InEnums_ParamOperator = map[ApiRequestParams_ParameterOperator]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true}

func (this *ApiRequestParams) _xxx_xxx_Validator_Validate_param_operator() error {
	if !(this.ParamOperator > 0) {
		return protovalidator.FieldError1("ApiRequestParams", "the value of field 'param_operator' must be greater than '0'", protovalidator.Int32ToString(int32(this.ParamOperator)))
	}
	if !(_xxx_xxx_Validator_ApiRequestParams_InEnums_ParamOperator[this.ParamOperator]) {
		return protovalidator.FieldError1("ApiRequestParams", "the value of field 'param_operator' must in enums of '[0 1 2 3 4 5 6 7]'", protovalidator.Int32ToString(int32(this.ParamOperator)))
	}
	return nil
}

var _xxx_xxx_Validator_ApiRequestParams_InEnums_ParamPosition = map[ApiRequestParams_ParameterPosition]bool{0: true, 1: true, 2: true, 3: true, 4: true}

func (this *ApiRequestParams) _xxx_xxx_Validator_Validate_param_position() error {
	if !(this.ParamPosition > 0) {
		return protovalidator.FieldError1("ApiRequestParams", "the value of field 'param_position' must be greater than '0'", protovalidator.Int32ToString(int32(this.ParamPosition)))
	}
	if !(_xxx_xxx_Validator_ApiRequestParams_InEnums_ParamPosition[this.ParamPosition]) {
		return protovalidator.FieldError1("ApiRequestParams", "the value of field 'param_position' must in enums of '[0 1 2 3 4]'", protovalidator.Int32ToString(int32(this.ParamPosition)))
	}
	return nil
}

// Set default value for message model.ApiRequestParams
func (this *ApiRequestParams) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_param_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_api_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_column_name(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_default_value(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_example_value(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_data_type(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_param_description(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_param_name(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_param_operator(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_param_position(); err != nil {
		return err
	}
	return nil
}

func (this *ApiResponseParams) _xxx_xxx_Validator_Validate_param_id() error {
	if !(len(this.ParamId) == 20) {
		return protovalidator.FieldError1("ApiResponseParams", "the byte length of field 'param_id' must be equal to '20'", protovalidator.StringByteLenToString(this.ParamId))
	}
	if !(strings.HasPrefix(this.ParamId, "rsp-")) {
		return protovalidator.FieldError1("ApiResponseParams", "the value of field 'param_id' must start with string 'rsp-'", this.ParamId)
	}
	return nil
}

func (this *ApiResponseParams) _xxx_xxx_Validator_Validate_api_id() error {
	if !(len(this.ApiId) == 20) {
		return protovalidator.FieldError1("ApiResponseParams", "the byte length of field 'api_id' must be equal to '20'", protovalidator.StringByteLenToString(this.ApiId))
	}
	if !(strings.HasPrefix(this.ApiId, "dsa-")) {
		return protovalidator.FieldError1("ApiResponseParams", "the value of field 'api_id' must start with string 'dsa-'", this.ApiId)
	}
	return nil
}

func (this *ApiResponseParams) _xxx_xxx_Validator_Validate_column_name() error {
	if !(len(this.ColumnName) > 0) {
		return protovalidator.FieldError1("ApiResponseParams", "the byte length of field 'column_name' must be greater than '0'", protovalidator.StringByteLenToString(this.ColumnName))
	}
	if !(len(this.ColumnName) <= 64) {
		return protovalidator.FieldError1("ApiResponseParams", "the byte length of field 'column_name' must be less than or equal to '64'", protovalidator.StringByteLenToString(this.ColumnName))
	}
	return nil
}

func (this *ApiResponseParams) _xxx_xxx_Validator_Validate_default_value() error {
	if !(len(this.DefaultValue) <= 257) {
		return protovalidator.FieldError1("ApiResponseParams", "the byte length of field 'default_value' must be less than or equal to '257'", protovalidator.StringByteLenToString(this.DefaultValue))
	}
	return nil
}

func (this *ApiResponseParams) _xxx_xxx_Validator_Validate_example_value() error {
	if !(len(this.ExampleValue) <= 257) {
		return protovalidator.FieldError1("ApiResponseParams", "the byte length of field 'example_value' must be less than or equal to '257'", protovalidator.StringByteLenToString(this.ExampleValue))
	}
	return nil
}

var _xxx_xxx_Validator_ApiResponseParams_InEnums_DataType = map[ParameterDataType]bool{0: true, 1: true, 2: true, 3: true, 4: true}

func (this *ApiResponseParams) _xxx_xxx_Validator_Validate_data_type() error {
	if !(this.DataType > 0) {
		return protovalidator.FieldError1("ApiResponseParams", "the value of field 'data_type' must be greater than '0'", protovalidator.Int32ToString(int32(this.DataType)))
	}
	if !(_xxx_xxx_Validator_ApiResponseParams_InEnums_DataType[this.DataType]) {
		return protovalidator.FieldError1("ApiResponseParams", "the value of field 'data_type' must in enums of '[0 1 2 3 4]'", protovalidator.Int32ToString(int32(this.DataType)))
	}
	return nil
}

func (this *ApiResponseParams) _xxx_xxx_Validator_Validate_param_description() error {
	if !(len(this.ParamDescription) <= 257) {
		return protovalidator.FieldError1("ApiResponseParams", "the byte length of field 'param_description' must be less than or equal to '257'", protovalidator.StringByteLenToString(this.ParamDescription))
	}
	return nil
}

func (this *ApiResponseParams) _xxx_xxx_Validator_Validate_param_name() error {
	if !(len(this.ParamName) > 0) {
		return protovalidator.FieldError1("ApiResponseParams", "the byte length of field 'param_name' must be greater than '0'", protovalidator.StringByteLenToString(this.ParamName))
	}
	if !(len(this.ParamName) <= 64) {
		return protovalidator.FieldError1("ApiResponseParams", "the byte length of field 'param_name' must be less than or equal to '64'", protovalidator.StringByteLenToString(this.ParamName))
	}
	return nil
}

// Set default value for message model.ApiResponseParams
func (this *ApiResponseParams) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_param_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_api_id(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_column_name(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_default_value(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_example_value(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_data_type(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_param_description(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_param_name(); err != nil {
		return err
	}
	return nil
}
