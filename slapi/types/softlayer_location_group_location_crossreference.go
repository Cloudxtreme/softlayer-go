package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Location_Group_Location_CrossReference - <nil>
type SoftLayer_Location_Group_Location_CrossReference struct {

	// Location - <nil>
	Location *SoftLayer_Location `json:"location"`

	// LocationGroup - <nil>
	LocationGroup *SoftLayer_Location_Group `json:"locationGroup"`

	// LocationGroupId - <nil>
	LocationGroupId int `json:"locationGroupId"`

	// LocationId - <nil>
	LocationId int `json:"locationId"`

	// Priority - If set, this is the priority of this cross reference record in the group.
	Priority int `json:"priority"`
}

func (softlayer_location_group_location_crossreference *SoftLayer_Location_Group_Location_CrossReference) String() string {
	return "SoftLayer_Location_Group_Location_CrossReference"
}
