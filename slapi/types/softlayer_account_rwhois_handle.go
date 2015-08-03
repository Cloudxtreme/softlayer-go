package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Rwhois_Handle - Provides a means of tracking handle identifiers at the various
// regional internet registries (RIRs). These objects are used by the
// [[SoftLayer_Network_Subnet_Registration (type)|SoftLayer_Network_Subnet_Registration]] objects to
// identify a customer or organization when a subnet is registered.
type SoftLayer_Account_Rwhois_Handle struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AccountId - The handle object's associated [[SoftLayer_Account|account]] id
	AccountId int `json:"accountId"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Handle - The handle object's unique identifier as assigned by the
	Handle string `json:"handle"`

	// Id - no documentation
	Id int `json:"id"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`
}

func (softlayer_account_rwhois_handle *SoftLayer_Account_Rwhois_Handle) String() string {
	return "SoftLayer_Account_Rwhois_Handle"
}
