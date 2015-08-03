package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Hardware_Note - <nil>
type SoftLayer_Hardware_Note struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// HardwareId - <nil>
	HardwareId int `json:"hardwareId"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// TypeId - <nil>
	TypeId int `json:"typeId"`

	// UserRecordId - <nil>
	UserRecordId int `json:"userRecordId"`

	// Id - <nil>
	Id int `json:"id"`

	// Note - <nil>
	Note string `json:"note"`
}

func (softlayer_hardware_note *SoftLayer_Hardware_Note) String() string {
	return "SoftLayer_Hardware_Note"
}

// SoftLayer_Hardware_Note_Extended is SoftLayer_Hardware_Note with all maskable types.
type SoftLayer_Hardware_Note_Extended struct {
	SoftLayer_Hardware_Note

	// Hardware - <nil>
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// Employee - <nil>
	Employee *SoftLayer_User_Employee `json:"employee"`

	// Type - <nil>
	Type *SoftLayer_Hardware_Note_Type `json:"type"`

	// User - <nil>
	User *SoftLayer_User_Customer `json:"user"`
}

func (softlayer_hardware_note *SoftLayer_Hardware_Note_Extended) String() string {
	return "SoftLayer_Hardware_Note"
}
