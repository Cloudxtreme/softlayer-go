package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Location - Every piece of hardware and network connection owned by SoftLayer is tracked
// physically by location and stored in the SoftLayer_Location data type. SoftLayer locations exist in
// parent/child relationships, a convenient way to track equipment from it's city, datacenter, server
// room, rack, then slot. Network backbones are tied to datacenters only, not to a room, rack, or slot.
type SoftLayer_Location struct {

	// BackboneDependentCount - no documentation
	BackboneDependentCount uint64 `json:"backboneDependentCount"`

	// BackboneDependents - <nil>
	BackboneDependents []*SoftLayer_Network_Backbone_Location_Dependent `json:"backboneDependents"`

	// GroupCount - A count of a location can be a member of 1 or more groups. This will show which groups
	// to which a location belongs.
	GroupCount uint64 `json:"groupCount"`

	// Groups - A location can be a member of 1 or more groups. This will show which groups to which a
	// location belongs.
	Groups []*SoftLayer_Location_Group `json:"groups"`

	// HardwareFirewallCount - no documentation
	HardwareFirewallCount uint64 `json:"hardwareFirewallCount"`

	// HardwareFirewalls - <nil>
	HardwareFirewalls []*SoftLayer_Hardware `json:"hardwareFirewalls"`

	// Id - no documentation
	Id int `json:"id"`

	// LocationAddress - no documentation
	LocationAddress *SoftLayer_Account_Address `json:"locationAddress"`

	// LocationReservationMember - no documentation
	LocationReservationMember *SoftLayer_Location_Reservation_Rack_Member `json:"locationReservationMember"`

	// LocationStatus - no documentation
	LocationStatus *SoftLayer_Location_Status `json:"locationStatus"`

	// LongName - no documentation
	LongName string `json:"longName"`

	// Name - no documentation
	Name string `json:"name"`

	// NetworkConfigurationAttribute - <nil>
	NetworkConfigurationAttribute *SoftLayer_Hardware_Attribute `json:"networkConfigurationAttribute"`

	// OnlinePptpVpnUserCount - The total number of users online using SoftLayer's VPN service for a
	// location.
	OnlinePptpVpnUserCount int `json:"onlinePptpVpnUserCount"`

	// OnlineSslVpnUserCount - The total number of users online using SoftLayer's SSL VPN service for a
	// location.
	OnlineSslVpnUserCount int `json:"onlineSslVpnUserCount"`

	// PathString - <nil>
	PathString string `json:"pathString"`

	// PriceGroupCount - A count of a location can be a member of 1 or more Price Groups. This will show
	// which groups to which a location belongs.
	PriceGroupCount uint64 `json:"priceGroupCount"`

	// PriceGroups - A location can be a member of 1 or more Price Groups. This will show which groups to
	// which a location belongs.
	PriceGroups []*SoftLayer_Location_Group `json:"priceGroups"`

	// RegionCount - A count of a location can be a member of 1 or more regions. This will show which
	// regions to which a location belongs.
	RegionCount uint64 `json:"regionCount"`

	// Regions - A location can be a member of 1 or more regions. This will show which regions to which a
	// location belongs.
	Regions []*SoftLayer_Location_Region `json:"regions"`

	// StatusId - <nil>
	StatusId int `json:"statusId"`

	// Timezone - <nil>
	Timezone *SoftLayer_Locale_Timezone `json:"timezone"`

	// VdrGroup - A location can be a member of 1 Bandwidth Pooling Group. This will show which group to
	// which a location belongs.
	VdrGroup *SoftLayer_Location_Group_Location_CrossReference `json:"vdrGroup"`
}

func (softlayer_location *SoftLayer_Location) String() string {
	return "SoftLayer_Location"
}

// GetAvailableObjectStorageDatacenters - Object Storage is only available in select datacenters. This
// method will return all the datacenters where object storage is available.
func (softlayer_location *SoftLayer_Location) GetAvailableObjectStorageDatacenters(ctx *slapi.RequestContext) ([]*SoftLayer_Location, error) {
	var returnValue []*SoftLayer_Location
	return returnValue, nil
}

// GetDatacenters - Retrieve all datacenter locations. SoftLayer's datacenters exist in various cities
// and each contain one or more server rooms which house network and server infrastructure.
func (softlayer_location *SoftLayer_Location) GetDatacenters(ctx *slapi.RequestContext) ([]*SoftLayer_Location, error) {
	var returnValue []*SoftLayer_Location
	return returnValue, nil
}

// GetDatacentersWithVirtualImageStoreServiceResourceRecord - <nil>
func (softlayer_location *SoftLayer_Location) GetDatacentersWithVirtualImageStoreServiceResourceRecord(ctx *slapi.RequestContext) ([]*SoftLayer_Location, error) {
	var returnValue []*SoftLayer_Location
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_location *SoftLayer_Location) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Location, error) {
	var returnValue *SoftLayer_Location
	return returnValue, nil
}

// GetViewableDatacenters - Retrieve all datacenter locations. SoftLayer's datacenters exist in various
// cities and each contain one or more server rooms which house network and server infrastructure.
func (softlayer_location *SoftLayer_Location) GetViewableDatacenters(ctx *slapi.RequestContext) ([]*SoftLayer_Location, error) {
	var returnValue []*SoftLayer_Location
	return returnValue, nil
}

// GetViewablePopsAndDataCenters - Retrieve all viewable pop and datacenter locations.
func (softlayer_location *SoftLayer_Location) GetViewablePopsAndDataCenters(ctx *slapi.RequestContext) ([]*SoftLayer_Location, error) {
	var returnValue []*SoftLayer_Location
	return returnValue, nil
}

// GetViewablepointOfPresence - no documentation
func (softlayer_location *SoftLayer_Location) GetViewablepointOfPresence(ctx *slapi.RequestContext) ([]*SoftLayer_Location, error) {
	var returnValue []*SoftLayer_Location
	return returnValue, nil
}

// GetpointOfPresence - no documentation
func (softlayer_location *SoftLayer_Location) GetpointOfPresence(ctx *slapi.RequestContext) ([]*SoftLayer_Location, error) {
	var returnValue []*SoftLayer_Location
	return returnValue, nil
}
