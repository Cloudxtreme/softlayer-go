package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Contact - <nil>
type SoftLayer_Account_Contact struct {

	// City - <nil>
	City string `json:"city"`

	// FaxPhone - <nil>
	FaxPhone string `json:"faxPhone"`

	// Id - <nil>
	Id int `json:"id"`

	// JobTitle - <nil>
	JobTitle string `json:"jobTitle"`

	// LastName - <nil>
	LastName string `json:"lastName"`

	// OfficePhone - <nil>
	OfficePhone string `json:"officePhone"`

	// ProfileName - <nil>
	ProfileName string `json:"profileName"`

	// Url - <nil>
	Url string `json:"url"`

	// FirstName - <nil>
	FirstName string `json:"firstName"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// PostalCode - <nil>
	PostalCode string `json:"postalCode"`

	// Email - <nil>
	Email string `json:"email"`

	// AccountId - <nil>
	AccountId int `json:"accountId"`

	// Address1 - <nil>
	Address1 string `json:"address1"`

	// Address2 - <nil>
	Address2 string `json:"address2"`

	// Country - <nil>
	Country string `json:"country"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// State - <nil>
	State string `json:"state"`

	// CompanyName - <nil>
	CompanyName string `json:"companyName"`

	// TypeId - <nil>
	TypeId int `json:"typeId"`

	// AlternatePhone - <nil>
	AlternatePhone string `json:"alternatePhone"`
}

func (softlayer_account_contact *SoftLayer_Account_Contact) String() string {
	return "SoftLayer_Account_Contact"
}

// SoftLayer_Account_Contact_Extended is SoftLayer_Account_Contact with all maskable types.
type SoftLayer_Account_Contact_Extended struct {
	SoftLayer_Account_Contact

	// Type - <nil>
	Type *SoftLayer_Account_Contact_Type `json:"type"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`
}

func (softlayer_account_contact *SoftLayer_Account_Contact_Extended) String() string {
	return "SoftLayer_Account_Contact"
}
