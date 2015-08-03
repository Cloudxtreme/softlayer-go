package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Info_Ach - <nil>
type SoftLayer_Billing_Info_Ach struct {

	// BankTransitNumber - <nil>
	BankTransitNumber string `json:"bankTransitNumber"`

	// Id - <nil>
	Id int `json:"id"`

	// PhoneNumber - <nil>
	PhoneNumber string `json:"phoneNumber"`

	// VerifiedDate - <nil>
	VerifiedDate *time.Time `json:"verifiedDate"`

	// AccountId - <nil>
	AccountId int `json:"accountId"`

	// AccountType - <nil>
	AccountType string `json:"accountType"`

	// City - <nil>
	City string `json:"city"`

	// Postalcode - <nil>
	Postalcode string `json:"postalcode"`

	// Street1 - <nil>
	Street1 string `json:"street1"`

	// AccountNumber - <nil>
	AccountNumber string `json:"accountNumber"`

	// Country - <nil>
	Country string `json:"country"`

	// LastName - <nil>
	LastName string `json:"lastName"`

	// FirstName - <nil>
	FirstName string `json:"firstName"`

	// State - <nil>
	State string `json:"state"`

	// Status - <nil>
	Status string `json:"status"`

	// Street2 - <nil>
	Street2 string `json:"street2"`
}

// SoftLayer_Billing_Info_Ach_Extended is SoftLayer_Billing_Info_Ach with all maskable types.
type SoftLayer_Billing_Info_Ach_Extended struct {
	SoftLayer_Billing_Info_Ach

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`
}

func (softlayer_billing_info_ach *SoftLayer_Billing_Info_Ach) String() string {
	return "SoftLayer_Billing_Info_Ach"
}
