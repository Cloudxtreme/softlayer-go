package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Network_LoadBalancer_VirtualIpAddress - The
// SoftLayer_Network_LoadBalancer_VirtualIpAddress data type contains all the information relating to a
// specific load balancer assigned to a customer account. Information retained on the object itself is
// the virtual IP address, load balancing method, and any notes that are related to the load balancer.
// There is also an array of SoftLayer_Network_LoadBalancer_Service objects, which represent the load
// balancer services, explained more fully in the SoftLayer_Network_LoadBalancer_Service documentation.
type SoftLayer_Network_LoadBalancer_VirtualIpAddress struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// ConnectionLimit - Connection limit on this Can be upgraded through the upgradeConnectionLimit()
	// function
	ConnectionLimit int `json:"connectionLimit"`

	// CustomerManagedFlag - If false, this VIP and associated services may be edited via the portal or the
	// If true, you must configure this VIP manually on the device.
	CustomerManagedFlag int `json:"customerManagedFlag"`

	// Id - Unique ID for this object, used for the getObject method, and must be set if you are editing
	// this object.
	Id int `json:"id"`

	// LoadBalancingMethod - The load balancing method that determines which server is used "next" by the
	// load balancer. The method is stored in an abbreviated form, represented in parentheses after the
	// full name. Methods include: Round Robin (Value "rr"): Each server is used sequentially in a circular
	// queue Shortest Response (Value "sr"): The server with the lowest ping at the last health check gets
	// the next request Least Connections (Value "lc"): The server with the least current connections is
	// given the next request Persistent IP - Round Robin (Value "pi"): The same server will be returned to
	// a request during a users session. Servers are chosen through round robin. Persistent IP - Shortest
	// Response (Value "pi-sr"): The same server will be returned to a request during a users session.
	// Servers are chosen through shortest response. Persistent IP - Least Connections (Value "pi-lc"): The
	// same server will be returned to a request during a users session. Servers are chosen through least
	// connections. Insert Cookie - Round Robin (Value "ic"): Inserts a cookie into the stream that will
	// tie that client to a particular balanced server. Servers are chosen through round robin. Insert
	// Cookie - Shortest Response (Value "ic-sr"): Inserts a cookie into the stream that will tie that
	// client to a particular balanced server. Servers are chosen through shortest response. Insert Cookie
	// - Least Connections (Value "ic-lc"): Inserts a cookie into the stream that will tie that client to a
	// particular balanced server. Servers are chosen through least connections.
	LoadBalancingMethod string `json:"loadBalancingMethod"`

	// LoadBalancingMethodFullName - A human readable version of loadBalancingMethod, intended mainly for
	// API users.
	LoadBalancingMethodFullName string `json:"loadBalancingMethodFullName"`

	// ManagedResourceFlag - A flag indicating that the load balancer is a managed resource.
	ManagedResourceFlag bool `json:"managedResourceFlag"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`

	// Notes - no documentation
	Notes string `json:"notes"`

	// SecurityCertificateId - The unique identifier of the Security Certificate to be utilized when SSL
	// support is enabled.
	SecurityCertificateId int `json:"securityCertificateId"`

	// ServiceCount - no documentation
	ServiceCount uint64 `json:"serviceCount"`

	// Services - no documentation
	Services []*SoftLayer_Network_LoadBalancer_Service `json:"services"`

	// SourcePort - no documentation
	SourcePort int `json:"sourcePort"`

	// Type - The connection type of this Valid values are and
	Type string `json:"type"`

	// VirtualIpAddress - The virtual, public-facing IP address for your load balancer. This is the address
	// of all incoming traffic
	VirtualIpAddress string `json:"virtualIpAddress"`
}

func (softlayer_network_loadbalancer_virtualipaddress *SoftLayer_Network_LoadBalancer_VirtualIpAddress) String() string {
	return "SoftLayer_Network_LoadBalancer_VirtualIpAddress"
}

// Disable - Disable a Virtual IP Address, removing it from load balancer rotation and denying all
// connections to that IP address.
func (softlayer_network_loadbalancer_virtualipaddress *SoftLayer_Network_LoadBalancer_VirtualIpAddress) Disable(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - Like any other API object, the load balancers can have their exposed properties edited
// by passing in a modified version of the object. The load balancer object also can modify its
// services in this way. Simply request the load balancer object you wish to edit, then modify the
// objects in the services array and pass the modified object to this function. Services cannot be
// deleted in this manner, you must call deleteObject() on the service to physically remove them from
// the load balancer.
func (softlayer_network_loadbalancer_virtualipaddress *SoftLayer_Network_LoadBalancer_VirtualIpAddress) EditObject(ctx *slapi.RequestContext, templateObject SoftLayer_Network_LoadBalancer_VirtualIpAddress) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// Enable - Enable a disabled Virtual IP Address, allowing connections back to the IP address.
func (softlayer_network_loadbalancer_virtualipaddress *SoftLayer_Network_LoadBalancer_VirtualIpAddress) Enable(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Network_LoadBalancer_VirtualIpAddress object whose ID
// number corresponds to the ID number of the init parameter passed to the
// SoftLayer_Network_LoadBalancer_VirtualIpAddress service. You can only retrieve Load Balancers
// assigned to your account.
func (softlayer_network_loadbalancer_virtualipaddress *SoftLayer_Network_LoadBalancer_VirtualIpAddress) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_LoadBalancer_VirtualIpAddress, error) {
	var returnValue *SoftLayer_Network_LoadBalancer_VirtualIpAddress
	return returnValue, nil
}

// KickAllConnections - Quickly remove all active external connections to a Virtual IP Address.
func (softlayer_network_loadbalancer_virtualipaddress *SoftLayer_Network_LoadBalancer_VirtualIpAddress) KickAllConnections(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// UpgradeConnectionLimit - Upgrades the connection limit on the VirtualIp and changes the billing item
// on your account to reflect the change. This function will only upgrade you to the next "level" of
// service. The next level follows this pattern Current Level => Next Level 50 100 100 200 200 500 500
// 1000 1000 1200 1200 1500 1500 2000 2000 2500 2500 3000
func (softlayer_network_loadbalancer_virtualipaddress *SoftLayer_Network_LoadBalancer_VirtualIpAddress) UpgradeConnectionLimit(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
