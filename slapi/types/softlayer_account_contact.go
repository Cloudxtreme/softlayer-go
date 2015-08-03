package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Contact - <nil>
type SoftLayer_Account_Contact struct {

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// AccountId - <nil>
	AccountId int `json:"accountId"`

	// Address1 - <nil>
	Address1 string `json:"address1"`

	// Address2 - <nil>
	Address2 string `json:"address2"`

	// AlternatePhone - <nil>
	AlternatePhone string `json:"alternatePhone"`

	// City - <nil>
	City string `json:"city"`

	// CompanyName - <nil>
	CompanyName string `json:"companyName"`

	// Country - <nil>
	Country string `json:"country"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Email - <nil>
	Email string `json:"email"`

	// FaxPhone - <nil>
	FaxPhone string `json:"faxPhone"`

	// FirstName - <nil>
	FirstName string `json:"firstName"`

	// Id - <nil>
	Id int `json:"id"`

	// JobTitle - <nil>
	JobTitle string `json:"jobTitle"`

	// LastName - <nil>
	LastName string `json:"lastName"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// OfficePhone - <nil>
	OfficePhone string `json:"officePhone"`

	// PostalCode - <nil>
	PostalCode string `json:"postalCode"`

	// ProfileName - <nil>
	ProfileName string `json:"profileName"`

	// State - <nil>
	State string `json:"state"`

	// Type - <nil>
	Type *SoftLayer_Account_Contact_Type `json:"type"`

	// TypeId - <nil>
	TypeId int `json:"typeId"`

	// Url - <nil>
	Url string `json:"url"`
}

func (softlayer_account_contact *SoftLayer_Account_Contact) String() string {
	return "SoftLayer_Account_Contact"
}
