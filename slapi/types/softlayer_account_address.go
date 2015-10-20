package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_Address - The SoftLayer_Account_Address data type contains information on an
// address associated with a SoftLayer account.
type SoftLayer_Account_Address struct {

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// AccountId - <nil>
	AccountId int `json:"accountId,omitempty"`

	// Address1 - Line 1 of the address (normally the street address).
	Address1 string `json:"address1,omitempty"`

	// Address2 - no documentation
	Address2 string `json:"address2,omitempty"`

	// Country - no documentation
	Country string `json:"country,omitempty"`

	// LocationId - no documentation
	LocationId int `json:"locationId,omitempty"`

	// IsActive - no documentation
	IsActive int `json:"isActive,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// PostalCode - no documentation
	PostalCode string `json:"postalCode,omitempty"`

	// City - no documentation
	City string `json:"city,omitempty"`

	// ContactName - no documentation
	ContactName string `json:"contactName,omitempty"`

	// State - no documentation
	State string `json:"state,omitempty"`

	// ModifyEmployee - no documentation
	ModifyEmployee *SoftLayer_User_Employee `json:"modifyEmployee,omitempty"`

	// ModifyUser - no documentation
	ModifyUser *SoftLayer_User_Customer `json:"modifyUser,omitempty"`

	// Type - no documentation
	Type *SoftLayer_Account_Address_Type `json:"type,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// CreateUser - no documentation
	CreateUser *SoftLayer_User_Customer `json:"createUser,omitempty"`

	// Location - no documentation
	Location *SoftLayer_Location `json:"location,omitempty"`
}

func (softlayer_account_address *SoftLayer_Account_Address) String() string {
	return "SoftLayer_Account_Address"
}
