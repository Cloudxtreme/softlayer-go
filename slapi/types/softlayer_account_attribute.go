package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_Attribute - Many SoftLayer customer accounts have individual attributes assigned
// to them that describe features or special features for that account, such as special pricing,
// account statuses, and ordering instructions. The SoftLayer_Account_Attribute data type contains
// information relating to a single SoftLayer_Account attribute.
type SoftLayer_Account_Attribute struct {

	// AccountAttributeTypeId - The internal identifier of the type of attribute that a SoftLayer customer
	// account attribute belongs to.
	AccountAttributeTypeId int `json:"accountAttributeTypeId"`

	// AccountId - The internal identifier of the SoftLayer customer account that is assigned an account
	// attribute.
	AccountId int `json:"accountId"`

	// Id - A SoftLayer customer account attribute's internal identifier.
	Id int `json:"id"`

	// Value - no documentation
	Value string `json:"value"`
}

// SoftLayer_Account_Attribute_Extended is SoftLayer_Account_Attribute with all maskable types.
type SoftLayer_Account_Attribute_Extended struct {
	SoftLayer_Account_Attribute

	// AccountAttributeType - The type of attribute assigned to a SoftLayer customer account.
	AccountAttributeType *SoftLayer_Account_Attribute_Type `json:"accountAttributeType"`

	// Account - The SoftLayer customer account that has an attribute.
	Account *SoftLayer_Account `json:"account"`
}

func (softlayer_account_attribute *SoftLayer_Account_Attribute) String() string {
	return "SoftLayer_Account_Attribute"
}
