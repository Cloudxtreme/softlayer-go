package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Layout_Preference - The SoftLayer_Layout_Preference contains definitions for default
// layout item preferences
type SoftLayer_Layout_Preference struct {

	// Id - no documentation
	Id int `json:"id"`

	// LayoutPreferenceTypeId - The internal identifier of the related [[SoftLayer_Layout_Preference_Type]]
	LayoutPreferenceTypeId int `json:"layoutPreferenceTypeId"`

	// Value - no documentation
	Value string `json:"value"`
}

func (softlayer_layout_preference *SoftLayer_Layout_Preference) String() string {
	return "SoftLayer_Layout_Preference"
}

// SoftLayer_Layout_Preference_Extended is SoftLayer_Layout_Preference with all maskable types.
type SoftLayer_Layout_Preference_Extended struct {
	SoftLayer_Layout_Preference

	// LayoutPreferenceType - no documentation
	LayoutPreferenceType *SoftLayer_Layout_Preference_Type `json:"layoutPreferenceType"`
}

func (softlayer_layout_preference *SoftLayer_Layout_Preference_Extended) String() string {
	return "SoftLayer_Layout_Preference"
}
