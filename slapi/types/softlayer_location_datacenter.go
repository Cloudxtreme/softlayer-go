package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Location_Datacenter - SoftLayer_Location_Datacenter extends the [[SoftLayer_Location]]
// data type to include datacenter-specific properties.
type SoftLayer_Location_Datacenter struct {

	// LongName - no documentation
	LongName string `json:"longName,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// StatusId - <nil>
	StatusId int `json:"statusId,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`
}

func (softlayer_location_datacenter *SoftLayer_Location_Datacenter) String() string {
	return "SoftLayer_Location_Datacenter"
}

// SoftLayer_Location_Datacenter_Extended is SoftLayer_Location_Datacenter with all maskable types.
type SoftLayer_Location_Datacenter_Extended struct {
	SoftLayer_Location_Datacenter

	// HardwareRouters - <nil>
	HardwareRouters []*SoftLayer_Hardware `json:"hardwareRouters,omitempty"`

	// PresaleEvents - <nil>
	PresaleEvents []*SoftLayer_Sales_Presale_Event `json:"presaleEvents,omitempty"`

	// RoutableBoundSubnets - Retrieve all subnets that are eligible to be routed; those which the account
	// has permission to associate with a vlan.
	RoutableBoundSubnets []*SoftLayer_Network_Subnet `json:"routableBoundSubnets,omitempty"`

	// HardwareRouterCount - no documentation
	HardwareRouterCount uint64 `json:"hardwareRouterCount,omitempty"`

	// PathString - <nil>
	PathString string `json:"pathString,omitempty"`

	// GroupCount - A count of a location can be a member of 1 or more groups. This will show which groups
	// to which a location belongs.
	GroupCount uint64 `json:"groupCount,omitempty"`

	// ActivePresaleEvents - <nil>
	ActivePresaleEvents []*SoftLayer_Sales_Presale_Event `json:"activePresaleEvents,omitempty"`

	// PresaleEventCount - no documentation
	PresaleEventCount uint64 `json:"presaleEventCount,omitempty"`

	// PriceGroupCount - A count of a location can be a member of 1 or more Price Groups. This will show
	// which groups to which a location belongs.
	PriceGroupCount uint64 `json:"priceGroupCount,omitempty"`

	// HardwareFirewallCount - no documentation
	HardwareFirewallCount uint64 `json:"hardwareFirewallCount,omitempty"`

	// ActiveItemPresaleEvents - <nil>
	ActiveItemPresaleEvents []*SoftLayer_Sales_Presale_Event `json:"activeItemPresaleEvents,omitempty"`

	// RoutableBoundSubnetCount - A count of retrieve all subnets that are eligible to be routed; those
	// which the account has permission to associate with a vlan.
	RoutableBoundSubnetCount uint64 `json:"routableBoundSubnetCount,omitempty"`

	// BackboneDependents - <nil>
	BackboneDependents []*SoftLayer_Network_Backbone_Location_Dependent `json:"backboneDependents,omitempty"`

	// NetworkConfigurationAttribute - <nil>
	NetworkConfigurationAttribute *SoftLayer_Hardware_Attribute `json:"networkConfigurationAttribute,omitempty"`

	// ActivePresaleEventCount - no documentation
	ActivePresaleEventCount uint64 `json:"activePresaleEventCount,omitempty"`

	// FrontendHardwareRouterCount - no documentation
	FrontendHardwareRouterCount uint64 `json:"frontendHardwareRouterCount,omitempty"`

	// PriceGroups - A location can be a member of 1 or more Price Groups. This will show which groups to
	// which a location belongs.
	PriceGroups []*SoftLayer_Location_Group `json:"priceGroups,omitempty"`

	// Groups - A location can be a member of 1 or more groups. This will show which groups to which a
	// location belongs.
	Groups []*SoftLayer_Location_Group `json:"groups,omitempty"`

	// BackendHardwareRouterCount - no documentation
	BackendHardwareRouterCount uint64 `json:"backendHardwareRouterCount,omitempty"`

	// Timezone - <nil>
	Timezone *SoftLayer_Locale_Timezone `json:"timezone,omitempty"`

	// BackboneDependentCount - no documentation
	BackboneDependentCount uint64 `json:"backboneDependentCount,omitempty"`

	// BackendHardwareRouters - <nil>
	BackendHardwareRouters []*SoftLayer_Hardware `json:"backendHardwareRouters,omitempty"`

	// RegionalGroup - no documentation
	RegionalGroup *SoftLayer_Location_Group_Regional `json:"regionalGroup,omitempty"`

	// RegionCount - A count of a location can be a member of 1 or more regions. This will show which
	// regions to which a location belongs.
	RegionCount uint64 `json:"regionCount,omitempty"`

	// LocationAddress - no documentation
	LocationAddress *SoftLayer_Account_Address `json:"locationAddress,omitempty"`

	// LocationReservationMember - no documentation
	LocationReservationMember *SoftLayer_Location_Reservation_Rack_Member `json:"locationReservationMember,omitempty"`

	// Regions - A location can be a member of 1 or more regions. This will show which regions to which a
	// location belongs.
	Regions []*SoftLayer_Location_Region `json:"regions,omitempty"`

	// BoundSubnetCount - A count of subnets which are directly bound to one or more routers in a given
	// datacenter, and currently allow routing.
	BoundSubnetCount uint64 `json:"boundSubnetCount,omitempty"`

	// HardwareFirewalls - <nil>
	HardwareFirewalls []*SoftLayer_Hardware `json:"hardwareFirewalls,omitempty"`

	// OnlineSslVpnUserCount - The total number of users online using SoftLayer's SSL VPN service for a
	// location.
	OnlineSslVpnUserCount int `json:"onlineSslVpnUserCount,omitempty"`

	// LocationStatus - no documentation
	LocationStatus *SoftLayer_Location_Status `json:"locationStatus,omitempty"`

	// BoundSubnets - Subnets which are directly bound to one or more routers in a given datacenter, and
	// currently allow routing.
	BoundSubnets []*SoftLayer_Network_Subnet `json:"boundSubnets,omitempty"`

	// FrontendHardwareRouters - <nil>
	FrontendHardwareRouters []*SoftLayer_Hardware `json:"frontendHardwareRouters,omitempty"`

	// RegionalInternetRegistry - <nil>
	RegionalInternetRegistry *SoftLayer_Network_Regional_Internet_Registry `json:"regionalInternetRegistry,omitempty"`

	// ActiveItemPresaleEventCount - no documentation
	ActiveItemPresaleEventCount uint64 `json:"activeItemPresaleEventCount,omitempty"`

	// OnlinePptpVpnUserCount - The total number of users online using SoftLayer's VPN service for a
	// location.
	OnlinePptpVpnUserCount int `json:"onlinePptpVpnUserCount,omitempty"`

	// VdrGroup - A location can be a member of 1 Bandwidth Pooling Group. This will show which group to
	// which a location belongs.
	VdrGroup *SoftLayer_Location_Group_Location_CrossReference `json:"vdrGroup,omitempty"`
}

func (softlayer_location_datacenter *SoftLayer_Location_Datacenter_Extended) String() string {
	return "SoftLayer_Location_Datacenter"
}
