package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Software_Component_Password_History - This object allows you to find the history of
// password changes for a specific SoftLayer_Software Component
type SoftLayer_Software_Component_Password_History struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Notes - A note string stored for this username/password pair.
	Notes string `json:"notes"`

	// Password - The password part of this specific password history instance.
	Password string `json:"password"`

	// SoftwareComponentId - The id number for the Software Component this username/password pair is for.
	SoftwareComponentId int `json:"softwareComponentId"`

	// Username - The username part of this specific password history instance.
	Username string `json:"username"`
}

// SoftLayer_Software_Component_Password_History_Extended is SoftLayer_Software_Component_Password_History with all maskable types.
type SoftLayer_Software_Component_Password_History_Extended struct {
	SoftLayer_Software_Component_Password_History

	// SoftwareComponent - An installed and licensed instance of a piece of software
	SoftwareComponent *SoftLayer_Software_Component `json:"softwareComponent"`
}

func (softlayer_software_component_password_history *SoftLayer_Software_Component_Password_History) String() string {
	return "SoftLayer_Software_Component_Password_History"
}
