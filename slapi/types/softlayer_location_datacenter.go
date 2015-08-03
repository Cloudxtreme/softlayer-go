package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Location_Datacenter - SoftLayer_Location_Datacenter extends the [[SoftLayer_Location]]
// data type to include datacenter-specific properties.
type SoftLayer_Location_Datacenter struct {
}

// SoftLayer_Location_Datacenter_Extended is SoftLayer_Location_Datacenter with all maskable types.
type SoftLayer_Location_Datacenter_Extended struct {
	SoftLayer_Location_Datacenter

	// BoundSubnets - Subnets which are directly bound to one or more routers in a given datacenter, and
	// currently allow routing.
	BoundSubnets []*SoftLayer_Network_Subnet `json:"boundSubnets"`

	// BoundSubnetCount - A count of subnets which are directly bound to one or more routers in a given
	// datacenter, and currently allow routing.
	BoundSubnetCount uint64 `json:"boundSubnetCount"`

	// RegionalGroup - no documentation
	RegionalGroup *SoftLayer_Location_Group_Regional `json:"regionalGroup"`

	// RoutableBoundSubnets - Retrieve all subnets that are eligible to be routed; those which the account
	// has permission to associate with a vlan.
	RoutableBoundSubnets []*SoftLayer_Network_Subnet `json:"routableBoundSubnets"`

	// ActiveItemPresaleEventCount - no documentation
	ActiveItemPresaleEventCount uint64 `json:"activeItemPresaleEventCount"`

	// RoutableBoundSubnetCount - A count of retrieve all subnets that are eligible to be routed; those
	// which the account has permission to associate with a vlan.
	RoutableBoundSubnetCount uint64 `json:"routableBoundSubnetCount"`

	// ActiveItemPresaleEvents - <nil>
	ActiveItemPresaleEvents []*SoftLayer_Sales_Presale_Event `json:"activeItemPresaleEvents"`

	// FrontendHardwareRouters - <nil>
	FrontendHardwareRouters []*SoftLayer_Hardware `json:"frontendHardwareRouters"`

	// PresaleEvents - <nil>
	PresaleEvents []*SoftLayer_Sales_Presale_Event `json:"presaleEvents"`

	// RegionalInternetRegistry - <nil>
	RegionalInternetRegistry *SoftLayer_Network_Regional_Internet_Registry `json:"regionalInternetRegistry"`

	// BackendHardwareRouterCount - no documentation
	BackendHardwareRouterCount uint64 `json:"backendHardwareRouterCount"`

	// HardwareRouterCount - no documentation
	HardwareRouterCount uint64 `json:"hardwareRouterCount"`

	// ActivePresaleEvents - <nil>
	ActivePresaleEvents []*SoftLayer_Sales_Presale_Event `json:"activePresaleEvents"`

	// BackendHardwareRouters - <nil>
	BackendHardwareRouters []*SoftLayer_Hardware `json:"backendHardwareRouters"`

	// HardwareRouters - <nil>
	HardwareRouters []*SoftLayer_Hardware `json:"hardwareRouters"`

	// ActivePresaleEventCount - no documentation
	ActivePresaleEventCount uint64 `json:"activePresaleEventCount"`

	// FrontendHardwareRouterCount - no documentation
	FrontendHardwareRouterCount uint64 `json:"frontendHardwareRouterCount"`

	// PresaleEventCount - no documentation
	PresaleEventCount uint64 `json:"presaleEventCount"`
}

func (softlayer_location_datacenter *SoftLayer_Location_Datacenter) String() string {
	return "SoftLayer_Location_Datacenter"
}
