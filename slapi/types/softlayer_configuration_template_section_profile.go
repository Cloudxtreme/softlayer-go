package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Configuration_Template_Section_Profile - Some configuration templates let you create a
// unique configuration profiles. For example, you can create multiple configuration profiles to
// monitor multiple hard drives with "CPU/Memory/Disk Monitoring Agent".
// SoftLayer_Configuration_Template_Section_Profile help you keep track of custom configuration
// profiles.
type SoftLayer_Configuration_Template_Section_Profile struct {

	// Id - no documentation
	Id int `json:"id"`

	// Name - no documentation
	Name string `json:"name"`

	// SectionId - Internal identifier of a configuration section that this profile belongs to.
	SectionId int `json:"sectionId"`

	// AgentId - Internal identifier of a monitoring agent this profile belongs to.
	AgentId int `json:"agentId"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`
}

func (softlayer_configuration_template_section_profile *SoftLayer_Configuration_Template_Section_Profile) String() string {
	return "SoftLayer_Configuration_Template_Section_Profile"
}

// SoftLayer_Configuration_Template_Section_Profile_Extended is SoftLayer_Configuration_Template_Section_Profile with all maskable types.
type SoftLayer_Configuration_Template_Section_Profile_Extended struct {
	SoftLayer_Configuration_Template_Section_Profile

	// ConfigurationSection - <nil>
	ConfigurationSection *SoftLayer_Configuration_Template_Section `json:"configurationSection"`

	// MonitoringAgent - <nil>
	MonitoringAgent *SoftLayer_Monitoring_Agent `json:"monitoringAgent"`
}

func (softlayer_configuration_template_section_profile *SoftLayer_Configuration_Template_Section_Profile_Extended) String() string {
	return "SoftLayer_Configuration_Template_Section_Profile"
}
