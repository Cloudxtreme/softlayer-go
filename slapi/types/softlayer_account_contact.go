package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
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

// CreateObject - This method creates an account contact. The accountId is fixed, other properties can
// be set during creation. The typeId indicates the SoftLayer_Account_Contact_Type for the contact.
// This method returns the SoftLayer_Account_Contact object that is created.
func (softlayer_account_contact *SoftLayer_Account_Contact) CreateObject(ctx *slapi.RequestContext, templateObject SoftLayer_Account_Contact) (*SoftLayer_Account_Contact, error) {
	var returnValue *SoftLayer_Account_Contact
	return returnValue, nil
}

// DeleteObject - deleteObject permanently removes an account contact
func (softlayer_account_contact *SoftLayer_Account_Contact) DeleteObject(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - This method allows you to modify an account contact. Only master users are permitted to
// modify an account contact.
func (softlayer_account_contact *SoftLayer_Account_Contact) EditObject(ctx *slapi.RequestContext, templateObject SoftLayer_Account_Contact) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetAllContactTypes - This method will return an array of SoftLayer_Account_Contact_Type objects
// which can be used when creating or editing an account contact.
func (softlayer_account_contact *SoftLayer_Account_Contact) GetAllContactTypes(ctx *slapi.RequestContext) ([]*SoftLayer_Account_Contact_Type, error) {
	var returnValue []*SoftLayer_Account_Contact_Type
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_account_contact *SoftLayer_Account_Contact) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Account_Contact, error) {
	var returnValue *SoftLayer_Account_Contact
	return returnValue, nil
}
