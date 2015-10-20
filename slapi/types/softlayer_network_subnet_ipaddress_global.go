package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Subnet_IpAddress_Global - <nil>
type SoftLayer_Network_Subnet_IpAddress_Global struct {

	// IpAddressId - A Global IP Address' associated [[SoftLayer_Account|account]] ID
	IpAddressId int `json:"ipAddressId,omitempty"`

	// Description - no documentation
	Description int `json:"description,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// TypeId - A Global IP Address' associated type [[SoftLayer_Network_Subnet_IpAddress_Global_Type|id]]
	// ID
	TypeId int `json:"typeId,omitempty"`

	// DestinationIpAddressId - A Global IP Address' associated
	// [[SoftLayer_Network_Subnet_IpAddress|ipAddress]] ID
	DestinationIpAddressId int `json:"destinationIpAddressId,omitempty"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item_Network_Subnet_IpAddress_Global `json:"billingItem,omitempty"`

	// IpAddress - <nil>
	IpAddress *SoftLayer_Network_Subnet_IpAddress `json:"ipAddress,omitempty"`

	// ActiveTransaction - The active transaction associated with this Global
	ActiveTransaction *SoftLayer_Provisioning_Version1_Transaction `json:"activeTransaction,omitempty"`

	// DestinationIpAddress - <nil>
	DestinationIpAddress *SoftLayer_Network_Subnet_IpAddress `json:"destinationIpAddress,omitempty"`
}

func (softlayer_network_subnet_ipaddress_global *SoftLayer_Network_Subnet_IpAddress_Global) String() string {
	return "SoftLayer_Network_Subnet_IpAddress_Global"
}
