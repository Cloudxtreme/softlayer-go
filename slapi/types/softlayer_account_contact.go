package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Contact - <nil>
type SoftLayer_Account_Contact struct {

	// PostalCode - <nil>
	PostalCode string `json:"postalCode,omitempty"`

	// City - <nil>
	City string `json:"city,omitempty"`

	// Email - <nil>
	Email string `json:"email,omitempty"`

	// FaxPhone - <nil>
	FaxPhone string `json:"faxPhone,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// OfficePhone - <nil>
	OfficePhone string `json:"officePhone,omitempty"`

	// FirstName - <nil>
	FirstName string `json:"firstName,omitempty"`

	// State - <nil>
	State string `json:"state,omitempty"`

	// TypeId - <nil>
	TypeId int `json:"typeId,omitempty"`

	// Address2 - <nil>
	Address2 string `json:"address2,omitempty"`

	// AlternatePhone - <nil>
	AlternatePhone string `json:"alternatePhone,omitempty"`

	// CompanyName - <nil>
	CompanyName string `json:"companyName,omitempty"`

	// Country - <nil>
	Country string `json:"country,omitempty"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// AccountId - <nil>
	AccountId int `json:"accountId,omitempty"`

	// Url - <nil>
	Url string `json:"url,omitempty"`

	// ProfileName - <nil>
	ProfileName string `json:"profileName,omitempty"`

	// Address1 - <nil>
	Address1 string `json:"address1,omitempty"`

	// JobTitle - <nil>
	JobTitle string `json:"jobTitle,omitempty"`

	// LastName - <nil>
	LastName string `json:"lastName,omitempty"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate,omitempty"`
}

func (softlayer_account_contact *SoftLayer_Account_Contact) String() string {
	return "SoftLayer_Account_Contact"
}

// SoftLayer_Account_Contact_Extended is SoftLayer_Account_Contact with all maskable types.
type SoftLayer_Account_Contact_Extended struct {
	SoftLayer_Account_Contact

	// Type - <nil>
	Type *SoftLayer_Account_Contact_Type `json:"type,omitempty"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`
}

func (softlayer_account_contact *SoftLayer_Account_Contact_Extended) String() string {
	return "SoftLayer_Account_Contact"
}
