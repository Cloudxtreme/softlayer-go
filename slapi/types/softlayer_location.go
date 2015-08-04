package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Location - Every piece of hardware and network connection owned by SoftLayer is tracked
// physically by location and stored in the SoftLayer_Location data type. SoftLayer locations exist in
// parent/child relationships, a convenient way to track equipment from it's city, datacenter, server
// room, rack, then slot. Network backbones are tied to datacenters only, not to a room, rack, or slot.
type SoftLayer_Location struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// StatusId - <nil>
	StatusId int `json:"statusId,omitempty"`

	// LongName - no documentation
	LongName string `json:"longName,omitempty"`
}

func (softlayer_location *SoftLayer_Location) String() string {
	return "SoftLayer_Location"
}

// SoftLayer_Location_Extended is SoftLayer_Location with all maskable types.
type SoftLayer_Location_Extended struct {
	SoftLayer_Location

	// LocationAddress - no documentation
	LocationAddress *SoftLayer_Account_Address `json:"locationAddress,omitempty"`

	// LocationReservationMember - no documentation
	LocationReservationMember *SoftLayer_Location_Reservation_Rack_Member `json:"locationReservationMember,omitempty"`

	// Timezone - <nil>
	Timezone *SoftLayer_Locale_Timezone `json:"timezone,omitempty"`

	// Groups - A location can be a member of 1 or more groups. This will show which groups to which a
	// location belongs.
	Groups []*SoftLayer_Location_Group `json:"groups,omitempty"`

	// NetworkConfigurationAttribute - <nil>
	NetworkConfigurationAttribute *SoftLayer_Hardware_Attribute `json:"networkConfigurationAttribute,omitempty"`

	// Regions - A location can be a member of 1 or more regions. This will show which regions to which a
	// location belongs.
	Regions []*SoftLayer_Location_Region `json:"regions,omitempty"`

	// BackboneDependentCount - no documentation
	BackboneDependentCount uint64 `json:"backboneDependentCount,omitempty"`

	// HardwareFirewallCount - no documentation
	HardwareFirewallCount uint64 `json:"hardwareFirewallCount,omitempty"`

	// HardwareFirewalls - <nil>
	HardwareFirewalls []*SoftLayer_Hardware `json:"hardwareFirewalls,omitempty"`

	// OnlinePptpVpnUserCount - The total number of users online using SoftLayer's VPN service for a
	// location.
	OnlinePptpVpnUserCount int `json:"onlinePptpVpnUserCount,omitempty"`

	// OnlineSslVpnUserCount - The total number of users online using SoftLayer's SSL VPN service for a
	// location.
	OnlineSslVpnUserCount int `json:"onlineSslVpnUserCount,omitempty"`

	// PathString - <nil>
	PathString string `json:"pathString,omitempty"`

	// VdrGroup - A location can be a member of 1 Bandwidth Pooling Group. This will show which group to
	// which a location belongs.
	VdrGroup *SoftLayer_Location_Group_Location_CrossReference `json:"vdrGroup,omitempty"`

	// RegionCount - A count of a location can be a member of 1 or more regions. This will show which
	// regions to which a location belongs.
	RegionCount uint64 `json:"regionCount,omitempty"`

	// BackboneDependents - <nil>
	BackboneDependents []*SoftLayer_Network_Backbone_Location_Dependent `json:"backboneDependents,omitempty"`

	// LocationStatus - no documentation
	LocationStatus *SoftLayer_Location_Status `json:"locationStatus,omitempty"`

	// PriceGroups - A location can be a member of 1 or more Price Groups. This will show which groups to
	// which a location belongs.
	PriceGroups []*SoftLayer_Location_Group `json:"priceGroups,omitempty"`

	// GroupCount - A count of a location can be a member of 1 or more groups. This will show which groups
	// to which a location belongs.
	GroupCount uint64 `json:"groupCount,omitempty"`

	// PriceGroupCount - A count of a location can be a member of 1 or more Price Groups. This will show
	// which groups to which a location belongs.
	PriceGroupCount uint64 `json:"priceGroupCount,omitempty"`
}

func (softlayer_location *SoftLayer_Location_Extended) String() string {
	return "SoftLayer_Location"
}
