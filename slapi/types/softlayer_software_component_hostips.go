package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Software_Component_HostIps - This object specifies a specific type of Software Component:
// A Host Intrusion Protection System instance.
type SoftLayer_Software_Component_HostIps struct {
}

// GetCurrentHostIpsPolicies - no documentation
func (softlayer_software_component_hostips *SoftLayer_Software_Component_HostIps) GetCurrentHostIpsPolicies(commonOptions *slapi.CommonOptions) ([]*SoftLayer_Container_Software_Component_HostIps_Policy, error) {
	var returnValue []*SoftLayer_Container_Software_Component_HostIps_Policy
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_software_component_hostips *SoftLayer_Software_Component_HostIps) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Software_Component_HostIps, error) {
	var returnValue *SoftLayer_Software_Component_HostIps
	return returnValue, nil
}

// UpdateHipsPolicies - Update the Host IPS policies. To retrieve valid policy options you must use the
// provided relationships.
func (softlayer_software_component_hostips *SoftLayer_Software_Component_HostIps) UpdateHipsPolicies(commonOptions *slapi.CommonOptions, newIpsMode string, newIpsProtection string, newFirewallMode string, newFirewallRuleset string, newApplicationMode string, newApplicationRuleset string, newEnforcementPolicy string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
