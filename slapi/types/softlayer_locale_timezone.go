package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Locale_Timezone - Each User is assigned a timezone allowing for a precise local timestamp.
type SoftLayer_Locale_Timezone struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// LongName - A timezone's long name. For example, "(GMT-06:00) America/Dallas -
	LongName string `json:"longName,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// Offset - A timezone's offset based on the GMT standard. For example, Central Standard Time's offset
	// is "-0600" from GMT=0000.
	Offset string `json:"offset,omitempty"`

	// ShortName - A timezone's common abbreviation. For example, Central Standard Time's abbreviation is
	ShortName string `json:"shortName,omitempty"`
}

func (softlayer_locale_timezone *SoftLayer_Locale_Timezone) String() string {
	return "SoftLayer_Locale_Timezone"
}
