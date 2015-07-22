package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_Application_Delivery_Controller - The
// SoftLayer_Network_Application_Delivery_Controller data type models a single instance of an
// application delivery controller. Local properties are read only, except for a ''notes'' property,
// which can be used to describe your application delivery controller service. The type's relational
// properties provide more information to the service's function and login information to the
// controller's backend management if advanced view is enabled.
type SoftLayer_Network_Application_Delivery_Controller struct {

	// Account - The SoftLayer customer account that owns an application delivery controller record.
	Account *SoftLayer_Account `json:"account"`

	// AccountId - The unique identifier of the SoftLayer customer account that owns an application
	// delivery controller record
	AccountId int `json:"accountId"`

	// AverageDailyPublicBandwidthUsage - The average daily public bandwidth usage for the current billing
	// cycle.
	AverageDailyPublicBandwidthUsage float32 `json:"averageDailyPublicBandwidthUsage"`

	// BillingItem - The billing item for a Application Delivery Controller.
	BillingItem *SoftLayer_Billing_Item_Network_Application_Delivery_Controller `json:"billingItem"`

	// ConfigurationHistory - Previous configurations for an Application Delivery Controller.
	ConfigurationHistory []*SoftLayer_Network_Application_Delivery_Controller_Configuration_History `json:"configurationHistory"`

	// ConfigurationHistoryCount - A count of previous configurations for an Application Delivery
	// Controller.
	ConfigurationHistoryCount uint64 `json:"configurationHistoryCount"`

	// CreateDate - The date that an application delivery controller record was created
	CreateDate *time.Time `json:"createDate"`

	// Datacenter - The datacenter that the application delivery controller resides in.
	Datacenter *SoftLayer_Location `json:"datacenter"`

	// Description - A brief description of an application delivery controller record.
	Description string `json:"description"`

	// Id - An application delivery controller's unique identifier
	Id int `json:"id"`

	// LicenseExpirationDate - The date in which the license for this application delivery controller will
	// expire.
	LicenseExpirationDate *time.Time `json:"licenseExpirationDate"`

	// LoadBalancerCount - A count of the virtual IP address records that belong to an application delivery
	// controller based load balancer.
	LoadBalancerCount uint64 `json:"loadBalancerCount"`

	// LoadBalancers - The virtual IP address records that belong to an application delivery controller
	// based load balancer.
	LoadBalancers []*SoftLayer_Network_LoadBalancer_VirtualIpAddress `json:"loadBalancers"`

	// ManagedResourceFlag - A flag indicating that this Application Delivery Controller is a managed
	// resource.
	ManagedResourceFlag bool `json:"managedResourceFlag"`

	// ManagementIpAddress - An application delivery controller's management ip address.
	ManagementIpAddress string `json:"managementIpAddress"`

	// ModifyDate - The date that an application delivery controller record was last modified
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`

	// NetworkVlan - The network that an application delivery controller resides on.
	NetworkVlan *SoftLayer_Network_Vlan `json:"networkVlan"`

	// NetworkVlanCount - A count of the network VLANs that an application delivery controller resides on.
	NetworkVlanCount uint64 `json:"networkVlanCount"`

	// NetworkVlans - The network VLANs that an application delivery controller resides on.
	NetworkVlans []*SoftLayer_Network_Vlan `json:"networkVlans"`

	// Notes - Editable notes used to describe an application delivery controller's function
	Notes string `json:"notes"`

	// OutboundPublicBandwidthUsage - The total public outbound bandwidth for the current billing cycle.
	OutboundPublicBandwidthUsage float64 `json:"outboundPublicBandwidthUsage"`

	// Password - The password used to connect to an application delivery controller's management interface
	// when it is operating in advanced view mode.
	Password *SoftLayer_Software_Component_Password `json:"password"`

	// PrimaryIpAddress - An application delivery controller's primary public IP address.
	PrimaryIpAddress string `json:"primaryIpAddress"`

	// ProjectedPublicBandwidthUsage - The projected public outbound bandwidth for the current billing
	// cycle.
	ProjectedPublicBandwidthUsage float32 `json:"projectedPublicBandwidthUsage"`

	// SubnetCount - A count of a network application controller's subnets. A subnet is a group of IP
	// addresses
	SubnetCount uint64 `json:"subnetCount"`

	// Subnets - A network application controller's subnets. A subnet is a group of IP addresses
	Subnets []*SoftLayer_Network_Subnet `json:"subnets"`

	// TagReferenceCount - no documentation
	TagReferenceCount uint64 `json:"tagReferenceCount"`

	// TagReferences - <nil>
	TagReferences []*SoftLayer_Tag_Reference `json:"tagReferences"`

	// Type - <nil>
	Type *SoftLayer_Network_Application_Delivery_Controller_Type `json:"type"`

	// TypeId - <nil>
	TypeId int `json:"typeId"`

	// VirtualIpAddressCount - no documentation
	VirtualIpAddressCount uint64 `json:"virtualIpAddressCount"`

	// VirtualIpAddresses - <nil>
	VirtualIpAddresses []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress `json:"virtualIpAddresses"`
}

// CreateLiveLoadBalancer - Create or add to an application delivery controller based load balancer
// service. The loadBalancer parameter must have its ''name'', ''type'', ''sourcePort'', and
// ''virtualIpAddress'' properties populated. Changes are reflected immediately in the application
// delivery controller.
func (softlayer_network_application_delivery_controller *SoftLayer_Network_Application_Delivery_Controller) CreateLiveLoadBalancer(ctx *slapi.RequestContext, loadBalancer SoftLayer_Network_LoadBalancer_VirtualIpAddress) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeleteLiveLoadBalancer - Remove a virtual IP address from an application delivery controller based
// load balancer. Only the ''name'' property in the loadBalancer parameter must be populated. Changes
// are reflected immediately in the application delivery controller.
func (softlayer_network_application_delivery_controller *SoftLayer_Network_Application_Delivery_Controller) DeleteLiveLoadBalancer(ctx *slapi.RequestContext, loadBalancer SoftLayer_Network_LoadBalancer_VirtualIpAddress) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeleteLiveLoadBalancerService - Remove an entire load balancer service, including all virtual IP
// addresses, from and application delivery controller based load balancer. The ''name'' property the
// and ''name'' property within the ''vip'' property of the service parameter must be provided. Changes
// are reflected immediately in the application delivery controller.
func (softlayer_network_application_delivery_controller *SoftLayer_Network_Application_Delivery_Controller) DeleteLiveLoadBalancerService(ctx *slapi.RequestContext, service SoftLayer_Network_LoadBalancer_Service) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - Edit an applications delivery controller record. Currently only a controller's notes
// property is editable.
func (softlayer_network_application_delivery_controller *SoftLayer_Network_Application_Delivery_Controller) EditObject(ctx *slapi.RequestContext, templateObject SoftLayer_Network_Application_Delivery_Controller) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetBandwidthDataByDate - <nil>
func (softlayer_network_application_delivery_controller *SoftLayer_Network_Application_Delivery_Controller) GetBandwidthDataByDate(ctx *slapi.RequestContext, startDateTime time.Time, endDateTime time.Time, networkType string) ([]*SoftLayer_Metric_Tracking_Object_Data, error) {
	var returnValue []*SoftLayer_Metric_Tracking_Object_Data
	return returnValue, nil
}

// GetBandwidthImageByDate - Use this method when needing a bandwidth image for a single application
// delivery controller. It will gather the correct input parameters for the generic graphing utility
// based on the date ranges
func (softlayer_network_application_delivery_controller *SoftLayer_Network_Application_Delivery_Controller) GetBandwidthImageByDate(ctx *slapi.RequestContext, startDateTime time.Time, endDateTime time.Time, networkType string) (*SoftLayer_Container_Bandwidth_GraphOutputs, error) {
	var returnValue *SoftLayer_Container_Bandwidth_GraphOutputs
	return returnValue, nil
}

// GetCustomBandwidthDataByDate - no documentation
func (softlayer_network_application_delivery_controller *SoftLayer_Network_Application_Delivery_Controller) GetCustomBandwidthDataByDate(ctx *slapi.RequestContext, graphData SoftLayer_Container_Graph) (*SoftLayer_Container_Graph, error) {
	var returnValue *SoftLayer_Container_Graph
	return returnValue, nil
}

// GetLiveLoadBalancerServiceGraphImage - Get the graph image for an application delivery controller
// service based on the supplied graph type and metric. The available graph types are: 'connections'
// and 'status', and the available metrics are: 'day', 'week' and 'month'. This method returns the raw
// binary image data.
func (softlayer_network_application_delivery_controller *SoftLayer_Network_Application_Delivery_Controller) GetLiveLoadBalancerServiceGraphImage(ctx *slapi.RequestContext, service SoftLayer_Network_LoadBalancer_Service, graphType string, metric string) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Network_Application_Delivery_Controller object whose
// ID number corresponds to the ID number of the init parameter passed to the
// SoftLayer_Network_Application_Delivery_Controller service. You can only retrieve application
// delivery controllers that are associated with your SoftLayer customer account.
func (softlayer_network_application_delivery_controller *SoftLayer_Network_Application_Delivery_Controller) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Application_Delivery_Controller, error) {
	var returnValue *SoftLayer_Network_Application_Delivery_Controller
	return returnValue, nil
}

// RestoreBaseConfiguration - Restore an application delivery controller's base configuration state.
// The configuration will be set to what it was when initially provisioned.
func (softlayer_network_application_delivery_controller *SoftLayer_Network_Application_Delivery_Controller) RestoreBaseConfiguration(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RestoreConfiguration - Restore an application delivery controller's configuration state.
func (softlayer_network_application_delivery_controller *SoftLayer_Network_Application_Delivery_Controller) RestoreConfiguration(ctx *slapi.RequestContext, configurationHistoryId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// SaveCurrentConfiguration - Save an application delivery controller's configuration state. The notes
// property for this method is optional.
func (softlayer_network_application_delivery_controller *SoftLayer_Network_Application_Delivery_Controller) SaveCurrentConfiguration(ctx *slapi.RequestContext, notes string) (*SoftLayer_Network_Application_Delivery_Controller_Configuration_History, error) {
	var returnValue *SoftLayer_Network_Application_Delivery_Controller_Configuration_History
	return returnValue, nil
}

// UpdateLiveLoadBalancer - Update the the virtual IP address interface within an application delivery
// controller based load balancer identified by the ''name'' property in the loadBalancer parameter.
// You only need to set the properties in the loadBalancer parameter that you wish to change. Any
// virtual IP properties omitted or left empty are ignored. Changes are reflected immediately in the
// application delivery controller.
func (softlayer_network_application_delivery_controller *SoftLayer_Network_Application_Delivery_Controller) UpdateLiveLoadBalancer(ctx *slapi.RequestContext, loadBalancer SoftLayer_Network_LoadBalancer_VirtualIpAddress) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// UpdateNetScalerLicense - Update the NetScaler VPX License. This service will create a transaction to
// update a NetScaler VPX License. After the license is updated the load balancer will reboot in order
// to apply the newly issued license The load balancer will be unavailable during the reboot.
func (softlayer_network_application_delivery_controller *SoftLayer_Network_Application_Delivery_Controller) UpdateNetScalerLicense(ctx *slapi.RequestContext) (*SoftLayer_Provisioning_Version1_Transaction, error) {
	var returnValue *SoftLayer_Provisioning_Version1_Transaction
	return returnValue, nil
}
