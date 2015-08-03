package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Configuration_Template_Section_Definition - Configuration definition gives you details of
// the value that you're setting. Some monitoring agents requires values unique to your system. If
// value type is defined as "Resource Specific Values", you will have to make an additional API call to
// retrieve your system specific values. See
// [[SoftLayer_Monitoring_Agent::getAvailableConfigurationValues|Monitoring Agent]] service to retrieve
// your system specific values.
type SoftLayer_Configuration_Template_Section_Definition struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// MinimumValue - no documentation
	MinimumValue string `json:"minimumValue"`

	// Sort - no documentation
	Sort int `json:"sort"`

	// Description - no documentation
	Description string `json:"description"`

	// GroupId - no documentation
	GroupId string `json:"groupId"`

	// MaximumValue - no documentation
	MaximumValue string `json:"maximumValue"`

	// Path - no documentation
	Path string `json:"path"`

	// RequireValueFlag - Indicates if a configuration value is required for this definition.
	RequireValueFlag int `json:"requireValueFlag"`

	// Id - no documentation
	Id int `json:"id"`

	// TypeId - Internal identifier of a configuration definition type.
	TypeId int `json:"typeId"`

	// EnumerationValues - no documentation
	EnumerationValues string `json:"enumerationValues"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`

	// SectionId - no documentation
	SectionId int `json:"sectionId"`

	// ShortName - no documentation
	ShortName string `json:"shortName"`
}

// SoftLayer_Configuration_Template_Section_Definition_Extended is SoftLayer_Configuration_Template_Section_Definition with all maskable types.
type SoftLayer_Configuration_Template_Section_Definition_Extended struct {
	SoftLayer_Configuration_Template_Section_Definition

	// Attributes - <nil>
	Attributes []*SoftLayer_Configuration_Template_Section_Definition_Attribute `json:"attributes"`

	// MonitoringDataFlag - <nil>
	MonitoringDataFlag bool `json:"monitoringDataFlag"`

	// AttributeCount - no documentation
	AttributeCount uint64 `json:"attributeCount"`

	// DefaultValue - <nil>
	DefaultValue *SoftLayer_Configuration_Template_Section_Definition_Value `json:"defaultValue"`

	// Group - <nil>
	Group *SoftLayer_Configuration_Template_Section_Definition_Group `json:"group"`

	// Section - <nil>
	Section *SoftLayer_Configuration_Template_Section `json:"section"`

	// ValueType - <nil>
	ValueType *SoftLayer_Configuration_Template_Section_Definition_Type `json:"valueType"`
}

func (softlayer_configuration_template_section_definition *SoftLayer_Configuration_Template_Section_Definition) String() string {
	return "SoftLayer_Configuration_Template_Section_Definition"
}
