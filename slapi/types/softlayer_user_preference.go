package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_User_Preference - The SoftLayer_User_Preference data type contains a single user
// preference to a specific preference type.
type SoftLayer_User_Preference struct {

	// Value - no documentation
	Value string `json:"value,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// Type - no documentation
	Type *SoftLayer_User_Preference_Type `json:"type,omitempty"`
}

func (softlayer_user_preference *SoftLayer_User_Preference) String() string {
	return "SoftLayer_User_Preference"
}
