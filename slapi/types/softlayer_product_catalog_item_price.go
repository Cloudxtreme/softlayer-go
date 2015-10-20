package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Product_Catalog_Item_Price - The SoftLayer_Product_Catalog_Item_Price type assigns an Item
// Price to a Catalog. This relation defines the composition of Item Prices in a Catalog.
type SoftLayer_Product_Catalog_Item_Price struct {

	// CatalogId - no documentation
	CatalogId int `json:"catalogId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// ModifyDate - The time the Item Price was changed for the Catalog
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// PriceId - The id of the Item Price that is part of the Catalog.
	PriceId int `json:"priceId,omitempty"`

	// Catalog - no documentation
	Catalog *SoftLayer_Product_Catalog `json:"catalog,omitempty"`

	// Price - no documentation
	Price *SoftLayer_Product_Item_Price `json:"price,omitempty"`
}

func (softlayer_product_catalog_item_price *SoftLayer_Product_Catalog_Item_Price) String() string {
	return "SoftLayer_Product_Catalog_Item_Price"
}
