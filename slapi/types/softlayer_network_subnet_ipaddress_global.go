package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Subnet_IpAddress_Global - <nil>
type SoftLayer_Network_Subnet_IpAddress_Global struct {

	// Description - no documentation
	Description int `json:"description"`

	// DestinationIpAddressId - A Global IP Address' associated
	// [[SoftLayer_Network_Subnet_IpAddress|ipAddress]] ID
	DestinationIpAddressId int `json:"destinationIpAddressId"`

	// IpAddressId - A Global IP Address' associated [[SoftLayer_Account|account]] ID
	IpAddressId int `json:"ipAddressId"`

	// Id - no documentation
	Id int `json:"id"`

	// TypeId - A Global IP Address' associated type [[SoftLayer_Network_Subnet_IpAddress_Global_Type|id]]
	// ID
	TypeId int `json:"typeId"`
}

// SoftLayer_Network_Subnet_IpAddress_Global_Extended is SoftLayer_Network_Subnet_IpAddress_Global with all maskable types.
type SoftLayer_Network_Subnet_IpAddress_Global_Extended struct {
	SoftLayer_Network_Subnet_IpAddress_Global

	// DestinationIpAddress - <nil>
	DestinationIpAddress *SoftLayer_Network_Subnet_IpAddress `json:"destinationIpAddress"`

	// IpAddress - <nil>
	IpAddress *SoftLayer_Network_Subnet_IpAddress `json:"ipAddress"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// ActiveTransaction - The active transaction associated with this Global
	ActiveTransaction *SoftLayer_Provisioning_Version1_Transaction `json:"activeTransaction"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item_Network_Subnet_IpAddress_Global `json:"billingItem"`
}

func (softlayer_network_subnet_ipaddress_global *SoftLayer_Network_Subnet_IpAddress_Global) String() string {
	return "SoftLayer_Network_Subnet_IpAddress_Global"
}
