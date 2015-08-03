package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Router - The SoftLayer_Hardware_Router data type contains general information
// relating to a single SoftLayer router.
type SoftLayer_Hardware_Router struct {
}

// SoftLayer_Hardware_Router_Extended is SoftLayer_Hardware_Router with all maskable types.
type SoftLayer_Hardware_Router_Extended struct {
	SoftLayer_Hardware_Router

	// SanStorageCapabilityFlag - A flag indicating that a on the router can be assigned to a host that has
	// SAN disk functionality.
	SanStorageCapabilityFlag bool `json:"sanStorageCapabilityFlag"`

	// BoundSubnetCount - no documentation
	BoundSubnetCount uint64 `json:"boundSubnetCount"`

	// BoundSubnets - no documentation
	BoundSubnets []*SoftLayer_Network_Subnet `json:"boundSubnets"`

	// LocalDiskStorageCapabilityFlag - A flag indicating that a on the router can be assigned to a host
	// that has local disk functionality.
	LocalDiskStorageCapabilityFlag bool `json:"localDiskStorageCapabilityFlag"`
}

func (softlayer_hardware_router *SoftLayer_Hardware_Router) String() string {
	return "SoftLayer_Hardware_Router"
}
