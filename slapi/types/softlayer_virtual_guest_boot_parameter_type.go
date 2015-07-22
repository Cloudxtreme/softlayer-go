package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Virtual_Guest_Boot_Parameter_Type - Describes a virtual guest boot parameter. In this the
// word class is used in the context of arguments sent to cloud computing instances such as single user
// mode and boot into bash.
type SoftLayer_Virtual_Guest_Boot_Parameter_Type struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Description - no documentation
	Description string `json:"description"`

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
}

// GetAllObjects - <nil>
func (softlayer_virtual_guest_boot_parameter_type *SoftLayer_Virtual_Guest_Boot_Parameter_Type) GetAllObjects(commonOptions *slapi.CommonOptions) ([]*SoftLayer_Virtual_Guest_Boot_Parameter_Type, error) {
	var returnValue []*SoftLayer_Virtual_Guest_Boot_Parameter_Type
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_virtual_guest_boot_parameter_type *SoftLayer_Virtual_Guest_Boot_Parameter_Type) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Virtual_Guest_Boot_Parameter_Type, error) {
	var returnValue *SoftLayer_Virtual_Guest_Boot_Parameter_Type
	return returnValue, nil
}
