package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Location_Group_Regional - <nil>
type SoftLayer_Location_Group_Regional struct {

	// LocationGroupTypeId - <nil>
	LocationGroupTypeId int `json:"locationGroupTypeId,omitempty"`

	// Name - <nil>
	Name string `json:"name,omitempty"`

	// SecurityLevelId - <nil>
	SecurityLevelId int `json:"securityLevelId,omitempty"`

	// Description - <nil>
	Description string `json:"description,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`
}

func (softlayer_location_group_regional *SoftLayer_Location_Group_Regional) String() string {
	return "SoftLayer_Location_Group_Regional"
}

// SoftLayer_Location_Group_Regional_Extended is SoftLayer_Location_Group_Regional with all maskable types.
type SoftLayer_Location_Group_Regional_Extended struct {
	SoftLayer_Location_Group_Regional

	// Datacenters - no documentation
	Datacenters []*SoftLayer_Location `json:"datacenters,omitempty"`

	// DatacenterCount - no documentation
	DatacenterCount uint64 `json:"datacenterCount,omitempty"`

	// LocationCount - no documentation
	LocationCount uint64 `json:"locationCount,omitempty"`

	// LocationGroupType - no documentation
	LocationGroupType *SoftLayer_Location_Group_Type `json:"locationGroupType,omitempty"`

	// PreferredDatacenter - no documentation
	PreferredDatacenter *SoftLayer_Location_Datacenter `json:"preferredDatacenter,omitempty"`

	// Locations - no documentation
	Locations []*SoftLayer_Location `json:"locations,omitempty"`
}

func (softlayer_location_group_regional *SoftLayer_Location_Group_Regional_Extended) String() string {
	return "SoftLayer_Location_Group_Regional"
}
