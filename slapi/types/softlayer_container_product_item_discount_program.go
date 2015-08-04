package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Item_Discount_Program - The
// SoftLayer_Container_Product_Item_Discount_Program data type represents the information about a
// discount that is related to a specific product item.
type SoftLayer_Container_Product_Item_Discount_Program struct {

	// OneTimeTax - The tax amount on the one time fees (one time, setup and labor) of the prices of this
	// container mulitiplied by the applicable quantity of this container.
	OneTimeTax float64 `json:"oneTimeTax,omitempty"`

	// ProratedOneTimeAmount - The sum of the one time fees (one time, setup and labor) of the prices of
	// this container multiplied by the applicable quantity of this container with the proration factor
	// applied.
	ProratedOneTimeAmount float64 `json:"proratedOneTimeAmount,omitempty"`

	// ProratedRecurringAmount - The sum of the recurring fees of the prices of this container multiplied
	// by the applicable quantity of this container with the proration factor applied.
	ProratedRecurringAmount float64 `json:"proratedRecurringAmount,omitempty"`

	// ProratedRecurringTax - The tax amount on the recurring fees of the prices of this container
	// mulitiplied by the applicable quantity of this container with the proration factor applied.
	ProratedRecurringTax float64 `json:"proratedRecurringTax,omitempty"`

	// RecurringTax - The tax amount on the recurring fees of the prices of this container mulitiplied by
	// the applicable quantity of this container.
	RecurringTax float64 `json:"recurringTax,omitempty"`

	// ApplicableQuantity - The number of times the item discount(s) may be applied for that order
	// container. At a minimum the number will be 1 and at most, it will match the quantity of the order
	// container.
	ApplicableQuantity int `json:"applicableQuantity,omitempty"`

	// Item - no documentation
	Item *SoftLayer_Product_Item `json:"item,omitempty"`

	// OneTimeAmount - The sum of the one time fees (one time, setup and labor) of the prices of this
	// container multiplied by the applicable quantity of this container.
	OneTimeAmount float64 `json:"oneTimeAmount,omitempty"`

	// Prices - The item prices that contain the amount of the discount in the recurringFee field. There
	// may be one or more prices.
	Prices []*SoftLayer_Product_Item_Price `json:"prices,omitempty"`

	// ProratedOneTimeTax - The tax amount on the one time fees (one time, setup and labor) of the prices
	// of this container mulitiplied by the applicable quantity of this container with the proration factor
	// applied.
	ProratedOneTimeTax float64 `json:"proratedOneTimeTax,omitempty"`

	// RecurringAmount - The sum of the recurring fees of the prices of this container multiplied by the
	// applicable quantity of this container.
	RecurringAmount float64 `json:"recurringAmount,omitempty"`
}

func (softlayer_container_product_item_discount_program *SoftLayer_Container_Product_Item_Discount_Program) String() string {
	return "SoftLayer_Container_Product_Item_Discount_Program"
}
