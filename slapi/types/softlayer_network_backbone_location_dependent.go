package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_Backbone_Location_Dependent - <nil>
type SoftLayer_Network_Backbone_Location_Dependent struct {

	// DependentLocation - <nil>
	DependentLocation *SoftLayer_Location `json:"dependentLocation"`

	// DependentLocationId - <nil>
	DependentLocationId int `json:"dependentLocationId"`

	// Id - <nil>
	Id int `json:"id"`

	// SourceLocation - <nil>
	SourceLocation *SoftLayer_Location `json:"sourceLocation"`

	// SourceLocationId - <nil>
	SourceLocationId int `json:"sourceLocationId"`
}

// GetAllObjects - <nil>
func (softlayer_network_backbone_location_dependent *SoftLayer_Network_Backbone_Location_Dependent) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_Network_Backbone_Location_Dependent, error) {
	var returnValue []*SoftLayer_Network_Backbone_Location_Dependent
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_backbone_location_dependent *SoftLayer_Network_Backbone_Location_Dependent) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Backbone_Location_Dependent, error) {
	var returnValue *SoftLayer_Network_Backbone_Location_Dependent
	return returnValue, nil
}

// GetSourceDependentsByName - <nil>
func (softlayer_network_backbone_location_dependent *SoftLayer_Network_Backbone_Location_Dependent) GetSourceDependentsByName(ctx *slapi.RequestContext, locationName string) (*SoftLayer_Location, error) {
	var returnValue *SoftLayer_Location
	return returnValue, nil
}
