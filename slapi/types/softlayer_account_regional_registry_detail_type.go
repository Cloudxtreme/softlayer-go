package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Regional_Registry_Detail_Type - Subnet Registration Detail Type objects describe
// the nature of a [[SoftLayer_Account_Regional_Registry_Detail]] object. The standard values for these
// objects are as follows: - The detail object represents the information for a
// [[SoftLayer_Network_Subnet|subnet]] NETWORK6 - The detail object represents the information for an
// [[SoftLayer_Network_Subnet_Version6|IPv6 subnet]] - The detail object represents the information for
// a customer with the
type SoftLayer_Account_Regional_Registry_Detail_Type struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// KeyName - no documentation
	KeyName string `json:"keyName,omitempty"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`
}

func (softlayer_account_regional_registry_detail_type *SoftLayer_Account_Regional_Registry_Detail_Type) String() string {
	return "SoftLayer_Account_Regional_Registry_Detail_Type"
}
