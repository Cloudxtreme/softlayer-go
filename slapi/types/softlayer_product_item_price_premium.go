package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Product_Item_Price_Premium - <nil>
type SoftLayer_Product_Item_Price_Premium struct {

	// LocationId - <nil>
	LocationId int `json:"locationId,omitempty"`

	// MonthlyModifier - <nil>
	MonthlyModifier slapi.Float64 `json:"monthlyModifier,omitempty"`

	// PackageId - <nil>
	PackageId int `json:"packageId,omitempty"`

	// HourlyModifier - <nil>
	HourlyModifier slapi.Float64 `json:"hourlyModifier,omitempty"`

	// ItemPriceId - <nil>
	ItemPriceId int `json:"itemPriceId,omitempty"`
}

func (softlayer_product_item_price_premium *SoftLayer_Product_Item_Price_Premium) String() string {
	return "SoftLayer_Product_Item_Price_Premium"
}

// SoftLayer_Product_Item_Price_Premium_Extended is SoftLayer_Product_Item_Price_Premium with all maskable types.
type SoftLayer_Product_Item_Price_Premium_Extended struct {
	SoftLayer_Product_Item_Price_Premium

	// ItemPrice - <nil>
	ItemPrice *SoftLayer_Product_Item_Price `json:"itemPrice,omitempty"`

	// Location - <nil>
	Location *SoftLayer_Location `json:"location,omitempty"`

	// Package - <nil>
	Package *SoftLayer_Product_Package `json:"package,omitempty"`
}

func (softlayer_product_item_price_premium *SoftLayer_Product_Item_Price_Premium_Extended) String() string {
	return "SoftLayer_Product_Item_Price_Premium"
}
