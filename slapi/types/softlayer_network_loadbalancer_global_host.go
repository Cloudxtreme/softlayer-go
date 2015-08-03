package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_LoadBalancer_Global_Host - The SoftLayer_Network_LoadBalancer_Global_Host data
// type represents a single host that belongs to a global load balancer account's load balancing pool.
// The destination IP address of a host must be one that belongs to your SoftLayer customer account, or
// to a datacenter load balancer virtual ip that belongs to your SoftLayer customer account. The
// destination IP address and port of a global load balancer host is a required field and must exist
// during creation and can not be removed. The acceptable values for the health check type are 'none',
// 'http', and 'tcp'. The status property is updated in 5 minute intervals and the hits property is
// updated in 10 minute intervals. The order of the host is only important if you are using the
// 'failover' load balance method, and the weight is only important if you are using the 'weighted
// round robin' load balance method.
type SoftLayer_Network_LoadBalancer_Global_Host struct {

	// DestinationIp - The IP address of the host that will be returned by the global load balancers in
	// response to a dns request.
	DestinationIp string `json:"destinationIp"`

	// DestinationPort - The port of the host that will be used for health checks.
	DestinationPort int `json:"destinationPort"`

	// Enabled - Whether the host is enabled or not. The value can be '0' for disabled, or '1' for enabled.
	Enabled int `json:"enabled"`

	// HealthCheck - The health check type of a host. Valid values include 'none', 'http', and 'tcp'.
	HealthCheck string `json:"healthCheck"`

	// Hits - The number of times the host was selected by the load balance method.
	Hits float32 `json:"hits"`

	// Id - The unique identifier of a global load balancer host.
	Id int `json:"id"`

	// LoadBalanceOrder - The order of this host within the load balance pool. This is only significant if
	// the load balance method is set to failover.
	LoadBalanceOrder int `json:"loadBalanceOrder"`

	// LoadBalancerAccount - The global load balancer account a host belongs to.
	LoadBalancerAccount *SoftLayer_Network_LoadBalancer_Global_Account `json:"loadBalancerAccount"`

	// Location - The location of a host in a datacenter.serverRoom format.
	Location string `json:"location"`

	// Status - The health status of a host. The status can be either or null which could mean that the
	// health check type is set to 'none' or an update to the ip, port, or health check type was recently
	// done and the host is waiting for the new status.
	Status string `json:"status"`

	// Weight - The load balance weight of a host. The total weight of all hosts in the load balancing pool
	// must not exceed 100.
	Weight int `json:"weight"`
}

func (softlayer_network_loadbalancer_global_host *SoftLayer_Network_LoadBalancer_Global_Host) String() string {
	return "SoftLayer_Network_LoadBalancer_Global_Host"
}
