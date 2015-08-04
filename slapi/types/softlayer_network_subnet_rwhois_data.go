package types

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

	// AccountId - An account's data's associated account identifier.
	AccountId int `json:"accountId,omitempty"`

	// Address1 - The first line of the mailing address associated with an account's data.
	Address1 string `json:"address1,omitempty"`

	// State - A two-letter abbreviation of the state of the mailing address associated with an account's
	// data. If an account does not reside in a province then this is typically blank.
	State string `json:"state,omitempty"`

	// AbuseEmail - An email address associated with an account's data that is responsible for responding
	// to network abuse queries about malicious traffic coming from your servers' IP addresses.
	AbuseEmail string `json:"abuseEmail,omitempty"`

	// LastName - no documentation
	LastName string `json:"lastName,omitempty"`

	// PrivateResidenceFlag - Whether an account's data refers to a private residence or not.
	PrivateResidenceFlag bool `json:"privateResidenceFlag,omitempty"`

	// Address2 - The second line of the mailing address associated with an account's data.
	Address2 string `json:"address2,omitempty"`

	// City - The city of the mailing address associated with an account's data.
	City string `json:"city,omitempty"`

	// Country - A two-letter abbreviation of the country of the mailing address associated with an
	// account's data.
	Country string `json:"country,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// FirstName - no documentation
	FirstName string `json:"firstName,omitempty"`

	// CompanyName - The company name associated with an account's data.
	CompanyName string `json:"companyName,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// PostalCode - The postal code of the mailing address associated with an account's data.
	PostalCode string `json:"postalCode,omitempty"`
}

func (softlayer_network_subnet_rwhois_data *SoftLayer_Network_Subnet_Rwhois_Data) String() string {
	return "SoftLayer_Network_Subnet_Rwhois_Data"
}

// SoftLayer_Network_Subnet_Rwhois_Data_Extended is SoftLayer_Network_Subnet_Rwhois_Data with all maskable types.
type SoftLayer_Network_Subnet_Rwhois_Data_Extended struct {
	SoftLayer_Network_Subnet_Rwhois_Data

	// Account - The SoftLayer customer account associated with this reverse data.
	Account *SoftLayer_Account `json:"account,omitempty"`
}

func (softlayer_network_subnet_rwhois_data *SoftLayer_Network_Subnet_Rwhois_Data_Extended) String() string {
	return "SoftLayer_Network_Subnet_Rwhois_Data"
}
