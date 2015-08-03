package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Price_Premium - <nil>
type SoftLayer_Product_Item_Price_Premium struct {

	// HourlyModifier - <nil>
	HourlyModifier float64 `json:"hourlyModifier"`

	// ItemPrice - <nil>
	ItemPrice *SoftLayer_Product_Item_Price `json:"itemPrice"`

	// ItemPriceId - <nil>
	ItemPriceId int `json:"itemPriceId"`

	// Location - <nil>
	Location *SoftLayer_Location `json:"location"`

	// LocationId - <nil>
	LocationId int `json:"locationId"`

	// MonthlyModifier - <nil>
	MonthlyModifier float64 `json:"monthlyModifier"`

	// Package - <nil>
	Package *SoftLayer_Product_Package `json:"package"`

	// PackageId - <nil>
	PackageId int `json:"packageId"`
}

func (softlayer_product_item_price_premium *SoftLayer_Product_Item_Price_Premium) String() string {
	return "SoftLayer_Product_Item_Price_Premium"
}
