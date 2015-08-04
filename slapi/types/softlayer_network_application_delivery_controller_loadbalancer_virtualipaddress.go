package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress - <nil>
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress struct {

	// AccountId - The unique identifier of the SoftLayer customer account that owns the virtual IP address
	AccountId int `json:"accountId,omitempty"`

	// ConnectionLimit - no documentation
	ConnectionLimit int `json:"connectionLimit,omitempty"`

	// SecurityCertificateId - The unique identifier of the Security Certificate to be utilized when SSL
	// support is enabled.
	SecurityCertificateId int `json:"securityCertificateId,omitempty"`

	// SslEnabledFlag - Determines if the VIP is _allowed_ to utilize SSL acceleration
	SslEnabledFlag bool `json:"sslEnabledFlag,omitempty"`

	// Id - The unique identifier of the virtual IP address record
	Id int `json:"id,omitempty"`

	// IpAddressId - no documentation
	IpAddressId int `json:"ipAddressId,omitempty"`

	// DedicatedFlag - A flag that determines if a VIP is dedicated or not. This is used to override the
	// connection limit and use an unlimited value.
	DedicatedFlag bool `json:"dedicatedFlag,omitempty"`

	// Notes - User-created notes for this load balancer virtual IP address
	Notes string `json:"notes,omitempty"`

	// SslActiveFlag - Determines if the VIP currently has SSL acceleration enabled
	SslActiveFlag bool `json:"sslActiveFlag,omitempty"`

	// ConnectionLimitUnits - no documentation
	ConnectionLimitUnits string `json:"connectionLimitUnits,omitempty"`
}

func (softlayer_network_application_delivery_controller_loadbalancer_virtualipaddress *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress"
}

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_Extended is SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress with all maskable types.
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_Extended struct {
	SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress

	// LoadBalancerHardware - <nil>
	LoadBalancerHardware []*SoftLayer_Hardware `json:"loadBalancerHardware,omitempty"`

	// HighAvailabilityFlag - Denotes whether the virtual IP is configured within a high availability
	// cluster.
	HighAvailabilityFlag bool `json:"highAvailabilityFlag,omitempty"`

	// ApplicationDeliveryControllers - A virtual IP address's associated application delivery controllers.
	ApplicationDeliveryControllers []*SoftLayer_Network_Application_Delivery_Controller `json:"applicationDeliveryControllers,omitempty"`

	// BillingItem - The current billing item for the load balancer virtual IP. This is only valid when
	// dedicatedFlag is false. This is an independent virtual IP, and if canceled, will only affect the
	// associated virtual
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`

	// DedicatedBillingItem - The current billing item for the load balancing device housing the virtual
	// IP. This billing item represents a device which could contain other virtual IPs. Caution should be
	// taken when canceling. This is only valid when dedicatedFlag is true.
	DedicatedBillingItem *SoftLayer_Billing_Item_Network_LoadBalancer `json:"dedicatedBillingItem,omitempty"`

	// ApplicationDeliveryControllerCount - A count of a virtual IP address's associated application
	// delivery controllers.
	ApplicationDeliveryControllerCount uint64 `json:"applicationDeliveryControllerCount,omitempty"`

	// SecureTransportProtocolCount - A count of the list of secure transport protocols enabled for this
	// virtual IP address
	SecureTransportProtocolCount uint64 `json:"secureTransportProtocolCount,omitempty"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// LoadBalancerHardwareCount - no documentation
	LoadBalancerHardwareCount uint64 `json:"loadBalancerHardwareCount,omitempty"`

	// ApplicationDeliveryController - A virtual IP address's associated application delivery controller.
	ApplicationDeliveryController *SoftLayer_Network_Application_Delivery_Controller `json:"applicationDeliveryController,omitempty"`

	// ManagedResourceFlag - A flag indicating that the load balancer is a managed resource.
	ManagedResourceFlag bool `json:"managedResourceFlag,omitempty"`

	// VirtualServers - <nil>
	VirtualServers []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer `json:"virtualServers,omitempty"`

	// SecurityCertificateEntry - The SSL certificate currently associated with the Provides chosen
	// certificate visibility to unprivileged users.
	SecurityCertificateEntry *SoftLayer_Security_Certificate_Entry `json:"securityCertificateEntry,omitempty"`

	// VirtualServerCount - no documentation
	VirtualServerCount uint64 `json:"virtualServerCount,omitempty"`

	// IpAddress - <nil>
	IpAddress *SoftLayer_Network_Subnet_IpAddress `json:"ipAddress,omitempty"`

	// SecureTransportProtocols - The list of secure transport protocols enabled for this virtual IP
	// address
	SecureTransportProtocols []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_SecureTransportProtocol `json:"secureTransportProtocols,omitempty"`

	// SecurityCertificate - no documentation
	SecurityCertificate *SoftLayer_Security_Certificate `json:"securityCertificate,omitempty"`
}

func (softlayer_network_application_delivery_controller_loadbalancer_virtualipaddress *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_Extended) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress"
}
