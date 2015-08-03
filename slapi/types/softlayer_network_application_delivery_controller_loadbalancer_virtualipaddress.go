package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress - <nil>
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress struct {

	// Notes - User-created notes for this load balancer virtual IP address
	Notes string `json:"notes"`

	// SslEnabledFlag - Determines if the VIP is _allowed_ to utilize SSL acceleration
	SslEnabledFlag bool `json:"sslEnabledFlag"`

	// Id - The unique identifier of the virtual IP address record
	Id int `json:"id"`

	// ConnectionLimit - no documentation
	ConnectionLimit int `json:"connectionLimit"`

	// ConnectionLimitUnits - no documentation
	ConnectionLimitUnits string `json:"connectionLimitUnits"`

	// DedicatedFlag - A flag that determines if a VIP is dedicated or not. This is used to override the
	// connection limit and use an unlimited value.
	DedicatedFlag bool `json:"dedicatedFlag"`

	// SecurityCertificateId - The unique identifier of the Security Certificate to be utilized when SSL
	// support is enabled.
	SecurityCertificateId int `json:"securityCertificateId"`

	// IpAddressId - no documentation
	IpAddressId int `json:"ipAddressId"`

	// SslActiveFlag - Determines if the VIP currently has SSL acceleration enabled
	SslActiveFlag bool `json:"sslActiveFlag"`

	// AccountId - The unique identifier of the SoftLayer customer account that owns the virtual IP address
	AccountId int `json:"accountId"`
}

func (softlayer_network_application_delivery_controller_loadbalancer_virtualipaddress *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress"
}

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_Extended is SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress with all maskable types.
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_Extended struct {
	SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// VirtualServers - <nil>
	VirtualServers []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer `json:"virtualServers"`

	// ApplicationDeliveryControllers - A virtual IP address's associated application delivery controllers.
	ApplicationDeliveryControllers []*SoftLayer_Network_Application_Delivery_Controller `json:"applicationDeliveryControllers"`

	// LoadBalancerHardware - <nil>
	LoadBalancerHardware []*SoftLayer_Hardware `json:"loadBalancerHardware"`

	// ManagedResourceFlag - A flag indicating that the load balancer is a managed resource.
	ManagedResourceFlag bool `json:"managedResourceFlag"`

	// HighAvailabilityFlag - Denotes whether the virtual IP is configured within a high availability
	// cluster.
	HighAvailabilityFlag bool `json:"highAvailabilityFlag"`

	// IpAddress - <nil>
	IpAddress *SoftLayer_Network_Subnet_IpAddress `json:"ipAddress"`

	// SecureTransportProtocols - The list of secure transport protocols enabled for this virtual IP
	// address
	SecureTransportProtocols []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportProtocol `json:"secureTransportProtocols"`

	// SecureTransportProtocolCount - A count of the list of secure transport protocols enabled for this
	// virtual IP address
	SecureTransportProtocolCount uint64 `json:"secureTransportProtocolCount"`

	// SecurityCertificate - no documentation
	SecurityCertificate *SoftLayer_Security_Certificate `json:"securityCertificate"`

	// VirtualServerCount - no documentation
	VirtualServerCount uint64 `json:"virtualServerCount"`

	// SecurityCertificateEntry - The SSL certificate currently associated with the Provides chosen
	// certificate visibility to unprivileged users.
	SecurityCertificateEntry *SoftLayer_Security_Certificate_Entry `json:"securityCertificateEntry"`

	// ApplicationDeliveryController - A virtual IP address's associated application delivery controller.
	ApplicationDeliveryController *SoftLayer_Network_Application_Delivery_Controller `json:"applicationDeliveryController"`

	// LoadBalancerHardwareCount - no documentation
	LoadBalancerHardwareCount uint64 `json:"loadBalancerHardwareCount"`

	// DedicatedBillingItem - The current billing item for the load balancing device housing the virtual
	// IP. This billing item represents a device which could contain other virtual IPs. Caution should be
	// taken when canceling. This is only valid when dedicatedFlag is true.
	DedicatedBillingItem *SoftLayer_Billing_Item_Network_LoadBalancer `json:"dedicatedBillingItem"`

	// BillingItem - The current billing item for the load balancer virtual IP. This is only valid when
	// dedicatedFlag is false. This is an independent virtual IP, and if canceled, will only affect the
	// associated virtual
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// ApplicationDeliveryControllerCount - A count of a virtual IP address's associated application
	// delivery controllers.
	ApplicationDeliveryControllerCount uint64 `json:"applicationDeliveryControllerCount"`
}

func (softlayer_network_application_delivery_controller_loadbalancer_virtualipaddress *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_Extended) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress"
}
