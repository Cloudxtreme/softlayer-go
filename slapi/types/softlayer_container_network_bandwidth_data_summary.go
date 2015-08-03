package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Bandwidth_Data_Summary -
// SoftLayer_Container_Network_Bandwidth_Data_Summary models an interface's overall bandwidth usage
// during it's current billing cycle.
type SoftLayer_Container_Network_Bandwidth_Data_Summary struct {

	// AllowedUsage - The amount of bandwidth a server has allocated to it in it's current billing period.
	AllowedUsage float32 `json:"allowedUsage"`

	// EstimatedUsage - The amount of bandwidth that a server has used within it's current billing period.
	EstimatedUsage float32 `json:"estimatedUsage"`

	// ProjectedUsage - The amount of bandwidth a server is projected to use within its billing period,
	// based on it's current usage.
	ProjectedUsage float32 `json:"projectedUsage"`

	// UsageUnits - The unit of measurement used in a bandwidth data summary.
	UsageUnits string `json:"usageUnits"`
}

func (softlayer_container_network_bandwidth_data_summary *SoftLayer_Container_Network_Bandwidth_Data_Summary) String() string {
	return "SoftLayer_Container_Network_Bandwidth_Data_Summary"
}
