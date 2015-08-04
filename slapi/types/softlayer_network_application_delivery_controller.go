package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

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
	AccountId int `json:"accountId,omitempty"`

	// Id - An application delivery controller's unique identifier
	Id int `json:"id,omitempty"`

	// ModifyDate - The date that an application delivery controller record was last modified
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// CreateDate - The date that an application delivery controller record was created
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Notes - Editable notes used to describe an application delivery controller's function
	Notes string `json:"notes,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// TypeId - <nil>
	TypeId int `json:"typeId,omitempty"`

	// Type - <nil>
	Type *SoftLayer_Network_Application_Delivery_Controller_Type `json:"type,omitempty"`

	// Description - A brief description of an application delivery controller record.
	Description string `json:"description,omitempty"`

	// PrimaryIpAddress - An application delivery controller's primary public IP address.
	PrimaryIpAddress string `json:"primaryIpAddress,omitempty"`

	// TagReferences - <nil>
	TagReferences []*SoftLayer_Tag_Reference `json:"tagReferences,omitempty"`

	// VirtualIpAddressCount - no documentation
	VirtualIpAddressCount uint64 `json:"virtualIpAddressCount,omitempty"`

	// OutboundPublicBandwidthUsage - The total public outbound bandwidth for the current billing cycle.
	OutboundPublicBandwidthUsage slapi.Float64 `json:"outboundPublicBandwidthUsage,omitempty"`

	// ConfigurationHistory - Previous configurations for an Application Delivery Controller.
	ConfigurationHistory []*SoftLayer_Network_Application_Delivery_Controller_Configuration_History `json:"configurationHistory,omitempty"`

	// LoadBalancers - The virtual IP address records that belong to an application delivery controller
	// based load balancer.
	LoadBalancers []*SoftLayer_Network_LoadBalancer_VirtualIpAddress `json:"loadBalancers,omitempty"`

	// ManagedResourceFlag - A flag indicating that this Application Delivery Controller is a managed
	// resource.
	ManagedResourceFlag bool `json:"managedResourceFlag,omitempty"`

	// SubnetCount - A count of a network application controller's subnets. A subnet is a group of IP
	// addresses
	SubnetCount uint64 `json:"subnetCount,omitempty"`

	// NetworkVlan - The network that an application delivery controller resides on.
	NetworkVlan *SoftLayer_Network_Vlan `json:"networkVlan,omitempty"`

	// Password - The password used to connect to an application delivery controller's management interface
	// when it is operating in advanced view mode.
	Password *SoftLayer_Software_Component_Password `json:"password,omitempty"`

	// ProjectedPublicBandwidthUsage - The projected public outbound bandwidth for the current billing
	// cycle.
	ProjectedPublicBandwidthUsage slapi.Float64 `json:"projectedPublicBandwidthUsage,omitempty"`

	// ConfigurationHistoryCount - A count of previous configurations for an Application Delivery
	// Controller.
	ConfigurationHistoryCount uint64 `json:"configurationHistoryCount,omitempty"`

	// NetworkVlanCount - A count of the network VLANs that an application delivery controller resides on.
	NetworkVlanCount uint64 `json:"networkVlanCount,omitempty"`

	// Account - The SoftLayer customer account that owns an application delivery controller record.
	Account *SoftLayer_Account `json:"account,omitempty"`

	// Subnets - A network application controller's subnets. A subnet is a group of IP addresses
	Subnets []*SoftLayer_Network_Subnet `json:"subnets,omitempty"`

	// BillingItem - The billing item for a Application Delivery Controller.
	BillingItem *SoftLayer_Billing_Item_Network_Application_Delivery_Controller `json:"billingItem,omitempty"`

	// LicenseExpirationDate - The date in which the license for this application delivery controller will
	// expire.
	LicenseExpirationDate *time.Time `json:"licenseExpirationDate,omitempty"`

	// ManagementIpAddress - An application delivery controller's management ip address.
	ManagementIpAddress string `json:"managementIpAddress,omitempty"`

	// LoadBalancerCount - A count of the virtual IP address records that belong to an application delivery
	// controller based load balancer.
	LoadBalancerCount uint64 `json:"loadBalancerCount,omitempty"`

	// AverageDailyPublicBandwidthUsage - The average daily public bandwidth usage for the current billing
	// cycle.
	AverageDailyPublicBandwidthUsage slapi.Float64 `json:"averageDailyPublicBandwidthUsage,omitempty"`

	// Datacenter - The datacenter that the application delivery controller resides in.
	Datacenter *SoftLayer_Location `json:"datacenter,omitempty"`

	// NetworkVlans - The network VLANs that an application delivery controller resides on.
	NetworkVlans []*SoftLayer_Network_Vlan `json:"networkVlans,omitempty"`

	// VirtualIpAddresses - <nil>
	VirtualIpAddresses []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress `json:"virtualIpAddresses,omitempty"`

	// TagReferenceCount - no documentation
	TagReferenceCount uint64 `json:"tagReferenceCount,omitempty"`
}

func (softlayer_network_application_delivery_controller *SoftLayer_Network_Application_Delivery_Controller) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller"
}
