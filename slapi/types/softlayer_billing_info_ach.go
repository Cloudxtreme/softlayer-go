package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Info_Ach - <nil>
type SoftLayer_Billing_Info_Ach struct {

	// FirstName - <nil>
	FirstName string `json:"firstName,omitempty"`

	// Street2 - <nil>
	Street2 string `json:"street2,omitempty"`

	// AccountNumber - <nil>
	AccountNumber string `json:"accountNumber,omitempty"`

	// BankTransitNumber - <nil>
	BankTransitNumber string `json:"bankTransitNumber,omitempty"`

	// City - <nil>
	City string `json:"city,omitempty"`

	// VerifiedDate - <nil>
	VerifiedDate *time.Time `json:"verifiedDate,omitempty"`

	// AccountId - <nil>
	AccountId int `json:"accountId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// PhoneNumber - <nil>
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// Status - <nil>
	Status string `json:"status,omitempty"`

	// Country - <nil>
	Country string `json:"country,omitempty"`

	// LastName - <nil>
	LastName string `json:"lastName,omitempty"`

	// State - <nil>
	State string `json:"state,omitempty"`

	// Street1 - <nil>
	Street1 string `json:"street1,omitempty"`

	// AccountType - <nil>
	AccountType string `json:"accountType,omitempty"`

	// Postalcode - <nil>
	Postalcode string `json:"postalcode,omitempty"`
}

func (softlayer_billing_info_ach *SoftLayer_Billing_Info_Ach) String() string {
	return "SoftLayer_Billing_Info_Ach"
}

// SoftLayer_Billing_Info_Ach_Extended is SoftLayer_Billing_Info_Ach with all maskable types.
type SoftLayer_Billing_Info_Ach_Extended struct {
	SoftLayer_Billing_Info_Ach

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`
}

func (softlayer_billing_info_ach *SoftLayer_Billing_Info_Ach_Extended) String() string {
	return "SoftLayer_Billing_Info_Ach"
}
