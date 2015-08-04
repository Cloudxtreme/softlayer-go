package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Layout_Preference - The SoftLayer_Layout_Preference contains definitions for default
// layout item preferences
type SoftLayer_Layout_Preference struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// LayoutPreferenceTypeId - The internal identifier of the related [[SoftLayer_Layout_Preference_Type]]
	LayoutPreferenceTypeId int `json:"layoutPreferenceTypeId,omitempty"`

	// Value - no documentation
	Value string `json:"value,omitempty"`

	// LayoutPreferenceType - no documentation
	LayoutPreferenceType *SoftLayer_Layout_Preference_Type `json:"layoutPreferenceType,omitempty"`
}

func (softlayer_layout_preference *SoftLayer_Layout_Preference) String() string {
	return "SoftLayer_Layout_Preference"
}
