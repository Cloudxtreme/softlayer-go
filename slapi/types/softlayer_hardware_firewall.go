package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Firewall - The SoftLayer_Hardware_Firewall data type contains general information
// relating to a single SoftLayer firewall.
type SoftLayer_Hardware_Firewall struct {
}

// SoftLayer_Hardware_Firewall_Extended is SoftLayer_Hardware_Firewall with all maskable types.
type SoftLayer_Hardware_Firewall_Extended struct {
	SoftLayer_Hardware_Firewall

	// UserCount - A count of a list of users that have access to this hardware firewall.
	UserCount uint64 `json:"userCount"`

	// Users - A list of users that have access to this hardware firewall.
	Users []*SoftLayer_User_Customer `json:"users"`
}

func (softlayer_hardware_firewall *SoftLayer_Hardware_Firewall) String() string {
	return "SoftLayer_Hardware_Firewall"
}
