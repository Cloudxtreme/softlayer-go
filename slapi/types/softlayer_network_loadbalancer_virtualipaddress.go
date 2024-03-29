package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_LoadBalancer_VirtualIpAddress - The
// SoftLayer_Network_LoadBalancer_VirtualIpAddress data type contains all the information relating to a
// specific load balancer assigned to a customer account. Information retained on the object itself is
// the virtual IP address, load balancing method, and any notes that are related to the load balancer.
// There is also an array of SoftLayer_Network_LoadBalancer_Service objects, which represent the load
// balancer services, explained more fully in the SoftLayer_Network_LoadBalancer_Service documentation.
type SoftLayer_Network_LoadBalancer_VirtualIpAddress struct {

	// Id - Unique ID for this object, used for the getObject method, and must be set if you are editing
	// this object.
	Id int `json:"id,omitempty"`

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
	LoadBalancingMethod string `json:"loadBalancingMethod,omitempty"`

	// LoadBalancingMethodFullName - A human readable version of loadBalancingMethod, intended mainly for
	// API users.
	LoadBalancingMethodFullName string `json:"loadBalancingMethodFullName,omitempty"`

	// SecurityCertificateId - The unique identifier of the Security Certificate to be utilized when SSL
	// support is enabled.
	SecurityCertificateId int `json:"securityCertificateId,omitempty"`

	// SourcePort - no documentation
	SourcePort int `json:"sourcePort,omitempty"`

	// Type - The connection type of this Valid values are and
	Type string `json:"type,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Notes - no documentation
	Notes string `json:"notes,omitempty"`

	// VirtualIpAddress - The virtual, public-facing IP address for your load balancer. This is the address
	// of all incoming traffic
	VirtualIpAddress string `json:"virtualIpAddress,omitempty"`

	// ConnectionLimit - Connection limit on this Can be upgraded through the upgradeConnectionLimit()
	// function
	ConnectionLimit int `json:"connectionLimit,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// ServiceCount - no documentation
	ServiceCount uint64 `json:"serviceCount,omitempty"`

	// ManagedResourceFlag - A flag indicating that the load balancer is a managed resource.
	ManagedResourceFlag bool `json:"managedResourceFlag,omitempty"`

	// Services - no documentation
	Services []*SoftLayer_Network_LoadBalancer_Service `json:"services,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`

	// CustomerManagedFlag - If false, this VIP and associated services may be edited via the portal or the
	// If true, you must configure this VIP manually on the device.
	CustomerManagedFlag int `json:"customerManagedFlag,omitempty"`
}

func (softlayer_network_loadbalancer_virtualipaddress *SoftLayer_Network_LoadBalancer_VirtualIpAddress) String() string {
	return "SoftLayer_Network_LoadBalancer_VirtualIpAddress"
}
