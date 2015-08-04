package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Bandwidth_Version1_Usage_Detail_Type - The
// SoftLayer_Network_Bandwidth_Version1_Usage_Detail_Type data type contains generic information
// relating to the types of bandwidth records available, currently just public and private.
type SoftLayer_Network_Bandwidth_Version1_Usage_Detail_Type struct {

	// Alias - Database key associated with this bandwidth detail type.
	Alias string `json:"alias,omitempty"`
}

func (softlayer_network_bandwidth_version1_usage_detail_type *SoftLayer_Network_Bandwidth_Version1_Usage_Detail_Type) String() string {
	return "SoftLayer_Network_Bandwidth_Version1_Usage_Detail_Type"
}
