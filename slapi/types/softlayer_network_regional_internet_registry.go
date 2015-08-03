package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Regional_Internet_Registry - Regional Internet Registries are the organizations
// who delegate IP address blocks to other groups or organizations around the Internet. The information
// contained in this data type is used throughout the networking-related services in our systems.
type SoftLayer_Network_Regional_Internet_Registry struct {

	// Id - no documentation
	Id int `json:"id"`

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// Name - no documentation
	Name string `json:"name"`
}

func (softlayer_network_regional_internet_registry *SoftLayer_Network_Regional_Internet_Registry) String() string {
	return "SoftLayer_Network_Regional_Internet_Registry"
}
