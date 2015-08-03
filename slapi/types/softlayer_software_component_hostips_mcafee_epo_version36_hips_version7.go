package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version7 - The
// SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version7 data type represents a
// single McAfee Secure Host IPS software component for version 7 of the Host IPS client and uses the
// ePolicy Orchestrator version 3.6 backend.
type SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version7 struct {
}

func (softlayer_software_component_hostips_mcafee_epo_version36_hips_version7 *SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version7) String() string {
	return "SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version7"
}

// SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version7_Extended is SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version7 with all maskable types.
type SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version7_Extended struct {
	SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version7

	// BlockedApplicationEvents - The blocked application events for this software component.
	BlockedApplicationEvents []*McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_BlockedApplicationEvent `json:"blockedApplicationEvents"`

	// IpsEvents - no documentation
	IpsEvents []*McAfee_Epolicy_Orchestrator_Version36_Hips_Version7_IPSEvent `json:"ipsEvents"`

	// BlockedApplicationEventCount - A count of the blocked application events for this software
	// component.
	BlockedApplicationEventCount uint64 `json:"blockedApplicationEventCount"`

	// IpsEventCount - A count of the host IPS events for this software component.
	IpsEventCount uint64 `json:"ipsEventCount"`
}

func (softlayer_software_component_hostips_mcafee_epo_version36_hips_version7 *SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version7_Extended) String() string {
	return "SoftLayer_Software_Component_HostIps_Mcafee_Epo_Version36_Hips_Version7"
}
