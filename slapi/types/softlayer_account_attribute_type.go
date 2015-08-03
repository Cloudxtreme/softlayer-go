package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_Attribute_Type - SoftLayer_Account_Attribute_Type models the type of attribute
// that can be assigned to a SoftLayer customer account.
type SoftLayer_Account_Attribute_Type struct {

	// Description - A brief description of a SoftLayer account attribute type.
	Description string `json:"description"`

	// Id - A SoftLayer account attribute type's internal identifier.
	Id int `json:"id"`

	// KeyName - A SoftLayer account attribute type's key name. This is typically a shorter version of an
	// attribute type's name.
	KeyName string `json:"keyName"`

	// Name - no documentation
	Name string `json:"name"`
}

func (softlayer_account_attribute_type *SoftLayer_Account_Attribute_Type) String() string {
	return "SoftLayer_Account_Attribute_Type"
}
