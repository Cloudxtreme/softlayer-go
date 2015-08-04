package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference -
// SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference class holds the reference
// information, essentially a SQL join, between a monitoring configuration group and agent
// configuration templates.
type SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference struct {

	// ConfigurationTemplateId - no documentation
	ConfigurationTemplateId int `json:"configurationTemplateId,omitempty"`

	// Id - Internal identifier of a configuration group reference record
	Id int `json:"id,omitempty"`

	// TemplateGroupId - Internal identifier of a monitoring agent configuration group
	TemplateGroupId int `json:"templateGroupId,omitempty"`
}

func (softlayer_monitoring_agent_configuration_template_group_reference *SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference) String() string {
	return "SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference"
}

// SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference_Extended is SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference with all maskable types.
type SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference_Extended struct {
	SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference

	// TemplateGroup - <nil>
	TemplateGroup *SoftLayer_Monitoring_Agent_Configuration_Template_Group `json:"templateGroup,omitempty"`

	// ConfigurationTemplate - <nil>
	ConfigurationTemplate *SoftLayer_Configuration_Template `json:"configurationTemplate,omitempty"`
}

func (softlayer_monitoring_agent_configuration_template_group_reference *SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference_Extended) String() string {
	return "SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference"
}
