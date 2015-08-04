package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Container_Product_Item_Discount_Program - The
// SoftLayer_Container_Product_Item_Discount_Program data type represents the information about a
// discount that is related to a specific product item.
type SoftLayer_Container_Product_Item_Discount_Program struct {

	// ProratedOneTimeAmount - The sum of the one time fees (one time, setup and labor) of the prices of
	// this container multiplied by the applicable quantity of this container with the proration factor
	// applied.
	ProratedOneTimeAmount slapi.Float64 `json:"proratedOneTimeAmount,omitempty"`

	// ProratedRecurringTax - The tax amount on the recurring fees of the prices of this container
	// mulitiplied by the applicable quantity of this container with the proration factor applied.
	ProratedRecurringTax slapi.Float64 `json:"proratedRecurringTax,omitempty"`

	// RecurringAmount - The sum of the recurring fees of the prices of this container multiplied by the
	// applicable quantity of this container.
	RecurringAmount slapi.Float64 `json:"recurringAmount,omitempty"`

	// OneTimeTax - The tax amount on the one time fees (one time, setup and labor) of the prices of this
	// container mulitiplied by the applicable quantity of this container.
	OneTimeTax slapi.Float64 `json:"oneTimeTax,omitempty"`

	// Prices - The item prices that contain the amount of the discount in the recurringFee field. There
	// may be one or more prices.
	Prices []*SoftLayer_Product_Item_Price `json:"prices,omitempty"`

	// ProratedOneTimeTax - The tax amount on the one time fees (one time, setup and labor) of the prices
	// of this container mulitiplied by the applicable quantity of this container with the proration factor
	// applied.
	ProratedOneTimeTax slapi.Float64 `json:"proratedOneTimeTax,omitempty"`

	// ProratedRecurringAmount - The sum of the recurring fees of the prices of this container multiplied
	// by the applicable quantity of this container with the proration factor applied.
	ProratedRecurringAmount slapi.Float64 `json:"proratedRecurringAmount,omitempty"`

	// RecurringTax - The tax amount on the recurring fees of the prices of this container mulitiplied by
	// the applicable quantity of this container.
	RecurringTax slapi.Float64 `json:"recurringTax,omitempty"`

	// ApplicableQuantity - The number of times the item discount(s) may be applied for that order
	// container. At a minimum the number will be 1 and at most, it will match the quantity of the order
	// container.
	ApplicableQuantity int `json:"applicableQuantity,omitempty"`

	// Item - no documentation
	Item *SoftLayer_Product_Item `json:"item,omitempty"`

	// OneTimeAmount - The sum of the one time fees (one time, setup and labor) of the prices of this
	// container multiplied by the applicable quantity of this container.
	OneTimeAmount slapi.Float64 `json:"oneTimeAmount,omitempty"`
}

func (softlayer_container_product_item_discount_program *SoftLayer_Container_Product_Item_Discount_Program) String() string {
	return "SoftLayer_Container_Product_Item_Discount_Program"
}
