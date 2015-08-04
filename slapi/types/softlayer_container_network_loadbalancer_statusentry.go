package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_LoadBalancer_StatusEntry - The LoadBalancer_StatusEntry object stores
// information about the current status of a particular load balancer service. It is a data container
// that cannot be edited, deleted, or saved. It is returned exclusively by the getStatus method on the
// [[SoftLayer_Network_LoadBalancer_Service]] service
type SoftLayer_Container_Network_LoadBalancer_StatusEntry struct {

	// Content - no documentation
	Content string `json:"content,omitempty"`

	// Label - no documentation
	Label string `json:"label,omitempty"`
}

func (softlayer_container_network_loadbalancer_statusentry *SoftLayer_Container_Network_LoadBalancer_StatusEntry) String() string {
	return "SoftLayer_Container_Network_LoadBalancer_StatusEntry"
}
