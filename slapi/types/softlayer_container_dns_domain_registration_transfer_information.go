package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Dns_Domain_Registration_Transfer_Information - Transfer Information container
// for domain registration
type SoftLayer_Container_Dns_Domain_Registration_Transfer_Information struct {

	// Transferrable - no documentation
	Transferrable int `json:"transferrable,omitempty"`

	// Reason - no documentation
	Reason string `json:"reason,omitempty"`

	// RegistrantEmail - no documentation
	RegistrantEmail string `json:"registrantEmail,omitempty"`

	// Status - no documentation
	Status string `json:"status,omitempty"`

	// TimeStamp - The date and time of the most recent update to the state of the transfer.
	TimeStamp *time.Time `json:"timeStamp,omitempty"`
}

func (softlayer_container_dns_domain_registration_transfer_information *SoftLayer_Container_Dns_Domain_Registration_Transfer_Information) String() string {
	return "SoftLayer_Container_Dns_Domain_Registration_Transfer_Information"
}
