package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_Address - The SoftLayer_Account_Address data type contains information on an
// address associated with a SoftLayer account.
type SoftLayer_Account_Address struct {

	// Address1 - Line 1 of the address (normally the street address).
	Address1 string `json:"address1,omitempty"`

	// Address2 - no documentation
	Address2 string `json:"address2,omitempty"`

	// PostalCode - no documentation
	PostalCode string `json:"postalCode,omitempty"`

	// State - no documentation
	State string `json:"state,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// IsActive - no documentation
	IsActive int `json:"isActive,omitempty"`

	// AccountId - <nil>
	AccountId int `json:"accountId,omitempty"`

	// City - no documentation
	City string `json:"city,omitempty"`

	// ContactName - no documentation
	ContactName string `json:"contactName,omitempty"`

	// Country - no documentation
	Country string `json:"country,omitempty"`

	// LocationId - no documentation
	LocationId int `json:"locationId,omitempty"`
}

func (softlayer_account_address *SoftLayer_Account_Address) String() string {
	return "SoftLayer_Account_Address"
}

// SoftLayer_Account_Address_Extended is SoftLayer_Account_Address with all maskable types.
type SoftLayer_Account_Address_Extended struct {
	SoftLayer_Account_Address

	// Location - no documentation
	Location *SoftLayer_Location `json:"location,omitempty"`

	// Type - no documentation
	Type *SoftLayer_Account_Address_Type `json:"type,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// CreateUser - no documentation
	CreateUser *SoftLayer_User_Customer `json:"createUser,omitempty"`

	// ModifyUser - no documentation
	ModifyUser *SoftLayer_User_Customer `json:"modifyUser,omitempty"`

	// ModifyEmployee - no documentation
	ModifyEmployee *SoftLayer_User_Employee `json:"modifyEmployee,omitempty"`
}

func (softlayer_account_address *SoftLayer_Account_Address_Extended) String() string {
	return "SoftLayer_Account_Address"
}
