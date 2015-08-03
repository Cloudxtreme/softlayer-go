package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_Address - The SoftLayer_Account_Address data type contains information on an
// address associated with a SoftLayer account.
type SoftLayer_Account_Address struct {

	// PostalCode - no documentation
	PostalCode string `json:"postalCode"`

	// AccountId - <nil>
	AccountId int `json:"accountId"`

	// Id - no documentation
	Id int `json:"id"`

	// IsActive - no documentation
	IsActive int `json:"isActive"`

	// Address2 - no documentation
	Address2 string `json:"address2"`

	// ContactName - no documentation
	ContactName string `json:"contactName"`

	// Description - no documentation
	Description string `json:"description"`

	// State - no documentation
	State string `json:"state"`

	// LocationId - no documentation
	LocationId int `json:"locationId"`

	// Address1 - Line 1 of the address (normally the street address).
	Address1 string `json:"address1"`

	// City - no documentation
	City string `json:"city"`

	// Country - no documentation
	Country string `json:"country"`
}

func (softlayer_account_address *SoftLayer_Account_Address) String() string {
	return "SoftLayer_Account_Address"
}

// SoftLayer_Account_Address_Extended is SoftLayer_Account_Address with all maskable types.
type SoftLayer_Account_Address_Extended struct {
	SoftLayer_Account_Address

	// Location - no documentation
	Location *SoftLayer_Location `json:"location"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// CreateUser - no documentation
	CreateUser *SoftLayer_User_Customer `json:"createUser"`

	// ModifyEmployee - no documentation
	ModifyEmployee *SoftLayer_User_Employee `json:"modifyEmployee"`

	// ModifyUser - no documentation
	ModifyUser *SoftLayer_User_Customer `json:"modifyUser"`

	// Type - no documentation
	Type *SoftLayer_Account_Address_Type `json:"type"`
}

func (softlayer_account_address *SoftLayer_Account_Address_Extended) String() string {
	return "SoftLayer_Account_Address"
}
