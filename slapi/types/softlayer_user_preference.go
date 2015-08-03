package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_User_Preference - The SoftLayer_User_Preference data type contains a single user
// preference to a specific preference type.
type SoftLayer_User_Preference struct {

	// Value - no documentation
	Value string `json:"value"`
}

// SoftLayer_User_Preference_Extended is SoftLayer_User_Preference with all maskable types.
type SoftLayer_User_Preference_Extended struct {
	SoftLayer_User_Preference

	// Description - no documentation
	Description string `json:"description"`

	// Type - no documentation
	Type *SoftLayer_User_Preference_Type `json:"type"`
}

func (softlayer_user_preference *SoftLayer_User_Preference) String() string {
	return "SoftLayer_User_Preference"
}
