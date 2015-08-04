package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Service_Resource_Attribute - <nil>
type SoftLayer_Network_Service_Resource_Attribute struct {

	// Value - <nil>
	Value string `json:"value,omitempty"`

	// AttributeType - <nil>
	AttributeType *SoftLayer_Network_Service_Resource_Attribute_Type `json:"attributeType,omitempty"`

	// ServiceResource - <nil>
	ServiceResource *SoftLayer_Network_Service_Resource `json:"serviceResource,omitempty"`
}

func (softlayer_network_service_resource_attribute *SoftLayer_Network_Service_Resource_Attribute) String() string {
	return "SoftLayer_Network_Service_Resource_Attribute"
}
