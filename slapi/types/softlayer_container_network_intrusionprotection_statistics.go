package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_IntrusionProtection_Statistics - The IntrusionProtection_Statistics Type
// is used as a container for SoftLayer_Container_Network_IntrusionProtection_Statistic objects. The
// SoftLayer_Container_Network_IntrusionProtection_Statistics class holds the "header" information,
// like the item being queried (either account or data center), the time frame, and the grand total of
// the attacks.
type SoftLayer_Container_Network_IntrusionProtection_Statistics struct {

	// TotalAttacks - Total attacks for this $target over this time frame
	TotalAttacks int `json:"totalAttacks,omitempty"`

	// Target - The actual target, either a datacenter name, an account ID, or a subnet
	Target string `json:"target,omitempty"`

	// TargetType - The type of the target, right now either "datacenter", "account", or "subnet"
	TargetType string `json:"targetType,omitempty"`

	// TimeFrame - The time frame of the attack, in string form, like "Last 24 hours"
	TimeFrame string `json:"timeFrame,omitempty"`

	// TopAttacks - The top attacks for this target over this time frame
	TopAttacks []*SoftLayer_Container_Network_IntrusionProtection_Statistic `json:"topAttacks,omitempty"`
}

func (softlayer_container_network_intrusionprotection_statistics *SoftLayer_Container_Network_IntrusionProtection_Statistics) String() string {
	return "SoftLayer_Container_Network_IntrusionProtection_Statistics"
}
