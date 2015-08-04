package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Subnet_Registration - The subnet registration data type contains general
// information relating to a single subnet registration instance. These registration instances can be
// updated to reflect changes, and will record the changes in the
// [[SoftLayer_Network_Subnet_Registration_Event|events]].
type SoftLayer_Network_Subnet_Registration struct {

	// Cidr - no documentation
	Cidr int `json:"cidr,omitempty"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// NetworkIdentifier - no documentation
	NetworkIdentifier string `json:"networkIdentifier,omitempty"`

	// RegionalInternetRegistryHandleId - The registration object's associated
	// [[SoftLayer_Account_Rwhois_Handle|RIR handle]] id
	RegionalInternetRegistryHandleId int `json:"regionalInternetRegistryHandleId,omitempty"`

	// RegionalInternetRegistryId - The registration object's associated
	// [[SoftLayer_Network_Regional_Internet_Registry|RIR]] id
	RegionalInternetRegistryId int `json:"regionalInternetRegistryId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// NetworkHandle - The RIR-specific handle or name of the registered subnet. This field is read-only.
	NetworkHandle string `json:"networkHandle,omitempty"`

	// StatusId - The registration object's associated
	// [[SoftLayer_Network_Subnet_Registration_Status|status]] id
	StatusId int `json:"statusId,omitempty"`

	// AccountId - The registration object's associated [[SoftLayer_Account|account]] id
	AccountId int `json:"accountId,omitempty"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Events - no documentation
	Events []*SoftLayer_Network_Subnet_Registration_Event `json:"events,omitempty"`

	// NetworkDetail - no documentation
	NetworkDetail *SoftLayer_Account_Regional_Registry_Detail `json:"networkDetail,omitempty"`

	// Status - no documentation
	Status *SoftLayer_Network_Subnet_Registration_Status `json:"status,omitempty"`

	// DetailReferences - The cross-reference records that tie the
	// [[SoftLayer_Account_Regional_Registry_Detail]] objects to the registration object.
	DetailReferences []*SoftLayer_Network_Subnet_Registration_Details `json:"detailReferences,omitempty"`

	// PersonDetail - no documentation
	PersonDetail *SoftLayer_Account_Regional_Registry_Detail `json:"personDetail,omitempty"`

	// Subnet - no documentation
	Subnet *SoftLayer_Network_Subnet `json:"subnet,omitempty"`

	// EventCount - no documentation
	EventCount uint64 `json:"eventCount,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// DetailReferenceCount - A count of the cross-reference records that tie the
	// [[SoftLayer_Account_Regional_Registry_Detail]] objects to the registration object.
	DetailReferenceCount uint64 `json:"detailReferenceCount,omitempty"`

	// RegionalInternetRegistry - no documentation
	RegionalInternetRegistry *SoftLayer_Network_Regional_Internet_Registry `json:"regionalInternetRegistry,omitempty"`

	// RegionalInternetRegistryHandle - The RIR handle that this registration object belongs to. This field
	// may not be populated until the registration is complete.
	RegionalInternetRegistryHandle *SoftLayer_Account_Rwhois_Handle `json:"regionalInternetRegistryHandle,omitempty"`
}

func (softlayer_network_subnet_registration *SoftLayer_Network_Subnet_Registration) String() string {
	return "SoftLayer_Network_Subnet_Registration"
}
