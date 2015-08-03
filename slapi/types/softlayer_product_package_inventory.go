package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Product_Package_Inventory - SoftLayer keeps near real-time track of the number of items
// available in it's product catalog inventory. The SoftLayer_Product_Package_Inventory data type
// models one of these inventory records. SoftLayer tracks inventory per product package and item per
// datacenter. This type is useful if you need to purchase specific servers in a specific location, and
// wish to check their availability before ordering. The data from this type is used primarily on the
// SoftLayer outlet website.
type SoftLayer_Product_Package_Inventory struct {

	// ModifyDate - The date that an inventory record was last updated.
	ModifyDate *time.Time `json:"modifyDate"`

	// ItemId - The unique identifier of the product item that an inventory record is associated with.
	ItemId int `json:"itemId"`

	// LocationId - The unique identifier of the datacenter that an inventory record is located in.
	LocationId int `json:"locationId"`

	// OverstockFlag - Whether an inventory record is marked as "overstock". Overstock records appear at
	// the top portion of the SoftLayer outlet website.
	OverstockFlag int `json:"overstockFlag"`

	// PackageId - The unique identifier of the product package that an inventory record is associated
	// with.
	PackageId int `json:"packageId"`

	// AvailableInventoryCount - The number of units available for purchase in SoftLayer's inventory for a
	// single item in a single datacenter.
	AvailableInventoryCount int `json:"availableInventoryCount"`
}

// SoftLayer_Product_Package_Inventory_Extended is SoftLayer_Product_Package_Inventory with all maskable types.
type SoftLayer_Product_Package_Inventory_Extended struct {
	SoftLayer_Product_Package_Inventory

	// Item - The product package item that is associated with an inventory record.
	Item *SoftLayer_Product_Item `json:"item"`

	// Location - The datacenter that an inventory record is located in.
	Location *SoftLayer_Location `json:"location"`

	// Package - The product package that is associated with an inventory record.
	Package *SoftLayer_Product_Package `json:"package"`
}

func (softlayer_product_package_inventory *SoftLayer_Product_Package_Inventory) String() string {
	return "SoftLayer_Product_Package_Inventory"
}
