package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Location_Group - <nil>
type SoftLayer_Location_Group struct {

	// SecurityLevelId - <nil>
	SecurityLevelId int `json:"securityLevelId"`

	// Description - <nil>
	Description string `json:"description"`

	// Id - <nil>
	Id int `json:"id"`

	// LocationGroupTypeId - <nil>
	LocationGroupTypeId int `json:"locationGroupTypeId"`

	// Name - <nil>
	Name string `json:"name"`
}

// SoftLayer_Location_Group_Extended is SoftLayer_Location_Group with all maskable types.
type SoftLayer_Location_Group_Extended struct {
	SoftLayer_Location_Group

	// Locations - no documentation
	Locations []*SoftLayer_Location `json:"locations"`

	// LocationCount - no documentation
	LocationCount uint64 `json:"locationCount"`

	// LocationGroupType - no documentation
	LocationGroupType *SoftLayer_Location_Group_Type `json:"locationGroupType"`
}

func (softlayer_location_group *SoftLayer_Location_Group) String() string {
	return "SoftLayer_Location_Group"
}
