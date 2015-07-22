package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Locale_Timezone - Each User is assigned a timezone allowing for a precise local timestamp.
type SoftLayer_Locale_Timezone struct {

	// Id - no documentation
	Id int `json:"id"`

	// LongName - A timezone's long name. For example, "(GMT-06:00) America/Dallas -
	LongName string `json:"longName"`

	// Name - no documentation
	Name string `json:"name"`

	// Offset - A timezone's offset based on the GMT standard. For example, Central Standard Time's offset
	// is "-0600" from GMT=0000.
	Offset string `json:"offset"`

	// ShortName - A timezone's common abbreviation. For example, Central Standard Time's abbreviation is
	ShortName string `json:"shortName"`
}

// GetAllObjects - no documentation
func (softlayer_locale_timezone *SoftLayer_Locale_Timezone) GetAllObjects(commonOptions *slapi.CommonOptions) ([]*SoftLayer_Locale_Timezone, error) {
	var returnValue []*SoftLayer_Locale_Timezone
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Locale_Timezone object whose ID number corresponds to
// the ID number of the init parameter passed to the SoftLayer_Locale_Timezone service.
func (softlayer_locale_timezone *SoftLayer_Locale_Timezone) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Locale_Timezone, error) {
	var returnValue *SoftLayer_Locale_Timezone
	return returnValue, nil
}
