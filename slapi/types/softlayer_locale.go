package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Locale - Every SoftLayer_User_Customer is associated with a locale. SoftLayer_Locale holds
// user's language and region information.
type SoftLayer_Locale struct {

	// LanguageTag - <nil>
	LanguageTag string `json:"languageTag"`

	// Name - no documentation
	Name string `json:"name"`

	// FriendlyName - <nil>
	FriendlyName string `json:"friendlyName"`

	// Id - no documentation
	Id int `json:"id"`
}

func (softlayer_locale *SoftLayer_Locale) String() string {
	return "SoftLayer_Locale"
}
