package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Software_Component_HostIps_Policy - The
// SoftLayer_Container_Software_Component_HostIps_Policy container holds the title and value of a
// current host ips policy.
type SoftLayer_Container_Software_Component_HostIps_Policy struct {

	// PolicyTitle - no documentation
	PolicyTitle string `json:"policyTitle"`

	// Policy - no documentation
	Policy string `json:"policy"`
}

func (softlayer_container_software_component_hostips_policy *SoftLayer_Container_Software_Component_HostIps_Policy) String() string {
	return "SoftLayer_Container_Software_Component_HostIps_Policy"
}
