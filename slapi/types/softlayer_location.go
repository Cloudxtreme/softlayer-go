package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Location - Every piece of hardware and network connection owned by SoftLayer is tracked
// physically by location and stored in the SoftLayer_Location data type. SoftLayer locations exist in
// parent/child relationships, a convenient way to track equipment from it's city, datacenter, server
// room, rack, then slot. Network backbones are tied to datacenters only, not to a room, rack, or slot.
type SoftLayer_Location struct {

	// StatusId - <nil>
	StatusId int `json:"statusId"`

	// Id - no documentation
	Id int `json:"id"`

	// LongName - no documentation
	LongName string `json:"longName"`

	// Name - no documentation
	Name string `json:"name"`
}

// SoftLayer_Location_Extended is SoftLayer_Location with all maskable types.
type SoftLayer_Location_Extended struct {
	SoftLayer_Location

	// PathString - <nil>
	PathString string `json:"pathString"`

	// BackboneDependentCount - no documentation
	BackboneDependentCount uint64 `json:"backboneDependentCount"`

	// GroupCount - A count of a location can be a member of 1 or more groups. This will show which groups
	// to which a location belongs.
	GroupCount uint64 `json:"groupCount"`

	// HardwareFirewallCount - no documentation
	HardwareFirewallCount uint64 `json:"hardwareFirewallCount"`

	// RegionCount - A count of a location can be a member of 1 or more regions. This will show which
	// regions to which a location belongs.
	RegionCount uint64 `json:"regionCount"`

	// OnlineSslVpnUserCount - The total number of users online using SoftLayer's SSL VPN service for a
	// location.
	OnlineSslVpnUserCount int `json:"onlineSslVpnUserCount"`

	// LocationStatus - no documentation
	LocationStatus *SoftLayer_Location_Status `json:"locationStatus"`

	// PriceGroups - A location can be a member of 1 or more Price Groups. This will show which groups to
	// which a location belongs.
	PriceGroups []*SoftLayer_Location_Group `json:"priceGroups"`

	// Timezone - <nil>
	Timezone *SoftLayer_Locale_Timezone `json:"timezone"`

	// PriceGroupCount - A count of a location can be a member of 1 or more Price Groups. This will show
	// which groups to which a location belongs.
	PriceGroupCount uint64 `json:"priceGroupCount"`

	// LocationReservationMember - no documentation
	LocationReservationMember *SoftLayer_Location_Reservation_Rack_Member `json:"locationReservationMember"`

	// LocationAddress - no documentation
	LocationAddress *SoftLayer_Account_Address `json:"locationAddress"`

	// VdrGroup - A location can be a member of 1 Bandwidth Pooling Group. This will show which group to
	// which a location belongs.
	VdrGroup *SoftLayer_Location_Group_Location_CrossReference `json:"vdrGroup"`

	// BackboneDependents - <nil>
	BackboneDependents []*SoftLayer_Network_Backbone_Location_Dependent `json:"backboneDependents"`

	// HardwareFirewalls - <nil>
	HardwareFirewalls []*SoftLayer_Hardware `json:"hardwareFirewalls"`

	// NetworkConfigurationAttribute - <nil>
	NetworkConfigurationAttribute *SoftLayer_Hardware_Attribute `json:"networkConfigurationAttribute"`

	// OnlinePptpVpnUserCount - The total number of users online using SoftLayer's VPN service for a
	// location.
	OnlinePptpVpnUserCount int `json:"onlinePptpVpnUserCount"`

	// Regions - A location can be a member of 1 or more regions. This will show which regions to which a
	// location belongs.
	Regions []*SoftLayer_Location_Region `json:"regions"`

	// Groups - A location can be a member of 1 or more groups. This will show which groups to which a
	// location belongs.
	Groups []*SoftLayer_Location_Group `json:"groups"`
}

func (softlayer_location *SoftLayer_Location) String() string {
	return "SoftLayer_Location"
}
