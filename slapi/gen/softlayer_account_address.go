package sl

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

// CreateObject - Create a new address record. The ''typeId'', ''accountId'', ''description'',
// ''address1'', ''city'', ''state'', ''country'', and ''postalCode'' properties in the templateObject
// parameter are required properties and may not be null or empty. Users will be restricted to creating
// addresses for their account.
func (softlayer_account_address *SoftLayer_Account_Address) CreateObject(templateObject SoftLayer_Account_Address) (*SoftLayer_Account_Address, error) {
	var returnValue *SoftLayer_Account_Address
	return returnValue, nil
}

// EditObject - Edit the properties of an address record by passing in a modified instance of a
// SoftLayer_Account_Address object. Users will be restricted to modifying addresses for their account.
func (softlayer_account_address *SoftLayer_Account_Address) EditObject(templateObject SoftLayer_Account_Address) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetAllDataCenters - no documentation
func (softlayer_account_address *SoftLayer_Account_Address) GetAllDataCenters() ([]*SoftLayer_Account_Address, error) {
	var returnValue []*SoftLayer_Account_Address
	return returnValue, nil
}

// GetNetworkAddress - no documentation
func (softlayer_account_address *SoftLayer_Account_Address) GetNetworkAddress(name string) ([]*SoftLayer_Account_Address, error) {
	var returnValue []*SoftLayer_Account_Address
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_account_address *SoftLayer_Account_Address) GetObject() (*SoftLayer_Account_Address, error) {
	var returnValue *SoftLayer_Account_Address
	return returnValue, nil
}
