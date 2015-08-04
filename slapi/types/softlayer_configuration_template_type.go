package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Configuration_Template_Type - The SoftLayer_Configuration_Template_Type data type contains
// configuration template type information.
type SoftLayer_Configuration_Template_Type struct {

	// Id - Internal identifier of a configuration template type
	Id int `json:"id,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`
}

func (softlayer_configuration_template_type *SoftLayer_Configuration_Template_Type) String() string {
	return "SoftLayer_Configuration_Template_Type"
}
