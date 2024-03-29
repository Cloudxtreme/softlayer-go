package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Backbone_Location_Dependent - <nil>
type SoftLayer_Network_Backbone_Location_Dependent struct {

	// DependentLocationId - <nil>
	DependentLocationId int `json:"dependentLocationId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// SourceLocationId - <nil>
	SourceLocationId int `json:"sourceLocationId,omitempty"`

	// DependentLocation - <nil>
	DependentLocation *SoftLayer_Location `json:"dependentLocation,omitempty"`

	// SourceLocation - <nil>
	SourceLocation *SoftLayer_Location `json:"sourceLocation,omitempty"`
}

func (softlayer_network_backbone_location_dependent *SoftLayer_Network_Backbone_Location_Dependent) String() string {
	return "SoftLayer_Network_Backbone_Location_Dependent"
}
