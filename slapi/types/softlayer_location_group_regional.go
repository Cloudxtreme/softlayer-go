package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Location_Group_Regional - <nil>
type SoftLayer_Location_Group_Regional struct {

	// SecurityLevelId - <nil>
	SecurityLevelId int `json:"securityLevelId,omitempty"`

	// Description - <nil>
	Description string `json:"description,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// LocationGroupTypeId - <nil>
	LocationGroupTypeId int `json:"locationGroupTypeId,omitempty"`

	// Name - <nil>
	Name string `json:"name,omitempty"`

	// Datacenters - no documentation
	Datacenters []*SoftLayer_Location `json:"datacenters,omitempty"`

	// PreferredDatacenter - no documentation
	PreferredDatacenter *SoftLayer_Location_Datacenter `json:"preferredDatacenter,omitempty"`

	// Locations - no documentation
	Locations []*SoftLayer_Location `json:"locations,omitempty"`

	// LocationGroupType - no documentation
	LocationGroupType *SoftLayer_Location_Group_Type `json:"locationGroupType,omitempty"`

	// DatacenterCount - no documentation
	DatacenterCount uint64 `json:"datacenterCount,omitempty"`

	// LocationCount - no documentation
	LocationCount uint64 `json:"locationCount,omitempty"`
}

func (softlayer_location_group_regional *SoftLayer_Location_Group_Regional) String() string {
	return "SoftLayer_Location_Group_Regional"
}
