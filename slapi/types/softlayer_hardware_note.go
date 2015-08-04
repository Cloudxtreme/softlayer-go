package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Hardware_Note - <nil>
type SoftLayer_Hardware_Note struct {

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Note - <nil>
	Note string `json:"note,omitempty"`

	// UserRecordId - <nil>
	UserRecordId int `json:"userRecordId,omitempty"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// HardwareId - <nil>
	HardwareId int `json:"hardwareId,omitempty"`

	// TypeId - <nil>
	TypeId int `json:"typeId,omitempty"`

	// Type - <nil>
	Type *SoftLayer_Hardware_Note_Type `json:"type,omitempty"`

	// User - <nil>
	User *SoftLayer_User_Customer `json:"user,omitempty"`

	// Employee - <nil>
	Employee *SoftLayer_User_Employee `json:"employee,omitempty"`

	// Hardware - <nil>
	Hardware *SoftLayer_Hardware `json:"hardware,omitempty"`
}

func (softlayer_hardware_note *SoftLayer_Hardware_Note) String() string {
	return "SoftLayer_Hardware_Note"
}
