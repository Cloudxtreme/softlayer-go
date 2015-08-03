package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

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

func (softlayer_network_backbone_location_dependent *SoftLayer_Network_Backbone_Location_Dependent) String() string {
	return "SoftLayer_Network_Backbone_Location_Dependent"
}
