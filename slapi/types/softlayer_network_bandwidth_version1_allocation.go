package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Bandwidth_Version1_Allocation - The
// SoftLayer_Network_Bandwidth_Version1_Allocation data type contains general information relating to a
// single bandwidth allocation record.
type SoftLayer_Network_Bandwidth_Version1_Allocation struct {

	// Amount - no documentation
	Amount float64 `json:"amount"`

	// Id - no documentation
	Id int `json:"id"`
}

// SoftLayer_Network_Bandwidth_Version1_Allocation_Extended is SoftLayer_Network_Bandwidth_Version1_Allocation with all maskable types.
type SoftLayer_Network_Bandwidth_Version1_Allocation_Extended struct {
	SoftLayer_Network_Bandwidth_Version1_Allocation

	// AllotmentDetail - no documentation
	AllotmentDetail *SoftLayer_Network_Bandwidth_Version1_Allotment_Detail `json:"allotmentDetail"`

	// BillingItem - Billing item associated with this hardware allocation.
	BillingItem *SoftLayer_Billing_Item_Hardware `json:"billingItem"`
}

func (softlayer_network_bandwidth_version1_allocation *SoftLayer_Network_Bandwidth_Version1_Allocation) String() string {
	return "SoftLayer_Network_Bandwidth_Version1_Allocation"
}
