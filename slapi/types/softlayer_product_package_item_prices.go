package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Package_Item_Prices - The SoftLayer_Product_Package_Item_Prices contains price to
// package cross references Relates a category, price and item to a bundle. Match bundle ids to see all
// items and prices in a particular bundle.
type SoftLayer_Product_Package_Item_Prices struct {

	// Id - The unique identifier for SoftLayer_Product_Package_Item_Price. This is only needed as a
	// reference. The important data is the itemPriceId property.
	Id int `json:"id"`

	// ItemPriceId - The SoftLayer_Product_Item_Price id. This value is to be used when placing orders. To
	// get more information about this item price, go from the item price to the item description
	ItemPriceId int `json:"itemPriceId"`

	// PackageId - The Package ID to which this price reference belongs
	PackageId int `json:"packageId"`
}

func (softlayer_product_package_item_prices *SoftLayer_Product_Package_Item_Prices) String() string {
	return "SoftLayer_Product_Package_Item_Prices"
}

// SoftLayer_Product_Package_Item_Prices_Extended is SoftLayer_Product_Package_Item_Prices with all maskable types.
type SoftLayer_Product_Package_Item_Prices_Extended struct {
	SoftLayer_Product_Package_Item_Prices

	// ItemPrice - The item price to which this object belongs. The item price has details regarding cost
	// for the item it belongs to.
	ItemPrice *SoftLayer_Product_Item_Price `json:"itemPrice"`

	// Package - no documentation
	Package *SoftLayer_Product_Package `json:"package"`
}

func (softlayer_product_package_item_prices *SoftLayer_Product_Package_Item_Prices_Extended) String() string {
	return "SoftLayer_Product_Package_Item_Prices"
}
