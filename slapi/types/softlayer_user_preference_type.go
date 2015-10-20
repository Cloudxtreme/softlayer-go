package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_User_Preference_Type - The SoftLayer_User_Preference_Type data type contains a single
// preference type including the accepted values.
type SoftLayer_User_Preference_Type struct {

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// ValueExample - no documentation
	ValueExample string `json:"valueExample,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`
}

func (softlayer_user_preference_type *SoftLayer_User_Preference_Type) String() string {
	return "SoftLayer_User_Preference_Type"
}
