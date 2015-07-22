package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Configuration_Template_Section_Profile - Some configuration templates let you create a
// unique configuration profiles. For example, you can create multiple configuration profiles to
// monitor multiple hard drives with "CPU/Memory/Disk Monitoring Agent".
// SoftLayer_Configuration_Template_Section_Profile help you keep track of custom configuration
// profiles.
type SoftLayer_Configuration_Template_Section_Profile struct {

	// AgentId - Internal identifier of a monitoring agent this profile belongs to.
	AgentId int `json:"agentId"`

	// ConfigurationSection - <nil>
	ConfigurationSection *SoftLayer_Configuration_Template_Section `json:"configurationSection"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Id - no documentation
	Id int `json:"id"`

	// MonitoringAgent - <nil>
	MonitoringAgent *SoftLayer_Monitoring_Agent `json:"monitoringAgent"`

	// Name - no documentation
	Name string `json:"name"`

	// SectionId - Internal identifier of a configuration section that this profile belongs to.
	SectionId int `json:"sectionId"`
}

// GetObject - <nil>
func (softlayer_configuration_template_section_profile *SoftLayer_Configuration_Template_Section_Profile) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Configuration_Template_Section_Profile, error) {
	var returnValue *SoftLayer_Configuration_Template_Section_Profile
	return returnValue, nil
}
