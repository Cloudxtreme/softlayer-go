package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Layout_Preference - The SoftLayer_Layout_Preference contains definitions for default
// layout item preferences
type SoftLayer_Layout_Preference struct {

	// Id - no documentation
	Id int `json:"id"`

	// LayoutPreferenceType - no documentation
	LayoutPreferenceType *SoftLayer_Layout_Preference_Type `json:"layoutPreferenceType"`

	// LayoutPreferenceTypeId - The internal identifier of the related [[SoftLayer_Layout_Preference_Type]]
	LayoutPreferenceTypeId int `json:"layoutPreferenceTypeId"`

	// Value - no documentation
	Value string `json:"value"`
}
