package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Software_Component_AntivirusSpyware - This object specifies a specific type of Software
// Component: An Anti-virus/spyware instance. Anti-virus/spyware installations have specific properties
// and methods such as SoftLayer_Software_Component_AntivirusSpyware::updateAntivirusSpywarePolicy.
// Defaults are initiated by this object.
type SoftLayer_Software_Component_AntivirusSpyware struct {
}

// GetObject - <nil>
func (softlayer_software_component_antivirusspyware *SoftLayer_Software_Component_AntivirusSpyware) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Software_Component_AntivirusSpyware, error) {
	var returnValue *SoftLayer_Software_Component_AntivirusSpyware
	return returnValue, nil
}

// UpdateAntivirusSpywarePolicy - Update an anti-virus/spyware policy. The policy options that it
// accepts are the following: *1 - Minimal *2 - Relaxed *3 - Default *4 - High *5 - Ultimate
func (softlayer_software_component_antivirusspyware *SoftLayer_Software_Component_AntivirusSpyware) UpdateAntivirusSpywarePolicy(commonOptions *slapi.CommonOptions, newPolicy string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}