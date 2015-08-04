package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Location_Group_Pricing - <nil>
type SoftLayer_Location_Group_Pricing struct {

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// LocationGroupTypeId - <nil>
	LocationGroupTypeId int `json:"locationGroupTypeId,omitempty"`

	// SecurityLevelId - <nil>
	SecurityLevelId int `json:"securityLevelId,omitempty"`

	// Description - <nil>
	Description string `json:"description,omitempty"`

	// Name - <nil>
	Name string `json:"name,omitempty"`

	// Locations - no documentation
	Locations []*SoftLayer_Location `json:"locations,omitempty"`

	// Prices - The prices that this pricing location group limits. All of these prices will only be
	// available in the locations defined by this pricing location group.
	Prices []*SoftLayer_Product_Item_Price `json:"prices,omitempty"`

	// PriceCount - A count of the prices that this pricing location group limits. All of these prices will
	// only be available in the locations defined by this pricing location group.
	PriceCount uint64 `json:"priceCount,omitempty"`

	// LocationCount - no documentation
	LocationCount uint64 `json:"locationCount,omitempty"`

	// LocationGroupType - no documentation
	LocationGroupType *SoftLayer_Location_Group_Type `json:"locationGroupType,omitempty"`
}

func (softlayer_location_group_pricing *SoftLayer_Location_Group_Pricing) String() string {
	return "SoftLayer_Location_Group_Pricing"
}
