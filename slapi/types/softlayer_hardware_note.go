package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Hardware_Note - <nil>
type SoftLayer_Hardware_Note struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Note - <nil>
	Note string `json:"note,omitempty"`

	// TypeId - <nil>
	TypeId int `json:"typeId,omitempty"`

	// HardwareId - <nil>
	HardwareId int `json:"hardwareId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// UserRecordId - <nil>
	UserRecordId int `json:"userRecordId,omitempty"`
}

func (softlayer_hardware_note *SoftLayer_Hardware_Note) String() string {
	return "SoftLayer_Hardware_Note"
}

// SoftLayer_Hardware_Note_Extended is SoftLayer_Hardware_Note with all maskable types.
type SoftLayer_Hardware_Note_Extended struct {
	SoftLayer_Hardware_Note

	// Type - <nil>
	Type *SoftLayer_Hardware_Note_Type `json:"type,omitempty"`

	// User - <nil>
	User *SoftLayer_User_Customer `json:"user,omitempty"`

	// Employee - <nil>
	Employee *SoftLayer_User_Employee `json:"employee,omitempty"`

	// Hardware - <nil>
	Hardware *SoftLayer_Hardware `json:"hardware,omitempty"`
}

func (softlayer_hardware_note *SoftLayer_Hardware_Note_Extended) String() string {
	return "SoftLayer_Hardware_Note"
}
