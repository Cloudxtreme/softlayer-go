package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Location_Group_Pricing - <nil>
type SoftLayer_Location_Group_Pricing struct {

	// PriceCount - A count of the prices that this pricing location group limits. All of these prices will
	// only be available in the locations defined by this pricing location group.
	PriceCount uint64 `json:"priceCount"`

	// Prices - The prices that this pricing location group limits. All of these prices will only be
	// available in the locations defined by this pricing location group.
	Prices []*SoftLayer_Product_Item_Price `json:"prices"`
}

func (softlayer_location_group_pricing *SoftLayer_Location_Group_Pricing) String() string {
	return "SoftLayer_Location_Group_Pricing"
}

// GetAllObjects - <nil>
func (softlayer_location_group_pricing *SoftLayer_Location_Group_Pricing) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_Location_Group, error) {
	var returnValue []*SoftLayer_Location_Group
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_location_group_pricing *SoftLayer_Location_Group_Pricing) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Location_Group_Pricing, error) {
	var returnValue *SoftLayer_Location_Group_Pricing
	return returnValue, nil
}
