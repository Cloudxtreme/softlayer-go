package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Container_Network_Bandwidth_Data_Summary -
// SoftLayer_Container_Network_Bandwidth_Data_Summary models an interface's overall bandwidth usage
// during it's current billing cycle.
type SoftLayer_Container_Network_Bandwidth_Data_Summary struct {

	// ProjectedUsage - The amount of bandwidth a server is projected to use within its billing period,
	// based on it's current usage.
	ProjectedUsage slapi.Float64 `json:"projectedUsage,omitempty"`

	// UsageUnits - The unit of measurement used in a bandwidth data summary.
	UsageUnits string `json:"usageUnits,omitempty"`

	// AllowedUsage - The amount of bandwidth a server has allocated to it in it's current billing period.
	AllowedUsage slapi.Float64 `json:"allowedUsage,omitempty"`

	// EstimatedUsage - The amount of bandwidth that a server has used within it's current billing period.
	EstimatedUsage slapi.Float64 `json:"estimatedUsage,omitempty"`
}

func (softlayer_container_network_bandwidth_data_summary *SoftLayer_Container_Network_Bandwidth_Data_Summary) String() string {
	return "SoftLayer_Container_Network_Bandwidth_Data_Summary"
}
