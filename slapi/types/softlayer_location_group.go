package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

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

// GetAllObjects - <nil>
func (softlayer_location_group *SoftLayer_Location_Group) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_Location_Group, error) {
	var returnValue []*SoftLayer_Location_Group
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_location_group *SoftLayer_Location_Group) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Location_Group, error) {
	var returnValue *SoftLayer_Location_Group
	return returnValue, nil
}
