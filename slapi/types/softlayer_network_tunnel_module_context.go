package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_Tunnel_Module_Context - The SoftLayer_Network_Tunnel_Module_Context data type
// contains general information relating to a single SoftLayer network tunnel. The
// SoftLayer_Network_Tunnel_Module_Context is useful to gather information such as related customer
// subnets (remote) and internal subnets (local) associated with the network tunnel as well as other
// information needed to manage the network tunnel. Account and billing information related to the
// network tunnel can also be retrieved.
type SoftLayer_Network_Tunnel_Module_Context struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// ActiveTransaction - The transaction that is currently applying configurations for the network
	// tunnel.
	ActiveTransaction *SoftLayer_Provisioning_Version1_Transaction `json:"activeTransaction"`

	// AddressTranslationCount - A count of a network tunnel's address translations.
	AddressTranslationCount uint64 `json:"addressTranslationCount"`

	// AddressTranslations - no documentation
	AddressTranslations []*SoftLayer_Network_Tunnel_Module_Context_Address_Translation `json:"addressTranslations"`

	// AdvancedConfigurationFlag - A flag used to specify when advanced configurations, complex
	// configurations that require manual setup, are being applied to network devices for a network tunnel.
	// When the flag is set to true (1), a network tunnel cannot be configured through the management
	// portal nor the
	AdvancedConfigurationFlag int `json:"advancedConfigurationFlag"`

	// AllAvailableServiceSubnetCount - A count of subnets that provide access to SoftLayer services such
	// as the management portal and the SoftLayer
	AllAvailableServiceSubnetCount uint64 `json:"allAvailableServiceSubnetCount"`

	// AllAvailableServiceSubnets - Subnets that provide access to SoftLayer services such as the
	// management portal and the SoftLayer
	AllAvailableServiceSubnets []*SoftLayer_Network_Subnet `json:"allAvailableServiceSubnets"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// CustomerPeerIpAddress - The remote end of a network tunnel. This end of the network tunnel resides
	// on an outside network and will be sending and receiving the IPSec packets.
	CustomerPeerIpAddress string `json:"customerPeerIpAddress"`

	// CustomerSubnetCount - A count of remote subnets that are allowed access through a network tunnel.
	CustomerSubnetCount uint64 `json:"customerSubnetCount"`

	// CustomerSubnets - Remote subnets that are allowed access through a network tunnel.
	CustomerSubnets []*SoftLayer_Network_Customer_Subnet `json:"customerSubnets"`

	// Datacenter - The datacenter location for one end of the network tunnel that allows access to
	// account's private subnets.
	Datacenter *SoftLayer_Location `json:"datacenter"`

	// FriendlyName - no documentation
	FriendlyName string `json:"friendlyName"`

	// Id - no documentation
	Id int `json:"id"`

	// InternalPeerIpAddress - The local end of a network tunnel. This end of the network tunnel resides on
	// the SoftLayer networks and allows access to remote end of the tunnel to subnets on SoftLayer
	// networks.
	InternalPeerIpAddress string `json:"internalPeerIpAddress"`

	// InternalSubnetCount - A count of private subnets that can be accessed through the network tunnel.
	InternalSubnetCount uint64 `json:"internalSubnetCount"`

	// InternalSubnets - Private subnets that can be accessed through the network tunnel.
	InternalSubnets []*SoftLayer_Network_Subnet `json:"internalSubnets"`

	// ModifyDate - The date a network tunnel was last modified. This date should NOT be used to determine
	// when the network tunnel configurations were last applied to the network device.
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - A network tunnel's unique name used on the network device.
	Name string `json:"name"`

	// PhaseOneAuthentication - Authentication used to generate keys for protecting the negotiations for a
	// network tunnel.
	PhaseOneAuthentication string `json:"phaseOneAuthentication"`

	// PhaseOneDiffieHellmanGroup - Determines the strength of the key used in the key exchange process.
	// The higher the group number the stronger the key is and the more secure it is. However, processing
	// time will increase as the strength of the key increases. Both peers in the must use the
	// Diffie-Hellman Group.
	PhaseOneDiffieHellmanGroup int `json:"phaseOneDiffieHellmanGroup"`

	// PhaseOneEncryption - Encryption used to generate keys for protecting the negotiations for a network
	// tunnel.
	PhaseOneEncryption string `json:"phaseOneEncryption"`

	// PhaseOneKeylife - Amount of time (in seconds) allowed to pass before the encryption key expires. A
	// new key is generated without interrupting service. Valid times are from 120 to 172800 seconds.
	PhaseOneKeylife int `json:"phaseOneKeylife"`

	// PhaseTwoAuthentication - The authentication used in phase 2 proposal negotiation process.
	PhaseTwoAuthentication string `json:"phaseTwoAuthentication"`

	// PhaseTwoDiffieHellmanGroup - Determines the strength of the key used in the key exchange process.
	// The higher the group number the stronger the key is and the more secure it is. However, processing
	// time will increase as the strength of the key increases. Both peers must use the Diffie-Hellman
	// Group.
	PhaseTwoDiffieHellmanGroup int `json:"phaseTwoDiffieHellmanGroup"`

	// PhaseTwoEncryption - The encryption used in phase 2 proposal negotiation process.
	PhaseTwoEncryption string `json:"phaseTwoEncryption"`

	// PhaseTwoKeylife - Amount of time (in seconds) allowed to pass before the encryption key expires. A
	// new key is generated without interrupting service. Valid times are from 120 to 172800 seconds.
	PhaseTwoKeylife int `json:"phaseTwoKeylife"`

	// PhaseTwoPerfectForwardSecrecy - Determines if the generated keys are made from previous keys. When
	// PFS is specified, a Diffie-Hellman exchange occurs each time a new security association is
	// negotiated.
	PhaseTwoPerfectForwardSecrecy int `json:"phaseTwoPerfectForwardSecrecy"`

	// PresharedKey - A key used so that peers authenticate each other. This key is hashed by using the
	// phase one encryption and phase one authentication.
	PresharedKey string `json:"presharedKey"`

	// ServiceSubnetCount - A count of service subnets that can be access through the network tunnel.
	ServiceSubnetCount uint64 `json:"serviceSubnetCount"`

	// ServiceSubnets - Service subnets that can be access through the network tunnel.
	ServiceSubnets []*SoftLayer_Network_Subnet `json:"serviceSubnets"`

	// StaticRouteSubnetCount - A count of subnets used for a network tunnel's address translations.
	StaticRouteSubnetCount uint64 `json:"staticRouteSubnetCount"`

	// StaticRouteSubnets - Subnets used for a network tunnel's address translations.
	StaticRouteSubnets []*SoftLayer_Network_Subnet `json:"staticRouteSubnets"`

	// TransactionHistory - no documentation
	TransactionHistory []*SoftLayer_Provisioning_Version1_Transaction `json:"transactionHistory"`

	// TransactionHistoryCount - A count of the transaction history for this network tunnel.
	TransactionHistoryCount uint64 `json:"transactionHistoryCount"`
}

// AddCustomerSubnetToNetworkTunnel - Associates a remote subnet to the network tunnel. When a remote
// subnet is associated, a network tunnel will allow the customer (remote) network to communicate with
// the private and service subnets on the SoftLayer network which are on the other end of this network
// tunnel. A network tunnel's configurations must be applied to the network device in order for the
// association described above to take effect.
func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context) AddCustomerSubnetToNetworkTunnel(ctx *slapi.RequestContext, subnetId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AddPrivateSubnetToNetworkTunnel - Associates a private subnet to the network tunnel. When a private
// subnet is associated, the network tunnel will allow the customer (remote) network to access the
// private subnet. A network tunnel's configurations must be applied to the network device in order for
// the association described above to take effect.
func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context) AddPrivateSubnetToNetworkTunnel(ctx *slapi.RequestContext, subnetId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AddServiceSubnetToNetworkTunnel - Associates a service subnet to the network tunnel. When a service
// subnet is associated, a network tunnel will allow the customer (remote) network to communicate with
// the private and service subnets on the SoftLayer network which are on the other end of this network
// tunnel. Service subnets provide access to SoftLayer services such as the customer management portal
// and the SoftLayer A network tunnel's configurations must be applied to the network device in order
// for the association described above to take effect.
func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context) AddServiceSubnetToNetworkTunnel(ctx *slapi.RequestContext, subnetId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// ApplyConfigurationsToDevice - A transaction will be created to apply the IPSec network tunnel's
// configuration to SoftLayer network devices. During this time, an IPSec network tunnel cannot be
// modified in anyway. Only one network tunnel configuration transaction can be created. If a
// transaction has been created or is running, a new transaction cannot be created until the previous
// transaction completes.
func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context) ApplyConfigurationsToDevice(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// CreateAddressTranslation - Create an address translation for a network tunnel. To create an address
// translation, ip addresses from an assigned /30 static route subnet are used. Address translations
// deliver packets to a destination ip address that is on a customer (remote) subnet. A network
// tunnel's configurations must be applied to the network device in order for an address translation to
// be created.
func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context) CreateAddressTranslation(ctx *slapi.RequestContext, translation SoftLayer_Network_Tunnel_Module_Context_Address_Translation) (*SoftLayer_Network_Tunnel_Module_Context_Address_Translation, error) {
	var returnValue *SoftLayer_Network_Tunnel_Module_Context_Address_Translation
	return returnValue, nil
}

// CreateAddressTranslations - This has the same functionality as the
// SoftLayer_Network_Tunnel_Module_Context::createAddressTranslation. However, it allows multiple
// translations to be passed in for creation. A network tunnel's configurations must be applied to the
// network device in order for the address translations to be created.
func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context) CreateAddressTranslations(ctx *slapi.RequestContext, translations []SoftLayer_Network_Tunnel_Module_Context_Address_Translation) ([]*SoftLayer_Network_Tunnel_Module_Context_Address_Translation, error) {
	var returnValue []*SoftLayer_Network_Tunnel_Module_Context_Address_Translation
	return returnValue, nil
}

// DeleteAddressTranslation - Remove an existing address translation from a network tunnel. Address
// translations deliver packets to a destination ip address that is on a customer subnet (remote). A
// network tunnel's configurations must be applied to the network device in order for an address
// translation to be deleted.
func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context) DeleteAddressTranslation(ctx *slapi.RequestContext, translationId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DownloadAddressTranslationConfigurations - Provides all of the address translation configurations
// for an IPSec VPN tunnel in a text file
func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context) DownloadAddressTranslationConfigurations(ctx *slapi.RequestContext) (*SoftLayer_Container_Utility_File_Entity, error) {
	var returnValue *SoftLayer_Container_Utility_File_Entity
	return returnValue, nil
}

// DownloadParameterConfigurations - Provides all of the configurations for an IPSec VPN network tunnel
// in a text file
func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context) DownloadParameterConfigurations(ctx *slapi.RequestContext) (*SoftLayer_Container_Utility_File_Entity, error) {
	var returnValue *SoftLayer_Container_Utility_File_Entity
	return returnValue, nil
}

// EditAddressTranslation - Edit name, source (SoftLayer IP) ip address and/or destination (Customer
// IP) ip address for an existing address translation for a network tunnel. Address translations
// deliver packets to a destination ip address that is on a customer (remote) subnet. A network
// tunnel's configurations must be applied to the network device in order for an address translation to
// be created.
func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context) EditAddressTranslation(ctx *slapi.RequestContext, translation SoftLayer_Network_Tunnel_Module_Context_Address_Translation) (*SoftLayer_Network_Tunnel_Module_Context_Address_Translation, error) {
	var returnValue *SoftLayer_Network_Tunnel_Module_Context_Address_Translation
	return returnValue, nil
}

// EditAddressTranslations - Edit name, source (SoftLayer IP) ip address and/or destination (Customer
// IP) ip address for existing address translations for a network tunnel. Address translations deliver
// packets to a destination ip address that is on a customer (remote) subnet. A network tunnel's
// configurations must be applied to the network device in order for an address translation to be
// modified.
func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context) EditAddressTranslations(ctx *slapi.RequestContext, translations []SoftLayer_Network_Tunnel_Module_Context_Address_Translation) ([]*SoftLayer_Network_Tunnel_Module_Context_Address_Translation, error) {
	var returnValue []*SoftLayer_Network_Tunnel_Module_Context_Address_Translation
	return returnValue, nil
}

// EditObject - Negotiation parameters for both phases one and two are editable. Here are the phase one
// and two parameters that can modified: *Phase One **Authentication ***Default value is set to MD5.
// ***Valid Options are: MD5, SHA1, SHA256. **Encryption ***Default value is set to 3DES. ***Valid
// Options are: 3DES, AES128, AES192, AES256. **Diffie-Hellman Group ***Default value is set to 2.
// ***Valid Options are: 0 (None), 1, 2, 5. **Keylife ***Default value is set to 3600. ***Limits are:
// MIN = 120, MAX = 172800 **Preshared Key *Phase Two **Authentication ***Default value is set to MD5.
// ***Valid Options are: MD5, SHA1, SHA256. **Encryption ***Default value is set to 3DES. ***Valid
// Options are: 3DES, AES128, AES192, AES256. **Diffie-Hellman Group ***Default value is set to 2.
// ***Valid Options are: 0 (None), 1, 2, 5. **Keylife ***Default value is set to 28800. ***Limits are:
// MIN = 120, MAX = 172800 **Perfect Forward Secrecy ***Valid Options are: Off = 0, On = 1. If perfect
// forward secrecy is turned On (set to 1), then a phase 2 diffie-hellman group is required. The remote
// peer address for the network tunnel may also be modified if needed. Invalid options will not be
// accepted and will cause an exception to be thrown. There are properties that provide valid options
// and limits for each negotiation parameter. Those properties are as follows: * encryptionDefault *
// encryptionOptions * authenticationDefault * authenticationOptions * diffieHellmanGroupDefault *
// diffieHellmanGroupOptions * phaseOneKeylifeDefault * phaseTwoKeylifeDefault * keylifeLimits
// Configurations cannot be modified if a network tunnel's requires complex manual setups/configuration
// modifications by the SoftLayer Network department. If the former is required, the configurations for
// the network tunnel will be locked until the manual configurations are complete. A network tunnel's
// configurations are applied via a transaction. If a network tunnel configuration change transaction
// is currently running, the network tunnel's setting cannot be modified until the running transaction
// completes. A network tunnel's configurations must be applied to the network device in order for the
// modifications made to take effect.
func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context) EditObject(ctx *slapi.RequestContext, templateObject SoftLayer_Network_Tunnel_Module_Context) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetAddressTranslationConfigurations - The address translations will be returned. All the
// translations will be formatted so that the configurations can be copied into a host file. Format:
// {address translation SoftLayer IP Address} {address translation name}
func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context) GetAddressTranslationConfigurations(ctx *slapi.RequestContext) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetAuthenticationDefault - The default authentication type used for both phases of the negotiation
// process. The default value is set to MD5.
func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context) GetAuthenticationDefault(ctx *slapi.RequestContext) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetAuthenticationOptions - Authentication options available for both phases of the negotiation
// process. The authentication options are as follows: * MD5 * SHA1 * SHA256
func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context) GetAuthenticationOptions(ctx *slapi.RequestContext) ([]string, error) {
	var returnValue []string
	return returnValue, nil
}

// GetDiffieHellmanGroupDefault - The default Diffie-Hellman group used for both phases of the
// negotiation process. The default value is set to 2.
func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context) GetDiffieHellmanGroupDefault(ctx *slapi.RequestContext) (int, error) {
	var returnValue int
	return returnValue, nil
}

// GetDiffieHellmanGroupOptions - The Diffie-Hellman group options used for both phases of the
// negotiation process. The diffie-hellman group options are as follows: * 0 (None) * 1 * 2 * 5
func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context) GetDiffieHellmanGroupOptions(ctx *slapi.RequestContext) ([]int, error) {
	var returnValue []int
	return returnValue, nil
}

// GetEncryptionDefault - The default encryption type used for both phases of the negotiation process.
// The default value is set to 3DES.
func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context) GetEncryptionDefault(ctx *slapi.RequestContext) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetEncryptionOptions - Encryption options available for both phases of the negotiation process. The
// valid encryption options are as follows: * DES * 3DES * AES128 * AES192 * AES256
func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context) GetEncryptionOptions(ctx *slapi.RequestContext) ([]string, error) {
	var returnValue []string
	return returnValue, nil
}

// GetKeylifeLimits - The keylife limits. Keylife max limit is set to 120. Keylife min limit is set to
// 172800.
func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context) GetKeylifeLimits(ctx *slapi.RequestContext) ([]int, error) {
	var returnValue []int
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Network_Tunnel_Module_Context object whose ID number
// corresponds to the ID number of the init parameter passed to the
// SoftLayer_Network_Tunnel_Module_Context service. The IPSec network tunnel will be returned if it is
// associated with the account and the user has proper permission to manage network tunnels.
func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Tunnel_Module_Context, error) {
	var returnValue *SoftLayer_Network_Tunnel_Module_Context
	return returnValue, nil
}

// GetParameterConfigurationsForCustomerView - All of the IPSec VPN tunnel's configurations will be
// returned. It will list all of phase one and two negotiation parameters. Both remote and local
// subnets will be provided as well. This is useful when the configurations need to be passed on to
// another team and/or company for internal network configuration.
func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context) GetParameterConfigurationsForCustomerView(ctx *slapi.RequestContext) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetPhaseOneKeylifeDefault - The default phase 1 keylife used if a value is not provided. The default
// value is set to 3600.
func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context) GetPhaseOneKeylifeDefault(ctx *slapi.RequestContext) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetPhaseTwoKeylifeDefault - The default phase 2 keylife used if a value is not provided. The default
// value is set to 28800.
func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context) GetPhaseTwoKeylifeDefault(ctx *slapi.RequestContext) (string, error) {
	var returnValue string
	return returnValue, nil
}

// RemoveCustomerSubnetFromNetworkTunnel - Disassociate a customer subnet (remote) from a network
// tunnel. When a remote subnet is disassociated, that subnet will not able to communicate with private
// and service subnets on the SoftLayer network. A network tunnel's configurations must be applied to
// the network device in order for the disassociation described above to take effect.
func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context) RemoveCustomerSubnetFromNetworkTunnel(ctx *slapi.RequestContext, subnetId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemovePrivateSubnetFromNetworkTunnel - Disassociate a private subnet from a network tunnel. When a
// private subnet is disassociated, the customer (remote) subnet on the other end of the tunnel will
// not able to communicate with the private subnet that was just disassociated. A network tunnel's
// configurations must be applied to the network device in order for the disassociation described above
// to take effect.
func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context) RemovePrivateSubnetFromNetworkTunnel(ctx *slapi.RequestContext, subnetId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveServiceSubnetFromNetworkTunnel - Disassociate a service subnet from a network tunnel. When a
// service subnet is disassociated, that customer (remote) subnet on the other end of the network
// tunnel will not able to communicate with that service subnet on the SoftLayer network. A network
// tunnel's configurations must be applied to the network device in order for the disassociation
// described above to take effect.
func (softlayer_network_tunnel_module_context *SoftLayer_Network_Tunnel_Module_Context) RemoveServiceSubnetFromNetworkTunnel(ctx *slapi.RequestContext, subnetId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
