package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_LoadBalancer_Global_Account - The SoftLayer_Network_LoadBalancer_Global_Account
// data type contains the properties for a single global load balancer account. The properties you are
// able to edit are fallbackIp, loadBalanceTypeId, and notes. The hosts relational property can be used
// for creating and editing hosts that belong to the global load balancer account. The
// [[SoftLayer_Network_LoadBalancer_Global_Account::editObject|editObject]] method contains details on
// creating and edited hosts through the hosts relational property.
type SoftLayer_Network_LoadBalancer_Global_Account struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AllowedNumberOfHosts - The maximum number of hosts that a global load balancer account is allowed to
	// have.
	AllowedNumberOfHosts int `json:"allowedNumberOfHosts"`

	// AverageConnectionsPerSecond - The average amount of connections per second used within the current
	// billing cycle. This number is updated daily.
	AverageConnectionsPerSecond float32 `json:"averageConnectionsPerSecond"`

	// BillingItem - The current billing item for a Global Load Balancer account.
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// ConnectionsPerSecond - The amount of connections per second a global load balancer account may use
	// within a billing cycle without being billed for an overage.
	ConnectionsPerSecond int `json:"connectionsPerSecond"`

	// FallbackIp - The IP address that will be return to a DNS request when none of the hosts for a global
	// load balancer account could be returned.
	FallbackIp string `json:"fallbackIp"`

	// HostCount - A count of the hosts in the load balancing pool for a global load balancer account.
	HostCount uint64 `json:"hostCount"`

	// Hostname - The hostname of a global load balancer account that is being load balanced.
	Hostname string `json:"hostname"`

	// Hosts - The hosts in the load balancing pool for a global load balancer account.
	Hosts []*SoftLayer_Network_LoadBalancer_Global_Host `json:"hosts"`

	// Id - The unique identifier of a global load balancer account.
	Id int `json:"id"`

	// LoadBalanceType - The load balance method of a global load balancer account
	LoadBalanceType *SoftLayer_Network_LoadBalancer_Global_Type `json:"loadBalanceType"`

	// LoadBalanceTypeId - The identifier of the load balance method for a global load balancer account.
	LoadBalanceTypeId int `json:"loadBalanceTypeId"`

	// ManagedResourceFlag - A flag indicating that the global load balancer is a managed resource.
	ManagedResourceFlag bool `json:"managedResourceFlag"`

	// Notes - Additional customer defined information for a global load balancer account.
	Notes string `json:"notes"`
}

// AddNsRecord - If your globally load balanced domain is hosted on the SoftLayer nameservers this
// method will add the required NS resource record to your DNS zone file and remove any A records that
// match the host portion of a global load balancer account hostname. A NS resource record is required
// to be able to use your SoftLayer global load balancer account. Please make sure the zone file for
// the hostname listed on your SoftLayer global load balancer account is setup prior to using this
// method. If your globally load balanced domain is hosted on any other nameservers this method will
// not be able to add the required NS record.
func (softlayer_network_loadbalancer_global_account *SoftLayer_Network_LoadBalancer_Global_Account) AddNsRecord(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - Edit the properties of a global load balancer account by passing in a modified instance
// of the object. The global load balancer account properties you are able to edit are: fallback ip,
// load balance type id, and notes. Hosts that belong to your SoftLayer global load balancer account
// are created and modified through this method. An example templateObject that updates global load
// balancer account properties, updates the properties of a host, and adds a new host is shown below: *
// id: 2 * loadBalanceTypeId: 2 * notes: Notes updated * fallbackIp: 1.1.1.1 * hosts: ** id: 19 **
// destinationIp: 2.2.2.2 ** weight: 25 ** healthCheck: http ** destinationPort: 80 ** enabled: 1 **
// destinationIp: 3.3.3.3 ** weight: 25 ** healthCheck: http ** destinationPort: 80 ** enabled: 1 The
// first section contains the properties of the global load balancer account that will be updated,
// while the second section contains the elements of the 'hosts' property of the global load balancer
// account. The first host listed will have its properties updated because the 'id' property of the
// host is set, meaning the global load balancer host with an id of 19 will be updated. The second host
// listed will be created because it lacks the 'id' property. There is a limit to the maximum number of
// hosts that you are allowed to add, and is defined by the allowedNumberOfHosts property on the global
// load balancer account. The destination IP address of a host must be an IP address that belongs to
// your SoftLayer Account, or a local load balancer virtual IP address that belongs to your account.
// The destination IP address and destination port are required and must be provided when creating a
// host.
func (softlayer_network_loadbalancer_global_account *SoftLayer_Network_LoadBalancer_Global_Account) EditObject(ctx *slapi.RequestContext, templateObject SoftLayer_Network_LoadBalancer_Global_Account) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Network_LoadBalancer_Global_Account object whose ID
// number corresponds to the ID number of the init paramater passed to the
// SoftLayer_Network_LoadBalancer_Global_Account service. You can only retrieve a global load balancer
// account that is assigned to your SoftLayer customer account.
func (softlayer_network_loadbalancer_global_account *SoftLayer_Network_LoadBalancer_Global_Account) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_LoadBalancer_Global_Account, error) {
	var returnValue *SoftLayer_Network_LoadBalancer_Global_Account
	return returnValue, nil
}

// RemoveNsRecord - If your globally load balanced domain is hosted on the SoftLayer nameservers this
// method will remove the NS resource record from your DNS zone file. Removing the NS resource record
// will basically disable your global load balancer account since no DNS requests will be forwarded to
// the global load balancers. Any A records that were removed when the NS resource record was added
// will not be created for you. If your globally load balanced domain is hosted on any other
// nameservers this method will not be able to remove the required NS record.
func (softlayer_network_loadbalancer_global_account *SoftLayer_Network_LoadBalancer_Global_Account) RemoveNsRecord(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
