package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Backbone - A SoftLayer_Network_Backbone represents a single backbone connection
// from SoftLayer to the public Internet, from the Internet to the SoftLayer private network, or a link
// that connects the private networks between SoftLayer's datacenters. The SoftLayer_Network_Backbone
// data type is a collection of data associated with one of those connections.
type SoftLayer_Network_Backbone struct {

	// Capacity - The numeric portion of the bandwidth capacity of a SoftLayer backbone. For instance, if a
	// backbone is rated at "1 GigE" capacity then the capacity property of the backbone is 1.
	Capacity int `json:"capacity,omitempty"`

	// Type - Whether a SoftLayer backbone connects to the public Internet, to the private network, or
	// connecting the private networks of SoftLayer's datacenters. Type is either the string "public",
	// "private", or "private-interconnect".
	Type string `json:"type,omitempty"`

	// CapacityUnits - The unit portion of the bandwidth capacity of a SoftLayer backbone. For instance, if
	// a backbone is rated at "10 G" capacity then the capacityUnits property of the backbone is
	CapacityUnits string `json:"capacityUnits,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Name - A backbone's name. This is usually the name of the backbone's network provider followed by a
	// number in case SoftLayer uses more than one backbone from a provider. Backbone provider numbers
	// start with the number one and increment from there.
	Name string `json:"name,omitempty"`

	// NetworkComponentId - The internal identifier of the network component that backbone is connected to.
	NetworkComponentId int `json:"networkComponentId,omitempty"`

	// NetworkComponent - no documentation
	NetworkComponent *SoftLayer_Network_Component `json:"networkComponent,omitempty"`

	// Health - no documentation
	Health string `json:"health,omitempty"`

	// Location - Which of the SoftLayer datacenters a backbone is connected to.
	Location *SoftLayer_Location `json:"location,omitempty"`
}

func (softlayer_network_backbone *SoftLayer_Network_Backbone) String() string {
	return "SoftLayer_Network_Backbone"
}
