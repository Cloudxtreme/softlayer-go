package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Locale_Timezone - Each User is assigned a timezone allowing for a precise local timestamp.
type SoftLayer_Locale_Timezone struct {

	// ShortName - A timezone's common abbreviation. For example, Central Standard Time's abbreviation is
	ShortName string `json:"shortName"`

	// Id - no documentation
	Id int `json:"id"`

	// LongName - A timezone's long name. For example, "(GMT-06:00) America/Dallas -
	LongName string `json:"longName"`

	// Name - no documentation
	Name string `json:"name"`

	// Offset - A timezone's offset based on the GMT standard. For example, Central Standard Time's offset
	// is "-0600" from GMT=0000.
	Offset string `json:"offset"`
}

func (softlayer_locale_timezone *SoftLayer_Locale_Timezone) String() string {
	return "SoftLayer_Locale_Timezone"
}
