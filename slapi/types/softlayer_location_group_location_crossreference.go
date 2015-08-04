package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Location_Group_Location_CrossReference - <nil>
type SoftLayer_Location_Group_Location_CrossReference struct {

	// LocationGroupId - <nil>
	LocationGroupId int `json:"locationGroupId,omitempty"`

	// LocationId - <nil>
	LocationId int `json:"locationId,omitempty"`

	// Priority - If set, this is the priority of this cross reference record in the group.
	Priority int `json:"priority,omitempty"`
}

func (softlayer_location_group_location_crossreference *SoftLayer_Location_Group_Location_CrossReference) String() string {
	return "SoftLayer_Location_Group_Location_CrossReference"
}

// SoftLayer_Location_Group_Location_CrossReference_Extended is SoftLayer_Location_Group_Location_CrossReference with all maskable types.
type SoftLayer_Location_Group_Location_CrossReference_Extended struct {
	SoftLayer_Location_Group_Location_CrossReference

	// Location - <nil>
	Location *SoftLayer_Location `json:"location,omitempty"`

	// LocationGroup - <nil>
	LocationGroup *SoftLayer_Location_Group `json:"locationGroup,omitempty"`
}

func (softlayer_location_group_location_crossreference *SoftLayer_Location_Group_Location_CrossReference_Extended) String() string {
	return "SoftLayer_Location_Group_Location_CrossReference"
}
