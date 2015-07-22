package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Hardware_Note - <nil>
type SoftLayer_Hardware_Note struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Employee - <nil>
	Employee *SoftLayer_User_Employee `json:"employee"`

	// Hardware - <nil>
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// HardwareId - <nil>
	HardwareId int `json:"hardwareId"`

	// Id - <nil>
	Id int `json:"id"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// Note - <nil>
	Note string `json:"note"`

	// Type - <nil>
	Type *SoftLayer_Hardware_Note_Type `json:"type"`

	// TypeId - <nil>
	TypeId int `json:"typeId"`

	// User - <nil>
	User *SoftLayer_User_Customer `json:"user"`

	// UserRecordId - <nil>
	UserRecordId int `json:"userRecordId"`
}
