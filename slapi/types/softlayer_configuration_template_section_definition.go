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

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// MaximumValue - no documentation
	MaximumValue string `json:"maximumValue,omitempty"`

	// MinimumValue - no documentation
	MinimumValue string `json:"minimumValue,omitempty"`

	// Sort - no documentation
	Sort int `json:"sort,omitempty"`

	// TypeId - Internal identifier of a configuration definition type.
	TypeId int `json:"typeId,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// GroupId - no documentation
	GroupId string `json:"groupId,omitempty"`

	// RequireValueFlag - Indicates if a configuration value is required for this definition.
	RequireValueFlag int `json:"requireValueFlag,omitempty"`

	// EnumerationValues - no documentation
	EnumerationValues string `json:"enumerationValues,omitempty"`

	// Path - no documentation
	Path string `json:"path,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// SectionId - no documentation
	SectionId int `json:"sectionId,omitempty"`

	// ShortName - no documentation
	ShortName string `json:"shortName,omitempty"`
}

func (softlayer_configuration_template_section_definition *SoftLayer_Configuration_Template_Section_Definition) String() string {
	return "SoftLayer_Configuration_Template_Section_Definition"
}

// SoftLayer_Configuration_Template_Section_Definition_Extended is SoftLayer_Configuration_Template_Section_Definition with all maskable types.
type SoftLayer_Configuration_Template_Section_Definition_Extended struct {
	SoftLayer_Configuration_Template_Section_Definition

	// DefaultValue - <nil>
	DefaultValue *SoftLayer_Configuration_Template_Section_Definition_Value `json:"defaultValue,omitempty"`

	// Group - <nil>
	Group *SoftLayer_Configuration_Template_Section_Definition_Group `json:"group,omitempty"`

	// MonitoringDataFlag - <nil>
	MonitoringDataFlag bool `json:"monitoringDataFlag,omitempty"`

	// AttributeCount - no documentation
	AttributeCount uint64 `json:"attributeCount,omitempty"`

	// ValueType - <nil>
	ValueType *SoftLayer_Configuration_Template_Section_Definition_Type `json:"valueType,omitempty"`

	// Attributes - <nil>
	Attributes []*SoftLayer_Configuration_Template_Section_Definition_Attribute `json:"attributes,omitempty"`

	// Section - <nil>
	Section *SoftLayer_Configuration_Template_Section `json:"section,omitempty"`
}

func (softlayer_configuration_template_section_definition *SoftLayer_Configuration_Template_Section_Definition_Extended) String() string {
	return "SoftLayer_Configuration_Template_Section_Definition"
}
