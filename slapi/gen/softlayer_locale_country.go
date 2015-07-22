package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Locale_Country - <nil>
type SoftLayer_Locale_Country struct {

	// LongName - <nil>
	LongName string `json:"longName"`

	// ShortName - <nil>
	ShortName string `json:"shortName"`

	// States - <nil>
	States []*SoftLayer_Locale_StateProvince `json:"states"`
}

// GetAvailableCountries - Use this method to retrieve a list of countries and locale information
// available to the current user.
func (softlayer_locale_country *SoftLayer_Locale_Country) GetAvailableCountries() ([]*SoftLayer_Locale_Country, error) {
	var returnValue []*SoftLayer_Locale_Country
	return returnValue, nil
}

// GetCountries - Use this method to retrieve a list of countries and locale information such as
// country code and state/provinces.
func (softlayer_locale_country *SoftLayer_Locale_Country) GetCountries() ([]*SoftLayer_Locale_Country, error) {
	var returnValue []*SoftLayer_Locale_Country
	return returnValue, nil
}
