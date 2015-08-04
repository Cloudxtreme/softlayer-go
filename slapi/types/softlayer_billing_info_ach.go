package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Info_Ach - <nil>
type SoftLayer_Billing_Info_Ach struct {

	// AccountId - <nil>
	AccountId int `json:"accountId,omitempty"`

	// AccountNumber - <nil>
	AccountNumber string `json:"accountNumber,omitempty"`

	// AccountType - <nil>
	AccountType string `json:"accountType,omitempty"`

	// Status - <nil>
	Status string `json:"status,omitempty"`

	// VerifiedDate - <nil>
	VerifiedDate *time.Time `json:"verifiedDate,omitempty"`

	// City - <nil>
	City string `json:"city,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// PhoneNumber - <nil>
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// Street1 - <nil>
	Street1 string `json:"street1,omitempty"`

	// LastName - <nil>
	LastName string `json:"lastName,omitempty"`

	// Postalcode - <nil>
	Postalcode string `json:"postalcode,omitempty"`

	// State - <nil>
	State string `json:"state,omitempty"`

	// BankTransitNumber - <nil>
	BankTransitNumber string `json:"bankTransitNumber,omitempty"`

	// Country - <nil>
	Country string `json:"country,omitempty"`

	// FirstName - <nil>
	FirstName string `json:"firstName,omitempty"`

	// Street2 - <nil>
	Street2 string `json:"street2,omitempty"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`
}

func (softlayer_billing_info_ach *SoftLayer_Billing_Info_Ach) String() string {
	return "SoftLayer_Billing_Info_Ach"
}
