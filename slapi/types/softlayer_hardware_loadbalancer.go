package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_LoadBalancer - <nil>
type SoftLayer_Hardware_LoadBalancer struct {
}

// SoftLayer_Hardware_LoadBalancer_Extended is SoftLayer_Hardware_LoadBalancer with all maskable types.
type SoftLayer_Hardware_LoadBalancer_Extended struct {
	SoftLayer_Hardware_LoadBalancer

	// ModelFamily - <nil>
	ModelFamily string `json:"modelFamily"`

	// Users - A list of users that have access to this hardware load balancer.
	Users []*SoftLayer_User_Customer `json:"users"`

	// UserCount - A count of a list of users that have access to this hardware load balancer.
	UserCount uint64 `json:"userCount"`
}

func (softlayer_hardware_loadbalancer *SoftLayer_Hardware_LoadBalancer) String() string {
	return "SoftLayer_Hardware_LoadBalancer"
}
