package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Virtual_Guest_Boot_Parameter_Type - Describes a virtual guest boot parameter. In this the
// word class is used in the context of arguments sent to cloud computing instances such as single user
// mode and boot into bash.
type SoftLayer_Virtual_Guest_Boot_Parameter_Type struct {

	// Id - no documentation
	Id int `json:"id"`

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`

	// Value - no documentation
	Value string `json:"value"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Description - no documentation
	Description string `json:"description"`
}

func (softlayer_virtual_guest_boot_parameter_type *SoftLayer_Virtual_Guest_Boot_Parameter_Type) String() string {
	return "SoftLayer_Virtual_Guest_Boot_Parameter_Type"
}
