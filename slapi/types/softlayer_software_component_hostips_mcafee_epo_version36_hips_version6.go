package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version6 - The
// SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version6 data type represents a
// single McAfee Secure Host IPS software component for version 6 of the Host IPS client and uses the
// ePolicy Orchestrator version 3.6 backend.
type SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version6 struct {

	// BlockedApplicationEventCount - A count of the blocked application events for this software
	// component.
	BlockedApplicationEventCount uint64 `json:"blockedApplicationEventCount"`

	// BlockedApplicationEvents - The blocked application events for this software component.
	BlockedApplicationEvents []*McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_BlockedApplicationEvent `json:"blockedApplicationEvents"`

	// IpsEventCount - A count of the host IPS events for this software component.
	IpsEventCount uint64 `json:"ipsEventCount"`

	// IpsEvents - no documentation
	IpsEvents []*McAfee_Epolicy_Orchestrator_Version36_Hips_Version6_IPSEvent `json:"ipsEvents"`
}

func (softlayer_software_component_hostips_mcafee_epo_version36_hips_version6 *SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version6) String() string {
	return "SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version6"
}
