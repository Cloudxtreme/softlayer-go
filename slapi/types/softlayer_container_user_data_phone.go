package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_User_Data_Phone - no documentation
type SoftLayer_Container_User_Data_Phone struct {

	// CountryCode - Country code number for the phone number Default: 1 (United States & Canada +1)
	CountryCode int `json:"countryCode"`

	// Extension - Phone extension code. It can be digits, commas, *, and # are allowed.
	Extension string `json:"extension"`

	// Phone - Phone number can be a mobile phone number, desk phone number, or some other option. The
	// phone number format must match the format selected in the country code.
	Phone string `json:"phone"`

	// PhoneType - Type of phone number such as "primary", "office" or "home"
	PhoneType string `json:"phoneType"`
}
