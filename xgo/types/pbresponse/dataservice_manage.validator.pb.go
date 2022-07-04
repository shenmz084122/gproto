// Code generated by protoc-gen-govalidator. DO NOT EDIT.
// versions:
// 		protoc-gen-govalidator 0.0.1
// source: proto/types/response/dataservice_manage.proto

package pbresponse

import (
	_ "github.com/DataWorkbench/gproto/xgo/types/pbmodel"
)

func (this *ListDataServiceClusters) _xxx_xxx_Validator_Validate_infos() error {
	for _, item := range this.Infos {
		_ = item // To avoid unused panics.
		if dt, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := dt.Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}

// Set default value for message response.ListDataServiceClusters
func (this *ListDataServiceClusters) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_infos(); err != nil {
		return err
	}
	return nil
}

func (this *DescribeDataServiceCluster) _xxx_xxx_Validator_Validate_info() error {
	if dt, ok := interface{}(this.Info).(interface{ Validate() error }); ok {
		if err := dt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Set default value for message response.DescribeDataServiceCluster
func (this *DescribeDataServiceCluster) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_info(); err != nil {
		return err
	}
	return nil
}

// Set default value for message response.CreateDataServiceCluster
func (this *CreateDataServiceCluster) Validate() error {
	if this == nil {
		return nil
	}
	return nil
}

func (this *ListApiGroups) _xxx_xxx_Validator_Validate_infos() error {
	for _, item := range this.Infos {
		_ = item // To avoid unused panics.
		if dt, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := dt.Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}

// Set default value for message response.ListApiGroups
func (this *ListApiGroups) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_infos(); err != nil {
		return err
	}
	return nil
}

// Set default value for message response.CreateApiGroup
func (this *CreateApiGroup) Validate() error {
	if this == nil {
		return nil
	}
	return nil
}

func (this *ListApiConfigs) _xxx_xxx_Validator_Validate_infos() error {
	for _, item := range this.Infos {
		_ = item // To avoid unused panics.
		if dt, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := dt.Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}

// Set default value for message response.ListApiConfigs
func (this *ListApiConfigs) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_infos(); err != nil {
		return err
	}
	return nil
}

func (this *DescribeApiConfig) _xxx_xxx_Validator_Validate_api_config() error {
	if dt, ok := interface{}(this.ApiConfig).(interface{ Validate() error }); ok {
		if err := dt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (this *DescribeApiConfig) _xxx_xxx_Validator_Validate_data_source() error {
	if dt, ok := interface{}(this.DataSource).(interface{ Validate() error }); ok {
		if err := dt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (this *DescribeApiConfig) _xxx_xxx_Validator_Validate_api_group() error {
	if dt, ok := interface{}(this.ApiGroup).(interface{ Validate() error }); ok {
		if err := dt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (this *DescribeApiConfig) _xxx_xxx_Validator_Validate_service_cluster() error {
	if dt, ok := interface{}(this.ServiceCluster).(interface{ Validate() error }); ok {
		if err := dt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Set default value for message response.DescribeApiConfig
func (this *DescribeApiConfig) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_api_config(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_data_source(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_api_group(); err != nil {
		return err
	}
	if err := this._xxx_xxx_Validator_Validate_service_cluster(); err != nil {
		return err
	}
	return nil
}

// Set default value for message response.CreateApiConfig
func (this *CreateApiConfig) Validate() error {
	if this == nil {
		return nil
	}
	return nil
}

func (this *DescribeDataServiceApiVersion) _xxx_xxx_Validator_Validate_api_version() error {
	if dt, ok := interface{}(this.ApiVersion).(interface{ Validate() error }); ok {
		if err := dt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Set default value for message response.DescribeDataServiceApiVersion
func (this *DescribeDataServiceApiVersion) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_api_version(); err != nil {
		return err
	}
	return nil
}

func (this *ListDataServiceApiVersions) _xxx_xxx_Validator_Validate_infos() error {
	for _, item := range this.Infos {
		_ = item // To avoid unused panics.
		if dt, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := dt.Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}

// Set default value for message response.ListDataServiceApiVersions
func (this *ListDataServiceApiVersions) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_infos(); err != nil {
		return err
	}
	return nil
}

func (this *ListPublishedApiVersionsByClusterId) _xxx_xxx_Validator_Validate_infos() error {
	for _, item := range this.Infos {
		_ = item // To avoid unused panics.
		if dt, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := dt.Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}

// Set default value for message response.ListPublishedApiVersionsByClusterId
func (this *ListPublishedApiVersionsByClusterId) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_infos(); err != nil {
		return err
	}
	return nil
}

// Set default value for message response.TestDataServiceApi
func (this *TestDataServiceApi) Validate() error {
	if this == nil {
		return nil
	}
	return nil
}

func (this *DescribeServiceDataSourceKinds) _xxx_xxx_Validator_Validate_Kinds() error {
	for _, item := range this.Kinds {
		_ = item // To avoid unused panics.
		if dt, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := dt.Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}

// Set default value for message response.DescribeServiceDataSourceKinds
func (this *DescribeServiceDataSourceKinds) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_Kinds(); err != nil {
		return err
	}
	return nil
}
