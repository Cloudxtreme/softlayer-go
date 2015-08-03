package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Location_Group - <nil>
type SoftLayer_Location_Group struct {

	// Description - <nil>
	Description string `json:"description"`

	// Id - <nil>
	Id int `json:"id"`

	// LocationCount - no documentation
	LocationCount uint64 `json:"locationCount"`

	// LocationGroupType - no documentation
	LocationGroupType *SoftLayer_Location_Group_Type `json:"locationGroupType"`

	// LocationGroupTypeId - <nil>
	LocationGroupTypeId int `json:"locationGroupTypeId"`

	// Locations - no documentation
	Locations []*SoftLayer_Location `json:"locations"`

	// Name - <nil>
	Name string `json:"name"`

	// SecurityLevelId - <nil>
	SecurityLevelId int `json:"securityLevelId"`
}

func (softlayer_location_group *SoftLayer_Location_Group) String() string {
	return "SoftLayer_Location_Group"
}
