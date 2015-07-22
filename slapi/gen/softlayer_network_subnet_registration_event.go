package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Subnet_Registration_Event - Each time a
// [[SoftLayer_Network_Subnet_Registration|subnet registration]] object is created or modified, the
// system will generate an event for it. Additional actions that would create an event include RIR
// responses and error cases. *
type SoftLayer_Network_Subnet_Registration_Event struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Id - no documentation
	Id int `json:"id"`

	// Message - A string message indicating what took place during this event
	Message string `json:"message"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// Registration - no documentation
	Registration *SoftLayer_Network_Subnet_Registration `json:"registration"`

	// RegistrationId - The numeric ID of the related [[SoftLayer_Network_Subnet_Registration]] object
	RegistrationId int `json:"registrationId"`

	// Type - no documentation
	Type *SoftLayer_Network_Subnet_Registration_Event_Type `json:"type"`

	// TypeId - The numeric ID of the associated [[SoftLayer_Network_Subnet_Registration_Event_Type|event
	// type]] object
	TypeId int `json:"typeId"`
}