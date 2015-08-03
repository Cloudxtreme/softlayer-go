package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Application_Delivery_Controller - The
// SoftLayer_Network_Application_Delivery_Controller data type models a single instance of an
// application delivery controller. Local properties are read only, except for a ''notes'' property,
// which can be used to describe your application delivery controller service. The type's relational
// properties provide more information to the service's function and login information to the
// controller's backend management if advanced view is enabled.
type SoftLayer_Network_Application_Delivery_Controller struct {

	// AccountId - The unique identifier of the SoftLayer customer account that owns an application
	// delivery controller record
	AccountId int `json:"accountId"`

	// Notes - Editable notes used to describe an application delivery controller's function
	Notes string `json:"notes"`

	// ModifyDate - The date that an application delivery controller record was last modified
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`

	// TypeId - <nil>
	TypeId int `json:"typeId"`

	// Id - An application delivery controller's unique identifier
	Id int `json:"id"`

	// CreateDate - The date that an application delivery controller record was created
	CreateDate *time.Time `json:"createDate"`
}

// SoftLayer_Network_Application_Delivery_Controller_Extended is SoftLayer_Network_Application_Delivery_Controller with all maskable types.
type SoftLayer_Network_Application_Delivery_Controller_Extended struct {
	SoftLayer_Network_Application_Delivery_Controller

	// ManagedResourceFlag - A flag indicating that this Application Delivery Controller is a managed
	// resource.
	ManagedResourceFlag bool `json:"managedResourceFlag"`

	// AverageDailyPublicBandwidthUsage - The average daily public bandwidth usage for the current billing
	// cycle.
	AverageDailyPublicBandwidthUsage float32 `json:"averageDailyPublicBandwidthUsage"`

	// ConfigurationHistoryCount - A count of previous configurations for an Application Delivery
	// Controller.
	ConfigurationHistoryCount uint64 `json:"configurationHistoryCount"`

	// LoadBalancerCount - A count of the virtual IP address records that belong to an application delivery
	// controller based load balancer.
	LoadBalancerCount uint64 `json:"loadBalancerCount"`

	// VirtualIpAddresses - <nil>
	VirtualIpAddresses []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress `json:"virtualIpAddresses"`

	// NetworkVlans - The network VLANs that an application delivery controller resides on.
	NetworkVlans []*SoftLayer_Network_Vlan `json:"networkVlans"`

	// OutboundPublicBandwidthUsage - The total public outbound bandwidth for the current billing cycle.
	OutboundPublicBandwidthUsage float64 `json:"outboundPublicBandwidthUsage"`

	// PrimaryIpAddress - An application delivery controller's primary public IP address.
	PrimaryIpAddress string `json:"primaryIpAddress"`

	// ProjectedPublicBandwidthUsage - The projected public outbound bandwidth for the current billing
	// cycle.
	ProjectedPublicBandwidthUsage float32 `json:"projectedPublicBandwidthUsage"`

	// TagReferences - <nil>
	TagReferences []*SoftLayer_Tag_Reference `json:"tagReferences"`

	// SubnetCount - A count of a network application controller's subnets. A subnet is a group of IP
	// addresses
	SubnetCount uint64 `json:"subnetCount"`

	// Datacenter - The datacenter that the application delivery controller resides in.
	Datacenter *SoftLayer_Location `json:"datacenter"`

	// ManagementIpAddress - An application delivery controller's management ip address.
	ManagementIpAddress string `json:"managementIpAddress"`

	// ConfigurationHistory - Previous configurations for an Application Delivery Controller.
	ConfigurationHistory []*SoftLayer_Network_Application_Delivery_Controller_Configuration_History `json:"configurationHistory"`

	// Password - The password used to connect to an application delivery controller's management interface
	// when it is operating in advanced view mode.
	Password *SoftLayer_Software_Component_Password `json:"password"`

	// Type - <nil>
	Type *SoftLayer_Network_Application_Delivery_Controller_Type `json:"type"`

	// Account - The SoftLayer customer account that owns an application delivery controller record.
	Account *SoftLayer_Account `json:"account"`

	// Description - A brief description of an application delivery controller record.
	Description string `json:"description"`

	// LicenseExpirationDate - The date in which the license for this application delivery controller will
	// expire.
	LicenseExpirationDate *time.Time `json:"licenseExpirationDate"`

	// Subnets - A network application controller's subnets. A subnet is a group of IP addresses
	Subnets []*SoftLayer_Network_Subnet `json:"subnets"`

	// TagReferenceCount - no documentation
	TagReferenceCount uint64 `json:"tagReferenceCount"`

	// VirtualIpAddressCount - no documentation
	VirtualIpAddressCount uint64 `json:"virtualIpAddressCount"`

	// BillingItem - The billing item for a Application Delivery Controller.
	BillingItem *SoftLayer_Billing_Item_Network_Application_Delivery_Controller `json:"billingItem"`

	// NetworkVlan - The network that an application delivery controller resides on.
	NetworkVlan *SoftLayer_Network_Vlan `json:"networkVlan"`

	// NetworkVlanCount - A count of the network VLANs that an application delivery controller resides on.
	NetworkVlanCount uint64 `json:"networkVlanCount"`

	// LoadBalancers - The virtual IP address records that belong to an application delivery controller
	// based load balancer.
	LoadBalancers []*SoftLayer_Network_LoadBalancer_VirtualIpAddress `json:"loadBalancers"`
}

func (softlayer_network_application_delivery_controller *SoftLayer_Network_Application_Delivery_Controller) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller"
}
