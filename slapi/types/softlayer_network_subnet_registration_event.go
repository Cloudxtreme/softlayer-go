package types

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
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Message - A string message indicating what took place during this event
	Message string `json:"message,omitempty"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// RegistrationId - The numeric ID of the related [[SoftLayer_Network_Subnet_Registration]] object
	RegistrationId int `json:"registrationId,omitempty"`

	// TypeId - The numeric ID of the associated [[SoftLayer_Network_Subnet_Registration_Event_Type|event
	// type]] object
	TypeId int `json:"typeId,omitempty"`

	// Registration - no documentation
	Registration *SoftLayer_Network_Subnet_Registration `json:"registration,omitempty"`

	// Type - no documentation
	Type *SoftLayer_Network_Subnet_Registration_Event_Type `json:"type,omitempty"`
}

func (softlayer_network_subnet_registration_event *SoftLayer_Network_Subnet_Registration_Event) String() string {
	return "SoftLayer_Network_Subnet_Registration_Event"
}
