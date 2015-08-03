package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Subnet_Registration_Event_Type - Subnet Registration Event Type objects describe
// the nature of a [[SoftLayer_Network_Subnet_Registration_Event]] The standard values for these
// objects are as follows: - Indicates that the registration has been created - Indicates that the
// registration has been updated - Indicates that the registration has been cancelled - Indicates that
// an action taken against the RIR has produced a response. More details will be provided in the event
// message. - Indicates that an error has been encountered. More details will be provided in the event
// message. - An employee or other system has entered a note regarding the registration. The note
// content will be provided in the event message.
type SoftLayer_Network_Subnet_Registration_Event_Type struct {

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Id - no documentation
	Id int `json:"id"`
}

func (softlayer_network_subnet_registration_event_type *SoftLayer_Network_Subnet_Registration_Event_Type) String() string {
	return "SoftLayer_Network_Subnet_Registration_Event_Type"
}
