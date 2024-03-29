package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_Attribute_Type - SoftLayer_Account_Attribute_Type models the type of attribute
// that can be assigned to a SoftLayer customer account.
type SoftLayer_Account_Attribute_Type struct {

	// KeyName - A SoftLayer account attribute type's key name. This is typically a shorter version of an
	// attribute type's name.
	KeyName string `json:"keyName,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// Description - A brief description of a SoftLayer account attribute type.
	Description string `json:"description,omitempty"`

	// Id - A SoftLayer account attribute type's internal identifier.
	Id int `json:"id,omitempty"`
}

func (softlayer_account_attribute_type *SoftLayer_Account_Attribute_Type) String() string {
	return "SoftLayer_Account_Attribute_Type"
}
