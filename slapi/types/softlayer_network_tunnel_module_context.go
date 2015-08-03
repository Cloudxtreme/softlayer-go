package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Tunnel_Module_Context - The SoftLayer_Network_Tunnel_Module_Context data type
// contains general information relating to a single SoftLayer network tunnel. The
// SoftLayer_Network_Tunnel_Module_Context is useful to gather information such as related customer
// subnets (remote) and internal subnets (local) associated with the network tunnel as well as other
// information needed to manage the network tunnel. Account and billing information related to the
// network tunnel can also be retrieved.
type SoftLayer_Network_Tunnel_Module_Context struct {

	// CustomerPeerIpAddress - The remote end of a network tunnel. This end of the network tunnel resides
	// on an outside network and will be sending and receiving the IPSec packets.
	CustomerPeerIpAddress string `json:"customerPeerIpAddress"`

	// Id - no documentation
	Id int `json:"id"`

	// Name - A network tunnel's unique name used on the network device.
	Name string `json:"name"`

	// PhaseTwoDiffieHellmanGroup - Determines the strength of the key used in the key exchange process.
	// The higher the group number the stronger the key is and the more secure it is. However, processing
	// time will increase as the strength of the key increases. Both peers must use the Diffie-Hellman
	// Group.
	PhaseTwoDiffieHellmanGroup int `json:"phaseTwoDiffieHellmanGroup"`

	// PhaseTwoPerfectForwardSecrecy - Determines if the generated keys are made from previous keys. When
	// PFS is specified, a Diffie-Hellman exchange occurs each time a new security association is
	// negotiated.
	PhaseTwoPerfectForwardSecrecy int `json:"phaseTwoPerfectForwardSecrecy"`

	// PhaseOneAuthentication - Authentication used to generate keys for protecting the negotiations for a
	// network tunnel.
	PhaseOneAuthentication string `json:"phaseOneAuthentication"`

	// PhaseTwoAuthentication - The authentication used in phase 2 proposal negotiation process.
	PhaseTwoAuthentication string `json:"phaseTwoAuthentication"`

	// PhaseTwoEncryption - The encryption used in phase 2 proposal negotiation process.
	PhaseTwoEncryption string `json:"phaseTwoEncryption"`

	// FriendlyName - no documentation
	FriendlyName string `json:"friendlyName"`

	// ModifyDate - The date a network tunnel was last modified. This date should NOT be used to determine
	// when the network tunnel configurations were last applied to the network device.
	ModifyDate *time.Time `json:"modifyDate"`

	// PhaseOneDiffieHellmanGroup - Determines the strength of the key used in the key exchange process.
	// The higher the group number the stronger the key is and the more secure it is. However, processing
	// time will increase as the strength of the key increases. Both peers in the must use the
	// Diffie-Hellman Group.
	PhaseOneDiffieHellmanGroup int `json:"phaseOneDiffieHellmanGroup"`

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// PhaseOneEncryption - Encryption used to generate keys for protecting the negotiations for a network
	// tunnel.
	PhaseOneEncryption string `json:"phaseOneEncryption"`

	// PhaseOneKeylife - Amount of time (in seconds) allowed to pass before the encryption key expires. A
	// new key is generated without interrupting service. Valid times are from 120 to 172800 seconds.
	PhaseOneKeylife int `json:"phaseOneKeylife"`

	// PhaseTwoKeylife - Amount of time (in seconds) allowed to pass before the encryption key expires. A
	// new key is generated without interrupting service. Valid times are from 120 to 172800 seconds.
	PhaseTwoKeylife int `json:"phaseTwoKeylife"`

	// PresharedKey - A key used so that peers authenticate each other. This key is hashed by using the
	// phase one encryption and phase one authentication.
	PresharedKey string `json:"presharedKey"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// AdvancedConfigurationFlag - A flag used to specify when advanced configurations, complex
	// configurations that require manual setup, are being applied to network devices for a network tunnel.
	// When the flag is set to true (1), a network tunnel cannot be configured through the management
	// portal nor the
	AdvancedConfigurationFlag int `json:"advancedConfigurationFlag"`

	// InternalPeerIpAddress - The local end of a network tunnel. This end of the network tunnel resides on
	// the SoftLayer networks and allows access to remote end of the tunnel to subnets on SoftLayer
	// networks.
	InternalPeerIpAddress string `json:"internalPeerIpAddress"`
}

func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context) String() string {
	return "SoftLayer_Network_Tunnel_Module_Context"
}

// SoftLayer_Network_Tunnel_Module_Context_Extended is SoftLayer_Network_Tunnel_Module_Context with all maskable types.
type SoftLayer_Network_Tunnel_Module_Context_Extended struct {
	SoftLayer_Network_Tunnel_Module_Context

	// StaticRouteSubnetCount - A count of subnets used for a network tunnel's address translations.
	StaticRouteSubnetCount uint64 `json:"staticRouteSubnetCount"`

	// AllAvailableServiceSubnets - Subnets that provide access to SoftLayer services such as the
	// management portal and the SoftLayer
	AllAvailableServiceSubnets []*SoftLayer_Network_Subnet `json:"allAvailableServiceSubnets"`

	// CustomerSubnets - Remote subnets that are allowed access through a network tunnel.
	CustomerSubnets []*SoftLayer_Network_Customer_Subnet `json:"customerSubnets"`

	// ServiceSubnets - Service subnets that can be access through the network tunnel.
	ServiceSubnets []*SoftLayer_Network_Subnet `json:"serviceSubnets"`

	// AddressTranslations - no documentation
	AddressTranslations []*SoftLayer_Network_Tunnel_Module_Context_Address_Translation `json:"addressTranslations"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// CustomerSubnetCount - A count of remote subnets that are allowed access through a network tunnel.
	CustomerSubnetCount uint64 `json:"customerSubnetCount"`

	// StaticRouteSubnets - Subnets used for a network tunnel's address translations.
	StaticRouteSubnets []*SoftLayer_Network_Subnet `json:"staticRouteSubnets"`

	// AddressTranslationCount - A count of a network tunnel's address translations.
	AddressTranslationCount uint64 `json:"addressTranslationCount"`

	// InternalSubnetCount - A count of private subnets that can be accessed through the network tunnel.
	InternalSubnetCount uint64 `json:"internalSubnetCount"`

	// ActiveTransaction - The transaction that is currently applying configurations for the network
	// tunnel.
	ActiveTransaction *SoftLayer_Provisioning_Version1_Transaction `json:"activeTransaction"`

	// Datacenter - The datacenter location for one end of the network tunnel that allows access to
	// account's private subnets.
	Datacenter *SoftLayer_Location `json:"datacenter"`

	// TransactionHistoryCount - A count of the transaction history for this network tunnel.
	TransactionHistoryCount uint64 `json:"transactionHistoryCount"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// TransactionHistory - no documentation
	TransactionHistory []*SoftLayer_Provisioning_Version1_Transaction `json:"transactionHistory"`

	// ServiceSubnetCount - A count of service subnets that can be access through the network tunnel.
	ServiceSubnetCount uint64 `json:"serviceSubnetCount"`

	// InternalSubnets - Private subnets that can be accessed through the network tunnel.
	InternalSubnets []*SoftLayer_Network_Subnet `json:"internalSubnets"`

	// AllAvailableServiceSubnetCount - A count of subnets that provide access to SoftLayer services such
	// as the management portal and the SoftLayer
	AllAvailableServiceSubnetCount uint64 `json:"allAvailableServiceSubnetCount"`
}

func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context_Extended) String() string {
	return "SoftLayer_Network_Tunnel_Module_Context"
}
