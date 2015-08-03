package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Dns_Domain_ResourceRecord_SrvType - SoftLayer_Dns_Domain_ResourceRecord_SrvType is a
// SoftLayer_Dns_Domain_ResourceRecord object whose ''type'' property is set to "srv" and defines a DNS
// SRV record on a SoftLayer hosted domain.
type SoftLayer_Dns_Domain_ResourceRecord_SrvType struct {

	// Port - The TCP or UDP port on which the service is to be found.
	Port int `json:"port"`

	// Priority - The priority of the target host, lower value means more preferred.
	Priority int `json:"priority"`

	// Protocol - The protocol of the desired service; this is usually either TCP or
	Protocol string `json:"protocol"`

	// Service - no documentation
	Service string `json:"service"`

	// Weight - A relative weight for records with the same priority.
	Weight int `json:"weight"`
}

func (softlayer_dns_domain_resourcerecord_srvtype *SoftLayer_Dns_Domain_ResourceRecord_SrvType) String() string {
	return "SoftLayer_Dns_Domain_ResourceRecord_SrvType"
}
