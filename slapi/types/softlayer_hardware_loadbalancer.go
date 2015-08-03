package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_LoadBalancer - <nil>
type SoftLayer_Hardware_LoadBalancer struct {

	// ModelFamily - <nil>
	ModelFamily string `json:"modelFamily"`

	// UserCount - A count of a list of users that have access to this hardware load balancer.
	UserCount uint64 `json:"userCount"`

	// Users - A list of users that have access to this hardware load balancer.
	Users []*SoftLayer_User_Customer `json:"users"`
}

func (softlayer_hardware_loadbalancer *SoftLayer_Hardware_LoadBalancer) String() string {
	return "SoftLayer_Hardware_LoadBalancer"
}
