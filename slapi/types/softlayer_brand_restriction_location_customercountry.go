package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Brand_Restriction_Location_CustomerCountry - The
// [[SoftLayer_Brand_Restriction_Location_CustomerCountry]] data type defines the relationship between
// brands, locations and countries associated with a user's account that are ineligible when ordering
// products. For example, the India datacenter may not be available on the SoftLayer US brand for
// customers that live in Great Britain.
type SoftLayer_Brand_Restriction_Location_CustomerCountry struct {

	// BrandId - no documentation
	BrandId int `json:"brandId,omitempty"`

	// CustomerCountryCode - no documentation
	CustomerCountryCode string `json:"customerCountryCode,omitempty"`

	// LocationId - no documentation
	LocationId int `json:"locationId,omitempty"`

	// Brand - This references the brand that has a brand-location-country restriction setup.
	Brand *SoftLayer_Brand `json:"brand,omitempty"`

	// Location - This references the datacenter that has a brand-location-country restriction setup. For
	// example, if a datacenter is listed with a restriction for Canada, a Canadian customer may not be
	// eligible to order services at that location.
	Location *SoftLayer_Location `json:"location,omitempty"`
}

func (softlayer_brand_restriction_location_customercountry *SoftLayer_Brand_Restriction_Location_CustomerCountry) String() string {
	return "SoftLayer_Brand_Restriction_Location_CustomerCountry"
}
