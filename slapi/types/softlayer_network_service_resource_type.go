package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Service_Resource_Type - <nil>
type SoftLayer_Network_Service_Resource_Type struct {

	// ServiceResourceCount - no documentation
	ServiceResourceCount uint64 `json:"serviceResourceCount"`

	// ServiceResources - <nil>
	ServiceResources []*SoftLayer_Network_Service_Resource `json:"serviceResources"`

	// Type - <nil>
	Type string `json:"type"`
}

func (softlayer_network_service_resource_type *SoftLayer_Network_Service_Resource_Type) String() string {
	return "SoftLayer_Network_Service_Resource_Type"
}
