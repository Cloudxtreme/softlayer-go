package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Virtual_Guest_Boot_Parameter - <nil>
type SoftLayer_Virtual_Guest_Boot_Parameter struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Guest - <nil>
	Guest *SoftLayer_Virtual_Guest `json:"guest"`

	// GuestBootParameterType - <nil>
	GuestBootParameterType *SoftLayer_Virtual_Guest_Boot_Parameter_Type `json:"guestBootParameterType"`

	// GuestBootParameterTypeId - <nil>
	GuestBootParameterTypeId int `json:"guestBootParameterTypeId"`

	// GuestId - <nil>
	GuestId int `json:"guestId"`

	// Id - <nil>
	Id int `json:"id"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`
}

// CreateObject - <nil>
func (softlayer_virtual_guest_boot_parameter *SoftLayer_Virtual_Guest_Boot_Parameter) CreateObject(templateObject SoftLayer_Virtual_Guest_Boot_Parameter) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeleteObject - <nil>
func (softlayer_virtual_guest_boot_parameter *SoftLayer_Virtual_Guest_Boot_Parameter) DeleteObject() (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - <nil>
func (softlayer_virtual_guest_boot_parameter *SoftLayer_Virtual_Guest_Boot_Parameter) EditObject(templateObject SoftLayer_Virtual_Guest_Boot_Parameter) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_virtual_guest_boot_parameter *SoftLayer_Virtual_Guest_Boot_Parameter) GetObject() (*SoftLayer_Virtual_Guest_Boot_Parameter, error) {
	var returnValue *SoftLayer_Virtual_Guest_Boot_Parameter
	return returnValue, nil
}
