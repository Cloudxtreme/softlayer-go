package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Network_Backbone - A SoftLayer_Network_Backbone represents a single backbone connection
// from SoftLayer to the public Internet, from the Internet to the SoftLayer private network, or a link
// that connects the private networks between SoftLayer's datacenters. The SoftLayer_Network_Backbone
// data type is a collection of data associated with one of those connections.
type SoftLayer_Network_Backbone struct {

	// Capacity - The numeric portion of the bandwidth capacity of a SoftLayer backbone. For instance, if a
	// backbone is rated at "1 GigE" capacity then the capacity property of the backbone is 1.
	Capacity int `json:"capacity"`

	// CapacityUnits - The unit portion of the bandwidth capacity of a SoftLayer backbone. For instance, if
	// a backbone is rated at "10 G" capacity then the capacityUnits property of the backbone is
	CapacityUnits string `json:"capacityUnits"`

	// Health - no documentation
	Health string `json:"health"`

	// Id - no documentation
	Id int `json:"id"`

	// Location - Which of the SoftLayer datacenters a backbone is connected to.
	Location *SoftLayer_Location `json:"location"`

	// Name - A backbone's name. This is usually the name of the backbone's network provider followed by a
	// number in case SoftLayer uses more than one backbone from a provider. Backbone provider numbers
	// start with the number one and increment from there.
	Name string `json:"name"`

	// NetworkComponent - no documentation
	NetworkComponent *SoftLayer_Network_Component `json:"networkComponent"`

	// NetworkComponentId - The internal identifier of the network component that backbone is connected to.
	NetworkComponentId int `json:"networkComponentId"`

	// Type - Whether a SoftLayer backbone connects to the public Internet, to the private network, or
	// connecting the private networks of SoftLayer's datacenters. Type is either the string "public",
	// "private", or "private-interconnect".
	Type string `json:"type"`
}

func (softlayer_network_backbone *SoftLayer_Network_Backbone) String() string {
	return "SoftLayer_Network_Backbone"
}

// GetAllBackbones - Retrieve a list of all SoftLayer backbone connections. Use this method if you need
// all backbones or don't know the id number of a specific backbone.
func (softlayer_network_backbone *SoftLayer_Network_Backbone) GetAllBackbones(ctx *slapi.RequestContext) ([]*SoftLayer_Network_Backbone, error) {
	var returnValue []*SoftLayer_Network_Backbone
	return returnValue, nil
}

// GetBackbonesForLocationName - Retrieve a list of all SoftLayer backbone connections for a location
// name.
func (softlayer_network_backbone *SoftLayer_Network_Backbone) GetBackbonesForLocationName(ctx *slapi.RequestContext, locationName string) ([]*SoftLayer_Network_Backbone, error) {
	var returnValue []*SoftLayer_Network_Backbone
	return returnValue, nil
}

// GetGraphImage - Retrieve a graph of a SoftLayer backbone's last 24 hours of activity. getGraphImage
// returns a PNG image measuring 827 pixels by 293 pixels.
func (softlayer_network_backbone *SoftLayer_Network_Backbone) GetGraphImage(ctx *slapi.RequestContext) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetObject - Retrieve an individual SoftLayer_Network_Backbone record. Use the getAllBackbones()
// method to retrieve a list of all SoftLayer network backbones.
func (softlayer_network_backbone *SoftLayer_Network_Backbone) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Backbone, error) {
	var returnValue *SoftLayer_Network_Backbone
	return returnValue, nil
}
