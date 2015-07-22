package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Subnet_Rwhois_Data - Every SoftLayer customer account has contact information
// associated with it for reverse purposes. An account's data, modeled by the
// SoftLayer_Network_Subnet_Rwhois_Data data type, is used by SoftLayer's reverse server as well as for
// transactions. SoftLayer's reverse servers respond to queries for IP addresses belonging to a
// customer's servers, returning this data. A SoftLayer customer's data may not necessarily match their
// account or portal users' contact information.
type SoftLayer_Network_Subnet_Rwhois_Data struct {

	// AbuseEmail - An email address associated with an account's data that is responsible for responding
	// to network abuse queries about malicious traffic coming from your servers' IP addresses.
	AbuseEmail string `json:"abuseEmail"`

	// Account - The SoftLayer customer account associated with this reverse data.
	Account *SoftLayer_Account `json:"account"`

	// AccountId - An account's data's associated account identifier.
	AccountId int `json:"accountId"`

	// Address1 - The first line of the mailing address associated with an account's data.
	Address1 string `json:"address1"`

	// Address2 - The second line of the mailing address associated with an account's data.
	Address2 string `json:"address2"`

	// City - The city of the mailing address associated with an account's data.
	City string `json:"city"`

	// CompanyName - The company name associated with an account's data.
	CompanyName string `json:"companyName"`

	// Country - A two-letter abbreviation of the country of the mailing address associated with an
	// account's data.
	Country string `json:"country"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// FirstName - no documentation
	FirstName string `json:"firstName"`

	// Id - no documentation
	Id int `json:"id"`

	// LastName - no documentation
	LastName string `json:"lastName"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// PostalCode - The postal code of the mailing address associated with an account's data.
	PostalCode string `json:"postalCode"`

	// PrivateResidenceFlag - Whether an account's data refers to a private residence or not.
	PrivateResidenceFlag bool `json:"privateResidenceFlag"`

	// State - A two-letter abbreviation of the state of the mailing address associated with an account's
	// data. If an account does not reside in a province then this is typically blank.
	State string `json:"state"`
}

// EditObject - Edit the record by passing in a modified version of the record object. All fields are
// editable.
func (softlayer_network_subnet_rwhois_data *SoftLayer_Network_Subnet_Rwhois_Data) EditObject(templateObject SoftLayer_Network_Subnet_Rwhois_Data) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Network_Subnet_Rwhois_Data object whose ID corresponds
// to the ID number of the init parameter passed to the SoftLayer_Network_Subnet_Rwhois_Data service.
// The best way to get Rwhois Data for an account is through getRhwoisData on the Account service.
func (softlayer_network_subnet_rwhois_data *SoftLayer_Network_Subnet_Rwhois_Data) GetObject() (*SoftLayer_Network_Subnet_Rwhois_Data, error) {
	var returnValue *SoftLayer_Network_Subnet_Rwhois_Data
	return returnValue, nil
}
