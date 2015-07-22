package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_Subnet_IpAddress_Global - <nil>
type SoftLayer_Network_Subnet_IpAddress_Global struct {

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// ActiveTransaction - The active transaction associated with this Global
	ActiveTransaction *SoftLayer_Provisioning_Version1_Transaction `json:"activeTransaction"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item_Network_Subnet_IpAddress_Global `json:"billingItem"`

	// Description - no documentation
	Description int `json:"description"`

	// DestinationIpAddress - <nil>
	DestinationIpAddress *SoftLayer_Network_Subnet_IpAddress `json:"destinationIpAddress"`

	// DestinationIpAddressId - A Global IP Address' associated
	// [[SoftLayer_Network_Subnet_IpAddress|ipAddress]] ID
	DestinationIpAddressId int `json:"destinationIpAddressId"`

	// Id - no documentation
	Id int `json:"id"`

	// IpAddress - <nil>
	IpAddress *SoftLayer_Network_Subnet_IpAddress `json:"ipAddress"`

	// IpAddressId - A Global IP Address' associated [[SoftLayer_Account|account]] ID
	IpAddressId int `json:"ipAddressId"`

	// TypeId - A Global IP Address' associated type [[SoftLayer_Network_Subnet_IpAddress_Global_Type|id]]
	// ID
	TypeId int `json:"typeId"`
}

// GetObject - <nil>
func (softlayer_network_subnet_ipaddress_global *SoftLayer_Network_Subnet_IpAddress_Global) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Network_Subnet_IpAddress_Global, error) {
	var returnValue *SoftLayer_Network_Subnet_IpAddress_Global
	return returnValue, nil
}

// Route - This function is used to create a new transaction to modify a global IP route. Routes are
// updated in one to two minutes depending on the number of transactions that are pending for a router.
func (softlayer_network_subnet_ipaddress_global *SoftLayer_Network_Subnet_IpAddress_Global) Route(commonOptions *slapi.CommonOptions, newEndPointIpAddress string) (*SoftLayer_Provisioning_Version1_Transaction, error) {
	var returnValue *SoftLayer_Provisioning_Version1_Transaction
	return returnValue, nil
}

// Unroute - This function is used to create a new transaction to unroute a global IP address. Routes
// are updated in one to two minutes depending on the number of transactions that are pending for a
// router.
func (softlayer_network_subnet_ipaddress_global *SoftLayer_Network_Subnet_IpAddress_Global) Unroute(commonOptions *slapi.CommonOptions) (*SoftLayer_Provisioning_Version1_Transaction, error) {
	var returnValue *SoftLayer_Provisioning_Version1_Transaction
	return returnValue, nil
}
