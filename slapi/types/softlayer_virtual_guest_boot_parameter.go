package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Virtual_Guest_Boot_Parameter - <nil>
type SoftLayer_Virtual_Guest_Boot_Parameter struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// GuestBootParameterTypeId - <nil>
	GuestBootParameterTypeId int `json:"guestBootParameterTypeId"`

	// GuestId - <nil>
	GuestId int `json:"guestId"`

	// Id - <nil>
	Id int `json:"id"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`
}

func (softlayer_virtual_guest_boot_parameter *SoftLayer_Virtual_Guest_Boot_Parameter) String() string {
	return "SoftLayer_Virtual_Guest_Boot_Parameter"
}

// SoftLayer_Virtual_Guest_Boot_Parameter_Extended is SoftLayer_Virtual_Guest_Boot_Parameter with all maskable types.
type SoftLayer_Virtual_Guest_Boot_Parameter_Extended struct {
	SoftLayer_Virtual_Guest_Boot_Parameter

	// Guest - <nil>
	Guest *SoftLayer_Virtual_Guest `json:"guest"`

	// GuestBootParameterType - <nil>
	GuestBootParameterType *SoftLayer_Virtual_Guest_Boot_Parameter_Type `json:"guestBootParameterType"`
}

func (softlayer_virtual_guest_boot_parameter *SoftLayer_Virtual_Guest_Boot_Parameter_Extended) String() string {
	return "SoftLayer_Virtual_Guest_Boot_Parameter"
}
