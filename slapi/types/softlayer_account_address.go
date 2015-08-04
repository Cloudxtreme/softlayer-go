package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_Address - The SoftLayer_Account_Address data type contains information on an
// address associated with a SoftLayer account.
type SoftLayer_Account_Address struct {

	// AccountId - <nil>
	AccountId int `json:"accountId,omitempty"`

	// Address2 - no documentation
	Address2 string `json:"address2,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// IsActive - no documentation
	IsActive int `json:"isActive,omitempty"`

	// City - no documentation
	City string `json:"city,omitempty"`

	// Address1 - Line 1 of the address (normally the street address).
	Address1 string `json:"address1,omitempty"`

	// Country - no documentation
	Country string `json:"country,omitempty"`

	// LocationId - no documentation
	LocationId int `json:"locationId,omitempty"`

	// PostalCode - no documentation
	PostalCode string `json:"postalCode,omitempty"`

	// State - no documentation
	State string `json:"state,omitempty"`

	// ContactName - no documentation
	ContactName string `json:"contactName,omitempty"`

	// CreateUser - no documentation
	CreateUser *SoftLayer_User_Customer `json:"createUser,omitempty"`

	// ModifyUser - no documentation
	ModifyUser *SoftLayer_User_Customer `json:"modifyUser,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// ModifyEmployee - no documentation
	ModifyEmployee *SoftLayer_User_Employee `json:"modifyEmployee,omitempty"`

	// Type - no documentation
	Type *SoftLayer_Account_Address_Type `json:"type,omitempty"`

	// Location - no documentation
	Location *SoftLayer_Location `json:"location,omitempty"`
}

func (softlayer_account_address *SoftLayer_Account_Address) String() string {
	return "SoftLayer_Account_Address"
}
