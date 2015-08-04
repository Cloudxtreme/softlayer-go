package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Virtual_Guest_Boot_Parameter - <nil>
type SoftLayer_Virtual_Guest_Boot_Parameter struct {

	// GuestBootParameterTypeId - <nil>
	GuestBootParameterTypeId int `json:"guestBootParameterTypeId,omitempty"`

	// GuestId - <nil>
	GuestId int `json:"guestId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Guest - <nil>
	Guest *SoftLayer_Virtual_Guest `json:"guest,omitempty"`

	// GuestBootParameterType - <nil>
	GuestBootParameterType *SoftLayer_Virtual_Guest_Boot_Parameter_Type `json:"guestBootParameterType,omitempty"`
}

func (softlayer_virtual_guest_boot_parameter *SoftLayer_Virtual_Guest_Boot_Parameter) String() string {
	return "SoftLayer_Virtual_Guest_Boot_Parameter"
}
