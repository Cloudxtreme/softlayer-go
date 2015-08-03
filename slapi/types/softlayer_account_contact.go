package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Contact - <nil>
type SoftLayer_Account_Contact struct {

	// FirstName - <nil>
	FirstName string `json:"firstName"`

	// JobTitle - <nil>
	JobTitle string `json:"jobTitle"`

	// ProfileName - <nil>
	ProfileName string `json:"profileName"`

	// Url - <nil>
	Url string `json:"url"`

	// AlternatePhone - <nil>
	AlternatePhone string `json:"alternatePhone"`

	// Address1 - <nil>
	Address1 string `json:"address1"`

	// CompanyName - <nil>
	CompanyName string `json:"companyName"`

	// Country - <nil>
	Country string `json:"country"`

	// Email - <nil>
	Email string `json:"email"`

	// FaxPhone - <nil>
	FaxPhone string `json:"faxPhone"`

	// AccountId - <nil>
	AccountId int `json:"accountId"`

	// State - <nil>
	State string `json:"state"`

	// OfficePhone - <nil>
	OfficePhone string `json:"officePhone"`

	// PostalCode - <nil>
	PostalCode string `json:"postalCode"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Id - <nil>
	Id int `json:"id"`

	// LastName - <nil>
	LastName string `json:"lastName"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// TypeId - <nil>
	TypeId int `json:"typeId"`

	// Address2 - <nil>
	Address2 string `json:"address2"`

	// City - <nil>
	City string `json:"city"`
}

// SoftLayer_Account_Contact_Extended is SoftLayer_Account_Contact with all maskable types.
type SoftLayer_Account_Contact_Extended struct {
	SoftLayer_Account_Contact

	// Type - <nil>
	Type *SoftLayer_Account_Contact_Type `json:"type"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`
}

func (softlayer_account_contact *SoftLayer_Account_Contact) String() string {
	return "SoftLayer_Account_Contact"
}
