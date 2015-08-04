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

	// EnumerationValues - no documentation
	EnumerationValues string `json:"enumerationValues,omitempty"`

	// GroupId - no documentation
	GroupId string `json:"groupId,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// Sort - no documentation
	Sort int `json:"sort,omitempty"`

	// MinimumValue - no documentation
	MinimumValue string `json:"minimumValue,omitempty"`

	// RequireValueFlag - Indicates if a configuration value is required for this definition.
	RequireValueFlag int `json:"requireValueFlag,omitempty"`

	// TypeId - Internal identifier of a configuration definition type.
	TypeId int `json:"typeId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// MaximumValue - no documentation
	MaximumValue string `json:"maximumValue,omitempty"`

	// Path - no documentation
	Path string `json:"path,omitempty"`

	// SectionId - no documentation
	SectionId int `json:"sectionId,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// ShortName - no documentation
	ShortName string `json:"shortName,omitempty"`
}

func (softlayer_configuration_template_section_definition *SoftLayer_Configuration_Template_Section_Definition) String() string {
	return "SoftLayer_Configuration_Template_Section_Definition"
}

// SoftLayer_Configuration_Template_Section_Definition_Extended is SoftLayer_Configuration_Template_Section_Definition with all maskable types.
type SoftLayer_Configuration_Template_Section_Definition_Extended struct {
	SoftLayer_Configuration_Template_Section_Definition

	// Group - <nil>
	Group *SoftLayer_Configuration_Template_Section_Definition_Group `json:"group,omitempty"`

	// MonitoringDataFlag - <nil>
	MonitoringDataFlag bool `json:"monitoringDataFlag,omitempty"`

	// Section - <nil>
	Section *SoftLayer_Configuration_Template_Section `json:"section,omitempty"`

	// AttributeCount - no documentation
	AttributeCount uint64 `json:"attributeCount,omitempty"`

	// Attributes - <nil>
	Attributes []*SoftLayer_Configuration_Template_Section_Definition_Attribute `json:"attributes,omitempty"`

	// DefaultValue - <nil>
	DefaultValue *SoftLayer_Configuration_Template_Section_Definition_Value `json:"defaultValue,omitempty"`

	// ValueType - <nil>
	ValueType *SoftLayer_Configuration_Template_Section_Definition_Type `json:"valueType,omitempty"`
}

func (softlayer_configuration_template_section_definition *SoftLayer_Configuration_Template_Section_Definition_Extended) String() string {
	return "SoftLayer_Configuration_Template_Section_Definition"
}
