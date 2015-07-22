package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Dns_Domain_Registration_Transfer_Information - Transfer Information container
// for domain registration
type SoftLayer_Container_Dns_Domain_Registration_Transfer_Information struct {

	// Reason - no documentation
	Reason string `json:"reason"`

	// RegistrantEmail - no documentation
	RegistrantEmail string `json:"registrantEmail"`

	// Status - no documentation
	Status string `json:"status"`

	// TimeStamp - The date and time of the most recent update to the state of the transfer.
	TimeStamp *time.Time `json:"timeStamp"`

	// Transferrable - no documentation
	Transferrable int `json:"transferrable"`
}
