package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Network_Bandwidth_Version1_Allocation - The
// SoftLayer_Network_Bandwidth_Version1_Allocation data type contains general information relating to a
// single bandwidth allocation record.
type SoftLayer_Network_Bandwidth_Version1_Allocation struct {

	// Amount - no documentation
	Amount slapi.Float64 `json:"amount,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// AllotmentDetail - no documentation
	AllotmentDetail *SoftLayer_Network_Bandwidth_Version1_Allotment_Detail `json:"allotmentDetail,omitempty"`

	// BillingItem - Billing item associated with this hardware allocation.
	BillingItem *SoftLayer_Billing_Item_Hardware `json:"billingItem,omitempty"`
}

func (softlayer_network_bandwidth_version1_allocation *SoftLayer_Network_Bandwidth_Version1_Allocation) String() string {
	return "SoftLayer_Network_Bandwidth_Version1_Allocation"
}
