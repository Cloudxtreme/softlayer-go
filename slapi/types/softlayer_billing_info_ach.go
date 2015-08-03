package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Info_Ach - <nil>
type SoftLayer_Billing_Info_Ach struct {

	// Country - <nil>
	Country string `json:"country"`

	// FirstName - <nil>
	FirstName string `json:"firstName"`

	// Id - <nil>
	Id int `json:"id"`

	// PhoneNumber - <nil>
	PhoneNumber string `json:"phoneNumber"`

	// Street1 - <nil>
	Street1 string `json:"street1"`

	// BankTransitNumber - <nil>
	BankTransitNumber string `json:"bankTransitNumber"`

	// Status - <nil>
	Status string `json:"status"`

	// AccountId - <nil>
	AccountId int `json:"accountId"`

	// AccountType - <nil>
	AccountType string `json:"accountType"`

	// LastName - <nil>
	LastName string `json:"lastName"`

	// Postalcode - <nil>
	Postalcode string `json:"postalcode"`

	// State - <nil>
	State string `json:"state"`

	// VerifiedDate - <nil>
	VerifiedDate *time.Time `json:"verifiedDate"`

	// AccountNumber - <nil>
	AccountNumber string `json:"accountNumber"`

	// City - <nil>
	City string `json:"city"`

	// Street2 - <nil>
	Street2 string `json:"street2"`
}

func (softlayer_billing_info_ach *SoftLayer_Billing_Info_Ach) String() string {
	return "SoftLayer_Billing_Info_Ach"
}

// SoftLayer_Billing_Info_Ach_Extended is SoftLayer_Billing_Info_Ach with all maskable types.
type SoftLayer_Billing_Info_Ach_Extended struct {
	SoftLayer_Billing_Info_Ach

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`
}

func (softlayer_billing_info_ach *SoftLayer_Billing_Info_Ach_Extended) String() string {
	return "SoftLayer_Billing_Info_Ach"
}
