package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Backbone_Location_Dependent - <nil>
type SoftLayer_Network_Backbone_Location_Dependent struct {

	// DependentLocationId - <nil>
	DependentLocationId int `json:"dependentLocationId"`

	// Id - <nil>
	Id int `json:"id"`

	// SourceLocationId - <nil>
	SourceLocationId int `json:"sourceLocationId"`
}

// SoftLayer_Network_Backbone_Location_Dependent_Extended is SoftLayer_Network_Backbone_Location_Dependent with all maskable types.
type SoftLayer_Network_Backbone_Location_Dependent_Extended struct {
	SoftLayer_Network_Backbone_Location_Dependent

	// SourceLocation - <nil>
	SourceLocation *SoftLayer_Location `json:"sourceLocation"`

	// DependentLocation - <nil>
	DependentLocation *SoftLayer_Location `json:"dependentLocation"`
}

func (softlayer_network_backbone_location_dependent *SoftLayer_Network_Backbone_Location_Dependent) String() string {
	return "SoftLayer_Network_Backbone_Location_Dependent"
}
