package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Regional_Registry_Detail - <nil>
type SoftLayer_Account_Regional_Registry_Detail struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// DetailTypeId - The detail object's associated
	// [[SoftLayer_Account_Regional_Registry_Detail_Type|type]] id
	DetailTypeId int `json:"detailTypeId"`

	// RegionalInternetRegistryHandleId - The detail object's associated
	// [[SoftLayer_Account_Rwhois_Handle|RIR handle]] id
	RegionalInternetRegistryHandleId int `json:"regionalInternetRegistryHandleId"`

	// AccountId - The detail object's associated [[SoftLayer_Account|account]] id
	AccountId int `json:"accountId"`

	// Id - no documentation
	Id int `json:"id"`

	// ModifyDate - The date and time the detail object was last modified
	ModifyDate *time.Time `json:"modifyDate"`
}

// SoftLayer_Account_Regional_Registry_Detail_Extended is SoftLayer_Account_Regional_Registry_Detail with all maskable types.
type SoftLayer_Account_Regional_Registry_Detail_Extended struct {
	SoftLayer_Account_Regional_Registry_Detail

	// PropertyCount - A count of the individual properties that define this detail object's values.
	PropertyCount uint64 `json:"propertyCount"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// DetailType - no documentation
	DetailType *SoftLayer_Account_Regional_Registry_Detail_Type `json:"detailType"`

	// RegionalInternetRegistryHandle - The associated RWhois handle of this detail object. Used only when
	// detailed reassignments are necessary.
	RegionalInternetRegistryHandle *SoftLayer_Account_Rwhois_Handle `json:"regionalInternetRegistryHandle"`

	// DetailCount - A count of references to the [[SoftLayer_Network_Subnet_Registration|registration
	// objects]] that consume this detail object.
	DetailCount uint64 `json:"detailCount"`

	// Details - References to the [[SoftLayer_Network_Subnet_Registration|registration objects]] that
	// consume this detail object.
	Details []*SoftLayer_Network_Subnet_Registration_Details `json:"details"`

	// Properties - The individual properties that define this detail object's values.
	Properties []*SoftLayer_Account_Regional_Registry_Detail_Property `json:"properties"`
}

func (softlayer_account_regional_registry_detail *SoftLayer_Account_Regional_Registry_Detail) String() string {
	return "SoftLayer_Account_Regional_Registry_Detail"
}
