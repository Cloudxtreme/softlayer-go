package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

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

func (softlayer_product_package_server_option *SoftLayer_Product_Package_Server_Option) String() string {
	return "SoftLayer_Product_Package_Server_Option"
}
