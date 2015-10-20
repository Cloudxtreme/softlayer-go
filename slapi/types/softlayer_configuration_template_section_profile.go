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

	// AgentId - Internal identifier of a monitoring agent this profile belongs to.
	AgentId int `json:"agentId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// SectionId - Internal identifier of a configuration section that this profile belongs to.
	SectionId int `json:"sectionId,omitempty"`

	// ConfigurationSection - <nil>
	ConfigurationSection *SoftLayer_Configuration_Template_Section `json:"configurationSection,omitempty"`

	// MonitoringAgent - <nil>
	MonitoringAgent *SoftLayer_Monitoring_Agent `json:"monitoringAgent,omitempty"`
}

func (softlayer_configuration_template_section_profile *SoftLayer_Configuration_Template_Section_Profile) String() string {
	return "SoftLayer_Configuration_Template_Section_Profile"
}
