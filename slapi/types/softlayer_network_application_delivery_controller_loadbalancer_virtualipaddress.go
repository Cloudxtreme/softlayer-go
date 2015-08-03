package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress - <nil>
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress struct {

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// AccountId - The unique identifier of the SoftLayer customer account that owns the virtual IP address
	AccountId int `json:"accountId"`

	// ApplicationDeliveryController - A virtual IP address's associated application delivery controller.
	ApplicationDeliveryController *SoftLayer_Network_Application_Delivery_Controller `json:"applicationDeliveryController"`

	// ApplicationDeliveryControllerCount - A count of a virtual IP address's associated application
	// delivery controllers.
	ApplicationDeliveryControllerCount uint64 `json:"applicationDeliveryControllerCount"`

	// ApplicationDeliveryControllers - A virtual IP address's associated application delivery controllers.
	ApplicationDeliveryControllers []*SoftLayer_Network_Application_Delivery_Controller `json:"applicationDeliveryControllers"`

	// BillingItem - The current billing item for the load balancer virtual IP. This is only valid when
	// dedicatedFlag is false. This is an independent virtual IP, and if canceled, will only affect the
	// associated virtual
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// ConnectionLimit - no documentation
	ConnectionLimit int `json:"connectionLimit"`

	// ConnectionLimitUnits - no documentation
	ConnectionLimitUnits string `json:"connectionLimitUnits"`

	// DedicatedBillingItem - The current billing item for the load balancing device housing the virtual
	// IP. This billing item represents a device which could contain other virtual IPs. Caution should be
	// taken when canceling. This is only valid when dedicatedFlag is true.
	DedicatedBillingItem *SoftLayer_Billing_Item_Network_LoadBalancer `json:"dedicatedBillingItem"`

	// DedicatedFlag - A flag that determines if a VIP is dedicated or not. This is used to override the
	// connection limit and use an unlimited value.
	DedicatedFlag bool `json:"dedicatedFlag"`

	// HighAvailabilityFlag - Denotes whether the virtual IP is configured within a high availability
	// cluster.
	HighAvailabilityFlag bool `json:"highAvailabilityFlag"`

	// Id - The unique identifier of the virtual IP address record
	Id int `json:"id"`

	// IpAddress - <nil>
	IpAddress *SoftLayer_Network_Subnet_IpAddress `json:"ipAddress"`

	// IpAddressId - no documentation
	IpAddressId int `json:"ipAddressId"`

	// LoadBalancerHardware - <nil>
	LoadBalancerHardware []*SoftLayer_Hardware `json:"loadBalancerHardware"`

	// LoadBalancerHardwareCount - no documentation
	LoadBalancerHardwareCount uint64 `json:"loadBalancerHardwareCount"`

	// ManagedResourceFlag - A flag indicating that the load balancer is a managed resource.
	ManagedResourceFlag bool `json:"managedResourceFlag"`

	// Notes - User-created notes for this load balancer virtual IP address
	Notes string `json:"notes"`

	// SecureTransportProtocolCount - A count of the list of secure transport protocols enabled for this
	// virtual IP address
	SecureTransportProtocolCount uint64 `json:"secureTransportProtocolCount"`

	// SecureTransportProtocols - The list of secure transport protocols enabled for this virtual IP
	// address
	SecureTransportProtocols []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportProtocol `json:"secureTransportProtocols"`

	// SecurityCertificate - no documentation
	SecurityCertificate *SoftLayer_Security_Certificate `json:"securityCertificate"`

	// SecurityCertificateEntry - The SSL certificate currently associated with the Provides chosen
	// certificate visibility to unprivileged users.
	SecurityCertificateEntry *SoftLayer_Security_Certificate_Entry `json:"securityCertificateEntry"`

	// SecurityCertificateId - The unique identifier of the Security Certificate to be utilized when SSL
	// support is enabled.
	SecurityCertificateId int `json:"securityCertificateId"`

	// SslActiveFlag - Determines if the VIP currently has SSL acceleration enabled
	SslActiveFlag bool `json:"sslActiveFlag"`

	// SslEnabledFlag - Determines if the VIP is _allowed_ to utilize SSL acceleration
	SslEnabledFlag bool `json:"sslEnabledFlag"`

	// VirtualServerCount - no documentation
	VirtualServerCount uint64 `json:"virtualServerCount"`

	// VirtualServers - <nil>
	VirtualServers []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer `json:"virtualServers"`
}

func (softlayer_network_application_delivery_controller_loadbalancer_virtualipaddress *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress"
}

// EditObject - Like any other API object, the load balancers can have their exposed properties edited
// by passing in a modified version of the object. The load balancer object also can modify its
// services in this way. Simply request the load balancer object you wish to edit, then modify the
// objects in the services array and pass the modified object to this function. Services cannot be
// deleted in this manner, you must call deleteObject() on the service to physically remove them from
// the load balancer.
func (softlayer_network_application_delivery_controller_loadbalancer_virtualipaddress *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) EditObject(ctx *slapi.RequestContext, templateObject SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetAvailableSecureTransportProtocols - <nil>
func (softlayer_network_application_delivery_controller_loadbalancer_virtualipaddress *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) GetAvailableSecureTransportProtocols(ctx *slapi.RequestContext) ([]*SoftLayer_Security_SecureTransportProtocol, error) {
	var returnValue []*SoftLayer_Security_SecureTransportProtocol
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_application_delivery_controller_loadbalancer_virtualipaddress *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress, error) {
	var returnValue *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress
	return returnValue, nil
}

// StartSsl - Start SSL acceleration on all SSL virtual services (those with a type of This action
// should be taken only after configuring an SSL certificate for the virtual IP.
func (softlayer_network_application_delivery_controller_loadbalancer_virtualipaddress *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) StartSsl(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// StopSsl - Stop SSL acceleration on all SSL virtual services (those with a type of
func (softlayer_network_application_delivery_controller_loadbalancer_virtualipaddress *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) StopSsl(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// UpgradeConnectionLimit - Upgrades the connection limit on the Virtual IP to Address to the next,
// higher connection limit of the same product.
func (softlayer_network_application_delivery_controller_loadbalancer_virtualipaddress *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) UpgradeConnectionLimit(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
