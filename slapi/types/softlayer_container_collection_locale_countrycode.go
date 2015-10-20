package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Collection_Locale_CountryCode - This container is used to hold country locale
// information.
type SoftLayer_Container_Collection_Locale_CountryCode struct {

	// ShortName - <nil>
	ShortName string `json:"shortName,omitempty"`

	// StateCodes - <nil>
	StateCodes []*SoftLayer_Container_Collection_Locale_StateCode `json:"stateCodes,omitempty"`

	// LongName - <nil>
	LongName string `json:"longName,omitempty"`
}

func (softlayer_container_collection_locale_countrycode *SoftLayer_Container_Collection_Locale_CountryCode) String() string {
	return "SoftLayer_Container_Collection_Locale_CountryCode"
}
