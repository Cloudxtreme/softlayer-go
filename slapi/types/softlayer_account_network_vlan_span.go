package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Network_Vlan_Span - The SoftLayer_Account_Network_Vlan_Span data type exposes the
// setting which controls the automatic spanning of private VLANs attached to a given customers
// account.
type SoftLayer_Account_Network_Vlan_Span struct {

	// EnabledFlag - Flag indicating whether the customer wishes to have all private network VLANs
	// associated with account automatically joined [0 or 1]
	EnabledFlag bool `json:"enabledFlag"`

	// Id - The unique internal identifier of the SoftLayer_Account_Network_Vlan_Span object.
	Id int `json:"id"`

	// LastAppliedDate - Timestamp of the last time the ACL for this account was applied.
	LastAppliedDate *time.Time `json:"lastAppliedDate"`

	// LastVerifiedDate - Timestamp of the last time the subnet hash was verified for this span record.
	LastVerifiedDate *time.Time `json:"lastVerifiedDate"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`
}

// SoftLayer_Account_Network_Vlan_Span_Extended is SoftLayer_Account_Network_Vlan_Span with all maskable types.
type SoftLayer_Account_Network_Vlan_Span_Extended struct {
	SoftLayer_Account_Network_Vlan_Span

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`
}

func (softlayer_account_network_vlan_span *SoftLayer_Account_Network_Vlan_Span) String() string {
	return "SoftLayer_Account_Network_Vlan_Span"
}
