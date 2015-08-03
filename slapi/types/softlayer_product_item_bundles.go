package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Bundles - The SoftLayer_Product_Item_Bundles contains item to price cross
// references Relates a category, price and item to a bundle. Match bundle ids to see all items and
// prices in a particular bundle.
type SoftLayer_Product_Item_Bundles struct {

	// Id - no documentation
	Id int `json:"id"`

	// ItemPriceId - no documentation
	ItemPriceId int `json:"itemPriceId"`

	// BundleItemId - no documentation
	BundleItemId int `json:"bundleItemId"`
}

// SoftLayer_Product_Item_Bundles_Extended is SoftLayer_Product_Item_Bundles with all maskable types.
type SoftLayer_Product_Item_Bundles_Extended struct {
	SoftLayer_Product_Item_Bundles

	// Category - no documentation
	Category *SoftLayer_Product_Item_Category `json:"category"`

	// ItemPrice - no documentation
	ItemPrice *SoftLayer_Product_Item_Price `json:"itemPrice"`

	// BundleItem - no documentation
	BundleItem *SoftLayer_Product_Item `json:"bundleItem"`
}

func (softlayer_product_item_bundles *SoftLayer_Product_Item_Bundles) String() string {
	return "SoftLayer_Product_Item_Bundles"
}
