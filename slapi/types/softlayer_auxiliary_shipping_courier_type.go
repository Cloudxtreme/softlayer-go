package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Auxiliary_Shipping_Courier_Type - <nil>
type SoftLayer_Auxiliary_Shipping_Courier_Type struct {

	// Courier - <nil>
	Courier []*SoftLayer_Auxiliary_Shipping_Courier `json:"courier"`

	// CourierCount - no documentation
	CourierCount uint64 `json:"courierCount"`

	// Description - <nil>
	Description string `json:"description"`

	// Id - <nil>
	Id int `json:"id"`

	// KeyName - <nil>
	KeyName string `json:"keyName"`

	// Name - <nil>
	Name string `json:"name"`
}

// GetObject - <nil>
func (softlayer_auxiliary_shipping_courier_type *SoftLayer_Auxiliary_Shipping_Courier_Type) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Auxiliary_Shipping_Courier_Type, error) {
	var returnValue *SoftLayer_Auxiliary_Shipping_Courier_Type
	return returnValue, nil
}

// GetTypeByKeyName - <nil>
func (softlayer_auxiliary_shipping_courier_type *SoftLayer_Auxiliary_Shipping_Courier_Type) GetTypeByKeyName(ctx *slapi.RequestContext, keyName string) (*SoftLayer_Auxiliary_Shipping_Courier_Type, error) {
	var returnValue *SoftLayer_Auxiliary_Shipping_Courier_Type
	return returnValue, nil
}
