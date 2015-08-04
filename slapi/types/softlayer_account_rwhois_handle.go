package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Rwhois_Handle - Provides a means of tracking handle identifiers at the various
// regional internet registries (RIRs). These objects are used by the
// [[SoftLayer_Network_Subnet_Registration (type)|SoftLayer_Network_Subnet_Registration]] objects to
// identify a customer or organization when a subnet is registered.
type SoftLayer_Account_Rwhois_Handle struct {

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// AccountId - The handle object's associated [[SoftLayer_Account|account]] id
	AccountId int `json:"accountId,omitempty"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Handle - The handle object's unique identifier as assigned by the
	Handle string `json:"handle,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`
}

func (softlayer_account_rwhois_handle *SoftLayer_Account_Rwhois_Handle) String() string {
	return "SoftLayer_Account_Rwhois_Handle"
}
