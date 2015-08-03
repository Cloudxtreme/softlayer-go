package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Resource_Group_Member_Hardware - <nil>
type SoftLayer_Resource_Group_Member_Hardware struct {
}

// SoftLayer_Resource_Group_Member_Hardware_Extended is SoftLayer_Resource_Group_Member_Hardware with all maskable types.
type SoftLayer_Resource_Group_Member_Hardware_Extended struct {
	SoftLayer_Resource_Group_Member_Hardware

	// ServerHidden - A resource group hardware member's associated server hidden state.
	ServerHidden *SoftLayer_Resource_Group_Member_Attribute `json:"serverHidden"`

	// ServerPriority - A resource group hardware member's associated server priority.
	ServerPriority *SoftLayer_Resource_Group_Member_Attribute `json:"serverPriority"`

	// ServerSlaveDelay - A resource group hardware member's associated server slave delay (in seconds).
	ServerSlaveDelay *SoftLayer_Resource_Group_Member_Attribute `json:"serverSlaveDelay"`

	// ServerTags - A resource group hardware member's associated server tags (in format).
	ServerTags *SoftLayer_Resource_Group_Member_Attribute `json:"serverTags"`

	// ServerVotes - A resource group hardware member's associated server vote count.
	ServerVotes *SoftLayer_Resource_Group_Member_Attribute `json:"serverVotes"`

	// Resource - no documentation
	Resource *SoftLayer_Hardware `json:"resource"`

	// ServerArbiterOnly - A resource group hardware member's associated server arbiter-only state.
	ServerArbiterOnly *SoftLayer_Resource_Group_Member_Attribute `json:"serverArbiterOnly"`
}

func (softlayer_resource_group_member_hardware *SoftLayer_Resource_Group_Member_Hardware) String() string {
	return "SoftLayer_Resource_Group_Member_Hardware"
}
