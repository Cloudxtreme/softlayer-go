package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_LoadBalancer_Global_Account - The SoftLayer_Network_LoadBalancer_Global_Account
// data type contains the properties for a single global load balancer account. The properties you are
// able to edit are fallbackIp, loadBalanceTypeId, and notes. The hosts relational property can be used
// for creating and editing hosts that belong to the global load balancer account. The
// [[SoftLayer_Network_LoadBalancer_Global_Account::editObject|editObject]] method contains details on
// creating and edited hosts through the hosts relational property.
type SoftLayer_Network_LoadBalancer_Global_Account struct {

	// Id - The unique identifier of a global load balancer account.
	Id int `json:"id"`

	// Notes - Additional customer defined information for a global load balancer account.
	Notes string `json:"notes"`

	// AverageConnectionsPerSecond - The average amount of connections per second used within the current
	// billing cycle. This number is updated daily.
	AverageConnectionsPerSecond float32 `json:"averageConnectionsPerSecond"`

	// FallbackIp - The IP address that will be return to a DNS request when none of the hosts for a global
	// load balancer account could be returned.
	FallbackIp string `json:"fallbackIp"`

	// AllowedNumberOfHosts - The maximum number of hosts that a global load balancer account is allowed to
	// have.
	AllowedNumberOfHosts int `json:"allowedNumberOfHosts"`

	// ConnectionsPerSecond - The amount of connections per second a global load balancer account may use
	// within a billing cycle without being billed for an overage.
	ConnectionsPerSecond int `json:"connectionsPerSecond"`

	// Hostname - The hostname of a global load balancer account that is being load balanced.
	Hostname string `json:"hostname"`

	// LoadBalanceTypeId - The identifier of the load balance method for a global load balancer account.
	LoadBalanceTypeId int `json:"loadBalanceTypeId"`
}

func (softlayer_network_loadbalancer_global_account *SoftLayer_Network_LoadBalancer_Global_Account) String() string {
	return "SoftLayer_Network_LoadBalancer_Global_Account"
}

// SoftLayer_Network_LoadBalancer_Global_Account_Extended is SoftLayer_Network_LoadBalancer_Global_Account with all maskable types.
type SoftLayer_Network_LoadBalancer_Global_Account_Extended struct {
	SoftLayer_Network_LoadBalancer_Global_Account

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// HostCount - A count of the hosts in the load balancing pool for a global load balancer account.
	HostCount uint64 `json:"hostCount"`

	// BillingItem - The current billing item for a Global Load Balancer account.
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// LoadBalanceType - The load balance method of a global load balancer account
	LoadBalanceType *SoftLayer_Network_LoadBalancer_Global_Type `json:"loadBalanceType"`

	// ManagedResourceFlag - A flag indicating that the global load balancer is a managed resource.
	ManagedResourceFlag bool `json:"managedResourceFlag"`

	// Hosts - The hosts in the load balancing pool for a global load balancer account.
	Hosts []*SoftLayer_Network_LoadBalancer_Global_Host `json:"hosts"`
}

func (softlayer_network_loadbalancer_global_account *SoftLayer_Network_LoadBalancer_Global_Account_Extended) String() string {
	return "SoftLayer_Network_LoadBalancer_Global_Account"
}
