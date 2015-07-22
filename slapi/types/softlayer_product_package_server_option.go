package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Product_Package_Server_Option - The [[SoftLayer_Product_Package_Server_Option]] data type
// contains various data points associated with package servers that can be used in selection criteria.
type SoftLayer_Product_Package_Server_Option struct {

	// CatalogId - no documentation
	CatalogId int `json:"catalogId"`

	// Description - no documentation
	Description string `json:"description"`

	// Id - no documentation
	Id int `json:"id"`

	// Type - no documentation
	Type string `json:"type"`

	// Value - no documentation
	Value string `json:"value"`
}

// GetAllOptions - This method will grab all the package server options.
func (softlayer_product_package_server_option *SoftLayer_Product_Package_Server_Option) GetAllOptions(commonOptions *slapi.CommonOptions) ([]*SoftLayer_Product_Package_Server_Option, error) {
	var returnValue []*SoftLayer_Product_Package_Server_Option
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_product_package_server_option *SoftLayer_Product_Package_Server_Option) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Product_Package_Server_Option, error) {
	var returnValue *SoftLayer_Product_Package_Server_Option
	return returnValue, nil
}

// GetOptions - This method will grab all the package server options for the specified type.
func (softlayer_product_package_server_option *SoftLayer_Product_Package_Server_Option) GetOptions(commonOptions *slapi.CommonOptions, type_ string) ([]*SoftLayer_Product_Package_Server_Option, error) {
	var returnValue []*SoftLayer_Product_Package_Server_Option
	return returnValue, nil
}
