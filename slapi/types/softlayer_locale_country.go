package types

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

func (softlayer_locale_country *SoftLayer_Locale_Country) String() string {
	return "SoftLayer_Locale_Country"
}
