package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_Address - The SoftLayer_Account_Address data type contains information on an
// address associated with a SoftLayer account.
type SoftLayer_Account_Address struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AccountId - <nil>
	AccountId int `json:"accountId"`

	// Address1 - Line 1 of the address (normally the street address).
	Address1 string `json:"address1"`

	// Address2 - no documentation
	Address2 string `json:"address2"`

	// City - no documentation
	City string `json:"city"`

	// ContactName - no documentation
	ContactName string `json:"contactName"`

	// Country - no documentation
	Country string `json:"country"`

	// CreateUser - no documentation
	CreateUser *SoftLayer_User_Customer `json:"createUser"`

	// Description - no documentation
	Description string `json:"description"`

	// Id - no documentation
	Id int `json:"id"`

	// IsActive - no documentation
	IsActive int `json:"isActive"`

	// Location - no documentation
	Location *SoftLayer_Location `json:"location"`

	// LocationId - no documentation
	LocationId int `json:"locationId"`

	// ModifyEmployee - no documentation
	ModifyEmployee *SoftLayer_User_Employee `json:"modifyEmployee"`

	// ModifyUser - no documentation
	ModifyUser *SoftLayer_User_Customer `json:"modifyUser"`

	// PostalCode - no documentation
	PostalCode string `json:"postalCode"`

	// State - no documentation
	State string `json:"state"`

	// Type - no documentation
	Type *SoftLayer_Account_Address_Type `json:"type"`
}

func (softlayer_account_address *SoftLayer_Account_Address) String() string {
	return "SoftLayer_Account_Address"
}
