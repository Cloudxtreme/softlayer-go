package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Account_Regional_Registry_Detail_Type - Subnet Registration Detail Type objects describe
// the nature of a [[SoftLayer_Account_Regional_Registry_Detail]] object. The standard values for these
// objects are as follows: - The detail object represents the information for a
// [[SoftLayer_Network_Subnet|subnet]] NETWORK6 - The detail object represents the information for an
// [[SoftLayer_Network_Subnet_Version6|IPv6 subnet]] - The detail object represents the information for
// a customer with the
type SoftLayer_Account_Regional_Registry_Detail_Type struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Id - no documentation
	Id int `json:"id"`

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`
}

func (softlayer_account_regional_registry_detail_type *SoftLayer_Account_Regional_Registry_Detail_Type) String() string {
	return "SoftLayer_Account_Regional_Registry_Detail_Type"
}

// GetAllObjects - <nil>
func (softlayer_account_regional_registry_detail_type *SoftLayer_Account_Regional_Registry_Detail_Type) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_Account_Regional_Registry_Detail_Type, error) {
	var returnValue []*SoftLayer_Account_Regional_Registry_Detail_Type
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_account_regional_registry_detail_type *SoftLayer_Account_Regional_Registry_Detail_Type) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Account_Regional_Registry_Detail_Type, error) {
	var returnValue *SoftLayer_Account_Regional_Registry_Detail_Type
	return returnValue, nil
}
