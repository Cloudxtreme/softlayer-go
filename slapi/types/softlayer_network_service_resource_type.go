package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Service_Resource_Type - <nil>
type SoftLayer_Network_Service_Resource_Type struct {

	// Type - <nil>
	Type string `json:"type,omitempty"`
}

func (softlayer_network_service_resource_type *SoftLayer_Network_Service_Resource_Type) String() string {
	return "SoftLayer_Network_Service_Resource_Type"
}

// SoftLayer_Network_Service_Resource_Type_Extended is SoftLayer_Network_Service_Resource_Type with all maskable types.
type SoftLayer_Network_Service_Resource_Type_Extended struct {
	SoftLayer_Network_Service_Resource_Type

	// ServiceResources - <nil>
	ServiceResources []*SoftLayer_Network_Service_Resource `json:"serviceResources,omitempty"`

	// ServiceResourceCount - no documentation
	ServiceResourceCount uint64 `json:"serviceResourceCount,omitempty"`
}

func (softlayer_network_service_resource_type *SoftLayer_Network_Service_Resource_Type_Extended) String() string {
	return "SoftLayer_Network_Service_Resource_Type"
}
