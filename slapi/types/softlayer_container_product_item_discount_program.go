package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Item_Discount_Program - The
// SoftLayer_Container_Product_Item_Discount_Program data type represents the information about a
// discount that is related to a specific product item.
type SoftLayer_Container_Product_Item_Discount_Program struct {

	// RecurringAmount - The sum of the recurring fees of the prices of this container multiplied by the
	// applicable quantity of this container.
	RecurringAmount float64 `json:"recurringAmount"`

	// RecurringTax - The tax amount on the recurring fees of the prices of this container mulitiplied by
	// the applicable quantity of this container.
	RecurringTax float64 `json:"recurringTax"`

	// Item - no documentation
	Item *SoftLayer_Product_Item `json:"item"`

	// Prices - The item prices that contain the amount of the discount in the recurringFee field. There
	// may be one or more prices.
	Prices []*SoftLayer_Product_Item_Price `json:"prices"`

	// ProratedOneTimeAmount - The sum of the one time fees (one time, setup and labor) of the prices of
	// this container multiplied by the applicable quantity of this container with the proration factor
	// applied.
	ProratedOneTimeAmount float64 `json:"proratedOneTimeAmount"`

	// ProratedOneTimeTax - The tax amount on the one time fees (one time, setup and labor) of the prices
	// of this container mulitiplied by the applicable quantity of this container with the proration factor
	// applied.
	ProratedOneTimeTax float64 `json:"proratedOneTimeTax"`

	// ProratedRecurringAmount - The sum of the recurring fees of the prices of this container multiplied
	// by the applicable quantity of this container with the proration factor applied.
	ProratedRecurringAmount float64 `json:"proratedRecurringAmount"`

	// ProratedRecurringTax - The tax amount on the recurring fees of the prices of this container
	// mulitiplied by the applicable quantity of this container with the proration factor applied.
	ProratedRecurringTax float64 `json:"proratedRecurringTax"`

	// ApplicableQuantity - The number of times the item discount(s) may be applied for that order
	// container. At a minimum the number will be 1 and at most, it will match the quantity of the order
	// container.
	ApplicableQuantity int `json:"applicableQuantity"`

	// OneTimeAmount - The sum of the one time fees (one time, setup and labor) of the prices of this
	// container multiplied by the applicable quantity of this container.
	OneTimeAmount float64 `json:"oneTimeAmount"`

	// OneTimeTax - The tax amount on the one time fees (one time, setup and labor) of the prices of this
	// container mulitiplied by the applicable quantity of this container.
	OneTimeTax float64 `json:"oneTimeTax"`
}

func (softlayer_container_product_item_discount_program *SoftLayer_Container_Product_Item_Discount_Program) String() string {
	return "SoftLayer_Container_Product_Item_Discount_Program"
}
