package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Software_Component_Password_History - This object allows you to find the history of
// password changes for a specific SoftLayer_Software Component
type SoftLayer_Software_Component_Password_History struct {

	// Notes - A note string stored for this username/password pair.
	Notes string `json:"notes,omitempty"`

	// Password - The password part of this specific password history instance.
	Password string `json:"password,omitempty"`

	// SoftwareComponentId - The id number for the Software Component this username/password pair is for.
	SoftwareComponentId int `json:"softwareComponentId,omitempty"`

	// Username - The username part of this specific password history instance.
	Username string `json:"username,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// SoftwareComponent - An installed and licensed instance of a piece of software
	SoftwareComponent *SoftLayer_Software_Component `json:"softwareComponent,omitempty"`
}

func (softlayer_software_component_password_history *SoftLayer_Software_Component_Password_History) String() string {
	return "SoftLayer_Software_Component_Password_History"
}
