package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Locale_Country - <nil>
type SoftLayer_Locale_Country struct {

	// LongName - <nil>
	LongName string `json:"longName,omitempty"`

	// ShortName - <nil>
	ShortName string `json:"shortName,omitempty"`

	// States - no documentation
	States []*SoftLayer_Locale_StateProvince `json:"states,omitempty"`

	// StateCount - no documentation
	StateCount uint64 `json:"stateCount,omitempty"`
}

func (softlayer_locale_country *SoftLayer_Locale_Country) String() string {
	return "SoftLayer_Locale_Country"
}
