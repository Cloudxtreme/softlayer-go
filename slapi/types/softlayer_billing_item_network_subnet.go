package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Network_Subnet - The SoftLayer_Billing_Item_Network_Subnet data type contains
// general information relating to a single SoftLayer billing item whose item category code is one of
// the following: * pri_ip_address * static_sec_ip_addresses (static secondary) * sov_sec_ip_addresses
// (secondary on vlan, also known as "portable ips") * sov_sec_ip_addresses_pub (sov_sec_ip_addresses
// public only) * sov_sec_ip_addresses_priv (sov_sec_ip_addresses private only) * sec_ip_addresses (old
// style, secondary ip addresses) These item categories denote that the billing item has subnet
// information attached.
type SoftLayer_Billing_Item_Network_Subnet struct {

	// Resource - no documentation
	Resource *SoftLayer_Network_Subnet `json:"resource"`

	// ResourceName - no documentation
	ResourceName string `json:"resourceName"`

	// ResourceTableId - The resource (unique identifier) for a server billing item.
	ResourceTableId int `json:"resourceTableId"`
}

func (softlayer_billing_item_network_subnet *SoftLayer_Billing_Item_Network_Subnet) String() string {
	return "SoftLayer_Billing_Item_Network_Subnet"
}
