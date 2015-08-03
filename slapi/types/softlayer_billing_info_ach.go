package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Info_Ach - <nil>
type SoftLayer_Billing_Info_Ach struct {

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// AccountId - <nil>
	AccountId int `json:"accountId"`

	// AccountNumber - <nil>
	AccountNumber string `json:"accountNumber"`

	// AccountType - <nil>
	AccountType string `json:"accountType"`

	// BankTransitNumber - <nil>
	BankTransitNumber string `json:"bankTransitNumber"`

	// City - <nil>
	City string `json:"city"`

	// Country - <nil>
	Country string `json:"country"`

	// FirstName - <nil>
	FirstName string `json:"firstName"`

	// Id - <nil>
	Id int `json:"id"`

	// LastName - <nil>
	LastName string `json:"lastName"`

	// PhoneNumber - <nil>
	PhoneNumber string `json:"phoneNumber"`

	// Postalcode - <nil>
	Postalcode string `json:"postalcode"`

	// State - <nil>
	State string `json:"state"`

	// Status - <nil>
	Status string `json:"status"`

	// Street1 - <nil>
	Street1 string `json:"street1"`

	// Street2 - <nil>
	Street2 string `json:"street2"`

	// VerifiedDate - <nil>
	VerifiedDate *time.Time `json:"verifiedDate"`
}

func (softlayer_billing_info_ach *SoftLayer_Billing_Info_Ach) String() string {
	return "SoftLayer_Billing_Info_Ach"
}
