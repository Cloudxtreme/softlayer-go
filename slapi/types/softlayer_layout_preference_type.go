package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Layout_Preference_Type - The SoftLayer_Layout_Preference_Type contains definitions for
// preference types
type SoftLayer_Layout_Preference_Type struct {

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// ValueExpression - A regular expression used to validate the related [[SoftLayer_Layout_Preference]]
	ValueExpression string `json:"valueExpression,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Keyname - The unique key name of the item type, used primarily for programmatic purposes
	Keyname string `json:"keyname,omitempty"`
}

func (softlayer_layout_preference_type *SoftLayer_Layout_Preference_Type) String() string {
	return "SoftLayer_Layout_Preference_Type"
}
