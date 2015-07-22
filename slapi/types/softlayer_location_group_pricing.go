package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Location_Group_Pricing - <nil>
type SoftLayer_Location_Group_Pricing struct {

	// PriceCount - no documentation
	PriceCount uint64 `json:"priceCount"`

	// Prices - <nil>
	Prices []*SoftLayer_Product_Item_Price `json:"prices"`
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
