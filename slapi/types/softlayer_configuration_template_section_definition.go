package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Configuration_Template_Section_Definition - Configuration definition gives you details of
// the value that you're setting. Some monitoring agents requires values unique to your system. If
// value type is defined as "Resource Specific Values", you will have to make an additional API call to
// retrieve your system specific values. See
// [[SoftLayer_Monitoring_Agent::getAvailableConfigurationValues|Monitoring Agent]] service to retrieve
// your system specific values.
type SoftLayer_Configuration_Template_Section_Definition struct {

	// AttributeCount - no documentation
	AttributeCount uint64 `json:"attributeCount"`

	// Attributes - <nil>
	Attributes []*SoftLayer_Configuration_Template_Section_Definition_Attribute `json:"attributes"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// DefaultValue - <nil>
	DefaultValue *SoftLayer_Configuration_Template_Section_Definition_Value `json:"defaultValue"`

	// Description - no documentation
	Description string `json:"description"`

	// EnumerationValues - no documentation
	EnumerationValues string `json:"enumerationValues"`

	// Group - <nil>
	Group *SoftLayer_Configuration_Template_Section_Definition_Group `json:"group"`

	// GroupId - no documentation
	GroupId string `json:"groupId"`

	// Id - no documentation
	Id int `json:"id"`

	// MaximumValue - no documentation
	MaximumValue string `json:"maximumValue"`

	// MinimumValue - no documentation
	MinimumValue string `json:"minimumValue"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// MonitoringDataFlag - <nil>
	MonitoringDataFlag bool `json:"monitoringDataFlag"`

	// Name - no documentation
	Name string `json:"name"`

	// Path - no documentation
	Path string `json:"path"`

	// RequireValueFlag - Indicates if a configuration value is required for this definition.
	RequireValueFlag int `json:"requireValueFlag"`

	// Section - <nil>
	Section *SoftLayer_Configuration_Template_Section `json:"section"`

	// SectionId - no documentation
	SectionId int `json:"sectionId"`

	// ShortName - no documentation
	ShortName string `json:"shortName"`

	// Sort - no documentation
	Sort int `json:"sort"`

	// TypeId - Internal identifier of a configuration definition type.
	TypeId int `json:"typeId"`

	// ValueType - <nil>
	ValueType *SoftLayer_Configuration_Template_Section_Definition_Type `json:"valueType"`
}

// GetObject - <nil>
func (softlayer_configuration_template_section_definition *SoftLayer_Configuration_Template_Section_Definition) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Configuration_Template_Section_Definition, error) {
	var returnValue *SoftLayer_Configuration_Template_Section_Definition
	return returnValue, nil
}
