package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference -
// SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference class holds the reference
// information, essentially a SQL join, between a monitoring configuration group and agent
// configuration templates.
type SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference struct {

	// ConfigurationTemplate - <nil>
	ConfigurationTemplate *SoftLayer_Configuration_Template `json:"configurationTemplate"`

	// ConfigurationTemplateId - no documentation
	ConfigurationTemplateId int `json:"configurationTemplateId"`

	// Id - Internal identifier of a configuration group reference record
	Id int `json:"id"`

	// TemplateGroup - <nil>
	TemplateGroup *SoftLayer_Monitoring_Agent_Configuration_Template_Group `json:"templateGroup"`

	// TemplateGroupId - Internal identifier of a monitoring agent configuration group
	TemplateGroupId int `json:"templateGroupId"`
}

// CreateObject - This method creates a monitoring agent configuration template group reference by
// passing in an object with the SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference
// structure as the $templateObject parameter.
func (softlayer_monitoring_agent_configuration_template_group_reference *SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference) CreateObject(ctx *slapi.RequestContext, templateObject SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference) (*SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference, error) {
	var returnValue *SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference
	return returnValue, nil
}

// CreateObjects - This method creates monitoring agent configuration template group references by
// passing in an array of objects with the
// SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference structure as the $templateObjects
// parameter. Setting the $bulkCommit parameter to true will commit the changes in one transaction,
// false will commit after each object is created.
func (softlayer_monitoring_agent_configuration_template_group_reference *SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference) CreateObjects(ctx *slapi.RequestContext, templateObjects []SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - This method updates a SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference
// record by passing in a modified instance of the object.
func (softlayer_monitoring_agent_configuration_template_group_reference *SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference) EditObject(ctx *slapi.RequestContext, templateObject SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObjects - This method updates a set of
// SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference records by passing in an array of
// modified instances of the objects. Setting the $bulkCommit parameter to true will commit the changes
// in one transaction, false will commit after each object is updated.
func (softlayer_monitoring_agent_configuration_template_group_reference *SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference) EditObjects(ctx *slapi.RequestContext, templateObjects []SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetAllObjects - This method retrieves all
// SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference objects accessible to the active
// user.
func (softlayer_monitoring_agent_configuration_template_group_reference *SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference, error) {
	var returnValue []*SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference
	return returnValue, nil
}

// GetObject - This method retrieves a monitoring agent configuration template group reference whose
// identifier corresponds to the value provided in the initialization parameter passed to the
// SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference service.
func (softlayer_monitoring_agent_configuration_template_group_reference *SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference, error) {
	var returnValue *SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference
	return returnValue, nil
}
