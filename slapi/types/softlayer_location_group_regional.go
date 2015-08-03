package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Location_Group_Regional - <nil>
type SoftLayer_Location_Group_Regional struct {
}

func (softlayer_location_group_regional *SoftLayer_Location_Group_Regional) String() string {
	return "SoftLayer_Location_Group_Regional"
}

// SoftLayer_Location_Group_Regional_Extended is SoftLayer_Location_Group_Regional with all maskable types.
type SoftLayer_Location_Group_Regional_Extended struct {
	SoftLayer_Location_Group_Regional

	// Datacenters - no documentation
	Datacenters []*SoftLayer_Location `json:"datacenters"`

	// PreferredDatacenter - no documentation
	PreferredDatacenter *SoftLayer_Location_Datacenter `json:"preferredDatacenter"`

	// DatacenterCount - no documentation
	DatacenterCount uint64 `json:"datacenterCount"`
}

func (softlayer_location_group_regional *SoftLayer_Location_Group_Regional_Extended) String() string {
	return "SoftLayer_Location_Group_Regional"
}
