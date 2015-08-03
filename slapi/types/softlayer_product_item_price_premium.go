package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Price_Premium - <nil>
type SoftLayer_Product_Item_Price_Premium struct {

	// HourlyModifier - <nil>
	HourlyModifier float64 `json:"hourlyModifier"`

	// ItemPriceId - <nil>
	ItemPriceId int `json:"itemPriceId"`

	// LocationId - <nil>
	LocationId int `json:"locationId"`

	// MonthlyModifier - <nil>
	MonthlyModifier float64 `json:"monthlyModifier"`

	// PackageId - <nil>
	PackageId int `json:"packageId"`
}

// SoftLayer_Product_Item_Price_Premium_Extended is SoftLayer_Product_Item_Price_Premium with all maskable types.
type SoftLayer_Product_Item_Price_Premium_Extended struct {
	SoftLayer_Product_Item_Price_Premium

	// ItemPrice - <nil>
	ItemPrice *SoftLayer_Product_Item_Price `json:"itemPrice"`

	// Location - <nil>
	Location *SoftLayer_Location `json:"location"`

	// Package - <nil>
	Package *SoftLayer_Product_Package `json:"package"`
}

func (softlayer_product_item_price_premium *SoftLayer_Product_Item_Price_Premium) String() string {
	return "SoftLayer_Product_Item_Price_Premium"
}
