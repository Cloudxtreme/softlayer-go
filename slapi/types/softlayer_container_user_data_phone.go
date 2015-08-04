package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_User_Data_Phone - no documentation
type SoftLayer_Container_User_Data_Phone struct {

	// Extension - Phone extension code. It can be digits, commas, *, and # are allowed.
	Extension string `json:"extension,omitempty"`

	// Phone - Phone number can be a mobile phone number, desk phone number, or some other option. The
	// phone number format must match the format selected in the country code.
	Phone string `json:"phone,omitempty"`

	// PhoneType - Type of phone number such as "primary", "office" or "home"
	PhoneType string `json:"phoneType,omitempty"`

	// CountryCode - Country code number for the phone number Default: 1 (United States & Canada +1)
	CountryCode int `json:"countryCode,omitempty"`
}

func (softlayer_container_user_data_phone *SoftLayer_Container_User_Data_Phone) String() string {
	return "SoftLayer_Container_User_Data_Phone"
}
