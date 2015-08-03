package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Product_Catalog_Item_Price - The SoftLayer_Product_Catalog_Item_Price type assigns an Item
// Price to a Catalog. This relation defines the composition of Item Prices in a Catalog.
type SoftLayer_Product_Catalog_Item_Price struct {

	// PriceId - The id of the Item Price that is part of the Catalog.
	PriceId int `json:"priceId"`

	// CatalogId - no documentation
	CatalogId int `json:"catalogId"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// ModifyDate - The time the Item Price was changed for the Catalog
	ModifyDate *time.Time `json:"modifyDate"`
}

func (softlayer_product_catalog_item_price *SoftLayer_Product_Catalog_Item_Price) String() string {
	return "SoftLayer_Product_Catalog_Item_Price"
}

// SoftLayer_Product_Catalog_Item_Price_Extended is SoftLayer_Product_Catalog_Item_Price with all maskable types.
type SoftLayer_Product_Catalog_Item_Price_Extended struct {
	SoftLayer_Product_Catalog_Item_Price

	// Catalog - no documentation
	Catalog *SoftLayer_Product_Catalog `json:"catalog"`

	// Price - no documentation
	Price *SoftLayer_Product_Item_Price `json:"price"`
}

func (softlayer_product_catalog_item_price *SoftLayer_Product_Catalog_Item_Price_Extended) String() string {
	return "SoftLayer_Product_Catalog_Item_Price"
}
