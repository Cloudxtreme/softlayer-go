package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Location_Datacenter - SoftLayer_Location_Datacenter extends the [[SoftLayer_Location]]
// data type to include datacenter-specific properties.
type SoftLayer_Location_Datacenter struct {

	// ActiveItemPresaleEventCount - no documentation
	ActiveItemPresaleEventCount uint64 `json:"activeItemPresaleEventCount"`

	// ActiveItemPresaleEvents - <nil>
	ActiveItemPresaleEvents []*SoftLayer_Sales_Presale_Event `json:"activeItemPresaleEvents"`

	// ActivePresaleEventCount - no documentation
	ActivePresaleEventCount uint64 `json:"activePresaleEventCount"`

	// ActivePresaleEvents - <nil>
	ActivePresaleEvents []*SoftLayer_Sales_Presale_Event `json:"activePresaleEvents"`

	// BackendHardwareRouterCount - no documentation
	BackendHardwareRouterCount uint64 `json:"backendHardwareRouterCount"`

	// BackendHardwareRouters - <nil>
	BackendHardwareRouters []*SoftLayer_Hardware `json:"backendHardwareRouters"`

	// BoundSubnetCount - A count of subnets which are directly bound to one or more routers in a given
	// datacenter, and currently allow routing.
	BoundSubnetCount uint64 `json:"boundSubnetCount"`

	// BoundSubnets - Subnets which are directly bound to one or more routers in a given datacenter, and
	// currently allow routing.
	BoundSubnets []*SoftLayer_Network_Subnet `json:"boundSubnets"`

	// FrontendHardwareRouterCount - no documentation
	FrontendHardwareRouterCount uint64 `json:"frontendHardwareRouterCount"`

	// FrontendHardwareRouters - <nil>
	FrontendHardwareRouters []*SoftLayer_Hardware `json:"frontendHardwareRouters"`

	// HardwareRouterCount - no documentation
	HardwareRouterCount uint64 `json:"hardwareRouterCount"`

	// HardwareRouters - <nil>
	HardwareRouters []*SoftLayer_Hardware `json:"hardwareRouters"`

	// PresaleEventCount - no documentation
	PresaleEventCount uint64 `json:"presaleEventCount"`

	// PresaleEvents - <nil>
	PresaleEvents []*SoftLayer_Sales_Presale_Event `json:"presaleEvents"`

	// RegionalGroup - no documentation
	RegionalGroup *SoftLayer_Location_Group_Regional `json:"regionalGroup"`

	// RegionalInternetRegistry - <nil>
	RegionalInternetRegistry *SoftLayer_Network_Regional_Internet_Registry `json:"regionalInternetRegistry"`

	// RoutableBoundSubnetCount - A count of retrieve all subnets that are eligible to be routed; those
	// which the account has permission to associate with a vlan.
	RoutableBoundSubnetCount uint64 `json:"routableBoundSubnetCount"`

	// RoutableBoundSubnets - Retrieve all subnets that are eligible to be routed; those which the account
	// has permission to associate with a vlan.
	RoutableBoundSubnets []*SoftLayer_Network_Subnet `json:"routableBoundSubnets"`
}

// GetObject - <nil>
func (softlayer_location_datacenter *SoftLayer_Location_Datacenter) GetObject() (*SoftLayer_Location_Datacenter, error) {
	var returnValue *SoftLayer_Location_Datacenter
	return returnValue, nil
}

// GetStatisticsGraphImage - Retrieve a graph of a SoftLayer datacenter's last 48 hours of network
// activity. Statistics graphs show traffic outbound from a datacenter on top and inbound traffic on
// the bottom followed by a legend of the network services tracked in the graph.
// getStatisticsGraphImage returns a PNG image of variable width and height depending on the number of
// services reported in the image.
func (softlayer_location_datacenter *SoftLayer_Location_Datacenter) GetStatisticsGraphImage() (string, error) {
	var returnValue string
	return returnValue, nil
}
